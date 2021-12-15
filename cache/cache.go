package cache

import (
	"fmt"
)

type Vault struct {
	data map[string]*Record
}

//Get возвращает запись по заданному ключу
func (v Vault) Get(key string) (*Record, error) {
	res := v.data[key]
	if res != nil {
		return res, nil
	} else {
		return nil, fmt.Errorf("невозможно получить запись. Запись с заданим ключем не существует, ключ:%s", key)
	}
}

//GetAll возвращает все записи из хранилища в виде списка
func (v Vault) GetAll() []*Record {
	res := make([]*Record,0)
	for _, i2 := range v.data {
		res = append(res, i2)
	}
	return res
}

//Add добавляет запись в хранилище
func (v Vault) Add(rec *Record) {
	v.data[rec.key] = rec
}

//Remove удаляет значение по заданному ключу. При неудаче возвращает ошибку
func (v Vault) Remove(key string) error {
	rec := v.data[key]
	if rec != nil {
		delete(v.data, key)
		return nil
	}else {
		return fmt.Errorf("невозможно удалить запись. Запись с заданим ключем не существует, ключ:%s", key)
	}
}

// Flush очищает хранилище с оперативными данными
func (v Vault) Flush() {
	v.data = make(map[string]*Record, 0)
}
