package main

// confirmExit preserves the original Pascal control flow.
func (m *gameState) confirmExit() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.runtime(0x6ea, 0x244)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((2) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.clock(16)
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
	m.si = (40960) & 65535
	m.clock(16)
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
	m.di = (48000) & 65535
	m.clock(16)
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
	m.di = (18633) & 65535
	m.push(m.cs)
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
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (62) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
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
	m.ax = (61) & 65535
	m.clock(16)
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
	m.di = (18647) & 65535
	m.clock(16)
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
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (81) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
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
	m.clock(16)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (80) & 65535
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
	m.clock(16)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = (18661) & 65535
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
	m.clock(16)
	m.ax = (68) & 65535
	m.push(m.ax)
	m.ax = (101) & 65535
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
	m.di = (48000) & 65535
	m.clock(16)
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (68) & 65535
	m.push(m.ax)
	m.ax = (100) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.di = (18673) & 65535
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
	m.ax = (76) & 65535
	m.push(m.ax)
	m.ax = (111) & 65535
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
	m.ax = (76) & 65535
	m.push(m.ax)
	m.ax = (110) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
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
	m.clock(16)
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (18683) & 65535
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
	m.ax = (57) & 65535
	m.push(m.ax)
	m.ax = (181) & 65535
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
	m.ax = (57) & 65535
	m.push(m.ax)
	m.ax = (180) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
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
	m.di = (18698) & 65535
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
	m.ax = (33) & 65535
	m.push(m.ax)
	m.ax = (191) & 65535
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
	m.ax = (33) & 65535
	m.push(m.ax)
	m.ax = (190) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
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
	m.push(m.cs)
	m.push(0)
	m.screenTransition()
	m.ax = (500) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.temp = m.r16(m.ds, (5680)&65535)
	m.es = m.r16(m.ds, ((5680)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.clock(16)
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (214) & 65535
	m.push(m.ax)
	m.ax = (30) & 65535
	m.push(m.ax)
	m.ax = (16) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.ax = (200) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.temp = m.r16(m.ds, (5692)&65535)
	m.es = m.r16(m.ds, ((5692)&65535)+2)
	m.di = (m.temp) & 65535
	m.clock(16)
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (214) & 65535
	m.push(m.ax)
	m.ax = (30) & 65535
	m.push(m.ax)
	m.ax = (16) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.clock(16)
	m.ax = (200) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 6) & 65535
	return
}
