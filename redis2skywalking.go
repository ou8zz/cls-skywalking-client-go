package cls_skywalking_client_go

import (
	"fmt"
	"time"

	"errors"
	"github.com/liuyungen1988/go2sky"
	"github.com/liuyungen1988/go2sky/propagation"
	v3 "github.com/liuyungen1988/go2sky/reporter/grpc/language-agent"
	"strconv"
	"gopkg.in/redis.v5"
	"github.com/labstack/echo/v4"
)

type RedisProxy struct {
	RedisCache *redis.Client
}

// “构造基类”
func NewRedisProxy(redisCache *redis.Client) *RedisProxy {
	return &RedisProxy{
		RedisCache: redisCache,
	}
}

func (f RedisProxy) getRedisCache() *redis.Client {
	return f.RedisCache
}

func (f RedisProxy) Exists(key string) *redis.BoolCmd {
	span, _ := StartSpantoSkyWalkingForRedis("Exists "+key, f.getRedisCache().String())

	cmd := f.getRedisCache().Exists(key)

	_, err := cmd.Result()
	defer processResult(span, "Exists "+key,
		err)
	return cmd
}

func (f RedisProxy) Get(key string) *redis.StringCmd {
	span, _ := StartSpantoSkyWalkingForRedis("Get "+key, f.getRedisCache().String())

	cmd := f.getRedisCache().Get(key)

	_, err := cmd.Result()

	if(err != nil) {
	   errStr := fmt.Sprintf("%s", err)
	   if(errStr == "redis: nil"){
	   	  err = nil
	   }
	}
	defer processResult(span, "Get "+key,
		err)
	return cmd
}

func (f RedisProxy) GetRange(key string, start, end int64) *redis.StringCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("GetRange %s, start %s, end %s", key, strconv.FormatInt(start, 10), strconv.FormatInt(end, 10)), f.getRedisCache().String())

	cmd := f.getRedisCache().GetRange(key, start, end)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("GetRange %s, start %s, end %s", key, strconv.FormatInt(start, 10), strconv.FormatInt(end, 10)),
		err)
	return cmd
}

func (f RedisProxy) GetSet(key string, value interface{}) *redis.StringCmd {
	span, _ := StartSpantoSkyWalkingForRedis("GetSet "+key, f.getRedisCache().String())

	cmd := f.getRedisCache().GetSet(key, value)

	_, err := cmd.Result()
	defer processResult(span, "GetSet "+key,
		err)
	return cmd
}

func (f RedisProxy) MGet(keys ...string) *redis.SliceCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("MGet %v", keys), f.getRedisCache().String())

	cmd := f.getRedisCache().MGet(keys...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("MGet %v", keys),
		err)
	return cmd
}

func (f RedisProxy) MSet(pairs ...interface{}) *redis.StatusCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("MSet %v", pairs), f.getRedisCache().String())

	cmd := f.getRedisCache().MSet(pairs...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("MSet %v", pairs),
		err)
	return cmd
}

func (f RedisProxy) MSetNX(pairs ...interface{}) *redis.BoolCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("MSetNX %v", pairs), f.getRedisCache().String())

	cmd := f.getRedisCache().MSetNX(pairs...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("MSetNX %v", pairs),
		err)
	return cmd
}

// Redis `SET key value [expiration]` command.
//
// Use expiration for `SETEX`-like behavior.
// Zero expiration means the key has no expiration time.
func (f RedisProxy) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("Set %s, value %v", key,
		value), f.getRedisCache().String())

	cmd := f.getRedisCache().Set(key, value, expiration)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("Set %s, value %v", key,
		value),
		err)
	return cmd
}

func (f RedisProxy) SetRange(key string, offset int64, value string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("SetRange %s, offset %s,  value %s", key, strconv.FormatInt(offset, 10), value), f.getRedisCache().String())

	cmd := f.getRedisCache().SetRange(key, offset, value)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("SetRange %s, offset %s,  value %s", key, strconv.FormatInt(offset, 10), value),
		err)
	return cmd
}

func (f RedisProxy) HDel(key string, fields ...string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("HDel %s, fields %v", key, fields), f.getRedisCache().String())

	cmd := f.getRedisCache().HDel(key, fields...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("HDel %s, fields %v", key, fields),
		err)
	return cmd
}

