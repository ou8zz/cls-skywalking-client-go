package test

import (
	"github.com/liuyungen1988/cls-skywalking-client-go"
	"fmt"
	"testing"
)
func TestGetDbUrl(t *testing.T) {
	dbUrl :="cailianpress_dba:Cailianpress_888@tcp(172.21.0.120:3306)/cls_1508?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"

	result := cls_skywalking_client_go.GetDbUrl(dbUrl)

	expectecDbUrl :=  "cailianpress_dba(172.21.0.120:3306)/cls_1508?charset=utf8mb4&parseTime=True&loc=Asia%!F(MISSING)Shanghai"

	if result != expectecDbUrl{
		t.Error(fmt.Sprintf("error, result is %s", result))
	}
}

func TestGetDbUrl2(t *testing.T) {
	dbUrl :="cls_readonly:xxx@xxx#x7xCm@tcp(192.168.7.33:3306)/cls_1508?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai&readTimeout=15000s&timeout=5000s"
	result := cls_skywalking_client_go.GetDbUrl(dbUrl)

	expectecDbUrl :=  "cls_readonly(192.168.7.33:3306)/cls_1508?charset=utf8mb4&parseTime=True&loc=Asia%!F(MISSING)Shanghai&readTimeout=15000s&timeout=5000s"

	if result != expectecDbUrl{
		t.Error(fmt.Sprintf("error, result is %s", result))
	}
}

