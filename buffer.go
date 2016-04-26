package executil

import (
	"bytes"
	"sync"
)

type threadsafeBuffer struct {
	sync.Mutex
	bytes.Buffer
}

func (buffer *threadsafeBuffer) Write(data []byte) (int, error) {
	buffer.Lock()
	defer buffer.Unlock()

	return buffer.Buffer.Write(data)
}
