// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	http "net/http"
)

func AnyHttpResponseWriter() http.ResponseWriter {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(http.ResponseWriter))(nil)).Elem()))
	var nullValue http.ResponseWriter
	return nullValue
}

func EqHttpResponseWriter(value http.ResponseWriter) http.ResponseWriter {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue http.ResponseWriter
	return nullValue
}

func NotEqHttpResponseWriter(value http.ResponseWriter) http.ResponseWriter {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue http.ResponseWriter
	return nullValue
}

func HttpResponseWriterThat(matcher pegomock.ArgumentMatcher) http.ResponseWriter {
	pegomock.RegisterMatcher(matcher)
	var nullValue http.ResponseWriter
	return nullValue
}
