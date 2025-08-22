package models

import (
	"time"

	"gorm.io/gorm"
)

// ---------- Core Reference ----------

type Organization struct {
	OrgID       string  `gorm:"primaryKey;type:char(36);column:OrgID" json:"orgId"`
	Title       string  `gorm:"type:varchar(50);column:Title" json:"title"`
	ParentOrgID *string `gorm:"type:char(36);column:ParentOrgID" json:"parentOrgId,omitempty"`
	MetaSoft
	AuditTimes
}

func (Organization) TableName() string { return "Organization" }

type Site struct {
	SiteID string `gorm:"primaryKey;type:char(36);column:SiteID" json:"siteId"`
	Title  string `gorm:"type:varchar(50);column:Title" json:"title"`
	MetaSoft
	AuditTimes
}

func (Site) TableName() string { return "Site" }

type Page struct {
	PageID string  `gorm:"primaryKey;type:char(36);column:PageID" json:"pageId"`
	Title  string  `gorm:"type:varchar(50);column:Title" json:"title"`
	Status *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	SiteID *string `gorm:"type:char(36);column:SiteID" json:"siteId,omitempty"`
	MetaSoft
	AuditTimes
}

func (Page) TableName() string { return "Page" }

type Language struct {
	LangCode string `gorm:"primaryKey;type:char(3);column:LangCode" json:"langCode"`
	Title    string `gorm:"type:varchar(30);column:Title" json:"title"`
	MetaSoft
	AuditTimes
}

func (Language) TableName() string { return "Language" }

