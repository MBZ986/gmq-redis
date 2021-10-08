package main

import (
	"github.com/wuzhc/gmq-redis/mq"
)

func main() {
	q := mq.NewGmq("conf.ini")
	q.Run()
}
