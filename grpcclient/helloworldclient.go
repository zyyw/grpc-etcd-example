package main

import (
	"flag"
	"fmt"
	"github.com/zyyw/grpc-etcd-example/expample/pb2"
	"time"

	grpclb "github.com/zyyw/grpc-etcd-example/balance"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	serv = flag.String("service", "user_service", "service name")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()
	fmt.Println("serv", *serv)
	r := grpclb.NewResolver(*serv)
	b := grpc.RoundRobin(r)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	fmt.Println("conn...")

	//ticker := time.NewTicker(1 * time.Second)
	//for range ticker.C {
		client := pb2.NewUserServiceClient(conn)
		resp, err := client.SetUserAutopay(context.Background(), &pb2.AutopayRequest{
			UserAddress: "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF",
			Code:        pb2.AutopayRequest_YES,
		})
		if err == nil {
			fmt.Printf("Reply is %s\n", resp.Message)
		} else {
			fmt.Println("AAAA")
			fmt.Println(err)
		}
	//}

	//ticker := time.NewTicker(1 * time.Second)
	//for t := range ticker.C {
	//	client := pb.NewGreeterClient(conn)
	//	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
	//	if err == nil {
	//		fmt.Printf("%v: Reply is %s\n", t, resp.Message)
	//	} else {
	//		fmt.Println("AAAA")
	//		fmt.Println(err)
	//	}
	//}
}
