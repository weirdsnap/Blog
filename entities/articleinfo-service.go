package entities

// import ("fmt")
//ArticleInfoAtomicService .
type ArticleInfoAtomicService struct{}

//ArticleInfoService .
var ArticleInfoService = ArticleInfoAtomicService{}

// Save .
func (*ArticleInfoAtomicService) Save(u *ArticleInfo) error {

	tx, err := mydb.Begin()
	checkErr(err)
	// fmt.Println("开始保存")
	dao := ArticleInfoDao{tx}
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

// FindByUsername .
func (*UserInfoAtomicService) FindByUsername(username string) *UserInfo {
	dao := userInfoDao{mydb}
	return dao.FindByUsername(username)
}