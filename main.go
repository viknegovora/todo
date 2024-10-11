package main

 import (

 "github.com/viknegovora/todo-app"
 "github.com/viknegovora/todo-app/pkg/handler"
 "github.com/viknegovora/todo-app/pkg/repository"
 "github.com/viknegovora/todo-app/pkg/service"
 "log"

 )

func main() {
    if err := initConfig(); err != nil {
        log.Fatalf("error initializing configs: %s", err.Error())
    }
    repos := repository.NewRepository()
    services := service.NewService(repos)
    handlers := handler.NewHandler(services)

    srv := new(todo.Server)
    if err := srv.Run(viper.getString("8000"), handlers.initRoutes(); err != nil {
            log.Fatalf("error occurred while running http server: %s", err.Error())
    }
}

func initConfig() error {
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}