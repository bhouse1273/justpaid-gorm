package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreateUserAndGet(t *testing.T) {
	h, _ := NewTestHandler(t)
	e := echo.New()
	e.POST("/users", h.CreateUser)
	e.GET("/users/:id", h.GetUser)

	// create
	payload := map[string]interface{}{"email": "test@example.com"}
	b, _ := json.Marshal(payload)
	rec := EchoRequest(e, http.MethodPost, "/users", bytes.NewBuffer(b))
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 creating user, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// parse returned id
	var created map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to parse create response: %v", err)
	}
	id, ok := created["userId"].(string)
	if !ok || id == "" {
		t.Fatalf("unexpected userId in response: %v", created)
	}

	// get
	rec = EchoRequest(e, http.MethodGet, "/users/"+id, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 getting user, got %d, body=%s", rec.Code, rec.Body.String())
	}
}
