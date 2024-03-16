package logger

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// captureOutput is a helper function to capture the output of logger methods.
func captureOutput(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // execute function

	w.Close()
	os.Stdout = old // restore the real stdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func TestLogger_Info(t *testing.T) {
	logger := New(false) // Creating a logger instance with verbosity off

	expected := "[INFO] Test message\n"
	result := captureOutput(func() {
		logger.Info("Test message")
	})

	if !strings.Contains(result, expected) {
		t.Errorf("Expected %q to contain %q", result, expected)
	}
}

func TestLogger_Detail_WhenVerbose(t *testing.T) {
	logger := New(true) // Creating a logger instance with verbosity on

	expected := "------- Verbose test message\n"
	result := captureOutput(func() {
		logger.Detail("Verbose test message")
	})

	if !strings.Contains(result, expected) {
		t.Errorf("Expected %q to contain %q", result, expected)
	}
}

func TestLogger_Detail_WhenNotVerbose(t *testing.T) {
	logger := New(false) // Creating a logger instance with verbosity off

	expected := "" // Expecting no output when verbosity is off
	result := captureOutput(func() {
		logger.Detail("This message should not be displayed")
	})

	if strings.Contains(result, "This message should not be displayed") {
		t.Errorf("Expected %q to be %q", result, expected)
	}
}

func TestLogger_Warn(t *testing.T) {
	logger := New(false) // Creating a logger instance

	expected := "[WARNING] Warning message"
	result := captureOutput(func() {
		logger.Warn("Warning message")
	})

	if !strings.Contains(result, expected) {
		t.Errorf("Expected %q to contain %q", result, expected)
	}
}

func TestLogger_Error(t *testing.T) {
	logger := New(false) // Creating a logger instance

	expected := "[ERROR] Error message"
	result := captureOutput(func() {
		logger.Error("Error message")
	})

	if !strings.Contains(result, expected) {
		t.Errorf("Expected %q to contain %q", result, expected)
	}
}

func TestLogger_Success(t *testing.T) {
	logger := New(false) // Creating a logger instance

	expected := "[SUCCESS] Success message"
	result := captureOutput(func() {
		logger.Success("Success message")
	})

	if !strings.Contains(result, expected) {
		t.Errorf("Expected %q to contain %q", result, expected)
	}
}

// Note: Testing Fatal and Debug methods which call os.Exit will require a different approach