type Role struct {
	RoleID string  `gorm:"primaryKey;type:char(36);column:RoleID" json:"roleId"`
	Title  string  `gorm:"type:varchar(50);column:Title" json:"title"`
	Status *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (Role) TableName() string { return "Role" }

type User struct {
	UserID         string  `gorm:"primaryKey;type:char(36);column:UserID" json:"userId"`
	Email          string  `gorm:"type:varchar(100);column:Email" json:"email"`
	IAgreedToTerms *string `gorm:"type:char(1);column:IAgreedToTerms" json:"iAgreedToTerms,omitempty"`
	Status         *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (User) TableName() string { return "User" }

type UserRole struct {
	UserRoleID uint    `gorm:"primaryKey;autoIncrement;column:UserRoleID" json:"userRoleId"`
	RoleID     string  `gorm:"type:char(36);column:RoleID" json:"roleId"`
	UserID     string  `gorm:"type:char(36);column:UserID" json:"userId"`
	Status     *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (UserRole) TableName() string { return "UserRole" }

// ---------- Org / Property / Content ----------

type Property struct {
	PropertyID string  `gorm:"primaryKey;type:char(36);column:PropertyID" json:"propertyId"`
	Title      string  `gorm:"type:varchar(50);column:Title" json:"title"`
	OrgID      *string `gorm:"type:char(36);column:OrgID" json:"orgId,omitempty"`
	MetaSoft
	AuditTimes
}

func (Property) TableName() string { return "Property" }

type Offer struct {
	OfferID string  `gorm:"primaryKey;type:char(36);column:OfferID" json:"offerId"`
	Title   string  `gorm:"type:varchar(50);column:Title" json:"title"`
	OrgID   string  `gorm:"type:char(36);column:OrgID" json:"orgId"`
	Rules   *string `gorm:"type:mediumtext;column:Rules" json:"rules,omitempty"`
	Status  *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (Offer) TableName() string { return "Offer" }

type PaymentPlan struct {
	PaymentPlanID string  `gorm:"primaryKey;type:char(36);column:PaymentPlanID" json:"paymentPlanId"`
	OfferID       *string `gorm:"type:char(36);column:OfferID" json:"offerId,omitempty"`
	Rules         *string `gorm:"type:mediumtext;column:Rules" json:"rules,omitempty"`
	MetaSoft
	AuditTimes
}

func (PaymentPlan) TableName() string { return "PaymentPlan" }

type LeaseMod struct {
	LeaseModID string  `gorm:"primaryKey;type:char(36);column:LeaseModID" json:"leaseModId"`
	OfferID    *string `gorm:"type:char(36);column:OfferID" json:"offerId,omitempty"`
	Rules      *string `gorm:"type:mediumtext;column:Rules" json:"rules,omitempty"`
	MetaSoft
	AuditTimes
}

func (LeaseMod) TableName() string { return "LeaseMod" }

type PlanLeaseMod struct {
	PlanLeaseModID uint    `gorm:"primaryKey;autoIncrement;column:PlanLeaseModID"`
	PaymentPlanID  string  `gorm:"type:char(36);column:PaymentPlanID"`
	LeaseModID     string  `gorm:"type:char(36);column:LeaseModID"`
	Status         *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (PlanLeaseMod) TableName() string { return "PlanLeaseMod" }

type Content struct {
	ContentID   string  `gorm:"primaryKey;type:char(36);column:ContentID" json:"contentId"`
	Title       *string `gorm:"type:varchar(30);column:Title" json:"title,omitempty"`
	ContentText *string `gorm:"type:mediumtext;column:ContentText" json:"contentText,omitempty"`
	LangCode    *string `gorm:"type:char(3);column:LangCode" json:"langCode,omitempty"`
	OrgID       *string `gorm:"type:char(36);column:OrgID" json:"orgId,omitempty"`
	PageID      *string `gorm:"type:char(36);column:PageID" json:"pageId,omitempty"`
	MetaSoft
	AuditTimes
}

func (Content) TableName() string { return "Content" }

type OfferText struct {
	OfferTextID  string  `gorm:"primaryKey;type:char(36);column:OfferTextID" json:"offerTextId"`
	OfferContent *string `gorm:"type:mediumtext;column:OfferContent" json:"offerContent,omitempty"`
	LangCode     *string `gorm:"type:char(3);column:LangCode" json:"langCode,omitempty"`
	OfferID      *string `gorm:"type:char(36);column:OfferID" json:"offerId,omitempty"`
	MetaSoft
	AuditTimes
}

func (OfferText) TableName() string { return "OfferText" }

// ---------- People / Tenancy ----------

type Address struct {
	AddressID string  `gorm:"primaryKey;type:char(36);column:AddressID" json:"addressId"`
	OwnerID   *string `gorm:"type:char(36);column:OwnerID" json:"ownerId,omitempty"`
	Address1  *string `gorm:"type:varchar(50);column:Address1" json:"address1,omitempty"`
	Address2  *string `gorm:"type:varchar(50);column:Address2" json:"address2,omitempty"`
	Address3  *string `gorm:"type:varchar(50);column:Address3" json:"address3,omitempty"`
	City      *string `gorm:"type:varchar(30);column:City" json:"city,omitempty"`
	County    *string `gorm:"type:varchar(30);column:County" json:"county,omitempty"`
	StateCode *string `gorm:"type:char(2);column:StateCode" json:"stateCode,omitempty"`
	ZipCode   *string `gorm:"type:varchar(10);column:ZipCode" json:"zipCode,omitempty"`
	MetaSoft
	AuditTimes
}

func (Address) TableName() string { return "Address" }

type Tenant struct {
	TenantID   string  `gorm:"primaryKey;type:char(36);column:TenantID" json:"tenantId"`
	PropertyID *string `gorm:"type:char(36);column:PropertyID" json:"propertyId,omitempty"`
	AddressID  *string `gorm:"type:char(36);column:AddressID" json:"addressId,omitempty"`
	UserID     *string `gorm:"type:char(36);column:UserID" json:"userId,omitempty"`
	MetaSoft
	AuditTimes
}

func (Tenant) TableName() string { return "Tenant" }

type Account struct {
	AccountID string  `gorm:"primaryKey;type:char(36);column:AccountID" json:"accountId"`
	AddressID *string `gorm:"type:char(36);column:AddressID" json:"addressId,omitempty"`
	TenantID  *string `gorm:"type:char(36);column:TenantID" json:"tenantId,omitempty"`
	Unit      *string `gorm:"type:varchar(20);column:Unit" json:"unit,omitempty"`
	Status    *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (Account) TableName() string { return "Account" }

type Phone struct {
	PhoneID     string  `gorm:"primaryKey;type:char(36);column:PhoneID" json:"phoneId"`
	PhoneType   *string `gorm:"type:varchar(10);column:PhoneType" json:"phoneType,omitempty"`
	CountryCode *string `gorm:"type:varchar(3);column:CountryCode" json:"countryCode,omitempty"`
	Number      *string `gorm:"type:varchar(20);column:Number" json:"number,omitempty"`
	Extension   *string `gorm:"type:varchar(8);column:Extension" json:"extension,omitempty"`
	OwnerID     *string `gorm:"type:char(36);column:OwnerID" json:"ownerId,omitempty"`
	MetaSoft
	AuditTimes
}

func (Phone) TableName() string { return "Phone" }

// ---------- Payments / Events ----------

type Method struct {
	MethodID   string  `gorm:"primaryKey;type:char(36);column:MethodID" json:"methodId"`
	UserID     *string `gorm:"type:char(36);column:UserID" json:"userId,omitempty"`
	MethodType *string `gorm:"type:varchar(10);column:MethodType" json:"methodType,omitempty"`
	LastFour   *string `gorm:"type:char(4);column:LastFour" json:"lastFour,omitempty"`
	NameOnCard *string `gorm:"type:varchar(50);column:NameOnCard" json:"nameOnCard,omitempty"`
	Expiration *string `gorm:"type:char(4);column:Expiration" json:"expiration,omitempty"`
	CardNumber *string `gorm:"type:varchar(100);column:CardNumber" json:"cardNumber,omitempty"`
	Status     *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (Method) TableName() string { return "Method" }

type TransCode struct {
	TransCodeID string  `gorm:"primaryKey;type:char(3);column:TransCodeID" json:"transCodeId"`
	Title       string  `gorm:"type:varchar(30);column:Title" json:"title"`
	TransCode   *string `gorm:"type:varchar(10);column:TransCode" json:"transCode,omitempty"`
	MetaCommon
	AuditTimes
}

func (TransCode) TableName() string { return "TransCode" }

type Decision struct {
	DecisionID    string  `gorm:"primaryKey;type:char(36);column:DecisionID" json:"decisionId"`
	OfferID       *string `gorm:"type:char(36);column:OfferID" json:"offerId,omitempty"`
	PaymentPlanID *string `gorm:"type:char(36);column:PaymentPlanID" json:"paymentPlanId,omitempty"`
	Decision      *string `gorm:"type:varchar(20);column:Decision" json:"decision,omitempty"`
	Results       *string `gorm:"type:mediumtext;column:Results" json:"results,omitempty"`
	MetaSoft
	AuditTimes
}

func (Decision) TableName() string { return "Decision" }

type Acceptance struct {
	AcceptanceID   string     `gorm:"primaryKey;type:char(36);column:AcceptanceID" json:"acceptanceId"`
	DecisionID     *string    `gorm:"type:char(36);column:DecisionID" json:"decisionId,omitempty"`
	AcceptanceDate *time.Time `gorm:"type:date;column:AcceptanceDate" json:"acceptanceDate,omitempty"`
	OfferID        *string    `gorm:"type:char(36);column:OfferID" json:"offerId,omitempty"`
	AccountID      *string    `gorm:"type:char(36);column:AccountID" json:"accountId,omitempty"`
	PaymentPlanID  *string    `gorm:"type:char(36);column:PaymentPlanID" json:"paymentPlanId,omitempty"`
	Contract       *string    `gorm:"type:mediumtext;column:Contract" json:"contract,omitempty"`
	MetaSoft
	AuditTimes
}

func (Acceptance) TableName() string { return "Acceptance" }

type Payment struct {
	PaymentID     string     `gorm:"primaryKey;type:char(36);column:PaymentID" json:"paymentId"`
	AccountID     *string    `gorm:"type:char(36);column:AccountID" json:"accountId,omitempty"`
	PaymentPlanID *string    `gorm:"type:char(36);column:PaymentPlanID" json:"paymentPlanId,omitempty"`
	MethodID      *string    `gorm:"type:char(36);column:MethodID" json:"methodId,omitempty"`
	TransDate     *time.Time `gorm:"type:date;column:TransDate" json:"transDate,omitempty"`
	Memo          *string    `gorm:"type:varchar(100);column:Memo" json:"memo,omitempty"`
	Amount        *Money     `gorm:"type:decimal(10,4);column:Amount" json:"amount,omitempty"`
	Status        *string    `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (Payment) TableName() string { return "Payment" }

type Transaction struct {
	TransactionID string     `gorm:"primaryKey;type:char(36);column:TransactionID" json:"transactionId"`
	AccountID     *string    `gorm:"type:char(36);column:AccountID" json:"accountId,omitempty"`
	PaymentID     *string    `gorm:"type:char(36);column:PaymentID" json:"paymentId,omitempty"`
	TransDate     *time.Time `gorm:"type:date;column:TransDate" json:"transDate,omitempty"`
	TransCode     *string    `gorm:"type:char(3);column:TransCode" json:"transCode,omitempty"`
	Amount        *Money     `gorm:"type:decimal(10,4);column:Amount" json:"amount,omitempty"`
	MetaCommon
	AuditTimes
}

func (Transaction) TableName() string { return "Transaction" }

type EventType struct {
	EventTypeID string  `gorm:"primaryKey;type:char(36);column:EventTypeID" json:"eventTypeId"`
	Title       *string `gorm:"type:varchar(40);column:Title" json:"title,omitempty"`
	MetaSoft
	AuditTimes
}

func (EventType) TableName() string { return "EventType" }

type EventLog struct {
	EventID     uint       `gorm:"primaryKey;autoIncrement;column:EventID" json:"eventId"`
	EventTypeID *string    `gorm:"type:char(36);column:EventTypeID" json:"eventTypeId,omitempty"`
	EventDate   *time.Time `gorm:"type:datetime(6);column:EventDate" json:"eventDate,omitempty"`
	EventData   JSON       `gorm:"type:json;column:EventData" json:"eventData,omitempty"`
	Status      *string    `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	OwnerID     *string    `gorm:"type:char(36);column:OwnerID" json:"ownerId,omitempty"`
	MetaCommon
	AuditTimes
}

func (EventLog) TableName() string { return "EventLog" }

// ---------- Portfolio / Campaign / Workflow ----------

type Portfolio struct {
	PortfolioID    string  `gorm:"primaryKey;type:char(36);column:PortfolioID" json:"portfolioId"`
	OrgID          *string `gorm:"type:char(36);column:OrgID" json:"orgId,omitempty"`
	PortTemplateID *string `gorm:"type:char(36);column:PortTemplateID" json:"portTemplateId,omitempty"`
	Title          string  `gorm:"type:varchar(80);column:Title" json:"title"`
	Description    *string `gorm:"type:varchar(255);column:Description" json:"description,omitempty"`
	Status         *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (Portfolio) TableName() string { return "Portfolio" }

type PortTemplate struct {
	PortTemplateID string         `gorm:"primaryKey;type:char(36);column:PortTemplateID" json:"portTemplateId"`
	OrgID          *string        `gorm:"type:char(36);column:OrgID" json:"orgId,omitempty"`
	Title          string         `gorm:"type:varchar(80);column:Title" json:"title"`
	Description    *string        `gorm:"type:varchar(255);column:Description" json:"description,omitempty"`
	Selection      *string        `gorm:"type:varchar(1024);column:Selection" json:"selection,omitempty"`
	Status         *string        `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	UpdatedBy      *string        `gorm:"column:UpdatedBy;type:char(36)" json:"updatedBy,omitempty"`
	DeletedAt      gorm.DeletedAt `gorm:"column:DeletedAt;index" json:"-"`
	Cas            *uint32        `gorm:"column:Cas" json:"cas,omitempty"`
	CreatedAt      *time.Time     `gorm:"column:CreatedAt;type:datetime(6)" json:"createdAt,omitempty"`
	UpdatedAt      *time.Time     `gorm:"column:UpdatedAt;type:datetime(6)" json:"updatedAt,omitempty"`
}

func (PortTemplate) TableName() string { return "PortTemplate" }

type PortfolioAccount struct {
	PortfolioAccountID uint       `gorm:"primaryKey;autoIncrement;column:PortfolioAccountID" json:"portfolioAccountId"`
	PortfolioID        string     `gorm:"type:char(36);column:PortfolioID" json:"portfolioId"`
	AccountID          string     `gorm:"type:char(36);column:AccountID" json:"accountId"`
	AssignedAt         *time.Time `gorm:"type:datetime(6);column:AssignedAt" json:"assignedAt,omitempty"`
	RemovedAt          *time.Time `gorm:"type:datetime(6);column:RemovedAt" json:"removedAt,omitempty"`
	Status             *string    `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (PortfolioAccount) TableName() string { return "PortfolioAccount" }

type ActionType struct {
	ActionTypeID string `gorm:"primaryKey;type:varchar(20);column:ActionTypeID" json:"actionTypeId"`
	Title        string `gorm:"type:varchar(40);column:Title" json:"title"`
	MetaSoft
	AuditTimes
}

func (ActionType) TableName() string { return "ActionType" }

type WorkflowTemplate struct {
	WorkflowTemplateID string  `gorm:"primaryKey;type:char(36);column:WorkflowTemplateID" json:"workflowTemplateId"`
	OrgID              *string `gorm:"type:char(36);column:OrgID" json:"orgId,omitempty"`
	Title              string  `gorm:"type:varchar(80);column:Title" json:"title"`
	Description        *string `gorm:"type:varchar(255);column:Description" json:"description,omitempty"`
	SegmentRules       JSON    `gorm:"type:json;column:SegmentRules" json:"segmentRules,omitempty"`
	Status             *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (WorkflowTemplate) TableName() string { return "WorkflowTemplate" }

type WorkflowStep struct {
	WorkflowStepID     uint    `gorm:"primaryKey;autoIncrement;column:WorkflowStepID" json:"workflowStepId"`
	WorkflowTemplateID string  `gorm:"type:char(36);column:WorkflowTemplateID" json:"workflowTemplateId"`
	SeqNo              int     `gorm:"column:SeqNo" json:"seqNo"`
	ActionTypeID       string  `gorm:"type:varchar(20);column:ActionTypeID" json:"actionTypeId"`
	OffsetDays         int     `gorm:"column:OffsetDays" json:"offsetDays"`
	Payload            JSON    `gorm:"type:json;column:Payload" json:"payload,omitempty"`
	Status             *string `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	MetaSoft
	AuditTimes
}

func (WorkflowStep) TableName() string { return "WorkflowStep" }

type Campaign struct {
	CampaignID         string     `gorm:"primaryKey;type:char(36);column:CampaignID" json:"campaignId"`
	OrgID              *string    `gorm:"type:char(36);column:OrgID" json:"orgId,omitempty"`
	PortfolioID        string     `gorm:"type:char(36);column:PortfolioID" json:"portfolioId"`
	WorkflowTemplateID string     `gorm:"type:char(36);column:WorkflowTemplateID" json:"workflowTemplateId"`
	Title              string     `gorm:"type:varchar(80);column:Title" json:"title"`
	Description        *string    `gorm:"type:varchar(255);column:Description" json:"description,omitempty"`
	CampaignType       *string    `gorm:"type:varchar(20);column:CampaignType" json:"campaignType,omitempty"`
	StartDate          *time.Time `gorm:"type:date;column:StartDate" json:"startDate,omitempty"`
	EndDate            *time.Time `gorm:"type:date;column:EndDate" json:"endDate,omitempty"`
	Status             *string    `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	Limits             JSON       `gorm:"type:json;column:Limits" json:"limits,omitempty"`
	MetaSoft
	AuditTimes
}

func (Campaign) TableName() string { return "Campaign" }

type CampaignAction struct {
	CampaignActionID string     `gorm:"primaryKey;type:char(36);column:CampaignActionID" json:"campaignActionId"`
	CampaignID       string     `gorm:"type:char(36);column:CampaignID" json:"campaignId"`
	WorkflowStepID   uint       `gorm:"column:WorkflowStepID" json:"workflowStepId"`
	AccountID        *string    `gorm:"type:char(36);column:AccountID" json:"accountId,omitempty"`
	OwnerID          *string    `gorm:"type:char(36);column:OwnerID" json:"ownerId,omitempty"`
	ActionTypeID     string     `gorm:"type:varchar(20);column:ActionTypeID" json:"actionTypeId"`
	ScheduledAt      *time.Time `gorm:"type:datetime(6);column:ScheduledAt" json:"scheduledAt,omitempty"`
	ExecutedAt       *time.Time `gorm:"type:datetime(6);column:ExecutedAt" json:"executedAt,omitempty"`
	Status           *string    `gorm:"type:char(1);column:Status" json:"status,omitempty"`
	Result           JSON       `gorm:"type:json;column:Result" json:"result,omitempty"`
	EventID          *uint      `gorm:"column:EventID" json:"eventId,omitempty"`
	MetaSoft
	AuditTimes
}

func (CampaignAction) TableName() string { return "CampaignAction" }

// ---------- Views (read-only helpers) ----------

// VPortfolioAccounts maps to the view v_portfolio_accounts
// Useful for read queries and reporting. Not intended for inserts/updates.
type VPortfolioAccounts struct {
	PortfolioID         string     `gorm:"column:PortfolioID" json:"portfolioId"`
	PortfolioTitle      string     `gorm:"column:PortfolioTitle" json:"portfolioTitle"`
	OrgID               *string    `gorm:"column:OrgID" json:"orgId,omitempty"`
	AccountID           string     `gorm:"column:AccountID" json:"accountId"`
	TenantID            *string    `gorm:"column:TenantID" json:"tenantId,omitempty"`
	AccountStatus       *string    `gorm:"column:AccountStatus" json:"accountStatus,omitempty"`
	AssignedAt          *time.Time `gorm:"column:AssignedAt" json:"assignedAt,omitempty"`
	RemovedAt           *time.Time `gorm:"column:RemovedAt" json:"removedAt,omitempty"`
	IsActiveInPortfolio bool       `gorm:"column:IsActiveInPortfolio" json:"isActiveInPortfolio"`
}

func (VPortfolioAccounts) TableName() string { return "v_portfolio_accounts" }

// VCampaignSchedule maps to the view v_campaign_schedule
type VCampaignSchedule struct {
	CampaignID       string     `gorm:"column:CampaignID" json:"campaignId"`
	CampaignTitle    string     `gorm:"column:CampaignTitle" json:"campaignTitle"`
	PortfolioID      string     `gorm:"column:PortfolioID" json:"portfolioId"`
	StartDate        *time.Time `gorm:"column:StartDate" json:"startDate,omitempty"`
	CampaignStatus   *string    `gorm:"column:CampaignStatus" json:"campaignStatus,omitempty"`
	CampaignActionID string     `gorm:"column:CampaignActionID" json:"campaignActionId"`
	WorkflowStepID   uint       `gorm:"column:WorkflowStepID" json:"workflowStepId"`
	AccountID        *string    `gorm:"column:AccountID" json:"accountId,omitempty"`
	OwnerID          *string    `gorm:"column:OwnerID" json:"ownerId,omitempty"`
	ActionTypeID     string     `gorm:"column:ActionTypeID" json:"actionTypeId"`
	ScheduledAt      *time.Time `gorm:"column:ScheduledAt" json:"scheduledAt,omitempty"`
	ExecutedAt       *time.Time `gorm:"column:ExecutedAt" json:"executedAt,omitempty"`
	ActionStatus     *string    `gorm:"column:ActionStatus" json:"actionStatus,omitempty"`
}

func (VCampaignSchedule) TableName() string { return "v_campaign_schedule" }
