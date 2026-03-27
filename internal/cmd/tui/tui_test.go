// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package tui

import (
	"context"
	"os"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/gittuf/gittuf/pkg/gitinterface"
	"github.com/muesli/termenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	lipgloss.SetColorProfile(termenv.Ascii)
}

func TestTUIStartAndQuit(t *testing.T) {
	tmpDir := t.TempDir()
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(currentDir) //nolint:errcheck

	gitinterface.CreateTestGitRepository(t, tmpDir, false)

	o := &options{
		readOnly: true,
	}

	m, err := initialModel(context.Background(), o)
	if err != nil {
		t.Fatalf("failed to create initial model: %v", err)
	}

	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))
	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second*5))
}

func TestSplitAndTrim(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"a, b, c", []string{"a", "b", "c"}},
		{"a", []string{"a"}},
		{" a ,b, c ", []string{"a", "b", "c"}},
		{"", []string{""}},
	}

	for _, test := range tests {
		result := splitAndTrim(test.input)
		assert.Equal(t, test.expected, result)
	}
}
