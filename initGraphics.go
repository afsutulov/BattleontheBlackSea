package main

// initGraphics preserves the original Pascal control flow.
func (m *gameState) initGraphics() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.runtime(0x6ea, 0x244)
	m.w16(m.ds, (5536)&65535, 13)
	m.ax = (m.ax & 65280) | ((16) & 255)
	m.push(m.ax)
	m.di = (5536) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x681, 0xb)
	m.w16(m.ds, (5536)&65535, 2)
	m.clock(16)
	m.ax = (m.ax & 65280) | ((51) & 255)
	m.push(m.ax)
	m.di = (5536) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x681, 0xb)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
