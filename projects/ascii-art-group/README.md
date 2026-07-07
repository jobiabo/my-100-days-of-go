# ASCII Art Generator

A modular Command-Line Interface (CLI) application written in Go that converts text into stylized ASCII art using banner template files. The project focuses on clean code structure, input validation, and cross-platform compatibility.

---

## Key Features

* **Clean Project Structure:** Separates file loading, rendering, utility helpers, and application control into dedicated source files.
* **Flexible Banner Selection:** Supports banner names with or without the `.txt` extension.
* **Cross-Platform Compatibility:** Normalizes Windows (`\r\n`) and Unix (`\n`) line endings to ensure consistent rendering across operating systems.
* **Input Preprocessing:** Converts literal `\n` sequences into real line breaks and supports escaped tab formatting.
* **Banner Validation:** Verifies banner file structure before rendering to prevent malformed template errors.
* **Buffered File Scanning:** Implements a custom buffered scanner using os.read to efficiently and safely process large input files.

---


## Project Architecture

The project follows a modular structure that separates banner assets, utility helpers, and application logic into dedicated directories for better maintainability and scalability.

```text
ascii-art/
├── main.go                 # CLI entry point and application workflow
├── go.mod                  # Go module definition
├── README.md               # Project documentation
│
├── ascii/
│   ├── loadbanner.go       # Loads and validates banner template files
│   ├── render.go           # ASCII rendering engine
│   ├── utils.go            # Input preprocessing and helper functions
│   └── ascii_test.go       # Unit tests for the ascii package
│
└── banners/
    ├── standard.txt        # Default ASCII banner template
    ├── shawdow.txt         # Shadow-style ASCII banner template
    └── thinkertoy.txt      # Thinkertoy-style ASCII banner template
```

### Architecture Overview

* **main.go** handles CLI arguments, application flow, and rendering coordination.
* The **ascii/** package contains reusable helper utilities for preprocessing and validation.
* The **banners/** directory stores all ASCII template files used during rendering.
* Banner templates are loaded dynamically, making it easy to add new styles without changing the rendering logic.

```

---

## Data Flow

```text
[CLI Arguments]
        │
        ▼
 ValidateArgs()
        │
        ▼
 PrepareInput()
        │
        ▼
   LoadBanner()
        │
        ▼
     Render()
        │
        ▼
[Terminal Output]
```

---

## Usage

### Basic Usage

```bash
go run . "Hello"
```

### Using a Different Banner

```bash
go run . "Hello" shadow
```

### Multi-Line Rendering

```bash
go run . "Hello\nWorld"
```

## Supported Banner Files

* `standard`
* `shadow`
* `thinkertoy`

## Example

Input:

```bash
go run . "Go"
```

Output:

```text
  ____       
 / ___| ___  
| |  _ / _ \ 
| |_| | (_) |
 \____|\___/
```

---

## Technical Notes

* Banner templates are parsed into rune-to-pattern mappings for efficient rendering.
* Rendering is performed line-by-line using `strings.Builder` to reduce unnecessary string allocations.
* Scanner buffer limits are extended to avoid `bufio.Scanner: token too long` errors on larger files.

---

## Contributors

* **Elizabeth Odeh(eliodeh)** — Testing, debugging, and CLI argument validation.

* **Owoicho Charles Echumba(oechumba)** — Project structure, rendering engine, input preprocessing, and documentation.
* **Joel Obiabo(jobiabo)** — Banner parsing, validation logic, and file handling.




