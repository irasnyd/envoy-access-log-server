package sink

import (
	"io"
	"log"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/metrics/v3"
	"github.com/gogo/protobuf/jsonpb"
)

type server struct {
	marshaler jsonpb.Marshaler
}

var _ v3.MetricsServiceServer = &server{}

// New ...
func New() v3.MetricsServiceServer {
	return &server{}
}

func (s *server) StreamMetrics(stream v3.MetricsService_StreamMetricsServer) error {
	//log.Println("Started stream")
	for {
		in, err := stream.Recv()
		//log.Println("Received value")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		str, _ := s.marshaler.MarshalToString(in)
		log.Println(str)
	}
}
