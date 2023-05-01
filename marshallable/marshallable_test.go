package marshallable

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type conf struct {
	Duration Duration `json:"duration"`
}

func Test_Unmarshal(t *testing.T) {
	const s = `{"duration": "1s"}`
	var c conf
	err := json.Unmarshal([]byte(s), &c)
	assert.Nil(t, err)

	assert.Equal(t, Duration{1 * time.Second}, c.Duration)
}

func Test_Marshal(t *testing.T) {
	c := conf{
		Duration: Duration{1 * time.Second},
	}

	bs, err := json.Marshal(c)
	assert.Nil(t, err)

	assert.Equal(t, `{"duration":"1s"}`, string(bs))

	var c2 conf
	err = json.Unmarshal(bs, &c2)
	assert.Nil(t, err)
	assert.Equal(t, c, c2)
}
