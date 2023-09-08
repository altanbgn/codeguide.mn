package config

import (
	"os"
	"strconv"
	"time"
)

type DB struct {
  Host string
  Port int
  User string
  Password string
  Name string
  Debug bool
  SslMode string
  MaxOpenConn int
  MaxIdleConn int
  MaxConnLifetime time.Duration
}

var db = &DB{}

func DBConfig() *DB {
  return db
}

func LoadDBConfig() {
  db.Host = os.Getenv("DB_HOST")
  db.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
  db.User = os.Getenv("DB_USER")
  db.Password = os.Getenv("DB_PASSWORD")
  db.Name = os.Getenv("DB_NAME")
  db.SslMode = os.Getenv("DB_SSL_MODE")
  db.Debug, _ = strconv.ParseBool(os.Getenv("DB_DEBUG"))
  db.MaxOpenConn, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
  db.MaxIdleConn, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
  lifetime, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
  db.MaxConnLifetime = time.Duration(lifetime) * time.Second
}
