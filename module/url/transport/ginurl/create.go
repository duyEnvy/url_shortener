package ginurl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"url-shortener/common"
	"url-shortener/component/appctx"
	urlbiz "url-shortener/module/url/biz"
	urlmodel "url-shortener/module/url/model"
	urlstorage "url-shortener/module/url/storage"
)

func CreateUrl(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data urlmodel.UrlCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := urlstorage.NewSQLStore(db)
		biz := urlbiz.NewCreateUrlBiz(store)
		result, err := biz.CreateUrl(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(os.Getenv("LOCAL_PORT")+result.ShortCode))
	}
}
