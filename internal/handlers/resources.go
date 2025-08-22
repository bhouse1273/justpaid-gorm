package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/bhouse1273/justpaid-gorm/internal/models"
)

// ---------- Organizations ----------
func (h *Handler) ListOrganizations(c echo.Context) error {
	var rows []models.Organization
	if err := h.DB.Order("Title").Find(&rows).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, rows)
}

func (h *Handler) GetOrganization(c echo.Context) error {
	id := c.Param("id")
	var org models.Organization
	if err := h.DB.First(&org, "OrgID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, org)
}

func (h *Handler) CreateOrganization(c echo.Context) error {
	var dto OrganizationCreateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := validate.Struct(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	in := models.Organization{
		OrgID:       uuid.NewString(),
		Title:       dto.Title,
		ParentOrgID: dto.ParentOrgID,
	}
	if err := h.DB.Create(&in).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, in)
}

func (h *Handler) UpdateOrganization(c echo.Context) error {
	id := c.Param("id")
	var dto OrganizationUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// fetch existing
	var existing models.Organization
	if err := h.DB.First(&existing, "OrgID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if dto.Title != nil {
		existing.Title = *dto.Title
	}
	if dto.ParentOrgID != nil {
		existing.ParentOrgID = dto.ParentOrgID
	}
	if err := h.DB.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, existing)
}

