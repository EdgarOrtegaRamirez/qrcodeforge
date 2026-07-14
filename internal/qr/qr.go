package qr

import (
	"image/color"
	"os"

	"github.com/skip2/go-qrcode"
)

// Encode generates a QR code from text.
func Encode(text string, level qrcode.RecoveryLevel, format string, output string) error {
	size := 256

	switch format {
	case "png":
		if output == "" {
			output = "qrcode.png"
		}
		return qrcode.WriteFile(text, level, size, output)
	case "svg":
		if output == "" {
			output = "qrcode.png"
		}
		return qrcode.WriteFile(text, level, size, output)
	default:
		// Default: ANSI output to stdout
		data, err := qrcode.Encode(text, level, size)
		if err != nil {
			return err
		}
		os.Stdout.Write(data)
		return nil
	}
}

// EncodeWithColors generates a QR code with custom foreground/background colors.
func EncodeWithColors(text, output string, fgColor, bgColor color.Color) error {
	if output == "" {
		output = "qrcode.png"
	}
	return qrcode.WriteColorFile(text, qrcode.Medium, 256, fgColor, bgColor, output)
}

// EncodeVariableSize generates a QR code with variable module size.
// Use negative sizes for variable module sizing (e.g., -4 for 4px modules).
func EncodeVariableSize(text string, size int, output string) error {
	if output == "" {
		output = "qrcode.png"
	}
	return qrcode.WriteFile(text, qrcode.Medium, size, output)
}
