package dataCollecting

import (
	"bufio"
)

type DataCollector interface {
	CollectData(reader *bufio.Reader) error
}
