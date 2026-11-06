package main

// drawPlacementCursor preserves the original Pascal control flow.
func (m *gameState) drawPlacementCursor() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (2) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu(opSub, m.sp, 2, 16)) & 65535
	m.temp = m.r16(m.ds, (5684)&65535)
	m.es = m.r16(m.ds, ((5684)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.clock(16)
	m.push(m.di)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (16) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.temp = m.r16(m.ds, (5692)&65535)
	m.es = m.r16(m.ds, ((5692)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.clock(16)
	m.push(m.es)
	m.push(m.di)
	m.ax = (212) & 65535
	m.push(m.ax)
	m.ax = (30) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.w16(m.ss, (m.bp+-2)&65535, 9)
L30ce:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.clock(16)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.alu(opSub, m.r8(m.ds, (m.di+5353)&65535), 0, 8)
	if !m.z {
		goto L312c
	}
	m.alu(opSub, m.r16(m.ss, (m.bp+-2)&65535), 1, 16)
	if m.z || m.s != m.o {
		goto L312c
	}
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.di = (m.ax) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.ax = (m.alu(opAdd, m.ax, 4, 16)) & 65535
	m.di = (m.ax) & 65535
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
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.cx = (4) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.dx = (m.ax) & 65535
	m.clock(16)
	m.ax = (160) & 65535
	m.ax = (m.alu(opSub, m.ax, m.dx, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.w16(m.ss, (m.bp+-2)&65535, m.step(m.r16(m.ss, (m.bp+-2)&65535), -1, 16))
	goto L30ce
L312c:
	m.clock(16)
	m.di = (12316) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.clock(16)
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (47) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.clock(16)
	m.ax = (46) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.clock(16)
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (12338) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.clock(16)
	m.push(m.di)
	m.ax = (31) & 65535
	m.push(m.ax)
	m.ax = (173) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.clock(16)
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (30) & 65535
	m.push(m.ax)
	m.ax = (172) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (12363) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.clock(16)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (61) & 65535
	m.push(m.ax)
	m.ax = (182) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.clock(16)
	m.push(m.di)
	m.ax = (60) & 65535
	m.push(m.ax)
	m.ax = (181) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.clock(16)
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (12382) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.clock(16)
	m.push(m.ds)
	m.push(m.di)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.ax = (191) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.clock(16)
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (190) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.clock(16)
	m.sp = (m.sp + 6) & 65535
	return
}
