package main

import "github.com/kuaileniu/zlog"

func init() {
	zlog.InitLogger(zlog.LogConfig{
		Filename: "./logs/transferfiles_server.log",
		Level:    "debug",
		// Level:      "info",
		MaxSize:    5,
		MaxBackups: 10,
		MaxAge:     10,
		Console:    true,
	})
}

func main(){
	
}