package cfg

import (
	"io/fs"
	"os"

	"github.com/joho/godotenv"
)

type Env int

const (
	Dev Env = iota
	Prod
)

var (
	ENV            Env = Dev
	ENCRYPTION_KEY []byte
)

const RW_ONLY_BY_APP fs.FileMode = 0600

func init() {
	godotenv.Load(".env")

	envStr := os.Getenv("ENV")
	var envEnum Env

	if envStr == "prod" {
		envEnum = Prod
	} else {
		envEnum = Dev
	}
	// Get encryption key from environment variables
	keyStr := os.Getenv("ENCRYPTION_KEY")
	// HARD CODING A DEVELOPMENT KEY FOR TESTING
	if envEnum == Dev {
		keyStr = "7458156cc861ef8305be92cf5b11d6b95dfcf40be62f6f781db550c91b85ecc2"
	}
	if len(keyStr) != 64 { // 64 hex characters = 32 bytes = 256 bits
		panic("Error: ENCRYPTION_KEY must be 64 hex characters long")
	}
	ENV = envEnum
	ENCRYPTION_KEY = []byte(keyStr)
}
