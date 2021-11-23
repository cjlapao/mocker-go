package templates

import (
	"bytes"
	"text/template"

	"github.com/cjlapao/mocker-go/mocker"
)

type Company struct {
}

func (c Company) Name() string {
	return mocker.New().Company().Name()
}

func (c Company) JobTitle() string {
	return mocker.New().Company().JobTitle()
}

type Test struct {
	NameCalculated string
	Company        Company
	Date           Date
	Lorem          Lorem
}

func (t Test) Name() string {
	return mocker.New().Names().Name()
}

func JustDoIt(templateContent string) string {
	td := Test{}
	td.NameCalculated = mocker.New().Names().Name()
	t, err := template.New("parsedTemplate").Parse(templateContent)

	if err != nil {
		return "error in template"
	}

	var buff bytes.Buffer
	t.Execute(&buff, td)
	return string(buff.String())
}
