package server_b

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/phev8/go_service_B/pkg/api"
	"google.golang.org/grpc"
)

type serviceBServer struct {
	counter int64
}

func NewServiceBServer() api.ServiceBServer {
	return &serviceBServer{
		counter: 0,
	}
}

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, port string,
) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// register service
	server := grpc.NewServer()
	api.RegisterServiceBServer(server, NewServiceBServer())

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	log.Println("wait connections on port " + port)
	return server.Serve(lis)
}
