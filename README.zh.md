# 一些流行的go、Rust、Nodejs的Web服务端框架的基准测试

### 做一些简单的测试，方便了解每种框架在一般情况下能达到的qps水平
### 说明
- go-zero或者kratos这样的重型框架来跟什么都没有的框架比，多少有点不公平。这里你只需要知道自己在用这个框架的时候，单机这个框架能达到什么样的大概水平就行。

### `bench-reports/01`包含如下框架的测试
- gin (go)
- fiber (go)
- koa (Nodejs)
- go-zero (go)
- actix-web (Rust)
- tokamak (zig)
- flask (python)
- kratos (go)
- hertz (go)
- rocket (Rust)
- salvo (Rust)
- net/http (go)
- zap (zig)

### `bench-reports/02`之后只包含如下框架的测试
- gin (go)
- fiber (go)
- actix-web (Rust)
- tokamak (zig)
- rocket (Rust)


### bench-reports下存放一些测试报告
- 01.腾讯云2C2G的Ecs的Localhost(本地请求)Benchmarks
- 02.腾讯云2C2G的Ecs的同机房(同一城市同一地区)Benchmarks
- 03.腾讯云4C4G的Ecs的同机房(同一城市同一地区)Benchmarks
- 04.阿里云4C4G的Ecs的同机房(同一城市同一地区)Benchmarks
- 05.腾讯云16C16G的Ecs的同机房(同一城市同一地区)Benchmarks
- 06.阿里云16C16G的Ecs的同机房(同一城市同一地区)Benchmarks
- 07.腾讯云16C16G的Ecs的同城不同机房(同一城市不同地区)Benchmarks
- 08.腾讯云16C16G的Ecs的异地机房(北京wrk->上海服务端)Benchmarks
- 09.腾讯云16C16G的Ecs的异地机房(北京wrk->广州服务端)Benchmarks
- 10.腾讯云16C16G的Ecs的海外机房(北京wrk->新加坡服务端)Benchmarks
- 11.腾讯云16C16G的Ecs的海外机房(北京wrk->硅谷服务端)Benchmarks
- 12.如何优化腾讯云
- 13.如何优化阿里云
