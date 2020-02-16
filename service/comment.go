package service

import (
	"github.com/treant5612/ytvc-web/db/sqldb"
	"github.com/treant5612/ytvc-web/model"
	"strconv"
)

func InsertComment(nick string, content string, addr string) (err error) {
	comment := &model.Comment{
		Content:  content,
		NickName: nick,
		Addr:     addr,
	}
	return sqldb.InsertComment(comment)
}

func ListComment(pageNo, pageSize string) (comments []*model.Comment, err error) {
	pNo, err := strconv.Atoi(pageNo)
	if err != nil {
		return nil, err
	}
	pSize, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, err
	}
	return sqldb.ListComment(pNo, pSize)
}
