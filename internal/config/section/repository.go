package section

import (
	"fmt"
	"time"
)

type Repository struct {
	Postgres RepositoryPostgres
	Redis    RepositoryRedis
}

type RepositoryPostgres struct {
	Host         string
	Port         string
	Username     string
	Password     string
	Name         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (p RepositoryPostgres) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, p.Username, p.Password, p.Name,
	)
}

type RepositoryRedis struct {
	Address  string
	Password string
	DB       int
}
