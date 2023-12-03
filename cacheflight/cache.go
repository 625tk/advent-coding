package cacheflight

import (
	"errors"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

var (
	ErrExpired  = errors.New("expired")
	ErrNotfound = errors.New("not found")
)

type Cache struct {
	v map[string]cacheVal
	m sync.RWMutex
	g singleflight.Group
}

type cacheVal struct {
	v   any
	exp int64
}

func New() Cache {
	return Cache{
		v: map[string]cacheVal{},
		m: sync.RWMutex{},
		g: singleflight.Group{},
	}
}

func (r *Cache) Get(key string, getFunc func() (v any, err error), expire time.Duration) (any, error) {
	// get
	val, err := r.get(key, time.Now().Unix())
	if err == nil {
		return val, nil
	}

	// update
	res, dErr, _ := r.g.Do(key, func() (interface{}, error) {
		got, err := getFunc()
		if err != nil {
			return nil, err
		}
		r.m.Lock()
		defer r.m.Unlock()
		r.v[key] = cacheVal{
			v:   got,
			exp: time.Now().Add(expire).Unix(),
		}
		return got, nil
	})

	if dErr != nil {
		if errors.Is(err, ErrExpired) {
			return val, ErrExpired
		} else {
			return nil, dErr
		}
	}

	return res, nil
}

func (r *Cache) get(key string, now int64) (any, error) {
	r.m.Lock()
	defer r.m.Unlock()
	if v, ok := r.v[key]; ok {
		if now < v.exp {
			return v.v, nil
		} else {
			return v.v, ErrExpired
		}
	}
	return nil, ErrNotfound
}
