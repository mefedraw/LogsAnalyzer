package Service

import (
	"NginxLogsAnalyzer/Rendering"
	"errors"
)

type RenderService struct {
}

func NewRenderService() *RenderService {
	return &RenderService{}
}

func (rs *RenderService) GetRender(renderType string) (Rendering.Render, error) {
	switch renderType {
	case "markdown":
		return Rendering.NewMarkdownRenderer(), nil
	case "adoc":
		return Rendering.NewAdocRender(), nil
	default:
		return nil, errors.New("unknown")
	}
}
