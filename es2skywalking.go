package cls_skywalking_client_go

import (
	"errors"
	"github.com/liuyungen1988/go2sky"
	"github.com/liuyungen1988/go2sky/propagation"
	v3 "github.com/liuyungen1988/go2sky/reporter/grpc/language-agent"
	"github.com/labstack/echo/v4"
	"reflect"
	"time"
	_ "unsafe"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"net/url"
	"log"
)

type  EsSearchServiceProxy struct {
	searchService *elastic.SearchService
}

func NewEsSearchServiceProxy(searchService *elastic.SearchService) *EsSearchServiceProxy {
	return &EsSearchServiceProxy{
		searchService: searchService,
	}
}

//go:linkname buildURL github.com/olivere/elastic.(*SearchService).buildURL
func buildURL(s *elastic.SearchService) (string, url.Values, error)

func (s *EsSearchServiceProxy) Do(ctx context.Context) (*elastic.SearchResult, error) {
	if err := s.searchService.Validate(); err != nil {
		return nil, err
	}

	// Get URL for request
	path, _, err := buildURL(s.searchService)
	if err != nil {
		return nil, err
	}

	reqSpan, spanErr := StartSpantoSkyWalkingForES(path)
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb Prepare error: %v \n", spanErr)
	}

	result, err := s.searchService.Do(ctx)

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, path, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, path, true, nil)
	}

	return result , err
}

func StartSpantoSkyWalkingForES(queryStr string) (go2sky.Span, error) {
	originCtx := GetContext()
	if originCtx == nil {
		return nil, errors.New(fmt.Sprintf("can not get context, queryStr %s", queryStr))
	}
	ctx := originCtx.(echo.Context)
	// op_name 是每一个操作的名称
	tracerFromCtx := ctx.Get("tracer")
	if tracerFromCtx == nil {
		return nil, errors.New(fmt.Sprintf("can not get tracer, queryStr %s", queryStr))
	}
	tracer := tracerFromCtx.(*go2sky.Tracer)
	reqSpan, err := tracer.CreateExitSpan(ctx.Request().Context(), queryStr, "ES", func(header string) error {
		if(reflect.TypeOf(ctx.Get("header")) != nil) {
			ctx.Get("header").(*SafeHeader).Set(propagation.Header, header)
		}

		return nil
	})

	if(err != nil) {
		return nil, errors.New(fmt.Sprintf("StartSpantoSkyWalkingForES CreateExitSpan error: %s", err))
	}

	reqSpan.SetComponent(5)
	reqSpan.SetSpanLayer(v3.SpanLayer_Http) // rpc 调用
	reqSpan.Log(time.Now(), "[ESRequest]", fmt.Sprintf("开始请求,请求服务:%s,请求地址:%s", "ES", queryStr))

	return reqSpan, err
}

func EndSpantoSkywalkingForEs(reqSpan go2sky.Span, queryStr string, isNormal bool, err error) {
	if reqSpan == nil {
		return
	}
	reqSpan.Tag(go2sky.TagURL, queryStr)
	if !isNormal {
		reqSpan.Error(time.Now(), "[ES Response]", fmt.Sprintf("结束请求,响应结果: %s", err))
	} else {
		reqSpan.Log(time.Now(), "[ES Response]", "结束请求")
	}
	reqSpan.End()
}





