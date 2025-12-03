# go_cliux

A simple Go library for enhancing command-line interface (CLI) user experience with styled boxes, colored labels, sections, and dividers.

## Features

- **Boxed**: Display text in a bordered box with a title.
- **Label**: Print colored labels for info, error, and success messages.
- **Section**: Create sections with headers and content.
- **Divider**: Add horizontal dividers to separate content.

## Installation

To install the library, use `go get`:

```bash
go get github.com/Pjdur/go_cliux
```

## Usage

Import the library in your Go code:

```go
import "github.com/Pjdur/go_cliux"
```

### Example

Here's a basic example demonstrating the functions:

```go
package main

import (
    "github.com/Pjdur/go_cliux"
)

func main() {
    go_cliux.Boxed("Google made Go", "Cliux", 40)
    go_cliux.Label("Cliux", "success")
    go_cliux.Label("Cliux", "error")
    go_cliux.Label("Cliux", "info")
    go_cliux.Section("Cliux", "Go is made by Google", 30)
    go_cliux.Divider("0", 25)
}
```

This will output styled CLI elements.

## API Documentation

### Boxed(text string, name string, width int)

Displays the given `text` in a bordered box with the specified `name` as the title. The `width` parameter controls the width of the box.

- `text`: The content to display inside the box.
- `name`: The title of the box.
- `width`: The width of the box in characters.

### Label(content string, style string)

Prints the `content` with color based on the `style`.

- `content`: The text to label.
- `style`: The style of the label. Supported values: "info" (blue), "error" (red), "success" (green).

### Section(content string, name string, width int)

Creates a section with a header `name` followed by the `content`, separated by a divider of the specified `width`.

- `content`: The main content of the section.
- `name`: The header of the section.
- `width`: The width of the divider in characters.

### Divider(style string, width int)

Prints a horizontal divider using the specified `style` character, repeated for the given `width`.

- `style`: The character to use for the divider (e.g., "-", "0").
- `width`: The length of the divider in characters.

## Dependencies

- [github.com/fatih/color](https://github.com/fatih/color) for colored output.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
