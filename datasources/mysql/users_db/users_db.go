package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tupt0101/bookstore_users-api/utils/config"
)

var (
	Client *sql.DB
)

func init() {

	config, configErr := config.LoadConfig(".")
	if configErr != nil {
		log.Fatal("Cannot load config", configErr)
	}

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBSchema,
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
