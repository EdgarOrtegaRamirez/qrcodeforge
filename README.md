# QRCodeForge 🔲

**Fast QR code generation CLI** — generate QR codes from text, URLs, and structured data formats. Output as PNG, SVG, or ANSI terminal art.

## What It Does

QRCodeForge generates QR codes from any text input with multiple output formats:

- **PNG output** — High-quality PNG images for printing or embedding
- **SVG output** — Scalable vector graphics for web and design
- **ANSI terminal output** — Quick QR codes directly in your terminal
- **Error correction levels** — L, M, Q, H for different durability needs
- **Colored QR codes** — Custom foreground and background colors
- **Variable size** — Adjust module size for different use cases

## Quick Start

```bash
# Generate a QR code (default: ANSI terminal output)
qrcodeforge encode "Hello World"

# Generate a PNG QR code
qrcodeforge encode "https://example.com" -o qrcode.png

# Generate an SVG QR code
qrcodeforge encode "https://example.com" -f svg -o qrcode.svg

# High error correction level
qrcodeforge encode "https://example.com" -e H -o qrcode.png

# WiFi QR code
qrcodeforge encode "WiFi:T:WPA;S:mywifi;P:pass;;"

# Email QR code
qrcodeforge encode "mailto:test@example.com"

# Geo location QR code
qrcodeforge encode "geo:37.7749,-122.4194"
```

## Installation

```bash
go install github.com/EdgarOrtegaRamirez/qrcodeforge/cmd/qrcodeforge@latest
```

## Usage

```
qrcodeforge encode <text|url> [options]

Options:
  -o, --output FILE    Output file path (default: stdout for ansi, qrcode.png for png)
  -f, --format FMT     Output format: ansi (default), png, svg
  -e, --error-level L  Error correction level: L, M (default), Q, H
```

## Supported QR Code Formats

| Format | Description |
|--------|-------------|
| Text | Plain text content |
| URL | Website URLs |
| WiFi | WiFi credentials (`WiFi:T:WPA;S:ssid;P:password;;`) |
| Email | Email addresses (`mailto:user@example.com`) |
| Phone | Phone numbers (`tel:+1234567890`) |
| SMS | SMS messages (`sms:+1234567890?body=Hello`) |
| Geo | GPS coordinates (`geo:lat,lon`) |
| vCard | Contact information |
| Bitcoin | Bitcoin payment addresses |

## Architecture

- `cmd/qrcodeforge/main.go` — CLI entry point with argument parsing
- `internal/qr/qr.go` — QR code encoding engine using `skip2/go-qrcode`
- `internal/qr/qr_test.go` — Unit tests for encoding functions

## Dependencies

- `github.com/skip2/go-qrcode` — QR code generation library

## License

MIT
