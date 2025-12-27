# sql.Out

开发基本用不到，主要是用来接收存储过程的 OUTPUT / INPUT 参数值，不用于普通 SELECT / UPDATE / INSERT，只用于：
调用存储过程并取输出参数

> 同样要求数据库和驱动支持，基本用不到，不再讲解

```
type Out struct {

	// Dest is a pointer to the value that will be set to the result of the
	// stored procedure's OUTPUT parameter.
	Dest any

	// In is whether the parameter is an INOUT parameter. If so, the input value to the stored
	// procedure is the dereferenced value of Dest's pointer, which is then replaced with
	// the output value.
	In bool
	// contains filtered or unexported fields
}
```