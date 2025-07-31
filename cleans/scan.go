package cleans

import (
	"fmt"
)

var scanPath = "c:/"

func CleanScans(paths []string, formats []string) (err error) {
	if len(paths) == 0 || len(formats) == 0 {
		return fmt.Errorf("scans or formats is empty")
	}
	for _, path := range paths {
		for _, format := range formats {
			err = removeFilesByMask(fmt.Sprintf(scanPath+path), format)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
