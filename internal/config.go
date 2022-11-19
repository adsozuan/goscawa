package goscawa

import "fmt"

type environment int

const (
	bad environment = iota
	dev
	test
	staging
	prod
)

var values = [...]string{"bad", "dev", "testing", "staging", "prod"}

func (e environment) String() string {
	return values[e]
}

func FromString(input string) (environment, error) {
	for i, v := range values {
		if v == input {
			return environment(i), nil
		}
	}
	return environment(0), fmt.Errorf("%s is not a possible value", input)
}

type config struct {
	port        int32
	environment string
}

func (c *config) check() {

}
