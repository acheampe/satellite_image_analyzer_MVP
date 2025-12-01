# 🌊 Satellite Image Analyzer (Go MVP)

A Go-based implementation of a simplified satellite image analyzer.  
This MVP focuses on grid-based algorithms such as water-body detection,  
flood-fill operations, border-connected expansion, and enclosed-lake identification.

The project is structured using **idiomatic Go module layout** and is written using a **test-driven development (TDD)** workflow.

---

## 🚀 Features

### 1. Count distinct water bodies  
Uses 8-directional adjacency (horizontal, vertical, diagonal).

### 2. Find the largest water body  
Returns the maximum connected area of water cells.

### 3. Detect enclosed lakes  
A water body is considered enclosed if it does **not** touch the border.

### 4. Simulate flood-risk expansion  
Starting from border-connected water bodies, expand outward using 8-directional adjacency.

---

## 🧱 Project Structure (Go Standard Layout)

```
satellite_image_analyzer/
│
├── cmd/
│   └── analyzer/            # CLI entry point (main.go)
│
├── internal/
│   └── grid/                # Core algorithm implementations
│       ├── connectivity.go
│       ├── floodfill.go
│       ├── count_water.go
│       ├── max_water.go
│       ├── enclosed_lakes.go
│       └── flood_risk.go
│
├── pkg/
│   └── gridtools/           # Public utilities
│
├── TDD/                   # Test Driven Development
│   ├── connectivity_test.go
│   ├── floodfill_test.go
│   ├── count_water_test.go
│   ├── enclosed_lakes_test.go
│   ├── max_water_test.go
│   └── flood_risk_test.go
│
├── go.mod
└── README.md
```

---

## 🧪 Test-Driven Development Workflow

> **Why a dedicated `TDD/` directory?**  

The `TDD/` folder contains the initial tests written for each feature *before* any 
production code exists.  
 
This emphasizes the project's TDD-driven workflow:  
1. Write a failing test  
2. Implement the minimal solution  
3. Refactor while keeping tests green  
 
Keeping TDD tests distinct highlights the iterative design process and demonstrates 
intentional, test-first engineering practices for portfolio reviewers.

Run tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test ./... -cover
```

---

## ▶️ Running the CLI

The CLI entry point will live at:

```
cmd/analyzer/main.go
```

Run it with:

```bash
go run ./cmd/analyzer
```

---

## 📦 Installation (For Contributors)

```bash
git clone https://github.com/acheampe/satellite_image_analyzer_MVP.git
cd satellite_image_analyzer_MVP
go mod tidy
```

---

## 📚 Future Enhancements

- Visualization tools (ASCII maps or web UI)  
- Parallel flood-fill using goroutines  
- Heatmap generation  
- REST API wrappers  
- S3/GCP integration for image ingestion  

---

## 👨‍💻 Author

Manny Acheampong  
Software Engineer + Physical Therapist
Focused on Go, TDD, and distributed systems in health-tech and simulation tooling.

---

## 📝 License

MIT License (or your preferred license)
