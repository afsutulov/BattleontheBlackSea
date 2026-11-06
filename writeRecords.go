package main

// writeRecords preserves the original Pascal control flow.
func (m *gameState) writeRecords() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (388) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu("sub", m.sp, 388, 16)) & 65535
	m.di = (174) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = ((m.bp + -256) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (255) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = ((m.bp + -256) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.runtime(0x6ea, 0x12f0)
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (17) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x1327)
	m.clock(16)
	m.runtime(0x6ea, 0x20e)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.w16(m.ss, (m.bp+-388)&65535, m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-388)&65535), 16)
	if !m.z && m.s == m.o {
		goto L15e
	}
	m.w16(m.ss, (m.bp+-386)&65535, m.ax)
	goto L117
L113:
	m.clock(16)
	m.w16(m.ss, (m.bp+-386)&65535, m.step(m.r16(m.ss, (m.bp+-386)&65535), 1, 16))
L117:
	m.clock(16)
	m.push(m.r16(m.ss, (m.bp+-386)&65535))
	m.push(m.cs)
	m.clock(16)
	m.push(0)
	m.recordChecksum()
	m.cx = (m.cx & 65280) | ((m.ax & 255) & 255)
	m.ax = (m.r16(m.ss, (m.bp+-386)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.w8(m.ds, (m.di+122)&65535, (m.cx & 255))
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (m.r16(m.ss, (m.bp+-386)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, 106, 16)) & 65535
	m.clock(16)
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x6ea, 0x13da)
	m.sp = (m.alu("add", m.sp, 4, 16)) & 65535
	m.runtime(0x6ea, 0x20e)
	m.ax = (m.r16(m.ss, (m.bp+-386)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-388)&65535), 16)
	if !m.z {
		goto L113
	}
L15e:
	m.clock(16)
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.runtime(0x6ea, 0x139f)
	m.runtime(0x6ea, 0x20e)
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
