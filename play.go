package main

// play preserves the original Pascal control flow.
func (m *gameState) play() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.runtime(0x6ea, 0x244)
	m.w16(m.ds, (4458)&65535, 5)
	m.w16(m.ds, (4460)&65535, 5)
	m.w16(m.ds, (4462)&65535, 1)
	m.w16(m.ds, (4464)&65535, 1)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (4466)&65535, m.ax)
L2ff2:
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (5526)&65535), 1, 16)
	if !m.z {
		goto L2fff
	}
	m.push(m.cs)
	m.push(0)
	m.playerTurn()
	goto L3003
L2fff:
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.computerTurn()
L3003:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 27, 8)
	if m.z {
		goto L3018
	}
	m.alu(opSub, m.r16(m.ds, (4472)&65535), 0, 16)
	if m.z {
		goto L3018
	}
	m.alu(opSub, m.r16(m.ds, (4474)&65535), 0, 16)
	m.clock(16)
	if !m.z {
		goto L2ff2
	}
L3018:
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
