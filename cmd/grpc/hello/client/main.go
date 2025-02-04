package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"1d1go/cmd/grpc/hello/pb"
)

const (
	address     = "localhost:12345"
	defaultName = "world"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			// keepalive settings - https://github.com/grpc/grpc/blob/master/doc/keepalive.md
			Time:                10 * time.Second,
			Timeout:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		return fmt.Errorf("grpc connection failed: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return fmt.Errorf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return fmt.Errorf("could not greet again: %v", err)
	}
	log.Printf("Greeting again: %s", r.Message)

	var goroErr error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(60 * time.Second)
			r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
			if err != nil {
				goroErr = fmt.Errorf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", r.GetMessage())
		}
	}()

	prev := conn.GetState()
	for i := 0; i < 1000000000; i++ {
		time.Sleep(time.Nanosecond)
		if curr := conn.GetState(); curr != prev {
			log.Println(curr)
			prev = curr
		}
		if i%10000 == 0 {
			fmt.Printf(".")
		}
	}

	log.Println("end")
	time.Sleep(time.Second)

	wg.Wait()
	return goroErr
}
