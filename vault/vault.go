package vault

import (
	"context"
	"fmt"
	"log"
	"src/golang_testWork2/vault/record"
	"time"
)

type Vault struct {
	data    map[string]*record.Record
	context context.Context
	ticker  time.Ticker
}

func NewVault(ctx context.Context, ticker time.Ticker) *Vault {
	return &Vault{
		data:    make(map[string]*record.Record, 0),
		context: ctx,
		ticker:  ticker,
	}
}

//Get возвращает запись по заданному ключу
func (v Vault) Get(key string) (*record.Record, error) {
	res := v.data[key]
	if res != nil {
		return res, nil
	} else {
		return nil, fmt.Errorf("невозможно получить запись. Запись с заданим ключем не существует, ключ:%s", key)
	}
}

//GetAll возвращает все записи из хранилища в виде списка
func (v Vault) GetAll() []*record.Record {
	res := make([]*record.Record, 0)
	for _, i2 := range v.data {
		res = append(res, i2)
	}
	return res
}

//Add добавляет запись в хранилище
func (v Vault) Add(rec *record.Record) {
	v.data[rec.Key] = rec
	log.Print("added key:", rec.Key)
}

//Remove удаляет значение по заданному ключу. При неудаче возвращает ошибку
func (v Vault) Remove(key string) error {
	rec := v.data[key]
	if rec != nil {
		delete(v.data, key)
		log.Print("removed key:", key)
		return nil
	} else {
		return fmt.Errorf("невозможно удалить запись. Запись с заданим ключем не существует, ключ:%s", key)
	}
}

// Flush очищает хранилище с оперативными данными
func (v Vault) Flush() {
	v.data = make(map[string]*record.Record, 0)
}

//ProcessTimer запускается в горутине и проверяет каждую запись на истечение по времени
func (v Vault) ProcessTimer() {
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
				//log.Print("done")
			}()
		case <-v.context.Done():
			return
		}
	}
}
