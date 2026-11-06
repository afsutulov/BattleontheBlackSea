package main

import (
	"bytes"
	"io"
	"sync"
	"testing"
	"time"
)

// Oto preloads 50 ms synchronously on Windows. Update can only supply
// one 1/60 s frame until Play returns. Waiting for another frame deadlocks.
func TestFirstFrameCanPrefillAudioPlayer(t *testing.T) {
	s := newSoundStream()
	defer s.close()
	frame := bytes.Repeat([]byte{0x12, 0x34, 0x12, 0x34}, 800)
	s.append(frame)
	buf := make([]byte, 48000*4/20)
	done := make(chan error, 1)
	go func() { _, err := io.ReadFull(s, buf); done <- err }()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(250 * time.Millisecond):
		s.close()
		<-done
		t.Fatal("audio prefill waits for future Update frames: startup deadlock")
	}
	if !bytes.Equal(buf[:len(frame)], frame) {
		t.Fatal("queued PCM changed")
	}
	if !bytes.Equal(buf[len(frame):], make([]byte, len(buf)-len(frame))) {
		t.Fatal("underrun did not produce silence")
	}
}
func TestEmptyAudioReadReturnsSilence(t *testing.T) {
	s := newSoundStream()
	defer s.close()
	buf := bytes.Repeat([]byte{255}, 4096)
	done := make(chan int, 1)
	go func() {
		n, err := s.Read(buf)
		if err != nil {
			n = -1
		}
		done <- n
	}()
	select {
	case n := <-done:
		if n != len(buf) {
			t.Fatalf("read %d, want %d", n, len(buf))
		}
	case <-time.After(250 * time.Millisecond):
		s.close()
		<-done
		t.Fatal("empty audio buffer blocks playback")
	}
	if !bytes.Equal(buf, make([]byte, len(buf))) {
		t.Fatal("stale samples returned instead of silence")
	}
}
func TestAudioResumeAndClose(t *testing.T) {
	s := newSoundStream()
	buf := make([]byte, 4096)
	// A temporary underrun must not turn the live stream into EOF.
	if n, err := s.Read(buf); err != nil || n != len(buf) {
		t.Fatal(n, err)
	}
	pcm := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	s.append(pcm)
	if n, err := s.Read(buf); err != nil || n != len(buf) {
		t.Fatal(n, err)
	}
	if !bytes.Equal(buf[:len(pcm)], pcm) {
		t.Fatal("PCM did not resume after underrun")
	}
	s.close()
	s.close()
	s.append(pcm)
	if n, err := s.Read(buf); n != 0 || err != io.EOF {
		t.Fatal(n, err)
	}
}
func TestAudioConcurrentProducerAndConsumer(t *testing.T) {
	s := newSoundStream()
	defer s.close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 2000; i++ {
			s.append(make([]byte, 3200))
		}
	}()
	go func() {
		defer wg.Done()
		buf := make([]byte, 9600)
		for i := 0; i < 2000; i++ {
			s.Read(buf)
		}
	}()
	wg.Wait()
}
