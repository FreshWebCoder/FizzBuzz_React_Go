package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/starikode430/bunzz-test/services/fizzbuzz/fizzbuzz"
)

var errCannotDecode = errors.New("cannot_decode")

type handler struct {
	fizzbuzzService fizzbuzz.Service
}

type FizzbuzzHandler interface {
	Fizzbuzz(res http.ResponseWriter, req *http.Request)
	HealthCheck(res http.ResponseWriter, req *http.Request)
}

//nolint:ireturn,nolintlint // reason: returning interface is useful for mocking.
func NewHandler(service fizzbuzz.Service) FizzbuzzHandler {
	return &handler{
		fizzbuzzService: service,
	}
}

type fizzbuzzRequest struct {
	Count int `json:"count,omitempty" validate:"required"`
}

type fizzbuzzResponse struct {
	Message string `json:"message"`
}

func (h *handler) Fizzbuzz(res http.ResponseWriter, req *http.Request) {
	validate := validator.New()

	defer req.Body.Close()

	var reqBody fizzbuzzRequest
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		http.Error(res, errCannotDecode.Error(), http.StatusBadRequest)
		return
	}

	// use the validator library to validate required fields
	if validationErr := validate.Struct(&reqBody); validationErr != nil {
		http.Error(res, validationErr.Error(), http.StatusBadRequest)
		return
	}

	respBody := fizzbuzzResponse{
		Message: h.fizzbuzzService.Fizzbuzz(reqBody.Count),
	}
	//nolint:errchkjson,errchkjson //reason: json.NewEncoder(res).Encode() cannot mock.
	_ = json.NewEncoder(res).Encode(&respBody)
}

func (h *handler) HealthCheck(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	//nolint:errcheck //reason: res.Write() cannot mock.
	res.Write([]byte("success"))
}
