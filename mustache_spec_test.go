// generated by go run gen.go -output mustache_spec_test.go; DO NOT EDIT

package mustache

import (
	"strings"
	"testing"
)

func convertHTMLCharsToExpectedFormat(s string) string {
	return strings.Replace(s, "&#34;", "&quot;", -1)
}

func testSpec(t *testing.T,
	template string,
	expected string,
	context interface{}) {
	output := convertHTMLCharsToExpectedFormat(Render(template, context))
	if output != expected {
		t.Errorf("%q expected <%q> but got <%q>",
			template, expected, output)
	}
}

func TestCommentsInline(t *testing.T) {
	testSpec(t,
		"12345{{! Comment Block! }}67890",
		"1234567890",
		map[string]interface{}{})
}

func TestCommentsMultiline(t *testing.T) {
	testSpec(t,
		"12345{{!\n  This is a\n  multi-line comment...\n}}67890\n",
		"1234567890\n",
		map[string]interface{}{})
}

func TestCommentsStandalone(t *testing.T) {
	testSpec(t,
		"Begin.\n{{! Comment Block! }}\nEnd.\n",
		"Begin.\nEnd.\n",
		map[string]interface{}{})
}

func TestCommentsIndentedStandalone(t *testing.T) {
	testSpec(t,
		"Begin.\n  {{! Indented Comment Block! }}\nEnd.\n",
		"Begin.\nEnd.\n",
		map[string]interface{}{})
}

func TestCommentsStandaloneLineEndings(t *testing.T) {
	testSpec(t,
		"|\r\n{{! Standalone Comment }}\r\n|",
		"|\r\n|",
		map[string]interface{}{})
}

func TestCommentsStandaloneWithoutPreviousLine(t *testing.T) {
	testSpec(t,
		"  {{! I'm Still Standalone }}\n!",
		"!",
		map[string]interface{}{})
}

func TestCommentsStandaloneWithoutNewline(t *testing.T) {
	testSpec(t,
		"!\n  {{! I'm Still Standalone }}",
		"!\n",
		map[string]interface{}{})
}

func TestCommentsMultilineStandalone(t *testing.T) {
	testSpec(t,
		"Begin.\n{{!\nSomething's going on here...\n}}\nEnd.\n",
		"Begin.\nEnd.\n",
		map[string]interface{}{})
}

func TestCommentsIndentedMultilineStandalone(t *testing.T) {
	testSpec(t,
		"Begin.\n  {{!\n    Something's going on here...\n  }}\nEnd.\n",
		"Begin.\nEnd.\n",
		map[string]interface{}{})
}

func TestCommentsIndentedInline(t *testing.T) {
	testSpec(t,
		"  12 {{! 34 }}\n",
		"  12 \n",
		map[string]interface{}{})
}

func TestCommentsSurroundingWhitespace(t *testing.T) {
	testSpec(t,
		"12345 {{! Comment Block! }} 67890",
		"12345  67890",
		map[string]interface{}{})
}

func TestInterpolationNoInterpolation(t *testing.T) {
	testSpec(t,
		"Hello from {Mustache}!\n",
		"Hello from {Mustache}!\n",
		map[string]interface{}{})
}

func TestInterpolationBasicInterpolation(t *testing.T) {
	testSpec(t,
		"Hello, {{subject}}!\n",
		"Hello, world!\n",
		map[string]interface{}{"subject": "world"})
}

func TestInterpolationHTMLEscaping(t *testing.T) {
	testSpec(t,
		"These characters should be HTML escaped: {{forbidden}}\n",
		"These characters should be HTML escaped: &amp; &quot; &lt; &gt;\n",
		map[string]interface{}{"forbidden": "& \" < >"})
}

func TestInterpolationTripleMustache(t *testing.T) {
	testSpec(t,
		"These characters should not be HTML escaped: {{{forbidden}}}\n",
		"These characters should not be HTML escaped: & \" < >\n",
		map[string]interface{}{"forbidden": "& \" < >"})
}

func TestInterpolationAmpersand(t *testing.T) {
	testSpec(t,
		"These characters should not be HTML escaped: {{&forbidden}}\n",
		"These characters should not be HTML escaped: & \" < >\n",
		map[string]interface{}{"forbidden": "& \" < >"})
}

func TestInterpolationBasicIntegerInterpolation(t *testing.T) {
	testSpec(t,
		"\"{{mph}} miles an hour!\"",
		"\"85 miles an hour!\"",
		map[string]interface{}{"mph": 85})
}

func TestInterpolationTripleMustacheIntegerInterpolation(t *testing.T) {
	testSpec(t,
		"\"{{{mph}}} miles an hour!\"",
		"\"85 miles an hour!\"",
		map[string]interface{}{"mph": 85})
}

func TestInterpolationAmpersandIntegerInterpolation(t *testing.T) {
	testSpec(t,
		"\"{{&mph}} miles an hour!\"",
		"\"85 miles an hour!\"",
		map[string]interface{}{"mph": 85})
}

func TestInterpolationBasicDecimalInterpolation(t *testing.T) {
	testSpec(t,
		"\"{{power}} jiggawatts!\"",
		"\"1.21 jiggawatts!\"",
		map[string]interface{}{"power": 1.21})
}

