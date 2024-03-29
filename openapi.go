package goapi

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/iancoleman/strcase"
	ff "github.com/ysmood/goapi/lib/flat-fields"
	"github.com/ysmood/goapi/lib/openapi"
	"github.com/ysmood/jschema"
	"github.com/ysmood/vary"
)

// ConfigOpenAPI is a function to modify the generated OpenAPI doc.
type ConfigOpenAPI func(doc *openapi.Operation)

// Interfaces is the global interface set.
var Interfaces = vary.NewInterfaces()

// Interface create a interface set of i. ts are the types that implement i.
// For golang runtime we can't reflect all the implementations of an interface,
// with it goapi can find out all the possible response type of an endpoint.
func Interface(i any, ts ...any) *vary.Interface {
	return Interfaces.New(i, ts...)
}

// AddInterfaces to the global interface set.
func AddInterfaces(is vary.Interfaces) {
	for k, v := range is {
		Interfaces[k] = v
	}
}

// Descriptioner is an interface that is use to specify the description in openapi.
type Descriptioner interface {
	Description() string
}

var tDescriptioner = reflect.TypeOf((*Descriptioner)(nil)).Elem()

// ContentTyper is an interface that is use to specify the http Content-Type in openapi.
type ContentTyper interface {
	ContentType() string
}

var tContentTyper = reflect.TypeOf((*ContentTyper)(nil)).Elem()

// OpenAPI is a shortcut for [Router.OpenAPI].
func (g *Group) OpenAPI() *openapi.Document {
	return g.router.OpenAPI()
}

// OpenAPI returns the OpenAPI doc of the router.
// You can use [json.Marshal] to convert it to a JSON string.
func (r *Router) OpenAPI() *openapi.Document {
	doc := &openapi.Document{
		Paths: map[string]openapi.Path{},
	}

	for _, op := range r.operations {
		if op.override != nil {
			continue
		}

		if _, has := doc.Paths[op.path.path]; !has {
			doc.Paths[op.path.path] = openapi.Path{}
		}

		doc.Paths[op.path.path][op.method] = operationDoc(r.Schemas, op)
	}

	doc.Components.Schemas = r.Schemas.JSON()

	return doc
}

// OpenAPI sets the config function to modify the generated OpenAPI doc.
func (op *Operation) OpenAPI(config ConfigOpenAPI) {
	op.configOpenAPI = config
}

func operationDoc(s jschema.Schemas, op *Operation) openapi.Operation {
	doc := openapi.Operation{
		OperationID: op.name,
		Parameters:  []openapi.Parameter{},
		Responses:   map[openapi.StatusCode]openapi.Response{},
	}

	for _, p := range op.params {
		var params []openapi.Parameter

		switch p.in {
		case inHeader:
			params = append(params, headerParamDoc(s, p)...)

		case inURL:
			params = append(params, urlParamDoc(s, p)...)

		case inBody:
			doc.RequestBody = &openapi.RequestBody{
				Content: &openapi.Content{
					getContentType(p.param, openapi.ContentTypeJSON): &openapi.Schema{
						Schema: s.DefineT(p.param),
					},
				},
				Required: true,
			}
		}

		doc.Parameters = append(doc.Parameters, params...)
	}

	doc.Responses = resDoc(s, op)

	if op.configOpenAPI != nil {
		op.configOpenAPI(&doc)
	}

	return doc
}

func urlParamDoc(s jschema.Schemas, p *parsedParam) []openapi.Parameter {
	arr := []openapi.Parameter{}

	for _, f := range p.fields {
		in := openapi.QUERY

		if f.InPath {
			in = openapi.PATH
		}

		schema := fieldSchema(s, f.flatField.Field)
		desc := schema.Description
		schema.Description = ""
		examples := map[string]openapi.Example{}

		if len(schema.Examples) > 0 {
			for i, e := range schema.Examples {
				b, _ := json.Marshal(e)
				k := strconv.Itoa(i)

				examples[k] = openapi.Example{
					Summary: string(b),
					Value:   e,
				}
			}
		}

		arr = append(arr, openapi.Parameter{
			Name:        f.name,
			In:          in,
			Schema:      schema,
			Description: desc,
			Required:    f.required,
			Examples:    examples,
		})
	}

	return arr
}

