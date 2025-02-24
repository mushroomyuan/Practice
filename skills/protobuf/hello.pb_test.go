package protobuf_test

import (
	"testing"

	"github.com/mushroomyuan/Practice/skills/protobuf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestHelloRequest(t *testing.T) {
	req := &protobuf.HelloRequest{
		MyName: "YFZ",
		MyAge:  22,
	}
	data, err := proto.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Marshaled data: %s", string(data))

	req2 := &protobuf.HelloRequest{}
	err = proto.Unmarshal(data, req2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Unmarshaled request: %+v", req2)
}

func TestAny(t *testing.T) {
	// 1.Pack
	req := &protobuf.Event{
		Type:    protobuf.EVENT_TYPE_ECS,
		Message: "test for any",
		Detail:  []*anypb.Any{},
	}
	ecsAny, err := anypb.New(&protobuf.EVENT_ECS{Message: "1"})
	if err != nil {
		t.Fatal(err)
	}
	req.Detail = append(req.Detail, ecsAny)
	t.Logf("ECS: %+v", req)
	// 2.Unpack
	t.Log(req.Detail[0])
	eceEvent := &protobuf.EVENT_ECS{}
	if err := req.Detail[0].UnmarshalTo(eceEvent); err != nil {
		t.Fatal(err)
	}
	t.Log(eceEvent)
}
