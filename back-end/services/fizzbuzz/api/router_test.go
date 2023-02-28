package api

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	router := mux.NewRouter()
	NewFizzbuzzRouter(router, &handler{})
	assertWithTest.NotNil(router)
}
