package idea

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
	smaconfig "jpmenezes.com/idebo/gen/sma_config"

	// Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// smaConfig service example implementation.
// The example methods log the requests and return zero values.
type smaConfigsrvc struct {
	logger *log.Logger
}

// NewSmaConfig returns the smaConfig service implementation.
func NewSmaConfig(logger *log.Logger) smaconfig.Service {
	return &smaConfigsrvc{logger}
}

func getGeonetworkDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=XPTO port=5432 user=postgres dbname=idea_master password="+postgresPassword+" sslmode=disable")
	if err != nil {
		log.Println("failed to connect database")
		return nil
	}
	db.SingularTable(true)

	return db
}

// Update the SMA Configuration
func (s *smaConfigsrvc) Update(ctx context.Context, p *smaconfig.SMAConfig) (err error) {
	s.logger.Print("smaConfig.update")

	db := getGeonetworkDB()
	if db == nil {
		return
	}
	defer db.Close()

	db.Exec("UPDATE geonetwork.settings SET value = '?' WHERE name = 'system/site/name';", p.Title)

	return
}