func TestInterpolationTripleMustacheDecimalInterpolation(t *testing.T) {
	testSpec(t,
		"\"{{{power}}} jiggawatts!\"",
		"\"1.21 jiggawatts!\"",
		map[string]interface{}{"power": 1.21})
}

func TestInterpolationAmpersandDecimalInterpolation(t *testing.T) {
	testSpec(t,
		"\"{{&power}} jiggawatts!\"",
		"\"1.21 jiggawatts!\"",
		map[string]interface{}{"power": 1.21})
}

func TestInterpolationBasicContextMissInterpolation(t *testing.T) {
	testSpec(t,
		"I ({{cannot}}) be seen!",
		"I () be seen!",
		map[string]interface{}{})
}

func TestInterpolationTripleMustacheContextMissInterpolation(t *testing.T) {
	testSpec(t,
		"I ({{{cannot}}}) be seen!",
		"I () be seen!",
		map[string]interface{}{})
}

func TestInterpolationAmpersandContextMissInterpolation(t *testing.T) {
	testSpec(t,
		"I ({{&cannot}}) be seen!",
		"I () be seen!",
		map[string]interface{}{})
}

func TestInterpolationDottedNamesBasicInterpolation(t *testing.T) {
	testSpec(t,
		"\"{{person.name}}\" == \"{{#person}}{{name}}{{/person}}\"",
		"\"Joe\" == \"Joe\"",
		map[string]interface{}{"person": map[string]interface{}{"name": "Joe"}})
}

func TestInterpolationDottedNamesTripleMustacheInterpolation(t *testing.T) {
	testSpec(t,
		"\"{{{person.name}}}\" == \"{{#person}}{{{name}}}{{/person}}\"",
		"\"Joe\" == \"Joe\"",
		map[string]interface{}{"person": map[string]interface{}{"name": "Joe"}})
}

func TestInterpolationDottedNamesAmpersandInterpolation(t *testing.T) {
	testSpec(t,
		"\"{{&person.name}}\" == \"{{#person}}{{&name}}{{/person}}\"",
		"\"Joe\" == \"Joe\"",
		map[string]interface{}{"person": map[string]interface{}{"name": "Joe"}})
}

func TestInterpolationDottedNamesArbitraryDepth(t *testing.T) {
	testSpec(t,
		"\"{{a.b.c.d.e.name}}\" == \"Phil\"",
		"\"Phil\" == \"Phil\"",
		map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": map[string]interface{}{"d": map[string]interface{}{"e": map[string]interface{}{"name": "Phil"}}}}}})
}

func TestInterpolationDottedNamesBrokenChains(t *testing.T) {
	testSpec(t,
		"\"{{a.b.c}}\" == \"\"",
		"\"\" == \"\"",
		map[string]interface{}{"a": map[string]interface{}{}})
}

func TestInterpolationDottedNamesBrokenChainResolution(t *testing.T) {
	testSpec(t,
		"\"{{a.b.c.name}}\" == \"\"",
		"\"\" == \"\"",
		map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{}}, "c": map[string]interface{}{"name": "Jim"}})
}

func TestInterpolationDottedNamesInitialResolution(t *testing.T) {
	testSpec(t,
		"\"{{#a}}{{b.c.d.e.name}}{{/a}}\" == \"Phil\"",
		"\"Phil\" == \"Phil\"",
		map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": map[string]interface{}{"d": map[string]interface{}{"e": map[string]interface{}{"name": "Phil"}}}}}, "b": map[string]interface{}{"c": map[string]interface{}{"d": map[string]interface{}{"e": map[string]interface{}{"name": "Wrong"}}}}})
}

func TestInterpolationInterpolationSurroundingWhitespace(t *testing.T) {
	testSpec(t,
		"| {{string}} |",
		"| --- |",
		map[string]interface{}{"string": "---"})
}

func TestInterpolationTripleMustacheSurroundingWhitespace(t *testing.T) {
	testSpec(t,
		"| {{{string}}} |",
		"| --- |",
		map[string]interface{}{"string": "---"})
}

func TestInterpolationAmpersandSurroundingWhitespace(t *testing.T) {
	testSpec(t,
		"| {{&string}} |",
		"| --- |",
		map[string]interface{}{"string": "---"})
}

func TestInterpolationInterpolationStandalone(t *testing.T) {
	testSpec(t,
		"  {{string}}\n",
		"  ---\n",
		map[string]interface{}{"string": "---"})
}

func TestInterpolationTripleMustacheStandalone(t *testing.T) {
	testSpec(t,
		"  {{{string}}}\n",
		"  ---\n",
		map[string]interface{}{"string": "---"})
}

func TestInterpolationAmpersandStandalone(t *testing.T) {
	testSpec(t,
		"  {{&string}}\n",
		"  ---\n",
		map[string]interface{}{"string": "---"})
}

func TestInterpolationInterpolationWithPadding(t *testing.T) {
	testSpec(t,
		"|{{ string }}|",
		"|---|",
		map[string]interface{}{"string": "---"})
}

func TestInterpolationTripleMustacheWithPadding(t *testing.T) {
	testSpec(t,
		"|{{{ string }}}|",
		"|---|",
		map[string]interface{}{"string": "---"})
}

func TestInterpolationAmpersandWithPadding(t *testing.T) {
	testSpec(t,
		"|{{& string }}|",
		"|---|",
		map[string]interface{}{"string": "---"})
}
