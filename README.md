# Go Practice Projects Collection

A collection of beginner-to-intermediate Go projects built while learning the Go programming language.  
This repository contains practical exercises focused on:

- User input handling
- Manual string processing
- Package creation and imports
- Error handling
- Pointer basics
- Random number generation
- CLI applications
- Mathematical calculations
- Structuring Go projects professionally

---

## 📁 Project Structure

.
├── bin/
│   └── executable-file
├── calculatepaintneeded/
│   └── calculatepaintneeded.go
├── greetings/
│   ├── greetings.go
│   └── deutsch/
│       └── deutsch.go
├── Keyboard/
│   └── keyboard.go
├── pass_fail/
│   └── pass_fail.go
├── Trimspace/
│   └── trimspace.go
├── main.go
└── README.md

---

## 🚀 Features

### ✅ Grade Checker
Checks whether a student passed or failed based on score input.

Example:
Type your score here: 75
passed

---

### ✅ Manual TrimSpace Function
A custom implementation of string trimming without using `strings.TrimSpace()`.

Handles:
- Spaces
- Tabs
- New lines
- Carriage returns

---

### ✅ Keyboard Input Package
Custom reusable package for:
- Reading user input
- Parsing float values safely
- Error handling

---

### ✅ Guessing Game
A command-line number guessing game using Go's random package.

Features:
- Random number generation
- Limited attempts
- High/low hints
- Success/failure tracking

---

### ✅ Paint Calculator
Calculates the amount of paint needed based on wall dimensions.

Formula used:

Paint Needed = Area / Paint Coverage

---

### ✅ Temperature Converter
Converts Fahrenheit to Celsius.

Example:
Enter temperature in Fahrenheit: 98
36.66667 celsius

---

### ✅ Table Formatting
Demonstrates formatted terminal output using `fmt.Printf`.

Example:

Product      | Amount in Cent
-------------------------------
Rice         | 60
Gun          | 90
Goat         | 890

---

### ✅ Pointer Practice
Demonstrates:
- Memory addresses
- Pointer referencing
- Dereferencing

---

### ✅ Custom Packages
This project demonstrates how to:
- Create reusable Go packages
- Import local packages
- Organize Go projects properly

---

## 🛠 Technologies Used

- Go (Golang)
- Standard Library Packages:
  - fmt
  - strconv
  - os
  - bufio
  - math/rand
  - reflect
  - log

---

## ▶️ How To Run

### Clone the Repository

```bash
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
go run .
```
📚 Future Improvements

Planned additions include:

ASCII art generator
File handling
HTTP servers
Web applications with Go
Goroutines and concurrency
Unit testing
Docker support
👨‍💻 Author

John Ayandipe

Aspiring Software Developer focused on:

Go (Golang)
Backend Engineering
Problem Solving
CLI and Web Applications

📄 License

This project is open-source and available under the MIT License.
