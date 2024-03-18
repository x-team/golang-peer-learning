package tests

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckEndpoint(t *testing.T) {
	client := resty.New()

	t.Run("healthcheck", func(t *testing.T) {
		resp, err := client.R().Get("http://localhost:4000/healthcheck")
		assert.NoError(t, err)

		assert.Equal(t, 200, resp.StatusCode())
	})
}
