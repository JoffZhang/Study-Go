//https://www.w3cschool.cn/yqbmht/9anweozt.html
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据

	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./Web/login.gtpl") //解析模板
		t.Execute(w, nil)                               //渲染模板，并发送
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		//解析表单
		r.ParseForm()
		validateForm(r)
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

func validateForm(r *http.Request) bool {
	//r.Form.Get()只能获取单个的值
	if len(r.Form["username"][0]) == 0 {
		//验证 为空
	}
	getint, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		//数字转化出错，可能不是数字
	}
	if getint > 100 {

	}
	//正则匹配
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {

	}
	//中文
	if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
		return false
	}
	//英文
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
		return false
	}
	//邮箱手机号  ^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$    ^(1[3|4|5|8][0-9]\d{4,8})$

	//下拉
	slice := []string{"apple", "pear", "banane"}
	for _, v := range slice {
		if v == r.Form.Get("fruit") {
			return true
		}
	}
	//单选按钮
	slice = []int{1, 2}
	for _, v := range slice {
		if v == r.Form.Get("gender") {
			return true
		}
	}
	//复选框
	//slice:=[]string{"football","basketball","tennis"}
	//a:=Slice_diff(r.Form["interest"],slice)
	//if a == nil{
	//	return true
	//}

	//验证15位身份证，15位的是全部数字
	if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
		return false
	}

	//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		return false
	}
	return true
}

//预防跨站脚本
func crossSiteScripting(w http.ResponseWriter, r *http.Request) {
	//Go的html/template里面带有下面几个函数可以帮你转义
	//func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
	//func HTMLEscapeString(s string) string //转义s之后返回结果字符串
	//func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串

	fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
	fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
	template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端

	//我们输入的username是<script>alert()</script>

	//Go的html/template包默认帮你过滤了html标签，但是有时候你只想要输出这个<script>alert()</script>看起来正常的信息
	//请使用text/template
	t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	_ = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
	//或者使用template.HTML类型   "html/template"
	t, _ = template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	_ = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>!"))
	//转换成​template.HTML​后，变量的内容也不会被转义
	//转义的例子 "html/template"
	t, _ = template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	_ = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
}

//文件上传
//上传文件主要三步处理：
//表单中增加enctype="multipart/form-data"
//服务端调用r.ParseMultipartForm,把上传的文件存储在内存和临时文件中
//使用r.FormFile获取文件句柄，然后对文件进行存储等处理。
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	r.ParseMultipartForm(32 << 20)   //里面的参数表示maxMemory
	//调用ParseMultipartForm之后，上传的文件存储在maxMemory大小的内存里面，如果文件大小超过了maxMemory，那么剩下的部分将存储在系统的临时文件中。我们可以通过r.FormFile获取上面的文件句柄，然后实例中使用了io.Copy来存储文件
	//获取其他非文件字段信息的时候就不需要调用r.ParseForm，因为在需要的时候Go自动会去调用。而且ParseMultipartForm调用一次之后，后面再次调用不会再有效果。
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

}
func main() {
	http.HandleFunc("/", sayHelloName) //设置访问路由
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
