# Go-Asyncer

`Go-Asyncer` leverages `github.com/hibiken/asynq` to run asynchronous tasks. It essentially has two functions that wraps around `asynq`'s `Enqueue`method:

* `EnqueueTask` adds tasks to a queue and run them immediately if workers are available
* `ScheduleTask` adds tasks to a queue but doesn't run them immediately. It runs them at a specified time.

## Installation
```
go get github.com/allenakinkunle/go-asyncer
```