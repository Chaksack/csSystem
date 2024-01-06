package database

import (
	"log"
	"os"
	"xactscore/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("xactscore.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	//TODO: Add migrations
	db.AutoMigrate(
		&models.User{},
		&models.Staff{},
		&models.Ghcard{},
		&models.UserTin{},
		&models.Business{},
		&models.Role{},
		&models.Permissions{},
		&models.Bank{},
		&models.UserBankDetails{},
		&models.Momo{},
	)
	// db.Migrator().DropColumn(&models.BankUser{}, "bank") // temporary fix
	// db.Model(&models.User{}).Association("Banks")

	Database = DbInstance{Db: db}
}
