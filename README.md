# 🔍 Go Log Analyzer CLI

A production-style log analyzer CLI written in Go to demonstrate clean architecture, testing best practices, and performance optimization techniques used by professional Go engineers.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [What This Tool Does](#what-this-tool-does)
- [Why This Project Matters](#why-this-project-matters)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Benchmarking](#benchmarking)
- [Profiling](#profiling)
- [Example Output](#example-output)
- [Concepts Demonstrated](#concepts-demonstrated)
- [Learning Outcomes](#learning-outcomes)
- [Technologies Used](#technologies-used)
- [Contributing](#contributing)
- [License](#license)
- [Author](#author)

## 🎯 Overview

This project is a **learning milestone** designed to deeply understand how real Go engineers write, test, benchmark, and optimize production-grade code. It's not a toy project—it mimics real backend/DevOps/SRE tooling used in production environments.

## ✨ Features

- **Clean Architecture**: Follows industry-standard Go project layout
- **Comprehensive Testing**: Unit tests with table-driven test patterns
- **Performance Benchmarking**: Measure execution time and memory usage
- **CPU & Memory Profiling**: Identify bottlenecks using `pprof`
- **Production-Ready**: Demonstrates real-world coding practices

## 📌 What This Tool Does

This CLI reads a log file and generates a comprehensive analysis report with:

- ✅ **Total requests** processed
- ❌ **Total error count** (HTTP status ≥ 500)
- 🌐 **Most frequent IP address**
- 🔗 **Most requested endpoint**
- ⏱️ **Average response time** in milliseconds

### Log Format

Each log line follows this format:

```
192.168.0.1 GET /api/users 200 120
```

**Format breakdown:**

```
<IP_ADDRESS> <HTTP_METHOD> <ENDPOINT> <STATUS_CODE> <RESPONSE_TIME_MS>
```

### Example Log File

```
192.168.0.1 GET /api/users 200 120
192.168.0.2 POST /api/login 200 95
192.168.0.1 GET /api/products 200 200
192.168.0.3 GET /api/users 500 150
192.168.0.1 DELETE /api/products/5 204 80
```

## 🧠 Why This Project Matters

This project forces you to use:

| Aspect                       | Implementation                             |
| ---------------------------- | ------------------------------------------ |
| **Multiple Go packages**     | Clean separation of concerns               |
| **Proper folder structure**  | Follows `cmd/` and `internal/` conventions |
| **Testing best practices**   | Table-driven tests, edge cases             |
| **Performance measurement**  | Benchmarking critical paths                |
| **Profiling & optimization** | Using `pprof` for bottleneck analysis      |

This mirrors **real-world backend/DevOps/SRE tooling** and prepares you for production Go development.

## 📁 Project Structure

```
go-log-analyzer/
├── cmd/
│   └── analyzer/
│       └── main.go          # CLI entry point
├── internal/
│   ├── parser/
│   │   ├── parser.go        # Log parsing logic
│   │   └── parser_test.go   # Parser unit tests
│   └── analyzer/
│       ├── analyzer.go      # Log analysis logic
│       └── analyzer_test.go # Analyzer unit tests
├── testdata/
│   └── sample.log           # Sample log file for testing
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
├── README.md                # This file
└── LICENSE                  # MIT License

```

**Why this structure?**

- `cmd/analyzer/`: Entry point for the CLI application
- `internal/`: Private packages that can't be imported by other projects
- `testdata/`: Test fixtures and sample data
- Follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

## 🚀 Installation

### Prerequisites

- Go 1.21 or higher
- Git

### Clone the Repository

```bash
git clone https://github.com/yourusername/go-log-analyzer.git
cd go-log-analyzer
```

### Install Dependencies

```bash
go mod download
```

### Build the Binary (Optional)

```bash
go build -o analyzer cmd/analyzer/main.go
```

## 💻 Usage

### Run Directly

```bash
go run cmd/analyzer/main.go
```

### Run Compiled Binary

```bash
./analyzer
```

### Analyze a Custom Log File

```bash
go run cmd/analyzer/main.go -file=/path/to/your/logfile.log
```

## 🧪 Testing

### Run All Tests

```bash
go test ./...
```

### Run Tests with Verbose Output

```bash
go test -v ./...
```

### Run Tests with Coverage

```bash
go test -cover ./...
```

### Generate Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### What's Being Tested?

- ✅ **Unit tests** for individual functions
- ✅ **Table-driven tests** for edge cases
- ✅ **Error handling** validation
- ✅ **Boundary conditions** testing

## 🏎️ Benchmarking

### Run All Benchmarks

```bash
go test -bench=. -benchmem ./...
```

### Sample Output

```
BenchmarkParseLine-8        5000000    250 ns/op    128 B/op    3 allocs/op
BenchmarkAnalyze-8           100000   1200 ns/op   1024 B/op   12 allocs/op
```

**Understanding the output:**

- `5000000`: Number of iterations
- `250 ns/op`: Average time per operation (nanoseconds)
- `128 B/op`: Memory allocated per operation (bytes)
- `3 allocs/op`: Number of allocations per operation

### Benchmark Specific Functions

```bash
go test -bench=BenchmarkParseLine -benchmem ./internal/parser/
```

## 🔥 Profiling

### CPU Profiling

#### 1. Enable CPU Profiling in Code

The application automatically generates `cpu.prof` when run.

#### 2. Run the Application

```bash
go run cmd/analyzer/main.go
```

#### 3. Analyze the Profile

```bash
go tool pprof cpu.prof
```

#### 4. Inside pprof Interactive Mode

```bash
(pprof) top
(pprof) top10
(pprof) list <function_name>
(pprof) web
```

**Commands:**

- `top`: Show top CPU consumers
- `list`: Show source code for a function
- `web`: Generate a graph visualization (requires Graphviz)

### Memory Profiling

#### 1. Run with Memory Profiling

```bash
go run cmd/analyzer/main.go
```

The application generates `mem.prof` automatically.

#### 2. Analyze Memory Usage

```bash
go tool pprof mem.prof
```

#### 3. Find Memory Hotspots

```bash
(pprof) top
(pprof) list <function_name>
```

### Installing Graphviz (for web visualization)

**macOS:**

```bash
brew install graphviz
```

**Ubuntu/Debian:**

```bash
sudo apt-get install graphviz
```

**Windows:**
Download from [graphviz.org](https://graphviz.org/download/)

## 📊 Example Output

```
===== Log Analysis Report =====
Total Requests:        5
Total Errors:          1
Average Response Time: 154 ms
Top IP Address:        192.168.0.1 (3 requests)
Top Endpoint:          /api/users (2 requests)
================================
```

## 🎓 Concepts Demonstrated

| Concept                | Implementation               | Location                 |
| ---------------------- | ---------------------------- | ------------------------ |
| **Go Modules**         | Dependency management        | `go.mod`, `go.sum`       |
| **Project Structure**  | Clean architecture           | `cmd/`, `internal/`      |
| **Unit Testing**       | Testing package              | `*_test.go` files        |
| **Table-Driven Tests** | Structured test cases        | `parser_test.go`         |
| **Benchmarking**       | Performance measurement      | `Benchmark*` functions   |
| **CPU Profiling**      | Runtime performance analysis | `pprof` CPU profiling    |
| **Memory Profiling**   | Memory usage analysis        | `pprof` memory profiling |
| **Error Handling**     | Idiomatic Go error handling  | Throughout codebase      |
| **Interfaces**         | Abstraction and testability  | `internal/` packages     |
| **Documentation**      | Code comments and README     | All `.go` files          |

## 🎯 Learning Outcomes

After building and studying this project, you will understand:

### 1. Professional Go Project Structure

- How to organize code using `cmd/` and `internal/`
- When to use packages vs. single files
- How to structure testable code

### 2. Testing in Go

- Writing unit tests with the `testing` package
- Table-driven test patterns
- Achieving high test coverage
- Testing error conditions

### 3. Performance Optimization

- Measuring performance with benchmarks
- Interpreting benchmark results
- Understanding memory allocations
- Identifying performance bottlenecks

### 4. Profiling and Debugging

- Using `pprof` for CPU profiling
- Analyzing memory consumption
- Reading flame graphs
- Optimizing hot paths

### 5. Production-Ready Code

- Error handling patterns
- Code documentation
- Module management with `go.mod`
- Building maintainable software

## 🛠️ Technologies Used

- **[Go](https://golang.org/)** - Primary programming language
- **[testing](https://pkg.go.dev/testing)** - Built-in Go testing framework
- **[pprof](https://pkg.go.dev/runtime/pprof)** - Performance profiling tool
- **Go Modules** - Dependency management
- **Go Benchmarking** - Performance measurement

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

Please ensure:

- ✅ All tests pass (`go test ./...`)
- ✅ Code is formatted (`go fmt ./...`)
- ✅ Code is linted (`golangci-lint run`)
- ✅ Benchmarks show no regression

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 💬 Author

**Saad Islam**

- Backend Developer
- Learning Go for Systems & Performance Engineering
- GitHub: [@yourusername](https://github.com/yourusername)
- LinkedIn: [Your LinkedIn](https://linkedin.com/in/yourprofile)

## ⭐ Star This Repository

If you found this project useful for learning Go, please give it a star! ⭐

It helps others discover this resource and motivates continued development.

## 🙏 Acknowledgments

- Go community for excellent documentation
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://go.dev/doc/effective_go)
- All contributors and learners

---

**Happy Learning! 🚀**

Built with ❤️ for the Go community
