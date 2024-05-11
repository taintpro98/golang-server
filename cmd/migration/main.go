package main

import (
	"fmt"
	"golang-server/config"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func migrateDB(db *gorm.DB, migrationDir string) error {
	//migrator := db.Migrator()
	//
	//files, err := os.ReadDir(migrationDir)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("xxx", files)
	//return nil

	// Sort the migration files based on their version number
	return filepath.Walk(migrationDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Mode().IsRegular() && info.Name() != ".gitkeep" {
			contents, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Execute the SQL statements in the migration file
			rawSQL := string(contents)
			if err := db.Exec(rawSQL).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func test() {
	cnf := config.Init()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cnf.Database.Host,
		cnf.Database.Username,
		cnf.Database.Password,
		cnf.Database.DatabaseName,
		cnf.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	err = migrateDB(db, "./migrations")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database migration completed successfully!")
}
