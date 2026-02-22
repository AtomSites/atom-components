package modal_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"

	"github.com/AtomSites/atom-components/modal"
)

func TestModal(t *testing.T) {
	var buf bytes.Buffer
	m := modal.Modal("confirm", "Confirm Action")
	err := m.Render(templ.WithChildren(context.Background(), templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, e := io.WriteString(w, "<p>Are you sure?</p>")
		return e
	})), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `id="confirm"`) {
		t.Error("expected modal id")
	}
	if !strings.Contains(html, "ac-modal-overlay") {
		t.Error("expected ac-modal-overlay class")
	}
	if !strings.Contains(html, "ac-modal") {
		t.Error("expected ac-modal class")
	}
	if !strings.Contains(html, "Confirm Action") {
		t.Error("expected modal title")
	}
	if !strings.Contains(html, "data-modal-close") {
		t.Error("expected close button with data-modal-close")
	}
	if !strings.Contains(html, "Are you sure?") {
		t.Error("expected children content")
	}
	if !strings.Contains(html, `role="dialog"`) {
		t.Error("expected dialog role")
	}
}

func TestModalWithFooter(t *testing.T) {
	footer := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, e := io.WriteString(w, `<button class="btn">OK</button>`)
		return e
	})

	var buf bytes.Buffer
	m := modal.ModalWithFooter("dlg", "Dialog", footer)
	err := m.Render(templ.WithChildren(context.Background(), templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, e := io.WriteString(w, "<p>Body content</p>")
		return e
	})), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "ac-modal-footer") {
		t.Error("expected ac-modal-footer class")
	}
	if !strings.Contains(html, `class="btn"`) {
		t.Error("expected footer button")
	}
	if !strings.Contains(html, "Body content") {
		t.Error("expected body content")
	}
}
