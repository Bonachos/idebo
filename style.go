package idea

import (
	"context"
	"log"

	style "jpmenezes.com/idebo/gen/style"
)

// style service example implementation.
// The example methods log the requests and return zero values.
type stylesrvc struct {
	logger *log.Logger
}

// NewStyle returns the style service implementation.
func NewStyle(logger *log.Logger) style.Service {
	return &stylesrvc{logger}
}

// List all stored styles
func (s *stylesrvc) List(ctx context.Context) (res style.StyleResultCollection, err error) {
	s.logger.Print("style.list")
	return
}

// Show style by ID
func (s *stylesrvc) Show(ctx context.Context, p *style.ShowPayload) (res *style.StyleResult, view string, err error) {
	res = &style.StyleResult{}
	view = "default"
	s.logger.Print("style.show")
	return
}

// Add new style and return its ID.
func (s *stylesrvc) Add(ctx context.Context, p *style.Style) (res string, err error) {
	s.logger.Print("style.add")
	return
}

// Remove style
func (s *stylesrvc) Remove(ctx context.Context, p *style.RemovePayload) (err error) {
	s.logger.Print("style.remove")
	return
}
