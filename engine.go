package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// gameState is the recovered Pascal program's data and arithmetic state.
// All game procedures are compiled Go functions. There is no instruction
// stream, CPU interpreter, DOS executable, DOS emulator or virtual machine.
// The 16-bit data layout retains the original array indexing and overflow.
type gameState struct {
	recordsReady                                         bool
	memory                                               [1 << 20]byte
	ax, bx, cx, dx, si, di, bp, sp, ds, es, ss, cs, temp int
	c, z, s, o                                           bool
	planes                                               [4][65536]byte
	latch                                                [4]byte
	gc                                                   [16]byte
	gcIndex                                              int
	pixels                                               [320 * 200 * 4]byte
	files                                                map[int]*legacyFile
	nextHandle, heap                                     int
	frames                                               chan frame
	resume                                               chan inputState
	keys                                                 []int
	mouseX, mouseY, mouseButtons                         int
	pendingScan                                          int
	frequency                                            int
	elapsed                                              float64
	toneBase                                             int
	toneEvents                                           []toneEvent
	phase                                                float64
	seed                                                 uint32
	closing                                              bool
	failure                                              error
	// untimed > 0 while drawing must not consume emulated CPU time.
	untimed int
	// pcm is reused for every audio frame: soundStream.append copies it.
	pcm [800 * 4]byte
}
type legacyFile struct {
	name        string
	data        []byte
	pos, record int
	writable    bool
}
type toneEvent struct {
	at        float64
	frequency int
}
type frame struct {
	pixels    []byte
	frequency int
	duration  float64
	pcm       []byte
}
type inputState struct {
	keys                         []int
	mouseX, mouseY, mouseButtons int
	quit                         bool
}

func unpack(s string) []byte {
	b, e := base64.StdEncoding.DecodeString(s)
	if e != nil {
		panic(e)
	}
	r, e := zlib.NewReader(bytes.NewReader(b))
	if e != nil {
		panic(e)
	}
	defer r.Close()
	b, e = io.ReadAll(r)
	if e != nil {
		panic(e)
	}
	return b
}
func newState() *gameState {
	m := &gameState{ds: 0x842, ss: 0x2000, sp: 0xf000, heap: 0x40000, files: map[int]*legacyFile{}, nextHandle: 5, frames: make(chan frame), resume: make(chan inputState), seed: uint32(time.Now().UnixNano())}
	m.gc[8] = 255
	m.initLiterals()
	font := unpack(bitmapFont)
	copy(m.memory[0xffa6e:], font[:1024])
	copy(m.memory[0xf0000:], font[1024:])
	m.memory[0x7c] = 0
	m.memory[0x7d] = 0
	m.memory[0x7e] = 0
	m.memory[0x7f] = 0xf0
	return m
}
func (m *gameState) r8(seg, off int) int {
	a := ((seg << 4) + (off & 65535)) & 0xfffff
	if a >= 0xa0000 && a < 0xb0000 {
		p := a - 0xa0000
		for i := 0; i < 4; i++ {
			m.latch[i] = m.planes[i][p]
		}
		return int(m.latch[int(m.gc[4])&3])
	}
	return int(m.memory[a])
}
func (m *gameState) r16(seg, off int) int { lo := m.r8(seg, off); return lo | m.r8(seg, off+1)<<8 }
func (m *gameState) w8(seg, off, value int) {
	a := ((seg << 4) + (off & 65535)) & 0xfffff
	if a >= 0xa0000 && a < 0xb0000 {
		p := a - 0xa0000
		mask := m.gc[8]
		mode := m.gc[5] & 3
		for i := 0; i < 4; i++ {
			v := byte(value)
			switch mode {
			case 1:
				m.planes[i][p] = m.latch[i]
				continue
			case 2:
				v = 0
				if value&(1<<i) != 0 {
					v = 255
				}
			default:
				if m.gc[1]&(1<<i) != 0 {
					v = 0
					if m.gc[0]&(1<<i) != 0 {
						v = 255
					}
				}
			}
			switch (m.gc[3] >> 3) & 3 {
			case 1:
				v &= m.latch[i]
			case 2:
				v |= m.latch[i]
			case 3:
				v ^= m.latch[i]
			}
			m.planes[i][p] = (v & mask) | (m.latch[i] &^ mask)
		}
		return
	}
	m.memory[a] = byte(value)
}
func (m *gameState) w16(seg, off, v int) { m.w8(seg, off, v); m.w8(seg, off+1, v>>8) }
func (m *gameState) push(v int)          { m.sp = (m.sp - 2) & 65535; m.w16(m.ss, m.sp, v) }
func (m *gameState) pop() int            { v := m.r16(m.ss, m.sp); m.sp = (m.sp + 2) & 65535; return v }

