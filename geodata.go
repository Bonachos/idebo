package idea

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	geodata "jpmenezes.com/idebo/gen/geodata"
)

var (
	geodataFolder = os.Getenv("GEODATAFOLDER")
)

// geodata service example implementation.
// The example methods log the requests and return zero values.
type geodatasrvc struct {
	logger *log.Logger
}

// NewGeodata returns the geodata service implementation.
func NewGeodata(logger *log.Logger) geodata.Service {
	return &geodatasrvc{logger}
}

// List all stored geodata
func (s *geodatasrvc) List(ctx context.Context, p *geodata.ListPayload) (res geodata.GeodataResultCollection, view string, err error) {
	view = "default"
	s.logger.Print("geodata.list")

	var files []string

	geodataFolderLength := len(geodataFolder)
	err = filepath.Walk(geodataFolder, func(path string, info os.FileInfo, err error) error {
		fileName := info.Name()
		if strings.Index(fileName, ".geojson") > 0 {
			geojsonFilePath := path[geodataFolderLength:]
			slashPosition := strings.Index(geojsonFilePath, "/")
			entity := "default"
			if slashPosition > 0 {
				entity = geojsonFilePath[:slashPosition]
				geojsonFilePath = geojsonFilePath[slashPosition+1:]
			}
			files = append(files, geojsonFilePath)

			geodataResult := &geodata.GeodataResult{Name: geojsonFilePath, ID: uint(len(res)), Entity: entity}
			geodataResult.Entityname = &geodataResult.Entity
			entityName := getEntityName(geodataResult.Entity)
			if entityName != "" {
				geodataResult.Entityname = &entityName
			}

			res = append(res, geodataResult)
		}
		return nil
	})
	if err != nil {
		log.Println(err.Error())
	}

	// DEBUG
	for _, file := range files {
		log.Println(file)
	}
	return
}

// Upload implements upload.
func (s *geodatasrvc) Upload(ctx context.Context, p *geodata.FilesUpload) (res string, err error) {
	s.logger.Print("geodata.upload")
	s.logger.Print(*p.Files[0].Name, *p.Files[0].Type)
	return
}

// Remove geodata
func (s *geodatasrvc) Remove(ctx context.Context, p *geodata.RemovePayload) (err error) {
	methodName := "geodata.remove"
	s.logger.Print(methodName)

	path := filepath.Join(geodataFolder, strings.Replace(p.ID, "_-_", "/", -1))
	os.Remove(path)

	return
}