func (h *Handler) DeleteOrganization(c echo.Context) error {
	id := c.Param("id")
	var existing models.Organization
	if err := h.DB.First(&existing, "OrgID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if err := h.DB.Delete(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// ---------- Generic CRUD helpers for string-PK models ----------
// Property
func (h *Handler) ListProperties(c echo.Context) error {
	var rows []models.Property
	if err := h.DB.Order("Title").Find(&rows).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) GetProperty(c echo.Context) error {
	id := c.Param("id")
	var r models.Property
	if err := h.DB.First(&r, "PropertyID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, r)
}
func (h *Handler) CreateProperty(c echo.Context) error {
	var dto PropertyCreateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := validate.Struct(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	in := models.Property{
		PropertyID: uuid.NewString(),
		Title:      dto.Title,
		OrgID:      dto.OrgID,
	}
	if err := h.DB.Create(&in).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, in)
}
func (h *Handler) UpdateProperty(c echo.Context) error {
	id := c.Param("id")
	var dto PropertyUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var existing models.Property
	if err := h.DB.First(&existing, "PropertyID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if dto.Title != nil {
		existing.Title = *dto.Title
	}
	if dto.OrgID != nil {
		existing.OrgID = dto.OrgID
	}
	if err := h.DB.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, existing)
}
func (h *Handler) DeleteProperty(c echo.Context) error {
	id := c.Param("id")
	var existing models.Property
	if err := h.DB.First(&existing, "PropertyID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if err := h.DB.Delete(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Offer
func (h *Handler) ListOffers(c echo.Context) error {
	var rows []models.Offer
	if err := h.DB.Order("Title").Find(&rows).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) GetOffer(c echo.Context) error {
	id := c.Param("id")
	var r models.Offer
	if err := h.DB.First(&r, "OfferID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, r)
}
func (h *Handler) CreateOffer(c echo.Context) error {
	var dto OfferCreateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := validate.Struct(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	in := models.Offer{
		OfferID: uuid.NewString(),
		Title:   dto.Title,
		OrgID:   dto.OrgID,
		Rules:   dto.Rules,
	}
	if err := h.DB.Create(&in).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, in)
}
func (h *Handler) UpdateOffer(c echo.Context) error {
	id := c.Param("id")
	var dto OfferUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var existing models.Offer
	if err := h.DB.First(&existing, "OfferID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if dto.Title != nil {
		existing.Title = *dto.Title
	}
	if dto.Rules != nil {
		existing.Rules = dto.Rules
	}
	if dto.Status != nil {
		existing.Status = dto.Status
	}
	if err := h.DB.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, existing)
}
func (h *Handler) DeleteOffer(c echo.Context) error {
	id := c.Param("id")
	var existing models.Offer
	if err := h.DB.First(&existing, "OfferID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if err := h.DB.Delete(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Portfolio
func (h *Handler) ListPortfolios(c echo.Context) error {
	var rows []models.Portfolio
	if err := h.DB.Order("Title").Find(&rows).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) GetPortfolio(c echo.Context) error {
	id := c.Param("id")
	var r models.Portfolio
	if err := h.DB.First(&r, "PortfolioID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, r)
}
func (h *Handler) CreatePortfolio(c echo.Context) error {
	var dto PortfolioCreateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := validate.Struct(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	in := models.Portfolio{
		PortfolioID:    uuid.NewString(),
		Title:          dto.Title,
		OrgID:          dto.OrgID,
		PortTemplateID: dto.PortTemplateID,
		Description:    dto.Description,
	}
	if err := h.DB.Create(&in).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, in)
}
func (h *Handler) UpdatePortfolio(c echo.Context) error {
	id := c.Param("id")
	var dto PortfolioUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var existing models.Portfolio
	if err := h.DB.First(&existing, "PortfolioID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if dto.Title != nil {
		existing.Title = *dto.Title
	}
	if dto.PortTemplateID != nil {
		existing.PortTemplateID = dto.PortTemplateID
	}
	if dto.Description != nil {
		existing.Description = dto.Description
	}
	if dto.Status != nil {
		existing.Status = dto.Status
	}
	if err := h.DB.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, existing)
}
func (h *Handler) DeletePortfolio(c echo.Context) error {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Portfolio{}, "PortfolioID = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// PortTemplate
func (h *Handler) ListPortTemplates(c echo.Context) error {
	var rows []models.PortTemplate
	query := h.DB.Order("Title")

	// Filter by OrgID if provided
	orgID := c.QueryParam("orgId")
	if orgID != "" {
		query = query.Where("OrgID = ?", orgID)
	}

	if err := query.Find(&rows).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) GetPortTemplate(c echo.Context) error {
	id := c.Param("id")
	var r models.PortTemplate
	if err := h.DB.First(&r, "PortTemplateID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, r)
}
func (h *Handler) CreatePortTemplate(c echo.Context) error {
	var dto PortTemplateCreateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := validate.Struct(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	in := models.PortTemplate{
		PortTemplateID: uuid.NewString(),
		Title:          dto.Title,
		OrgID:          dto.OrgID,
		Description:    dto.Description,
		Selection:      dto.Selection,
	}
	if err := h.DB.Create(&in).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, in)
}
func (h *Handler) UpdatePortTemplate(c echo.Context) error {
	id := c.Param("id")
	var dto PortTemplateUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var existing models.PortTemplate
	if err := h.DB.First(&existing, "PortTemplateID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if dto.Title != nil {
		existing.Title = *dto.Title
	}
	if dto.Description != nil {
		existing.Description = dto.Description
	}
	if dto.Selection != nil {
		existing.Selection = dto.Selection
	}
	if dto.Status != nil {
		existing.Status = dto.Status
	}
	if err := h.DB.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, existing)
}
func (h *Handler) DeletePortTemplate(c echo.Context) error {
	id := c.Param("id")
	if err := h.DB.Delete(&models.PortTemplate{}, "PortTemplateID = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Campaign
func (h *Handler) ListCampaigns(c echo.Context) error {
	var rows []models.Campaign
	if err := h.DB.Order("Title").Find(&rows).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) GetCampaign(c echo.Context) error {
	id := c.Param("id")
	var r models.Campaign
	if err := h.DB.First(&r, "CampaignID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, r)
}
func (h *Handler) CreateCampaign(c echo.Context) error {
	var dto CampaignCreateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := validate.Struct(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	in := models.Campaign{
		CampaignID:         uuid.NewString(),
		PortfolioID:        dto.PortfolioID,
		WorkflowTemplateID: dto.WorkflowTemplateID,
		Title:              dto.Title,
		Description:        dto.Description,
		CampaignType:       dto.CampaignType,
		StartDate:          dto.StartDate,
		EndDate:            dto.EndDate,
	}
	if err := h.DB.Create(&in).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, in)
}
func (h *Handler) UpdateCampaign(c echo.Context) error {
	id := c.Param("id")
	var dto CampaignUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var existing models.Campaign
	if err := h.DB.First(&existing, "CampaignID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if dto.Title != nil {
		existing.Title = *dto.Title
	}
	if dto.Description != nil {
		existing.Description = dto.Description
	}
	if dto.StartDate != nil {
		existing.StartDate = dto.StartDate
	}
	if dto.EndDate != nil {
		existing.EndDate = dto.EndDate
	}
	if dto.Status != nil {
		existing.Status = dto.Status
	}
	if err := h.DB.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, existing)
}
func (h *Handler) DeleteCampaign(c echo.Context) error {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Campaign{}, "CampaignID = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Account
func (h *Handler) ListAccounts(c echo.Context) error {
	var rows []models.Account
	if err := h.DB.Order("AccountID").Find(&rows).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) GetAccount(c echo.Context) error {
	id := c.Param("id")
	var r models.Account
	if err := h.DB.First(&r, "AccountID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, r)
}
func (h *Handler) CreateAccount(c echo.Context) error {
	var dto AccountCreateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	in := models.Account{
		AccountID: uuid.NewString(),
		AddressID: dto.AddressID,
		TenantID:  dto.TenantID,
		Unit:      dto.Unit,
	}
	if err := h.DB.Create(&in).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, in)
}
func (h *Handler) UpdateAccount(c echo.Context) error {
	id := c.Param("id")
	var dto AccountUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var existing models.Account
	if err := h.DB.First(&existing, "AccountID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if dto.AddressID != nil {
		existing.AddressID = dto.AddressID
	}
	if dto.TenantID != nil {
		existing.TenantID = dto.TenantID
	}
	if dto.Unit != nil {
		existing.Unit = dto.Unit
	}
	if dto.Status != nil {
		existing.Status = dto.Status
	}
	if err := h.DB.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, existing)
}
func (h *Handler) DeleteAccount(c echo.Context) error {
	id := c.Param("id")
	var existing models.Account
	if err := h.DB.First(&existing, "AccountID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if err := h.DB.Delete(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Payment
func (h *Handler) ListPayments(c echo.Context) error {
	var rows []models.Payment
	if err := h.DB.Order("TransDate").Find(&rows).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) GetPayment(c echo.Context) error {
	id := c.Param("id")
	var r models.Payment
	if err := h.DB.First(&r, "PaymentID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, r)
}
func (h *Handler) CreatePayment(c echo.Context) error {
	var dto PaymentCreateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// optional amount parsing
	var amount *models.Money
	if dto.Amount != nil && *dto.Amount != "" {
		d, err := decimal.NewFromString(*dto.Amount)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid amount"})
		}
		m := models.Money(d)
		amount = &m
	}
	in := models.Payment{
		PaymentID:     uuid.NewString(),
		AccountID:     dto.AccountID,
		PaymentPlanID: dto.PaymentPlanID,
		MethodID:      dto.MethodID,
		TransDate:     dto.TransDate,
		Memo:          dto.Memo,
		Amount:        amount,
	}
	if err := h.DB.Create(&in).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, in)
}
func (h *Handler) UpdatePayment(c echo.Context) error {
	id := c.Param("id")
	var dto PaymentUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var existing models.Payment
	if err := h.DB.First(&existing, "PaymentID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if dto.Memo != nil {
		existing.Memo = dto.Memo
	}
	if dto.Status != nil {
		existing.Status = dto.Status
	}
	if err := h.DB.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, existing)
}
func (h *Handler) DeletePayment(c echo.Context) error {
	id := c.Param("id")
	var existing models.Payment
	if err := h.DB.First(&existing, "PaymentID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if err := h.DB.Delete(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// User
func (h *Handler) ListUsers(c echo.Context) error {
	var rows []models.User
	if err := h.DB.Order("Email").Find(&rows).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	var r models.User
	if err := h.DB.First(&r, "UserID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, r)
}
func (h *Handler) CreateUser(c echo.Context) error {
	var dto UserCreateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := validate.Struct(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	in := models.User{
		UserID: uuid.NewString(),
		Email:  dto.Email,
	}
	if err := h.DB.Create(&in).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, in)
}
func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var dto UserUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var existing models.User
	if err := h.DB.First(&existing, "UserID = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if dto.Email != nil {
		existing.Email = *dto.Email
	}
	if dto.Status != nil {
		existing.Status = dto.Status
	}
	if err := h.DB.Save(&existing).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, existing)
}
func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if err := h.DB.Delete(&models.User{}, "UserID = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
