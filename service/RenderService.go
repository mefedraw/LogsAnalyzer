package service

import (
	"NginxLogsAnalyzer/rendering"
	"errors"
)

type RenderService struct {
}

func NewRenderService() *RenderService {
	return &RenderService{}
}

func (rs *RenderService) GetRender(renderType string) (rendering.Render, error) {
	switch renderType {
	case "markdown":
		return rendering.NewMarkdownRenderer(), nil
	case "adoc":
		return rendering.NewAdocRender(), nil
	default:
		return nil, errors.New("unknown")
	}
}
