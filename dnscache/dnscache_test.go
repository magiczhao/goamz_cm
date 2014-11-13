package dnscache

import (
    "testing"
)

func TestDNSCacheRefresh(t *testing.T) {
    cache := NewDNSCache("http://www.visenze.com")
    if cache.GetIp() == "" {
        t.Error("Cache ip failed!")
    }

    last_modify := cache.last_update

    if cache.access_count != 1 {
        t.Error("access_count failed!")
    }
    for i := 0; i < 10; i++ {
        cache.GetIp()
    }
    if cache.access_count != 11 {
        t.Error("access_count failed!")
    }

    for i := 0; i < 100; i++ {
        cache.GetIp()
    }

    cache.Refresh()
    if cache.last_update.Sub(last_modify).Nanoseconds() == 0 {
        t.Error("Update failed!")
    }
}
