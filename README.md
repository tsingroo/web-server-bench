# Benchmark Tests of Popular Go, Rust, and Node.js Web Server Frameworks

### Simple tests to understand the QPS level each framework can achieve under normal circumstances

### Notes
- Comparing heavy frameworks like go-zero or kratos with frameworks that have minimal features is somewhat unfair. Here you just need to know what approximate performance level a single machine can achieve when using this framework.

### `bench-reports/01` includes tests for the following frameworks
- gin (go)
- fiber (go)
- koa (Node.js)
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

### `bench-reports/02` and onwards only include tests for the following frameworks
- gin (go)
- fiber (go)
- actix-web (Rust)
- tokamak (zig)
- rocket (Rust)

### Test reports stored under bench-reports
- 01. Tencent Cloud 2C2G ECS Localhost (local requests) Benchmarks
- 02. Tencent Cloud 2C2G ECS Same Data Center (same city, same region) Benchmarks
- 03. Tencent Cloud 4C4G ECS Same Data Center (same city, same region) Benchmarks
- 04. Alibaba Cloud 4C4G ECS Same Data Center (same city, same region) Benchmarks
- 05. Tencent Cloud 16C16G ECS Same Data Center (same city, same region) Benchmarks
- 06. Alibaba Cloud 16C16G ECS Same Data Center (same city, same region) Benchmarks
- 07. Tencent Cloud 16C16G ECS Same City Different Data Center (same city, different region) Benchmarks
- 08. Tencent Cloud 16C16G ECS Different Region Data Center (Beijing wrk -> Shanghai server) Benchmarks
- 09. Tencent Cloud 16C16G ECS Different Region Data Center (Beijing wrk -> Guangzhou server) Benchmarks
- 10. Tencent Cloud 16C16G ECS Overseas Data Center (Beijing wrk -> Singapore server) Benchmarks
- 11. Tencent Cloud 16C16G ECS Overseas Data Center (Beijing wrk -> Silicon Valley server) Benchmarks
- 12. How to Optimize Tencent Cloud
- 13. How to Optimize Alibaba Cloud
