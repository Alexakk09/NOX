package voice

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type STT struct {
	whisperPath string
	modelPath   string

	cmd     *exec.Cmd
	stdout  io.ReadCloser
	scanner *bufio.Scanner
}

func NewSTT(whisperPath string, modelPath string) *STT {
	return &STT{
		whisperPath: whisperPath,
		modelPath:   modelPath,
	}
}

func (s *STT) Start() error {
	if s.cmd != nil {
		return nil
	}

	s.cmd = exec.Command(
		s.whisperPath,
		"-m", s.modelPath,
		"-t", "4",
		"-c", "1",
		"--step", "0",
		"--length", "30000",
		"-vth", "0.6",
	)

	stdout, err := s.cmd.StdoutPipe()
	if err != nil {
		s.cmd = nil
		return fmt.Errorf("failed to create Whisper output pipe: %w", err)
	}

	if err := s.cmd.Start(); err != nil {
		s.cmd = nil
		return fmt.Errorf("failed to start Whisper: %w", err)
	}

	s.stdout = stdout
	s.scanner = bufio.NewScanner(stdout)

	return nil
}

func (s *STT) Listen() (string, error) {
	if s.cmd == nil || s.scanner == nil {
		return "", fmt.Errorf("Whisper is not running")
	}

	var transcription strings.Builder
	inTranscription := false

	for s.scanner.Scan() {
		line := strings.TrimSpace(s.scanner.Text())

		if strings.Contains(line, "### Transcription") &&
			strings.Contains(line, "START") {
			inTranscription = true
			transcription.Reset()
			continue
		}

		if strings.Contains(line, "### Transcription") &&
			strings.Contains(line, "END") {
			if transcription.Len() > 0 {
				return strings.TrimSpace(transcription.String()), nil
			}

			inTranscription = false
			continue
		}

		if !inTranscription {
			continue
		}

		if strings.HasPrefix(line, "[") &&
			strings.Contains(line, "-->") {

			parts := strings.SplitN(line, "]", 2)

			if len(parts) != 2 {
				continue
			}

			text := strings.TrimSpace(parts[1])

			if text == "" {
				continue
			}

			transcription.WriteString(text)
			transcription.WriteString(" ")
		}
	}

	if err := s.scanner.Err(); err != nil {
		return "", fmt.Errorf("failed reading Whisper output: %w", err)
	}

	return "", fmt.Errorf("Whisper stopped unexpectedly")
}

func (s *STT) Close() error {
	if s.cmd == nil {
		return nil
	}

	err := s.cmd.Process.Kill()

	if s.stdout != nil {
		_ = s.stdout.Close()
	}

	s.cmd = nil
	s.stdout = nil
	s.scanner = nil

	return err
}