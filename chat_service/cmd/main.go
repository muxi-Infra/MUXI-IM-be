package main

import (
	bp "github.com/muxi-Infra/MUXI-IM-be/chat_service/grpc"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/ioc"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/repository"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/repository/dao"
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/service"
	"github.com/muxi-Infra/MUXI-IM-be/pkg/logger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	runMode := viper.GetString("run.mode")
	switch runMode {
	case "dev":
		// 本地测试
		StartChatServiceLocal()
	case "prod":
		// 生产环境
		StartChatService()
	default:
		panic("run.mode must be dev or prod")
	}
	
}

func StartChatServiceLocal() {
l := logger.NewNopLogger()
	db := ioc.InitDBwithSqlite()
	d := dao.NewDao(db)
	repo := repository.NewChatRepository(d)
	ser := service.NewChatService(repo)
	chatService := bp.NewChatRpcServer(ser)
	rpc := ioc.InitGRPCxKratosServer(chatService, nil, l)

	err := rpc.Serve()
	if err != nil {
		panic(err)
	}
}

func StartChatService() {
   l := logger.NewNopLogger()
	db := ioc.InitDB(l)
	d := dao.NewDao(db)
	repo := repository.NewChatRepository(d)
	ser := service.NewChatService(repo)
	client := ioc.InitEtcdClient()
	chatService := bp.NewChatRpcServer(ser)
	rpc := ioc.InitGRPCxKratosServer(chatService, client, l)

	err := rpc.Serve()
	if err != nil {
		panic(err)
	}
}

func init() {
	initViper()
}

func initViper() {
	cfile := pflag.String("config", "./chat_service/config.yaml", "配置文件路径")
	pflag.Parse()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(*cfile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
