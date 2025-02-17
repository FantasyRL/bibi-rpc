package rpc

import (
	"bibi/cmd/user/vector/vector"
	"context"
	"google.golang.org/grpc"
)

func InitConvertRPC() {
	conn, err := grpc.Dial("localhost:8011", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	pictureClient = vector.NewVectorClient(conn)
	//convertClient = p.NewClipServiceClient(conn)
}

func UserGetVector(ctx context.Context, req *vector.GetVectorRequest) (resp *vector.GetVectorResponse, err error) {
	resp, err = pictureClient.GetPictureVector(ctx, req)
	//fmt.Println(resp.Vector)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
