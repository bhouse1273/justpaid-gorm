package models

import "time"

// ---------- Core Reference ----------

type Organization struct {
	OrgID       string  `gorm:"primaryKey;type:char(36);column:OrgID"`
	Title       string  `gorm:"type:varchar(50);column:Title"`
	ParentOrgID *string `gorm:"type:char(36);column:ParentOrgID"`
	MetaSoft
	AuditTimes
}

func (Organization) TableName() string { return "Organization" }

type Site struct {
	SiteID string `gorm:"primaryKey;type:char(36);column:SiteID"`
	Title  string `gorm:"type:varchar(50);column:Title"`
	MetaSoft
	AuditTimes
}

func (Site) TableName() string { return "Site" }

type Page struct {
	PageID string  `gorm:"primaryKey;type:char(36);column:PageID"`
	Title  string  `gorm:"type:varchar(50);column:Title"`
	Status *string `gorm:"type:char(1);column:Status"`
	SiteID *string `gorm:"type:char(36);column:SiteID"`
	MetaSoft
	AuditTimes
}

func (Page) TableName() string { return "Page" }

type Language struct {
	LangCode string `gorm:"primaryKey;type:char(3);column:LangCode"`
	Title    string `gorm:"type:varchar(30);column:Title"`
	MetaSoft
	AuditTimes
}

func (Language) TableName() string { return "Language" }

