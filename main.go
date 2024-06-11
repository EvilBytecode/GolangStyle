package main

import (
	"github.com/EvilBytecode/GolangStyle/pkg"
)

func main() {
//console init, changing title clearing it and setting adjustement
if err := gostyle.Init(); err != nil {
	gostyle.Write("Failed to init console:", gostyle.RED_TO_BLACK, false)
	return
}
// clears console (anything in there will be cleaned aka cls)
gostyle.ClearConsole()
// hides console cursor 
gostyle.HideCursor()
// Shows Console Cursor
gostyle.ShowCursor()
// Writes an Ascii with gradient fade
gostyle.Write(`
   _____       _____ _         _      
  / ____|     / ____| |       | |     
 | |  __  ___| (___ | |_ _   _| | ___ 
 | | |_ |/ _ \\___ \| __| | | | |/ _ \
 | |__| | (_) |___) | |_| |_| | |  __/
  \_____|\___/_____/ \__|\__, |_|\___|
                          __/ |       
                         |___/        
`, gostyle.PURPLE_TO_BLUE, true)

gostyle.WriteColorized("Best Looking Color Library written in Go for Windows", "red", true)
gostyle.Write("Coded by Evilbytecode & MMCopyMemory", gostyle.BLUE_TO_PURPLE, true)
gostyle.Write(`
Follow us on github:
https://github.com/MmCopyMemory
https://github.com/Evilbytecode
`, gostyle.BLUE_TO_PURPLE, true)
gostyle.Write("Without MMCopyMemory this wouldnt exist, follow him!", nil, true)
}
