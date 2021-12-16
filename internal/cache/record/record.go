package record

import "time"

const defaultTTL = time.Minute

type Record struct {
	CreationTime time.Time
	TimeToLive   time.Duration
	Key          string
	Value        string
}

//New создает новую запись.
//Передавать time.Duration(0) для стандартного ttl
func New(key, value string, timeToLive time.Duration) *Record {
	if timeToLive != time.Duration(0) {
		return &Record{
			CreationTime: time.Now(),
			TimeToLive:   timeToLive,
			Key:          key,
			Value:        value,
		}
	} else {
		return &Record{
			CreationTime: time.Now(),
			TimeToLive:   defaultTTL,
			Key:          key,
			Value:        value,
		}
	}
}
