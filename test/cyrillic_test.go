package test

import (
	"github.com/Torebekov/L5/internals/filter"
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Surname string
	Age     int
	Student *Student
}
type Student struct {
	University *string
	Employee   Employee
	Address    string
}
type Employee struct {
	Company  string
	Position string
}

func TestSomeFunc(t *testing.T) {
	wrongUName := "KBTUуник"
	rightUName := "KBTU"
	test := struct {
		input    Person
		expected Person
	}{
		input: Person{
			"Johnбек",
			"Wickеландо",
			2021,
			&Student{
				&wrongUName,
				Employee{
					"OneLab",
					"StudentөҠ",
				},
				"Almatyқаласы",
			},
		},
		expected: Person{
			"John",
			"Wick",
			2021,
			&Student{
				&rightUName,
				Employee{
					"OneLab",
					"Student",
				},
				"Almaty",
			},
		},
	}

	filter.SomeFunc(&test.input)
	if !reflect.DeepEqual(test.input, test.expected) {
		t.Errorf("Incorrect result. Expect %v, got %v",
			test.expected, test.input)
	}
}
