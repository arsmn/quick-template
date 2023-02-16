package api

import (
	"net/http"

	"github.com/ory/herodot"
)

type writer struct {
	herodot.Writer
}

func newWriter() writer {
	return writer{
		Writer: herodot.NewJSONWriter(nil),
	}
}

func (w writer) Write(rw http.ResponseWriter, r *http.Request, e any) error {
	w.Writer.Write(rw, r, e)
	return nil
}

func (w writer) WriteCode(rw http.ResponseWriter, r *http.Request, code int, e any) error {
	w.Writer.WriteCode(rw, r, code, e)
	return nil
}

func (w writer) WriteCreated(rw http.ResponseWriter, r *http.Request, location string, e any) error {
	w.Writer.WriteCreated(rw, r, location, e)
	return nil
}

func (w writer) WriteError(rw http.ResponseWriter, r *http.Request, err error) error {
	w.Writer.WriteError(rw, r, err)
	return nil
}

func (w writer) WriteErrorCode(rw http.ResponseWriter, r *http.Request, code int, err error) error {
	w.Writer.WriteErrorCode(rw, r, code, err)
	return nil
}
