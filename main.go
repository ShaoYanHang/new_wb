package main

import {
	"app2/route"
	"app2/model"
}

func main() {
	// 引用数据库
	model.InitDb()
	// 引入路由组件
	route.InitRouter()
}
