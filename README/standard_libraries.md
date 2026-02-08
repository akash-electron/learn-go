# 🧰 The Donkey's Toolbelt: Standard Libraries

Go comes with a lot of built-in tools. You don't need to download them—they are already in your barn! Here are the ones every Go developer should know.

---

## 1. `fmt` (The Speaker 🗣️)

We already know this one. It's for talking to the screen and formatting strings.

- `fmt.Println()`: Say something.
- `fmt.Printf()`: Say something with fancy placeholders like `%v` or `%d`.

## 2. `net/http` (The Web Messenger 🌐)

This is Go's superpower. It lets you build web servers so easily, even a donkey could do it.

- Create a website server.
- Send requests to other websites.

## 3. `encoding/json` (The Secret Code 📝)

Computers love JSON. This tool turns Go boxes into JSON text and back again.

- `Marshal`: Go Box ➡️ JSON text.
- `Unmarshal`: JSON text ➡️ Go Box.

## 4. `os` (The Room Manager 🏠)

This tool lets you talk to your computer's operating system.

- `os.Open()`: Open a file.
- `os.Args`: See what words were typed when starting the program (command line arguments).
- `os.Exit()`: Shut everything down immediately.

## 5. `io` & `io/ioutil` (The Pipe 🚰)

Used for moving data around—like water through a pipe. Reading files, writing to strings, etc.

## 6. `time` (The Watch ⌚)

Need to wait? Need to know what day it is? Use this!

- `time.Now()`: What time is it right now?
- `time.Sleep()`: Take a nap for 5 seconds.

## 7. `strings` (The Word Surgeon ✂️)

Used for messing with text.

- `strings.Contains()`: Is the word "hay" inside this sentence?
- `strings.ToUpper()`: MAKE EVERYTHING LOUD.
- `strings.Split()`: Chop a sentence into a list of words.

## 8. `math` & `math/rand` (The Accountant 🔢)

- `math`: For big calculations like square roots or Pi.
- `math/rand`: For when you want a random number (like rolling a dice).

## 9. `slices` & `maps` (The Organizer 📂)

Newer versions of Go have extra tools to help you search and sort your collections.

## 10. `errors` (The Lifeguard 🛟)

In Go, we don't ignore problems. We use the `errors` package to create and handle messages when things go wrong.

---

### 🐴 How to use them?

Just put them at the top of your file!

```go
import (
    "fmt"
    "time"
    "strings"
)
```

---

_Now your toolbelt is heavy with powerful tools!_ 🫏🛠️
