// Code generated by "enumer -type=StatusCode -values -trimprefix=Status -json"; DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _StatusCodeName = "200201202204301302303304305307400401402403404405406407408409410411412413414415416417418421422423424425426428429431451500501502503504505506507508510511"
const _StatusCodeLowerName = "200201202204301302303304305307400401402403404405406407408409410411412413414415416417418421422423424425426428429431451500501502503504505506507508510511"

var _StatusCodeMap = map[StatusCode]string{
	200: _StatusCodeName[0:3],
	201: _StatusCodeName[3:6],
	202: _StatusCodeName[6:9],
	204: _StatusCodeName[9:12],
	301: _StatusCodeName[12:15],
	302: _StatusCodeName[15:18],
	303: _StatusCodeName[18:21],
	304: _StatusCodeName[21:24],
	305: _StatusCodeName[24:27],
	307: _StatusCodeName[27:30],
	400: _StatusCodeName[30:33],
	401: _StatusCodeName[33:36],
	402: _StatusCodeName[36:39],
	403: _StatusCodeName[39:42],
	404: _StatusCodeName[42:45],
	405: _StatusCodeName[45:48],
	406: _StatusCodeName[48:51],
	407: _StatusCodeName[51:54],
	408: _StatusCodeName[54:57],
	409: _StatusCodeName[57:60],
	410: _StatusCodeName[60:63],
	411: _StatusCodeName[63:66],
	412: _StatusCodeName[66:69],
	413: _StatusCodeName[69:72],
	414: _StatusCodeName[72:75],
	415: _StatusCodeName[75:78],
	416: _StatusCodeName[78:81],
	417: _StatusCodeName[81:84],
	418: _StatusCodeName[84:87],
	421: _StatusCodeName[87:90],
	422: _StatusCodeName[90:93],
	423: _StatusCodeName[93:96],
	424: _StatusCodeName[96:99],
	425: _StatusCodeName[99:102],
	426: _StatusCodeName[102:105],
	428: _StatusCodeName[105:108],
	429: _StatusCodeName[108:111],
	431: _StatusCodeName[111:114],
	451: _StatusCodeName[114:117],
	500: _StatusCodeName[117:120],
	501: _StatusCodeName[120:123],
	502: _StatusCodeName[123:126],
	503: _StatusCodeName[126:129],
	504: _StatusCodeName[129:132],
	505: _StatusCodeName[132:135],
	506: _StatusCodeName[135:138],
	507: _StatusCodeName[138:141],
	508: _StatusCodeName[141:144],
	510: _StatusCodeName[144:147],
	511: _StatusCodeName[147:150],
}

func (i StatusCode) String() string {
	if str, ok := _StatusCodeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("StatusCode(%d)", i)
}

