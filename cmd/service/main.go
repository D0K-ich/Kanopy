package main

import (
	"flag"
	"context"

	"github.com/D0K-ich/Kanopy"
	"github.com/D0K-ich/Kanopy/logs"
)

var (
	mainCtx, mainCancel = context.WithCancel(context.Background())
	configPath = flag.String("config", "..\\..\\templates\\kanopy.yml", "Config file path")
)

func init() {
	var err error

	var config *Kanopy.Config
	if config, err 			= Kanopy.NewConfig(configPath)	; err != nil {panic(err)}
	if logs.Log, err 	= logs.NewLog(config.Logger)	; err != nil {panic(err)}
}

func main() {
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
	logs.Log.Debug("test", "key", map[string]interface{}{"123": "asd"})
}