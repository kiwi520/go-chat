package main

import (
	"fmt"
)

var userId int
var userPwd string

func main()  {
	//接受用户的选择
	var key int
	//判断是否还继续显示菜单
	var loop bool = true

	for loop {
		fmt.Println("------------欢迎登陆多人聊天系统----------------")
		fmt.Println("\t\t\t 1 登陆聊天系统")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出聊天系统")
		fmt.Println("\t\t\t 请选择（1-3）:")

		fmt.Scanf("%d\n",&key)

		switch key {
		case 1:
			fmt.Println("登陆聊天系统")
			loop =false
		case 2:
			fmt.Println("注册用户")
			loop =false
		case 3:
			fmt.Println("退出系统")
			loop = false
		}
	}

	//根据用户输入的信息，显示新的提示信息

	if key == 1 {
		//说明用户已登陆
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n",&userPwd)

		var err = Login(userId, userPwd)
		if err != nil{
			fmt.Println("登陆失败\n")
		}else {
			fmt.Println("登陆成功\n")
		}


	}else if key == 2{
		fmt.Println("注册用户")
	}
}
