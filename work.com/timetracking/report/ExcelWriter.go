package report

import (
    "github.com/xlsx"
    "path/filepath"
)

type ExcelWriter struct {
    file *xlsx.File
    sheet *xlsx.Sheet
}

var excelWriter ExcelWriter

func NewExcelWriter() *ExcelWriter {
    return &excelWriter
}

func (this *ExcelWriter) Initialize(values []string) {
    this.file = xlsx.NewFile()
    this.sheet, _ = this.file.AddSheet("test")
}

func (this *ExcelWriter) PrintLine(teamName string, teamMembers string, projectname string, hours string, percent string) {
    row := this.sheet.AddRow()
    this.addCell(row, teamName)
    this.addCell(row, teamMembers)
    this.addCell(row, projectname)
    this.addCell(row, hours)
    this.addCell(row, percent)

}

func (this *ExcelWriter) addCell(row *xlsx.Row, value string) {
    cell := row.AddCell()
    cell.Value = value

}

func (this *ExcelWriter) Close() {
    pathToXLSX, _ := filepath.Abs("./test.xslx")
    this.file.Save(pathToXLSX)
}
