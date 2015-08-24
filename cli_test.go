package main_test

import (
	"github.com/shiwano/musta"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleRun_WithTemplateData() {
	args := []string{"foo=1", "bar.bar=a,b,c"}
	main.Run(args, "", "", "foo: {{foo}}, bar: {{#bar.bar}}{{value}}{{^last}},{{/last}}{{/bar.bar}}")
	// Output: foo: 1, bar: a,b,c
}

func ExampleRun_WithTemplateFile() {
	main.Run([]string{}, "{\"foo\": \"1\", \"bar\": \"2\"}", "./test.mustache", "")
	// Output: foo_file: 1, bar_file: 2
}

func TestParseArgs(t *testing.T) {
	args := []string{"foo=1", "bar.bar=2", "baz=100,200,300,"}
	jsonContainer := main.ParseArgs(args, "{\"foo\": \"100\", \"qux\": \"100\"}")

	assert.Equal(t, "1", jsonContainer.Path("foo").Data().(string))
	assert.Equal(t, "2", jsonContainer.Path("bar.bar").Data().(string))
	assert.Equal(t, "100", jsonContainer.Path("qux").Data().(string))
	assert.EqualValues(
		t,
		[]map[string]interface{}{
			map[string]interface{}{"index": 0, "first": true, "last": false, "value": "100"},
			map[string]interface{}{"index": 1, "first": false, "last": false, "value": "200"},
			map[string]interface{}{"index": 2, "first": false, "last": true, "value": "300"},
		},
		jsonContainer.Path("baz").Data().([]map[string]interface{}),
	)
}
