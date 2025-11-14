module example

go 1.25

// 告诉 Go：这个模块依赖一个名为 gee 的模块，版本号是 v0.0.0
require gee v0.0.0

// 替换依赖来源。它就在当前项目目录下的 ./gee 文件夹里
replace gee => ./gee