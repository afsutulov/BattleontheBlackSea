package main

import (
	"io"
	"sync"
)

// soundStream is a live PCM source. Reads must never wait for a future game
// frame: Oto preloads its buffer synchronously in Play on Windows, and Play
// runs inside Update. Waiting here would prevent Update from generating PCM.
type soundStream struct {
	mu     sync.Mutex
	buffer []byte
	closed bool
}

func newSoundStream() *soundStream { return &soundStream{} }
func (s *soundStream) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return 0, io.EOF
	}
	n := copy(p, s.buffer)
	s.buffer = s.buffer[n:]
	// Preserve queued samples and pad only an underrun with digital silence.
	// Returning a full buffer also avoids a busy retry loop in the audio mixer.
	clear(p[n:])
	return len(p), nil
}
func (s *soundStream) append(p []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return
	}
	if len(s.buffer) > 48000*4 {
		s.buffer = s.buffer[:0]
	}
	s.buffer = append(s.buffer, p...)
}
func (s *soundStream) close() {
	s.mu.Lock()
	s.closed = true
	s.buffer = nil
	s.mu.Unlock()
}
