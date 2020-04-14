// Package main implements a client for Greeter service.
package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	commpb "github.com/lpy-neo/micro-framework-test/comm_proto"
	pb "github.com/lpy-neo/micro-framework-test/grpc-gateway/proto"
	hellopb "github.com/lpy-neo/micro-framework-test/helloworld/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:9876"
	defaultName = "world"
)

func main() {
	grpcReq()
	restReq()
}

func grpcReq() {
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	// Set up a connection to the server.
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var req hellopb.HelloRequest
	req.Name = "lipengyong"
	bytes, _ := proto.Marshal(&req)

	r, err := c.GrpcReq(ctx, &commpb.GrpcRequest{Head: &commpb.GrpcRequestHead{Cmd: 1000, Uid: "uid123"}, Body: bytes})
	if err != nil {
		log.Fatalf("could not send grpc req: %v", err)
	}

	var rsp hellopb.HelloReply
	proto.Unmarshal(r.Data, &rsp)
	log.Printf("Greeting: %v", rsp.Message)
}

func restReq() {
	var req hellopb.HelloRequest
	req.Name = "lipengyong"
	bytes1, _ := json.Marshal(&req)

	commReq := &commpb.GrpcRequest{Head: &commpb.GrpcRequestHead{Cmd: 1000, Uid: "uid123", Encoding: 2}, Body: bytes1}
	commbytes, _ := json.Marshal(commReq)

	httpReq, err := http.NewRequest("POST", "http://localhost:9877/grpc_gateway.RestService", bytes.NewBuffer(commbytes))
	// req.Header.Set("X-Custom-Header", "myvalue")
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var rsp commpb.GrpcReply
	json.Unmarshal(body, &rsp)
	var helloRsp hellopb.HelloReply
	json.Unmarshal(rsp.Data, &helloRsp)
	log.Printf("Greeting: %v", helloRsp.Message)
}