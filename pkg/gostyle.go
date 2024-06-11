package gostyle

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

// please don't skid or you're a noob lol
var (
	k32                  = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleMode   = k32.NewProc("SetConsoleMode")
	procGetConsoleMode   = k32.NewProc("GetConsoleMode")
	pgsh                 = k32.NewProc("GetStdHandle")
	pgcci                = k32.NewProc("GetConsoleCursorInfo")
	pscci                = k32.NewProc("SetConsoleCursorInfo")
	// these are static colors, made for simple printing they are not assosicated with gradients colors at all.
	Reset                = "\033[0m"
	Red                  = "\033[31m"
	Green                = "\033[32m"
	Yellow               = "\033[33m"
	Blue                 = "\033[34m"
	Magenta              = "\033[35m"
	Cyan                 = "\033[36m"
	White                = "\033[37m"
	// these are gradient colors, there arent assosicated with static ones, these are the colors that have transitions
	BLACK_TO_WHITE       = []int{0, 0, 0, 255, 255, 255}
	BLACK_TO_RED         = []int{0, 0, 0, 255, 0, 0}
	BLACK_TO_GREEN       = []int{0, 0, 0, 0, 255, 0}
	BLACK_TO_BLUE        = []int{0, 0, 0, 0, 0, 255}
	WHITE_TO_BLACK       = []int{255, 255, 255, 0, 0, 0}
	WHITE_TO_RED         = []int{255, 255, 255, 255, 0, 0}
	WHITE_TO_GREEN       = []int{255, 255, 255, 0, 255, 0}
	WHITE_TO_BLUE        = []int{255, 255, 255, 0, 0, 255}
	RED_TO_BLACK         = []int{255, 0, 0, 0, 0, 0}
	RED_TO_WHITE         = []int{255, 0, 0, 255, 255, 255}
	RED_TO_YELLOW        = []int{255, 0, 0, 255, 255, 0}
	RED_TO_PURPLE        = []int{255, 0, 0, 255, 0, 255}
	GREEN_TO_BLACK       = []int{0, 255, 0, 0, 0, 0}
	GREEN_TO_WHITE       = []int{0, 255, 0, 255, 255, 255}
	GREEN_TO_YELLOW      = []int{0, 255, 0, 255, 255, 0}
	GREEN_TO_CYAN        = []int{0, 255, 0, 0, 255, 255}
	BLUE_TO_BLACK        = []int{0, 0, 255, 0, 0, 0}
	BLUE_TO_WHITE        = []int{0, 0, 255, 255, 255, 255}
	BLUE_TO_CYAN         = []int{0, 0, 255, 0, 255, 255}
	BLUE_TO_PURPLE       = []int{0, 0, 255, 255, 0, 255}
	YELLOW_TO_RED        = []int{255, 255, 0, 255, 0, 0}
	YELLOW_TO_GREEN      = []int{255, 255, 0, 0, 255, 0}
	PURPLE_TO_RED        = []int{255, 0, 255, 255, 0, 0}
	PURPLE_TO_BLUE       = []int{255, 0, 255, 0, 0, 255}
	CYAN_TO_GREEN        = []int{0, 255, 255, 0, 255, 0}
	CYAN_TO_BLUE         = []int{0, 255, 255, 0, 0, 255}

	// this is init, this makes sure that colors print out successfully
	ENABLE_VIRTUAL_TERMINAL_PROCESSING uint32 = 0x0004
)

/*
	GradientFade applies a gradient color effect to text.
*/
func GradientFade(text string, colors []int) string {
	var result string
	for i, char := range text {
		currentR := int(float64(colors[0]) + (float64(colors[3]-colors[0])/float64(len(text)-1))*float64(i))
		currentG := int(float64(colors[1]) + (float64(colors[4]-colors[1])/float64(len(text)-1))*float64(i))
		currentB := int(float64(colors[2]) + (float64(colors[5]-colors[2])/float64(len(text)-1))*float64(i))

		result += fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c", currentR, currentG, currentB, char)
	}
	result += Reset
	return result
}

