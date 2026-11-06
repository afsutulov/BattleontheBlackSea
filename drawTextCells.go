package main

// drawTextCells preserves the original Pascal control flow.
func (m *gameState) drawTextCells() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (276) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu("sub", m.sp, 276, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+20)&65535)
	m.es = m.r16(m.ss, ((m.bp+20)&65535)+2)
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
	m.w16(m.ss, (m.bp+-270)&65535, m.ax)
	m.w16(m.ss, (m.bp+-268)&65535, m.dx)
	m.ax = (m.r16(m.ss, (m.bp+16)&65535)) & 65535
	m.cx = (40) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+18)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-258)&65535, m.ax)
	m.clock(16)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((2) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-256)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-276)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-276)&65535), 16)
	if m.c || m.z {
		goto L7dd
	}
	goto Lad6
L7dd:
	m.clock(16)
	m.w16(m.ss, (m.bp+-262)&65535, m.ax)
	goto L7e7
L7e3:
	m.clock(16)
	m.w16(m.ss, (m.bp+-262)&65535, m.step(m.r16(m.ss, (m.bp+-262)&65535), 1, 16))
L7e7:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+m.di+-256)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-266)&65535, m.ax)
	m.alu("sub", m.r16(m.ss, (m.bp+-266)&65535), 127, 16)
	if m.c || m.z {
		goto L807
	}
	m.ax = (m.r16(m.ss, (m.bp+-266)&65535)) & 65535
	m.ax = (m.alu("sub", m.ax, 128, 16)) & 65535
	m.w16(m.ss, (m.bp+-266)&65535, m.ax)
L807:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-266)&65535)) & 65535
	m.cx = (3) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.w16(m.ss, (m.bp+-260)&65535, m.ax)
	m.di = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+m.di+-256)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-266)&65535, m.ax)
	m.alu("sub", m.r16(m.ss, (m.bp+14)&65535), 2, 16)
	if m.s != m.o {
		goto L82b
	}
	goto L935
L82b:
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-264)&65535, m.ax)
	goto L837
L833:
	m.clock(16)
	m.w16(m.ss, (m.bp+-264)&65535, m.step(m.r16(m.ss, (m.bp+-264)&65535), 1, 16))
L837:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-264)&65535), 4, 16)
	if !m.c {
		goto L847
	}
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+12)&65535)) & 255)
	m.w8(m.ss, (m.bp+-272)&65535, (m.ax & 255))
	goto L84e
L847:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+10)&65535)) & 255)
	m.clock(16)
	m.w8(m.ss, (m.bp+-272)&65535, (m.ax & 255))
L84e:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-266)&65535), 128, 16)
	if !m.c {
		goto L870
	}
	m.ax = (m.r16(m.ss, (m.bp+-260)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-264)&65535), 16)) & 65535
	m.di = (m.ax) & 65535
	m.si = (61440) & 65535
	m.es = (m.si) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+-1426)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
	goto L885
L870:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-260)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-264)&65535), 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+-270)&65535)
	m.es = m.r16(m.ss, ((m.bp+-270)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
L885:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-271)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.expandBits()
	m.w16(m.ss, (m.bp+-274)&65535, m.ax)
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.clock(16)
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-273)&65535)) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (40) & 65535
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (40) & 65535
	m.product("mul", m.cx)
	m.clock(16)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-272)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-274)&65535)) & 255)
	m.dx = (975) & 65535
	m.clock(16)
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (40) & 65535
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (40) & 65535
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-272)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.alu("sub", m.r16(m.ss, (m.bp+-264)&65535), 7, 16)
	if m.z {
		goto L932
	}
	goto L833
L932:
	m.clock(16)
	goto Labe
L935:
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-264)&65535, m.ax)
	goto L941
L93d:
	m.clock(16)
	m.w16(m.ss, (m.bp+-264)&65535, m.step(m.r16(m.ss, (m.bp+-264)&65535), 1, 16))
L941:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-264)&65535), 4, 16)
	if !m.c {
		goto L951
	}
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+12)&65535)) & 255)
	m.w8(m.ss, (m.bp+-272)&65535, (m.ax & 255))
	goto L958
L951:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+10)&65535)) & 255)
	m.w8(m.ss, (m.bp+-272)&65535, (m.ax & 255))
L958:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-266)&65535), 128, 16)
	if !m.c {
		goto L97a
	}
	m.ax = (m.r16(m.ss, (m.bp+-260)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-264)&65535), 16)) & 65535
	m.di = (m.ax) & 65535
	m.si = (61440) & 65535
	m.es = (m.si) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+-1426)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
	goto L98f
L97a:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-260)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-264)&65535), 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+-270)&65535)
	m.es = m.r16(m.ss, ((m.bp+-270)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
L98f:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-271)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.expandBits()
	m.w16(m.ss, (m.bp+-274)&65535, m.ax)
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.clock(16)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-273)&65535)) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (80) & 65535
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (80) & 65535
	m.clock(16)
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-272)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (80) & 65535
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.clock(16)
	m.ax = (m.alu("add", m.ax, 40, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (80) & 65535
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.ax = (m.alu("add", m.ax, 40, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-272)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-274)&65535)) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (80) & 65535
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (80) & 65535
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-272)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (80) & 65535
	m.product("mul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.ax = (m.alu("add", m.ax, 41, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.ss, (m.bp+-271)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-264)&65535)) & 65535
	m.cx = (80) & 65535
	m.product("mul", m.cx)
	m.clock(16)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-258)&65535), 16)) & 65535
	m.ax = (m.alu("add", m.ax, 41, 16)) & 65535
	m.temp = m.r16(m.ss, (m.bp+6)&65535)
	m.es = m.r16(m.ss, ((m.bp+6)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-272)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.alu("sub", m.r16(m.ss, (m.bp+-264)&65535), 7, 16)
	if m.z {
		goto Labe
	}
	goto L93d
Labe:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-258)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, 2, 16)) & 65535
	m.w16(m.ss, (m.bp+-258)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-262)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-276)&65535), 16)
	if m.z {
		goto Lad6
	}
	goto L7e3
Lad6:
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 22) & 65535
	return
}
