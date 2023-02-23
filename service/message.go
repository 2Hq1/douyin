package service

import (
	"encoding/json"
	"fmt"
	"github.com/RaymondCode/simple-demo/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"net"
	"simple-demo/models/pb"
	"sync"
)

type feedServer struct {
	pb.UnimplementedVideoServiceServer
}

var chatConnMap = sync.Map{}

func RunMessageServer() {
	creds, _ := credentials.NewServerTLSFromFile("D:\\goLandProjects\\byte\\simple-demo\\key\\test.pem", "D:\\goLandProjects\\byte\\simple-demo\\key\\test.key")

	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}
	// 创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	// 在grpc服务端中去注册自己编写的服务
	pb.RegisterVideoServiceServer(grpcServer, &feedServer{})

	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept conn failed: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue
		}

		var event = controller.MessageSendEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		fmt.Printf("Receive Message：%+v\n", event)

		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			fmt.Printf("User %d offline\n", event.ToUserId)
			continue
		}

		pushEvent := controller.MessagePushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent)
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			fmt.Printf("Push message failed: %v\n", err)
		}
	}
}
