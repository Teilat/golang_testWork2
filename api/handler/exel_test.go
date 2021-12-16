package api

import (
	"log"
	"src/golang_testWork2/cache/record"
	"testing"
	"time"
)

//тест для ручной проверки созданого файла
func TestExcel(t *testing.T) {
	tests := []struct {
		name string
		rec  []*record.Record
	}{
		{
			name: "case1",
			rec: []*record.Record{
				{CreationTime: time.Now(), TimeToLive: 600000000000, Key: "key1", Value: "val1"},
				{CreationTime: time.Now(), TimeToLive: 25000000000, Key: "key2", Value: "val2"},
				{CreationTime: time.Now(), TimeToLive: 10000000000, Key: "key3", Value: "val3"},
				{CreationTime: time.Now(), TimeToLive: 40000000000, Key: "key4", Value: "val4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Excel(tt.rec)
			if err := f.SaveAs("test.xlsx"); err != nil {
				log.Fatal(err)
			}
		})
	}
}
