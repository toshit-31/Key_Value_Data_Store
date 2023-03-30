package core

import (
	"in-memory-db/1/components"
	"in-memory-db/1/error_code"
)

type DataStoreKey interface {
	IsExpired()
	GetValue()
	Push()
	Pop()
	Peek()
}

type DataStore struct {
	keyStore   map[string]*components.Value
	queueStore map[string]*components.Queue
}

func (s *DataStore) Init() {
	s.keyStore = make(map[string]*components.Value)
	s.queueStore = make(map[string]*components.Queue)
}

func (s *DataStore) Query(query string) components.Result {
	result := components.Result{Err: nil, Value: ""}
	err, parsed := CommandParser(query)
	if err != nil {
		result.Err = err
		return result
	}
	switch parsed.Action {
	case "GET":
		{
			value, present := s.keyStore[parsed.key]
			if !present || value.IsExpired() {
				go delete(s.keyStore, parsed.key)
				result.Err = error_code.ErrKeyNotFound
			} else {
				result.Value = value.GetValue()
			}
			return result
		}
	case "SET":
		{
			key := parsed.key
			_, present := s.keyStore[key]
			value := new(components.Value).Build(parsed.value, int64(parsed.expiry), parsed.timestamp)
			cond := parsed.condition
			if cond == "" {
				s.keyStore[key] = value
			} else {
				if present && cond == "XX" {
					s.keyStore[key] = value
				}
				if !present && cond == "NX" {
					s.keyStore[key] = value
				}
			}
			return result
		}
	case "QPUSH":
		{
			q, present := s.queueStore[parsed.key]
			if present {
				q.Push(parsed.qvalue)
			} else {
				s.queueStore[parsed.key] = new(components.Queue)
				s.queueStore[parsed.key].Push(parsed.qvalue)
			}
			return result
		}
	case "QPOP":
		{
			q, present := s.queueStore[parsed.key]
			if !present {
				result.Err = error_code.ErrKeyNotFound
			} else {
				result.Value, result.Err = q.Pop()
			}
			return result
		}
	default:
		{
			result.Err = error_code.ErrInvCommand
			return result
		}
	}
}
