package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreateOrganizationValidation(t *testing.T) {
	h, _ := NewTestHandler(t)
	e := echo.New()
	e.POST("/organizations", h.CreateOrganization)

	// Missing title -> should 400
	body := bytes.NewBufferString(`{}`)
	rec := EchoRequest(e, http.MethodPost, "/organizations", body)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for missing title, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// Valid title -> 201
	payload := map[string]interface{}{"title": "Acme"}
	b, _ := json.Marshal(payload)
	rec = EchoRequest(e, http.MethodPost, "/organizations", bytes.NewBuffer(b))
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 for valid create, got %d, body=%s", rec.Code, rec.Body.String())
	}
}
