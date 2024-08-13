package helper

func ReportCreator(documentation string, security string, codeFix string, testcase string) string {
	htmlString := `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Audit Report</title>
    <style>
        * {
            line-height: 1.6;
            letter-spacing: 1.4px;
        }

        body {
            background-color: #000000;
            color: #ffffff;
            padding: 20px;
            font-family: Arial, Helvetica, sans-serif;
            box-sizing: border-box;
        }

        h1 {
            margin-top: 2.5rem;
            margin-bottom: 2.5rem;
            color: #FFFFFFE5;
            font-size: 2.25rem;
            font-weight: 600;
            line-height: 1.8125rem;
            letter-spacing: normal;
            text-align: left;
        }

        h2,
        h3,
        h4,
        h5 {
            color: #66bb6a;
        }
    
        hr{
        
        }

        pre {
            padding: 50px;
			font-size: 14;
        }
    </style>
</head>

<body>
		<h1>Documentation</h1>
        <hr/>
		<div class="headline"></div>
		` + documentation + `
		<h1>Securrity</h1>
		` + security + `
		<h1>Code Fix</h1>
		` + codeFix + `
		<h1>Testcase</h1>
		` + testcase + `
</body>

</html>`
	return htmlString
}
