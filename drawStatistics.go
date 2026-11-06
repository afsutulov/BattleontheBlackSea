package main

// drawStatistics preserves the original Pascal control flow.
func (m *gameState) drawStatistics() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (270) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu("sub", m.sp, 270, 16)) & 65535
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((2) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.clock(16)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((255) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((3) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.clock(16)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((3) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.temp = m.r16(m.ds, (5692)&65535)
	m.es = m.r16(m.ds, ((5692)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.clock(16)
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (214) & 65535
	m.push(m.ax)
	m.ax = (30) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.di = (15666) & 65535
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
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.clock(16)
	m.alu("sub", m.r16(m.ds, (4470)&65535), 0, 16)
	if !m.z {
		goto L3f19
	}
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ds, (4470)&65535, m.ax)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ds, (4468)&65535, m.ax)
	m.di = (4482) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.clock(16)
	m.di = (m.alu("add", m.di, 106, 16)) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.w16(m.ds, (m.di+117)&65535, 1)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (m.di+119)&65535, m.ax)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.w8(m.ds, (m.di+121)&65535, 99)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.w8(m.ds, (m.di+122)&65535, 0)
L3f19:
	m.clock(16)
	m.di = ((m.bp + -270) & 65535) & 65535
	m.clock(16)
	m.push(m.ss)
	m.push(m.di)
	m.di = (15686) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.runtime(0x6ea, 0x612)
	m.di = (4482) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x6ea, 0x69f)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.clock(16)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.ax = (21) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.clock(16)
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
	m.ax = (6) & 65535
	m.push(m.ax)
	m.ax = (20) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (13) & 65535
	m.push(m.ax)
	m.ax = (13) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.clock(16)
	m.di = (15706) & 65535
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
	m.push(m.di)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.ax = (41) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.clock(16)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (6) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.clock(16)
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.ax = (m.r16(m.ds, (4470)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (m.di+117)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xe6d)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (41) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = (6056) & 65535
	m.clock(16)
	m.push(m.ds)
	m.push(m.di)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (14) & 65535
	m.push(m.ax)
	m.ax = (14) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.clock(16)
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = (15713) & 65535
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
	m.ax = (15) & 65535
	m.push(m.ax)
	m.ax = (59) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
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
	m.ax = (14) & 65535
	m.push(m.ax)
	m.ax = (58) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.ax = (m.r16(m.ds, (4470)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.clock(16)
	m.di = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (m.di+119)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.push(m.dx)
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xe6d)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.clock(16)
	m.push(m.di)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (59) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.clock(16)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (58) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (14) & 65535
	m.push(m.ax)
	m.ax = (14) & 65535
	m.clock(16)
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = (15720) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.ax = (77) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
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
	m.ax = (6) & 65535
	m.push(m.ax)
	m.ax = (76) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.clock(16)
	m.push(0)
	m.drawText()
	m.ax = (m.r16(m.ds, (4470)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+121)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.clock(16)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xe6d)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (77) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.clock(16)
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (76) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (14) & 65535
	m.push(m.ax)
	m.ax = (14) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = (15727) & 65535
	m.push(m.cs)
	m.clock(16)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (111) & 65535
	m.push(m.ax)
	m.ax = (77) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
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
	m.clock(16)
	m.push(m.ds)
	m.push(m.di)
	m.ax = (110) & 65535
	m.push(m.ax)
	m.ax = (76) & 65535
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
	m.clock(16)
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.ax = (m.r16(m.ds, (4470)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+121)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.push(m.ax)
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xe6d)
	m.di = (15742) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (115) & 65535
	m.push(m.ax)
	m.ax = (87) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
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
	m.ax = (114) & 65535
	m.push(m.ax)
	m.ax = (86) & 65535
	m.clock(16)
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
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.clock(16)
	m.push(0)
	m.drawText()
	m.di = (15755) & 65535
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
	m.push(m.di)
	m.ax = (27) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (111) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (26) & 65535
	m.push(m.ax)
	m.ax = (110) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.di = (15778) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.clock(16)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (23) & 65535
	m.push(m.ax)
	m.ax = (120) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
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
	m.ax = (22) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (119) & 65535
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
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (15802) & 65535
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
	m.push(m.di)
	m.ax = (13) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (133) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.clock(16)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (12) & 65535
	m.push(m.ax)
	m.ax = (132) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.clock(16)
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
	m.w16(m.ds, (5560)&65535, 1)
	goto L4452
L444e:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
L4452:
	m.clock(16)
	m.di = (m.r16(m.ds, (5560)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.w16(m.ss, (m.bp+m.di+-16)&65535, 1)
	m.alu("sub", m.r16(m.ds, (5560)&65535), 6, 16)
	if !m.z {
		goto L444e
	}
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.w16(m.ss, (m.bp+-16)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-16)&65535), 16)
	if !m.z && m.s == m.o {
		goto L4492
	}
	m.w16(m.ds, (5560)&65535, m.ax)
	goto L447b
L4477:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
L447b:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.w8(m.ds, (m.di+122)&65535, 0)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-16)&65535), 16)
	if !m.z {
		goto L4477
	}
L4492:
	m.clock(16)
	m.w16(m.ds, (5528)&65535, 1)
	goto L449e
L449a:
	m.clock(16)
	m.w16(m.ds, (5528)&65535, m.step(m.r16(m.ds, (5528)&65535), 1, 16))
L449e:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, 1)
L44a4:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.alu("sub", m.r8(m.ds, (m.di+122)&65535), 1, 8)
	if !m.z {
		goto L44bb
	}
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
	goto L44a4
