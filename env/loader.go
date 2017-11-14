package env

// 環境設定を保持するパッケージ

import (
	"bytes"
	"fmt"
	"goa-context-sample/assets"
	"goa-context-sample/valueobjects/config"
	"log"

	"github.com/deadcheat/goacors"
	"github.com/spf13/viper"
)

// 環境設定として保持する情報の構造体
var (
	// .goで定義
	OnDevelopment  = false
	configfilename string

	// tomlで設定
	Server   *config.ServerConfig
	CorsConf goacors.GoaCORSConfig
)

// init パッケージの初期化処理
func init() {
	// 初期設定
	Initialize()
	InitializeEnv()
}

// InitializeEnv パッケージの初期化処理
func InitializeEnv() {
	// 設定ファイル
	confDir := "config"
	fileName := configfilename
	ext := "toml"
	filePath := fmt.Sprintf("%s/%s.%s", confDir, fileName, ext)

	// Asset経由でファイルを取得しに行く
	viper.SetConfigType(ext)
	data, err := assets.Asset(filePath)
	if err != nil {
		log.Panic(err, filePath)
	}

	// go-bindataのAsset、およびconfigファイルのbindataが存在する場合
	viper.ReadConfig(bytes.NewBuffer(data))
	// 設定読み込み
	_ = viper.UnmarshalKey("server", &Server)
	_ = viper.UnmarshalKey("cors", &CorsConf)

}
