package deserialize

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strconv"
)

func (ci *CustomInt64) UnmarshalJSON(data []byte) (err error) {
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}
	err = json.Unmarshal(data, &ci.Int64)
	if err != nil {
		return errors.New("CustomInt64: UnmarshalJSON: " + err.Error())
	}
	return
}

func (ci *CustomInt) UnmarshalJSON(data []byte) (err error) {
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}
	err = json.Unmarshal(data, &ci.Int)
	if err != nil {
		return errors.New("CustomInt: UnmarshalJSON: " + err.Error())
	}
	return
}

func (ci *CustomInt64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var data string
	err = d.DecodeElement(&data, &start)
	if err != nil {
		return err
	}
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}
	res, _ := strconv.Atoi(data)
	ci.Int64 = int64(res)
	return
}

func (ci *CustomInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var data string
	err = d.DecodeElement(&data, &start)
	if err != nil {
		return err
	}
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}
	res, _ := strconv.Atoi(data)
	ci.Int = res
	return
}

type CustomInt64 struct {
	Int64 int64
}
type CustomInt struct {
	Int int
}
type Users struct {
	Users []User `xml:"users"`
}
type User struct {
	ID      CustomInt64 `json:"id"      xml:"id"`
	Address Address     `json:"address" xml:"address"`
	Age     CustomInt   `json:"age"     xml:"age"`
}
type Address struct {
	CityID CustomInt64 `json:"city_id"  xml:"city_id"`
	Street string      `json:"street"   xml:"street"`
}
