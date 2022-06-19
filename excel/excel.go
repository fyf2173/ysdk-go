package excel

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shakinm/xlsReader/xls"
	"github.com/tealeg/xlsx"
)

const (
	Xlsx = ".xlsx"
	Xls  = ".xls"
)

// ReadExcelFormFile 读取表单上传的excel文件内容，支持 .xlsx、.xls格式的文件
// 如果上传的文件格式是.xlsx，则返回*xlsx.File，否则返回xls.Workbook
func ReadExcelFormFile(fh *multipart.FileHeader, ft *string) (interface{}, error) {
	f, err := fh.Open()
	if err != nil {
		return nil, err
	}

	*ft = filepath.Ext(fh.Filename)
	return openFromReader(f, *ft)
}

// ReadExcelLocalFile 读取本地上传的excel文件内容，支持 .xlsx、.xls格式的文件
// 如果上传的文件格式是.xlsx，则返回*xlsx.File，否则返回xls.Workbook
func ReadExcelLocalFile(filename string, ft *string) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	*ft = filepath.Ext(filename)
	return openFromReader(f, *ft)
}

// ReadRemoteExcelFile 读取表单上传的excel文件内容，支持 .xlsx、.xls格式的文件
// 如果上传的文件格式是.xlsx，则返回*xlsx.File，否则返回*xls.Workbook
func ReadRemoteExcelFile(furl string, ft *string) (interface{}, error) {
	*ft = filepath.Ext(furl)
	if *ft != Xlsx && *ft != Xls {
		return nil, fmt.Errorf("file not support")
	}
	resp, err := http.Get(furl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	defer resp.Body.Close()

	fBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if *ft == Xlsx {
		return xlsx.OpenBinary(fBody)
	}
	return xls.OpenReader(bytes.NewReader(fBody))
}

// openFromReader 根据excel的文件后缀读取文件内容
func openFromReader(r io.ReadSeeker, ft string) (interface{}, error) {
	if ft == Xlsx {
		fBody, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		return xlsx.OpenBinary(fBody)
	}
	if ft == Xls {
		wb, err := xls.OpenReader(r)
		if err != nil {
			return nil, err
		}
		return wb, nil
	}
	return nil, fmt.Errorf("invalid file extension")
}
