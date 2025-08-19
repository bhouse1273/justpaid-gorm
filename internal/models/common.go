package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

// Money uses DECIMAL(10,4) in MySQL. shopspring/decimal implements sql.Scanner/driver.Valuer.
type Money = decimal.Decimal

// JSON maps to MySQL JSON columns.
type JSON = datatypes.JSON

// DATETIME audit mixin (uses MySQL DEFAULT/ON UPDATE to auto-populate)
type AuditTimes struct {
	CreatedAt time.Time `gorm:"column:CreatedAt;type:datetime(6);autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt;type:datetime(6);autoUpdateTime"`
}

// Meta with UpdatedBy, Deleted, Cas
type MetaSoft struct {
	UpdatedBy *string `gorm:"column:UpdatedBy;type:char(36)"`
	Deleted   *string `gorm:"column:Deleted;type:char(1)"`
	Cas       *uint32 `gorm:"column:Cas"`
}

// Meta without Deleted (for a few tables like TransCode, EventLog, Transaction)
type MetaCommon struct {
	UpdatedBy *string `gorm:"column:UpdatedBy;type:char(36)"`
	Cas       *uint32 `gorm:"column:Cas"`
}
