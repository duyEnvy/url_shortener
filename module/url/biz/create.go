package urlbiz

import (
	"context"
	"strings"
	"url-shortener/common"
	urlmodel "url-shortener/module/url/model"
)

type CreateUrlStore interface {
	Create(ctx context.Context, data *urlmodel.UrlCreate) error
	FindDataWithCondition(
		context context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*urlmodel.Url, error)
}

type createUrlBiz struct {
	store CreateUrlStore
}

func NewCreateUrlBiz(store CreateUrlStore) *createUrlBiz {
	return &createUrlBiz{store: store}
}

func (biz *createUrlBiz) CreateUrl(ctx context.Context, data *urlmodel.UrlCreate) (*urlmodel.Url, error) {
	existedUrl, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"original_url": data.OriginalUrl})

	if err != nil {
		if err.Error() != "record not found" {
			return nil, err
		}
	}

	if existedUrl != nil {
		return existedUrl, nil
	}

	for {
		shortCode, err := common.GenerateShortCode()
		if err != nil {
			return nil, err
		}

		data.ShortCode = shortCode
		if err := biz.store.Create(ctx, data); err != nil {
			if strings.Contains(err.Error(), "Error 1062 (23000): Duplicate entry") {
				continue
			}
			return nil, err
		}
		break
	}

	return &urlmodel.Url{OriginalUrl: data.OriginalUrl, ShortCode: data.ShortCode}, nil
}
