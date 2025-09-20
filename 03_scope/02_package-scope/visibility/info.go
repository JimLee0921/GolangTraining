// Package visibility 演示包作用域内可导出与不可导出的标识符。
package visibility

var MyName = "JimLee" // 导出，可以在 main 包里访问

var YourName = "DSB" // 导出，可以在 main 包里访问

var secret = "13ff-3323-vcv5-cke1-ck22-gkk4" // secret（小写） -> 未导出，包外不能访问。
