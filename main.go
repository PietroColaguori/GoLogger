package main

import (
	"fmt"
	"os"
	"time"

	hook "github.com/robotn/gohook"
)

func main() {
	// Manage -h and --help CLI options
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		printHelp()
	}

	// Create output file
	file, errs := os.Create("logs.txt")
	if errs != nil {
		fmt.Println("Failed to create logs.txt: ", errs)
		return
	}
	defer file.Close()

	// Initialize event listener
	evChan := hook.Start()

	// Schedule cleanup of channel when main function ends
	defer hook.End()

	for ev := range evChan {
		// use to inspect raw output
		// fmt.Printf("Event: %v\n", ev)

		if formatEvent(ev) != "" {
			fmt.Println(formatEvent(ev))
		}

		if len(os.Args) > 1 && os.Args[1] == "-s" {
			_, errs = file.WriteString(formatEvent(ev))
			if errs != nil {
				fmt.Println("Failed to write to logs.txt: ", errs)
				return
			}
		}
		_, errs = file.WriteString(fmt.Sprintf("%s\n", formatEvent(ev)))
		if errs != nil {
			fmt.Println("Failed to write to logs.txt: ", errs)
			return
		}
	}
}

func formatEvent(ev hook.Event) string {
	switch ev.Kind {
	case hook.KeyDown:
		if len(os.Args) > 1 && os.Args[1] == "-t" {
			timestamp := time.Now()
			return fmt.Sprintf("Key pressed: %s [%s]", keyName(ev), timestamp)
		}
		if len(os.Args) > 1 && os.Args[1] == "-s" {
			return keyName(ev)
		}
		return fmt.Sprintf("Key pressed: %s", keyName(ev))
	case hook.MouseDown:
		if len(os.Args) > 1 && os.Args[1] == "-t" {
			timestamp := time.Now()
			return fmt.Sprintf("Mouse button pressed: %s [%s]", mouseButtonName(ev.Button), timestamp)
		}
		if len(os.Args) > 1 && os.Args[1] == "-s" {
			return mouseButtonName(ev.Button)
		}
		return fmt.Sprintf("Key pressed: %s", mouseButtonName(ev.Button))
	default:
		return ""
	}
}

// Identify key
func keyName(ev hook.Event) string {

	// Catch control characters and some extended ASCII characters
	switch ev.Keychar {
	case 1:
		return "[start of heading]"
	case 2:
		return "[start of text]"
	case 3:
		return "[end of text]"
	case 4:
		return "[end of transmission]"
	case 5:
		return "[enquiry]"
	case 6:
		return "[acknowledge]"
	case 7:
		return "[bell]"
	case 8:
		return "[backspace]"
	case 9:
		return "[horizontal tab]"
	case 10:
		return "[line feed]"
	case 11:
		return "[vertical tab]"
	case 12:
		return "[form feed]"
	case 13:
		return "[carriage return]"
	case 14:
		return "[shift out]"
	case 15:
		return "[shift in]"
	case 16:
		return "[data link escape]"
	case 17:
		return "[device control 1]"
	case 18:
		return "[device control 2]"
	case 19:
		return "[device control 3]"
	case 20:
		return "[device control 4]"
	case 21:
		return "[negative acknowledge]"
	case 22:
		return "[synchronous idle]"
	case 23:
		return "[end of transmission block]"
	case 24:
		return "[cancel]"
	case 25:
		return "[end of medium]"
	case 26:
		return "[substitute]"
	case 27:
		return "[escape]"
	case 28:
		return "[file separator]"
	case 29:
		return "[group separator]"
	case 30:
		return "[record separator]"
	case 31:
		return "[unit separator]"
	case 32:
		return "[space]"
	case 33:
		return "[exclamation mark]"
	case 34:
		return "[double quote]"
	case 35:
		return "[number sign]"
	case 36:
		return "[dollar sign]"
	case 37:
		return "[percent]"
	case 38:
		return "[ampersand]"
	case 39:
		return "[single quote]"
	case 40:
		return "[left parenthesis]"
	case 41:
		return "[right parenthesis]"
	case 42:
		return "[asterisk]"
	case 43:
		return "[plus sign]"
	case 44:
		return "[comma]"
	case 45:
		return "[hyphen]"
	case 46:
		return "[period]"
	case 47:
		return "[slash]"
	case 58:
		return "[colon]"
	case 59:
		return "[semicolon]"
	case 60:
		return "[less than]"
	case 61:
		return "[equals]"
	case 62:
		return "[greater than]"
	case 63:
		return "[question mark]"
	case 64:
		return "[at sign]"
	case 91:
		return "[left square bracket]"
	case 92:
		return "[backslash]"
	case 93:
		return "[right square bracket]"
	case 94:
		return "[caret]"
	case 95:
		return "[underscore]"
	case 96:
		return "[grave accent]"
	case 123:
		return "[left curly brace]"
	case 124:
		return "[pipe]"
	case 125:
		return "[right curly brace]"
	case 126:
		return "[tilde]"
	}

	// Special keys mapped to null
	if ev.Keychar == 0 {
		switch ev.Rawcode {
		case 65507:
			return "[left control]"
		case 65505:
			return "[left shift]"
		case 65509:
			return "[case lock]"
		case 65513:
			return "[alt]"
		case 65027:
			return "[alt gr]"
		}
	}

	// Printable ASCII characters not already caught
	if ev.Keychar > 32 && ev.Keychar <= 127 {
		return string(ev.Keychar)
	}

	return fmt.Sprintf("Unknown key (raw code %d, key char %d)", ev.Rawcode, ev.Keychar)
}

// identify mouse button
func mouseButtonName(button uint16) string {
	switch button {
	case 1:
		return "left"
	case 2:
		return "middle"
	case 3:
		return "right"
	default:
		return fmt.Sprintf("button: %d", button)
	}
}

func printHelp() {
	fmt.Println("-t: display the timestamps")
	fmt.Println("-s: display just characters")
	fmt.Println("-h, --help: display this help dialog")
	os.Exit(0)
}
