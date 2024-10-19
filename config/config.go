package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HttpPort           string
	PostgresUser       string
	PostgresPassword   string
	PostgresDatabase   string
	PostgresHost       string
	PostgresPort       string
	DefaultOffset      string
	DefaultLimit       string
	CsvFilePath        string
	AuthConfigPath     string
	SigningKey         string
	AwsRegion          string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	ClinicBucket       string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error connect .env loading: ", err.Error())
	}
	c := Config{}
	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "port"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres_username"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "password"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "db_name"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "host"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "port"))
	c.DefaultLimit = cast.ToString(getOrReturnDefault("DEFAULT_LIMIT", "10000000"))
	c.DefaultOffset = cast.ToString(getOrReturnDefault("DEFAULT_OFFSET", "0"))
	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_CONFIG_PATH", "file_path"))
	c.CsvFilePath = cast.ToString(getOrReturnDefault("CSV_FILE_PATH", "file_path"))
	c.SigningKey = cast.ToString(getOrReturnDefault("SIGNING_KEY", "traffilightsigningkey"))
	c.AwsRegion = cast.ToString(getOrReturnDefault("AWS_REGION", "REGION"))
	c.AwsAccessKeyId = cast.ToString(getOrReturnDefault("AWS_ACCESS_KEY_ID", "key_id"))
	c.AwsSecretAccessKey = cast.ToString(getOrReturnDefault("AWS_SECRET_ACCESS_KEY", "secret_key"))
	c.ClinicBucket = cast.ToString(getOrReturnDefault("TRAFFIC_LIGHT_BUCKET", "bucket"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
