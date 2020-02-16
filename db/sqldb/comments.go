package sqldb

import (
	"fmt"
	"github.com/treant5612/ytvc-web/model"
	"time"
)

func InsertComment(comment *model.Comment) (err error) {
	comment.Time = time.Now()
	return db.Create(comment).Error
}

func ListComment(pageNo int, pageSize int) (comments []*model.Comment, err error) {
	if pageNo < 1 || pageSize < 1 {
		return nil, fmt.Errorf("invalid parameter:pageNo=%d,pageSize=%d", pageNo, pageSize)
	}
	offset := (pageNo - 1) * pageSize
	err = db.Order("time desc").
		Offset(offset).
		Limit(pageSize).
		Find(&comments).Error
	return
}

func DeleteComment(comment *model.Comment) (err error) {
	if comment.ID == 0 {
		return
	}
	return db.Delete(comment).Error
}

func FindCommentById(id int) (comment *model.Comment, err error) {
	comment = new(model.Comment)
	err = db.First(comment, id).Error
	return
}
