package main

var m = map[string]string{
	"a": "x",
	"b": "y",
	"c": "z",
}

func transform(s string) string {
	v, ok := m[s]
	if !ok {
		return "invalid"
	}
	return v
}
