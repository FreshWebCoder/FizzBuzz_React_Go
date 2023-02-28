package fizzbuzz

import (
	"testing"

	"github.com/starikode430/bunzz-test/cmd/config"
	"github.com/stretchr/testify/assert"
)

var mockConfig = config.AppConfig{
	Fizz:     "Fizz",
	Buzz:     "Buzz",
	FizzBuzz: "FizzBuzz",
}

var testCases = []struct {
	Count   int
	Message string
}{
	{
		Count:   5,
		Message: "Buzz",
	},
	{
		Count:   3,
		Message: "Fizz",
	},
	{
		Count:   15,
		Message: "FizzBuzz",
	},
	{
		Count:   56,
		Message: "",
	},
}

func TestFizz(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	service := NewFizzbuzzService(mockConfig)
	for _, testCase := range testCases {
		messsage := service.Fizzbuzz(testCase.Count)
		assertWithTest.Equal(messsage, testCase.Message)
	}
}
