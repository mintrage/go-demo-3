package main

import "fmt"

type bookmarkMap map[string]string

func main() {
	bookmarks := bookmarkMap{}
	var input int
Menu:
	for {
		input = getMenu()
		switch input {
		case 1:
			showBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			fmt.Println("Выход из программы")
			break Menu
		}
	}

}

func showBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Нет закладок для отображения")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}

}

func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Println("Добавление закладки")
	fmt.Println("Введите название закладки:")
	fmt.Scan(&newBookmarkKey)
	fmt.Println("Введите URL закладки:")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBookmark(bookmarks bookmarkMap) {
	var name string
	fmt.Println("Удаление закладки")
	fmt.Println("Введите название закладки для удаления:")
	fmt.Scan(&name)
	delete(bookmarks, name)
}

func getMenu() int {
	var input int
	fmt.Println("Выберите действие: ")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&input)
	return input
}
