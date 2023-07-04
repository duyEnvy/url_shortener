package urlstorage

import (
	"context"
	urlmodel "url-shortener/module/url/model"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*urlmodel.Url, error) {
	var data urlmodel.Url

	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
