package cls_skywalking_client_go

import (
	"fmt"
	"github.com/petermattis/goid"
	"log"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
	"github.com/labstack/echo/v4"
)

var (
	contexts = map[int64]echo.Context{}
	rwm      sync.RWMutex
)

// Set 设置一个 context
func SetContext(context echo.Context) {
	if context == nil {
		return
	}
	goID := getGoID()
	rwm.Lock()
	defer rwm.Unlock()

	context.Set("time", time.Now())
	contexts[goID] = context
}

// Get 返回设置的 context
func GetContext() echo.Context {
	goID := getGoID()
	rwm.RLock()
	defer rwm.RUnlock()

	return contexts[goID]
}

// Delete 删除设置的 RequestID
func DeleteContext() {
	goID := getGoID()
	rwm.Lock()
	defer rwm.Unlock()

	delete(contexts, goID)
}

func getGoID() int64 {
	return goid.Get()
}

func ClearContextAtRegularTime() {
	t := time.NewTicker(120 * time.Second)
	defer t.Stop()
	for {
		<-t.C
		doClearContextAtRegularTime()
		t.Reset(120 * time.Second)
	}
}

func doClearContextAtRegularTime() {
	rwm.Lock()
	defer rwm.Unlock()
	sm, _ := time.ParseDuration("-2m")
	timeBefore := time.Now().Add(sm)

	newContexts := map[int64]echo.Context{}
	for k, v := range contexts {
		if v == nil {
			delete(contexts, k)
			continue
		}
		vTime := v.Get("time")
		if vTime != nil {
			contextTime := vTime.(time.Time)
			if contextTime.Unix() < timeBefore.Unix() {
				delete(contexts, k)
			} else {
				newContexts[k] = v
			}

		} else {
			delete(contexts, k)
		}
	}

	contexts = nil
	runtime.GC()
	debug.FreeOSMemory()

	contexts = map[int64]echo.Context{}
	for k, v := range newContexts {
		contexts[k] = v
	}

	newContexts = nil
	runtime.GC()
	debug.FreeOSMemory()

	printMemStats()
	log.Printf(fmt.Sprintf("contexts left: %s \n", strconv.Itoa(len(contexts))))
}


func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}

