package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type claimReq struct {
	Batch int `json:"batch"`
}

// POST /api/actions/claim
func (h *Handler) ClaimDueActions(c echo.Context) error {
	var req claimReq
	if err := c.Bind(&req); err != nil || req.Batch <= 0 { req.Batch = 100 }
	rows, err := h.DB.Raw("CALL sp_campaign_claim_due_actions(?)", req.Batch).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	// stream rows as JSON array (keeps it simple for now)
	types, _ := rows.ColumnTypes()
	cols, _ := rows.Columns()
	result := make([]map[string]any, 0, req.Batch)
	for rows.Next() {
		vals := make([]any, len(cols))
		for i, ct := range types {
			switch ct.DatabaseTypeName() {
			case "DATETIME", "TIMESTAMP", "DATE", "TIME":
				var s sql.NullTime; vals[i] = &s
			default:
				var s sql.NullString; vals[i] = &s
			}
		}
		if err := rows.Scan(vals...); err != nil { continue }
		m := map[string]any{}
		for i, name := range cols {
			switch v := vals[i].(type) {
			case *sql.NullTime:
				if v.Valid { m[name] = v.Time } else { m[name] = nil }
			case *sql.NullString:
				if v.Valid { m[name] = v.String } else { m[name] = nil }
			default:
				m[name] = v
			}
		}
		result = append(result, m)
	}
	return c.JSON(http.StatusOK, result)
}

type completeReq struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

// POST /api/actions/:id/complete
func (h *Handler) CompleteAction(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}
	var req completeReq
	if err := c.Bind(&req); err != nil { req.Success = true }
	if req.Result == nil { req.Result = map[string]any{"ok": req.Success} }
	if err := h.DB.Exec("CALL sp_campaign_complete_action(?, ?, JSON_OBJECT('payload', ?))", id, req.Success, req.Result).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"ok": true, "id": id})
}