func (f RedisProxy) HExists(key, field string) *redis.BoolCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("HExists %s, field %s", key, field), f.getRedisCache().String())

	cmd := f.getRedisCache().HExists(key, field)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("HExists %s, field %s", key, field),
		err)
	return cmd
}

func (f RedisProxy) HGet(key, field string) *redis.StringCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("HGet %s, field %s", key, field), f.getRedisCache().String())

	cmd := f.getRedisCache().HGet(key, field)

	_, err := cmd.Result()

	if(err != nil) {
		errStr := fmt.Sprintf("%s", err)
		if(errStr == "redis: nil"){
			err = nil
		}
	}
	defer processResult(span, fmt.Sprintf("HGet %s, field %s", key, field),
		err)
	return cmd
}

func (f RedisProxy) HGetAll(key string) *redis.StringStringMapCmd {
	span, _ := StartSpantoSkyWalkingForRedis("HGetAll  "+key, f.getRedisCache().String())

	cmd := f.getRedisCache().HGetAll(key)

	_, err := cmd.Result()
	defer processResult(span, "HGetAll  "+key,
		err)
	return cmd
}

func (f RedisProxy) HMGet(key string, fields ...string) *redis.SliceCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("HMGet %s, fields %v ", key, fields), f.getRedisCache().String())

	cmd := f.getRedisCache().HMGet(key, fields...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("HMGet %s, fields %v ", key, fields),
		err)
	return cmd
}

func (f RedisProxy) HMSet(key string, fields map[string]string) *redis.StatusCmd {
	span, _ := StartSpantoSkyWalkingForRedis("HMSet  "+key, f.getRedisCache().String())

	cmd := f.getRedisCache().HMSet(key, fields)

	_, err := cmd.Result()
	defer processResult(span, "HMGet  "+key,
		err)
	return cmd
}

func (f RedisProxy) HSet(key, field string, value interface{}) *redis.BoolCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("HSet %s, fields %s, value %v ", key, field, value), f.getRedisCache().String())

	cmd := f.getRedisCache().HSet(key, field, value)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("HSet %s, fields %s, value %v ", key, field, value),
		err)
	return cmd
}

func (f RedisProxy) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("HSetNX %s, fields %s, value %v ", key, field, value), f.getRedisCache().String())

	cmd := f.getRedisCache().HSetNX(key, field, value)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("HSetNX %s, fields %s, value %v ", key, field, value),
		err)
	return cmd
}

func (f RedisProxy) LPop(key string) *redis.StringCmd {
	span, _ := StartSpantoSkyWalkingForRedis("LPop  "+key, f.getRedisCache().String())

	cmd := f.getRedisCache().LPop(key)

	_, err := cmd.Result()
	defer processResult(span, "LPop  "+key,
		err)
	return cmd
}

func (f RedisProxy) LPush(key string, values ...interface{}) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("LPush %s, values %v ", key, values), f.getRedisCache().String())

	cmd := f.getRedisCache().LPush(key, values...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("LPush %s, values %v ", key, values),
		err)
	return cmd
}

func (f RedisProxy) LPushX(key string, value interface{}) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("LPushX %s, value %v ", key, value), f.getRedisCache().String())

	cmd := f.getRedisCache().LPushX(key, value)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("LPushX %s, value %v ", key, value),
		err)
	return cmd
}

func (f RedisProxy) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("LRange %s, start %s, stop %s ", key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)), f.getRedisCache().String())

	cmd := f.getRedisCache().LRange(key, start, stop)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("LRange %s, start %s, stop %s "+key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)),
		err)
	return cmd
}

func (f RedisProxy) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZRange %s, start %s, stop %s ", key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)), f.getRedisCache().String())

	cmd := f.getRedisCache().ZRange(key, start, stop)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZRange %s, start %s, stop %s ", key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)),
		err)
	return cmd
}

func (f RedisProxy) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZRangeWithScores %s, start %s, stop %s ", key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)), f.getRedisCache().String())

	cmd := f.getRedisCache().ZRangeWithScores(key, start, stop)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZRangeWithScores %s, start %s, stop %s ", key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)),
		err)
	return cmd
}

func (f RedisProxy) IncrBy(key string, value int64) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("IncrBy %s, value %s ", key, strconv.FormatInt(value, 10)), f.getRedisCache().String())

	cmd := f.getRedisCache().IncrBy(key, value)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("IncrBy %s, value %s ", key, strconv.FormatInt(value, 10)),
		err)
	return cmd
}

