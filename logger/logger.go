package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() (err error) {
	//通过viper获取logger的配置
	writeSyncer := getLogWriter(
		viper.GetString("log.filename"),
		viper.GetInt("log.maxsize"),
		viper.GetInt("log.maxbackup"),
		viper.GetInt("log.maxage"))
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(viper.GetString("log.level")))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	lg := zap.New(core, zap.AddCaller())
	//替换zap库中的全局logger
	zap.ReplaceGlobals(lg)
	return
}
