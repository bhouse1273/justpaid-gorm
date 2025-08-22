package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type Config struct {
	Port          string
	MySQLDSN      string
	JWTSecret     string
	CORSOrigins   []string
	WorkerEnabled bool
	WorkerBatch   int
	WorkerTickSec int
}

var (
	JPLogger *zap.Logger
	JPConfig *Config
)

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func New() *Config {
	cors := getenv("CORS_ORIGINS", "")
	origins := []string{}
	if cors != "" {
		for _, o := range strings.Split(cors, ",") {
			origins = append(origins, strings.TrimSpace(o))
		}
	}
	batch, _ := strconv.Atoi(getenv("WORKER_BATCH", "200"))
	secs, _ := strconv.Atoi(getenv("WORKER_TICK_SECONDS", "30"))
	worker := getenv("WORKER_ENABLED", "0") == "1"
	cfg := &Config{
		Port:          getenv("PORT", "8080"),
		MySQLDSN:      getenv("MYSQL_DSN", ""),
		JWTSecret:     getenv("JWT_SECRET", ""),
		CORSOrigins:   origins,
		WorkerEnabled: worker,
		WorkerBatch:   batch,
		WorkerTickSec: secs,
	}
	if cfg.MySQLDSN == "" {
		log.Println("WARNING: MYSQL_DSN is empty; set it in .env")
	}
	// Assign global reference variable
	JPConfig = cfg

	return cfg
}
