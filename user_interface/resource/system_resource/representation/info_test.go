package representation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfoMarshal(t *testing.T) {
	info := Info{
		Url:             "localhost",
		Version:         "4.0.0",
		CpuArchitecture: 64,
	}

	data, err := json.Marshal(info)

	assert.Nil(t, err)
	assert.Equal(t, `{"url":"localhost","version":"4.0.0","processorArchitecture":64}`, string(data))
}