L44bb:
	m.clock(16)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.w16(m.ss, (m.bp+-16)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-16)&65535), 16)
	if !m.z && m.s == m.o {
		goto L4511
	}
	m.w16(m.ds, (5558)&65535, m.ax)
	goto L44d2
L44ce:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L44d2:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.clock(16)
	m.di = (m.ax) & 65535
	m.alu("sub", m.r8(m.ds, (m.di+122)&65535), 0, 8)
	if !m.z {
		goto L4509
	}
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.cx = (m.r16(m.ds, (m.di+119)&65535)) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (m.di+119)&65535)) & 65535
	m.alu("sub", m.ax, m.cx, 16)
	if m.z || m.s != m.o {
		goto L4509
	}
	m.clock(16)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.w16(m.ds, (5560)&65535, m.ax)
L4509:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-16)&65535), 16)
	if !m.z {
		goto L44ce
	}
L4511:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.w8(m.ds, (m.di+122)&65535, 1)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.clock(16)
	m.di = (m.r16(m.ds, (5528)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.w16(m.ss, (m.bp+m.di+-16)&65535, m.ax)
	m.alu("sub", m.r16(m.ds, (5528)&65535), 6, 16)
	if m.z {
		goto L4536
	}
	goto L449a
L4536:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, 1)
	goto L4542
L453e:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
L4542:
	m.clock(16)
	m.di = (m.r16(m.ds, (5560)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.ax = (m.r16(m.ss, (m.bp+m.di+-16)&65535)) & 65535
	m.w16(m.ds, (5558)&65535, m.ax)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, 106, 16)) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.clock(16)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.di = (4482) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x6ea, 0x717)
	if !m.z {
		goto L4584
	}
	m.w16(m.ss, (m.bp+-2)&65535, 15)
	goto L4589
L4584:
	m.clock(16)
	m.w16(m.ss, (m.bp+-2)&65535, 14)
L4589:
	m.clock(16)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (9) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 135, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (1) & 65535
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
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (9) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 134, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.push(m.r16(m.ss, (m.bp+-2)&65535))
	m.push(m.r16(m.ss, (m.bp+-2)&65535))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.clock(16)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (m.di+119)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.push(m.dx)
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.clock(16)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xe6d)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (75) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (9) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 135, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
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
	m.clock(16)
	m.push(m.ds)
	m.push(m.di)
	m.ax = (74) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (9) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 134, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.push(m.r16(m.ss, (m.bp+-2)&65535))
	m.push(m.r16(m.ss, (m.bp+-2)&65535))
	m.clock(16)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.alu("sub", m.r16(m.ds, (5560)&65535), 6, 16)
	if m.z {
		goto L467e
	}
	goto L453e
