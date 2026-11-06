package main

// drawMap preserves the original Pascal control flow.
func (m *gameState) drawMap() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (2) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu(opSub, m.sp, 2, 16)) & 65535
	m.temp = m.r16(m.ds, (5768)&65535)
	m.es = m.r16(m.ds, ((5768)&65535)+2)
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
	m.di = (19548) & 65535
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
	m.ax = (100) & 65535
	m.push(m.ax)
	m.ax = (122) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.clock(16)
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
	m.drawText()
	m.di = (19558) & 65535
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
	m.ax = (22) & 65535
	m.push(m.ax)
	m.ax = (65) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.drawText()
	m.di = (19565) & 65535
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
	m.ax = (92) & 65535
	m.push(m.ax)
	m.ax = (72) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.clock(16)
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
	m.drawText()
	m.di = (19574) & 65535
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
	m.ax = (100) & 65535
	m.push(m.ax)
	m.ax = (103) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (40) & 65535
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
	m.drawText()
	m.di = (19586) & 65535
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
	m.ax = (160) & 65535
	m.push(m.ax)
	m.ax = (88) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (40) & 65535
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
	m.drawText()
	m.di = (19593) & 65535
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
	m.ax = (220) & 65535
	m.push(m.ax)
	m.ax = (86) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
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
	m.drawText()
	m.clock(16)
	m.di = (19606) & 65535
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
	m.ax = (215) & 65535
	m.push(m.ax)
	m.ax = (124) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
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
	m.clock(16)
	m.push(0)
	m.drawText()
	m.di = (19613) & 65535
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
	m.ax = (231) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (155) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
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
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (19619) & 65535
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
	m.ax = (2) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
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
	m.clock(16)
	m.push(0)
	m.drawTextCells()
	m.di = (19627) & 65535
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
	m.ax = (26) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (35) & 65535
	m.push(m.ax)
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
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.screenTransition()
L4f32:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
	m.w16(m.ds, (5558)&65535, 1)
	goto L4f67
L4f63:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L4f67:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di+90)&65535)) & 65535
	m.alu(opSub, m.ax, 0, 16)
	if !m.z {
		goto L4fc1
	}
	m.ax = (2) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 29, 16)) & 65535
	m.di = (m.ax) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.temp = m.r16(m.ds, (m.di+5572)&65535)
	m.es = m.r16(m.ds, ((m.di+5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, 13, 16)) & 65535
	m.push(m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di+14)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 3, 16)) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	goto L504d
L4fc1:
	m.clock(16)
	m.alu(opSub, m.ax, 1, 16)
	if !m.z {
		goto L5006
	}
	m.ax = (3) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 50, 16)) & 65535
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
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.push(m.r16(m.ds, (m.di)&65535))
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.push(m.r16(m.ds, (m.di+14)&65535))
	m.clock(16)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	goto L504d
L5006:
	m.clock(16)
	m.alu(opSub, m.ax, 2, 16)
	if !m.z {
		goto L504d
	}
	m.ax = (3) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 53, 16)) & 65535
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
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, 20, 16)) & 65535
	m.push(m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.clock(16)
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.push(m.r16(m.ds, (m.di+14)&65535))
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
L504d:
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 7, 16)
	if m.z {
		goto L5057
	}
	goto L4f63
L5057:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.clock(16)
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
	m.runtime(0x688, 0x2fa)
	m.ax = (m.ax & 65280) | ((m.alu(opOr, (m.ax & 255), (m.ax & 255), 8)) & 255)
	if !m.z {
		goto L508c
	}
	goto L4f32
L508c:
	m.clock(16)
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.alu(opSub, m.r16(m.ds, (90)&65535), 0, 16)
	if !m.z {
		goto L50a3
	}
	m.ax = (2) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.clock(16)
	m.push(0)
	m.confirmExit()
L50a3:
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (90)&65535), 8, 16)
	if !m.z {
		goto L50b2
	}
	m.ax = (1) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.confirmExit()
L50b2:
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
