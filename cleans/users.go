package cleans

import (
	"fmt"
	"path/filepath"
)

var usersPath = "C:/Users/"

func CleanUsers(users []string, formats []string) (err error) {
	if len(users) == 0 {
		users, _ = filepath.Glob(usersPath + "*.*")
	} else {
		for i, user := range users {
			users[i] = fmt.Sprintf(usersPath + user)
		}
	}
	fmt.Println(users)
	for _, user := range users {
		for _, format := range formats {
			fmt.Println(fmt.Sprintf("%s/%s", user, "Desktop"), fmt.Sprintf(".%s", format))
			err = removeFilesByMask(fmt.Sprintf("%s/%s", user, "Desktop"), fmt.Sprintf(".%s", format))
			if err != nil {
				return err
			}
			err = removeFilesByMask(fmt.Sprintf("%s/%s", user, "Documents"), fmt.Sprintf(".%s", format))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
