package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/taaaaakahiro/grpc-example/domain/entity"
	"github.com/taaaaakahiro/grpc-example/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	//if crash the code, get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("hw")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// connect to database
	collection := getMongoDB()

	var user entity.User
	err := collection.FindOne(context.TODO(), bson.D{{"id", 1}}).Decode(&user)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return &pb.HelloReply{Message: "Hello again/" + strconv.Itoa(user.ID) + "/" + user.Name}, nil
}

func getMongoDB() *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("DATABASE")).Collection("user")
	return collection
}
