package cache

import "time"

const defaultTTL = time.Minute

type Record struct {
	creationTime time.Time
	timeToLive time.Duration
	key string
	value string
}

//New создает новую запись.
//Передавать time.Duration(0) для стандартного времени жизни
func New(key, value string, timeToLive time.Duration) *Record {
	if timeToLive != time.Duration(0) {
		return &Record{
			creationTime: time.Now(),
			timeToLive:   timeToLive,
			key:          key,
			value:        value,
		}
	}else {
		return &Record{
			creationTime: time.Now(),
			timeToLive:   defaultTTL,
			key:          key,
			value:        value,
		}
	}
}