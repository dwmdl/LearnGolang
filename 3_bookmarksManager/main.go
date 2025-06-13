package main

import "fmt"

type stringMap = map[string]string

func main() {
	bookmarks := stringMap{}

Menu:
	for {
		userMenuChoice := getUserMenuChoice()

		switch userMenuChoice {
		case 1:
			showBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		default:
			break Menu
		}
	}
}

func getUserMenuChoice() (userChoice int) {
	fmt.Println("1. See all bookmarks")
	fmt.Println("2. Add a bookmark")
	fmt.Println("3. Delete a bookmark")
	fmt.Println("4. Exit")

	fmt.Scan(&userChoice)

	return
}

func showBookmarks(bookmarks stringMap) {
	if len(bookmarks) == 0 {
		fmt.Println("You have not bookmarks yet")
	}

	for key, value := range bookmarks {
		fmt.Printf("Bookmark key: %s, Bookmark content: %s \n", key, value)
	}
}

func addBookmark(bookmarks stringMap) {
	var bookmarkTitle string
	var bookmarkContent string

	fmt.Println("Enter a bookmark title")
	fmt.Scan(&bookmarkTitle)

	fmt.Println("Enter the bookmark content")
	fmt.Scan(&bookmarkContent)

	bookmarks[bookmarkTitle] = bookmarkContent
}

func deleteBookmark(bookmarks stringMap) {
	var bookmarkTitle string

	fmt.Println("Which bookmark do you want delete? Enter the bookmark title: ")
	fmt.Scan(&bookmarkTitle)

	delete(bookmarks, bookmarkTitle)
}
