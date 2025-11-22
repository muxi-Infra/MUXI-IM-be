package ioc

import (
	"time"

	"github.com/muxi-Infra/MUXI-IM-be/pkg/grpcx"

	"github.com/muxi-Infra/MUXI-IM-be/chat_service/grpc"
	"github.com/muxi-Infra/MUXI-IM-be/pkg/logger"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
)

func InitGRPCxKratosServer(ser *grpc.ChatRpcServer, ecli *clientv3.Client, l logger.Logger) grpcx.Server {
	type Config struct {
		Name    string `yaml:"name"`
		Weight  int    `yaml:"weight"`
		Addr    string `yaml:"addr"`
		EtcdTTL int64  `yaml:"etcdTTL"`
	}
	var cfg Config
	err := viper.UnmarshalKey("grpc.server", &cfg)
	if err != nil {
		panic(err)
	}
	server := kgrpc.NewServer(
		kgrpc.Address(cfg.Addr),
		kgrpc.Middleware(recovery.Recovery()),
	)

	if ecli == nil {
		return newLocalServer(server, cfg.Name)
	}

	ser.Register(server)

	return &grpcx.KratosServer{
		Server:     server,
		Name:       cfg.Name,
		Weight:     cfg.Weight,
		EtcdTTL:    time.Second * time.Duration(cfg.EtcdTTL),
		EtcdClient: ecli,
		L:          l,
	}
}


type localServer struct {
    server *kgrpc.Server
    stop   func() error
    name   string
}

func newLocalServer(s *kgrpc.Server, name string) grpcx.Server {
    return &localServer{server: s, name: name}
}

func (s *localServer) Serve() error {
    app := kratos.New(kratos.Name(s.name), kratos.Server(s.server))
    s.stop = app.Stop
    return app.Run()
}

func (s *localServer) Close() error {
    if s.stop != nil {
        return s.stop()
    }
    return nil
}