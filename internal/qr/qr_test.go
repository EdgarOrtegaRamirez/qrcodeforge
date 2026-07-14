package qr

import (
	"os"
	"testing"
)

func TestEncodePNG(t *testing.T) {
	tmpFile := "/tmp/test_qrcode.png"
	defer os.Remove(tmpFile)

	err := Encode("https://example.com", 0, "png", tmpFile)
	if err != nil {
		t.Fatalf("EncodePNG failed: %v", err)
	}

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("Output file is empty")
	}
}

func TestEncodeDefault(t *testing.T) {
	err := Encode("Hello", 0, "", "")
	if err != nil {
		t.Fatalf("Encode default failed: %v", err)
	}
}

func TestEncodeVariableSize(t *testing.T) {
	tmpFile := "/tmp/test_qrcode_var.png"
	defer os.Remove(tmpFile)

	err := EncodeVariableSize("test", -4, tmpFile)
	if err != nil {
		t.Fatalf("EncodeVariableSize failed: %v", err)
	}
}
