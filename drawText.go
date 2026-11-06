package main

// drawText preserves the original Pascal control flow.
func (m *gameState) drawText() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (280) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu("sub", m.sp, 280, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+22)&65535)
	m.es = m.r16(m.ss, ((m.bp+22)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.di = ((m.bp + -256) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (255) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.si = (m.alu("xor", m.si, m.si, 16)) & 65535
	m.es = (m.si) & 65535
	m.ax = (m.r16(m.es, (124)&65535)) & 65535
	m.si = (m.alu("xor", m.si, m.si, 16)) & 65535
	m.es = (m.si) & 65535
	m.dx = (m.r16(m.es, (126)&65535)) & 65535
	m.w16(m.ss, (m.bp+-260)&65535, m.ax)
	m.w16(m.ss, (m.bp+-258)&65535, m.dx)
	m.ax = (m.r16(m.ss, (m.bp+20)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (8) & 65535
	m.product("idiv", m.cx)
	m.w16(m.ss, (m.bp+-272)&65535, m.ax)
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+20)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (8) & 65535
	m.product("idiv", m.cx)
	m.temp = m.dx
	m.dx = (m.ax) & 65535
	m.ax = (m.temp) & 65535
	m.w8(m.ss, (m.bp+-273)&65535, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-273)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.dx = (m.ax) & 65535
	m.ax = (8) & 65535
	m.ax = (m.alu("sub", m.ax, m.dx, 16)) & 65535
	m.w8(m.ss, (m.bp+-274)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+18)&65535)) & 65535
	m.product("imul", m.r16(m.ss, (m.bp+14)&65535))
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-272)&65535), 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-262)&65535, m.ax)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((2) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-256)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-280)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-280)&65535), 16)
	if m.c || m.z {
		goto L2c1
	}
	goto L76a
L2c1:
	m.clock(16)
	m.w16(m.ss, (m.bp+-268)&65535, m.ax)
	goto L2cb
L2c7:
	m.clock(16)
	m.w16(m.ss, (m.bp+-268)&65535, m.step(m.r16(m.ss, (m.bp+-268)&65535), 1, 16))
L2cb:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-268)&65535)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+m.di+-256)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-266)&65535, m.ax)
	m.alu("sub", m.r16(m.ss, (m.bp+-266)&65535), 127, 16)
	if m.c || m.z {
		goto L2eb
	}
	m.ax = (m.r16(m.ss, (m.bp+-266)&65535)) & 65535
	m.ax = (m.alu("sub", m.ax, 128, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-266)&65535, m.ax)
L2eb:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-266)&65535)) & 65535
	m.cx = (3) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.w16(m.ss, (m.bp+-264)&65535, m.ax)
	m.di = (m.r16(m.ss, (m.bp+-268)&65535)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+m.di+-256)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-266)&65535, m.ax)
	m.alu("sub", m.r16(m.ss, (m.bp+16)&65535), 2, 16)
	if m.s != m.o {
		goto L30f
	}
	goto L48a
L30f:
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-270)&65535, m.ax)
	goto L31b
L317:
	m.clock(16)
	m.w16(m.ss, (m.bp+-270)&65535, m.step(m.r16(m.ss, (m.bp+-270)&65535), 1, 16))
L31b:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-270)&65535), 4, 16)
	if !m.c {
		goto L32b
	}
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+12)&65535)) & 255)
	m.w8(m.ss, (m.bp+-277)&65535, (m.ax & 255))
	goto L332
L32b:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+10)&65535)) & 255)
	m.w8(m.ss, (m.bp+-277)&65535, (m.ax & 255))
L332:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-266)&65535), 128, 16)
	if !m.c {
		goto L354
	}
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-270)&65535), 16)) & 65535
	m.di = (m.ax) & 65535
	m.si = (61440) & 65535
	m.es = (m.si) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+-1426)&65535)) & 255)
	m.w8(m.ss, (m.bp+-275)&65535, (m.ax & 255))
	goto L369
