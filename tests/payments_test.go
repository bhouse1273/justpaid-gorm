package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreatePaymentAmountParsing(t *testing.T) {
	h, _ := NewTestHandler(t)
	e := echo.New()
	e.POST("/payments", h.CreatePayment)

	// invalid amount
	payload := map[string]interface{}{"amount": "not-a-number"}
	b, _ := json.Marshal(payload)
	rec := EchoRequest(e, http.MethodPost, "/payments", bytes.NewBuffer(b))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for invalid amount, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// valid amount
	payload = map[string]interface{}{"amount": "123.45"}
	b, _ = json.Marshal(payload)
	rec = EchoRequest(e, http.MethodPost, "/payments", bytes.NewBuffer(b))
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 for valid amount, got %d, body=%s", rec.Code, rec.Body.String())
	}
}
