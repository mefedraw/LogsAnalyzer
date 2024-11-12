package bufferedSource

import "bufio"

type BufferedSourceProvider interface {
	DataBufferWrap(path string) (*bufio.Reader, error)
}
