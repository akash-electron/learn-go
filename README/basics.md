# 🫏 Go for Donkeys: The "I Don't Know Anything" Guide

Hello friend! If you are as confused as a donkey in a library, don't worry. Go (Golang) is simple. Let's break it down like we're eating hay.

---

## 1. The Package: Where do I live?

Every file must say what "group" it belongs to. Usually, for your main stuff, it's `main`.

```go
package main // I am in the 'main' group. Easy.
```

## 2. The Import: Asking for help

Go is a bit lazy. It doesn't know how to do anything unless you give it a tool.
`fmt` (Format) is the tool used to talk to the screen.

```go
import "fmt" // Give me the 'fmt' tool so I can speak!
```

## 3. The Main Function: The Start Button

Your computer needs to know where to start running. Imagine a big green button that says `func main()`.

```go
func main() {
    // Everything inside these curly { } brackets runs!
}
```

---

## 4. Talking to the Screen (Printing)

Use your `fmt` tool to shout things out.

```go
fmt.Println("Hee-haw!") // Prints a line and moves to the next
fmt.Print("I like hay")    // Prints but stays on the same line
```

---

## 5. Variables: Storage Boxes

Imagine you have boxes to put stuff in.

### The "Fancy" Way (Long)

You say: "I want a box, name it `X`, put a `Type` on it, and put `Value` inside."

```go
var name string = "Akash"
var age int = 22
var isHungry bool = true
```

### The "Lazy donkey" Way (Short)

Use `:=`. Go is smart enough to see what's inside and label the box for you.

```go
food := "Carrots" // Go sees " " and knows it is a STRING
price := 5.99    // Go sees . and knows it is a FLOAT
```

_Note: This only works INSIDE functions!_

---

## 6. Constants: The "Never Change" Rules

Some things never change, like a donkey's stubbornness. Use `const`.

```go
const species = "Equus africanus asinus"
// If you try to change this later, Go will yell at you!
```

---

## 7. Loops: Groundhog Day

In Go, there is only ONE way to repeat things: `for`. No `while`, no `do-while`. Just `for`.

```go
// Do this 5 times
for i := 0; i < 5; i++ {
    fmt.Println("Eating hay...")
}

// The "Keep going until I'm full" way (Like a while loop)
for condition {
    // keep doing it
}
```

---

## 8. If/Else: Decision Time

"If there is hay, eat it. Otherwise, sleep."

```go
if food == "hay" {
    fmt.Println("Yum!")
} else if food == "carrot" {
    fmt.Println("Better!")
} else {
    fmt.Println("I'm sad.")
}
```

---

## 9. Arrays: A Row of Boxes

An array is like a fence with a fixed number of slots. You CANNOT add more slots later.

```go
var fence [3]string
fence[0] = "Donkey 1"
fence[1] = "Donkey 2"
fence[2] = "Donkey 3"

// Or all at once:
colors := [2]string{"Grey", "Brown"}
```

---

## 10. Slices: The Magic Stretching Fence

Slices are like arrays, but they can grow! They are much more common in Go.

```go
myList := []string{"Apple", "Banana"} // Notice: No number in the []
myList = append(myList, "Carrot")     // Now it has 3 items! Magic!
```

---

---

## 11. Switch: The Sorting Gate

If you have many choices, instead of 100 `if` statements, use a `switch`. It's like a gate that sends you to the right path.

```go
day := "Sunday"

switch day {
case "Monday":
    fmt.Println("Back to work...")
case "Sunday":
    fmt.Println("Resting time!")
default:
    fmt.Println("Just another day.")
}
```

---

## 12. Functions: Training the Donkey

A function is just a set of instructions you can name. Instead of saying "Walk 10 steps, bend head, eat hay" every time, you just call the function `Eat()`.

```go
func Greet(name string) {
    fmt.Println("Hello", name)
}

// Or one that gives something back:
func Add(a int, b int) int {
    return a + b
}
```

---

## ⚠️ Important Donkey Rules

1. **Unused Variables**: If you create a variable and don't use it, Go will refuse to run. It hates waste!
2. **Curly Brackets**: The `{` MUST be on the same line as the `func` or `if`. Don't put it on the next line or the donkey will kick!

---

_Now go forth and code, you majestic creature!_ 🫏🚀
