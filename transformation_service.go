package idea

import (
	"context"
	"log"

	transformationservice "jpmenezes.com/idebo/gen/transformation_service"
)

// transformationService service example implementation.
// The example methods log the requests and return zero values.
type transformationServicesrvc struct {
	logger *log.Logger
}

// NewTransformationService returns the transformationService service
// implementation.
func NewTransformationService(logger *log.Logger) transformationservice.Service {
	return &transformationServicesrvc{logger}
}

// List all stored transformationServices
func (s *transformationServicesrvc) List(ctx context.Context) (res transformationservice.TransformationServiceResultCollection, err error) {
	methodName := "transformationService.list"
	s.logger.Print(methodName)

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	var transformationServices transformationservice.TransformationServiceResultCollection
	if err = db.Find(&transformationServices).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	// log.Println(viewers)

	return transformationServices, nil
}

// Show transformationService by ID
func (s *transformationServicesrvc) Show(ctx context.Context, p *transformationservice.ShowPayload) (res *transformationservice.TransformationServiceResult, view string, err error) {
	res = &transformationservice.TransformationServiceResult{}
	view = "default"
	s.logger.Print("transformationService.show")
	return
}

// Add new transformationService and return its ID.
func (s *transformationServicesrvc) Add(ctx context.Context, p *transformationservice.TransformationService) (res string, err error) {
	s.logger.Print("transformationService.add")
	return
}

// Remove transformationService
func (s *transformationServicesrvc) Remove(ctx context.Context, p *transformationservice.RemovePayload) (err error) {
	s.logger.Print("transformationService.remove")
	return
}
