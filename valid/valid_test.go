package valid

import (
	"github.com/stretchr/testify/assert"
	"github.com/willbtlr/w-lib/marshallable"
	"testing"
	"time"
)

type Conf struct {
	Timeout marshallable.Duration `default:"1s"`
	BaseURL string                `validate:"required"`
	Key     string                `default:"foo" validate:"required"`
}

func Test_Conform(t *testing.T) {
	c := Conf{BaseURL: "http"}
	err := Conform(&c)
	assert.Nil(t, err)
	assert.Equal(t, 1*time.Second, c.Timeout.Duration)
	assert.Equal(t, "foo", c.Key)
	assert.Equal(t, "http", c.BaseURL)
}
