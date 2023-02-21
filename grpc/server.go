package grpc

import (
	"context"
	"github.com/go-xmodule/utils/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
)

const (
	TCP = "tcp"
	UDP = "udp"
)

var GrpcServer = new(Server)

type ServerRegister func(server *grpc.Server)

type Server struct {
	network  string
	address  string
	token    string
	register ServerRegister
}

func NewServer() *Server {
	return new(Server)
}

// Init 初始化
func (g *Server) Init(network, address, token string) *Server {
	g.network = network
	g.address = address
	g.token = token
	return g
}
func (g *Server) RegisterServer(register ServerRegister) *Server {
	g.register = register
	return g
}

// Start 启动grpc 服务
func (g *Server) Start() error {
	var opts []grpc.ServerOption
	// 注册interceptor
	opts = append(opts, grpc.UnaryInterceptor(g.interceptor))
	listen, err := net.Listen(g.network, g.address)
	if err != nil {
		return err
	}
	// 实例化grpc Server
	s := grpc.NewServer(opts...)
	// 注册服务
	g.register(s)
	return s.Serve(listen)
}

// interceptor 拦截器
func (g *Server) interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	err := g.auth(ctx)
	if err != nil {
		return nil, err
	}
	// 继续处理请求
	return handler(ctx, req)
}

// 身份认证
func (g *Server) auth(context context.Context) error {
	md, ok := metadata.FromIncomingContext(context)
	if !ok {
		return status.Errorf(codes.Unauthenticated, global.NoTokenErr.String())
	}
	if _, ok := md["ts"]; !ok {
		return status.Error(codes.Unauthenticated, global.NoSignParamsErr.String())
	}
	newSign := RequestSign(md["ts"][0], g.token)
	if sign, ok := md["sign"]; !ok || sign[0] != newSign {
		return status.Error(codes.Unauthenticated, global.TokenErr.String())
	}
	return nil
}
