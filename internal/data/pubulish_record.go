package data

import (
	"database/sql"

	pkgErrors "github.com/pkg/errors"
)

type PublishRecord struct {
	Id           string
	AppId        string
	ArtifactPath string
}

func FindPublishRecordByAppId(appId string) (PublishRecord, error) {
	var record PublishRecord
	// err := db.QueryRow("select id, app_id, artifact_path from publish_record where app_id = ?", id).Scan(&record)

	// 在此处并没有办法处理这个错误，甚至都不知道 ErrNoRows 对于业务逻辑来说是不是一个错误
	err := ChaosError{err: sql.ErrNoRows}
	return record, pkgErrors.Wrap(err, "未能找到 AppId 的产物路径")
}

type SQLError interface {
	IsDataNotFound() bool
}

type ChaosError struct {
	err error
}

func (c ChaosError) Error() string {
	return c.err.Error()
}

func (c ChaosError) IsDataNotFound() bool {
	return pkgErrors.Is(c.err, sql.ErrNoRows)
}
