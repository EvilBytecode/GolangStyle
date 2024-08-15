# GoStyle (DONT USE NOT RELIABLE AT ALL + WIN ONLY)
Best Looking Windows Color Library made for Go users!
<a href="https://t.me/pulzetools"><img src="https://img.shields.io/badge/Join%20my%20Telegram%20group-2CA5E0?style=for-the-badge&logo=telegram&labelColor=db44ad&color=5e2775"></a>

![ConsoleLooks](https://github.com/EvilBytecode/GolangStyle/assets/151552809/d1bb0610-2c74-4049-b37e-797fb5c4d345)

## Install
- ```go mod init bestcolorlib```
- ```go get github.com/EvilBytecode/GolangStyle/pkg```
- ```go run .```

## GoStyle Functions

- Sample README that provides documentation for the functions in the GoStyle package.
## `Init() error`

- **Description:** Initializes the console, comes with simple error handle.
- **Sample Usage:** 
```go
if err := gostyle.Init(); err != nil {
    gostyle.Write("Failed to init console:", gostyle.RED_TO_BLACK, false)
    return
}
```

## `ClearConsole()`

- **Description:** Clears the console. Anything currently displayed will be removed.
- **Sample Usage:** 
```go
gostyle.ClearConsole()
```

## `HideCursor()`

- **Description:** Hides the console cursor.
- **Sample Usage:** 
```go
gostyle.HideCursor()
```

## `ShowCursor()`

- **Description:** Shows the console cursor.
- **Sample Usage:** 
```go
gostyle.ShowCursor()
```

## `Write(text string, gradient gostyle.Gradient, true/false)`

- **Description:** Writes text to the console with an optional gradient and centers it based on your choice.
- **Parameters:**
  - `text string`: Text to be written to the console.
  - `gradient gostyle.Gradient`: Optional gradient color for the text.
  - `true`: You Decide if it should be in center set true or false based on if you want the text to be centered or not.
- **Sample Usage:** 
```go
gostyle.Write("Your text here", gostyle.PURPLE_TO_BLUE, true)
```

## `WriteColorized(text string, color string, true/false)`

- **Description:** Writes text to the console with a specified color and centers it based on your choice.
- **Parameters:**
  - `text string`: Text to be written to the console.
  - `color string`: Color of the text (e.g., "red", "blue").
  - `true`: You Decide if it should be in center set true or false based on if you want the text to be centered or not.
- **Sample Usage:** 
```go
gostyle.WriteColorized("Your text here", "red", true)
```

### MORE IN ```main.go```

Remember to import the `gostyle` package before using these functions.
### issues with some ascii styles will be fixed soon, i just need to find time for it. but most of them are supported
### Credits:
- https://github.com/MmCopyMemory (Lots of help)
- https://github.com/billythegoat356 (GoStyle Idea, From Python to Go)
