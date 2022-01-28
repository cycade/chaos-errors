package biz

import (
	"database/sql"
	"errors"

	"github.com/cycade/chaos-errors/internal/data"
	pkgErrors "github.com/pkg/errors"
)

const MOCK_ARTIFACT = "path/to/python3"

// 寻找 AppId 所对应的产物路径，在允许使用兜底的情况下使用 MOCK_ARTIFACT
func GetArtifactPath(appId string, allowMock bool) (string, error) {
	record, err := data.FindPublishRecordByAppId(appId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if allowMock {
				return MOCK_ARTIFACT, nil
			} else {
				return "", pkgErrors.Wrap(err, "未能找到 AppId 现存的产物路径")
			}
		}

		return "", pkgErrors.Wrap(err, "未能找到 AppId 现存的产物路径")
	}

	return record.ArtifactPath, nil
}
