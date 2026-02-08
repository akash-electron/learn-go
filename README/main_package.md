# 👑 The "Boss" Package: Understanding `package main`

In Go, every file belongs to a family (package). But `package main` is the **Boss**. Without it, your code is just a bunch of papers lying around—nothing actually happens!

---

## 1. What is `package main`?

Think of it like the captain of a ship. You can have many sailors (other packages), but if there’s no captain (`package main`), the ship doesn't move.

- If you name a package `main`, you are telling Go: **"This is a program I want to RUN."**
- If you name it something else (like `package hay`), Go thinks: **"This is just a tool for others to use later."**

## 2. The `main()` Function: The Spark

Inside the Boss package, you MUST have a `main()` function. This is the exact spot where the computer starts reading.

```go
package main // The Boss Package

import "fmt"

func main() {
    fmt.Println("The donkey starts walking now!")
}
```

_If you don't have this, Go goes: "I don't know where to start!" and sits down._

---

## 3. The "Main" Family (Multiple Files)

You can have many files in the same folder all saying `package main`. They are all part of the same family!

Imagine two files in one folder:

- `main.go`
- `helper.go`

If both say `package main`, they can see each other's secrets without even asking!

**To run them together, you say:**

```bash
go run .
```

_(The dot means "Run everything in this family folder!")_

---

## 4. Executable vs. Library (The "Hee-Haw" Test)

| Type        | Package Name         | Does it have `main()`? | Result                                      |
| :---------- | :------------------- | :--------------------- | :------------------------------------------ |
| **Program** | `main`               | ✅ Yes                 | Creates a file you can run (.exe or binary) |
| **Library** | `util`, `math`, etc. | ❌ No                  | Just a helper for other programs            |

---

## 5. One Boss per Folder

**Crucial Donkey Rule:** You cannot have two different "Bosses" in one folder.
If you have two files that both have `func main()` in the same folder, Go will get confused and kick!

Each project or tool should usually have its own folder if it needs its own `main()` function.

---

_Now you know who's the boss!_ 🫏👑
