# 一些流行的go、Rust、Nodejs的Web服务端框架的基准测试

### 做一些简单的测试，方便了解每种框架在一般情况下能达到的qps水平
### 说明
- go-zero或者kratos这样的重型框架来跟什么都没有的框架比，多少有点不公平。这里你只需要知道自己在用这个框架的时候，单机这个框架能达到什么样的大概水平就行。

### 包含
- gin (go)
- fiber (go)
- koa (Nodejs)
- net/http (go)
- flask (python)
- go-zero (go)
- kratos (go)
- hertz (go)
- actix-web (Rust)
- rocket (Rust)
- pingora (Rust)


### ben-reports下存放一些测试报告
- 01.腾讯云2C2G的Ecs的Localhost(本地请求)Benchmarks
- 02.阿里云2C2G的Ecs的Localhost(本地请求)Benchmarks
- 03.腾讯云4C4G的Ecs的Localhost(本地请求)Benchmarks
- 04.腾讯云4C4G的Ecs的同机房(同一城市同一地区)Benchmarks
- 05.腾讯云16C16G的Ecs的同机房(同一城市同一地区)Benchmarks
- 06.腾讯云64C64G的Ecs的同机房(同一城市同一地区)Benchmarks
- 07.阿里云64C64G的Ecs的同机房(同一城市同一地区)Benchmarks
- 08.如何优化腾讯云
- 09.如何优化阿里云



