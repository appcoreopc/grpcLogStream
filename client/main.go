package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/appcoreopc/grpcLogStream/loggingstream"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:9001"
	defaultName = "world"
)

func main() {

	connecton, err := grpc.Dial(address, grpc.WithInsecure())

	defer connecton.Close()

	if err != nil {
		fmt.Println("error connecting")
	}

	client := pb.NewLoggingStreamClient(connecton)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, re := client.SendLog(ctx, &pb.LogRequest{Logfile: "test.dat"})

	if re != nil {
		fmt.Println("error getting response from server ")
	}

	for {

		f, e := r.Recv()

		if e == io.EOF {
			break
		}

		fmt.Println(f.Type)
	}
	//fmt.Println(r.Type)

}