// aluOp and mulOp select the operation of alu and product. They replace the
// former string arguments, which were compared on every emulated instruction.
type aluOp uint8

const (
	opAdd aluOp = iota
	opAdc
	opSub
	opSbb
	opAnd
	opOr
	opXor
)

type mulOp uint8

const (
	opMul mulOp = iota
	opImul
	opDiv
	opIdiv
)

func (m *gameState) alu(op aluOp, a, b, bits int) int {
	mask := (1 << bits) - 1
	sign := 1 << (bits - 1)
	a &= mask
	b &= mask
	r := 0
	carry := 0
	if m.c {
		carry = 1
	}
	switch op {
	case opAdd, opAdc:
		if op == opAdc {
			b += carry
		}
		r = a + b
		m.c = r > mask
		m.o = (^(a ^ b) & (a ^ r) & sign) != 0
	case opSub, opSbb:
		if op == opSbb {
			b += carry
		}
		r = a - b
		m.c = a < b
		m.o = ((a ^ b) & (a ^ r) & sign) != 0
	case opAnd:
		r = a & b
		m.c = false
		m.o = false
	case opOr:
		r = a | b
		m.c = false
		m.o = false
	case opXor:
		r = a ^ b
		m.c = false
		m.o = false
	}
	r &= mask
	m.z = r == 0
	m.s = r&sign != 0
	return r
}
func (m *gameState) step(v, n, bits int) int {
	c := m.c
	r := m.alu(opAdd, v, n, bits)
	m.c = c
	return r
}
func (m *gameState) shift(v, n, bits int, left bool) int {
	n &= 31
	if n == 0 {
		return v
	}
	// Equivalent to shifting one bit n times: the carry is the last bit
	// shifted out.
	if left {
		mask := (1 << bits) - 1
		m.c = n <= bits && v&(1<<(bits-n)) != 0
		v = (v << n) & mask
	} else {
		m.c = (v>>(n-1))&1 != 0
		v >>= n
	}
	m.z = v == 0
	m.s = v&(1<<(bits-1)) != 0
	return v
}
func (m *gameState) product(op mulOp, v int) {
	switch op {
	case opMul:
		p := uint32(m.ax) * uint32(v)
		m.ax = int(p & 65535)
		m.dx = int(p >> 16)
	case opImul:
		p := int32(int16(m.ax)) * int32(int16(v))
		m.ax = int(uint32(p) & 65535)
		m.dx = int(uint32(p) >> 16)
	case opDiv:
		p := uint32(m.dx)<<16 | uint32(m.ax)
		m.ax = int(p/uint32(v)) & 65535
		m.dx = int(p % uint32(v))
	case opIdiv:
		p := int32(uint32(m.dx)<<16 | uint32(m.ax))
		d := int32(int16(v))
		m.ax = int(p/d) & 65535
		m.dx = int(p%d) & 65535
	}
}
func (m *gameState) videoOut(port, v int) {
	switch port {
	case 0x3ce:
		m.gcIndex = v & 15
	case 0x3cf:
		m.gc[m.gcIndex] = byte(v)
	}
}
func (m *gameState) videoIn(port int) int { return 0 }
func (m *gameState) pascalString(seg, off int) string {
	n := m.r8(seg, off)
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(m.r8(seg, off+i+1))
	}
	return string(b)
}
func (m *gameState) putString(seg, off int, s string, max int) {
	if len(s) > max {
		s = s[:max]
	}
	m.w8(seg, off, len(s))
	for i := range s {
		m.w8(seg, off+1+i, int(s[i]))
	}
}
func (m *gameState) arg(n int) int            { return m.r16(m.ss, m.sp+n) }
func (m *gameState) pointer(n int) (int, int) { return m.arg(n + 2), m.arg(n) }
func (m *gameState) consume(n int)            { m.sp = (m.sp + n) & 65535 }
func (m *gameState) number() float64          { return realNumber(m.ax, m.bx, m.dx) }
func realNumber(ax, bx, dx int) float64 {
	if ax&255 == 0 {
		return 0
	}
	f := 1 + float64(uint64(dx&32767)<<24|uint64(bx)<<8|uint64(ax>>8))/float64(uint64(1)<<39)
	if dx&32768 != 0 {
		f = -f
	}
	return math.Ldexp(f, (ax&255)-129)
}
func (m *gameState) real(v float64) {
	if v == 0 {
		m.ax = 0
		m.bx = 0
		m.dx = 0
		return
	}
	sg := 0
	if v < 0 {
		sg = 32768
		v = -v
	}
	f, e := math.Frexp(v)
	bits := uint64(math.Round((f*2 - 1) * float64(uint64(1)<<39)))
	if bits == (uint64(1) << 39) {
		bits = 0
		e++
	}
	m.ax = ((e + 128) & 255) | int(bits&255)<<8
	m.bx = int((bits >> 8) & 65535)
	m.dx = int(bits>>24) | sg
}
func (m *gameState) randomWord() uint32 { m.seed = m.seed*134775813 + 1; return m.seed }
func (m *gameState) openFile(seg, off, rec int, write bool) {
	name := strings.ToLower(m.pascalString(seg, off))
	f := &legacyFile{name: name, record: rec, writable: write}
	switch name {
	case "sea.lib":
		f.data = unpack(pictureLibrary)
	case "sea.pic":
		f.data = unpack(initialScores)
		if b, e := os.ReadFile(scorePath()); e == nil && len(b)%17 == 0 && len(b) <= 4352 {
			f.data = b
		}
	default:
		panic("unexpected resource " + name)
	}
	if write {
		f.data = nil
	}
	h := m.nextHandle
	m.nextHandle++
	m.files[h] = f
	m.w16(seg, off, h)
	m.w16(seg, off+4, rec)
	m.w16(m.ds, 0x52, 0)
}
func scorePath() string {
	d, e := os.UserConfigDir()
	if e != nil {
		return "sea-scores.dat"
	}
	return filepath.Join(d, "BattleOnTheBlackSea", "scores.dat")
}
func (m *gameState) closeFile(seg, off int) {
	f := m.files[m.r16(seg, off)]
	if f != nil && f.writable {
		p := scorePath()
		if os.MkdirAll(filepath.Dir(p), 0755) == nil {
			tmp := p + ".tmp"
			if os.WriteFile(tmp, f.data, 0600) == nil {
				if os.Rename(tmp, p) != nil {
					_ = os.WriteFile(p, f.data, 0600)
					_ = os.Remove(tmp)
				}
			}
		}
	}
}
func (m *gameState) transfer(fs, fo, bs, bo, n int, write bool) int {
	f := m.files[m.r16(fs, fo)]
	if f == nil {
		panic("closed score file")
	}
	count := n * f.record
	if write {
		for i := 0; i < count; i++ {
			v := byte(m.r8(bs, bo+i))
			if f.pos < len(f.data) {
				f.data[f.pos] = v
			} else {
				f.data = append(f.data, v)
			}
			f.pos++
		}
	} else {
		if f.pos+count > len(f.data) {
			count = len(f.data) - f.pos
		}
		for i := 0; i < count; i++ {
			m.w8(bs, bo+i, int(f.data[f.pos+i]))
		}
		f.pos += count
	}
	return count / f.record
}
func (m *gameState) runtime(seg, fn int) {
	// These are native replacements for Pascal's library calls. The original
	// game never initializes or invokes a BGI graphics driver.
	if seg == 0x681 {
		is, io := m.pointer(0)
		interrupt := m.arg(4)
		if interrupt == 0x10 && m.r16(is, io) == 0xd {
			m.gc[5] = 0
			m.gc[8] = 255
		}
		if interrupt == 0x33 {
			switch m.r16(is, io) {
			case 0:
				m.w16(is, io, 65535)
				m.w16(is, io+2, 2)
			case 3:
				m.w16(is, io+2, m.mouseButtons)
				m.w16(is, io+4, m.mouseX*2)
				m.w16(is, io+6, m.mouseY)
			case 4:
				m.mouseX = m.r16(is, io+4) / 2
				m.mouseY = m.r16(is, io+6)
			}
		}
		m.consume(6)
		return
	}
	if seg == 0x688 {
		switch fn {
		case 0x29e:
			m.pause(float64(m.arg(0)) / 1000)
			m.consume(2)
		case 0x2c6:
			f := m.arg(0)
			if f > 18 {
				div := 1193181 / f
				m.tone(1193181 / div)
			}
			m.consume(2)
		case 0x2f3:
			m.tone(0)
		case 0x2fa:
			m.clock(1000)
			m.ax &= 65280
			if len(m.keys) > 0 || m.pendingScan != 0 {
				m.ax |= 1
			}
		case 0x30c:
			if m.pendingScan != 0 {
				m.ax = (m.ax & 65280) | m.pendingScan
				m.pendingScan = 0
				return
			}
			for len(m.keys) == 0 {
				m.yield()
			}
			k := m.keys[0]
			m.keys = m.keys[1:]
			m.ax &= 65280
			if k > 255 {
				m.pendingScan = k - 256
			} else {
				m.ax |= k
			}
		default:
			panic(fmt.Sprintf("CRT %x", fn))
		}
		return
	}
	switch fn {
	case 0xd8:
		panic(gameExit{})
	case 0x244: // Pascal stack check.
	case 0xd4e: // Random seed was initialized at startup.
	case 0x207:
		m.ax = m.r16(m.ds, 0x52)
		m.w16(m.ds, 0x52, 0)
	case 0x20e:
		if n := m.r16(m.ds, 0x52); n != 0 {
			panic(fmt.Sprintf("I/O error %d", n))
		}
	case 0x261:
		v := (uint32(m.dx)<<16 | uint32(m.ax)) * (uint32(m.bx)<<16 | uint32(m.cx))
		m.ax = int(v & 65535)
		m.dx = int(v >> 16)
	case 0x311:
		n := m.arg(0)
		a := m.heap
		m.heap += (n + 15) &^ 15
		if m.heap >= 0x90000 {
			panic("picture allocation overflow")
		}
		m.ax = a & 15
		m.dx = a >> 4
		m.consume(2)
	case 0x612:
		ss, so := m.pointer(0)
		ds, do := m.pointer(4)
		m.putString(ds, do, m.pascalString(ss, so), 255)
		m.consume(4)
	case 0x62c:
		ss, so := m.pointer(6)
		ds, do := m.pointer(2)
		m.putString(ds, do, m.pascalString(ss, so), m.arg(0))
		m.consume(10)
	case 0x65e:
		ss, so := m.pointer(4)
		ds, do := m.pointer(8)
		s := m.pascalString(ss, so)
		start := m.arg(2) - 1
		n := m.arg(0)
		if start < 0 {
			start = 0
		}
		if start > len(s) {
			start = len(s)
		}
		end := start + n
		if end > len(s) {
			end = len(s)
		}
		m.putString(ds, do, s[start:end], 255)
		m.consume(8)
	case 0x69f:
		ss, so := m.pointer(0)
		ds, do := m.pointer(4)
		s := m.pascalString(ds, do) + m.pascalString(ss, so)
		m.putString(ds, do, s, 255)
		m.consume(4)
	case 0x717:
		s1, o1 := m.pointer(4)
		s2, o2 := m.pointer(0)
		r := strings.Compare(m.pascalString(s1, o1), m.pascalString(s2, o2))
		m.z = r == 0
		m.c = r < 0
		m.s = r < 0
		m.o = false
		m.consume(8)
	case 0x742:
		ds, do := m.pointer(2)
		m.putString(ds, do, string([]byte{byte(m.arg(0))}), 255)
		m.consume(2)
	case 0x8c7:
		ds, do := m.pointer(0)
		v := m.arg(4) & 255
		m.z = m.r8(ds, do+v/8)&(1<<(v%8)) == 0
		m.consume(6)
	case 0xc81:
		m.real(m.number() - realNumber(m.cx, m.si, m.di))
	case 0xc8d:
		m.real(m.number() * realNumber(m.cx, m.si, m.di))
	case 0xc93:
		m.real(m.number() / realNumber(m.cx, m.si, m.di))
	case 0xc9d:
		a, b := m.number(), realNumber(m.cx, m.si, m.di)
		m.z = a == b
		m.c = a < b
		m.s = a < b
		m.o = false
	case 0xca1:
		m.real(float64(int32(uint32(m.dx)<<16 | uint32(m.ax))))
	case 0xcad:
		v := int32(math.Round(m.number()))
		m.ax = int(uint32(v) & 65535)
		m.dx = int(uint32(v) >> 16)
	case 0xcc7:
		n := m.arg(0)
		r := m.randomWord()
		if n == 0 {
			m.ax = 0
		} else {
			m.ax = int(r>>16) % n
		}
		m.consume(2)
	case 0xcdd:
		r := m.randomWord()
		m.real(float64(r) / 4294967296)
	case 0xe6d:
		limit := m.arg(0)
		ds, do := m.pointer(2)
		width := m.arg(6)
		v := int32(uint32(m.arg(10))<<16 | uint32(m.arg(8)))
		s := fmt.Sprint(v)
		if len(s) < width {
			s = strings.Repeat(" ", width-len(s)) + s
		}
		m.putString(ds, do, s, limit)
		m.consume(12)
	case 0x12f0:
		ss, so := m.pointer(0)
		ds, do := m.pointer(4)
		m.putString(ds, do, m.pascalString(ss, so), 79)
		m.consume(8)
	case 0x131e, 0x1327:
		ds, do := m.pointer(2)
		m.openFile(ds, do, m.arg(0), fn == 0x1327)
		m.consume(6)
	case 0x139f:
		ds, do := m.pointer(0)
		m.closeFile(ds, do)
		m.consume(4)
	case 0x13d3, 0x13da:
		bs, bo := m.pointer(0)
		fs, fo := m.pointer(4)
		m.transfer(fs, fo, bs, bo, 1, fn == 0x13da)
		m.consume(4)
	case 0x1409:
		rs, ro := m.pointer(0)
		n := m.arg(4)
		bs, bo := m.pointer(6)
		fs, fo := m.pointer(10)
		got := m.transfer(fs, fo, bs, bo, n, false)
		if rs|ro != 0 {
			m.w16(rs, ro, got)
		}
		m.consume(14)
	case 0x14d3:
		fs, fo := m.pointer(0)
		f := m.files[m.r16(fs, fo)]
		m.ax = 0
		if f.pos >= len(f.data) {
			m.ax = 1
		}
		m.consume(4)
	case 0x152d:
		n := m.arg(0)
		ds, do := m.pointer(2)
		ss, so := m.pointer(6)
		b := make([]byte, n)
		if ss >= 0xa000 && ss < 0xb000 && ds >= 0xa000 && ds < 0xb000 && m.gc[5]&3 == 1 {
			for j := 0; j < 4; j++ {
				copy(b, m.planes[j][so:so+n])
				copy(m.planes[j][do:do+n], b)
			}
		} else {
			for i := range b {
				b[i] = byte(m.r8(ss, so+i))
			}
			for i, v := range b {
				m.w8(ds, do+i, int(v))
			}
		}
		m.consume(10)
	case 0x1550:
		n := m.arg(2)
		v := m.arg(0)
		ds, do := m.pointer(4)
		for i := 0; i < n; i++ {
			m.w8(ds, do+i, v)
		}
		m.consume(8)
	case 0x1564:
		v := m.arg(0) & 255
		if v >= 'a' && v <= 'z' {
			v -= 32
		}
		m.ax = (m.ax & 65280) | v
		m.consume(2)
	case 0x1235, 0x12b2:
		panic("unexpected console output")
	default:
		panic(fmt.Sprintf("Pascal library %x:%x", seg, fn))
	}
}

