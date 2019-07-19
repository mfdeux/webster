package utils

import (
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/gorilla/schema"
)

// Reference: http://www.gorillatoolkit.org/pkg/schema
// https://github.com/google/go-querystring/issues/7

// MarshalQueryString marshals an interface into a query string
func MarshalQueryString(v interface{}) (string, error) {
	values, err := query.Values(v)
	if err != nil {
		return "", err
	}
	return values.Encode(), nil
}

// UnmarshalQueryString unmarshals a query string from an http request into a struct
// Uses the Gorilla schema package
// type Person struct {
// 	Name  string `schema:"name"` // custom name
// 	Admin bool   `schema:"-"`    // this field is never set
// }
func UnmarshalQueryString(r *http.Request, v interface{}) error {
	decoder := schema.NewDecoder()
	if err := decoder.Decode(v, r.URL.Query()); err != nil {
		return err
	}
	return nil
}
