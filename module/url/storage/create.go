package urlstorage

import (
	"context"
	urlmodel "url-shortener/module/url/model"
)

func (s *sqlStore) Create(ctx context.Context, data *urlmodel.UrlCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
