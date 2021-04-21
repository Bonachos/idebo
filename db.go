package idea

import (
	"fmt"
	"log"
	"os"
	"strings"

	catalogservice "jpmenezes.com/idebo/gen/catalog_service"
	downloadservice "jpmenezes.com/idebo/gen/download_service"
	entity "jpmenezes.com/idebo/gen/entity"
	geoserver "jpmenezes.com/idebo/gen/geoserver"
	transformationservice "jpmenezes.com/idebo/gen/transformation_service"
	viewservice "jpmenezes.com/idebo/gen/view_service"
	viewer "jpmenezes.com/idebo/gen/viewer"

	"github.com/jinzhu/gorm"

	// Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	postgresHost     = os.Getenv("POSTGRES_HOST")
	postgresPort     = os.Getenv("POSTGRES_PORT")
	postgresUser     = os.Getenv("POSTGRES_USER")
	postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase = os.Getenv("POSTGRES_DATABASE")
)

func getDB() (*gorm.DB, error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		// log.Println("defaultTableName: " + defaultTableName)
		posResult := strings.Index(defaultTableName, "_result")
		if posResult > 0 {
			return defaultTableName[:posResult]
		}
		return defaultTableName
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", postgresHost, postgresPort, postgresUser, postgresPassword, postgresDatabase))
	if err != nil {
		log.Println("failed to connect database")
		return nil, err
	}
	db.SingularTable(true)

	// Migrate the schema
	db.AutoMigrate(&viewer.Viewer{})
	db.AutoMigrate(&catalogservice.CatalogService{})
	db.AutoMigrate(&downloadservice.DownloadService{})
	db.AutoMigrate(&entity.Entity{})
	db.AutoMigrate(&geoserver.Geoserver{})
	db.AutoMigrate(&viewservice.ViewService{})
	db.AutoMigrate(&transformationservice.TransformationService{})
	return db, nil
}

func createDB() error {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", postgresHost, postgresPort, postgresUser, postgresPassword))
	if err != nil {
		return err
	}
	defer db.Close()
	db = db.Exec("CREATE DATABASE " + postgresDatabase + ";")
	if db.Error != nil {
		fmt.Println("Unable to create DB")
		return db.Error
	}
	return nil
}

func dropDB() error {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", postgresHost, postgresPort, postgresUser, postgresPassword))
	if err != nil {
		return err
	}
	defer db.Close()
	db = db.Exec("DROP DATABASE " + postgresDatabase + ";")
	if db.Error != nil {
		fmt.Println("Unable to drop DB")
		return db.Error
	}
	return nil
}
