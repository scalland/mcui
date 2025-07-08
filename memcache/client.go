package memcache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var Client *memcache.Client

func Init(server string) {
	Client = memcache.New(server)
}
