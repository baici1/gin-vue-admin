package core

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rabbitmq"
	"github.com/gin-contrib/pprof"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	//注册参数校验翻译器测试
	if err := utils.InitTrans("zh"); err != nil {
		global.GVA_LOG.Error("Validator Transition Tool Register Failed", zap.Error(err))
	}
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	//当当前模式为测试模式，开始pprof
	if global.GVA_CONFIG.System.Env == "test" {
		pprof.Register(Router)
	}
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	GVA讨论社区:https://support.qq.com/products/371961
	`)
	//启动rabbitmq服务处理消息
	go rabbitmq.CreateRabbitMqFactoryByRouting().StartHandleMsgByRoutingToStudent()
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
