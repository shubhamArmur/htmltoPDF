package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"example.com/armur/helper"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Audit struct {
	Name            string `json:"name"`
	Audit           string `json:"audit"`
	GasOptimisation string `json:"gasOptimisation"`
	Documentation   string `json:"documentation"`
	CodeFix         string `json:"codeFix"`
	Testcase        string `json:"testcase"`
	Security        string `json:"security"`
}

func main() {
	fileContent, err := os.ReadFile("audit.json")
	if err != nil {
		fmt.Println(err)
	}

	// unmarshal json content
	var audit Audit
	err = json.Unmarshal(fileContent, &audit)
	if err != nil {
		fmt.Println(err)
	}
	// Create a new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println(err)
		return
	}

	pdfg.MarginBottom.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)

	// get markdown as html
	documentation := helper.ConvertMarkDownToHTML([]byte(audit.Documentation))
	security := helper.ConvertMarkDownToHTML([]byte(audit.Security))
	codeFix := helper.ConvertMarkDownToHTML([]byte(audit.CodeFix))
	testcase := helper.ConvertMarkDownToHTML([]byte(audit.Testcase))

	htmlContent := helper.ReportCreator(documentation, security, codeFix, testcase)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(htmlContent)))

	// Create PDF document in buffer
	err = pdfg.Create()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./output-colored.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("PDF created successfully")
}
