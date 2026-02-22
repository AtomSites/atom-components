# Contributing to Atom Components

## Dev Setup

```bash
make install    # install Go deps + templ + golangci-lint
make generate   # generate templ files
make test       # generate templ + run all tests
make lint       # run golangci-lint
```

## Adding a Component

1. **Create a package** — add a `componentname/` directory with `componentname.templ` inside it.

2. **Add CSS** — add styles to `static/css/atom-components.css`. All classes must use the `ac-` prefix.

3. **Add JS (if needed)** — add to `static/js/atom-components.js`. Use `data-*` attributes for targeting. No framework dependencies.

4. **Write tests** — add `componentname_test.go` in the same package directory. Use the `componentname_test` package. Render the component to a buffer and assert on the HTML output.

5. **Run checks** — `make test && make lint` must pass.

## CSS Rules

- All classes **must** use the `ac-` prefix (e.g., `ac-my-widget`)
- Only reference CSS variables — no hardcoded colors. Use `var(--accent)`, `var(--glass-bg)`, etc.
- Follow glass-morphism patterns: `backdrop-filter: blur(var(--glass-blur))`, glass borders, etc.
- Add responsive styles at 768px and 480px breakpoints
- Keep specificity low — avoid nesting beyond 2 levels

## JS Rules

- Add to `static/js/atom-components.js` inside the IIFE
- Use `data-*` attributes for targeting elements (e.g., `data-modal-close`)
- Use event delegation on `document` — works with dynamically rendered content
- No framework dependencies — vanilla JS only
- Expose global helpers via `window.acFunctionName`

## Testing

Tests render components to a buffer and assert on the HTML structure:

```go
func TestMyComponent(t *testing.T) {
    var buf bytes.Buffer
    err := mypackage.MyComponent("arg").Render(context.Background(), &buf)
    if err != nil {
        t.Fatalf("render error: %v", err)
    }
    html := buf.String()

    if !strings.Contains(html, "ac-my-component") {
        t.Error("expected ac-my-component class")
    }
}
```

For components with `{ children... }`, use `templ.WithChildren`:

```go
component := mypackage.MyComponent("arg")
err := component.Render(templ.WithChildren(context.Background(), templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
    _, e := io.WriteString(w, "<p>Child content</p>")
    return e
})), &buf)
```

## Releasing

After merge to main, tag with semver:

```bash
git tag v0.X.0
git push --tags
```
