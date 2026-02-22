package contact

import (
	"net/http"
	"strings"
)

// DefaultFields returns the standard Name, Email, Subject, Message fields.
func DefaultFields() []Field {
	return []Field{
		{Name: "name", Label: "Name", Type: "text", Placeholder: "Your name", Required: true},
		{Name: "email", Label: "Email", Type: "email", Placeholder: "you@example.com", Required: true},
		{Name: "subject", Label: "Subject", Type: "text", Placeholder: "What is this about?", Required: true},
		{Name: "message", Label: "Message", Type: "textarea", Placeholder: "Your message...", Rows: 5, Required: true},
	}
}

// ParseForm reads values from *http.Request for each field (trimmed).
func ParseForm(r *http.Request, fields []Field) FormData {
	data := FormData{
		Values: make(map[string]string, len(fields)),
		Errors: make(map[string]string),
	}
	for _, f := range fields {
		data.Values[f.Name] = strings.TrimSpace(r.FormValue(f.Name))
	}
	return data
}

// ValidateRequired checks Required fields are non-empty, sets Errors.
// Returns true if valid.
func ValidateRequired(fields []Field, data *FormData) bool {
	if data.Errors == nil {
		data.Errors = make(map[string]string)
	}
	valid := true
	for _, f := range fields {
		if f.Required && data.Values[f.Name] == "" {
			data.Errors[f.Name] = f.Label + " is required"
			valid = false
		}
	}
	return valid
}

// textareaRows defaults 0 to 5.
func textareaRows(rows int) int {
	if rows == 0 {
		return 5
	}
	return rows
}
