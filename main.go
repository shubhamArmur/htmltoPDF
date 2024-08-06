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
	// fileContent, err := os.ReadFile("ReadMe.md")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// htmlContent := helper.Highli(fileContent)

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

	// get markdown as html
	documentation := helper.Highli([]byte(audit.Documentation))
	security := helper.Highli([]byte(audit.Security))
	codeFix := helper.Highli([]byte(audit.CodeFix))
	testcase := helper.Highli([]byte(audit.Testcase))

	// Set the HTML content
	htmlContent := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Audit Report</title>
	</head>
	<body>
		<h1>` + audit.Name + `</h1>
		<h2>Documentation</h2>
		<p>` + documentation + `</p>
		<h2>Audit</h2>
		<h2>Gas Optimisation</h2>
		<p>` + audit.GasOptimisation + `</p>
		<h2>Code Fix</h2>
		<p>` + codeFix + `</p>
		<h2>Testcase</h2>
		<p>` + testcase + `</p>
		<h2>Security</h2>
		<p>` + security + `</p>
	`
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
