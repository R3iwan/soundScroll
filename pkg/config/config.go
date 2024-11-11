package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresConfig PostgresConfig
	MongoConfig    MongoConfig
	RedisConfig    RedisConfig
	RabbitMQConfig RabbitMQConfig
	KafkaConfig    KafkaConfig
	ServerConfig   ServerConfig
	JWTConfig      JWTConfig
	APIConfig      APIConfig
	LoggingConfig  LoggingConfig
}

type PostgresConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     string
	SSLMode  string
}

type MongoConfig struct {
	URI      string
	Database string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type RabbitMQConfig struct {
	URI string
}

type KafkaConfig struct {
	Brokers    []string
	Topic      string
	GroupID    string
	AutoOffset string
}

type ServerConfig struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
}

type JWTConfig struct {
	SecretKey     string
	TokenExpiry   int
	RefreshExpiry int
}

type APIConfig struct {
	SpotifyAPIKey string
	LastFmAPIKey  string
}

type LoggingConfig struct {
	Level string
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("SERVER_READ_TIMEOUT", 15)
	viper.SetDefault("SERVER_WRITE_TIMEOUT", 15)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	return &Config{
		PostgresConfig: PostgresConfig{
			Host:     viper.GetString("POSTGRES_HOST"),
			Name:     viper.GetString("POSTGRES_NAME"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			Port:     viper.GetString("POSTGRES_PORT"),
			SSLMode:  viper.GetString("SSL_MODE"),
		},
		MongoConfig: MongoConfig{
			URI:      viper.GetString("MONGO_URI"),
			Database: viper.GetString("MONGO_DATABASE"),
		},
		RedisConfig: RedisConfig{
			Addr:     viper.GetString("REDIS_ADDR"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		},
		RabbitMQConfig: RabbitMQConfig{
			URI: viper.GetString("RABBIT_URI"),
		},
		KafkaConfig: KafkaConfig{
			Brokers:    viper.GetStringSlice("KAFKA_BROKERS"),
			Topic:      viper.GetString("KAFKA_TOPIC"),
			GroupID:    viper.GetString("KAFKA_GROUP_ID"),
			AutoOffset: viper.GetString("KAFKA_AUTO_OFFSET"),
		},
		ServerConfig: ServerConfig{
			Port:         viper.GetString("SERVER_PORT"),
			ReadTimeout:  viper.GetInt("SERVER_READ_TIMEOUT"),
			WriteTimeout: viper.GetInt("SERVER_WRITE_TIMEOUT"),
		},
		JWTConfig: JWTConfig{
			SecretKey:     viper.GetString("JWT_SECRET_KEY"),
			TokenExpiry:   viper.GetInt("JWT_TOKEN_EXPIRY"),
			RefreshExpiry: viper.GetInt("JWT_REFRESH_EXPIRY"),
		},
		APIConfig: APIConfig{
			SpotifyAPIKey: viper.GetString("API_SPOTIFY"),
			LastFmAPIKey:  viper.GetString("API_LAST_FM"),
		},
		LoggingConfig: LoggingConfig{
			Level: viper.GetString("LOGGING_LEVEL"),
		},
	}, nil
}
