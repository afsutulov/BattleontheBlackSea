package main

// drawPlayerCrosshair preserves the original Pascal control flow.
func (m *gameState) drawPlayerCrosshair() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
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
	m.alu(opSub, m.r16(m.ss, (m.di+-14)&65535), 0, 16)
	if !m.z && m.s == m.o {
		goto L1845
	}
	goto L18d5
L1845:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-14)&65535, m.step(m.r16(m.ss, (m.di+-14)&65535), 1, 16))
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.di+-16)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 2, 16)) & 65535
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-16)&65535, m.ax)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.di+-18)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-18)&65535, m.ax)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.alu(opSub, m.r16(m.ss, (m.di+-14)&65535), 10, 16)
	if m.z || m.s != m.o {
		goto L1881
	}
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-20)&65535, 21)
	goto L188a
L1881:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w16(m.ss, (m.di+-20)&65535, 20)
L188a:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.di+-14)&65535), 20, 16)
	if m.z || m.s != m.o {
		goto L189d
	}
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.di+-14)&65535, m.ax)
L189d:
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
	m.clock(16)
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.di+-16)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 3, 16)) & 65535
	m.push(m.ax)
	m.di = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.di+-18)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, 2, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.clock(16)
	m.push(0)
	m.drawSprite()
L18d5:
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
	m.clock(16)
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 14) & 65535
	return
}
