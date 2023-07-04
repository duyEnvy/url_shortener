package ginurl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/component/appctx"
	urlbiz "url-shortener/module/url/biz"
	urlmodel "url-shortener/module/url/model"
	urlstorage "url-shortener/module/url/storage"
)

func FindUrl(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		code := c.Param("short_code")

		var result *urlmodel.Url

		store := urlstorage.NewSQLStore(db)
		biz := urlbiz.NewFindUrlBiz(store)

		result, err := biz.FindUrl(c.Request.Context(), code)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Redirect(http.StatusMovedPermanently, result.OriginalUrl)
	}
}
