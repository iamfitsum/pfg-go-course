# 🚀 PFG Go Course

This repository contains the code written while following the [Introduction to Go for Beginners](https://pfglabs.com/) course from PFG Labs. Each chapter demonstrates fundamental Go concepts with practical examples and concludes with a capstone project integrating everything learned in the course.

## 📖 Table of Contents
- [Repository Structure](#%EF%B8%8F-repository-structure)
- [Chapters Overview](#%EF%B8%8F-chapters-overview)
  - [Main Program](#main-program)
  - [Chapter 1: Variables](#chapter-1-variables)
  - [Chapter 2: Slices](#chapter-2-slices)
  - [Chapter 3: Maps](#chapter-3-maps)
  - [Chapter 4: Structs](#chapter-4-structs)
  - [Chapter 5: Interfaces](#chapter-5-interfaces)
  - [Chapter 6: Capstone Project](#chapter-6-capstone-project)
- [Capstone Project: CLI Checkers Game](#-capstone-project-cli-checkers-game)
- [Getting Started](#-getting-started)
  - [Prerequisites](#prerequisites)
  - [Running Code for Each Chapter](#running-code-for-each-chapter)
  - [Running the Capstone Project](#running-the-capstone-project)
- [About the Course](#-about-the-course)
- [Acknowledgements](#-acknowledgements)

## 🗂️ Repository Structure

```
.
├── go.mod          # Dependency management
├── go.sum          # Dependency checksums
├── [interfaces/](./interfaces/)     # Chapter 5: Interfaces
│   └── main.go
├── [main.go](./main.go)         # Initial "Hello World" example
├── [maps/](./maps/)           # Chapter 3: Maps
│   └── main.go
├── [project/](./project/)        # Chapter 6: Capstone project - CLI Checkers game
│   └── main.go
├── [slices/](./slices/)         # Chapter 2: Slices
│   └── main.go
├── [structs/](./structs/)        # Chapter 4: Structs
│   └── main.go
└── [variables/](./variables/)      # Chapter 1: Variables
    └── main.go
```

## 🛠️ Chapters Overview

### **Main Program**
- **File:** `main.go`  
- **Description:** Prints a personalized "Hello" message to the console using a function. This serves as the foundation for understanding Go basics like packages, imports, and functions.

### **Chapter 1: Variables**
- **File:** `variables/main.go`  
- **Key Concepts:**  
  - Declaring and initializing variables using `var`.  
  - String concatenation with the `+` operator.  
  - Basic integer arithmetic (increment operation).

### **Chapter 2: Slices**
- **File:** `slices/main.go`  
- **Key Concepts:**  
  - Creating and modifying slices.  
  - Implementing slice deletion using the `slices.Delete` package.  
  - Exploring dynamic behavior of slices in Go.

### **Chapter 3: Maps**
- **File:** `maps/main.go`  
- **Key Concepts:**  
  - Creating and managing maps to store key-value pairs.  
  - Adding and deleting elements in a map.  
  - Practical use cases for maps in Go.

### **Chapter 4: Structs**
- **File:** `structs/main.go`  
- **Key Concepts:**  
  - Defining and using structs to group related data.  
  - Updating struct fields by value and by reference.  
  - Adding methods to manipulate struct instances.

### **Chapter 5: Interfaces**
- **File:** `interfaces/main.go`  
- **Key Concepts:**  
  - Implementing interfaces to define shared behavior across types.  
  - Using `Speak` methods to demonstrate polymorphism with different types (Person, Robot, Animal).  
  - Iterating through a collection of interface types.

### **Chapter 6: Capstone Project**
- **File:** `project/main.go`  
- **Key Concepts:**  
  - Building a command-line Checkers game.  
  - Combining structs, slices, maps, and interfaces into a cohesive project.  
  - Implementing game logic, board visualization, and move validation.

## 🎮 Capstone Project: CLI Checkers Game

The capstone project integrates concepts from all chapters to create a simple Checkers game playable via the command line.  
- **Features:**  
  - Diagonal movement and capture rules for Checkers.  
  - Board visualization in the terminal.  
  - Turn-based gameplay and move validation logic.  
- **Code Highlights:**  
  - Structs for `Piece`, `Position`, `Move`, and `Board`.  
  - Validation logic for moves, including captures and bounds checking.  
  - Use of constants, loops, and conditionals to implement game rules.

## 🏁 Getting Started

### Prerequisites
- Install [Go](https://go.dev/doc/install).

### Running Code for Each Chapter
1. Navigate to the respective directory (e.g., `cd variables`).  
2. Run the code using the Go command:  
   ```bash
   go run main.go
   ```

### Running the Capstone Project
1. Navigate to the `project` directory:  
   ```bash
   cd project
   ```
2. Run the Checkers game:  
   ```bash
   go run main.go
   ```

## 📚 About the Course

This repository is a hands-on implementation of the foundational concepts taught in the [PFG Labs](https://pfglabs.com/) Go course. It is structured to reinforce understanding of Go programming through progressively complex examples, culminating in a complete CLI-based project. 

What makes this course unique:
- **Beginner-Friendly:** Designed for those new to programming or Go.
- **Practical Approach:** Learn through real-world examples, not just theory.
- **Capstone Project:** A fully integrated CLI Checkers game to apply all concepts learned.

## 🙏 Acknowledgements

Special thanks to [PFG Labs](https://pfglabs.com/) and [MelkyDev](https://github.com/Melkeydev) for creating this excellent course.