func (f RedisProxy) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("HIncrBy %s, field %s, incr %s ", key, field, strconv.FormatInt(incr, 10)), f.getRedisCache().String())

	cmd := f.getRedisCache().HIncrBy(key, field, incr)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("HIncrBy %s, field %s, incr %s ", key, field, strconv.FormatInt(incr, 10)),
		err)
	return cmd
}

func (f RedisProxy) ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {

	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZRangeByScore %s ", key), f.getRedisCache().String())

	cmd := f.getRedisCache().ZRangeByScore(key, opt)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZRangeByScore %s", key),
		err)
	return cmd

}

func (f RedisProxy) Pipelined(fn func(*redis.Pipeline) error) ([]redis.Cmder, error) {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("start Pipelined"), f.getRedisCache().String())
	defer  EndSpantoSkywalkingForRedis(span, "end Pipelined", true, nil)
	return f.getRedisCache().Pipelined(fn)
}

func (f RedisProxy) Del(keys ...string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("Del keys: %v ", keys), f.getRedisCache().String())

	cmd := f.getRedisCache().Del(keys...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("Del keys %v", keys),
		err)
	return cmd


}

func (f RedisProxy) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZRemRangeByScore key: %s, min: %s, max: %s ", key, min, max), f.getRedisCache().String())

	cmd := f.getRedisCache().ZRemRangeByScore(key, min, max)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZRemRangeByScore key: %s, min: %s, max: %s ", key, min, max),
		err)
	return cmd
}

func (f RedisProxy) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZAdd key: %s", key), f.getRedisCache().String())


	cmd := f.getRedisCache().ZAdd(key, members...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZAdd key: %s", key),
		err)
	return cmd
}

func (f RedisProxy) ZCard(key string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZCard key: %s", key), f.getRedisCache().String())

	cmd := f.getRedisCache().ZCard(key)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZCard key: %s", key),
		err)
	return cmd
}

func (f RedisProxy) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZRemRangeByRank key: %s, start:%s, stop:%s", key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)), f.getRedisCache().String())

	cmd := f.getRedisCache().ZRemRangeByRank(key, start, stop)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZRemRangeByRank key: %s, start:%s, stop:%s", key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)),
		err)
	return cmd
}

func (f RedisProxy) ZRem(key string, members ...interface{}) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZRem key: %s", key), f.getRedisCache().String())

	cmd:=  f.getRedisCache().ZRem(key, members...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZRem key: %s", key),
		err)
	return cmd
}

func (f RedisProxy)  ZRank(key, member string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZRank key: %s, member:%s", key, member), f.getRedisCache().String())

	cmd:=  f.getRedisCache().ZRank(key, member)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZRank key: %s, member:%s", key, member),
		err)
	return cmd
}

func (f RedisProxy) Decr(key string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("Decr key: %s", key), f.getRedisCache().String())

	cmd:=  f.getRedisCache().Decr(key)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("Decr key: %s", key),
		err)
	return cmd
}

func (f RedisProxy) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZRevRange %s, start %s, stop %s ", key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)), f.getRedisCache().String())

	cmd:=  f.getRedisCache().ZRevRange(key, start, stop)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZRevRange %s, start %s, stop %s ", key, strconv.FormatInt(start, 10), strconv.FormatInt(stop, 10)),
		err)
	return cmd
}

func (f RedisProxy) ZCount(key, min, max string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZCount %s, min %s, max %s ", key, min, max), f.getRedisCache().String())

	cmd := f.getRedisCache().ZCount(key, min, max)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZCount %s, min %s, max %s ", key, min, max),
	err)
	return cmd
}

func (f RedisProxy) SAdd(key string, members ...interface{}) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("SAdd %s ", key), f.getRedisCache().String())

	cmd := f.getRedisCache().SAdd(key,  members...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("SAdd %s ", key),
		err)
	return cmd
}

func (f RedisProxy) SRem(key string, members ...interface{}) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("SRem %s ", key), f.getRedisCache().String())

	cmd := f.getRedisCache().SRem(key,  members...)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("SRem %s ", key),
		err)
	return cmd
}

