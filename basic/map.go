package main

import "fmt"

func main() {
	// 1. 创建map
	// 1.1 使用make
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 99
	scoreMap["李四"] = 89
	// 1.2
	userInfo := map[string]string{
		"username": "foo",
		"password": "foo123",
	}

	fmt.Println(userInfo)
	fmt.Println(scoreMap)

	// 2. map的常见操作
	// 2.1 判断键是否存在
	value, ok := userInfo["username"]
	fmt.Printf("value: %s\nok: %t\n", value, ok)

	// 2.2 map遍历
	for i, v := range userInfo {
		fmt.Printf("%s -> %s\n", i, v)
	}

	// 2.2 删除键值对
	userInfo["age"] = "12"
	delete(userInfo, "age")
	fmt.Println(userInfo)

	// 2.3 指定顺序遍历map

}
