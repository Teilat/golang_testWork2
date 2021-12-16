package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"src/golang_testWork2/cache"
	"src/golang_testWork2/cache/record"
	"time"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Handlers struct {
	cache cache.Cache
}

func New(cache *cache.Cache) *Handlers {
	return &Handlers{
		cache: *cache,
	}
}

func (h *Handlers) HandlerView() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)
		if args["key"] != "" {
			rec, err := h.cache.Get(args["key"])
			if err == nil {
				_, err := fmt.Fprintf(w, "%s:%s, %s\n", rec.Key, rec.Value, rec.CreationTime.Add(rec.TimeToLive).Sub(time.Now()))
				if err != nil {
					log.Println(err)
					return
				}
			} else {
				_, err := fmt.Fprintf(w, err.Error())
				if err != nil {
					log.Println(err)
					return
				}
			}
		} else {
			all := h.cache.GetAll()
			str := ""
			for _, d := range all {
				str = str + fmt.Sprintf("%s:%s, %s\n", d.Key, d.Value, d.CreationTime.Add(d.TimeToLive).Sub(time.Now()))
			}
			_, err := fmt.Fprintf(w, str)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (h *Handlers) HandlerAdd() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)
		if args["key"] != "" && args["value"] != "" {
			if args["duration"] != "" {
				ttl, err := time.ParseDuration(args["duration"] + "s")
				if err != nil {
					log.Println(err)
					_, err := fmt.Fprintf(w, err.Error())
					if err != nil {
						log.Println(err)
						return
					}
				}
				h.cache.Add(record.New(args["key"], args["value"], ttl))
			} else {
				h.cache.Add(record.New(args["key"], args["value"], time.Duration(0)))
			}
		} else {
			_, err := fmt.Fprintf(w, "No enough params\nUse /add?key=(yourKey)&val=(yourVal)")
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (h *Handlers) HandlerRemove() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)
		if args["key"] != "" {
			err := h.cache.Remove(args["key"])
			if err != nil {
				_, err := fmt.Fprintf(w, "error: %s", err)
				if err != nil {
					log.Println(err)
					return
				}
				return
			} else {
				_, err := fmt.Fprintf(w, "succes, key:%s", args["key"])
				if err != nil {
					log.Println(err)
					return
				}
			}
		} else {
			_, err := fmt.Fprintf(w, "No enough params\nUse /remove?key=(yourKey)")
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (h *Handlers) HandlerFlush() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.cache.Flush()
	}
}

func (h *Handlers) HandlerExel() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, filePath)
	}
}
