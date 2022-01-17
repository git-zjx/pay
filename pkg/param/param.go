package param

import (
	"encoding/xml"
	"io"
	"net/url"
	"pay/pkg/helper"
	"sort"
)

type Params map[string]interface{}

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (base Params) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(base) == 0 {
		return nil
	}
	start.Name = xml.Name{Local: "xml"}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for k, v := range base {
		err := e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: helper.Strval(v)})
		if err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

func (base *Params) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*base = Params{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*base)[e.XMLName.Local] = e.Value
	}
	return nil
}

func (base Params) ToUrlValue() string {
	var keys []string
	value := url.Values{}
	for k := range base {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		value.Add(k, helper.Strval(base[k]))
	}
	return value.Encode()
}
