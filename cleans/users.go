package cleans

import (
	"fmt"
	"path/filepath"
)

var usersPath = "C:/Users/"

func CleanUsers(userAdded []string, formats []string, DEBUG bool) (err error) {
	fmt.Println("Каталоги Пользователей!")
	var users []string
	users, _ = filepath.Glob(usersPath + "*.*")
	for _, user := range userAdded {
		users = append(users, usersPath+user)
	}

	fmt.Println("Пользователи: ", users)

	//for _, user := range users {
	//	for _, format := range formats {
	//
	//		fmt.Println(fmt.Sprintf("%s/%s", user, "Desktop"), fmt.Sprintf(".%s", format))
	//		err = removeFilesByMask(fmt.Sprintf("%s/%s", user, "Desktop"), fmt.Sprintf(".%s", format))
	//		if err != nil {
	//			return err
	//		}
	//
	//		err = removeFilesByMask(fmt.Sprintf("%s/%s", user, "Documents"), fmt.Sprintf(".%s", format))
	//		if err != nil {
	//			return err
	//		}
	//
	//	}
	//}
	return nil
}
