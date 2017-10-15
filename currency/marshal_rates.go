package currency

import "encoding/xml"

type xmlMapEntry struct {
	XMLName xml.Name
	Value   float64 `xml:"value,attr"`
}

// MarshalXML marshals rates to XML, with each key in the map being a
// tag and it's corresponding value being it's contents.
func (m Rates) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{
			XMLName: xml.Name{Local: k},
			Value:   v,
		})
	}

	return e.EncodeToken(start.End())
}
