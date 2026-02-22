package contact_test

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/AtomSites/atom-components/contact"
)

func TestContactForm(t *testing.T) {
	fields := contact.DefaultFields()
	var buf bytes.Buffer
	err := contact.ContactForm("/contact", fields, contact.FormData{}, "").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `action="/contact"`) {
		t.Error("expected form action")
	}
	if !strings.Contains(html, `method="POST"`) {
		t.Error("expected POST method")
	}
	if !strings.Contains(html, "ac-contact-form") {
		t.Error("expected ac-contact-form class")
	}
	if !strings.Contains(html, `name="name"`) {
		t.Error("expected name field")
	}
	if !strings.Contains(html, `name="email"`) {
		t.Error("expected email field")
	}
	if !strings.Contains(html, `name="subject"`) {
		t.Error("expected subject field")
	}
	if !strings.Contains(html, `name="message"`) {
		t.Error("expected message field")
	}
	if !strings.Contains(html, "ac-contact-submit") {
		t.Error("expected submit button")
	}
}

func TestContactFormWithData(t *testing.T) {
	fields := contact.DefaultFields()
	data := contact.FormData{
		Values: map[string]string{
			"name":    "Alice",
			"email":   "alice@example.com",
			"subject": "Hello",
			"message": "Hi there",
		},
		Errors: map[string]string{"email": "Invalid email"},
	}
	var buf bytes.Buffer
	err := contact.ContactForm("/send", fields, data, "").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "Alice") {
		t.Error("expected prefilled name")
	}
	if !strings.Contains(html, "alice@example.com") {
		t.Error("expected prefilled email")
	}
	if !strings.Contains(html, "Invalid email") {
		t.Error("expected email error message")
	}
	if !strings.Contains(html, "ac-input-error") {
		t.Error("expected error class on email input")
	}
}

func TestContactFormCustomFields(t *testing.T) {
	fields := []contact.Field{
		{Name: "full_name", Label: "Full Name", Type: "text", Placeholder: "Jane Doe"},
		{Name: "body", Label: "Body", Type: "textarea", Placeholder: "Write here...", Rows: 10},
	}
	var buf bytes.Buffer
	err := contact.ContactForm("/custom", fields, contact.FormData{}, "").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `name="full_name"`) {
		t.Error("expected full_name field")
	}
	if !strings.Contains(html, `name="body"`) {
		t.Error("expected body field")
	}
	// Default fields should NOT be present
	if strings.Contains(html, `name="email"`) {
		t.Error("unexpected email field")
	}
	if strings.Contains(html, `name="subject"`) {
		t.Error("unexpected subject field")
	}
}

func TestParseForm(t *testing.T) {
	fields := []contact.Field{
		{Name: "name", Label: "Name", Type: "text", Required: true},
		{Name: "email", Label: "Email", Type: "email", Required: true},
	}

	body := strings.NewReader("name=Alice+Smith&email=alice%40example.com")
	req, err := http.NewRequest(http.MethodPost, "/contact", body)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	data := contact.ParseForm(req, fields)

	if data.Values["name"] != "Alice Smith" {
		t.Errorf("expected 'Alice Smith', got %q", data.Values["name"])
	}
	if data.Values["email"] != "alice@example.com" {
		t.Errorf("expected 'alice@example.com', got %q", data.Values["email"])
	}
}

func TestValidateRequired(t *testing.T) {
	fields := []contact.Field{
		{Name: "name", Label: "Name", Type: "text", Required: true},
		{Name: "note", Label: "Note", Type: "textarea", Required: false},
	}

	data := contact.FormData{
		Values: map[string]string{"name": "", "note": ""},
		Errors: make(map[string]string),
	}

	valid := contact.ValidateRequired(fields, &data)
	if valid {
		t.Error("expected validation to fail")
	}
	if data.Errors["name"] == "" {
		t.Error("expected error for required 'name' field")
	}
	if _, has := data.Errors["note"]; has {
		t.Error("unexpected error for optional 'note' field")
	}

	// Valid case
	data2 := contact.FormData{
		Values: map[string]string{"name": "Alice", "note": ""},
		Errors: make(map[string]string),
	}
	if !contact.ValidateRequired(fields, &data2) {
		t.Error("expected validation to pass")
	}
}

