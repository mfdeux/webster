package utils

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Options struct {
	Query   string `url:"q" schema:"q"`
	ShowAll bool   `url:"all" schema:"all"`
	Page    int    `url:"page" schema:"page"`
}

func TestMarshalQueryString(t *testing.T) {
	opt := Options{"foo", true, 2}
	qs, err := MarshalQueryString(opt)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expectedQS := "all=true&page=2&q=foo"
	if qs != expectedQS {
		t.Errorf("Payloads do not match: %s vs %s", qs, expectedQS)
		return
	}
}

func TestUnmarshalQueryString(t *testing.T) {
	opt := &Options{}
	expectedOpt := &Options{"foo", true, 2}
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	req.URL.RawQuery = "all=true&page=2&q=foo"
	err = UnmarshalQueryString(req, opt)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if !cmp.Equal(opt, expectedOpt) {
		t.Error("Did not get expected options")
		return
	}
}
