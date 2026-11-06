package main

// exitGame preserves the original Pascal control flow.
func (m *gameState) exitGame() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.runtime(0x6ea, 0x244)
	m.runtime(0x688, 0x2f3)
	m.w16(m.ds, (5536)&65535, 3)
	m.ax = (m.ax & 65280) | ((16) & 255)
	m.push(m.ax)
	m.di = (5536) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x681, 0xb)
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.writeRecords()
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.runtime(0x6ea, 0xd8)
	return
}
