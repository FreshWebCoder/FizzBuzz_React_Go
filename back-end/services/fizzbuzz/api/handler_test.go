package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/starikode430/bunzz-test/cmd/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//nolint:ireturn,nolintlint,revive // reason: this is better for checking whether mock covers every function.
func newMockFizzbuzzService() *mockFizzbuzzService {
	sampleConfig := config.AppConfig{
		Fizz:     "Fizz",
		Buzz:     "Buzz",
		FizzBuzz: "FizzBuzz",
	}
	return &mockFizzbuzzService{
		config: sampleConfig,
	}
}

type mockFizzbuzzService struct {
	mock.Mock
	config config.AppConfig
}

func (m *mockFizzbuzzService) Fizzbuzz(count int) string {
	args := m.Called(count)
	return args.String(0)
}

func TestCannotDecode(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	service := newMockFizzbuzzService()
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"/fizzbuzz", http.NoBody)
	assertWithTest.Nil(err)
	res := httptest.NewRecorder()
	handler := NewHandler(service)
	handler.Fizzbuzz(res, req)
	assertWithTest.Equal(http.StatusBadRequest, res.Code)
}

func TestValidationFailed(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	service := newMockFizzbuzzService()
	type reqBody struct{}
	body, err := json.Marshal(&reqBody{})
	assertWithTest.Nil(err)
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"/fizzbuzz", bytes.NewReader(body))
	assertWithTest.Nil(err)
	res := httptest.NewRecorder()
	handler := NewHandler(service)
	handler.Fizzbuzz(res, req)
	assertWithTest.Equal(http.StatusBadRequest, res.Code)
}

func TestFizzSuccess(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	service := newMockFizzbuzzService()
	service.On("Fizzbuzz", 5).Return(service.config.Fizz)
	type reqBody struct {
		Count int `json:"count"`
	}
	requestBody := &reqBody{
		Count: 5,
	}
	body, err := json.Marshal(requestBody)
	assertWithTest.Nil(err)
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"/fizzbuzz", bytes.NewReader(body))
	assertWithTest.Nil(err)
	res := httptest.NewRecorder()
	handler := NewHandler(service)
	handler.Fizzbuzz(res, req)
	message := service.Fizzbuzz(requestBody.Count)
	assertWithTest.Equal(http.StatusOK, res.Code)
	assertWithTest.Equal(message, service.config.Fizz)
}

func TestBuzzSuccess(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	service := newMockFizzbuzzService()
	service.On("Fizzbuzz", 3).Return(service.config.Buzz)
	type reqBody struct {
		Count int `json:"count"`
	}
	requestBody := &reqBody{
		Count: 3,
	}
	body, err := json.Marshal(requestBody)
	assertWithTest.Nil(err)
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"/fizzbuzz", bytes.NewReader(body))
	assertWithTest.Nil(err)
	res := httptest.NewRecorder()
	handler := NewHandler(service)
	handler.Fizzbuzz(res, req)
	assertWithTest.Equal(http.StatusOK, res.Code)
}
