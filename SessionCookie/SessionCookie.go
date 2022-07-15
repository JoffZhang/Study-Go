package main

import (
	"fmt"
	"net/http"
	"time"
)

/**
知道session管理涉及到如下几个因素

	全局session管理器
	保证sessionid 的全局唯一性
	为每个客户关联一个session
	session 的存储(可以存储到内存、文件、数据库等)
	session 过期处理
*/
func cookieDemo(w http.ResponseWriter, r http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)

	cookie, _ = r.Cookie("username")
	fmt.Fprint(w, cookie)

}
