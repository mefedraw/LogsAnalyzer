package Redndering

import (
	"NginxLogsAnalyzer/FileModel"
)

type Render interface {
	Render(file *FileModel.FileModel)
}
