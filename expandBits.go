package main

// expandBits preserves the original Pascal control flow.
func (m *gameState) expandBits() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (6) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu(opSub, m.sp, 6, 16)) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-6)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	goto L1f4
L1f1:
	m.clock(16)
	m.w16(m.ss, (m.bp+-4)&65535, m.step(m.r16(m.ss, (m.bp+-4)&65535), 1, 16))
L1f4:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.cx = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, false)) & 65535
	m.ax = (m.alu(opAnd, m.ax, 1, 16)) & 65535
	m.cx = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.cx = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.alu(opOr, m.ax, m.r16(m.ss, (m.bp+-6)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-6)&65535, m.ax)
	m.alu(opSub, m.r16(m.ss, (m.bp+-4)&65535), 7, 16)
	if !m.z {
		goto L1f1
	}
	m.ax = (m.r16(m.ss, (m.bp+-6)&65535)) & 65535
	m.clock(16)
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-6)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-2)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 6) & 65535
	return
}
