package cleans

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	usersPath = "C:/Users/"
	dirs      = []string{"Desktop", "Downloads", "Documents"}
)

func CleanUsers(userAdded, dirsExcept, filesExcept, formats []string, DEBUG bool) (err error) {
	if DEBUG {
		fmt.Println("Пользователели!")
	}

	var (
		users []string
		temp  []string
	)
	temp, _ = filepath.Glob(usersPath + "*.*")
	for _, user := range userAdded {
		temp = append(temp, usersPath+user)
	}
	for _, user := range temp {
		if isDir(user) {
			users = append(users, user)
		}
	}

	if DEBUG {
		fmt.Println("Пользователи: ", users)
	}

	for _, user := range users {
		if DEBUG {
			fmt.Println()
			fmt.Println("Каталог: ", user)
		}
		for _, dir := range dirs {
			if DEBUG {
				fmt.Println()
				fmt.Println("Каталог: ", dir)
			}
			for _, format := range formats {
				path := fmt.Sprintf("%s\\%s", user, dir)
				//Каталог рабочего стола
				if DEBUG {
					fmt.Println("Каталог и удаляемый формат: ", path, format)
				}
				err = removeFilesByMaskUser(path, format, filesExcept, dirsExcept, DEBUG)
				if err != nil {
					fmt.Printf("%v\n", err.Error())
					break
				}

			}
		}
	}
	return nil
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
