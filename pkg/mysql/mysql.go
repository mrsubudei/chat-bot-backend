package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mrsubudei/chat-bot-backend/config"
)

type Mysql struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*Mysql, error) {
	dbHost := cfg.MySql.Host
	dbPort := cfg.MySql.Port
	dbUser := cfg.MySql.User
	dbPass := cfg.MySql.Password
	dbName := cfg.MySql.Name
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// val := url.Values{}
	// val.Add("parseTime", cfg.MySql.TimeZone)
	// val.Add("loc", cfg.MySql.Location)
	// dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open("mysql", connection)

	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	return &Mysql{
		DB: dbConn,
	}, nil
}

func (s *Mysql) Close() error {
	return s.DB.Close()
}
