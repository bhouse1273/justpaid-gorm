package main

import (
	"log"
	"time"

	"github.com/bhouse1273/justpaid-gorm/internal/config"
	"github.com/bhouse1273/justpaid-gorm/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.New()
	if !cfg.WorkerEnabled {
		log.Println("Worker disabled (set WORKER_ENABLED=1 to enable)")
		return
	}
	dbconn := db.MustOpen()
	log.Printf("Worker started: batch=%d tick=%ds", cfg.WorkerBatch, cfg.WorkerTickSec)
	ticker := time.NewTicker(time.Duration(cfg.WorkerTickSec) * time.Second)
	for range ticker.C {
		// Claim a batch
		type row struct{ CampaignActionID string }
		var claimed []row
		dbconn.Raw("CALL sp_campaign_claim_due_actions(?)", cfg.WorkerBatch).Scan(&claimed)
		if len(claimed) == 0 {
			continue
		}
		log.Printf("claimed %d actions", len(claimed))
		for _, r := range claimed {
			// TODO: perform the actual action (email/letter/dialer) using Result JSON
			_ = dbconn.Exec("CALL sp_campaign_complete_action(?, ?, JSON_OBJECT('worker','ok'))", r.CampaignActionID, 1).Error
		}
	}
}
