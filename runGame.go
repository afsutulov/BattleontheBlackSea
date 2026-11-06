package main

// runGame preserves the original Pascal control flow.
func (m *gameState) runGame() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.sp = (m.alu(opSub, m.sp, 256, 16)) & 65535
	m.runtime(0x6ea, 0xd4e)
	m.w16(m.ds, (5518)&65535, 1)
	m.w8(m.ds, (5570)&65535, 7)
	m.w16(m.ds, (5558)&65535, 1)
	goto L51a3
L519f:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L51a3:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.w16(m.ds, (m.di+90)&65535, 1)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 3, 16)
	if !m.z {
		goto L519f
	}
	m.w16(m.ds, (5558)&65535, 4)
	goto L51c2
L51be:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L51c2:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.w16(m.ds, (m.di+90)&65535, 2)
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 7, 16)
	if !m.z {
		goto L51be
	}
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (98)&65535, m.ax)
	m.w16(m.ds, (90)&65535, 4)
	m.push(m.cs)
	m.push(0)
	m.readRecords()
	m.push(m.cs)
	m.push(0)
	m.readPictures()
	m.push(m.cs)
	m.push(0)
	m.initGraphics()
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (5558)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (4478)&65535, m.ax)
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (4480)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (4476)&65535, m.ax)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((2) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((255) & 255)
	m.dx = (975) & 65535
	m.clock(16)
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
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.clock(16)
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
	m.push(m.es)
	m.push(m.di)
	m.ax = (214) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (30) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.di = (20662) & 65535
	m.push(m.cs)
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
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (62) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.clock(16)
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
	m.push(m.ax)
	m.ax = (2) & 65535
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
	m.drawTextCells()
	m.di = (20676) & 65535
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
	m.ax = (81) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
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
	m.drawTextCells()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.clock(16)
	m.di = (20690) & 65535
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
	m.ax = (68) & 65535
	m.push(m.ax)
	m.ax = (101) & 65535
	m.clock(16)
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
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.clock(16)
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
	m.di = (20702) & 65535
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
	m.ax = (76) & 65535
	m.push(m.ax)
	m.ax = (111) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.ax = (76) & 65535
	m.push(m.ax)
	m.ax = (110) & 65535
	m.clock(16)
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
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.clock(16)
	m.push(0)
	m.drawText()
	m.di = (20712) & 65535
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
	m.ax = (57) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.ax = (57) & 65535
	m.push(m.ax)
	m.ax = (180) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.di = (20727) & 65535
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
	m.ax = (33) & 65535
	m.push(m.ax)
	m.ax = (191) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.clock(16)
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
	m.ax = (33) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
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
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
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
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.clock(16)
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
	m.ax = (200) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.inputName()
L5529:
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.drawStatistics()
	m.push(m.cs)
	m.push(0)
	m.drawMap()
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (5558)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (4478)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (4480)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.w16(m.ds, (4476)&65535, m.ax)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (5558)&65535, m.ax)
	goto L555c
L5558:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L555c:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (5560)&65535, m.ax)
	goto L5567
L5563:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
L5567:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5214)&65535, 1)
	m.alu(opSub, m.r16(m.ds, (5560)&65535), 11, 16)
	if !m.z {
		goto L5563
	}
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 11, 16)
	if !m.z {
		goto L5558
	}
	m.w16(m.ds, (5558)&65535, 1)
	goto L55ad
L55a9:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L55ad:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, 1)
	goto L55b9
L55b5:
	m.clock(16)
	m.w16(m.ds, (5560)&65535, m.step(m.r16(m.ds, (5560)&65535), 1, 16))
L55b9:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (24) & 65535
	m.clock(16)
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (m.di+4494)&65535, m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (m.di+4782)&65535, m.ax)
	m.clock(16)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 0)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.clock(16)
	m.w8(m.ds, (m.di+5214)&65535, 0)
	m.alu(opSub, m.r16(m.ds, (5560)&65535), 10, 16)
	if !m.z {
		goto L55b5
	}
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 10, 16)
	if !m.z {
		goto L55a9
	}
	m.w16(m.ds, (5558)&65535, 1)
	goto L5631
L562d:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L5631:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.clock(16)
	m.w8(m.ds, (m.di+5352)&65535, 0)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 20, 16)
	if !m.z {
		goto L562d
	}
	m.w16(m.ds, (5558)&65535, 1)
	goto L5651