type gameExit struct{}

func (m *gameState) run() {
	defer func() {
		if e := recover(); e != nil {
			if _, ok := e.(gameExit); !ok {
				m.failure = fmt.Errorf("game: %v", e)
				fmt.Fprintln(os.Stderr, m.failure)
			}
		}
		close(m.frames)
	}()
	m.push(0)
	m.push(0)
	m.runGame()
}

var ega = [16][3]byte{{0, 0, 0}, {0, 0, 170}, {0, 170, 0}, {0, 170, 170}, {170, 0, 0}, {170, 0, 170}, {170, 85, 0}, {170, 170, 170}, {85, 85, 85}, {85, 85, 255}, {85, 255, 85}, {85, 255, 255}, {255, 85, 85}, {255, 85, 255}, {255, 255, 85}, {255, 255, 255}}

// egaRGBA is ega with an opaque alpha channel, ready to be copied as a pixel.
var egaRGBA = func() (t [16][4]byte) {
	for i, c := range ega {
		t[i] = [4]byte{c[0], c[1], c[2], 255}
	}
	return
}()

// planarToRGBA converts the 4 EGA bit planes (40 bytes per line, 8 pixels per
// byte) of the 320x200 screen into RGBA pixels.
func (m *gameState) planarToRGBA() {
	p0, p1, p2, p3 := &m.planes[0], &m.planes[1], &m.planes[2], &m.planes[3]
	px := m.pixels[:]
	for off := 0; off < 8000; off++ {
		b0, b1, b2, b3 := p0[off], p1[off], p2[off], p3[off]
		i := off * 32
		for bit := 7; bit >= 0; bit-- {
			c := (b0>>bit)&1 | ((b1>>bit)&1)<<1 | ((b2>>bit)&1)<<2 | ((b3>>bit)&1)<<3
			copy(px[i:i+4], egaRGBA[c][:])
			i += 4
		}
	}
}

