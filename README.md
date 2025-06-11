# Go Concurrency 🧵

A collection of Go programs exploring and implementing core concurrency concepts including goroutines, channels, pipelines, fan-in/fan-out, generators, and more.

## 🚀 Overview

This repository is my hands-on journey to mastering concurrency in Go. Each folder or file demonstrates a specific concept or pattern with explanations and examples.

> 🔧 Language: [Go (Golang)](https://golang.org)
> 📚 Goal: Learn and apply concurrency patterns
> 🧠 Outcome: Build a strong foundation for writing efficient, concurrent Go applications

---

## 📂 Structure

| Topic              | Description                                       |
| ------------------ | ------------------------------------------------- |
| `goroutines/`      | Basics of goroutines and scheduling               |
| `channels/`        | Unbuffered and buffered channels                  |
| `select/`          | Multiplexing with `select`                        |
| `done-channel/`    | Graceful goroutine cancellation                   |
| `for-select-loop/` | Event loop style pattern                          |
| `pipeline/`        | Chained processing with stages                    |
| `generators/`      | Lazy data production                              |
| `fan-in-fan-out/`  | Parallel work distribution and result aggregation |
| `problems/`        | Custom problems to test concurrency skills        |

---

## 🧪 How to Run

Make sure you have Go installed (v1.18+ recommended).

```bash
go run <filename>.go
```
