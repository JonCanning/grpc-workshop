package main

import (
	"context"
	"fmt"
	"main/proto"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
)

type greeterService struct {
	proto.UnimplementedGreeterServer
}

func (s *greeterService) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Printf("Received request from %s\n", request.Name)
	message := fmt.Sprintf("Hello %s from GO", request.Name)
	return &proto.HelloReply{Message: message}, nil
}

func startServer() {
	go func() {
		listener, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			panic(err)
		}
		server := grpc.NewServer()
		greeterService := greeterService{}
		proto.RegisterGreeterServer(server, &greeterService)
		println("Server started")
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()
}

func startClient() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithConnectParams(grpc.ConnectParams{Backoff: backoff.DefaultConfig}))
	if err != nil {
		panic(err)
	}
	client := proto.NewGreeterClient(conn)
	ctx := context.Background()
	request := proto.HelloRequest{Name: "GO"}
	response, err := client.SayHello(ctx, &request)
	if err != nil {
		panic(err)
	}
	println(response.Message)
}

func main() {
	fmt.Println("Press 1 to start server, 2 to start client")
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1)
	_, err = os.Stdin.Read(buf)
	if err != nil {
		panic(err)
	}
	term.Restore(int(os.Stdin.Fd()), oldState)
	arg, err := strconv.Atoi(string(buf))
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	switch arg {
	case 1:
		startServer()
	case 2:
		startClient()
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
		os.Exit(1)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
