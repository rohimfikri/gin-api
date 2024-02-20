package config

type Config struct {
	GIN_MODE         string
	PORT             string
	TRUSTED_PROXIES  string
	LOG_LEVEL        string
	LOG_PRETTY       uint8
	JWT_SECRET_KEY   string
	JWT_EXPIRED_HOUR uint8
	JWT_ISSUER       string
	DB_SYS_DRIVER    string
	DB_SYS_USER      string
	DB_SYS_PASS      string
	DB_SYS_URL       string
	DB_SYS_NAME      string
}
