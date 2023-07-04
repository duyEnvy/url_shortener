package urlbiz

import (
	"context"
	urlmodel "url-shortener/module/url/model"
)

type FindUrlStore interface {
	FindDataWithCondition(
		context context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*urlmodel.Url, error)
}

type findUrlBiz struct {
	store FindUrlStore
}

func NewFindUrlBiz(store FindUrlStore) *findUrlBiz {
	return &findUrlBiz{store: store}
}

func (biz *findUrlBiz) FindUrl(
	context context.Context,
	code string,
) (*urlmodel.Url, error) {
	result, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"short_code": code})

	if err != nil {
		return nil, err
	}

	return result, nil
}
