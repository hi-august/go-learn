##Go笔记
####以下情况会导致编译错误
http://www.qiukeke.com/2015/05/28/gotchas-and-common-mistakes-in-go-golang.html
1. {不能单独放一行
2. 未使用的变量
3. 未使用的包
4. 函数外使用 a := 3
5. 函数内精简赋值,重复赋值
6. 一个变量的作用域是一个代码块
7. 除非特别指定,或者不能使用nil对变量赋值
8. 初始值nil的slice可以进行添加操作,但是map不能进行添加操作
9. 创建map时可以指定map的长度,运行时无法使用cap重新指定map大小
10. 字符串无法为nil,为空值""
