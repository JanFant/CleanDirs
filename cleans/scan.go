package cleans

import (
	"fmt"
)

var scanPath = "c:/"

func CleanScans(paths, dirsExcept, exceptFiles, formats []string, DEBUG bool) (err error) {
	if DEBUG {
		fmt.Println("Сканы!")
	}

	if len(paths) == 0 || len(formats) == 0 {
		return fmt.Errorf("сакны и форматы в конфиге пусты")
	}
	for _, path := range paths {
		for _, format := range formats {
			err = removeFilesByMaskUser(fmt.Sprintf(scanPath+path), format, exceptFiles, dirsExcept, DEBUG)
			if err != nil {
				if DEBUG {
					fmt.Println("Не найден каталок: ", scanPath+path)
				}
				break
			}
		}
		if DEBUG {
			fmt.Println()
		}
	}
	return nil
}
