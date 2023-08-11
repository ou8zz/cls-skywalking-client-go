package cls_skywalking_client_go

import (
	"net/http"
	"sync"
)

type SafeHeader struct {
	sync.RWMutex
	header http.Header
}

func newSafeHeader(inputHeader http.Header) *SafeHeader {
	sm := new(SafeHeader)
	sm.header = inputHeader
	return sm

}

func (sh *SafeHeader) Get(key string) string {
	sh.RLock()
	defer sh.RUnlock()
	value := sh.header.Get(key)
	return value
}

func (sh *SafeHeader) Set(key string, value string) {
	sh.Lock()
	defer sh.Unlock()
	sh.header.Set(key, value)
}
