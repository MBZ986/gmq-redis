package mq

import (
	"context"
	"fmt"
	"github.com/bitly/go-simplejson"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Service struct {
	simplejson.Json
}

func (s *Service) Push(job *Job, reply *string) error {
	//delay, _ := strconv.Atoi(j["delay"])
	//TTR, _ := strconv.Atoi(j["TTR"])
	//job := &Job{
	//	Id:    j["id"],
	//	Body:  j["body"],
	//	Topic: j["topic"],
	//	Delay: delay,
	//	TTR:   TTR,
	//}

	err := gmq.dispatcher.AddToJobPool(job)
	if err != nil {
		*reply = err.Error()
	} else {
		*reply = "success"
	}
	return nil
}

func (s *Service) Pop(topic []string, reply *map[string]string) (err error) {
	*reply, err = Pop(topic...)
	return err
}

func (s *Service) Ack(id string, reply *bool) (err error) {
	*reply, err = Ack(id)
	return err
}

type RpcServer struct {
}

func (s *RpcServer) Run(ctx context.Context) {
	rpc.Register(new(Service))
	port := gmq.cfg.Section("server").Key("rpc_port").MustString("9503")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error("listen error:", err)
	} else {
		defer listener.Close()
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		select {
		case <-ctx.Done():
			log.Info("rpcServer exit")
			return
		default:
		}
		go jsonrpc.ServeConn(conn)
	}
}
