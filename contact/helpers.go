package contact

import (
	"net/http"
	"net/mail"
	"strings"
	"unicode/utf8"
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

// SanitizeNewlines strips \r and \n from all non-textarea field values.
// This prevents email header injection. Call after ParseForm, before validation.
func SanitizeNewlines(fields []Field, data *FormData) {
	for _, f := range fields {
		if f.Type == "textarea" {
			continue
		}
		v := data.Values[f.Name]
		v = strings.ReplaceAll(v, "\r", "")
		v = strings.ReplaceAll(v, "\n", "")
		data.Values[f.Name] = v
	}
}

// ValidateFormat checks email format and maximum field length.
// Email fields are validated with net/mail.ParseAddress (RFC 5322).
// All non-empty fields are checked against maxLen runes (0 = no limit).
// Errors are appended to data.Errors. Returns true if all checks pass.
func ValidateFormat(fields []Field, data *FormData, maxLen int) bool {
	if data.Errors == nil {
		data.Errors = make(map[string]string)
	}
	valid := true
	for _, f := range fields {
		v := data.Values[f.Name]
		if v == "" {
			continue
		}
		if f.Type == "email" {
			if _, err := mail.ParseAddress(v); err != nil {
				data.Errors[f.Name] = "Please enter a valid email address"
				valid = false
			}
		}
		if maxLen > 0 && utf8.RuneCountInString(v) > maxLen {
			data.Errors[f.Name] = f.Label + " is too long"
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
