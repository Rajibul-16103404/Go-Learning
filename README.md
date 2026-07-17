# Go Learning Journey 🐹

Welcome to my Go (Golang) practice and learning repository! I am using this space to write, test, and save all my practice codes as I follow along with the tutorial series.

## 📺 Tutorial Source
* **Channel:** Go With Habib
* **Video:** [[Golang] 010 - Introduction to Functions](https://youtu.be/UY439sqE1tM)

---

## 📂 Repository Structure

The project directory is structured as follows:

```text
Go Learning/
│
├── Hello world/
│   └── main.go           # Basic Hello World setup & standard entry point
│
└── Conditions/
    └── switch/
        └── switch.go     # Practice files covering conditional switch logic
```

---

## 🛠️ Environment Setup & Quick Reference

### 1. Initializing Git Identity
Before making your first commit, tell Git who you are:
```powershell
git config --global user.email "your.email@example.com"
git config --global user.name "Your Name"
```

### 2. Basic Go Commands
To run a Go file directly without building a binary:
```powershell
go run main.go
```

---

## 🚀 How to Sync Changes to GitHub

Whenever you complete a code example, follow these commands to push it to your repository:

### For the very first push (Initial Setup):
```powershell
# Navigate to your project root or specific package folder
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin https://github.com/Rajibul-16103404/Go-Learning.git
git push -u origin main
```

### For subsequent updates:
```powershell
git add .
git commit -m "Add practice code for functions/conditions"
git push
```

---

## 📝 Concepts Practiced (From Video 010)
* Understanding how Go manages variables inside RAM during execution.
* The basic anatomy of a function using the `func` keyword.
* Defining parameter types explicitly:
  ```go
  func add(number1 int, number2 int) { ... }
  ```
* Scope and lifecycles of variable stacks within a custom function frame.

Happy Coding! 🚀
