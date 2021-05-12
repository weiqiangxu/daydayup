package main

import (
	"context"
	"fmt"
)

func main() {
    ctx := context.Background()
    process(ctx)
	// 创建一个上下文对象携带键值对的ctx对象
    ctx = context.WithValue(ctx, "name", "jack")
    process(ctx)
}

func process(ctx context.Context) {
	// 获取上下文对象的键值对的一个“键”对应的值
    name, ok := ctx.Value("name").(string)
    if ok {
        fmt.Printf("process over. trace_id=%s\n", name)
    } else {
        fmt.Printf("process over. no trace_id\n")
    }
}