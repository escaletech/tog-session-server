package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/escaletech/tog-go/sessions"
	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

	"github.com/escaletech/tog-session-server/internal/config"
	"github.com/escaletech/tog-session-server/internal/handler"
)

var ctx = context.Background()

func TestHandler(t *testing.T) {
	defaultConfig := config.Config{RedisURL: "redis://localhost/2"}

	t.Run("without path prefix", testScenario(
		defaultConfig,
		"/some-ns/abc123",
		SessionParams{"some-ns", "abc123", sessions.SessionOptions{}},
	))

	t.Run("with path prefix", testScenario(
		config.Config{RedisURL: defaultConfig.RedisURL, PathPrefix: "/_sessions"},
		"/_sessions/some-ns/abc123",
		SessionParams{"some-ns", "abc123", sessions.SessionOptions{}},
	))

	t.Run("forcing enabled flags", testScenario(
		defaultConfig,
		"/some-ns/abc123?enable=one&enable=two",
		SessionParams{"some-ns", "abc123", sessions.SessionOptions{
			Force: sessions.Session{"one": true, "two": true},
		}},
	))

	t.Run("forcing disabled flags", testScenario(
		defaultConfig,
		"/some-ns/abc123?disable=one&disable=two",
		SessionParams{"some-ns", "abc123", sessions.SessionOptions{
			Force: sessions.Session{"one": false, "two": false},
		}},
	))

	t.Run("passing session traits", testScenario(
		defaultConfig,
		"/some-ns/abc123?traits=early-adopter",
		SessionParams{"some-ns", "abc123", sessions.SessionOptions{
			Traits: []string{"early-adopter"},
		}},
	))
}

type SessionParams struct {
	Namespace string
	SessionID string
	Options   sessions.SessionOptions
}

func testScenario(conf config.Config, uri string, expectedCall SessionParams) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()
		// Arrange
		sc := new(fakeSessionClient)

		app := fiber.New()
		handler.RegisterWithClient(app, conf, sc)

		expected := sessions.Session{"foo": true, "bar": false}
		sc.result = expected

		// Act
		res, err := app.Test(httptest.NewRequest("GET", "http://localhost"+uri, nil))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		// Assert
		assertResponse(t, res, http.StatusOK, expected)

		assert.Equal(t, expectedCall.Namespace, sc.params[0])
		assert.Equal(t, expectedCall.SessionID, sc.params[1])
		assert.Equal(t, expectedCall.Options, sc.params[2])
	}
}

func assertResponse(t *testing.T, res *http.Response, expectedStatus int, expectedBody sessions.Session) {
	t.Helper()
	assert.Equal(t, expectedStatus, res.StatusCode)
	var actualBody sessions.Session
	json.NewDecoder(res.Body).Decode(&actualBody)
	assert.Equal(t, expectedBody, actualBody)
}

type fakeSessionClient struct {
	params []interface{}
	result sessions.Session
}

func (c *fakeSessionClient) Session(ctx context.Context, ns, sessionID string, opt *sessions.SessionOptions) sessions.Session {
	c.params = []interface{}{ns, sessionID, *opt}
	return c.result
}
