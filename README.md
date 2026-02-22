# Atom Components

Reusable [Templ](https://templ.guide/) UI components for [Atom](https://github.com/AtomSites/atom-quickstart) projects. Modals, forms, cards, toasts — styled with glass-morphism and ready to drop in.

## Install

```bash
go get github.com/AtomSites/atom-components@latest
```

## Setup

### 1. Serve the component CSS and JS

Add to your `internal/web/server.go` after the existing `e.Static` line:

```go
import "github.com/AtomSites/atom-components/static"

e.GET("/static/css/atom-components.css", echo.WrapHandler(
    http.StripPrefix("/static/", http.FileServer(http.FS(static.Assets))),
))
e.GET("/static/js/atom-components.js", echo.WrapHandler(
    http.StripPrefix("/static/", http.FileServer(http.FS(static.Assets))),
))
```

### 2. Link the stylesheet and script

Add to your `internal/web/components/layout.templ`:

```html
<link rel="stylesheet" href="/static/css/atom-components.css"/>
<script src="/static/js/atom-components.js" defer></script>
```

## Components

### Modal

```go
import "github.com/AtomSites/atom-components/modal"

// Basic modal
@modal.Modal("confirm", "Are you sure?") {
    <p>This action cannot be undone.</p>
}

// Modal with footer buttons
@modal.ModalWithFooter("edit", "Edit Profile", myFooterComponent) {
    <p>Modal body content here.</p>
}
```

Open and close modals from JavaScript:

```js
acOpenModal("confirm");   // show modal
acCloseModal("confirm");  // hide modal
```

Modals also close when clicking the overlay background or pressing Escape.

### Toast

```go
import "github.com/AtomSites/atom-components/toast"

// Add the container once in your layout
@toast.Container()

// Render a static toast
@toast.Toast("Changes saved!", toast.Success)
```

Show toasts dynamically from JavaScript:

```js
acToast("File uploaded successfully", "success");
acToast("Something went wrong", "error");
acToast("Check your input", "warning");
acToast("New version available", "info");
```

Toast levels: `Success`, `Error`, `Warning`, `Info`.

### Form Inputs

```go
import "github.com/AtomSites/atom-components/form"

@form.TextInput("email", "email", "Email", "email", "you@example.com", "", "")
@form.TextArea("msg", "message", "Message", "Your message...", "", 5, "")
@form.Select("plan", "plan", "Plan", []form.SelectOption{
    {Value: "", Label: "Choose..."},
    {Value: "free", Label: "Free"},
    {Value: "pro", Label: "Pro", Selected: true},
}, "")
```

Pass an error message as the last argument to show validation errors:

```go
@form.TextInput("email", "email", "Email", "email", "", "bad", "Invalid email address")
```

### Contact Form

A complete contact form composed from the form primitives:

```go
import "github.com/AtomSites/atom-components/contact"

@contact.ContactForm("/contact", contact.FormData{})

// With prefilled data and validation errors
@contact.ContactForm("/contact", contact.FormData{
    Name:   "Alice",
    Email:  "alice@example.com",
    Errors: map[string]string{"email": "Invalid email"},
})
```

### Feature Card

```go
import "github.com/AtomSites/atom-components/card"

@card.FeatureCard("Fast", "Blazing fast performance") {
    // icon slot — put any HTML/SVG here
}
```

### Pricing Card

```go
@card.PricingCard(card.PricingTier{
    Name:        "Pro",
    Price:       "29",
    Currency:    "$",
    Period:      "/month",
    Features:    []string{"Unlimited projects", "Priority support", "Custom domains"},
    CTAText:     "Get Started",
    CTALink:     "/signup?plan=pro",
    Highlighted: true,
    Badge:       "Popular",
})
```

### Testimonial Card

```go
@card.TestimonialCard(
    "This product changed my workflow completely.",
    "Jane Doe",
    "CTO at TechCo",
    "https://example.com/avatar.jpg",
)
```

The avatar image is optional — pass an empty string to omit it.

## Packages

| Package | Import | Components |
|---|---|---|
| `modal` | `github.com/AtomSites/atom-components/modal` | `Modal`, `ModalWithFooter` |
| `toast` | `github.com/AtomSites/atom-components/toast` | `Container`, `Toast` |
| `form` | `github.com/AtomSites/atom-components/form` | `TextInput`, `TextArea`, `Select` |
| `contact` | `github.com/AtomSites/atom-components/contact` | `ContactForm` |
| `card` | `github.com/AtomSites/atom-components/card` | `FeatureCard`, `PricingCard`, `TestimonialCard` |
| `static` | `github.com/AtomSites/atom-components/static` | `Assets` (embedded CSS/JS) |

## CSS Variable Contract

All components reference CSS variables from the host project's `:root`. If you're using [atom-quickstart](https://github.com/AtomSites/atom-quickstart), these are already defined.

| Variable | Purpose |
|---|---|
| `--accent` | Primary brand color |
| `--accent-dark` | Darker brand shade |
| `--accent-light` | Lighter brand shade |
| `--bg-card` | Card background |
| `--bg-card-hover` | Card hover background |
| `--text-body` | Body text color |
| `--text-white` | Heading/emphasis text color |
| `--glass-bg` | Glass-morphism background |
| `--glass-bg-hover` | Glass-morphism hover background |
| `--glass-border` | Glass-morphism border |
| `--glass-border-hover` | Glass-morphism hover border |
| `--glass-blur` | Backdrop blur radius |
| `--border-subtle` | Subtle accent border |
| `--border-card` | Divider/card border |

Non-quickstart projects just need to define these variables in `:root` to use the components.

## CSS Class Prefix

All component classes use the `ac-` prefix to avoid collisions with the host project's styles.
