package service

import (
	"log"

	"github.com/cycade/chaos-errors/internal/biz"
)

func Publish(appId string, env string) error {
	allowMock := true
	if env == "prod" {
		allowMock = false
	}

	artifactPath, err := biz.GetArtifactPath(appId, allowMock)
	if err != nil {
		return err
	}

	// 所谓 publish 的本体是一次 log
	log.Printf("PUBLISH %s =============> \n", artifactPath)
	return nil
}
