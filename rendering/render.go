package rendering

import (
	"NginxLogsAnalyzer/fileModel"
)

type Render interface {
	Render(file *fileModel.FileModel)
}
