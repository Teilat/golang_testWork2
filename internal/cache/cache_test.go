package cache

import (
	"github.com/stretchr/testify/assert"
	"src/golang_testWork2/internal/cache/record"
	"testing"
	"time"
)

func TestCache_Get(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]*record.Record
		key     string
		want    *record.Record
		wantErr bool
	}{
		{
			name: "case1",
			data: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"}},
			key:     "key2",
			want:    &record.Record{CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
			wantErr: false,
		},
		{
			name: "case2",
			data: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"}},
			key:     "key4",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Cache{
				data: tt.data,
			}
			got, err := v.Get(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCache_GetAll(t *testing.T) {

	tests := []struct {
		name string
		data map[string]*record.Record
		want []*record.Record
	}{
		{
			name: "case1",
			data: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"}},
			want: []*record.Record{
				{CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				{CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				{CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Cache{
				data: tt.data,
			}
			got := v.GetAll()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCache_Add(t *testing.T) {
	tests := []struct {
		name string
		data map[string]*record.Record
		rec  *record.Record
		want map[string]*record.Record
	}{
		{
			name: "case1",
			data: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"}},
			rec: &record.Record{CreationTime: time.Now(), TimeToLive: 400, Key: "key4", Value: "val"},
			want: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"},
				"key4": {CreationTime: time.Now(), TimeToLive: 400, Key: "key4", Value: "val"}},
		},
		{
			name: "case2",
			data: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"}},
			rec: &record.Record{CreationTime: time.Now(), TimeToLive: 400, Key: "key3", Value: "val343"},
			want: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 400, Key: "key3", Value: "val343"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Cache{
				data: tt.data,
			}
			v.Add(tt.rec)
			assert.Equal(t, tt.want, v.data)
		})
	}
}

func TestCache_Remove(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]*record.Record
		key     string
		want    map[string]*record.Record
		wantErr bool
	}{
		{
			name: "case1",
			data: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"}},
			key: "key1",
			want: map[string]*record.Record{
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"}},
			wantErr: false,
		},
		{
			name: "case2",
			data: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"}},
			key: "key4",
			want: map[string]*record.Record{
				"key1": {CreationTime: time.Now(), TimeToLive: 600, Key: "key1", Value: "val"},
				"key2": {CreationTime: time.Now(), TimeToLive: 200, Key: "key2", Value: "val"},
				"key3": {CreationTime: time.Now(), TimeToLive: 300, Key: "key3", Value: "val"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Cache{
				data: tt.data,
			}
			err := v.Remove(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, v.data)
		})
	}
}
