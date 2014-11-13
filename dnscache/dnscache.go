package dnscache

import (
    "time"
    "net"
    "net/url"
)
type DNSCache struct {
    endpoint_host string
    endpoint_ip string
    last_update time.Time
    access_count int
}

func NewDNSCache(endpoint_host string) *DNSCache {
    dns_cache := &DNSCache{endpoint_host,
        "",
        time.Time{},
        0,
    }
    dns_cache.Refresh()
    return dns_cache
}

func (d *DNSCache) Refresh() {
    now := time.Now()
    if now.Sub(d.last_update).Seconds() < 30 && d.access_count < 100 {
        // less than 30 second, use cached result or
        // used less than 100 times, use cached result
        return
    }
    d.access_count = 0
    u, err := url.Parse(d.endpoint_host)
    if err != nil {
        return
    }
    ip, err := net.LookupHost(u.Host)
    if err != nil {
        return
    }
    // always use the first ip address
    if len(ip) > 0 {
        d.endpoint_ip = ip[0]
        d.last_update = now
    }
}

func (d *DNSCache) GetIp() string {
    d.access_count += 1
    return d.endpoint_ip
}
