package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Money uses DECIMAL(10,4) in MySQL. shopspring/decimal implements sql.Scanner/driver.Valuer.
type Money = decimal.Decimal

// JSON maps to MySQL JSON columns.
type JSON = datatypes.JSON

// DATETIME audit mixin (uses MySQL DEFAULT/ON UPDATE to auto-populate)
type AuditTimes struct {
	CreatedAt time.Time `gorm:"column:CreatedAt;type:datetime(6);autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt;type:datetime(6);autoUpdateTime" json:"updatedAt"`
}

// Meta with UpdatedBy, soft-delete (DeletedAt), Cas
type MetaSoft struct {
	UpdatedBy *string        `gorm:"column:UpdatedBy;type:char(36)" json:"updatedBy,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:DeletedAt;index" json:"-"`
	Cas       *uint32        `gorm:"column:Cas" json:"cas,omitempty"`
}

// Meta without Deleted (for a few tables like TransCode, EventLog, Transaction)
type MetaCommon struct {
	UpdatedBy *string `gorm:"column:UpdatedBy;type:char(36)" json:"updatedBy,omitempty"`
	Cas       *uint32 `gorm:"column:Cas" json:"cas,omitempty"`
}
