package cls_skywalking_client_go

import (
	"github.com/liuyungen1988/go2sky"
	"log"
	"strings"
	"time"
	"github.com/labstack/echo/v4"
)

func Log(str ...string) {
	LogWithSearch("", str...)
}

/**
searchableTagKeys: userId=a,telephone=12333333
*/
func LogWithSearch(searchableTagKeys string, str ...string) {
	originCtx := GetContext()
	if originCtx == nil {
		log.Printf("can not get Context\n")
		return
	}
	ctx := originCtx.(echo.Context)
	tracerFromCtx := ctx.Get("tracer")
	if tracerFromCtx == nil {
		log.Printf("can not get tracer\n")
		return
	}
	tracer := tracerFromCtx.(*go2sky.Tracer)

	subSpan, _, err := tracer.CreateLocalSpan(ctx.Request().Context(), go2sky.WithOperationName("Log Info"))

	if err != nil {
		log.Printf("can not CreateLocalSpan\n")
		return
	}

	subSpan.Log(time.Now(), searchableTagKeys)
	subSpan.Log(time.Now(), str...)
	addTags(subSpan, searchableTagKeys)

	subSpan.End()
}

func Error(str ...string) {
	ErrorWithSearch("", str...)
}

/**
searchableTagKeys: userId=a,phone=12333333
*/
func ErrorWithSearch(searchableTagKeys string, str ...string) {
	originCtx := GetContext()
	if originCtx == nil {
		log.Printf("can not get Context\n")
		return
	}
	ctx := originCtx.(echo.Context)
	tracerFromCtx := ctx.Get("tracer")
	if tracerFromCtx == nil {
		log.Printf("can not get tracer\n")
		return
	}
	tracer := tracerFromCtx.(*go2sky.Tracer)

	subSpan, _, err := tracer.CreateLocalSpan(ctx.Request().Context(), go2sky.WithOperationName("Log Error"))

	if err != nil {
		log.Printf("can not CreateLocalSpan\n")
		return
	}

	subSpan.Error(time.Now(), searchableTagKeys)
	subSpan.Error(time.Now(), str...)
	addTags(subSpan, searchableTagKeys)
	subSpan.End()
}

func addTags(span go2sky.Span, searchableTagKeys string) {
	if len(searchableTagKeys) == 0 || strings.Index(searchableTagKeys, "=") == -1 {
		return
	}
	tagKeys := strings.Split(searchableTagKeys, ",")
	for tagKeyIndex := range tagKeys {
		tagKeyAndValue := strings.Split(tagKeys[tagKeyIndex], "=")
		if len(tagKeyAndValue) == 2 {
			span.Tag(go2sky.Tag(tagKeyAndValue[0]), tagKeyAndValue[1])
		}
	}
}
