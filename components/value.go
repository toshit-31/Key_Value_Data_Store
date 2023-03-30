package components

import "time"

type Value struct {
	value     string
	timeout   int64
	timestamp int64
}

func (v *Value) Build(value string, expiry int64, timestamp int64) *Value {
	v.value = value
	v.timeout = expiry + timestamp
	v.timestamp = timestamp
	return v
}

func (v *Value) GetValue() string {
	return v.value
}

func (v *Value) IsExpired() bool {
	return (!(v.timeout < v.timestamp) &&
		time.Now().UnixMilli()/1000 > v.timeout)
}
