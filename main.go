package main

import "fmt"

func main() {
	bookmarks := map[string]string{}
	var input int
Menu:
	for {
		input = getMenu()
		switch input {
		case 1:
			showBookmarks(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			fmt.Println("Выход из программы")
			break Menu
		}
	}

}

func showBookmarks(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("Нет закладок для отображения")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}

}

func addBookmark(bookmarks map[string]string) map[string]string {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Println("Добавление закладки")
	fmt.Println("Введите название закладки:")
	fmt.Scan(&newBookmarkKey)
	fmt.Println("Введите URL закладки:")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
	return bookmarks
}

func deleteBookmark(bookmarks map[string]string) map[string]string {
	var name string
	fmt.Println("Удаление закладки")
	fmt.Println("Введите название закладки для удаления:")
	fmt.Scan(&name)
	delete(bookmarks, name)
	return bookmarks
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
