package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCommandProcessing(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantCmd string
	}{
		{
			name:    "simaple command",
			input:   "ls\n",
			wantCmd: "ls",
		},
		{
			name:    "command with argument",
			input:   "ls -la\n",
			wantCmd: "ls -la",
		},
		{
			name:    "command with space",
			input:   "  ls  \n",
			wantCmd: "ls",
		},
		{
			name:    "empty command",
			input:   "\n",
			wantCmd: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			command, err := reader.ReadString('\n')
			if err != nil {
				t.Fatalf("Read failed: %v", err)
			}

			command = strings.TrimSpace(command)
			if command != tt.wantCmd {
				t.Errorf("Got %q, want %q", command, tt.wantCmd)
			}
		})
	}
}

func TestEOF(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	_, err := reader.ReadString('\n')
	if err == nil {
		t.Error("Expected EOF error, got nil")
	}
}

// Benchmark to test performance
func BenchmarkCommandRead(b *testing.B) {
	input := strings.Repeat("ls -la\n", 1000)
	reader := bufio.NewReader(strings.NewReader(input))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader = bufio.NewReader(strings.NewReader(input))
		for {
			command, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			_ = strings.TrimSpace(command)
		}
	}
}
