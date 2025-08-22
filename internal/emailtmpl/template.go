package emailtmpl

import (
	"bytes"
	"embed"
	"html/template"
	"log"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/valentino7504/tax-automation-go/internal/models"
)

//go:embed template.html
var emailTemplate embed.FS

// formatNaira formats a numeric amount represented by a float64 to a comma delimited
// Naira representation with two decimal places.
func formatNaira(amount float64) string {
	return message.NewPrinter(language.English).Sprintf("₦%.2f", amount)
}

// Render takes a recipient and applies the email template, returning the
// rendered HTML string
func Render(data models.Recipient) (*string, error) {
	var buf bytes.Buffer
	funcMap := template.FuncMap{
		"fmtAmt": formatNaira,
	}
	tmpl, err := template.New("template.html").Funcs(funcMap).ParseFS(
		emailTemplate,
		"template.html",
	)
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Fatalf("Error executing template %v", err)
	}
	htmlBody := buf.String()
	return &htmlBody, nil
}
