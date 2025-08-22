package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPropertyCRUD(t *testing.T) {
	h, _ := NewTestHandler(t)
	e := echo.New()
	e.POST("/properties", h.CreateProperty)
	e.GET("/properties/:id", h.GetProperty)
	e.PUT("/properties/:id", h.UpdateProperty)
	e.DELETE("/properties/:id", h.DeleteProperty)

	// create
	payload := map[string]interface{}{"title": "My Property"}
	b, _ := json.Marshal(payload)
	rec := EchoRequest(e, http.MethodPost, "/properties", bytes.NewBuffer(b))
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 creating property, got %d, body=%s", rec.Code, rec.Body.String())
	}
	var created map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to parse create response: %v", err)
	}
	id, ok := created["propertyId"].(string)
	if !ok || id == "" {
		t.Fatalf("unexpected propertyId in response: %v", created)
	}

	// get
	rec = EchoRequest(e, http.MethodGet, "/properties/"+id, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 getting property, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// update
	up := map[string]interface{}{"title": "My Property v2"}
	ub, _ := json.Marshal(up)
	rec = EchoRequest(e, http.MethodPut, "/properties/"+id, bytes.NewBuffer(ub))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 updating property, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// delete
	rec = EchoRequest(e, http.MethodDelete, "/properties/"+id, nil)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected 204 deleting property, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// get after delete -> not found (handler returns NoContent 404)
	rec = EchoRequest(e, http.MethodGet, "/properties/"+id, nil)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404 after delete, got %d, body=%s", rec.Code, rec.Body.String())
	}
}

func TestOfferCRUD(t *testing.T) {
	h, _ := NewTestHandler(t)
	e := echo.New()
	// need org for offer
	e.POST("/organizations", h.CreateOrganization)
	e.POST("/offers", h.CreateOffer)
	e.GET("/offers/:id", h.GetOffer)
	e.PUT("/offers/:id", h.UpdateOffer)
	e.DELETE("/offers/:id", h.DeleteOffer)

	// create org
	org := map[string]interface{}{"title": "Org For Offer"}
	ob, _ := json.Marshal(org)
	rec := EchoRequest(e, http.MethodPost, "/organizations", bytes.NewBuffer(ob))
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 creating org, got %d, body=%s", rec.Code, rec.Body.String())
	}
	var ocreated map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &ocreated); err != nil {
		t.Fatalf("failed to parse org create response: %v", err)
	}
	orgId, ok := ocreated["orgId"].(string)
	if !ok || orgId == "" {
		t.Fatalf("unexpected orgId in response: %v", ocreated)
	}

	// create offer
	offer := map[string]interface{}{"title": "Special Offer", "orgId": orgId}
	fb, _ := json.Marshal(offer)
	rec = EchoRequest(e, http.MethodPost, "/offers", bytes.NewBuffer(fb))
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 creating offer, got %d, body=%s", rec.Code, rec.Body.String())
	}
	var fcreated map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &fcreated); err != nil {
		t.Fatalf("failed to parse offer create response: %v", err)
	}
	id, ok := fcreated["offerId"].(string)
	if !ok || id == "" {
		t.Fatalf("unexpected offerId in response: %v", fcreated)
	}

	// get
	rec = EchoRequest(e, http.MethodGet, "/offers/"+id, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 getting offer, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// update
	up := map[string]interface{}{"rules": "new rules"}
	ub, _ := json.Marshal(up)
	rec = EchoRequest(e, http.MethodPut, "/offers/"+id, bytes.NewBuffer(ub))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 updating offer, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// delete
	rec = EchoRequest(e, http.MethodDelete, "/offers/"+id, nil)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected 204 deleting offer, got %d, body=%s", rec.Code, rec.Body.String())
	}
}

func TestAccountCRUD(t *testing.T) {
	h, _ := NewTestHandler(t)
	e := echo.New()
	e.POST("/accounts", h.CreateAccount)
	e.GET("/accounts/:id", h.GetAccount)
	e.PUT("/accounts/:id", h.UpdateAccount)
	e.DELETE("/accounts/:id", h.DeleteAccount)

	// create
	payload := map[string]interface{}{"unit": "101"}
	b, _ := json.Marshal(payload)
	rec := EchoRequest(e, http.MethodPost, "/accounts", bytes.NewBuffer(b))
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 creating account, got %d, body=%s", rec.Code, rec.Body.String())
	}
	var created map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to parse create response: %v", err)
	}
	id, ok := created["accountId"].(string)
	if !ok || id == "" {
		t.Fatalf("unexpected accountId in response: %v", created)
	}

	// update status
	up := map[string]interface{}{"status": "A"}
	ub, _ := json.Marshal(up)
	rec = EchoRequest(e, http.MethodPut, "/accounts/"+id, bytes.NewBuffer(ub))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 updating account, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// delete
	rec = EchoRequest(e, http.MethodDelete, "/accounts/"+id, nil)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected 204 deleting account, got %d, body=%s", rec.Code, rec.Body.String())
	}
}

func TestPaymentCreateGet(t *testing.T) {
	h, _ := NewTestHandler(t)
	e := echo.New()
	e.POST("/payments", h.CreatePayment)
	e.GET("/payments/:id", h.GetPayment)

	// create
	payload := map[string]interface{}{"amount": "12.34"}
	b, _ := json.Marshal(payload)
	rec := EchoRequest(e, http.MethodPost, "/payments", bytes.NewBuffer(b))
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 creating payment, got %d, body=%s", rec.Code, rec.Body.String())
	}
	var created map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to parse create response: %v", err)
	}
	id, ok := created["paymentId"].(string)
	if !ok || id == "" {
		t.Fatalf("unexpected paymentId in response: %v", created)
	}

	// get
	rec = EchoRequest(e, http.MethodGet, "/payments/"+id, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 getting payment, got %d, body=%s", rec.Code, rec.Body.String())
	}
}
