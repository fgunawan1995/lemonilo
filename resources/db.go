package resources

import (
	"fmt"

	"github.com/fgunawan1995/lemonilo/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //using postgres
)

//ConnectDb connect to accounting db
func ConnectDb(cfg *config.Config) (*sqlx.DB, error) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Pass, cfg.DB.Name, cfg.DB.SSLMode)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}
