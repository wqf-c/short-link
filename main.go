package main

import (
	"fmt"
	linkDao "short-link/dao"
	"short-link/model"
)

func main() {
	link := model.Link{LongLink: "xxx", ShortLink: "yyy"}
	linkDao.Insert(link)
	linkDao.SelectByLongLink(&link)
	fmt.Println(link)
}