L564d:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L5651:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w8(m.ds, (m.di+5353)&65535, 0)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 20, 16)
	m.clock(16)
	if !m.z {
		goto L564d
	}
	m.w16(m.ds, (5558)&65535, 1)
	goto L5671
L566d:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L5671:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.alu(opSub, m.ax, 1, 16)
	if m.z {
		goto L567e
	}
	m.alu(opSub, m.ax, 11, 16)
	if !m.z {
		goto L568d
	}
L567e:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w8(m.ds, (m.di+5351)&65535, 4)
	goto L56f4
L568d:
	m.clock(16)
	m.alu(opSub, m.ax, 2, 16)
	if m.s != m.o {
		goto L5697
	}
	m.alu(opSub, m.ax, 3, 16)
	if m.z || m.s != m.o {
		goto L56a1
	}
L5697:
	m.clock(16)
	m.alu(opSub, m.ax, 12, 16)
	if m.s != m.o {
		goto L56b0
	}
	m.alu(opSub, m.ax, 13, 16)
	m.clock(16)
	if !m.z && m.s == m.o {
		goto L56b0
	}
L56a1:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w8(m.ds, (m.di+5351)&65535, 3)
	goto L56f4
L56b0:
	m.clock(16)
	m.alu(opSub, m.ax, 4, 16)
	if m.s != m.o {
		goto L56ba
	}
	m.alu(opSub, m.ax, 6, 16)
	if m.z || m.s != m.o {
		goto L56c4
	}
L56ba:
	m.clock(16)
	m.alu(opSub, m.ax, 14, 16)
	if m.s != m.o {
		goto L56d3
	}
	m.alu(opSub, m.ax, 16, 16)
	if !m.z && m.s == m.o {
		goto L56d3
	}
L56c4:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w8(m.ds, (m.di+5351)&65535, 2)
	goto L56f4
L56d3:
	m.clock(16)
	m.alu(opSub, m.ax, 7, 16)
	if m.s != m.o {
		goto L56dd
	}
	m.clock(16)
	m.alu(opSub, m.ax, 10, 16)
	if m.z || m.s != m.o {
		goto L56e7
	}
L56dd:
	m.clock(16)
	m.alu(opSub, m.ax, 17, 16)
	if m.s != m.o {
		goto L56f4
	}
	m.alu(opSub, m.ax, 20, 16)
	if !m.z && m.s == m.o {
		goto L56f4
	}
L56e7:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w8(m.ds, (m.di+5351)&65535, 1)
L56f4:
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 20, 16)
	if m.z {
		goto L56fe
	}
	goto L566d