L354:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-270)&65535), 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+-260)&65535)
	m.es = m.r16(m.ss, ((m.bp+-260)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-275)&65535, (m.ax & 255))
L369:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-273)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.dx = (m.ax) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-275)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.clock(16)
	m.cx = (m.dx) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, false)) & 65535
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.clock(16)
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-276)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-277)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.alu("sub", m.r8(m.ss, (m.bp+-274)&65535), 8, 8)
	if !m.z {
		goto L3f2
	}
	m.clock(16)
	goto L47d
L3f2:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-274)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.dx = (m.ax) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-275)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.cx = (m.dx) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, 1, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, 0, 16)) & 65535
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-276)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.clock(16)
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, 1, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, 0, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-277)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
L47d:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-270)&65535), 7, 16)
	if m.z {
		goto L487
	}
	goto L317
L487:
	m.clock(16)
	goto L754
L48a:
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-270)&65535, m.ax)
	goto L496
L492:
	m.clock(16)
	m.w16(m.ss, (m.bp+-270)&65535, m.step(m.r16(m.ss, (m.bp+-270)&65535), 1, 16))
L496:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-270)&65535), 4, 16)
	if !m.c {
		goto L4a6
	}
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+12)&65535)) & 255)
	m.w8(m.ss, (m.bp+-277)&65535, (m.ax & 255))
	goto L4ad
L4a6:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+10)&65535)) & 255)
	m.w8(m.ss, (m.bp+-277)&65535, (m.ax & 255))
L4ad:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-266)&65535), 128, 16)
	if !m.c {
		goto L4cf
	}
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-270)&65535), 16)) & 65535
	m.di = (m.ax) & 65535
	m.clock(16)
	m.si = (61440) & 65535
	m.es = (m.si) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+-1426)&65535)) & 255)
	m.w8(m.ss, (m.bp+-275)&65535, (m.ax & 255))
	goto L4e4
L4cf:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-270)&65535), 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+-260)&65535)
	m.es = m.r16(m.ss, ((m.bp+-260)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-275)&65535, (m.ax & 255))
L4e4:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-273)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.dx = (m.ax) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-275)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.cx = (m.dx) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, false)) & 65535
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (2) & 65535
	m.bx = (m.alu("xor", m.bx, m.bx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.clock(16)
	m.w8(m.ss, (m.bp+-276)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (2) & 65535
	m.bx = (m.alu("xor", m.bx, m.bx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.clock(16)
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-277)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (2) & 65535
	m.bx = (m.alu("xor", m.bx, m.bx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.clock(16)
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-276)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (2) & 65535
	m.clock(16)
	m.bx = (m.alu("xor", m.bx, m.bx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.clock(16)
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-277)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.alu("sub", m.r8(m.ss, (m.bp+-274)&65535), 8, 8)
	if !m.z {
		goto L610
	}
	goto L74a
L610:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.clock(16)
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-274)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.dx = (m.ax) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-275)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.cx = (m.dx) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.clock(16)
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (2) & 65535
	m.bx = (m.alu("xor", m.bx, m.bx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, 1, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, 0, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-276)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (2) & 65535
	m.bx = (m.alu("xor", m.bx, m.bx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.clock(16)
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, 1, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, 0, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-277)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.clock(16)
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (2) & 65535
	m.bx = (m.alu("xor", m.bx, m.bx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.clock(16)
	m.ax = (m.alu("add", m.ax, 1, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, 0, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-276)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.ax) & 65535
	m.clock(16)
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-270)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (2) & 65535
	m.bx = (m.alu("xor", m.bx, m.bx, 16)) & 65535
	m.runtime(0x6ea, 0x261)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, 1, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, 0, 16)) & 65535
	m.clock(16)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-277)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
L74a:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-270)&65535), 7, 16)
	if m.z {
		goto L754
	}
	goto L492
L754:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ss, (m.bp+-262)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-268)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-280)&65535), 16)
	if m.z {
		goto L76a
	}
	goto L2c7
L76a:
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 24) & 65535
	return
}
