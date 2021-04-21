package idea

import (
	"context"
	"log"

	downloadservice "jpmenezes.com/idebo/gen/download_service"
)

// downloadService service example implementation.
// The example methods log the requests and return zero values.
type downloadServicesrvc struct {
	logger *log.Logger
}

// NewDownloadService returns the downloadService service implementation.
func NewDownloadService(logger *log.Logger) downloadservice.Service {
	return &downloadServicesrvc{logger}
}

// List all stored downloadServices
func (s *downloadServicesrvc) List(ctx context.Context) (res downloadservice.DownloadServiceResultCollection, err error) {
	s.logger.Print("downloadService.list")
	return
}

// Show downloadService by ID
func (s *downloadServicesrvc) Show(ctx context.Context, p *downloadservice.ShowPayload) (res *downloadservice.DownloadServiceResult, view string, err error) {
	res = &downloadservice.DownloadServiceResult{}
	view = "default"
	s.logger.Print("downloadService.show")
	return
}

// Add new downloadService and return its ID.
func (s *downloadServicesrvc) Add(ctx context.Context, p *downloadservice.DownloadService) (res string, err error) {
	s.logger.Print("downloadService.add")
	return
}

// Remove downloadService
func (s *downloadServicesrvc) Remove(ctx context.Context, p *downloadservice.RemovePayload) (err error) {
	s.logger.Print("downloadService.remove")
	return
}
