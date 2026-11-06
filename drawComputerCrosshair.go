package main

// drawComputerCrosshair preserves the original Pascal control flow.
func (m *gameState) drawComputerCrosshair() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.runtime(0x6ea, 0x244)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.clock(16)
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
	m.push(m.cs)
	m.push(0)
	m.drawOpponentShips()
	m.di = (m.r16(m.ss, (m.bp+14)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.clock(16)
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.temp = m.r16(m.ds, (m.di+5572)&65535)
	m.es = m.r16(m.ds, ((m.di+5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.r16(m.ss, (m.bp+12)&65535))
	m.push(m.r16(m.ss, (m.bp+10)&65535))
	m.push(m.r16(m.ss, (m.bp+8)&65535))
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.di+-14)&65535), 0, 16)
	if !m.z && m.s == m.o {
		goto L24d4
	}
	goto L2576
L24d4:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-14)&65535, m.step(m.r16(m.ss, (m.di+-14)&65535), 1, 16))
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.di+-16)&65535)) & 65535
	m.ax = (m.alu("sub", m.ax, 2, 16)) & 65535
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-16)&65535, m.ax)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.di+-14)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.clock(16)
	m.cx = (5) & 65535
	m.product("idiv", m.cx)
	m.temp = m.dx
	m.dx = (m.ax) & 65535
	m.ax = (m.temp) & 65535
	m.ax = (m.alu("or", m.ax, m.ax, 16)) & 65535
	if !m.z {
		goto L250d
	}
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.di+-18)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-18)&65535, m.ax)
L250d:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.alu("sub", m.r16(m.ss, (m.di+-14)&65535), 15, 16)
	if m.z || m.s != m.o {
		goto L2522
	}
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-20)&65535, 21)
	goto L252b
L2522:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-20)&65535, 20)
L252b:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.alu("sub", m.r16(m.ss, (m.di+-14)&65535), 28, 16)
	if m.z || m.s != m.o {
		goto L253e
	}
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.di+-14)&65535, m.ax)
L253e:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.di = (m.r16(m.ss, (m.di+-20)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.temp = m.r16(m.ds, (m.di+5572)&65535)
	m.es = m.r16(m.ds, ((m.di+5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.di+-16)&65535)) & 65535
	m.clock(16)
	m.ax = (m.alu("add", m.ax, 3, 16)) & 65535
	m.push(m.ax)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.di+-18)&65535)) & 65535
	m.ax = (m.alu("sub", m.ax, 4, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
L2576:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5530)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.cx = (10) & 65535
	m.clock(16)
	m.product("div", m.cx)
	m.temp = m.dx
	m.dx = (m.ax) & 65535
	m.ax = (m.temp) & 65535
	m.alu("sub", m.ax, 5, 16)
	if !m.c {
		goto L259b
	}
	m.di = (9312) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	goto L25ae
L259b:
	m.clock(16)
	m.di = (9318) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
L25ae:
	m.clock(16)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (3) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu("add", m.ax, 6, 16)) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu("add", m.ax, 2, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.clock(16)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu("add", m.ax, 14, 16)) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu("add", m.ax, 14, 16)) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.clock(16)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.clock(16)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 14) & 65535
	return
}
