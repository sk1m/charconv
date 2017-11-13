package charconv

import (
	"io/ioutil"
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func Decode(other, enc string) string {
	e, _ := charset.Lookup(enc)
	if e == nil {
		return ""
	}
	s, _ := transformString(e.NewDecoder(), other)
	return s
}

func Encode(utf8, enc string) string {
	e, _ := charset.Lookup(enc)
	if e == nil {
		return ""
	}
	s, _ := transformString(e.NewEncoder(), utf8)
	return s
}

func transformString(t transform.Transformer, s string) (string, error) {
	r := transform.NewReader(strings.NewReader(s), t)
	b, err := ioutil.ReadAll(r)
	return string(b), err
}
