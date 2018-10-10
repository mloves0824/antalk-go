package main

import (
	"fmt"
	"io"
	"log"
	//	"math/rand"
	//	"time"

	"github.com/docopt/docopt-go"
	pb "github.com/mloves0824/antalk-go/proto/apigw"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var name string
var token string
var subChan chan pb.TopicType

const (
	address = "localhost:50051"
)

// Login ...
func login(client pb.ApigwServiceClient) error {
	//log.Printf("Calling Register RPC")

	resp, err := client.Login(context.Background(), &pb.LoginReq{Uid: name, Token: token})
	if err != nil {
		log.Fatalf("Login failed %v", err)
	}
	log.Println("Login Resp: ", resp, err)

	return nil
}

// subscribeToModeChanges ...
func subscribeTopicKickout(c chan pb.TopicType) {
	log.Println("Subscribing to Kickout")
	c <- pb.TopicType_KICKOUT
}

// recvNotification ...
func recvNotification(stream pb.ApigwService_SubscribeClient) {
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to recv %v", err)
		}
		if err == io.EOF {
			break
		}
		log.Print(resp, err)
	}
}

// subscribe ...
func subscribe(client pb.ApigwServiceClient, c chan pb.TopicType) {
	stream, err := client.Subscribe(context.Background())
	if err != nil {
		log.Fatalf("failed to subscribe %v", err)
	}

	go recvNotification(stream)

	for t := range c {
		//log.Printf("sending subscibe for topic %s", t.String())
		if err := stream.Send(&pb.Topic{Uid: name, Type: t}); err != nil {
			log.Fatalf("couldnt send %s", t.String())
		}
	}
	stream.CloseSend()
}

func Argument(d map[string]interface{}, name string) (string, bool) {
	if d[name] != nil {
		if s, ok := d[name].(string); ok {
			if s != "" {
				return s, true
			}
			log.Panicf("option %s requires an argument", name)
		} else {
			log.Panicf("option %s isn't a valid string", name)
		}
	}
	return "", false
}

func main() {

	const usage = `
Usage:
	client --uid=uid --token=token
`
	d, err := docopt.Parse(usage, nil, true, "", false)
	if err != nil {
		log.Panicf("parse arguments failed: %s", err)
	}

	if s, ok := Argument(d, "--uid"); ok {
		fmt.Printf("uid=%s\n", s)
		name = s
	}
	if t, ok := Argument(d, "--token"); ok {
		fmt.Printf("token=%s\n", t)
		token = t
	}
	// init important structures
	subChan = make(chan pb.TopicType, 10)
	//	rand.Seed(time.Now().UTC().UnixNano())
	//	name = fmt.Sprintf("%s:%d", "Client", rand.Intn(50))
	//name = s

	// Setup a connection with the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewApigwServiceClient(conn)

	login(client)

	subscribeTopicKickout(subChan)
	subscribe(client, subChan)
}
