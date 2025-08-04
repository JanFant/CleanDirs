package cleans

import (
	"fmt"
)

var scanPath = "c:/"

func CleanScans(paths []string, formats []string, exceptFiles []string, DEBUG bool) (err error) {
	if DEBUG {
		fmt.Println("Каталога Сканов!")
	}

	if len(paths) == 0 || len(formats) == 0 {
		return fmt.Errorf("scans or formats is empty")
	}
	for _, path := range paths {
		for _, format := range formats {
			err = removeFilesByMask(fmt.Sprintf(scanPath+path), format, exceptFiles, DEBUG)
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
