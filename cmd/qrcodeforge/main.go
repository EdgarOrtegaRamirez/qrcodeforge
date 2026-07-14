package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/EdgarOrtegaRamirez/qrcodeforge/internal/qr"
	"github.com/skip2/go-qrcode"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "encode":
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Error: encode requires text/URL input")
			printUsage()
			os.Exit(1)
		}
		text := args[0]
		var output string
		var format string = "ansi"
		var level qrcode.RecoveryLevel = qrcode.Medium

		for i := 1; i < len(args); i++ {
			switch args[i] {
			case "-o", "--output":
				if i+1 < len(args) {
					output = args[i+1]
					i++
				}
			case "-f", "--format":
				if i+1 < len(args) {
					format = args[i+1]
					i++
				}
			case "-e", "--error-level":
				if i+1 < len(args) {
					switch args[i+1] {
					case "L":
						level = qrcode.Low
					case "M":
						level = qrcode.Medium
					case "Q":
						level = qrcode.High
					case "H":
						level = qrcode.Highest
					}
					i++
				}
			}
		}

		if err := qr.Encode(text, level, format, output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

// parseColor parses a hex color string to a simple RGB struct.
func parseColor(hex string) interface{} {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 6 {
		r, _ := strconv.ParseUint(hex[0:2], 16, 8)
		g, _ := strconv.ParseUint(hex[2:4], 16, 8)
		b, _ := strconv.ParseUint(hex[4:6], 16, 8)
		return struct{ R, G, B, A uint8 }{uint8(r), uint8(g), uint8(b), 255}
	}
	return nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `qrcodeforge — QR code generation CLI

Usage:
  qrcodeforge encode <text|url> [options]

Options:
  -o, --output FILE    Output file path (default: stdout for ansi, qrcode.png for png)
  -f, --format FMT     Output format: ansi (default), png, svg
  -e, --error-level L  Error correction level: L, M (default), Q, H

Examples:
  qrcodeforge encode "Hello World"
  qrcodeforge encode "https://example.com" -o qrcode.png
  qrcodeforge encode "https://example.com" -f png -o qr.png
  qrcodeforge encode "data:text/plain,hello" -f svg -o qr.svg
  qrcodeforge encode "WiFi:T:WPA;S:mywifi;P:pass;;"
  qrcodeforge encode "mailto:test@example.com"
  qrcodeforge encode "geo:37.7749,-122.4194"
`)
}