L56fe:
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.placePlayerFleet()
	m.push(m.cs)
	m.push(0)
	m.placeComputerFleet()
	m.temp = m.r16(m.ds, (5572)&65535)
	m.es = m.r16(m.ds, ((5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.clock(16)
	m.push(m.es)
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
	m.temp = m.r16(m.ds, (5616)&65535)
	m.es = m.r16(m.ds, ((5616)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.clock(16)
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (294) & 65535
	m.push(m.ax)
	m.ax = (6) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.push(m.cs)
	m.push(0)
	m.drawPlayerShips()
	m.push(m.cs)
	m.push(0)
	m.screenTransition()
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (5530)&65535, m.ax)
	m.w8(m.ds, (5570)&65535, 7)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (4476)&65535, m.ax)
	m.ax = (m.r16(m.ds, (4470)&65535)) & 65535
	m.dx = (17) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.w16(m.ds, (m.di+117)&65535, m.step(m.r16(m.ds, (m.di+117)&65535), 1, 16))
	m.w16(m.ds, (5526)&65535, 1)
	m.push(m.cs)
	m.push(0)
	m.play()
	m.runtime(0x688, 0x2f3)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 27, 8)
	m.clock(16)
	if !m.z {
		goto L5782
	}
	goto L5c50
L5782:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((2) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((255) & 255)
	m.dx = (975) & 65535
	m.clock(16)
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
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.clock(16)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((3) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.temp = m.r16(m.ds, (5680)&65535)
	m.es = m.r16(m.ds, ((5680)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (212) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (30) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.di = (20748) & 65535
	m.push(m.cs)
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
	m.ax = (4) & 65535
	m.push(m.ax)
	m.ax = (61) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.clock(16)
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
	m.ax = (4) & 65535
	m.push(m.ax)
	m.ax = (60) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
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
	m.drawTextCells()
	m.di = (20752) & 65535
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
	m.ax = (82) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
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
	m.drawTextCells()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (81) & 65535
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
	m.di = (20690) & 65535
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
	m.ax = (10) & 65535
	m.push(m.ax)
	m.ax = (101) & 65535
	m.clock(16)
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
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.clock(16)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.ax = (100) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
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
	m.di = (20702) & 65535
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
	m.ax = (14) & 65535
	m.push(m.ax)
	m.ax = (111) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.ax = (14) & 65535
	m.push(m.ax)
	m.ax = (110) & 65535
	m.clock(16)
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
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.clock(16)
	m.push(0)
	m.drawText()
	m.di = (20712) & 65535
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
	m.clock(16)
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
	m.ax = (7) & 65535
	m.push(m.ax)
	m.ax = (180) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (9) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.di = (20759) & 65535
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
	m.ax = (3) & 65535
	m.push(m.ax)
	m.ax = (191) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.clock(16)
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
	m.ax = (3) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.alu(opSub, m.r16(m.ds, (4474)&65535), 0, 16)
	if m.z {
		goto L5a52
	}
	goto L5bae
L5a52:
	m.clock(16)
	m.ax = (m.r16(m.ds, (4476)&65535)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.clock(16)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xe6d)
	m.di = ((m.bp + -256) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = (20775) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.runtime(0x6ea, 0x612)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x6ea, 0x69f)
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
	m.ax = (5) & 65535
	m.push(m.ax)
	m.ax = (22) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.clock(16)
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
	m.drawTextCells()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (5) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (21) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (5) & 65535
	m.push(m.ax)
	m.ax = (5) & 65535
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
	m.ax = (m.r16(m.ds, (4470)&65535)) & 65535
	m.dx = (17) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.w16(m.ds, (m.di+119)&65535, m.step(m.r16(m.ds, (m.di+119)&65535), 1, 16))
	m.ax = (m.r16(m.ds, (4470)&65535)) & 65535
	m.dx = (17) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+121)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.alu(opSub, m.ax, m.r16(m.ds, (4476)&65535), 16)
	if m.c || m.z {
		goto L5b16
	}
	m.cx = (m.cx & 65280) | ((m.r8(m.ds, (4476)&65535)) & 255)
	m.clock(16)
	m.ax = (m.r16(m.ds, (4470)&65535)) & 65535
	m.dx = (17) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.w8(m.ds, (m.di+121)&65535, (m.cx & 255))
L5b16:
	m.clock(16)
	m.di = (m.r16(m.ds, (90)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.w16(m.ds, (m.di+90)&65535, 2)
	m.ax = (m.r16(m.ds, (90)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.w16(m.ds, (90)&65535, m.ax)
	m.alu(opSub, m.r16(m.ds, (90)&65535), 0, 16)
	if m.z || m.s != m.o {
		goto L5b3c
	}
	m.clock(16)
	m.di = (m.r16(m.ds, (90)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (m.di+90)&65535, m.ax)
L5b3c:
	m.clock(16)
	m.di = (20782) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.clock(16)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
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
	m.drawTextCells()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (12) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (12) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.push(m.cs)
	m.push(0)
	m.screenTransition()
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.runtime(0x688, 0x2f3)
L5bae:
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (4472)&65535), 0, 16)
	if m.z {
		goto L5bb8
	}
	goto L5c50
L5bb8:
	m.clock(16)
	m.di = (m.r16(m.ds, (90)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.w16(m.ds, (m.di+90)&65535, 1)
	m.ax = (m.r16(m.ds, (90)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ds, (90)&65535, m.ax)
	m.alu(opSub, m.r16(m.ds, (90)&65535), 8, 16)
	if m.s == m.o {
		goto L5bde
	}
	m.di = (m.r16(m.ds, (90)&65535)) & 65535
	m.clock(16)
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (m.di+90)&65535, m.ax)
L5bde:
	m.clock(16)
	m.di = (20799) & 65535
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
	m.ax = (4) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
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
	m.drawTextCells()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.push(m.cs)
	m.push(0)
	m.screenTransition()
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.runtime(0x688, 0x2f3)
L5c50:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (5570)&65535)) & 255)
	m.push(m.ax)
	m.di = (20817) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.runtime(0x6ea, 0x8c7)
	if !m.z {
		goto L5c63
	}
	goto L5529
L5c63:
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.exitGame()
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.runtime(0x6ea, 0xd8)
	return
}
