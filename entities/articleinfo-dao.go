package entities

type articleInfoDao DaoSource

var articleInfoInsertStmt = "INSERT INTO articleinfo(aid,title,class,content) values(?,?,?,?)"

// Save .
func (dao *articleInfoDao) Save(a *ArticleInfo) error {
	stmt, err := dao.Prepare(articleInfoInsertStmt)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(a.AID, a.Title, a.Class, a.Content)
	checkErr(err)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.AID = int(id)
	return nil
}

var articleInfoQueryAll = "SELECT * FROM articleinfo"
var articleInfoQueryByID = "SELECT * FROM articleinfo where aid = ?"
var articleInfoQueryByTitle = "SELECT * FROM articleinfo where title = ?"

// FindAll .
func (dao *articleInfoDao) FindAll() []ArticleInfo {
	rows, err := dao.Query(articleInfoQueryAll)
	checkErr(err)
	defer rows.Close()

	alist := make([]ArticleInfo, 0, 0)

	for rows.Next() {
		a := ArticleInfo{}
		err := rows.Scan(&a.AID, &a.Title, &a.Class, &a.Content)
		checkErr(err)
		alist = append(alist, a)
	}

	return alist
}

// FindByID .
func (dao *articleInfoDao) FindByID(id int) *ArticleInfo {
	stmt, err := dao.Prepare(articleInfoQueryByID)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(id)
	a := ArticleInfo{}
	err = row.Scan(&a.AID, &a.Title, &a.Class, &a.Content)
	checkErr(err)

	return &a
}

// FindByTitle .
func (dao *articleInfoDao) FindByTitle(Title string) *ArticleInfo {
	stmt, err := dao.Prepare(articleInfoQueryByTitle)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(Title)
	a := ArticleInfo{}
	err = row.Scan(&a.AID, &a.Title, &a.Class, &a.Content)
	checkErr(err)

	return &a
}