func (m *gameState) yield() {
	m.planarToRGBA()
	// The window thread copies the pixels and PCM in Update before it sends
	// the next resume, and this goroutine waits for that resume, so both
	// buffers can be reused without copying them for every frame.
	m.frames <- frame{pixels: m.pixels[:], frequency: m.frequency, duration: 1.0 / 60, pcm: m.audioFrame()}
	in := <-m.resume
	m.keys = append(m.keys, in.keys...)
	m.mouseX = in.mouseX
	m.mouseY = in.mouseY
	m.mouseButtons = in.mouseButtons
	m.elapsed = 0
	if in.quit {
		if !m.recordsReady {
			panic(gameExit{})
		}
		m.closing = true
		m.ds = 0x842
		m.cs = 0
		m.push(0)
		m.push(0)
		m.exitGame()
	}
}

const frameSeconds = 1.0 / 60

func (m *gameState) advance(seconds float64) {
	for seconds > 0 {
		remaining := frameSeconds - m.elapsed
		if seconds < remaining {
			m.elapsed += seconds
			return
		}
		m.elapsed = frameSeconds
		m.yield()
		seconds -= remaining
	}
}
func (m *gameState) pause(seconds float64) { m.advance(seconds) }
func (m *gameState) clock(cycles int) {
	if m.closing || m.untimed > 0 {
		return
	}
	m.advance(float64(cycles) / 3000000)
}
func (m *gameState) tone(frequency int) {
	m.frequency = frequency
	m.toneEvents = append(m.toneEvents, toneEvent{m.elapsed, frequency})
}

// drawSpriteUntimed draws a sprite like drawSprite but takes no emulated time.
func (m *gameState) drawSpriteUntimed() {
	m.untimed++
	defer func() { m.untimed-- }()
	m.drawSprite()
}
func (m *gameState) audioFrame() []byte {
	pcm := m.pcm[:]
	frequency := m.toneBase
	event := 0
	for i := 0; i < 800; i++ {
		for event < len(m.toneEvents) && m.toneEvents[event].at <= float64(i)/48000 {
			frequency = m.toneEvents[event].frequency
			m.phase = 0
			event++
		}
		v := int16(0)
		if frequency > 0 {
			v = 2600
			if m.phase >= 0.5 {
				v = -v
			}
			m.phase += float64(frequency) / 48000
			m.phase -= math.Floor(m.phase)
		}
		u := uint16(v)
		j := i * 4
		pcm[j] = byte(u)
		pcm[j+1] = byte(u >> 8)
		pcm[j+2] = byte(u)
		pcm[j+3] = byte(u >> 8)
	}
	m.toneEvents = m.toneEvents[:0]
	m.toneBase = m.frequency
	return pcm
}
