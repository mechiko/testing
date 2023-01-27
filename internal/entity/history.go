package entity

import "time"

type History interface {
	Length() int
	Push(s string)
	List() []string
	Clear()
	Last() string
}

type history struct {
	buf    []string
	head   int
	maxlen int
}

var _ History = &history{}

// New constructs and returns a new Queue.
func NewHistory(len int) *history {
	return &history{
		buf:    make([]string, len),
		head:   0,
		maxlen: len,
	}
}

// Length returns the number of elements currently stored in the queue.
func (h *history) Length() int {
	return len(h.buf)
}

func (h *history) Push(s string) {
	st := time.Now().Format("03:04:05") + " " + s
	if (h.head + 1) < len(h.buf) {
		h.buf[h.head] = st
		h.head += 1
	} else {
		h.buf = append(h.buf[1:], st)
	}
}

func (h *history) List() []string {
	return h.buf
}

func (h *history) Clear() {
	h.buf = make([]string, h.maxlen)
	h.head = 0
}

func (h *history) Last() string {
	if h.head > 0 {
		return h.buf[h.head-1]
	}
	return ""
}
