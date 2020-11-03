package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	welcomeJSON = `{"message":"Welcome to Yamabiko!!"}` + "\n"
	messageJSON = `{"message":"Hey"}` + "\n"
)

func TestGetWelcome(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")
	h := &Handler{}

	// Assertions
	if assert.NoError(t, h.GetWelcome(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, welcomeJSON, rec.Body.String())
	}
}

func TestPostEcho(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/echo", strings.NewReader(messageJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{}

	// Assertions
	if assert.NoError(t, h.PostEcho(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, messageJSON, rec.Body.String())
	}
}
