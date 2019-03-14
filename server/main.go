package main

// "github.com/appcoreopc/grpcLogStream/loggingStream/loggingStream"
import (
	"context"
	"fmt"
	"net"

	pb "github.com/appcoreopc/grpcLogStream/loggingstream"
	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

type LoggingService struct {
}

func (ls LoggingService) SendLog(context.Context, *pb.LogRequest) (*pb.LogResponse, error) {
	return nil, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Fail to listen")
	}

	// protobuf register server here //

	server := grpc.NewServer()

	pb.RegisterLoggingStreamServer(server, &LoggingService{})

	server.Serve(lis)
}
