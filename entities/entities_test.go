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
	if e := ArticleInfoService.FindByID(123); e != nil {
		t.Error("find测试没通过", e)
	} else {
		t.Log("find测试通过了")
	}
}
