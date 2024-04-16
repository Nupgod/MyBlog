package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

const postTable = "posts" 

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	conf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSLMode)
	db, err := sqlx.Open("postgres", conf)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	query := `
	DROP TABLE IF EXISTS posts;	
	CREATE TABLE IF NOT EXISTS posts
	(
		id serial not null unique,
		title varchar(255) not null,
		p_text text not null,
		p_time timestamp not null
	);
	`

	_, err = db.Query(query);
	if err != nil {
		return nil, err
	}

	return db, nil
}
