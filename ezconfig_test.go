package ezconfig

import (
	"os"
	"strconv"
	"testing"
)

func TestParse(t *testing.T) {
	type Enviroments struct {
		SomeName    string `env:"someName"`
		SomeInteger int    `env:"someInteger"`
	}
	ex := Enviroments{
		SomeName:    "helloworld",
		SomeInteger: 20,
	}
	os.Setenv("someName", ex.SomeName)
	os.Setenv("someInteger", strconv.Itoa(ex.SomeInteger))

	var re Enviroments
	err := Parse(&re)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if re.SomeName != ex.SomeName {
		t.Fatalf("'SomeName' mismatch, ex: %s, re: %s", ex.SomeName, re.SomeName)
	}
	if re.SomeInteger != ex.SomeInteger {
		t.Fatalf("'SomeInteger' mismatch, ex: %d, re: %d", ex.SomeInteger, re.SomeInteger)
	}
}
