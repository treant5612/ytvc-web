package sqldb

import (
	"github.com/jinzhu/gorm"
	"github.com/treant5612/ytvc-web/model"
	"os"
	"strconv"
	"strings"
	"testing"
)

func InitTestDb() {
	os.Remove("test.db")
	InitSqlite("test.db")
}

func TestInsertComment(t *testing.T) {
	InitTestDb()
	comment := &model.Comment{
		Content:  "test comment3",
		NickName: "pal",
		Addr:     "127.0.0.1",
	}
	err := InsertComment(comment)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListComment(t *testing.T) {
	InitTestDb()

	for i := 0; i < 100; i++ {
		InsertComment(&model.Comment{
			Content:  strings.Repeat(strconv.Itoa(i), 10),
			NickName: "com" + strconv.Itoa(i+1),
			Addr:     "addr",
		})
	}

	pageNo, pageSize := 1, 10
	comments, err := ListComment(pageNo, pageSize)
	if err != nil {
		t.Fatal(err)
	}
	if len(comments) != pageSize {
		t.Fatal(err)
	}
	for i, v := range comments {
		if v.ID != uint(100-(pageNo-1)*pageSize-i) {
			t.Fatalf("wrong id%v,page%d ,pageSize%d",v.ID,pageNo,pageSize)
		}
		if v.NickName != "com"+strconv.Itoa(int(v.ID)) {
			t.Fatal("wrong nickname",v.NickName)
		}
	}
}

func TestDeleteComment(t *testing.T) {
	 InitTestDb()
	for i := 0; i < 100; i++ {
		InsertComment(&model.Comment{
			Content:  strings.Repeat(strconv.Itoa(i), 10),
			NickName: "com" + strconv.Itoa(i+1),
			Addr:     "addr",
		})
	}
	_,err := FindCommentById(10)
	if err!=nil{
		t.Fatal(err)
	}
	err =DeleteComment(&model.Comment{ID:10})
	if err!=nil{
		t.Fatal("delete record error:",err)
	}
	_,err = FindCommentById(10)
	if err!=gorm.ErrRecordNotFound{
		t.Fatal("find record after deleting",err)

	}
}
