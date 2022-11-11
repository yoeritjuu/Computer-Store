package handlers

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"reflect"
	"text/template"

	"github.com/yoeritjuu/Computer-Store/pkg/parts"
)

type Node struct {
	Contact_id  int
	Employer_id int
	First_name  string
	Middle_name string
	Last_name   string
}

var templateFuncs = template.FuncMap{"rangeStruct": RangeStructer}

func GetPartsHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("computer_parts.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	partsList := parts.GetAllParts(data)

	t := template.New("t").Funcs(templateFuncs)
	t, err = t.Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, partsList)
	if err != nil {
		panic(err)
	}
}

// RangeStructer takes the first argument, which must be a struct, and
// returns the value of each field in a slice. It will return nil
// if there are no arguments or first argument is not a struct
func RangeStructer(args ...interface{}) []interface{} {
	if len(args) == 0 {
		return nil
	}

	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.Struct {
		return nil
	}

	out := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		out[i] = v.Field(i).Interface()
	}

	return out
}

var htmlTemplate = `
<head>
    <script type="text/javascript" charset="utf8" src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
    <link rel="stylesheet" type="text/css" href="//cdn.datatables.net/1.10.16/css/jquery.dataTables.css">
    <script type="text/javascript" charset="utf8" src="//cdn.datatables.net/1.10.16/js/jquery.dataTables.js"></script>
</head>

<h1>Parts in WebStore</h1>
<table id="fancytable" class="display">
	<col width="35%">
    <col width="65%">
	<tr>
		<th>ID</th>
		<th>Description</th>
		<th>Brand</th>
		<th>Color</th>      
		<th>Price (€)<th>
	</tr>
	
	{{range .}}
		<tr>
		{{range rangeStruct .}} 
			<td>{{.}}</td>
		{{end}}
		</tr>
	{{end}}
</table>
<script>$(document).ready(function() {
    $('#fancytable').DataTable();
} );</script>  
`
