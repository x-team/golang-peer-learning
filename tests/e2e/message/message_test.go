package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		fmt.Println(err)
	}

	return "Bearer " + tokenString
}

func TestMessageTestE2E(t *testing.T) {
	t.Run("When user is not authenticated", func(t *testing.T) {
		client := resty.New()

		resp, err := client.R().
			SetBody(`{"content": "MESSAGE FROM POST","toPhoneNumber": "+5541996740459"}`).
			Post("http://localhost:4000/api/v1/message")

		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})

	t.Run("When user is authenticated with an invalid token", func(t *testing.T) {
		client := resty.New()

		resp, err := client.R().
			SetHeader("Authorization", "bearer invalid_token").
			SetBody(`{"content": "MESSAGE FROM POST","toPhoneNumber": "+5541996740459"}`).
			Post("http://localhost:4000/api/v1/message")

		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})

	t.Run("Creates a new Message", func(t *testing.T) {
		client := resty.New()

		resp, err := client.R().
			SetHeader("Authorization", createToken()).
			SetBody(`{"content": "MESSAGE FROM POST","toPhoneNumber": "+5541996740459"}`).
			Post("http://localhost:4000/api/v1/message")

		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})
}
