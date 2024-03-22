// Code generated by goctl. DO NOT EDIT!
// Source: publish.proto

package publishrpc

import (
	"context"

	"DouYin/server/publish/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PublishActionReq  = __.PublishActionReq
	PublishActionResp = __.PublishActionResp
	PublishListReq    = __.PublishListReq
	PublishListResp   = __.PublishListResp
	Video             = __.Video

	PublishRpc interface {
		// 视频投稿
		PublishAction(ctx context.Context, in *PublishActionReq, opts ...grpc.CallOption) (*PublishActionResp, error)
		// 发布列表
		PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error)
	}

	defaultPublishRpc struct {
		cli zrpc.Client
	}
)

func NewPublishRpc(cli zrpc.Client) PublishRpc {
	return &defaultPublishRpc{
		cli: cli,
	}
}

// 视频投稿
func (m *defaultPublishRpc) PublishAction(ctx context.Context, in *PublishActionReq, opts ...grpc.CallOption) (*PublishActionResp, error) {
	client := __.NewPublishRpcClient(m.cli.Conn())
	return client.PublishAction(ctx, in, opts...)
}

// 发布列表
func (m *defaultPublishRpc) PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error) {
	client := __.NewPublishRpcClient(m.cli.Conn())
	return client.PublishList(ctx, in, opts...)
}
