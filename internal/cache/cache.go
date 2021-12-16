package cache

import (
	"context"
	"fmt"
	"log"
	"src/golang_testWork2/internal/cache/record"
	"time"
)

type Cache struct {
	data    map[string]*record.Record
	context context.Context
	ticker  time.Ticker
}

func New(ctx context.Context, ticker time.Ticker) *Cache {
	return &Cache{
		data:    make(map[string]*record.Record, 0),
		context: ctx,
		ticker:  ticker,
	}
}

//Get возвращает запись по заданному ключу
func (v Cache) Get(key string) (*record.Record, error) {
	rec := v.data[key]
	rec.TimeToLive = rec.CreationTime.Add(rec.TimeToLive).Sub(time.Now())
	if rec != nil {
		return rec, nil
	} else {
		return nil, fmt.Errorf("невозможно получить запись. Запись с заданим ключем не существует, ключ:%s", key)
	}
}

//GetAll возвращает все записи из хранилища в виде списка
func (v Cache) GetAll() []*record.Record {
	res := make([]*record.Record, 0)
	for _, rec := range v.data {
		rec.TimeToLive = rec.CreationTime.Add(rec.TimeToLive).Sub(time.Now())
		res = append(res, rec)
	}
	return res
}

//Add добавляет запись в хранилище
func (v Cache) Add(rec *record.Record) {
	v.data[rec.Key] = rec
	log.Print("added key:", rec.Key)
}

//Remove удаляет значение по заданному ключу. При неудаче возвращает ошибку
func (v Cache) Remove(key string) error {
	rec := v.data[key]
	if rec != nil {
		delete(v.data, key)
		log.Print("removed key:", key)
		return nil
	} else {
		return fmt.Errorf("невозможно удалить запись. Запись с заданим ключем не существует, ключ:%s", key)
	}
}

//ClearCache очищает хранилище с оперативными данными
func (v Cache) ClearCache() {
	for k := range v.data {
		err := v.Remove(k)
		if err != nil {
			log.Print(err)
			return
		}
	}
}

//ProcessTimer запускается в горутине и проверяет каждую запись на истечение по времени
func (v Cache) ProcessTimer() {
	for {
		select {
		case <-v.ticker.C:
			go func() {
				for s, r := range v.data {
					isExpired := r.CreationTime.Add(r.TimeToLive).Before(time.Now())
					if isExpired {
						err := v.Remove(s)
						if err != nil {
							log.Print(err)
							return
						}
					}
				}
			}()
		case <-v.context.Done():
			return
		}
	}
}
