package main

import (
	"RSs/common"
	"RSs/config"
	"RSs/plugins/etcd"
	"RSs/plugins/logs"
	pb "RSs/rpc"
	"context"
	"flag"
	"fmt"
	"github.com/micro/go-web"
	"google.golang.org/grpc"
	"os"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		http()
	}()
	go func() {
		rpc()
	}()
	wg.Wait()

}

func http(){
	port := flag.String("httpport", "7777", "http listen port")
	flag.Parse()

	service := web.NewService(
		web.Address("0.0.0.0:" + *port),
	)

	service.Init()
	service.Handle("/", config.GetRouterContainer())

	logs.Info("http has listen port on 7777")
	etcdAddr := fmt.Sprintf("%s:%d", common.EtcdHost, common.EccdPort)
	etcd.EtcdConfSet(etcdAddr)
	if err := service.Run(); err != nil {
		logs.Error(err)
	}
}
func rpc(){
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		logs.Error("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "defaultName"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		logs.Error("could not greet: %v", err)
	}
	logs.Info("Greeting: ", r.Message)
}
