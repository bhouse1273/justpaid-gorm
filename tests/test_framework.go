package tests

import (
	"context"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/bhouse1273/justpaid-gorm/internal/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewTestDB returns an in-memory sqlite DB and auto-migrates minimal models used by tests.
func NewTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	// Create minimal tables with SQLite-compatible DDL matching the models
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS Organization (OrgID TEXT PRIMARY KEY, Title TEXT, ParentOrgID TEXT, CreatedAt DATETIME, UpdatedAt DATETIME, UpdatedBy TEXT, DeletedAt DATETIME, Cas INTEGER);`,
		`CREATE TABLE IF NOT EXISTS "User" (UserID TEXT PRIMARY KEY, Email TEXT, IAgreedToTerms TEXT, Status TEXT, CreatedAt DATETIME, UpdatedAt DATETIME, UpdatedBy TEXT, DeletedAt DATETIME, Cas INTEGER);`,
		`CREATE TABLE IF NOT EXISTS Payment (PaymentID TEXT PRIMARY KEY, AccountID TEXT, PaymentPlanID TEXT, MethodID TEXT, TransDate DATETIME, Memo TEXT, Amount TEXT, Status TEXT, UpdatedBy TEXT, DeletedAt DATETIME, Cas INTEGER, CreatedAt DATETIME, UpdatedAt DATETIME);`,
		`CREATE TABLE IF NOT EXISTS Property (PropertyID TEXT PRIMARY KEY, Title TEXT, OrgID TEXT, CreatedAt DATETIME, UpdatedAt DATETIME, UpdatedBy TEXT, DeletedAt DATETIME, Cas INTEGER);`,
		`CREATE TABLE IF NOT EXISTS Offer (OfferID TEXT PRIMARY KEY, Title TEXT, OrgID TEXT, Rules TEXT, Status TEXT, CreatedAt DATETIME, UpdatedAt DATETIME, UpdatedBy TEXT, DeletedAt DATETIME, Cas INTEGER);`,
		`CREATE TABLE IF NOT EXISTS Account (AccountID TEXT PRIMARY KEY, AddressID TEXT, TenantID TEXT, Unit TEXT, Status TEXT, CreatedAt DATETIME, UpdatedAt DATETIME, UpdatedBy TEXT, DeletedAt DATETIME, Cas INTEGER);`,
	}
	for _, s := range stmts {
		if err := db.Exec(s).Error; err != nil {
			t.Fatalf("failed to create table: %v", err)
		}
	}
	return db
}

// EchoRequest executes an HTTP request against an Echo handler and returns recorder
func EchoRequest(e *echo.Echo, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec
}

// NewTestHandler constructs a handlers.Handler wired to an in-memory DB
func NewTestHandler(t *testing.T) (*handlers.Handler, *gorm.DB) {
	db := NewTestDB(t)
	return handlers.New(db), db
}

// Helper to run a subtest with context
func WithCtx(t *testing.T, fn func(ctx context.Context)) {
	ctx := context.Background()
	fn(ctx)
}
