package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// 创建一个无缓冲通道 - 消息创建者 - 生产者
    msg_chan := make(chan string)
	// 消费者
    done := make(chan bool)

    i := 0

    go func() {
        for  {
            i++
            time.Sleep(1*time.Second)
			// 生产者 - 生成一则消息
            msg_chan <- "on message"
			// 阻塞当前循环 - 等待消费
            <- done
        }
    }()

    go func() {
        for {
            select {
				// 监听生产者消息
            case msg := <- msg_chan :
				// 当生产者有消息产出是立即消费
                i++
				// 这里strconv.Itoa()将数字转为字符串类型
				// 但是string()会把整型数值转换为ascii码等于该数字的字符
                fmt.Println(msg + " " + strconv.Itoa(i))
                time.Sleep(2*time.Second)
				// 消费完以后给消息生成者让那个生产者继续执行
                done <- true
            }
        }

    }()


    time.Sleep(20*time.Second)
}