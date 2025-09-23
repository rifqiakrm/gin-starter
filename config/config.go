// Package config provides config
package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config holds configuration for the project.
type Config struct {
	Env     string `env:"APP_ENV,default=development"`
	AppName string `env:"APP_NAME,default=starter-api"`
	Host    string `env:"APP_HOST,default=localhost"`
	Port    string `env:"APP_PORT,default=8080"`

	HashID    HashID
	Postgres  Postgres
	MongoDB   MongoDB
	Redis     Redis
	RabbitMQ  RabbitMQ
	SMTP      SMTP
	JWTConfig JWTConfig
	Google    Google
	AWS       AWS
	Minio     Minio
	Image     Image
	URL       URL
	Jaeger    Jaeger
}

// HashID holds configuration for HashID.
type HashID struct {
	Salt      string `env:"HASHID_SALT"`
	MinLength int    `env:"HASHID_MIN_LENGTH,default=10"`
}

// Postgres holds all configuration for PostgreSQL.
type Postgres struct {
	Host            string `env:"POSTGRES_HOST,default=localhost"`
	Port            string `env:"POSTGRES_PORT,default=5432"`
	User            string `env:"POSTGRES_USER,required"`
	Password        string `env:"POSTGRES_PASSWORD,required"`
	Name            string `env:"POSTGRES_NAME,required"`
	MaxOpenConns    string `env:"POSTGRES_MAX_OPEN_CONNS,default=5"`
	MaxConnLifetime string `env:"POSTGRES_MAX_CONN_LIFETIME,default=10m"`
	MaxIdleLifetime string `env:"POSTGRES_MAX_IDLE_LIFETIME,default=5m"`
}

// MongoDB holds configuration for MongoDB.
type MongoDB struct {
	ConnString string `env:"MONGODB_CONN_STRING,required"`
	DBName     string `env:"MONGODB_DBNAME,required"`
}

// Redis holds configuration for the Redis.
type Redis struct {
	Address       string `env:"REDIS_ADDRESS"`
	Password      string `env:"REDIS_PASSWORD"`
	DatabaseIndex string `env:"REDIS_DATABASE_INDEX"`
}

// RabbitMQ holds configuration for RabbitMQ.
type RabbitMQ struct {
	Host     string `env:"RABBITMQ_HOST"`
	Username string `env:"RABBITMQ_USERNAME"`
	Password string `env:"RABBITMQ_PASSWORD"`
	Port     string `env:"RABBITMQ_PORT"`
	VHost    string `env:"RABBITMQ_VHOST"`
}

// SMTP holds configuration for smtp email.
type SMTP struct {
	Host string `env:"SMTP_HOST,required"`
	Port int    `env:"SMTP_PORT,default=587"`
	User string `env:"SMTP_USER,required"`
	Pass string `env:"SMTP_PASS,required"`
	From string `env:"SMTP_FROM"`
}

// JWTConfig holds configuration for jwt.
type JWTConfig struct {
	Public  string `env:"JWT_PUBLIC_KEY,required"`
	Private string `env:"JWT_PRIVATE_KEY,required"`
	Issuer  string `env:"JWT_ISSUER,required"`
}

// Google holds configuration for the Google.
type Google struct {
	ProjectID          string `env:"GOOGLE_PROJECT_ID"`
	ServiceAccountFile string `env:"GOOGLE_SA"`
	StorageBucketName  string `env:"GOOGLE_STORAGE_BUCKET_NAME"`
	StorageEndpoint    string `env:"GOOGLE_STORAGE_ENDPOINT"`
}

// AWS holds configuration for the AWS.
type AWS struct {
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	Region          string `env:"AWS_REGION"`
	BucketName      string `env:"AWS_BUCKET_NAME"`
}

// Minio holds configuration for Minio.
type Minio struct {
	Host         string `env:"MINIO_HOST"`
	AccessKey    string `env:"MINIO_ACCESS"`
	SecretKey    string `env:"MINIO_SECRET"`
	UseSSL       string `env:"MINIO_USE_SSL"`
	Bucket       string `env:"MINIO_BUCKET"`
	BucketPublic string `env:"MINIO_BUCKET_PUBLIC"`
}

// Image holds configuration for the Image.
type Image struct {
	Host string `env:"UPLOAD_SERVICE"`
}

// URL holds configuration for the URL.
type URL struct {
	BackendURL         string `env:"BACKEND_URL"`
	FrontendURL        string `env:"FRONTEND_URL"`
	ForgotPasswordURL  string `env:"FORGOT_PASSWORD_URL"`
	AjariAIURL         string `env:"AJARI_AI_URL"`
	ChatAppService     string `env:"CHAT_APP_SERVICE"`
	SeeuAppService     string `env:"SEEU_APP_SERVICE"`
	AIPreAssessmentURL string `env:"AI_PRE_ASSESSMENT_URL"`
}

// Jaeger holds configuration for the Jaeger.
type Jaeger struct {
	Address string `env:"JAEGER_ADDRESS"`
	Port    string `env:"JAEGER_PORT"`
}

// LoadConfig initiate load config either from env file or os env
func LoadConfig(env string) (*Config, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
