package postgresql

import (
	"Backend_Engineer_Interview_Assignment/common/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func PostgreSQLConfig(cfg *config.Config) *sql.DB {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.PostgreSQL.Host, cfg.PostgreSQL.Port, cfg.PostgreSQL.User, cfg.PostgreSQL.Pass, cfg.PostgreSQL.DbName)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	return db
}
