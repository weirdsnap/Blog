package entities

// import ("fmt")
// Article Info Atomic Service .
type ArticleInfoAtomicService struct{}

// Article Info Service .
var ArticleInfoService = ArticleInfoAtomicService{}

// Save
// function
// on the *ArticleInfoAtomicService
func (*ArticleInfoAtomicService) Save(u *ArticleInfo) error {

	tx, err := mydb.Begin()
	checkErr(err)
	// fmt.Println("开始保存")
	dao := articleInfoDao{tx}
	err = dao.Save(u)

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

// FindByTitle .
func (*ArticleInfoAtomicService) FindByTitle(title string) *ArticleInfo {
	dao := articleInfoDao{mydb}
	return dao.FindByTitle(title)
}
