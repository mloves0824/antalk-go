package dao

import (
	"antalk-go/internal/auth/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	c  *config.Config
	db *sql.DB
}

func New(c *config.Config) *Dao {
	d := &Dao{
		c:  c,
		db: NewDb(c),
	}
	return d
}

func NewDb(c *config.Config) *sql.DB {
	db, err := sql.Open(c.GetString("mysql.driverName"), c.GetString("mysql.dataSourceName"))
	if err != nil {
		return nil
	}
	return db
}
