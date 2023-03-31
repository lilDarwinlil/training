package main

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/lib/pq"
	"training/chans"
	"training/click"
	"training/mutexes"
)

func main() {
	fmt.Println("")
	mutexes.RunMutex()
	fmt.Println("")
	chans.RunChan()

	data := click.Read(123)
	fmt.Println(data)

	click.Write(123, "new message", 1233.343)
}
