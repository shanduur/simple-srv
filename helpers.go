package server

import (
	// embed is imported for template file
	_ "embed"
	"strconv"
	"strings"

	"log"
	"net/http"
)

//go:embed template/error-page.html
// TemplateErrorPage contains HTML of the Error Page.
//
// The default Error Page is defined as:
//	<head>
//	  <title>Error</title>
//	  <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
//	</head>
//	<body>
//	  <h1>ERROR {{ .StatusCode }}</h1>
//	  <p>
//	    <b>Cause:</b> {{ .Message }}
//	  </p>
//	</body>
var TemplateErrorPage string

func prepareErrorPage(err error, statusCode int) []byte {
	out := TemplateErrorPage
	out = strings.ReplaceAll(out, "{{ .StatusCode }}", strconv.Itoa(statusCode))
	out = strings.ReplaceAll(out, "{{ .Message }}", err.Error())

	return []byte(out)
}

// PrintError is a function that parses error in the handler functions.
// Its main functionality is to create cimple page displaying error code,
// as well as some basic information about the error to the end user.
func PrintError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)

	w.Write(prepareErrorPage(err, statusCode))

	log.Print(err.Error())
}
