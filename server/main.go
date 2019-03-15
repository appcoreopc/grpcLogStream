package main

// "github.com/appcoreopc/grpcLogStream/loggingStream/loggingStream"
import (
	"fmt"
	"net"
	"strconv"

	pb "github.com/appcoreopc/grpcLogStream/loggingstream"
	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

type LoggingService struct {
}

func (ls LoggingService) SendLog(request *pb.LogRequest, stream pb.LoggingStream_SendLogServer) error {

	fmt.Println("getting connection from client")

	for i := 0; i < 10; i++ {

		msg := &pb.LogResponse{Type: "MsgType" + strconv.Itoa(i)}
		stream.Send(msg)
	}

	return nil

}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Fail to listen")
	}

	// protobuf register server here //

	server := grpc.NewServer()

	pb.RegisterLoggingStreamServer(server, &LoggingService{})

	fmt.Println("serving up grpc services port", port)
	server.Serve(lis)
}
