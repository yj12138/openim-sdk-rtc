package common

import (
	"errors"
	"sync"
)

type RingBuffer struct {
	buffer []byte
	size   int
	read   int
	write  int
	mutex  sync.Mutex
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]byte, size),
		size:   size,
		read:   0,
		write:  0,
	}
}

func (rb *RingBuffer) Write(data []byte) (int, error) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	n := len(data)
	if n == 0 {
		return 0, nil
	}

	if rb.availableWrite() < n {
		return 0, errors.New("buffer is full")
	}

	firstPart := min(n, rb.size-rb.write)
	copy(rb.buffer[rb.write:], data[:firstPart])
	secondPart := n - firstPart
	if secondPart > 0 {
		copy(rb.buffer[:secondPart], data[firstPart:])
	}

	rb.write = (rb.write + n) % rb.size
	return n, nil
}

func (rb *RingBuffer) Read(p []byte) (int, error) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	n := len(p)
	if n == 0 {
		return 0, nil
	}

	if rb.availableRead() == 0 {
		return 0, errors.New("buffer is empty")
	}

	n = min(n, rb.availableRead())

	// 分两次读取以处理环绕情况
	firstPart := min(n, rb.size-rb.read)
	copy(p, rb.buffer[rb.read:rb.read+firstPart])
	secondPart := n - firstPart
	if secondPart > 0 {
		copy(p[firstPart:], rb.buffer[:secondPart])
	}

	rb.read = (rb.read + n) % rb.size
	return n, nil
}

func (rb *RingBuffer) availableRead() int {
	if rb.write >= rb.read {
		return rb.write - rb.read
	}
	return rb.size - rb.read + rb.write
}

func (rb *RingBuffer) availableWrite() int {
	return rb.size - rb.availableRead()
}
