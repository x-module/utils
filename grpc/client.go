/**
 * Created by PhpStorm.
 * @file   client.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/27 22:51
 * @desc   client.go
 */

package grpc

import (
	"context"
	"github.com/go-utils-module/utils/global"
	"github.com/go-utils-module/utils/utils/xlog"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

var GrpcClient = new(Client)

type Client struct {
	address string
	token   string
	connect *grpc.ClientConn
}

func NewClient() *Client {
	return new(Client)
}

// Init 初始化
func (g *Client) Init(address string, token string) *Client {
	g.address = address
	g.token = token
	return g
}

func (g *Client) Close() error {
	return g.connect.Close()
}

func (g *Client) Connect() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(NewCustomCredential(g.token)))
	// 指定客户端interceptor
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(g.address, opts...)
	if err != nil {
		xlog.Logger.WithField("err", err).Errorf(global.RPCLinkErr.String())
		return nil, err
	}
	g.connect = conn
	return conn, err
}

// CustomCredential 自定义认证
type CustomCredential struct {
	sign string
	ts   int64
}

func NewCustomCredential(token string) *CustomCredential {
	credential := &CustomCredential{}
	credential.sign, credential.ts = RpcSign(token)
	return credential
}

// GetRequestMetadata 实现自定义认证接口
func (c CustomCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"sign": c.sign,
		"ts":   strconv.FormatInt(c.ts, 10),
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c CustomCredential) RequireTransportSecurity() bool {
	return false
}

// interceptor 客户端拦截器
func interceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if xlog.Logger == nil {
		log.Printf("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	} else {
		xlog.Logger.Debugf("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	}
	return err
}
