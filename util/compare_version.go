package util

import (
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

	// compare first element, if not equal, then return ret, or continue compare
	var firstCompareResult = strings.Compare(oldvs[0], newvs[0])
	if firstCompareResult != 0 {
		return firstCompareResult
	}

	// compare the other element
	for i := 1; i <= len(newvs)-1; i++ {
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
