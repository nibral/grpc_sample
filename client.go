package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(flag.Arg(0), grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := NewHelloClient(conn)
	msg := &HelloRequest{}
	res, err := client.SayHello(context.TODO(), msg)
	if err != nil {
		fmt.Printf("error::%#v \n", err)
	}
	fmt.Println(res.Msg)
}
