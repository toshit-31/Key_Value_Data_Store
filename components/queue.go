package components

import (
	err "in-memory-db/1/error_code"
	"time"
)

type Queue struct {
	queue     [1000]string
	Count     int
	timestamp int64
}

func (q *Queue) Push(qvalues []string) bool {
	if (q.Count + len(qvalues)) > 1000 {
		return false
	}
	copy(q.queue[q.Count:], qvalues)
	q.Count += len(qvalues)
	q.timestamp = time.Now().UnixMilli() / 1000
	return true
}

func (q *Queue) Pop() (string, error) {
	if q.Count > 0 {
		q.Count--
		return q.queue[q.Count], nil
	} else {
		return "", err.ErrQueueEmpty
	}
}

func (q *Queue) Peek() (string, error) {
	if q.Count > 0 {
		return q.queue[q.Count-1], nil
	} else {
		return "", err.ErrQueueEmpty
	}
}
