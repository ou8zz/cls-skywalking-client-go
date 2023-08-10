package cls_skywalking_client_go

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLogWithSearchUseRequestParamMap(t *testing.T) {
	var requestParamMap = make(map[string]string) /*创建集合 */
	requestParamMap["sv"] = "4.7.0"
	requestParamMap["app"] = "ios"
	requestParamMap["cuid"] = "ios 设备号"

	searchableKeys := logWithSearchUseRequestParamMap(requestParamMap)

	if searchableKeys != "sv=4.7.0,app=ios,cuid=ios 设备号" {
		t.Errorf("错误响应")
	}
}

func TestLogWithSearchUseRequestParamMapWithoutSv(t *testing.T) {
	var requestParamMap = make(map[string]string) /*创建集合 */
	requestParamMap["app"] = "ios"
	requestParamMap["cuid"] = "ios 设备号"

	searchableKeys := logWithSearchUseRequestParamMap(requestParamMap)

	if searchableKeys != "app=ios,cuid=ios 设备号" {
		t.Errorf("错误响应")
	}
}

func TestLogWithSearchUseRequestParamMapWithoutData(t *testing.T) {
	var requestParamMap = make(map[string]string) /*创建集合 */

	searchableKeys := logWithSearchUseRequestParamMap(requestParamMap)

	if searchableKeys != "" && len(searchableKeys) == 0 {
		t.Errorf("错误响应")
	}
}

func TestIsNull(t *testing.T) {
	fmt.Println("nil:", reflect.TypeOf("fff") != nil)
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())
	fmt.Println("ddd")
}
