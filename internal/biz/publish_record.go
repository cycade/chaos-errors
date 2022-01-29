package biz

import (
	"errors"

	"github.com/cycade/chaos-errors/internal/data"
)

const MOCK_ARTIFACT = "path/to/python3"

// 寻找 AppId 所对应的产物路径，在 允许使用兜底(allowMock) 的情况下返回 MOCK_ARTIFACT
func GetArtifactPath(appId string, allowMock bool) (string, error) {
	record, err := data.FindPublishRecordByAppId(appId)
	if err != nil {
		var sqlError data.SQLError
		if errors.As(err, &sqlError) && sqlError.IsDataNotFound() && allowMock {
			return MOCK_ARTIFACT, nil
		}

		return "", err
	}

	return record.ArtifactPath, nil
}
