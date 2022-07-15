<html>
<head>
<title></title>
</head>
<body>
文件上传  enctype属性有如下三种情况:
application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
multipart/form-data   不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
text/plain    空格转换为 "+" 加号，但不对特殊字符编码。

<form action="/login" method="post"   enctype="multipart/form-data" >
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">

    年龄:<input type="text" name="age">
    姓名:<input type="text" name="realname">
    英文姓名:<input type="text" name="engname">

    <select name="fruit">
    <option value="apple">apple</option>
    <option value="pear">pear</option>
    <option value="banane">banane</option>

    </select>


<input type="radio" name="gender" value="1">男
<input type="radio" name="gender" value="2">女


<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球




    <input type="submit" value="登陆">

</form>
</body>
</html>