L467e:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, 1)
	goto L468a
L4686:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
L468a:
	m.clock(16)
	m.di = (m.r16(m.ds, (5560)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.w16(m.ss, (m.bp+m.di+-16)&65535, 1)
	m.alu("sub", m.r16(m.ds, (5560)&65535), 6, 16)
	if !m.z {
		goto L4686
	}
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.w16(m.ss, (m.bp+-16)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-16)&65535), 16)
	if !m.z && m.s == m.o {
		goto L46ca
	}
	m.w16(m.ds, (5560)&65535, m.ax)
	goto L46b3
L46af:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
L46b3:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.w8(m.ds, (m.di+122)&65535, 0)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-16)&65535), 16)
	if !m.z {
		goto L46af
	}
L46ca:
	m.clock(16)
	m.w16(m.ds, (5528)&65535, 1)
	m.clock(16)
	goto L46d6
L46d2:
	m.clock(16)
	m.w16(m.ds, (5528)&65535, m.step(m.r16(m.ds, (5528)&65535), 1, 16))
L46d6:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, 1)
L46dc:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.alu("sub", m.r8(m.ds, (m.di+122)&65535), 1, 8)
	if !m.z {
		goto L46f3
	}
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
	goto L46dc
L46f3:
	m.clock(16)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.w16(m.ss, (m.bp+-16)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-16)&65535), 16)
	if !m.z && m.s == m.o {
		goto L4749
	}
	m.w16(m.ds, (5558)&65535, m.ax)
	goto L470a
L4706:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L470a:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.alu("sub", m.r8(m.ds, (m.di+122)&65535), 0, 8)
	if !m.z {
		goto L4741
	}
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.cx = (m.cx & 65280) | ((m.r8(m.ds, (m.di+121)&65535)) & 255)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+121)&65535)) & 255)
	m.alu("sub", (m.ax & 255), (m.cx & 255), 8)
	if !m.c {
		goto L4741
	}
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.w16(m.ds, (5560)&65535, m.ax)
L4741:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-16)&65535), 16)
	if !m.z {
		goto L4706
	}
L4749:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.clock(16)
	m.di = (m.ax) & 65535
	m.w8(m.ds, (m.di+122)&65535, 1)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.di = (m.r16(m.ds, (5528)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.w16(m.ss, (m.bp+m.di+-16)&65535, m.ax)
	m.alu("sub", m.r16(m.ds, (5528)&65535), 6, 16)
	if m.z {
		goto L476e
	}
	goto L46d2
L476e:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, 1)
	goto L477a
L4776:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
L477a:
	m.clock(16)
	m.di = (m.r16(m.ds, (5560)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.ax = (m.r16(m.ss, (m.bp+m.di+-16)&65535)) & 65535
	m.w16(m.ds, (5558)&65535, m.ax)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, 106, 16)) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.clock(16)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.di = (4482) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x6ea, 0x717)
	if !m.z {
		goto L47bc
	}
	m.w16(m.ss, (m.bp+-2)&65535, 15)
	goto L47c1
L47bc:
	m.clock(16)
	m.w16(m.ss, (m.bp+-2)&65535, 14)
L47c1:
	m.clock(16)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (122) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (9) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 135, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
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
	m.ax = (121) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (9) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 134, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.push(m.r16(m.ss, (m.bp+-2)&65535))
	m.push(m.r16(m.ss, (m.bp+-2)&65535))
	m.si = (40960) & 65535
	m.clock(16)
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+121)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.push(m.ax)
	m.clock(16)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xe6d)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (215) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (9) & 65535
	m.clock(16)
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 135, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.clock(16)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (214) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (9) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 134, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.push(m.r16(m.ss, (m.bp+-2)&65535))
	m.push(m.r16(m.ss, (m.bp+-2)&65535))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.alu("sub", m.r16(m.ds, (5560)&65535), 6, 16)
	if m.z {
		goto L48b9
	}
	goto L4776
L48b9:
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.screenTransition()
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
