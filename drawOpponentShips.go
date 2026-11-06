package main

// drawOpponentShips preserves the original Pascal control flow.
//
// It runs on every animation step of a shell (from drawPlayerCrosshair and
// drawComputerCrosshair) and redraws a fire marker on every hit cell of both
// fields. In the DOS original the shell's speed was simply the CPU time this
// redraw took, so each new hit made every following flight slower: a step
// takes about 23 ms on an empty board and about 135 ms with 20 hits. The
// fire markers are now drawn without consuming emulated time, so every step
// lasts as long as on the empty board at the start of the game.
func (m *gameState) drawOpponentShips() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.runtime(0x6ea, 0x244)
	m.w16(m.ds, (5558)&65535, 1)
	goto L16a1
L169d:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L16a1:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, 1)
	m.clock(16)
	goto L16ad
L16a9:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
L16ad:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r16(m.ds, (m.di+4494)&65535), 100, 16)
	m.clock(16)
	if m.z || m.s != m.o {
		goto L1707
	}
	m.ax = (4) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 12, 16)) & 65535
	m.di = (m.ax) & 65535
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
	m.clock(16)
	m.push(m.di)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 15, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSpriteUntimed()
L1707:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r16(m.ds, (m.di+4782)&65535), 100, 16)
	if m.z || m.s != m.o {
		goto L1763
	}
	m.ax = (4) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 12, 16)) & 65535
	m.clock(16)
	m.di = (m.ax) & 65535
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
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 159, 16)) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 15, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSpriteUntimed()
L1763:
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (5560)&65535), 10, 16)
	if m.z {
		goto L176d
	}
	goto L16a9
L176d:
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 10, 16)
	if m.z {
		goto L1777
	}
	goto L169d
L1777:
	m.clock(16)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 53, 16)) & 65535
	m.di = (m.ax) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.temp = m.r16(m.ds, (m.di+5572)&65535)
	m.es = m.r16(m.ds, ((m.di+5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.clock(16)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (60) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.ax = (3) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 50, 16)) & 65535
	m.di = (m.ax) & 65535
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
	m.ax = (250) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
