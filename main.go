package main

import "fmt"

func main() {
	m := map[string]string{
		"PurpleSchool": "https://www.purpleschool.com",
	}
	fmt.Println(m)
	fmt.Println(m["PurpleSchool"])
	m["PurpleSchool"] = "https://www.purpleschool.ru"
	fmt.Println(m)
	m["Google"] = "https://www.google.com"
	fmt.Println(m)
	m["Yandex"] = "https://www.yandex.ru"
	fmt.Println(m)
	delete(m, "Yandex")
	fmt.Println(m)

}
