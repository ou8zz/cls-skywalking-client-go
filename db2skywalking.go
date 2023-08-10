package cls_skywalking_client_go

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"errors"
	"reflect"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/liuyungen1988/go2sky"
	"github.com/liuyungen1988/go2sky/propagation"
	v3 "github.com/liuyungen1988/go2sky/reporter/grpc/language-agent"
	"os"
	"github.com/labstack/echo/v4"
)

var (
    dbProxyMap = make(map[interface{}]string)
)


type DbProxy struct {
	Db sq.DBProxy
}

// “构造基类”
func NewDbProxy(db sq.DBProxy) *DbProxy {
	return &DbProxy{
		Db: db,
	}
}

func (f DbProxy) getDb() sq.DBProxy {
	return f.Db
}

//func (f DbProxy) Exec(insert interface) (sql.Result, error) {
//
//	result, err := f.getDb().Exec(insert)
//}

func PutDsn(Db sq.DBProxy, dsn string) {
	dbProxyMap[Db] = dsn
}

func GetDsn(Db sq.DBProxy) string {
	dsn := dbProxyMap[Db]
	if len(dsn) == 0 {
		dsn = os.Getenv("DB_URL")
	}
	return dsn
}

func (f DbProxy) Prepare(query string) (*sql.Stmt, error) {
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(query, GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb Prepare error: %v \n", spanErr)
	}

	result, err := f.getDb().Prepare(query)

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, query, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, query, true, nil)
	}

	return result, err
}

func (f DbProxy) Update(update sq.UpdateBuilder) (sql.Result, error) {
	updateStr, args, _ := update.ToSql()
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(updateStr+fmt.Sprintf("\r\n Parameters: %+v", args), GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb Update error: %v \n", spanErr)
	}

	result, err := update.RunWith(f.getDb()).Exec()

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, updateStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, updateStr, true, nil)
	}

	return result, err
}

func (f DbProxy) Delete(delete sq.DeleteBuilder) (sql.Result, error) {
	deleteStr, args, _ := delete.ToSql()
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(deleteStr+fmt.Sprintf("\r\n Parameters: %+v", args),  GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb Delete error: %v \n", spanErr)
	}

	result, err := delete.RunWith(f.getDb()).Exec()

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, deleteStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, deleteStr, true, nil)
	}

	return result, err
}

func (f DbProxy) Insert(insert sq.InsertBuilder) (sql.Result, error) {
	insertStr, args, _ := insert.ToSql()
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(insertStr+fmt.Sprintf("\r\n Parameters: %+v", args),  GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb Insert error: %v \n", spanErr)
	}

	result, err := insert.RunWith(f.getDb()).Exec()

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, insertStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, insertStr, true, nil)
	}

	return result, err
}

func (f DbProxy) Exec(queryStr string, args ...interface{}) (sql.Result, error) {
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr+fmt.Sprintf("\r\n Parameters: %+v", args),  GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb exec error: %v \n", spanErr)
	}

	result, err := f.getDb().Exec(queryStr, args...)

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, true, nil)
	}

	return result, err
}

func (f DbProxy) ExecWith(s sq.Sqlizer) (sql.Result, error) {
	query, args, err := s.ToSql()
	if err != nil {
		return nil, err
	}

	reqSpan, spanErr := StartSpantoSkyWalkingForDb(query+fmt.Sprintf("\r\n Parameters: %+v", args),  GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb exec error: %v \n", spanErr)
	}

	result, err := sq.ExecWith(f.getDb(), s)

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, query, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, query, true, nil)
	}

	return result, err
}

func (f DbProxy) QueryRowByStr(query string, args ...interface{}) squirrel.RowScanner {

	reqSpan, spanErr := StartSpantoSkyWalkingForDb(fmt.Sprintf(query+"\r\n Parameters%+v: ", args),  GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb error: %v \n", spanErr)
	}
	rowScanner := f.getDb().QueryRow(query, args...)

	EndSpantoSkywalkingForDb(reqSpan, query, true, nil)

	return rowScanner
}

func (f DbProxy) QueryRow(query squirrel.SelectBuilder) squirrel.RowScanner {
	queryStr, args, _ := query.ToSql()

	var temp = make([]string, len(args))
	for k, v := range args {
		temp[k] = fmt.Sprintf("%d", v)
	}
	var result = "[" + strings.Join(temp, ",") + "]"

	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr+"\r\n Parameters: "+result,  GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb error: %v \n", spanErr)
	}

	rowScanner := query.RunWith(f.getDb()).QueryRow()

	EndSpantoSkywalkingForDb(reqSpan, queryStr, true, nil)

	return rowScanner
}

func (f DbProxy) QueryByStr(query string, args ...interface{}) (*sql.Rows, error) {
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(fmt.Sprintf(query+"\r\n Parameters%+v: ", args),  GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb error: %v \n", spanErr)
	}

	rows, err := f.getDb().Query(query, args...)

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, query, false, err)
	}

	EndSpantoSkywalkingForDb(reqSpan, query, true, err)

	return rows, err
}

func (f DbProxy) Query(query squirrel.SelectBuilder) (*sql.Rows, error) {
	queryStr, args, _ := query.ToSql()

	var temp = make([]string, len(args))
	for k, v := range args {
		temp[k] = fmt.Sprintf("%d", v)
	}
	var result = "[" + strings.Join(temp, ",") + "]"

	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr+"\r\n Parameters: "+result, GetDsn(f.getDb()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb error: %v \n", spanErr)
	}

	rows, err := query.RunWith(f.getDb()).Query()

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, false, err)
	}

	EndSpantoSkywalkingForDb(reqSpan, queryStr, true, err)

	return rows, err
}

func StartSpantoSkyWalkingForDb(queryStr string, db string) (go2sky.Span, error) {
	db = GetDbUrl(db)
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
	reqSpan, err := tracer.CreateExitSpan(ctx.Request().Context(), queryStr, db, func(header string) error {
		if(reflect.TypeOf(ctx.Get("header")) != nil) {
			ctx.Get("header").(*SafeHeader).Set(propagation.Header, header)
		}

		return nil
	})

	if(err != nil) {
		return nil, errors.New(fmt.Sprintf("StartSpantoSkyWalkingForDb CreateExitSpan error: %s", err))
	}

	reqSpan.SetComponent(5)
	reqSpan.SetSpanLayer(v3.SpanLayer_Database) // rpc 调用
	reqSpan.Log(time.Now(), "[DBRequest]", fmt.Sprintf("开始请求,请求服务:%s,请求地址:%s", db, queryStr))

	return reqSpan, err
}

func EndSpantoSkywalkingForDb(reqSpan go2sky.Span, queryStr string, isNormal bool, err error) {
	if reqSpan == nil {
		return
	}
	reqSpan.Tag(go2sky.TagDBType, "MySql")
	reqSpan.Tag(go2sky.TagURL, queryStr)
	if !isNormal {
		reqSpan.Error(time.Now(), "[DB Response]", fmt.Sprintf("结束请求,响应结果: %s", err))
	} else {
		reqSpan.Log(time.Now(), "[DB Response]", "结束请求")
	}
	reqSpan.End()
}

func GetDbUrl(dbUrl string) string {
	if dbUrl != "" {
		start := strings.Index(dbUrl, ":")
		end := strings.Index(dbUrl, "(")
		dbUrl = fmt.Sprintf(dbUrl[0:start]) + fmt.Sprintf(dbUrl[end:len(dbUrl)])
	}
	return dbUrl
}
