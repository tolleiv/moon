package main

import (
	"testing"
)

var primitivesTests = []struct {
	in  string
	out token
}{
	{`"x"`, token{t_string, "x"}},
	{`"yes"`, token{t_string, "yes"}},
	{`"this one has spaces"`, token{t_string, "this one has spaces"}},
	{`"this one has \"quotes\" in it"`, token{t_string, `this one has "quotes" in it`}},
	{"`this one is delimited by backticks`", token{t_string, "this one is delimited by backticks"}},
}

func TestLexPrimities(t *testing.T) {
	for _, test := range primitivesTests {
		tokens, err := fullTokens(lexString(test.in))
		if err != nil {
			t.Error(err)
			continue
		}
		if len(tokens) > 1 {
			t.Errorf("expected 1 token, saw %d: %v", len(tokens), tokens)
			continue
		}
		t.Logf("OK: %s", test.in)
	}
}