func (f RedisProxy)  LLen(key string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("LLen %s ", key), f.getRedisCache().String())

	cmd := f.getRedisCache().LLen(key)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("LLen %s ", key),
		err)
	return cmd
}

func (f RedisProxy) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("LRem %s, count: %s, value: %S ", key, strconv.FormatInt(count, 10), value), f.getRedisCache().String())

	cmd := f.getRedisCache().LRem(key, count, value)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("LRem %s, count: %s, value: %S ", key, strconv.FormatInt(count, 10), value),
		err)
	return cmd
}



func (f RedisProxy) TTL(key string) *redis.DurationCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("TTL key: %s", key), f.getRedisCache().String())

	cmd := f.getRedisCache().TTL(key)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("TTL key: %s", key),
		err)
	return cmd
}

func (f RedisProxy) Incr(key string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("Incr key: %s", key), f.getRedisCache().String())

	cmd := f.getRedisCache().Incr(key)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("Incr key: %s", key),
		err)
	return cmd
}

func (f RedisProxy) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("Expire key: %s", key), f.getRedisCache().String())

	cmd := f.getRedisCache().Expire(key, expiration)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("Expire key: %s", key),
		err)
	return cmd
}

func (f RedisProxy) Keys(pattern string) *redis.StringSliceCmd {

	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("Keys pattern: %s", pattern), f.getRedisCache().String())

	cmd := f.getRedisCache().Keys(pattern)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("Keys pattern: %s", pattern),
		err)
	return cmd
}


func (f RedisProxy) ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("ZRevRangeByScore key: %s", key), f.getRedisCache().String())

	cmd :=  f.getRedisCache().ZRevRangeByScore(key, opt)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("ZRevRangeByScore key: %s", key),
		err)
	return cmd
}

func (f RedisProxy) Publish(channel, message string) *redis.IntCmd {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("Publish channel: %s, message: %s", channel, message), f.getRedisCache().String())

	cmd :=  f.getRedisCache().Publish(channel, message)

	_, err := cmd.Result()
	defer processResult(span, fmt.Sprintf("Publish channel: %s, message: %s", channel, message),
		err)
	return cmd
}

func (f RedisProxy) Pipeline() *redis.Pipeline {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("start Pipeline"), f.getRedisCache().String())
	defer  EndSpantoSkywalkingForRedis(span, "end Pipeline", true, nil)

	return f.getRedisCache().Pipeline()
}

func  (f RedisProxy) Close() error {
	span, _ := StartSpantoSkyWalkingForRedis(fmt.Sprintf("start Close"), f.getRedisCache().String())
	defer  EndSpantoSkywalkingForRedis(span, "end Close", true, nil)
	return f.getRedisCache().Close()
}


func StartSpantoSkyWalkingForRedis(queryStr string, db string) (go2sky.Span, error) {
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
	reqSpan, err := tracer.CreateExitSpan(ctx.Request().Context(), queryStr, db, func(header string) error {
		if ctx.Get("header") != nil {
			ctx.Get("header").(*SafeHeader).Set(propagation.Header, header)
		}
		return nil
	})
	if(err != nil) {
		return nil, errors.New(fmt.Sprintf("StartSpantoSkyWalkingForRedis CreateExitSpan error: %s", err))
	}
	reqSpan.SetComponent(7)
	reqSpan.SetSpanLayer(v3.SpanLayer_Cache) // cache
	reqSpan.Log(time.Now(), "[Redis Request]", fmt.Sprintf("开始请求,请求服务:%s,请求地址:%s", db, queryStr))

	return reqSpan, err
}

func EndSpantoSkywalkingForRedis(reqSpan go2sky.Span, queryStr string, isNormal bool, err error) {
	if reqSpan == nil {
		return
	}
	reqSpan.Tag(go2sky.TagURL, queryStr)
	if !isNormal {
		reqSpan.Error(time.Now(), "[Redis Response]", fmt.Sprintf("结束请求,响应结果: %s", err))
	} else {
		reqSpan.Log(time.Now(), "[Redis Response]", "结束请求")
	}
	reqSpan.End()
}

func processResult(span go2sky.Span, str string, err error) {
	if err == nil {
		EndSpantoSkywalkingForRedis(span, str, true, nil)
	} else {
		EndSpantoSkywalkingForRedis(span, str, false, err)
	}
}
