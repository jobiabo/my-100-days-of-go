# ASCII Banner Renderer (Go)

A simple Go program that converts text into ASCII art using a predefined banner font (`standard.txt`).

---

## ✨ Features

* Convert strings into ASCII art
* Supports full printable ASCII range (32–126)
* Input validation for safe rendering
* Clean modular structure (load, validate, render)

---

## 📦 Project Structure

```
ascii/
├── load.go        # Loads and parses banner file
├── render.go      # Renders ASCII art output
├── validate.go    # Validates input characters
main.go            # Entry point
standard.txt       # ASCII font file
```

---

## 🚀 How It Works

The program follows a simple pipeline:

```
LoadBanner → ValidateInput → RenderBanner → Print Output
```

Each character is mapped to an 8-line ASCII representation and combined horizontally to form the final output.

---

## 🧪 Example Usage

### Run the program:

```bash
go run . "hello"
```

### Output:

```
 _              _              _        
| |            | |            | |       
| |__     ___  | |__     ___  | |  ___  
|  _ \   / _ \ |  _ \   / _ \ | | / _ \ 
| | | | |  __/ | | | | |  __/ | || (_) |
|_| |_|  \___| |_| |_|  \___| |_| \___/ 
```

---

## ⚠️ Known Issue (Important Learning Point)

During development, a subtle issue appeared after refactoring the code into separate functions.

The ASCII banner file (`standard.txt`) contained **Windows-style line endings (`\r\n`)**, which introduced hidden carriage return characters (`\r`) into the program.

This caused unexpected terminal behavior where output appeared to overwrite itself, making it seem like only the last character was being rendered.

---

### Why it worked before refactoring

On some systems:

* The file used Unix-style line endings (`\n`)
* Or the issue was masked by less strict output handling

---

### Why it broke after refactoring

After introducing a clean pipeline:

```
LoadBanner → RenderBanner → Print Output
```

The program processed raw file data more strictly, exposing the hidden `\r` characters in the rendering stage.

---

## 🛠️ Fix

Normalize line endings when loading the banner file:

```go
clean := strings.ReplaceAll(string(data), "\r", "")
lines := strings.Split(clean, "\n")
```

This ensures consistent behavior across all platforms.

---

## 🧠 Key Lesson

This bug was not caused by refactoring itself.

Instead, refactoring made the program more deterministic and revealed a hidden data formatting issue.

> Clean architecture often exposes inconsistencies in external data sources.

---

## 📌 Summary

* ASCII rendering works as expected
* Issue was caused by `\r` carriage returns in banner file
* Refactoring exposed, not introduced, the bug
* Normalizing input data resolves cross-platform issues

---

If you want, I can also:

* add a **GIF/demo section**
* make it look like a **professional portfolio project**
* or convert it into a **42/Zone01-style README** (very common format there)
