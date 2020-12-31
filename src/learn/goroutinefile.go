package main

import (
    "fmt"
    "net/url"
)

func Producer(factor int, out chan<- int) {
    for i := 0; ; i++ {
        out <- i*factor
    }
}

// 消费者
func Consumer(in <-chan int) {
    for v := range in {
        fmt.Println(v)
    }
}
func main() {
    u := "/ad/forward?c=10001&header=100000&t=1"
    a,_ := url.Parse(u)
    m, _ := url.ParseQuery(a.RawQuery)
    fmt.Println(m["c"])
    fmt.Println(m["header"])
    fmt.Println(m["t"])
    /*ch := make(chan int, 64) // 成果队列

    go Producer(3, ch) // 生成 3 的倍数的序列
    go Producer(5, ch) // 生成 5 的倍数的序列
    go Consumer(ch)    // 消费 生成的队列

    // 运行一定时间后退出
    time.Sleep(5 * time.Second)*/
}
