# Atom Components

Reusable Templ UI components for Atom projects.

## Tech Stack

- **Language**: Go 1.25 with Templ
- **Templates**: Templ (`.templ` files in per-component sub-packages)
- **Styling**: Single CSS file at `static/css/atom-components.css`, embedded via `go:embed`
- **JavaScript**: Minimal vanilla JS at `static/js/atom-components.js`, embedded alongside CSS

## Project Structure

```
go.mod                           — module github.com/AtomSites/atom-components
static/
  static.go                      — //go:embed directive exposing embed.FS
  css/
    atom-components.css          — all component styles (ac-* prefixed)
  js/
    atom-components.js           — minimal vanilla JS (modal close, toast dismiss)
modal/
  modal.templ                    — Modal, ModalWithFooter
  modal_test.go
toast/
  toast.templ                    — Container, Toast
  toast_test.go
form/
  form.templ                     — TextInput, TextArea, Select
  helpers.go                     — intToString helper
  form_test.go
contact/
  contact.templ                  — ContactForm (composed from form primitives)
  contact_test.go
card/
  card.templ                     — FeatureCard, PricingCard, TestimonialCard
  card_test.go
```

## Dev Workflow

```bash
make install    # install Go deps + templ + golangci-lint
make generate   # generate templ files
make test       # generate templ + run all tests
make lint       # run golangci-lint
```

## CSS Rules

- All classes MUST use `ac-` prefix to avoid collisions with consuming projects
- Only reference CSS variables from the host project (--accent, --glass-bg, etc.) — no hardcoded colors
- Follow glass-morphism patterns matching atom-quickstart's design
- Add responsive styles at 768px and 480px breakpoints

## JS Rules

- Add to `static/js/atom-components.js`
- Use `data-*` attributes for targeting elements
- No framework dependencies — vanilla JS only
- Use event delegation on `document`

## Adding a Component

1. Create `componentname/componentname.templ` as a new sub-package
2. Add CSS to `static/css/atom-components.css` with `ac-` prefix
3. If JS needed, add to `static/js/atom-components.js` with `data-*` targeting
4. Add tests in `componentname/componentname_test.go`
