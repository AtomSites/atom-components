package contact_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/AtomSites/atom-components/contact"
)

func TestContactForm(t *testing.T) {
	var buf bytes.Buffer
	err := contact.ContactForm("/contact", contact.FormData{}).Render(context.Background(), &buf)
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
	data := contact.FormData{
		Name:    "Alice",
		Email:   "alice@example.com",
		Subject: "Hello",
		Message: "Hi there",
		Errors:  map[string]string{"email": "Invalid email"},
	}
	var buf bytes.Buffer
	err := contact.ContactForm("/send", data).Render(context.Background(), &buf)
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
