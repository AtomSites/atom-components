package form_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/AtomSites/atom-components/form"
)

func TestTextInput(t *testing.T) {
	var buf bytes.Buffer
	err := form.TextInput("email", "email", "Email", "email", "you@example.com", "", "").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `id="email"`) {
		t.Error("expected id attribute")
	}
	if !strings.Contains(html, `name="email"`) {
		t.Error("expected name attribute")
	}
	if !strings.Contains(html, `type="email"`) {
		t.Error("expected type attribute")
	}
	if !strings.Contains(html, "ac-label") {
		t.Error("expected ac-label class")
	}
	if !strings.Contains(html, "ac-input") {
		t.Error("expected ac-input class")
	}
	if strings.Contains(html, "ac-input-error") {
		t.Error("should not have error class when no error")
	}
	if strings.Contains(html, "ac-error-text") {
		t.Error("should not render error text when no error")
	}
}

func TestTextInputWithError(t *testing.T) {
	var buf bytes.Buffer
	err := form.TextInput("name", "name", "Name", "text", "", "", "Name is required").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "ac-input-error") {
		t.Error("expected ac-input-error class")
	}
	if !strings.Contains(html, "ac-error-text") {
		t.Error("expected error text element")
	}
	if !strings.Contains(html, "Name is required") {
		t.Error("expected error message text")
	}
}

func TestTextArea(t *testing.T) {
	var buf bytes.Buffer
	err := form.TextArea("msg", "message", "Message", "Type here...", "", 5, "").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `<textarea`) {
		t.Error("expected textarea element")
	}
	if !strings.Contains(html, "ac-textarea") {
		t.Error("expected ac-textarea class")
	}
	if !strings.Contains(html, `rows="5"`) {
		t.Error("expected rows attribute")
	}
}

func TestSelect(t *testing.T) {
	options := []form.SelectOption{
		{Value: "", Label: "Choose..."},
		{Value: "a", Label: "Option A"},
		{Value: "b", Label: "Option B", Selected: true},
	}
	var buf bytes.Buffer
	err := form.Select("sel", "selection", "Pick one", options, "").Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `<select`) {
		t.Error("expected select element")
	}
	if !strings.Contains(html, "ac-select") {
		t.Error("expected ac-select class")
	}
	if !strings.Contains(html, "Option A") {
		t.Error("expected Option A")
	}
	if !strings.Contains(html, "selected") {
		t.Error("expected selected attribute on Option B")
	}
}