func headerParamDoc(s jschema.Schemas, p *parsedParam) []openapi.Parameter {
	arr := []openapi.Parameter{}

	for _, f := range p.fields {
		schema := fieldSchema(s, f.flatField.Field)
		desc := schema.Description
		schema.Description = ""

		arr = append(arr, openapi.Parameter{
			Name:        f.name,
			In:          openapi.HEADER,
			Schema:      schema,
			Description: desc,
			Required:    f.required,
		})
	}

	return arr
}

func fieldSchema(s jschema.Schemas, f reflect.StructField) *jschema.Schema {
	if f.Type.Kind() == reflect.Ptr {
		f.Type = f.Type.Elem()
	}

	scm := s.DefineFieldT(f)
	scm = firstProp(scm)

	return scm
}

func resDoc(s jschema.Schemas, op *Operation) map[openapi.StatusCode]openapi.Response {
	list := map[openapi.StatusCode]openapi.Response{}

	add := func(t reflect.Type) {
		parsedRes := op.parseResponse(t)

		var content *openapi.Content

		if parsedRes.isStream { //nolint: gocritic,nestif
			content = &openapi.Content{
				getContentType(t, openapi.ContentTypeBin): &openapi.Schema{
					Schema: &jschema.Schema{
						Type:   jschema.TypeString,
						Format: "binary",
					},
				},
			}
		} else if parsedRes.isDirect {
			content = &openapi.Content{
				getContentType(t, openapi.ContentTypeJSON): &openapi.Schema{
					Schema: s.DefineT(parsedRes.data),
				},
			}
		} else if parsedRes.hasData || parsedRes.hasErr {
			scm := &jschema.Schema{
				Type:                 jschema.TypeObject,
				AdditionalProperties: ptr(false),
				Properties:           jschema.Properties{},
			}

			if parsedRes.hasErr { //nolint: gocritic
				scm.Properties["error"] = s.DefineT(parsedRes.err)
				scm.Required = []string{"error"}
			} else if parsedRes.hasMeta {
				scm.Properties["data"] = s.DefineT(parsedRes.data)
				scm.Properties["meta"] = s.DefineT(parsedRes.meta)
				scm.Required = []string{"data", "meta"}
			} else {
				scm.Properties["data"] = s.DefineT(parsedRes.data)
				scm.Required = []string{"data"}
			}

			content = &openapi.Content{
				getContentType(t, openapi.ContentTypeJSON): &openapi.Schema{
					Schema: scm,
				},
			}
		}

		code := openapi.StatusCode(parsedRes.statusCode)

		res := openapi.Response{
			Description: getDescription(t, code),
			Headers:     op.group.resHeaderDoc(s, parsedRes.header),
			Content:     content,
		}

		list[code] = res
	}

	if it, has := Interfaces[vary.ID(op.tRes)]; has {
		for _, t := range it.Implementations {
			add(t)
		}
	} else {
		add(op.tRes)
	}

	return list
}

func (g *Group) resHeaderDoc(s jschema.Schemas, t reflect.Type) openapi.Headers {
	if t == nil {
		return nil
	}

	headers := openapi.Headers{}

	for _, flat := range ff.Parse(t).Fields {
		f := parseHeaderField(g.router.Schemas, flat)
		headers[f.name] = openapi.Header{
			Description: f.schema.Description,
			Schema:      s.DefineT(f.item),
		}
	}

	return headers
}

func getDescription(t reflect.Type, code openapi.StatusCode) string {
	if t.Implements(tDescriptioner) {
		return reflect.New(t).Elem().Interface().(Descriptioner).Description()
	}

	return http.StatusText(int(code))
}

func getContentType(t reflect.Type, defaultType string) string {
	if t.Implements(tContentTyper) {
		return reflect.New(t).Elem().Interface().(ContentTyper).ContentType()
	}

	return defaultType
}

func ptr[T any](v T) *T {
	return &v
}

func toOperationName(name string) string {
	return strcase.ToLowerCamel(name)
}

func toHeaderName(name string) string {
	return strcase.ToKebab(name)
}

func toPathName(name string) string {
	return strcase.ToKebab(name)
}

func toQueryName(name string) string {
	return strcase.ToSnake(name)
}