type Role struct {
	RoleID string  `gorm:"primaryKey;type:char(36);column:RoleID"`
	Title  string  `gorm:"type:varchar(50);column:Title"`
	Status *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (Role) TableName() string { return "Role" }

type User struct {
	UserID         string  `gorm:"primaryKey;type:char(36);column:UserID"`
	Email          string  `gorm:"type:varchar(100);column:Email"`
	IAgreedToTerms *string `gorm:"type:char(1);column:IAgreedToTerms"`
	Status         *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (User) TableName() string { return "User" }

type UserRole struct {
	UserRoleID uint    `gorm:"primaryKey;autoIncrement;column:UserRoleID"`
	RoleID     string  `gorm:"type:char(36);column:RoleID"`
	UserID     string  `gorm:"type:char(36);column:UserID"`
	Status     *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (UserRole) TableName() string { return "UserRole" }

// ---------- Org / Property / Content ----------

type Property struct {
	PropertyID string  `gorm:"primaryKey;type:char(36);column:PropertyID"`
	Title      string  `gorm:"type:varchar(50);column:Title"`
	OrgID      *string `gorm:"type:char(36);column:OrgID"`
	MetaSoft
	AuditTimes
}

func (Property) TableName() string { return "Property" }

type Offer struct {
	OfferID string  `gorm:"primaryKey;type:char(36);column:OfferID"`
	Title   string  `gorm:"type:varchar(50);column:Title"`
	OrgID   string  `gorm:"type:char(36);column:OrgID"`
	Rules   *string `gorm:"type:mediumtext;column:Rules"`
	Status  *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (Offer) TableName() string { return "Offer" }

type PaymentPlan struct {
	PaymentPlanID string  `gorm:"primaryKey;type:char(36);column:PaymentPlanID"`
	OfferID       *string `gorm:"type:char(36);column:OfferID"`
	Rules         *string `gorm:"type:mediumtext;column:Rules"`
	MetaSoft
	AuditTimes
}

func (PaymentPlan) TableName() string { return "PaymentPlan" }

type LeaseMod struct {
	LeaseModID string  `gorm:"primaryKey;type:char(36);column:LeaseModID"`
	OfferID    *string `gorm:"type:char(36);column:OfferID"`
	Rules      *string `gorm:"type:mediumtext;column:Rules"`
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
	ContentID   string  `gorm:"primaryKey;type:char(36);column:ContentID"`
	Title       *string `gorm:"type:varchar(30);column:Title"`
	ContentText *string `gorm:"type:mediumtext;column:ContentText"`
	LangCode    *string `gorm:"type:char(3);column:LangCode"`
	OrgID       *string `gorm:"type:char(36);column:OrgID"`
	PageID      *string `gorm:"type:char(36);column:PageID"`
	MetaSoft
	AuditTimes
}

func (Content) TableName() string { return "Content" }

type OfferText struct {
	OfferTextID  string  `gorm:"primaryKey;type:char(36);column:OfferTextID"`
	OfferContent *string `gorm:"type:mediumtext;column:OfferContent"`
	LangCode     *string `gorm:"type:char(3);column:LangCode"`
	OfferID      *string `gorm:"type:char(36);column:OfferID"`
	MetaSoft
	AuditTimes
}

func (OfferText) TableName() string { return "OfferText" }

// ---------- People / Tenancy ----------

type Address struct {
	AddressID string  `gorm:"primaryKey;type:char(36);column:AddressID"`
	OwnerID   *string `gorm:"type:char(36);column:OwnerID"`
	Address1  *string `gorm:"type:varchar(50);column:Address1"`
	Address2  *string `gorm:"type:varchar(50);column:Address2"`
	Address3  *string `gorm:"type:varchar(50);column:Address3"`
	City      *string `gorm:"type:varchar(30);column:City"`
	County    *string `gorm:"type:varchar(30);column:County"`
	StateCode *string `gorm:"type:char(2);column:StateCode"`
	ZipCode   *string `gorm:"type:varchar(10);column:ZipCode"`
	MetaSoft
	AuditTimes
}

func (Address) TableName() string { return "Address" }

type Tenant struct {
	TenantID   string  `gorm:"primaryKey;type:char(36);column:TenantID"`
	PropertyID *string `gorm:"type:char(36);column:PropertyID"`
	AddressID  *string `gorm:"type:char(36);column:AddressID"`
	UserID     *string `gorm:"type:char(36);column:UserID"`
	MetaSoft
	AuditTimes
}

func (Tenant) TableName() string { return "Tenant" }

type Account struct {
	AccountID string  `gorm:"primaryKey;type:char(36);column:AccountID"`
	AddressID *string `gorm:"type:char(36);column:AddressID"`
	TenantID  *string `gorm:"type:char(36);column:TenantID"`
	Unit      *string `gorm:"type:varchar(20);column:Unit"`
	Status    *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (Account) TableName() string { return "Account" }

type Phone struct {
	PhoneID     string  `gorm:"primaryKey;type:char(36);column:PhoneID"`
	PhoneType   *string `gorm:"type:varchar(10);column:PhoneType"`
	CountryCode *string `gorm:"type:varchar(3);column:CountryCode"`
	Number      *string `gorm:"type:varchar(20);column:Number"`
	Extension   *string `gorm:"type:varchar(8);column:Extension"`
	OwnerID     *string `gorm:"type:char(36);column:OwnerID"`
	MetaSoft
	AuditTimes
}

func (Phone) TableName() string { return "Phone" }

// ---------- Payments / Events ----------

type Method struct {
	MethodID   string  `gorm:"primaryKey;type:char(36);column:MethodID"`
	UserID     *string `gorm:"type:char(36);column:UserID"`
	MethodType *string `gorm:"type:varchar(10);column:MethodType"`
	LastFour   *string `gorm:"type:char(4);column:LastFour"`
	NameOnCard *string `gorm:"type:varchar(50);column:NameOnCard"`
	Expiration *string `gorm:"type:char(4);column:Expiration"`
	CardNumber *string `gorm:"type:varchar(100);column:CardNumber"`
	Status     *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (Method) TableName() string { return "Method" }

type TransCode struct {
	TransCodeID string  `gorm:"primaryKey;type:char(3);column:TransCodeID"`
	Title       string  `gorm:"type:varchar(30);column:Title"`
	TransCode   *string `gorm:"type:varchar(10);column:TransCode"`
	MetaCommon
	AuditTimes
}

func (TransCode) TableName() string { return "TransCode" }

type Decision struct {
	DecisionID    string  `gorm:"primaryKey;type:char(36);column:DecisionID"`
	OfferID       *string `gorm:"type:char(36);column:OfferID"`
	PaymentPlanID *string `gorm:"type:char(36);column:PaymentPlanID"`
	Decision      *string `gorm:"type:varchar(20);column:Decision"`
	Results       *string `gorm:"type:mediumtext;column:Results"`
	MetaSoft
	AuditTimes
}

func (Decision) TableName() string { return "Decision" }

type Acceptance struct {
	AcceptanceID   string     `gorm:"primaryKey;type:char(36);column:AcceptanceID"`
	DecisionID     *string    `gorm:"type:char(36);column:DecisionID"`
	AcceptanceDate *time.Time `gorm:"type:date;column:AcceptanceDate"`
	OfferID        *string    `gorm:"type:char(36);column:OfferID"`
	AccountID      *string    `gorm:"type:char(36);column:AccountID"`
	PaymentPlanID  *string    `gorm:"type:char(36);column:PaymentPlanID"`
	Contract       *string    `gorm:"type:mediumtext;column:Contract"`
	MetaSoft
	AuditTimes
}

func (Acceptance) TableName() string { return "Acceptance" }

type Payment struct {
	PaymentID     string     `gorm:"primaryKey;type:char(36);column:PaymentID"`
	AccountID     *string    `gorm:"type:char(36);column:AccountID"`
	PaymentPlanID *string    `gorm:"type:char(36);column:PaymentPlanID"`
	MethodID      *string    `gorm:"type:char(36);column:MethodID"`
	TransDate     *time.Time `gorm:"type:date;column:TransDate"`
	Memo          *string    `gorm:"type:varchar(100);column:Memo"`
	Amount        *Money     `gorm:"type:decimal(10,4);column:Amount"`
	Status        *string    `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (Payment) TableName() string { return "Payment" }

type Transaction struct {
	TransactionID string     `gorm:"primaryKey;type:char(36);column:TransactionID"`
	AccountID     *string    `gorm:"type:char(36);column:AccountID"`
	PaymentID     *string    `gorm:"type:char(36);column:PaymentID"`
	TransDate     *time.Time `gorm:"type:date;column:TransDate"`
	TransCode     *string    `gorm:"type:char(3);column:TransCode"`
	Amount        *Money     `gorm:"type:decimal(10,4);column:Amount"`
	MetaCommon
	AuditTimes
}

func (Transaction) TableName() string { return "Transaction" }

type EventType struct {
	EventTypeID string  `gorm:"primaryKey;type:char(36);column:EventTypeID"`
	Title       *string `gorm:"type:varchar(40);column:Title"`
	MetaSoft
	AuditTimes
}

func (EventType) TableName() string { return "EventType" }

type EventLog struct {
	EventID     uint       `gorm:"primaryKey;autoIncrement;column:EventID"`
	EventTypeID *string    `gorm:"type:char(36);column:EventTypeID"`
	EventDate   *time.Time `gorm:"type:datetime(6);column:EventDate"`
	EventData   JSON       `gorm:"type:json;column:EventData"`
	Status      *string    `gorm:"type:char(1);column:Status"`
	OwnerID     *string    `gorm:"type:char(36);column:OwnerID"`
	MetaCommon
	AuditTimes
}

func (EventLog) TableName() string { return "EventLog" }

// ---------- Portfolio / Campaign / Workflow ----------

type Portfolio struct {
	PortfolioID string  `gorm:"primaryKey;type:char(36);column:PortfolioID"`
	OrgID       *string `gorm:"type:char(36);column:OrgID"`
	Title       string  `gorm:"type:varchar(80);column:Title"`
	Description *string `gorm:"type:varchar(255);column:Description"`
	Status      *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (Portfolio) TableName() string { return "Portfolio" }

type PortfolioAccount struct {
	PortfolioAccountID uint       `gorm:"primaryKey;autoIncrement;column:PortfolioAccountID"`
	PortfolioID        string     `gorm:"type:char(36);column:PortfolioID"`
	AccountID          string     `gorm:"type:char(36);column:AccountID"`
	AssignedAt         *time.Time `gorm:"type:datetime(6);column:AssignedAt"`
	RemovedAt          *time.Time `gorm:"type:datetime(6);column:RemovedAt"`
	Status             *string    `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (PortfolioAccount) TableName() string { return "PortfolioAccount" }

type ActionType struct {
	ActionTypeID string `gorm:"primaryKey;type:varchar(20);column:ActionTypeID"`
	Title        string `gorm:"type:varchar(40);column:Title"`
	MetaSoft
	AuditTimes
}

func (ActionType) TableName() string { return "ActionType" }

type WorkflowTemplate struct {
	WorkflowTemplateID string  `gorm:"primaryKey;type:char(36);column:WorkflowTemplateID"`
	OrgID              *string `gorm:"type:char(36);column:OrgID"`
	Title              string  `gorm:"type:varchar(80);column:Title"`
	Description        *string `gorm:"type:varchar(255);column:Description"`
	SegmentRules       JSON    `gorm:"type:json;column:SegmentRules"`
	Status             *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (WorkflowTemplate) TableName() string { return "WorkflowTemplate" }

type WorkflowStep struct {
	WorkflowStepID     uint    `gorm:"primaryKey;autoIncrement;column:WorkflowStepID"`
	WorkflowTemplateID string  `gorm:"type:char(36);column:WorkflowTemplateID"`
	SeqNo              int     `gorm:"column:SeqNo"`
	ActionTypeID       string  `gorm:"type:varchar(20);column:ActionTypeID"`
	OffsetDays         int     `gorm:"column:OffsetDays"`
	Payload            JSON    `gorm:"type:json;column:Payload"`
	Status             *string `gorm:"type:char(1);column:Status"`
	MetaSoft
	AuditTimes
}

func (WorkflowStep) TableName() string { return "WorkflowStep" }

type Campaign struct {
	CampaignID         string     `gorm:"primaryKey;type:char(36);column:CampaignID"`
	OrgID              *string    `gorm:"type:char(36);column:OrgID"`
	PortfolioID        string     `gorm:"type:char(36);column:PortfolioID"`
	WorkflowTemplateID string     `gorm:"type:char(36);column:WorkflowTemplateID"`
	Title              string     `gorm:"type:varchar(80);column:Title"`
	Description        *string    `gorm:"type:varchar(255);column:Description"`
	CampaignType       *string    `gorm:"type:varchar(20);column:CampaignType"`
	StartDate          *time.Time `gorm:"type:date;column:StartDate"`
	EndDate            *time.Time `gorm:"type:date;column:EndDate"`
	Status             *string    `gorm:"type:char(1);column:Status"`
	Limits             JSON       `gorm:"type:json;column:Limits"`
	MetaSoft
	AuditTimes
}

func (Campaign) TableName() string { return "Campaign" }

type CampaignAction struct {
	CampaignActionID string     `gorm:"primaryKey;type:char(36);column:CampaignActionID"`
	CampaignID       string     `gorm:"type:char(36);column:CampaignID"`
	WorkflowStepID   uint       `gorm:"column:WorkflowStepID"`
	AccountID        *string    `gorm:"type:char(36);column:AccountID"`
	OwnerID          *string    `gorm:"type:char(36);column:OwnerID"`
	ActionTypeID     string     `gorm:"type:varchar(20);column:ActionTypeID"`
	ScheduledAt      *time.Time `gorm:"type:datetime(6);column:ScheduledAt"`
	ExecutedAt       *time.Time `gorm:"type:datetime(6);column:ExecutedAt"`
	Status           *string    `gorm:"type:char(1);column:Status"`
	Result           JSON       `gorm:"type:json;column:Result"`
	EventID          *uint      `gorm:"column:EventID"`
	MetaSoft
	AuditTimes
}

func (CampaignAction) TableName() string { return "CampaignAction" }

// ---------- Views (read-only helpers) ----------

// VPortfolioAccounts maps to the view v_portfolio_accounts
// Useful for read queries and reporting. Not intended for inserts/updates.
type VPortfolioAccounts struct {
	PortfolioID         string     `gorm:"column:PortfolioID"`
	PortfolioTitle      string     `gorm:"column:PortfolioTitle"`
	OrgID               *string    `gorm:"column:OrgID"`
	AccountID           string     `gorm:"column:AccountID"`
	TenantID            *string    `gorm:"column:TenantID"`
	AccountStatus       *string    `gorm:"column:AccountStatus"`
	AssignedAt          *time.Time `gorm:"column:AssignedAt"`
	RemovedAt           *time.Time `gorm:"column:RemovedAt"`
	IsActiveInPortfolio bool       `gorm:"column:IsActiveInPortfolio"`
}

func (VPortfolioAccounts) TableName() string { return "v_portfolio_accounts" }

// VCampaignSchedule maps to the view v_campaign_schedule
type VCampaignSchedule struct {
	CampaignID       string     `gorm:"column:CampaignID"`
	CampaignTitle    string     `gorm:"column:CampaignTitle"`
	PortfolioID      string     `gorm:"column:PortfolioID"`
	StartDate        *time.Time `gorm:"column:StartDate"`
	CampaignStatus   *string    `gorm:"column:CampaignStatus"`
	CampaignActionID string     `gorm:"column:CampaignActionID"`
	WorkflowStepID   uint       `gorm:"column:WorkflowStepID"`
	AccountID        *string    `gorm:"column:AccountID"`
	OwnerID          *string    `gorm:"column:OwnerID"`
	ActionTypeID     string     `gorm:"column:ActionTypeID"`
	ScheduledAt      *time.Time `gorm:"column:ScheduledAt"`
	ExecutedAt       *time.Time `gorm:"column:ExecutedAt"`
	ActionStatus     *string    `gorm:"column:ActionStatus"`
}

func (VCampaignSchedule) TableName() string { return "v_campaign_schedule" }