func (StatusCode) Values() []string {
	return StatusCodeStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _StatusCodeNoOp() {
	var x [1]struct{}
	_ = x[Status200-(200)]
	_ = x[Status201-(201)]
	_ = x[Status202-(202)]
	_ = x[Status204-(204)]
	_ = x[Status301-(301)]
	_ = x[Status302-(302)]
	_ = x[Status303-(303)]
	_ = x[Status304-(304)]
	_ = x[Status305-(305)]
	_ = x[Status307-(307)]
	_ = x[Status400-(400)]
	_ = x[Status401-(401)]
	_ = x[Status402-(402)]
	_ = x[Status403-(403)]
	_ = x[Status404-(404)]
	_ = x[Status405-(405)]
	_ = x[Status406-(406)]
	_ = x[Status407-(407)]
	_ = x[Status408-(408)]
	_ = x[Status409-(409)]
	_ = x[Status410-(410)]
	_ = x[Status411-(411)]
	_ = x[Status412-(412)]
	_ = x[Status413-(413)]
	_ = x[Status414-(414)]
	_ = x[Status415-(415)]
	_ = x[Status416-(416)]
	_ = x[Status417-(417)]
	_ = x[Status418-(418)]
	_ = x[Status421-(421)]
	_ = x[Status422-(422)]
	_ = x[Status423-(423)]
	_ = x[Status424-(424)]
	_ = x[Status425-(425)]
	_ = x[Status426-(426)]
	_ = x[Status428-(428)]
	_ = x[Status429-(429)]
	_ = x[Status431-(431)]
	_ = x[Status451-(451)]
	_ = x[Status500-(500)]
	_ = x[Status501-(501)]
	_ = x[Status502-(502)]
	_ = x[Status503-(503)]
	_ = x[Status504-(504)]
	_ = x[Status505-(505)]
	_ = x[Status506-(506)]
	_ = x[Status507-(507)]
	_ = x[Status508-(508)]
	_ = x[Status510-(510)]
	_ = x[Status511-(511)]
}

var _StatusCodeValues = []StatusCode{Status200, Status201, Status202, Status204, Status301, Status302, Status303, Status304, Status305, Status307, Status400, Status401, Status402, Status403, Status404, Status405, Status406, Status407, Status408, Status409, Status410, Status411, Status412, Status413, Status414, Status415, Status416, Status417, Status418, Status421, Status422, Status423, Status424, Status425, Status426, Status428, Status429, Status431, Status451, Status500, Status501, Status502, Status503, Status504, Status505, Status506, Status507, Status508, Status510, Status511}

var _StatusCodeNameToValueMap = map[string]StatusCode{
	_StatusCodeName[0:3]:          Status200,
	_StatusCodeLowerName[0:3]:     Status200,
	_StatusCodeName[3:6]:          Status201,
	_StatusCodeLowerName[3:6]:     Status201,
	_StatusCodeName[6:9]:          Status202,
	_StatusCodeLowerName[6:9]:     Status202,
	_StatusCodeName[9:12]:         Status204,
	_StatusCodeLowerName[9:12]:    Status204,
	_StatusCodeName[12:15]:        Status301,
	_StatusCodeLowerName[12:15]:   Status301,
	_StatusCodeName[15:18]:        Status302,
	_StatusCodeLowerName[15:18]:   Status302,
	_StatusCodeName[18:21]:        Status303,
	_StatusCodeLowerName[18:21]:   Status303,
	_StatusCodeName[21:24]:        Status304,
	_StatusCodeLowerName[21:24]:   Status304,
	_StatusCodeName[24:27]:        Status305,
	_StatusCodeLowerName[24:27]:   Status305,
	_StatusCodeName[27:30]:        Status307,
	_StatusCodeLowerName[27:30]:   Status307,
	_StatusCodeName[30:33]:        Status400,
	_StatusCodeLowerName[30:33]:   Status400,
	_StatusCodeName[33:36]:        Status401,
	_StatusCodeLowerName[33:36]:   Status401,
	_StatusCodeName[36:39]:        Status402,
	_StatusCodeLowerName[36:39]:   Status402,
	_StatusCodeName[39:42]:        Status403,
	_StatusCodeLowerName[39:42]:   Status403,
	_StatusCodeName[42:45]:        Status404,
	_StatusCodeLowerName[42:45]:   Status404,
	_StatusCodeName[45:48]:        Status405,
	_StatusCodeLowerName[45:48]:   Status405,
	_StatusCodeName[48:51]:        Status406,
	_StatusCodeLowerName[48:51]:   Status406,
	_StatusCodeName[51:54]:        Status407,
	_StatusCodeLowerName[51:54]:   Status407,
	_StatusCodeName[54:57]:        Status408,
	_StatusCodeLowerName[54:57]:   Status408,
	_StatusCodeName[57:60]:        Status409,
	_StatusCodeLowerName[57:60]:   Status409,
	_StatusCodeName[60:63]:        Status410,
	_StatusCodeLowerName[60:63]:   Status410,
	_StatusCodeName[63:66]:        Status411,
	_StatusCodeLowerName[63:66]:   Status411,
	_StatusCodeName[66:69]:        Status412,
	_StatusCodeLowerName[66:69]:   Status412,
	_StatusCodeName[69:72]:        Status413,
	_StatusCodeLowerName[69:72]:   Status413,
	_StatusCodeName[72:75]:        Status414,
	_StatusCodeLowerName[72:75]:   Status414,
	_StatusCodeName[75:78]:        Status415,
	_StatusCodeLowerName[75:78]:   Status415,
	_StatusCodeName[78:81]:        Status416,
	_StatusCodeLowerName[78:81]:   Status416,
	_StatusCodeName[81:84]:        Status417,
	_StatusCodeLowerName[81:84]:   Status417,
	_StatusCodeName[84:87]:        Status418,
	_StatusCodeLowerName[84:87]:   Status418,
	_StatusCodeName[87:90]:        Status421,
	_StatusCodeLowerName[87:90]:   Status421,
	_StatusCodeName[90:93]:        Status422,
	_StatusCodeLowerName[90:93]:   Status422,
	_StatusCodeName[93:96]:        Status423,
	_StatusCodeLowerName[93:96]:   Status423,
	_StatusCodeName[96:99]:        Status424,
	_StatusCodeLowerName[96:99]:   Status424,
	_StatusCodeName[99:102]:       Status425,
	_StatusCodeLowerName[99:102]:  Status425,
	_StatusCodeName[102:105]:      Status426,
	_StatusCodeLowerName[102:105]: Status426,
	_StatusCodeName[105:108]:      Status428,
	_StatusCodeLowerName[105:108]: Status428,
	_StatusCodeName[108:111]:      Status429,
	_StatusCodeLowerName[108:111]: Status429,
	_StatusCodeName[111:114]:      Status431,
	_StatusCodeLowerName[111:114]: Status431,
	_StatusCodeName[114:117]:      Status451,
	_StatusCodeLowerName[114:117]: Status451,
	_StatusCodeName[117:120]:      Status500,
	_StatusCodeLowerName[117:120]: Status500,
	_StatusCodeName[120:123]:      Status501,
	_StatusCodeLowerName[120:123]: Status501,
	_StatusCodeName[123:126]:      Status502,
	_StatusCodeLowerName[123:126]: Status502,
	_StatusCodeName[126:129]:      Status503,
	_StatusCodeLowerName[126:129]: Status503,
	_StatusCodeName[129:132]:      Status504,
	_StatusCodeLowerName[129:132]: Status504,
	_StatusCodeName[132:135]:      Status505,
	_StatusCodeLowerName[132:135]: Status505,
	_StatusCodeName[135:138]:      Status506,
	_StatusCodeLowerName[135:138]: Status506,
	_StatusCodeName[138:141]:      Status507,
	_StatusCodeLowerName[138:141]: Status507,
	_StatusCodeName[141:144]:      Status508,
	_StatusCodeLowerName[141:144]: Status508,
	_StatusCodeName[144:147]:      Status510,
	_StatusCodeLowerName[144:147]: Status510,
	_StatusCodeName[147:150]:      Status511,
	_StatusCodeLowerName[147:150]: Status511,
}

var _StatusCodeNames = []string{
	_StatusCodeName[0:3],
	_StatusCodeName[3:6],
	_StatusCodeName[6:9],
	_StatusCodeName[9:12],
	_StatusCodeName[12:15],
	_StatusCodeName[15:18],
	_StatusCodeName[18:21],
	_StatusCodeName[21:24],
	_StatusCodeName[24:27],
	_StatusCodeName[27:30],
	_StatusCodeName[30:33],
	_StatusCodeName[33:36],
	_StatusCodeName[36:39],
	_StatusCodeName[39:42],
	_StatusCodeName[42:45],
	_StatusCodeName[45:48],
	_StatusCodeName[48:51],
	_StatusCodeName[51:54],
	_StatusCodeName[54:57],
	_StatusCodeName[57:60],
	_StatusCodeName[60:63],
	_StatusCodeName[63:66],
	_StatusCodeName[66:69],
	_StatusCodeName[69:72],
	_StatusCodeName[72:75],
	_StatusCodeName[75:78],
	_StatusCodeName[78:81],
	_StatusCodeName[81:84],
	_StatusCodeName[84:87],
	_StatusCodeName[87:90],
	_StatusCodeName[90:93],
	_StatusCodeName[93:96],
	_StatusCodeName[96:99],
	_StatusCodeName[99:102],
	_StatusCodeName[102:105],
	_StatusCodeName[105:108],
	_StatusCodeName[108:111],
	_StatusCodeName[111:114],
	_StatusCodeName[114:117],
	_StatusCodeName[117:120],
	_StatusCodeName[120:123],
	_StatusCodeName[123:126],
	_StatusCodeName[126:129],
	_StatusCodeName[129:132],
	_StatusCodeName[132:135],
	_StatusCodeName[135:138],
	_StatusCodeName[138:141],
	_StatusCodeName[141:144],
	_StatusCodeName[144:147],
	_StatusCodeName[147:150],
}

// StatusCodeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func StatusCodeString(s string) (StatusCode, error) {
	if val, ok := _StatusCodeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _StatusCodeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to StatusCode values", s)
}

// StatusCodeValues returns all values of the enum
func StatusCodeValues() []StatusCode {
	return _StatusCodeValues
}

// StatusCodeStrings returns a slice of all String values of the enum
func StatusCodeStrings() []string {
	strs := make([]string, len(_StatusCodeNames))
	copy(strs, _StatusCodeNames)
	return strs
}

// IsAStatusCode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i StatusCode) IsAStatusCode() bool {
	_, ok := _StatusCodeMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for StatusCode
func (i StatusCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatusCode
func (i *StatusCode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("StatusCode should be a string, got %s", data)
	}

	var err error
	*i, err = StatusCodeString(s)
	return err
}