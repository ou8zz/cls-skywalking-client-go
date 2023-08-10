package cls_skywalking_client_go

import (
	"net/http"
	"time"
	"github.com/liuyungen1988/cls-skywalking-client-go/util"

	"errors"
	"fmt"
	"github.com/liuyungen1988/go2sky"
	"github.com/liuyungen1988/go2sky/propagation"
	v3 "github.com/liuyungen1988/go2sky/reporter/grpc/language-agent"
	"github.com/labstack/echo/v4"
)

func StartSpantoSkyWalking(url string, params []string, remoteService string) (go2sky.Span, error) {
	originCtx := GetContext()
	if originCtx == nil {
		return nil, errors.New("can not get Context")
	}
	ctx := originCtx.(echo.Context)
	// op_name 是每一个操作的名称
	tracerFromCtx := ctx.Get("tracer")
	if tracerFromCtx == nil {
		return nil, errors.New("can not get tracer")
	}
	tracer := tracerFromCtx.(*go2sky.Tracer)
	reqSpan, err := tracer.CreateExitSpan(ctx.Request().Context(), url, remoteService, func(header string) error {
		if ctx.Get("header") != nil {
			ctx.Get("header").(*SafeHeader).Set(propagation.Header, header)
		}
		return nil
	})
	if(err != nil) {
		return nil, errors.New(fmt.Sprintf("StartSpantoSkyWalking CreateExitSpan error: %s", err))
	}
	reqSpan.SetComponent(2)                 //HttpClient,看 https://github.com/apache/skywalking/blob/master/docs/en/guides/Component-library-settings.md ， 目录在component-libraries.yml文件配置
	reqSpan.SetSpanLayer(v3.SpanLayer_Http) // rpc 调用
	reqSpan.Log(time.Now(), "[HttpRequest]", fmt.Sprintf("开始请求,请求服务:%s,请求地址:%s,请求参数:%+v", remoteService, util.ReplaceAccessKeyId(url), params))

	return reqSpan, err
}

func EndSpantoSkywalking(reqSpan go2sky.Span, url string, resp string, isNormal bool, err error) {
	if reqSpan == nil {
		return
	}
	reqSpan.Tag(go2sky.TagHTTPMethod, http.MethodPost)
	reqSpan.Tag(go2sky.TagURL, url)
	if !isNormal {
		reqSpan.Error(time.Now(), "[Http Request]", fmt.Sprintf("结束请求,返回异常: %s", err.Error()))
	} else {
		reqSpan.Log(time.Now(), "[Http Response]", fmt.Sprintf("结束请求,响应结果: %s", resp))
	}
	reqSpan.End()
}
