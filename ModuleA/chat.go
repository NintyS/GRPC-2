package chat

import (
	"log"

	"google.golang.org/grpc"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := save.NewChatServiceClient(conn)

	response, err := c.Save(context.Background(), &save.Message{Body: in.Body})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)

	return &Message{Body: "Aj Dont End Łyf Ju Jet!"}, nil
}
