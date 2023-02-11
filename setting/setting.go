package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init(filePath string) (err error) {
	viper.AddConfigPath(filePath)
	//viper.SetConfigFile("./conf/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		zap.L().Error("配置读取失败", zap.Error(err))
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Warn("配置文件修改了")
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper unmarshal faield", zap.Error(err))
		}
	})
	return
}

var Conf = new(AppConfig)
