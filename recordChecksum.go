package main

// recordChecksum preserves the original Pascal control flow.
func (m *gameState) recordChecksum() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (12) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu("sub", m.sp, 12, 16)) & 65535
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, 106, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-10)&65535, m.di)
	m.w16(m.ss, (m.bp+-8)&65535, m.ds)
	m.temp = m.r16(m.ss, (m.bp+-10)&65535)
	m.es = m.r16(m.ss, ((m.bp+-10)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-12)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-12)&65535), 16)
	if !m.c && !m.z {
		goto L5d
	}
	m.w16(m.ss, (m.bp+-6)&65535, m.ax)
	goto L42
L3f:
	m.clock(16)
	m.w16(m.ss, (m.bp+-6)&65535, m.step(m.r16(m.ss, (m.bp+-6)&65535), 1, 16))
L42:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-6)&65535)) & 65535
	m.temp = m.r16(m.ss, (m.bp+-10)&65535)
	m.es = m.r16(m.ss, ((m.bp+-10)&65535)+2)
	m.di = (m.temp) & 65535
	m.di = (m.alu("add", m.di, m.ax, 16)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-4)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-6)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-12)&65535), 16)
	if !m.z {
		goto L3f
	}
L5d:
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-10)&65535)
	m.es = m.r16(m.ss, ((m.bp+-10)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+13)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.clock(16)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.temp = m.r16(m.ss, (m.bp+-10)&65535)
	m.es = m.r16(m.ss, ((m.bp+-10)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+15)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.dx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.temp = m.r16(m.ss, (m.bp+-10)&65535)
	m.es = m.r16(m.ss, ((m.bp+-10)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.es, (m.di+11)&65535), 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.dx, 16)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.ax = (m.alu("add", m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu("adc", m.dx, m.bx, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.clock(16)
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.cx = (256) & 65535
	m.product("div", m.cx)
	m.temp = m.dx
	m.dx = (m.ax) & 65535
	m.ax = (m.temp) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.cx = (256) & 65535
	m.product("div", m.cx)
	m.temp = m.dx
	m.dx = (m.ax) & 65535
	m.ax = (m.temp) & 65535
	m.w16(m.ss, (m.bp+-2)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 6) & 65535
	return
}
