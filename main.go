package main

import (
	"fmt"
	"os"
	"simple-demo/controller"
	"simple-demo/dao"
	"simple-demo/models"
	"simple-demo/routers"
	"simple-demo/service"
	"simple-demo/setting"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./simple-demo conf/config.ini")
		return
	}

	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	go service.RunMessageServer()

	//r := gin.Default()
	// 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接

	//s := grpc.NewServer()
	//publish.RegisterPublishServiceServer(s)
	// 模型绑定
	dao.DB.AutoMigrate(&(controller.User{}), &(controller.Video{}), &(models.VideoInfo{}))
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
