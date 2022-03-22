package linkDao

import (
	"short-link/model"
	"short-link/utils"
)

func Insert(link model.Link) error {
	result := utils.Db.Create(&link)
	return result.Error
}

func SelectByShortLink(link *model.Link) error {
	result := utils.Db.Where("shortLink = ?", link.ShortLink).First(link)
	return result.Error
}

func SelectByLongLink(link *model.Link) error {
	result := utils.Db.Where("longLink = ?", link.LongLink).First(link)
	return result.Error
}
