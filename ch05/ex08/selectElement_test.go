package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	for _, test := range []struct {
		query, id, expected string
	}{
		//{"", ""},
		{`
<div id="hogefuga" name="hoge1fuga1">
	<div id="hoge" name="hoge1">
	</div>
	<div id="fuga" name="fuga1">
	</div>
</div>
`, "hoge",
			`<div id="hoge" name="hoge1">`},
		{`
<div id="hogefuga" name="hoge1fuga1">
	<div id="hoge" name="hoge1">
	</div>
	<div id="fuga" name="fuga1">
	</div>
</div>
`, "fuga",
			`<div id="fuga" name="fuga1">`},
		{`
<div id="hogefuga" name="hoge1fuga1">
	<div id="hoge" name="hoge1">
	</div>
	<div id="fuga" name="fuga1">
	</div>
</div>
`, "hogefuga",
			`<div id="hogefuga" name="hoge1fuga1">`},
	} {
		writer = new(bytes.Buffer)
		node, err := html.Parse(strings.NewReader(test.query))
		if err != nil {
			t.Errorf("parseErr \n %v", err)
		}
		printNode(ElementById(node, test.id))
		result := writer.(*bytes.Buffer).String()

		if result != test.expected {
			t.Errorf("result =\n%s, but expected was =\n%s", result, test.expected)
		}

	}
}