/*
	Init initializes the console with ANSI color support.
*/

func Init() error {
	handle := syscall.Handle(os.Stdout.Fd())

	var mode uint32
	ret, _, err := procGetConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&mode)))
	if ret == 0 {
		return err
	}

	mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING

	ret, _, err = procSetConsoleMode.Call(uintptr(handle), uintptr(mode))
	if ret == 0 {
		return err
	}

	return nil
}

const (
	soh = uint32(-11 & 0xFFFFFFFF)
)

type CONSOLE_CURSOR_INFO struct {
	Size    uint32
	Visible int32
}
// gets console handle
func getConsoleHandle() uintptr {
	handle, _, _ := pgsh.Call(uintptr(soh))
	return handle
}
// hide console cursror
func HideCursor() {
	handle := getConsoleHandle()
	var ci CONSOLE_CURSOR_INFO
	ci.Size = uint32(unsafe.Sizeof(ci))
	pgcci.Call(handle, uintptr(unsafe.Pointer(&ci)))
	ci.Visible = 0
	pscci.Call(handle, uintptr(unsafe.Pointer(&ci)))
}
// show console cursor
func ShowCursor() {
	handle := getConsoleHandle()
	var ci CONSOLE_CURSOR_INFO
	ci.Size = uint32(unsafe.Sizeof(ci))
	pgcci.Call(handle, uintptr(unsafe.Pointer(&ci)))
	ci.Visible = 1
	pscci.Call(handle, uintptr(unsafe.Pointer(&ci)))
}

/*
	CenterText centers text within the console window.
*/

func CenterText(text string) string {
    width, _ := getConsoleWindowSize()
    lines := strings.Split(text, "\n")
    var centeredText string
    for _, line := range lines {
        // Calculate visible length after applying color codes
        visibleLength := len(line) - strings.Count(line, "\033[")*7

        centeredText += fmt.Sprintf("%*s%s\n", (width-visibleLength)/2, "", line)
    }
    return centeredText
}



/*
	WriteColorized writes colorized text to the console, optionally centering it.
*/
func WriteColorized(text, color string, center bool) {
    if center {
        text = CenterText(text)
    }
    formattedText := Colorize(text, color)
    fmt.Println(formattedText)
}


/*
	getConsoleWindowSize retrieves the console window size.
*/
func getConsoleWindowSize() (width, height int) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("GetConsoleScreenBufferInfo")
	handle, _, _ := kernel32.NewProc("GetStdHandle").Call(uintptr(uint32(uintptr(0xfffffff5))))

	var info struct{ SizeX, SizeY int16 }

	ret, _, _ := proc.Call(handle, uintptr(unsafe.Pointer(&info)))
	if ret == 0 {
		return -1, -1
	}
	// wow no wayyy.. 
	return int(info.SizeX), int(info.SizeY)
}


/*
	Colorize applies color to text based on the provided color string.
*/
func Colorize(text, color string) string {
	color = strings.ToLower(color)
	switch color {
	case "red":
		return Red + text + Reset
	case "green":
		return Green + text + Reset
	case "yellow":
		return Yellow + text + Reset
	case "blue":
		return Blue + text + Reset
	case "magenta":
		return Magenta + text + Reset
	case "cyan":
		return Cyan + text + Reset
	case "white":
		return White + text + Reset
	default:
		return text
	}
}
/*
	Write writes text to the console, optionally applying colors and centering it.
*/
func Write(text string, colors []int, center bool) {
    if center {
        text = CenterText(text)
    }
    if len(colors) > 0 {
        text = GradientFade(text, colors)
    }
    lines := strings.Split(text, "\n")
    for _, line := range lines {
        fmt.Println(line)
    }
}
/*
	ClearConsole clears the console screen.
*/

func ClearConsole() {
cmd := exec.Command("cmd", "/c", "cls")
cmd.Stdout = os.Stdout
cmd.Run()
}
