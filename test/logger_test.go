package test

import (
	"github.com/jishulangcom/go-def"
	"github.com/jishulangcom/go-logger"
	"testing"
)

func Test(t *testing.T) {
	// 默认文件名为"./app.log"
	logger.Error("技术狼|jishulang.com")

	// 可自定义日志文件名
	logger.New("app2.log")

	// 带上链路日志ID
	logger.ErrorfTrace(def.Ctx, "技术狼|jishulang.com")
}
