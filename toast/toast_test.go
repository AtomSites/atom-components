package toast_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/AtomSites/atom-components/toast"
)

func TestContainer(t *testing.T) {
	var buf bytes.Buffer
	err := toast.Container().Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "ac-toast-container") {
		t.Error("expected ac-toast-container class")
	}
}

func TestToast(t *testing.T) {
	levels := []toast.Level{toast.Success, toast.Error, toast.Warning, toast.Info}
	for _, level := range levels {
		t.Run(string(level), func(t *testing.T) {
			var buf bytes.Buffer
			err := toast.Toast("Something happened", level).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render error: %v", err)
			}
			html := buf.String()

			if !strings.Contains(html, "ac-toast") {
				t.Error("expected ac-toast class")
			}
			if !strings.Contains(html, "ac-toast-"+string(level)) {
				t.Errorf("expected ac-toast-%s class", level)
			}
			if !strings.Contains(html, "Something happened") {
				t.Error("expected toast message")
			}
			if !strings.Contains(html, `role="alert"`) {
				t.Error("expected alert role")
			}
			if !strings.Contains(html, "data-toast-close") {
				t.Error("expected close button with data-toast-close")
			}
		})
	}
}
