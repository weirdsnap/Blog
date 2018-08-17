package entities

import (
	"testing"
)

func Test_init(t *testing.T) {
	init()
	// a := NewArticleInfo(ArticleInfo{AID: 123, Title: "asd", Class: "asd", Content: "asd"})
	
	// if e := ArticleInfoService.Save(a); e != nil {
	if e := ArticleInfoService.FindByTitle("asd"); e != nil {
		t.Error("save函数测试没通过")
	} else {
		t.Log("save函数测试通过了")
	}
}
