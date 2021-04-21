package idea

import (
	"context"
	"log"

	catalogservice "jpmenezes.com/idebo/gen/catalog_service"
)

// catalogService service example implementation.
// The example methods log the requests and return zero values.
type catalogServicesrvc struct {
	logger *log.Logger
}

// NewCatalogService returns the catalogService service implementation.
func NewCatalogService(logger *log.Logger) catalogservice.Service {
	return &catalogServicesrvc{logger}
}

// List all stored catalogServices
func (s *catalogServicesrvc) List(ctx context.Context) (res catalogservice.CatalogServiceResultCollection, err error) {
	s.logger.Print("catalogService.list")
	return
}

// Show catalogService by ID
func (s *catalogServicesrvc) Show(ctx context.Context, p *catalogservice.ShowPayload) (res *catalogservice.CatalogServiceResult, view string, err error) {
	res = &catalogservice.CatalogServiceResult{}
	view = "default"
	s.logger.Print("catalogService.show")
	return
}

// Add new catalogService and return its ID.
func (s *catalogServicesrvc) Add(ctx context.Context, p *catalogservice.CatalogService) (res string, err error) {
	s.logger.Print("catalogService.add")
	return
}

// Remove catalogService
func (s *catalogServicesrvc) Remove(ctx context.Context, p *catalogservice.RemovePayload) (err error) {
	s.logger.Print("catalogService.remove")
	return
}
