package goapi_test

import (
	"bytes"
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/ysmood/goapi"
	"github.com/ysmood/goapi/lib/middlewares"
	"github.com/ysmood/goapi/lib/middlewares/calm"
	"github.com/ysmood/goapi/lib/openapi"
	"github.com/ysmood/got"
)

type res interface {
	goapi.Response
}

var ires = goapi.Interface(new(res))

type resOK struct {
	goapi.StatusOK
	Data string
}

var _ = ires.Add(resOK{})

type resMeta struct {
	goapi.StatusOK
	Data string
	Meta string
}

type resErr struct {
	goapi.StatusBadRequest
	Error openapi.Error
}

type resHeader struct {
	goapi.StatusOK
	Data   int
	Header struct {
		X_UA string
	}
}

type resEncErr struct {
	goapi.StatusOK
	Data func()
}

type resEmpty struct {
	goapi.StatusOK
}

type ParamsValidate struct {
	goapi.InURL
}

type resForgetCreateInterface interface{ goapi.Response }

type resBin struct {
	goapi.StatusOK
	Data goapi.DataStream
}

type resDirect struct {
	goapi.StatusOK
	Data string `response:"direct"`
}

type resImage struct {
	goapi.StatusOK
	Header struct {
		ContentType string
	}
	Data goapi.DataStream
}

func (resImage) ContentType() string {
	return "image/*"
}

