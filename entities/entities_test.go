package entities

import (
	"testing"
)

func Test_init(t *testing.T) {
	//a := NewArticleInfo(ArticleInfo{AID: 123, Title: "asd", Class: "asd", Content: "asd"})

	// if e := ArticleInfoService.Save(a); e != nil {
	//} else {
	//	t.Log("save出错")
	//}
	//if e := ArticleInfoService.FindByTitle("asd"); e != nil {
	if a := ArticleInfoService.FindByID(123); a != nil {
		t.Log("find测试通过", a)
	} else {
		t.Error("find测试为通过了")
	}
}
