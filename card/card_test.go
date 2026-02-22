package card_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"

	"github.com/AtomSites/atom-components/card"
)

func TestFeatureCard(t *testing.T) {
	var buf bytes.Buffer
	c := card.FeatureCard("Fast", "Blazing fast performance")
	err := c.Render(templ.WithChildren(context.Background(), templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, e := io.WriteString(w, "⚡")
		return e
	})), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "ac-feature-card") {
		t.Error("expected ac-feature-card class")
	}
	if !strings.Contains(html, "Fast") {
		t.Error("expected card title")
	}
	if !strings.Contains(html, "Blazing fast performance") {
		t.Error("expected card description")
	}
	if !strings.Contains(html, "ac-feature-card-icon") {
		t.Error("expected icon slot")
	}
}

func TestPricingCard(t *testing.T) {
	tier := card.PricingTier{
		Name:        "Pro",
		Price:       "29",
		Currency:    "$",
		Period:      "/month",
		Features:    []string{"Unlimited projects", "Priority support", "Custom domains"},
		CTAText:     "Get Started",
		CTALink:     "/signup?plan=pro",
		Highlighted: true,
		Badge:       "Popular",
	}
	var buf bytes.Buffer
	err := card.PricingCard(tier).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "ac-pricing-card") {
		t.Error("expected ac-pricing-card class")
	}
	if !strings.Contains(html, "ac-pricing-card-highlighted") {
		t.Error("expected highlighted class")
	}
	if !strings.Contains(html, "Popular") {
		t.Error("expected badge text")
	}
	if !strings.Contains(html, "Pro") {
		t.Error("expected tier name")
	}
	if !strings.Contains(html, "29") {
		t.Error("expected price")
	}
	if !strings.Contains(html, "/month") {
		t.Error("expected period")
	}
	if !strings.Contains(html, "Unlimited projects") {
		t.Error("expected feature")
	}
	if !strings.Contains(html, "Get Started") {
		t.Error("expected CTA text")
	}
}

func TestPricingCardNotHighlighted(t *testing.T) {
	tier := card.PricingTier{
		Name:     "Free",
		Price:    "0",
		Currency: "$",
		Period:   "/month",
		Features: []string{"1 project"},
		CTAText:  "Start Free",
		CTALink:  "/signup",
	}
	var buf bytes.Buffer
	err := card.PricingCard(tier).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if strings.Contains(html, "ac-pricing-card-highlighted") {
		t.Error("should not have highlighted class")
	}
	if strings.Contains(html, "ac-pricing-badge") {
		t.Error("should not have badge")
	}
}

func TestTestimonialCard(t *testing.T) {
	var buf bytes.Buffer
	err := card.TestimonialCard(
		"This product changed my workflow completely.",
		"Jane Doe",
		"CTO at TechCo",
		"https://example.com/avatar.jpg",
	).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if !strings.Contains(html, "ac-testimonial-card") {
		t.Error("expected ac-testimonial-card class")
	}
	if !strings.Contains(html, "This product changed my workflow completely.") {
		t.Error("expected quote text")
	}
	if !strings.Contains(html, "Jane Doe") {
		t.Error("expected author name")
	}
	if !strings.Contains(html, "CTO at TechCo") {
		t.Error("expected author role")
	}
	if !strings.Contains(html, "ac-testimonial-avatar") {
		t.Error("expected avatar image")
	}
}

func TestTestimonialCardNoAvatar(t *testing.T) {
	var buf bytes.Buffer
	err := card.TestimonialCard(
		"Great tool!",
		"John",
		"Developer",
		"",
	).Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	html := buf.String()

	if strings.Contains(html, "ac-testimonial-avatar") {
		t.Error("should not render avatar when URL is empty")
	}
}
