package cleans

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//func formatRemove(path string, formats []string) (err error) {
//	for _, format := range formats {
//		pattern := fmt.Sprintf("%s/*.%s", path, format)
//		files, _ := filepath.Glob(pattern)
//		fmt.Println(files)
//		if len(files) > 0 {
//			for _, file := range files {
//				err = os.Remove(file)
//				if err != nil {
//					return err
//				}
//				fmt.Printf("Removed %s\n", file)
//			}
//		}
//	}
//	return nil
//}

//func removeFilesByMaskScan(root, mask string, exceptFiles []string, DEBUG bool) error {
//	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return fmt.Errorf("нет каталога %v", root)
//		}
//		if !info.IsDir() {
//			// Проверяем соответствие маске
//			if strings.Contains(info.Name(), mask) {
//				for _, except := range exceptFiles {
//					if info.Name() == except {
//						if DEBUG {
//							fmt.Println("Исключение: ", except)
//						}
//						return nil
//					}
//				}
//				//Удаляем файл
//				if DEBUG {
//					fmt.Printf("Удален файл: %s\n", path)
//				} else {
//					err := os.Remove(path)
//					if err != nil {
//						return fmt.Errorf("ошибка удаления файла %s: %w", path, err)
//					}
//				}
//			}
//		}
//		return nil
//	})
//}

func removeFilesByMaskUser(root, mask string, exceptFiles, exceptDirs []string, DEBUG bool) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("нет каталога %v", root)
		}
		for _, except := range exceptDirs {
			if strings.Contains(path, except) {
				return nil
			}
		}
		if !info.IsDir() {
			// Проверяем соответствие маске
			if strings.Contains(info.Name(), mask) {
				for _, except := range exceptFiles {
					if info.Name() == except {
						if DEBUG {
							fmt.Println("Исключение: ", except)
						}
						return nil
					}
				}
				//Удаляем файл
				if DEBUG {
					fmt.Printf("Удален файл: %s\n", path)
				} else {
					err := os.Remove(path)
					if err != nil {
						return fmt.Errorf("ошибка удаления файла %s: %w", path, err)
					}
				}
			}
		}
		return nil
	})
}
