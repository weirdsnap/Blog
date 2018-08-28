package entities

import (
	"fmt"
)

// Article Info Atomic Service .
type ArticleInfoAtomicService struct{}

// Article Info Service .
var ArticleInfoService = ArticleInfoAtomicService{}

// Save
// function
// on the *ArticleInfoAtomicService
func (*ArticleInfoAtomicService) Save(a *ArticleInfo) error {

	tx, err := mydb.Begin()
	checkErr(err)
	// fmt.Println("开始保存")
	dao := articleInfoDao{tx}
	err = dao.Save(a)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return nil
}

// FindAll .
func (*ArticleInfoAtomicService) FindAll() []ArticleInfo {
	dao := articleInfoDao{mydb}
	return dao.FindAll()
}

// FindByID .
func (*ArticleInfoAtomicService) FindByID(id int) *ArticleInfo {
	dao := articleInfoDao{mydb}
	fmt.Println("开始寻找", id)
	return dao.FindByID(id)
}

// FindByTitle .
func (*ArticleInfoAtomicService) FindByTitle(title string) *ArticleInfo {
	dao := articleInfoDao{mydb}
	return dao.FindByTitle(title)
}
