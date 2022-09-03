package valid

import (
	"fmt"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator"
	"os"
)

var customValidations = map[string]validator.Func{}

type conformConf struct {
	Validator *validator.Validate
}

type conformOpt func(opts *conformConf)

func WithValidator(v *validator.Validate) conformOpt {
	return func(opts *conformConf) {
		opts.Validator = v
	}
}

func MustConform(i interface{}, opts ...conformOpt) {
	err := Conform(i, opts...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func Conform(i interface{}, opts ...conformOpt) error {
	conf := conformConf{}
	for _, opt := range opts {
		opt(&conf)
	}

	var validate *validator.Validate
	if conf.Validator != nil {
		validate = conf.Validator
	} else {
		validate = validator.New()
	}

	for tagName, f := range customValidations {
		err := validate.RegisterValidation(tagName, f)
		if err != nil {
			return err
		}
	}

	err := defaults.Set(i)
	if err != nil {
		return err
	}

	return validate.Struct(i)
}
