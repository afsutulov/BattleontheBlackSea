package main

// drawPlayerShips preserves the original Pascal control flow.
func (m *gameState) drawPlayerShips() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (12) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu("sub", m.sp, 12, 16)) & 65535
	m.w16(m.ss, (m.bp+-2)&65535, 1)
	goto L15cc
L15c9:
	m.clock(16)
	m.w16(m.ss, (m.bp+-2)&65535, m.step(m.r16(m.ss, (m.bp+-2)&65535), 1, 16))
L15cc:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.di = (m.alu("add", m.di, 5350, 16)) & 65535
	m.w16(m.ss, (m.bp+-12)&65535, m.di)
	m.w16(m.ss, (m.bp+-10)&65535, m.ds)
	m.temp = m.r16(m.ss, (m.bp+-12)&65535)
	m.es = m.r16(m.ss, ((m.bp+-12)&65535)+2)
	m.di = (m.temp) & 65535
	m.alu("sub", m.r8(m.es, (m.di+3)&65535), 0, 8)
	if !m.c && !m.z {
		goto L15ea
	}
	goto L167e
L15ea:
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-12)&65535)
	m.es = m.r16(m.ss, ((m.bp+-12)&65535)+2)
	m.di = (m.temp) & 65535
	m.alu("sub", m.r8(m.es, (m.di+2)&65535), 0, 8)
	if !m.z {
		goto L1629
	}
	m.temp = m.r16(m.ss, (m.bp+-12)&65535)
	m.es = m.r16(m.ss, ((m.bp+-12)&65535)+2)
	m.di = (m.temp) & 65535
	m.clock(16)
	m.ax = (m.r16(m.es, (m.di+4)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("sub", m.ax, 5, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-12)&65535)
	m.es = m.r16(m.ss, ((m.bp+-12)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+6)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 13, 16)) & 65535
	m.w16(m.ss, (m.bp+-6)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-12)&65535)
	m.es = m.r16(m.ss, ((m.bp+-12)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+1)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.ax = (m.alu("add", m.ax, 4, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
	goto L1659
L1629:
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-12)&65535)
	m.es = m.r16(m.ss, ((m.bp+-12)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+4)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("sub", m.ax, 4, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-12)&65535)
	m.es = m.r16(m.ss, ((m.bp+-12)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+6)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 9, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-6)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-12)&65535)
	m.es = m.r16(m.ss, ((m.bp+-12)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+1)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
L1659:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.temp = m.r16(m.ds, (m.di+5572)&65535)
	m.es = m.r16(m.ds, ((m.di+5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.clock(16)
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.r16(m.ss, (m.bp+-4)&65535))
	m.push(m.r16(m.ss, (m.bp+-6)&65535))
	m.ax = (1) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
L167e:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-2)&65535), 10, 16)
	if m.z {
		goto L1687
	}
	goto L15c9
L1687:
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
