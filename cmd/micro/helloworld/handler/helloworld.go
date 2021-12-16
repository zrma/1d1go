package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	helloworld "helloworld/proto"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	logger.Info("Received Helloworld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Helloworld) Stream(ctx context.Context, req *helloworld.StreamingRequest, stream helloworld.Helloworld_StreamStream) error {
	logger.Infof("Received Helloworld.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Responding: %d", i)
		if err := stream.Send(&helloworld.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Helloworld) PingPong(ctx context.Context, stream helloworld.Helloworld_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&helloworld.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
