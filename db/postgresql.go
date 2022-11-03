package db

import (
	"akaflieg/fresskasse/v2/config"
	"akaflieg/fresskasse/v2/entities"
	"context"
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB

func InitDB() {
	config := config.GetConfig()
	conn := pgdriver.NewConnector(
		pgdriver.WithAddr(config.DB_ADDRESS),
		pgdriver.WithUser(config.DB_USER),
		pgdriver.WithPassword(config.DB_PASSWORD),
		pgdriver.WithDatabase(config.DB),
		pgdriver.WithInsecure(true),
	)

	sqldb := sql.OpenDB(conn)
	db = bun.NewDB(sqldb, pgdialect.New())

}

func SetupDB() {
	_, err := db.NewCreateTable().
		Model((*entities.User)(nil)).
		Exec(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.NewCreateTable().
		Model((*entities.Meal)(nil)).
		Exec(context.Background())

	if err != nil {
		log.Fatal(err)
	}
}
