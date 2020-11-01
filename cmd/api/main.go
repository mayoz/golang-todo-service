package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"service/internal"
	"service/internal/todo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	conn, err := sql.Open("mysql", dbDSN())
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	if err := migrate(conn); err != nil {
		log.Fatal(err)
	}

	store := todo.NewStore(conn)
	server := setupServer(store)

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := server.Start(address); err != nil {
		log.Fatal(err)
	}
}

func dbDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&multiStatements=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

func migrate(db *sql.DB) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := db.Exec(internal.Migration()); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
