package main

import (
	"log"

	"github.com/cycade/chaos-errors/internal/service"
)

func main() {
	appId := "42"
	env := "prod"
	err := service.Publish(appId, env)
	if err != nil {
		log.Fatalf("发布 %s 环境的应用 %s 失败:\n%+v\n", env, appId, err)
	}

	log.Printf("发布 %s 环境的应用 %s 成功", env, appId)
}
