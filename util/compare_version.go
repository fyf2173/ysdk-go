package util

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// CompareVersion compare oldv and newv, if oldv > newv return 1
// if oldv < newv return -1, otherwise return 0, and return -999 if
// version is bad format
func CompareVersion(oldv, newv string) int {
	if strings.Compare(oldv, newv) == 0 {
		return 0
	}
	var (
		oldvs = strings.Split(oldv, ".")
		newvs = strings.Split(newv, ".")

		appendEle = func(t *[]string, c int) {
			for i := 0; i < c; i++ {
				*t = append(*t, "0")
			}
		}
	)

	// if len(oldvs) not equal len(newvs), then make up "0" for the short
	if len(oldvs) >= len(newvs) {
		appendEle(&newvs, len(oldvs)-len(newvs))
	} else {
		appendEle(&oldvs, len(newvs)-len(oldvs))
	}

	// compare the other element
	for i := 0; i <= len(newvs)-1; i++ {
		var (
			err    error
			ov, nv int
		)
		ov, err = strconv.Atoi(oldvs[i])
		if err != nil {
			log.Printf("Bad format %+v \n", oldvs[i])
			return -999
		}
		nv, err = strconv.Atoi(newvs[i])
		if err != nil {
			log.Printf("Bad format %+v \n", newvs[i])
			return -999
		}
		if ov == nv {
			continue
		} else if ov > nv {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

// CompareVersion2 compare oldv and newv, if oldv > newv return 1
// if oldv < newv return -1, otherwise return 0, and return -999 if
// version is bad format
func CompareVersion2(version1, version2 string) int {
	if strings.Compare(version1, version2) == 0 {
		return 0
	}

	var (
		oldvs = strings.Split(version1, ".")
		newvs = strings.Split(version2, ".")

		appendEle = func(t *[]string, c int) {
			for i := 0; i < c; i++ {
				*t = append(*t, "0")
			}
		}
	)

	// if len(oldvs) not equal len(newvs), then make up "0" for the short
	if len(oldvs) >= len(newvs) {
		appendEle(&newvs, len(oldvs)-len(newvs))
	} else {
		appendEle(&oldvs, len(newvs)-len(oldvs))
	}

	// compare every phase version
	for i := 0; i <= len(newvs)-1; i++ {
		var maxLen = len(oldvs[i])
		if len(newvs[i]) > maxLen {
			maxLen = len(newvs[i])
		}

		format := fmt.Sprintf("%%0%ds", maxLen)
		result := strings.Compare(fmt.Sprintf(format, oldvs[i]), fmt.Sprintf(format, newvs[i]))
		if result != 0 {
			return result
		}
	}

	return 0
}
