package urlmodel

type Url struct {
	Id          int    `json:"id" gorm:"column:id;"`
	OriginalUrl string `json:"original_url" gorm:"column:original_url;"`
	ShortCode   string `json:"short_code" gorm:"column:short_code;"`
}

func (Url) TableName() string { return "urls" }

type UrlCreate struct {
	OriginalUrl string `json:"original_url" gorm:"column:original_url;"`
	ShortCode   string `json:"short_code" gorm:"column:short_code;"`
}

func (UrlCreate) TableName() string { return Url{}.TableName() }
