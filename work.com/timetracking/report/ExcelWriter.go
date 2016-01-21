package report

import (
    "github.com/xlsx"
    "path/filepath"
)

type ExcelWriter struct {
    file *xlsx.File
    sheet *xlsx.Sheet
}

const savedFilePath string = "./test.xlsx"

var excelWriter ExcelWriter

func NewExcelWriter() *ExcelWriter {
    return &excelWriter
}

func (this *ExcelWriter) Initialize(values []string, reportname string) {
    //pathToXLSX, _ := filepath.Abs(savedFilePath)
    //openFile, err := xlsx.OpenFile(pathToXLSX)
    //if err != nil {
        this.file = xlsx.NewFile()
/*    } else {
        this.file = openFile
    }*/
    this.sheet, _ = this.file.AddSheet(reportname)
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
    pathToXLSX, _ := filepath.Abs(savedFilePath)
    this.file.Save(pathToXLSX)
}
