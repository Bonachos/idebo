package idea

import (
	"context"
	viewservice "jpmenezes.com/idebo/gen/view_service"
	"log"
)

// viewService service example implementation.
// The example methods log the requests and return zero values.
type viewServicesrvc struct {
	logger *log.Logger
}

// NewViewService returns the viewService service implementation.
func NewViewService(logger *log.Logger) viewservice.Service {
	return &viewServicesrvc{logger}
}

// List all stored viewServices
func (s *viewServicesrvc) List(ctx context.Context) (res viewservice.ViewServiceResultCollection, err error) {
	s.logger.Print("viewService.list")
	return
}

// Show viewService by ID
func (s *viewServicesrvc) Show(ctx context.Context, p *viewservice.ShowPayload) (res *viewservice.ViewServiceResult, view string, err error) {
	res = &viewservice.ViewServiceResult{}
	view = "default"
	s.logger.Print("viewService.show")
	return
}

// Add new viewService and return its ID.
func (s *viewServicesrvc) Add(ctx context.Context, p *viewservice.ViewService) (res string, err error) {
	s.logger.Print("viewService.add")
	return
}

// Remove viewService
func (s *viewServicesrvc) Remove(ctx context.Context, p *viewservice.RemovePayload) (err error) {
	s.logger.Print("viewService.remove")
	return
}
