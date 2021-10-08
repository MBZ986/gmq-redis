package main

import (
	"github.com/jackiedong168/gmq-redis/mq"
)

func main() {
	q := mq.NewGmq("conf.ini")
	q.Run()
}
