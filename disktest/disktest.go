package disktest

import (
	"fmt"
	"github.com/oneclickvirt/disktest/disk"
	"runtime"
)

func DiskTest(language, testMethod, testPath string, isMultiCheck bool) {
	var res string
	if runtime.GOOS == "windows" {
		if testMethod != "winsat" && testMethod != "" {
			res = "Detected host is Windows, using Winsat for testing.\n"
		}
		res = disk.WinsatTest(language, isMultiCheck, testPath)
	} else {
		switch testMethod {
		case "fio":
			res = disk.FioTest(language, isMultiCheck, testPath)
			if res == "" {
				res = "Fio test failed, switching to DD for testing.\n"
				res += disk.DDTest(language, isMultiCheck, testPath)
			}
		case "dd":
			res = disk.DDTest(language, isMultiCheck, testPath)
			if res == "" {
				res = "DD test failed, switching to Fio for testing.\n"
				res += disk.FioTest(language, isMultiCheck, testPath)
			}
		default:
			res = "Unsupported test method specified.\n"
		}
	}
	//fmt.Println("--------------------------------------------------")
	fmt.Printf(res)
	//fmt.Println("--------------------------------------------------")
}
