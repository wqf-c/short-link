package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	linkDao "short-link/dao"
	"short-link/model"
)

func GetShortLink(url string) (string, error) {
	if len(url) == 0 {
		return "", errors.New("empty url")
	}
	link := model.Link{LongLink: url}
	err := linkDao.SelectByLongLink(&link)
	if err == nil {
		return link.ShortLink, nil
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uId := uuid.NewV4().String()
			link = model.Link{ShortLink: uId, LongLink: url}
			err = linkDao.Insert(link)
			if err != nil {
				return "", err
			}
			return uId, nil
		} else {
			return "", nil
		}
	}

}

func GetLongLinkBySLink(shortLink string) (*model.Link, error) {
	if len(shortLink) == 0 {
		return nil, errors.New("empty shortLink")
	}
	link := model.Link{ShortLink: shortLink}
	err := linkDao.SelectByShortLink(&link)
	if err != nil {
		return nil, err
	} else {
		return &link, nil
	}
}
