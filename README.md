# Project Description

This project is a simple Thai Baht text converter implemented in Go.  
It converts decimal numbers into Thai text representation with proper Baht and Satang currency formatting.

## Technology Stack

| Component    | Technology      |
| ------------ | --------------- |
| **Language** | Go 1.25.6       |
| **Testing**  | testify v1.11.1 |

## Project Structure

```
.
├── decimal
│   ├── decimal_test.go
│   └── decimal.go
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md

2 directories, 7 files
```

# Getting Started

## Prerequisites

- Go 1.25.6 or later

---

# Installation

Clone the repository:

```bash
git clone <repository-url>
cd <repository-name>
```

Install dependencies:

```bash
go mod tidy
```

---

# Run the Project

Using Makefile:

```bash
make run
```

Or directly with Go:

```bash
go run main.go
```

---

# Run Tests

Using Makefile:

```bash
make test
```

Or directly with Go:

```bash
go test ./... -cover
```

---
