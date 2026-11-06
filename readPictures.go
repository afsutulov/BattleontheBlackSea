package main

// readPictures preserves the original Pascal control flow.
func (m *gameState) readPictures() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (514) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu(opSub, m.sp, 514, 16)) & 65535
	m.di = (2780) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (255) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = ((m.bp + -512) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.runtime(0x6ea, 0x12f0)
	m.di = ((m.bp + -512) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (128) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x131e)
	m.clock(16)
	m.runtime(0x6ea, 0x207)
	m.ax = (m.alu(opOr, m.ax, m.ax, 16)) & 65535
	if m.z {
		goto Lb81
	}
	m.di = (7068) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.di = (2788) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x12b2)
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x12b2)
	m.di = (2794) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x12b2)
	m.runtime(0x6ea, 0x1235)
	m.runtime(0x6ea, 0x20e)
	m.ax = (1) & 65535
	m.runtime(0x6ea, 0xd8)
	return
Lb81:
	m.clock(16)
	m.di = ((m.bp + -512) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = ((m.bp + -128) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1409)
	m.runtime(0x6ea, 0x20e)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-128)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-514)&65535, m.ax)
	m.clock(16)
	m.ax = (1) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-514)&65535), 16)
	if !m.z && m.s == m.o {
		goto Lc1a
	}
	m.w16(m.ds, (5558)&65535, m.ax)
	goto Lbb9
Lbb5:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
Lbb9:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+m.di+-128)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ds, (5560)&65535, m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.clock(16)
	m.cx = (7) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x311)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.di = (m.ax) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.w16(m.ds, (m.di+5572)&65535, m.cx)
	m.w16(m.ds, (m.di+5574)&65535, m.bx)
	m.di = ((m.bp + -512) & 65535) & 65535
	m.push(m.ss)
	m.clock(16)
	m.push(m.di)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.di = (m.ax) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.temp = m.r16(m.ds, (m.di+5572)&65535)
	m.es = m.r16(m.ds, ((m.di+5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.r16(m.ds, (5560)&65535))
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1409)
	m.runtime(0x6ea, 0x20e)
	m.clock(16)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-514)&65535), 16)
	if !m.z {
		goto Lbb5
	}
Lc1a:
	m.clock(16)
	m.di = ((m.bp + -512) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.runtime(0x6ea, 0x139f)
	m.runtime(0x6ea, 0x20e)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
