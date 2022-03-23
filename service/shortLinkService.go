package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	linkDao "short-link/dao"
	"short-link/model"
)

func GetShortLink(url string) (string, error) {
	if len(url) == 0 {
		return "", errors.New("empty url")
	}
	uId := uuid.NewV4().String()
	link := model.Link{ShortLink: uId, LongLink: url}
	err := linkDao.Insert(link)
	if err != nil {
		return "", err
	}
	return uId, nil
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
