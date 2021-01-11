package parser

import (
	"fmt"
	"strings"
)
type Result struct {
    Name  string
    Array []uint8
}


func (t *Result) MarshalJSON() ([]byte, error) {
    var array string
    if t.Array == nil {
        array = "null"
    } else {
        array = strings.Join(strings.Fields(fmt.Sprintf("%d", t.Array)), ",")
    }
    jsonResult := fmt.Sprintf(`{"Name":%q,"Array":%s}`, t.Name, array)
    return []byte(jsonResult), nil
}