func TestContactFormWithCSRFToken(t *testing.T) {
	fields := contact.DefaultFields()
	var buf bytes.Buffer
	err := contact.ContactForm("/contact", fields, contact.FormData{}, "test-token-123").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `name="csrf_token"`) {
		t.Error("expected csrf_token hidden input")
	}
	if !strings.Contains(html, `value="test-token-123"`) {
		t.Error("expected csrf token value")
	}
	if !strings.Contains(html, `type="hidden"`) {
		t.Error("expected hidden input type")
	}
}

func TestContactFormWithoutCSRFToken(t *testing.T) {
	fields := contact.DefaultFields()
	var buf bytes.Buffer
	err := contact.ContactForm("/contact", fields, contact.FormData{}, "").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if strings.Contains(html, `name="csrf_token"`) {
		t.Error("csrf_token should not be rendered when token is empty")
	}
}

func TestValidateFormatInvalidEmail(t *testing.T) {
	fields := []contact.Field{
		{Name: "email", Label: "Email", Type: "email", Required: true},
	}
	data := contact.FormData{
		Values: map[string]string{"email": "not-an-email"},
		Errors: make(map[string]string),
	}

	valid := contact.ValidateFormat(fields, &data, 0)
	if valid {
		t.Error("expected validation to fail for invalid email")
	}
	if data.Errors["email"] == "" {
		t.Error("expected error for invalid email")
	}
}

func TestValidateFormatValidEmail(t *testing.T) {
	fields := []contact.Field{
		{Name: "email", Label: "Email", Type: "email", Required: true},
	}
	data := contact.FormData{
		Values: map[string]string{"email": "alice@example.com"},
		Errors: make(map[string]string),
	}

	valid := contact.ValidateFormat(fields, &data, 0)
	if !valid {
		t.Error("expected validation to pass for valid email")
	}
}

func TestValidateFormatMaxLength(t *testing.T) {
	fields := []contact.Field{
		{Name: "name", Label: "Name", Type: "text", Required: true},
	}
	data := contact.FormData{
		Values: map[string]string{"name": strings.Repeat("a", 11)},
		Errors: make(map[string]string),
	}

	valid := contact.ValidateFormat(fields, &data, 10)
	if valid {
		t.Error("expected validation to fail for too-long input")
	}
	if data.Errors["name"] == "" {
		t.Error("expected error for too-long name")
	}

	// Within limit
	data2 := contact.FormData{
		Values: map[string]string{"name": strings.Repeat("a", 10)},
		Errors: make(map[string]string),
	}
	if !contact.ValidateFormat(fields, &data2, 10) {
		t.Error("expected validation to pass for name at limit")
	}
}

func TestSanitizeNewlines(t *testing.T) {
	fields := []contact.Field{
		{Name: "name", Label: "Name", Type: "text"},
		{Name: "subject", Label: "Subject", Type: "text"},
		{Name: "message", Label: "Message", Type: "textarea"},
	}
	data := contact.FormData{
		Values: map[string]string{
			"name":    "Alice\r\nSmith",
			"subject": "Hello\nWorld",
			"message": "Line1\r\nLine2\nLine3",
		},
		Errors: make(map[string]string),
	}

	contact.SanitizeNewlines(fields, &data)

	if data.Values["name"] != "AliceSmith" {
		t.Errorf("expected newlines stripped from name, got %q", data.Values["name"])
	}
	if data.Values["subject"] != "HelloWorld" {
		t.Errorf("expected newlines stripped from subject, got %q", data.Values["subject"])
	}
	if data.Values["message"] != "Line1\r\nLine2\nLine3" {
		t.Errorf("expected newlines preserved in textarea, got %q", data.Values["message"])
	}
}
