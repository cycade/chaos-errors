package data

import "database/sql"

type PublishRecord struct {
	Id           string
	AppId        string
	ArtifactPath string
}

func FindPublishRecordByAppId(id string) (PublishRecord, error) {
	var record PublishRecord
	// err := db.QueryRow("select id, app_id, artifact_path from publish_record where app_id = ?", id).Scan(&record)

	// 在此处并没有办法处理这个错误，甚至都不知道 ErrNoRows 对于业务逻辑来说是不是一个错误
	// 由业务层在发现真正业务逻辑错误时，进行 wrap 获取堆栈信息，在此处 wrap 会导致堆栈信息冗余
	return record, sql.ErrNoRows
}
