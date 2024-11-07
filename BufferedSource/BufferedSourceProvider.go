package BufferedSource

import "bufio"

type BufferedSourceProvider interface {
	DataBufferWrap(path string) (*bufio.Reader, error)
}
