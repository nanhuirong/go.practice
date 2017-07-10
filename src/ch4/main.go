package main

import (
	"fmt"
)

func main()  {

	/**
	 切片
	 */
	slice := []int{10, 20, 30, 40}
	for _, elem := range slice {
		fmt.Println("%d", elem)
	}

	for index := 2; index < len(slice); index++ {
		fmt.Println("%d", slice[index])
	}

	//多维切片
	slice1 := [][]int{{10}, {100, 200}}
	slice1[0] = append(slice1[0], 20)

	/**
	 映射,用于存储无序键值对
	  1.数组,内部存储散列键的高8位,用于区分放入那个桶中
	  2.字节数组,用于存储键值对
	 */
	dict := make(map[string]int)
	//dict := map[string]int{}
	//dict := map[string]int{"nan": 1, "wang": 2}
	dict["nan"] = 1
	dict["wang"] = 2
	//value, exist := dict["nan"]
	//if exist {
	//	fmt.Println(value)
	//}

	for key, value := range dict{
		fmt.Printf("%s - %d\n", key, value)
	}
	delete(dict, "nan")

}
