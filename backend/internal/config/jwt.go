package config

import (
	"log"
	"os"
	"time"
)

type JWTConfig struct {
	Secret     string
	AccessTTL  time.Duration
}

func LoadJWT() JWTConfig {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" { log.Fatal("JWT_SECRET not set") }
	ttlStr := os.Getenv("JWT_ACCESS_TTL")
	if ttlStr == "" { ttlStr = "24h" }
	d, err := time.ParseDuration(ttlStr)
	if err != nil { d = 24 * time.Hour }
	return JWTConfig{ Secret: secret, AccessTTL: d }
}
