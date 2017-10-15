package convert

import (
	"encoding/json"
	"encoding/xml"
)

type contentType string

var ctJSON contentType = "application/json"
var ctXML contentType = "application/xml"

func (ct contentType) Marshal(val interface{}) ([]byte, error) {
	if ct == ctJSON {
		return json.Marshal(val)
	} else if ct == ctXML {
		return xml.Marshal(val)
	}
	panic("Wrong content type.")
}
