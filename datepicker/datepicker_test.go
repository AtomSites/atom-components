package datepicker_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/AtomSites/atom-components/datepicker"
)

func TestDatePickerBasic(t *testing.T) {
	var buf bytes.Buffer
	cfg := datepicker.DatePickerConfig{
		ID:    "dob",
		Name:  "date_of_birth",
		Label: "Date of Birth",
	}
	err := datepicker.DatePicker(cfg).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `id="dob"`) {
		t.Error("expected root id")
	}
	if !strings.Contains(html, `data-ac-datepicker`) {
		t.Error("expected data-ac-datepicker attribute")
	}
	if !strings.Contains(html, "Date of Birth") {
		t.Error("expected label text")
	}
	if !strings.Contains(html, "ac-datepicker-trigger") {
		t.Error("expected trigger input class")
	}
	if !strings.Contains(html, `name="date_of_birth"`) {
		t.Error("expected hidden input name")
	}
	if !strings.Contains(html, "ac-datepicker-modal") {
		t.Error("expected modal structure")
	}
	if !strings.Contains(html, "data-ac-datepicker-confirm") {
		t.Error("expected confirm button")
	}
	if !strings.Contains(html, "data-ac-datepicker-cancel") {
		t.Error("expected cancel button")
	}
	if !strings.Contains(html, "data-ac-datepicker-today") {
		t.Error("expected today button")
	}
}

func TestDatePickerPreSelectedValue(t *testing.T) {
	var buf bytes.Buffer
	cfg := datepicker.DatePickerConfig{
		ID:    "start",
		Name:  "start_date",
		Label: "Start Date",
		Value: "2026-01-15",
	}
	err := datepicker.DatePicker(cfg).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `value="2026-01-15"`) {
		t.Error("expected ISO value in hidden input")
	}
	if !strings.Contains(html, "Jan 15, 2026") {
		t.Error("expected formatted display date in trigger")
	}
}

func TestDatePickerErrorState(t *testing.T) {
	var buf bytes.Buffer
	cfg := datepicker.DatePickerConfig{
		ID:     "event",
		Name:   "event_date",
		Label:  "Event Date",
		ErrMsg: "Date is required",
	}
	err := datepicker.DatePicker(cfg).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "ac-input-error") {
		t.Error("expected ac-input-error class on trigger")
	}
	if !strings.Contains(html, "ac-error-text") {
		t.Error("expected error text element")
	}
	if !strings.Contains(html, "Date is required") {
		t.Error("expected error message text")
	}
}

func TestDatePickerNoError(t *testing.T) {
	var buf bytes.Buffer
	cfg := datepicker.DatePickerConfig{
		ID:   "clean",
		Name: "clean_date",
	}
	err := datepicker.DatePicker(cfg).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if strings.Contains(html, "ac-input-error") {
		t.Error("should not have error class when no error")
	}
	if strings.Contains(html, "ac-error-text") {
		t.Error("should not render error text when no error")
	}
}

func TestDatePickerCustomPlaceholder(t *testing.T) {
	var buf bytes.Buffer
	cfg := datepicker.DatePickerConfig{
		ID:          "custom",
		Name:        "custom_date",
		Placeholder: "Pick a date",
	}
	err := datepicker.DatePicker(cfg).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "Pick a date") {
		t.Error("expected custom placeholder")
	}
}

func TestDatePickerDefaultPlaceholder(t *testing.T) {
	var buf bytes.Buffer
	cfg := datepicker.DatePickerConfig{
		ID:   "default",
		Name: "default_date",
	}
	err := datepicker.DatePicker(cfg).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "Select a date...") {
		t.Error("expected default placeholder")
	}
}

func TestDatePickerYearRange(t *testing.T) {
	var buf bytes.Buffer
	cfg := datepicker.DatePickerConfig{
		ID:      "range",
		Name:    "range_date",
		MinYear: 2000,
		MaxYear: 2030,
	}
	err := datepicker.DatePicker(cfg).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `data-ac-datepicker-min-year="2000"`) {
		t.Error("expected min year data attribute")
	}
	if !strings.Contains(html, `data-ac-datepicker-max-year="2030"`) {
		t.Error("expected max year data attribute")
	}
}

func TestDatePickerAccessibility(t *testing.T) {
	var buf bytes.Buffer
	cfg := datepicker.DatePickerConfig{
		ID:    "a11y",
		Name:  "a11y_date",
		Label: "Appointment",
	}
	err := datepicker.DatePicker(cfg).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, `aria-haspopup="dialog"`) {
		t.Error("expected aria-haspopup on trigger")
	}
	if !strings.Contains(html, `aria-expanded="false"`) {
		t.Error("expected aria-expanded on trigger")
	}
	if !strings.Contains(html, `aria-controls="a11y-modal"`) {
		t.Error("expected aria-controls on trigger")
	}
	if !strings.Contains(html, `aria-label="Appointment date picker"`) {
		t.Error("expected aria-label on modal")
	}
}

func TestDatePickerNoLabel(t *testing.T) {
	var buf bytes.Buffer
	cfg := datepicker.DatePickerConfig{
		ID:   "nolabel",
		Name: "nolabel_date",
	}
	err := datepicker.DatePicker(cfg).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if strings.Contains(html, "ac-label") {
		t.Error("should not render label when Label is empty")
	}
}
