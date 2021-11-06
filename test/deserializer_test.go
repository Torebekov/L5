package test

import (
	"encoding/json"
	"encoding/xml"
	"github.com/Torebekov/L5/internals/deserialize"
	"reflect"
	"testing"
)

var rawJson = []byte(`[
  {
    "id": 1,
    "address": {
      "city_id": 5,
      "street": "Satbayev"
    },
    "Age": 20
  },
  {
    "id": 1,
    "address": {
      "city_id": "6",
      "street": "Al-Farabi"
    },
    "Age": "32"
  }
]`)
var rawXML = []byte(`
<Users>
  <users>
    <id>"1"</id>
	<address>
		<city_id>5</city_id>
		<street>Satbayev</street>
	</address>
	<age>20</age>
  </users>
  <users>
    <id>1</id>
	<address>
		<city_id>"6"</city_id>
		<street>Al-Farabi</street>
	</address>
	<age>"32"</age>
  </users>
</Users>`)

func TestUnmarshalJSON(t *testing.T) {
	test := struct {
		input    []byte
		expected []deserialize.User
	}{
		input:    rawJson,
		expected: expectedJSON,
	}
	var users []deserialize.User
	if err := json.Unmarshal(test.input, &users); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(users, test.expected) {
		t.Errorf("Incorrect result. Expect %v, got %v",
			test.expected, users)
	}
}
func TestUnmarshalXML(t *testing.T) {
	test := struct {
		input    []byte
		expected deserialize.Users
	}{
		input:    rawXML,
		expected: expectedXML,
	}
	var xmlUsers deserialize.Users
	if err := xml.Unmarshal(test.input, &xmlUsers); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(xmlUsers, test.expected) {
		t.Errorf("Incorrect result. Expect %v, got %v",
			test.expected, xmlUsers)
	}
}

var expectedJSON = []deserialize.User{
	{
		deserialize.CustomInt64{1},
		deserialize.Address{
			deserialize.CustomInt64{5},
			"Satbayev",
		},
		deserialize.CustomInt{20},
	},
	{
		deserialize.CustomInt64{1},
		deserialize.Address{
			deserialize.CustomInt64{6},
			"Al-Farabi",
		},
		deserialize.CustomInt{32},
	},
}
var expectedXML = deserialize.Users{
	expectedJSON,
}
