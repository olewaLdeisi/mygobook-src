package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	// 声明一个map
	var personDB map[string]PersonInfo
	// 创建map
	personDB = make(map[string]PersonInfo)
	/*
		// 创建并初始化
		var personDB map[string] PersonInfo{
			"1234": PersonInfo{"1", "Jack", "Room 101,..."},
		}
	*/

	// 插入数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{
		ID:      "1",
		Name:    "Jack",
		Address: "Room 101,...",
	}
	//
	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

	// 元素删除, 如果key不存在则什么都不发生，如果传入nil则抛异常panic
	delete(personDB, "1234")

}
