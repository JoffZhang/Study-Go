package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/**
	sql.Open()函数用来打开一个注册过的数据库驱动，go-sql-driver中注册了mysql这个数据库驱动，第二个参数是DSN(Data Source Name)，它是go-sql-driver定义的一些数据库链接和配置信息。它支持如下格式：
	user@unix(/path/to/socket)/dbname?charset=utf8
	user:password@tcp(localhost:5555)/dbname?charset=utf8
	user:password@/dbname
	user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

	db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。

	db.Query()函数用来直接执行Sql返回Rows结果。

	stmt.Exec()函数用来执行stmt准备好的SQL语句

	我们可以看到我们传入的参数都是=?对应的数据，这样做的方式可以一定程度上防止SQL注入。
	*/
	db, err := sql.Open("mysql", "root:123456@(192.168.120.143)/test?charset=utf8")
	checkErr(err)

	//插入
	stmt, err := db.Prepare("insert into s_user set login =? ,uname = ? ")
	checkErr(err)

	res, err := stmt.Exec("研发", "研发")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update s_user set uname=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT uid,uname,login FROM s_user order by uid desc limit 3")
	checkErr(err)

	for rows.Next() {
		var uid int
		var uname string
		var login string
		err = rows.Scan(&uid, &uname, &login)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(uname)
		fmt.Println(login)
	}

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
