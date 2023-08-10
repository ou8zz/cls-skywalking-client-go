cls-skywalking-client-go


release/v0.1.3
 增加标记搜索app=xx及version=xx
 
 release/v0.1.4
  升级go2sky到版本0.6.4
  
 release/v0.1.5
    将引用的github.com/SkyAPM/go2sky替换为 "codehub-cn-east-2.devcloud.huaweicloud.com/jgz00001/go2sky.git"
    
 release/v0.1.6
    增加采样率： USE_SKYWALKING_SAMPLE 可设置0.0~1.0
    
 release/v0.1.7  
    解决执行doClearContextAtRegularTime()时报错：panic: interface conversion: interface {} is nil, not time.Time
    
release/v0.1.8  
    1. echo的context运行一段时间后会reset, 导致获取其中的一些值为空， 原因未知。 暂时用判空方式解决。 
    2. 增加grpc的超时时间设置。
    
release/v0.1.9
    1. 根据设备号需要能获取tracing IOS cuid 解决， Android后续解决。
    2. 响应小于xxx打印响应。
    3. 修改判断错误的条件。
    
