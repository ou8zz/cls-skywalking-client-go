package cls_skywalking_client_go

import (
	"sync"
	"net/http"
)

type SafeHeader struct {
	sync.RWMutex
	header http.Header
}

func newSafeHeader(inputHeader http.Header ) *SafeHeader {
	sm := new(SafeHeader)
	sm.header = inputHeader
	return sm

}

func (sh *SafeHeader) Get(key string) string {
	sh.RLock()
	value := sh.header.Get(key)
	sh.RUnlock()
	return value
}

func (sh *SafeHeader) Set(key string, value string) {
	sh.Lock()
	sh.header.Set(key, value)
	sh.Unlock()
}
