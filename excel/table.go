package excel

import (
	"mime/multipart"

	"github.com/shakinm/xlsReader/xls"
	"github.com/tealeg/xlsx"
)

type Table struct {
	ft     string
	Sheets []*Sheet
}

type Sheet struct {
	Index   int
	Name    string
	Rows    []int
	RowsMap map[int][]Cell
	MaxCol  int
}

type Cell string

func NewFormTable(fh *multipart.FileHeader) *Table {
	t := new(Table)
	t.ReadFormFile(fh)
	return t
}

func NewLocalTable(filename string) *Table {
	t := new(Table)
	t.ReadLocalFile(filename)
	return t
}

func (t *Table) GetSheetByIndex(si int) *Sheet {
	for i, sheet := range t.Sheets {
		if i == si {
			return sheet
		}
	}
	return new(Sheet)
}

func (t *Table) GetSheetByName(sheetName string) *Sheet {
	for _, sheet := range t.Sheets {
		if sheet.Name == sheetName {
			return sheet
		}
	}
	return new(Sheet)
}

func (s *Sheet) GetRows() []int {
	return s.Rows
}

func (s *Sheet) GetRow(ri int) []Cell {
	return s.RowsMap[ri]
}

// ReadFormFile 从表单上传的文件读取excel内容，同时支持.xlsx,.xls格式文件
func (t *Table) ReadFormFile(fh *multipart.FileHeader) {
	sheets, err := ReadExcelFormFile(fh, &t.ft)
	if err != nil {
		return
	}
	t.readData(sheets)
	t.paddingCells()
	return
}

// ReadLocalFile 从本地文件读取excel内容，同时支持.xlsx,.xls格式文件
func (t *Table) ReadLocalFile(filename string) {
	sheets, err := ReadExcelLocalFile(filename, &t.ft)
	if err != nil {
		return
	}
	t.readData(sheets)
	t.paddingCells()
	return
}

func (t *Table) paddingCells() {
	for _, sheet := range t.Sheets {
		for ri, cells := range sheet.RowsMap {
			if len(cells) < sheet.MaxCol {
				for i := 0; i <= sheet.MaxCol-len(cells)-1; i++ {
					sheet.RowsMap[ri] = append(sheet.RowsMap[ri], "")
				}
			}
		}
	}
}

// readData 读取文件内容
func (t *Table) readData(sheets interface{}) {
	if t.ft == Xlsx {
		t.readXlsx(sheets.(*xlsx.File).Sheets)
		return
	}
	t.readXls(sheets.(xls.Workbook))
	return
}

// readXlsx 读取.xlsx格式的文件内容
func (t *Table) readXlsx(sheets []*xlsx.Sheet) {
	for i := 0; i <= len(sheets)-1; i++ {
		sheet := sheets[i]
		sht := new(Sheet)
		sht.Index = i
		sht.Name = sheet.Name
		sht.RowsMap = make(map[int][]Cell)
		for j := 0; j <= len(sheet.Rows)-1; j++ {
			row := sheet.Row(j)
			if sht.MaxCol < len(row.Cells) {
				sht.MaxCol = len(row.Cells)
			}
			sht.Rows = append(sht.Rows, j)
			for k := 0; k <= len(row.Cells)-1; k++ {
				col := row.Cells[k]
				sht.RowsMap[j] = append(sht.RowsMap[j], Cell(col.Value))
			}
		}

		t.Sheets = append(t.Sheets, sht)
	}
}

// readXls 读取.xls格式的文件内容
func (t *Table) readXls(wb xls.Workbook) {
	for i := 0; i <= len(wb.GetSheets())-1; i++ {
		sheet, err := wb.GetSheet(i)
		if err != nil {
			continue
		}
		sht := new(Sheet)
		sht.Index = i
		sht.Name = sheet.GetName()
		sht.RowsMap = make(map[int][]Cell)
		for j := 0; j <= len(sheet.GetRows())-1; j++ {
			row, _ := sheet.GetRow(j)
			if sht.MaxCol < len(row.GetCols()) {
				sht.MaxCol = len(row.GetCols())
			}
			sht.Rows = append(sht.Rows, j)
			for k := 0; k <= len(row.GetCols())-1; k++ {
				col, _ := row.GetCol(k)
				sht.RowsMap[j] = append(sht.RowsMap[j], Cell(col.GetString()))
			}
		}
		t.Sheets = append(t.Sheets, sht)
	}
}
