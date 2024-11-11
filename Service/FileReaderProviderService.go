package Service

import (
	"NginxLogsAnalyzer/BufferedSource"
	"NginxLogsAnalyzer/Errors/FileReaderProviderServiceError"
	"net/url"
	"os"
)

type FileReaderProviderService struct {
}

func NewFileReaderProviderService() *FileReaderProviderService {
	return &FileReaderProviderService{}
}

func IsURL(path string) bool {
	u, err := url.Parse(path)
	if err != nil {
		return false
	}

	return u.Scheme == "http" || u.Scheme == "https"
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (frp *FileReaderProviderService) GetReader(path string) (BufferedSource.BufferedSourceProvider, error) {
	if IsURL(path) {
		return BufferedSource.NewHttpResponseReaderProvider(), nil
	} else if IsFile(path) {
		return BufferedSource.NewFileReaderProvider(), nil
	} else {
		return nil, FileReaderProviderServiceError.NewErrorFileReaderProvider("incorrect file path")
	}
}
