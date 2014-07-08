package iglo

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	data, err := ParseMarkdown(strings.NewReader(dummyMarkdown))
	if err != nil {
		t.Errorf("ParseMarkdown() returned an error %s", err)
	}

	ioutil.WriteFile("x.json", data, os.ModePerm)

	a, err := ParseJSON(bytes.NewBuffer(data))
	if err == nil {
		if !reflect.DeepEqual(a, dummyAPI) {
			t.Errorf("ParseJSON() returned %+v, want %+v", a, dummyAPI)
		}
	}
}
