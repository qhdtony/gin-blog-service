package main

import (
	"github.com/gin-blog-service/internal/routers"
	"github.com/gin-blog-service/internal/model"
	"github.com/gin-blog-service/pkg/setting"
	"github.com/gin-blog-service/global"
	"github.com/gin-blog-service/pkg/logger"
	"net/http"
	"time"
	"log"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err =  setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEgine()
	if err != nil {
		log.Fatalf("init.setupDBEgine err: %v", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEgine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: fileName,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime:true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func main() {
	/*
	router := routers.NewRouter()
	s := &http.Server{
		Addr:			":8080",
		Handler:		router,
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
		MaxHeaderBytes:	1 << 20,
	}
	s.ListenAndServe()
	*/
	gin.SetMode(global.ServerSetting.RunMode)
	global.Logger.Infof("%s: book %s", "xx", "blog-service")

	router := routers.NewRouter()
	s := &http.Server{
		Addr:			":" + global.ServerSetting.HttpPort,
		Handler:		router,
		ReadTimeout:	global.ServerSetting.ReadTimeOut,
		WriteTimeout:	global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:	1 << 20,
	}
	s.ListenAndServe()
}
