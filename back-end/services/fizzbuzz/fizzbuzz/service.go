package fizzbuzz

import "github.com/starikode430/bunzz-test/cmd/config"

type service struct {
	config config.AppConfig
}

type Service interface {
	Fizzbuzz(count int) string
}

//nolint:ireturn,nolintlint // reason: this is necessary for mock.
func NewFizzbuzzService(appConfig config.AppConfig) Service {
	return &service{
		config: appConfig,
	}
}

func (s *service) Fizzbuzz(count int) string {
	if count%15 == 0 {
		return s.config.FizzBuzz
	}
	if count%3 == 0 {
		return s.config.Fizz
	}
	if count%5 == 0 {
		return s.config.Buzz
	}
	return ""
}
