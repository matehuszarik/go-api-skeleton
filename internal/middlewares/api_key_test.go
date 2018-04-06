package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var APIKeys = []string{"JB7i&buy3wHBe", "90n&vbi76gRGH"}

func createRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(APIKey(APIKeys))
	r.GET("/foo", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	return r
}
func TestAPIKeyMissing(t *testing.T) {
	r := createRouter()

	req, _ := http.NewRequest(http.MethodGet, "/foo", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusUnauthorized)
}

func TestAPIKeyInvalid(t *testing.T) {
	r := createRouter()

	req, _ := http.NewRequest(http.MethodGet, "/foo", nil)
	req.Header.Add(HeaderAPIKey, "abcdefghijkl")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusUnauthorized)
}

func TestAPIKeyValid(t *testing.T) {
	r := createRouter()

	req, _ := http.NewRequest(http.MethodGet, "/foo", nil)
	req.Header.Add(HeaderAPIKey, APIKeys[1])
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}
