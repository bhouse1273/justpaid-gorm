package handlers

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Organization DTOs
type OrganizationCreateDTO struct {
	Title       string  `json:"title" validate:"required"`
	ParentOrgID *string `json:"parentOrgId,omitempty"`
}

type OrganizationUpdateDTO struct {
	Title       *string `json:"title,omitempty"`
	ParentOrgID *string `json:"parentOrgId,omitempty"`
}

// Property DTOs
type PropertyCreateDTO struct {
	Title string  `json:"title" validate:"required"`
	OrgID *string `json:"orgId,omitempty"`
}
type PropertyUpdateDTO struct {
	Title *string `json:"title,omitempty"`
	OrgID *string `json:"orgId,omitempty"`
}

// Offer DTOs
type OfferCreateDTO struct {
	Title string  `json:"title" validate:"required"`
	OrgID string  `json:"orgId" validate:"required"`
	Rules *string `json:"rules,omitempty"`
}
type OfferUpdateDTO struct {
	Title  *string `json:"title,omitempty"`
	Rules  *string `json:"rules,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Portfolio DTOs
type PortfolioCreateDTO struct {
	Title          string  `json:"title" validate:"required"`
	OrgID          *string `json:"orgId,omitempty"`
	PortTemplateID *string `json:"portTemplateId,omitempty"`
	Description    *string `json:"description,omitempty"`
}
type PortfolioUpdateDTO struct {
	Title          *string `json:"title,omitempty"`
	PortTemplateID *string `json:"portTemplateId,omitempty"`
	Description    *string `json:"description,omitempty"`
	Status         *string `json:"status,omitempty"`
}

// PortTemplate DTOs
type PortTemplateCreateDTO struct {
	Title       string  `json:"title" validate:"required"`
	OrgID       *string `json:"orgId,omitempty"`
	Description *string `json:"description,omitempty"`
	Selection   *string `json:"selection,omitempty"`
}
type PortTemplateUpdateDTO struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Selection   *string `json:"selection,omitempty"`
	Status      *string `json:"status,omitempty"`
}

// Campaign DTOs
type CampaignCreateDTO struct {
	PortfolioID        string     `json:"portfolioId" validate:"required"`
	WorkflowTemplateID string     `json:"workflowTemplateId" validate:"required"`
	Title              string     `json:"title" validate:"required"`
	Description        *string    `json:"description,omitempty"`
	CampaignType       *string    `json:"campaignType,omitempty"`
	StartDate          *time.Time `json:"startDate,omitempty"`
	EndDate            *time.Time `json:"endDate,omitempty"`
}
type CampaignUpdateDTO struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Status      *string    `json:"status,omitempty"`
}

// Account DTOs
type AccountCreateDTO struct {
	AddressID *string `json:"addressId,omitempty"`
	TenantID  *string `json:"tenantId,omitempty"`
	Unit      *string `json:"unit,omitempty"`
}
type AccountUpdateDTO struct {
	AddressID *string `json:"addressId,omitempty"`
	TenantID  *string `json:"tenantId,omitempty"`
	Unit      *string `json:"unit,omitempty"`
	Status    *string `json:"status,omitempty"`
}

// Payment DTOs
type PaymentCreateDTO struct {
	AccountID     *string    `json:"accountId,omitempty"`
	PaymentPlanID *string    `json:"paymentPlanId,omitempty"`
	MethodID      *string    `json:"methodId,omitempty"`
	TransDate     *time.Time `json:"transDate,omitempty"`
	Memo          *string    `json:"memo,omitempty"`
	Amount        *string    `json:"amount,omitempty"` // string to accept decimal input
}
type PaymentUpdateDTO struct {
	Memo   *string `json:"memo,omitempty"`
	Status *string `json:"status,omitempty"`
}

// User DTOs
type UserCreateDTO struct {
	Email string `json:"email" validate:"required,email"`
}
type UserUpdateDTO struct {
	Email  *string `json:"email,omitempty"`
	Status *string `json:"status,omitempty"`
}
