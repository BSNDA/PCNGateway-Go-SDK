package cache

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/cache/proc"
	"time"
)

const (
	Expiration = "0.5h"
	Interval   = "3s"
)

var ca *proc.Cache

func InitCache() {
	// Default expiration time
	defaultExpiration, _ := time.ParseDuration(Expiration)
	//Recovery time, recovery expired
	gcInterval, _ := time.ParseDuration(Interval)
	ca = proc.NewCache(defaultExpiration, gcInterval)

}

//Set a sliding expiration cashe expiration time
func SetSlideValue(key string, value interface{}, expiration string) {
	if ca == nil {
		InitCache()
	}
	dur, _ := time.ParseDuration(expiration)
	ca.Set(key, value, dur, true)
}

//Set a cache with an expiration date
func SetValueByExpiration(key string, value interface{}, expiration string) {
	if ca == nil {
		InitCache()
	}
	expi, _ := time.ParseDuration(expiration)
	ca.Set(key, value, expi, false)
}

//Set a cache that never expires
func SetValue(key string, value interface{}) {
	if ca == nil {
		InitCache()
	}
	ca.Set(key, value, 0, false)
}

//Get the configured value
func GetValue(key string) (interface{}, bool) {
	if ca == nil {
		return nil, false
	}
	return ca.Get(key)

}

func ClearKey(key string) {
	ca.Delete(key)
}