func TestOperation(t *testing.T) {
	g := got.T(t)
	tr := g.Serve()
	r := goapi.New()

	r.Use(&calm.Calm{
		PrintStack: false,
	})

	{ // setup
		r.GET("/query", func(params struct {
			goapi.InURL
			A string
		},
		) res {
			return resOK{Data: params.A}
		})

		r.GET("/meta", func(params struct {
			goapi.InHeader
			A string `json:"x"`
		},
		) resMeta {
			return resMeta{Data: params.A, Meta: params.A}
		})

		r.Use(middlewares.Func(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := context.WithValue(r.Context(), "key", "ok") //nolint: staticcheck
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		}))

		r.GET("/context", func(c context.Context) resOK {
			return resOK{Data: c.Value("key").(string)}
		})

		r.GET("/request", func(r *http.Request) resOK {
			return resOK{Data: r.URL.Query().Get("a")}
		})

		r.GET("/params-time/{t}", func(params struct {
			goapi.InURL
			T time.Time
		},
		) res {
			return resOK{Data: params.T.String()}
		})

		r.POST("/req-body", func(params struct {
			A string
		},
		) res {
			return resOK{Data: params.A}
		})

		r.GET("/error-res", func() resErr {
			return resErr{
				Error: openapi.Error{
					Code: openapi.CodeInvalidParam,
				},
			}
		})

		r.GET("/override-res", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotModified)
		})

		r.GET("/override-header", func() resHeader {
			return resHeader{
				Header: struct{ X_UA string }{X_UA: "test-client"},
			}
		})

		r.GET("/res-enc-err", func() resEncErr {
			return resEncErr{
				Data: func() {},
			}
		})

		r.GET("/res-empty", func() resEmpty {
			return resEmpty{}
		})

		r.GET("/res-missed-type", func() res {
			return resEmpty{}
		})

		r.GET("/validate", func(params ParamsValidate) goapi.StatusOK { return goapi.StatusOK{} })

		r.GET("/forget-create-interface", func() resForgetCreateInterface {
			return struct{ goapi.StatusOK }{}
		})

		r.GET("/res-bin", func() resBin {
			return resBin{
				Data: bytes.NewBufferString("ok"),
			}
		})

		r.GET("/res-direct", func() resDirect {
			return resDirect{
				Data: "ok",
			}
		})

		r.GET("/res-image", func() resImage {
			return resImage{
				Header: struct{ ContentType string }{ContentType: "image/png"},
				Data:   g.Open(false, "LICENSE"),
			}
		})

		tr.Mux.Handle("/", r.Server())
	}

	g.Eq(g.Req("", tr.URL("/query?a=ok")).JSON(), map[string]any{"data": "ok"})

	g.Eq(g.Req("", tr.URL("/query")).StatusCode, http.StatusBadRequest)

	g.Eq(g.Req("", tr.URL("/meta"), http.Header{"x": {"ok"}}).JSON(), map[string]any{
		"data": "ok",
		"meta": "ok",
	})

	g.Eq(g.Req("", tr.URL("/context")).JSON(), map[string]any{"data": "ok"})

	g.Eq(g.Req("", tr.URL("/request?a=ok")).JSON(), map[string]any{"data": "ok"})

	g.Eq(g.Req("", tr.URL("/params-time/2023-09-05T14:09:01.123Z")).JSON(), map[string]any{
		"data": "2023-09-05 14:09:01.123 +0000 UTC",
	})

	g.Eq(g.Req(http.MethodPost, tr.URL("/req-body"), `{"a": "ok"}`).JSON(), map[string]any{
		"data": "ok",
	})

	g.Eq(g.Req("", tr.URL("/error-res")).JSON(), map[string]interface{}{
		"error": map[string]interface{}{
			"code": "invalid_param",
		},
	})

	g.Eq(g.Req("", tr.URL("/override-res")).StatusCode, http.StatusNotModified)

	g.Eq(g.Req("", tr.URL("/override-header")).Header.Get("x-ua"), "test-client")

	g.Eq(g.Panic(func() {
		r.GET("/[", func() res { return resOK{} })
	}).(error).Error(), "error parsing regexp: missing closing ]: `[$`")

	g.Eq(g.Panic(func() {
		r.GET("/", 10)
	}), "handler must be a function or a struct with Handle method")

	g.Eq(g.Panic(func() {
		r.GET("/", func() {})
	}), "handler must return a single value")

	g.Eq(g.Req("", tr.URL("/res-enc-err")).JSON(), map[string]interface{}{
		"error": map[string]interface{} /* len=2 */ {
			"code":    "internal_error",
			"message": `/res-enc-err json: unsupported type: func()`, /* len=43 */
		},
	})

	g.Eq(g.Req("", tr.URL("/res-empty")).String(), "")

	g.Eq(g.Req("", tr.URL("/res-missed-type")).JSON(), map[string]interface{}{
		"error": map[string]interface{} /* len=2 */ {
			"code": "internal_error",
			"message": "handler response of path `/res-missed-type` must " +
				"goapi.Interface(new(goapi_test.res), goapi_test.resEmpty{})",
		},
	})

	g.Eq(g.Req("", tr.URL("/forget-create-interface")).JSON(), map[string]interface{}{
		"error": map[string]interface{} /* len=2 */ {
			"code": "internal_error",
			"message": "handler response of path `/forget-create-interface` " +
				"must goapi.Interface(new(goapi_test.resForgetCreateInterface))",
		},
	})

	res := g.Req("", tr.URL("/res-bin"))
	g.Eq(res.String(), "ok")
	g.Eq(res.Header.Get("Content-Type"), "application/octet-stream")

	g.Eq(g.Req("", tr.URL("/res-direct")).JSON(), "ok")

	res = g.Req("", tr.URL("/res-image"))
	g.Gt(res.Bytes().Len(), 1000)
	g.Eq(res.Header.Get("Content-Type"), "image/png")
}

func TestHeader(t *testing.T) {
	g := got.T(t)

	tr := g.Serve()
	r := goapi.New()

	type Header struct {
		goapi.InHeader
		X string `default:""`
		Y string `default:""`
	}

	r.GET("/header", func(h Header) res {
		return resOK{Data: h.X}
	})

	tr.Mux.Handle("/", r.Server())

	g.Eq(g.Req("", tr.URL("/header"), http.Header{"x": {"ok"}}).JSON(), map[string]any{"data": "ok"})
}
