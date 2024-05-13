package main

import (
	"context"
	"flag"
	"golang-server/config"
	"golang-server/pkg/database"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	envi := flag.String("e", "", "Environment option")
	dir := flag.String("dir", "migrations", "Path to migrations")
	flag.Parse()
	args := flag.Args()
	command := args[0]

	ctx := context.Background()
	cnf := config.Init(*envi)
	dsn := database.GetDatabaseDSN(cnf.Database)
	db, err := goose.OpenDBWithDriver("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.RunContext(ctx, command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
