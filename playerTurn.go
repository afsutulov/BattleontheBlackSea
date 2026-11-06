package main

// playerTurn preserves the original Pascal control flow.
func (m *gameState) playerTurn() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (46) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu(opSub, m.sp, 46, 16)) & 65535
	m.di = (6404) & 65535
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
	m.ax = (23) & 65535
	m.push(m.ax)
	m.ax = (182) & 65535
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
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (22) & 65535
	m.push(m.ax)
	m.ax = (181) & 65535
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
	m.di = (6421) & 65535
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
	m.ax = (47) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (191) & 65535
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
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (46) & 65535
	m.push(m.ax)
	m.ax = (190) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6432) & 65535
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
	m.ax = (22) & 65535
	m.push(m.ax)
	m.ax = (187) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
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
	m.ax = (22) & 65535
	m.push(m.ax)
	m.ax = (186) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.ax = (m.r16(m.ds, (4458)&65535)) & 65535
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
	m.clock(16)
	m.ax = (m.r16(m.ds, (4460)&65535)) & 65535
	m.w16(m.ss, (m.bp+-10)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-14)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (5530)&65535, m.ax)
	m.w16(m.ds, (4476)&65535, m.step(m.r16(m.ds, (4476)&65535), 1, 16))
L1a82:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.clock(16)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
	m.push(m.cs)
	m.push(0)
	m.drawOpponentShips()
	m.clock(16)
	m.runtime(0x6ea, 0xcdd)
	m.cx = (26240) & 65535
	m.si = (26214) & 65535
	m.di = (26214) & 65535
	m.runtime(0x6ea, 0xc9d)
	if m.c || m.z {
		goto L1acb
	}
	m.w16(m.ss, (m.bp+-2)&65535, 18)
	goto L1ad0
L1acb:
	m.clock(16)
	m.w16(m.ss, (m.bp+-2)&65535, 17)
L1ad0:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
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
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 150, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.clock(16)
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 17, 16)) & 65535
	m.push(m.ax)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.ax = (m.r16(m.ds, (4476)&65535)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.di = (6056) & 65535
	m.clock(16)
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xe6d)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (35) & 65535
	m.push(m.ax)
	m.ax = (180) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (35) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (179) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.ax = (12) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.clock(16)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
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
	m.clock(16)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
	m.runtime(0x688, 0x2fa)
	m.ax = (m.ax & 65280) | ((m.alu(opOr, (m.ax & 255), (m.ax & 255), 8)) & 255)
	if !m.z {
		goto L1ba3
	}
	goto L2438
L1ba3:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, 1)
	goto L1baf
L1bab:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L1baf:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.product(opImul, m.r16(m.ds, (5558)&65535))
	m.cx = (2) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 20, 16)) & 65535
	m.product(opImul, m.r16(m.ds, (5518)&65535))
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 20, 16)
	if !m.z {
		goto L1bab
	}
	m.runtime(0x688, 0x2f3)
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 0, 8)
	m.clock(16)
	if !m.z {
		goto L1c50
	}
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 72, 8)
	if !m.z {
		goto L1bf5
	}
	m.w16(m.ss, (m.bp+-10)&65535, m.step(m.r16(m.ss, (m.bp+-10)&65535), -1, 16))
L1bf5:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 80, 8)
	if !m.z {
		goto L1bff
	}
	m.w16(m.ss, (m.bp+-10)&65535, m.step(m.r16(m.ss, (m.bp+-10)&65535), 1, 16))
L1bff:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 75, 8)
	if !m.z {
		goto L1c09
	}
	m.clock(16)
	m.w16(m.ss, (m.bp+-8)&65535, m.step(m.r16(m.ss, (m.bp+-8)&65535), -1, 16))
L1c09:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 77, 8)
	if !m.z {
		goto L1c13
	}
	m.w16(m.ss, (m.bp+-8)&65535, m.step(m.r16(m.ss, (m.bp+-8)&65535), 1, 16))
L1c13:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 59, 8)
	if !m.z {
		goto L1c24
	}
	m.ax = (1) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ds, (5518)&65535), 16)) & 65535
	m.w16(m.ds, (5518)&65535, m.ax)
L1c24:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-8)&65535), 1, 16)
	if m.s == m.o {
		goto L1c2f
	}
	m.w16(m.ss, (m.bp+-8)&65535, 10)
L1c2f:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-8)&65535), 10, 16)
	if m.z || m.s != m.o {
		goto L1c3a
	}
	m.w16(m.ss, (m.bp+-8)&65535, 1)
L1c3a:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-10)&65535), 1, 16)
	if m.s == m.o {
		goto L1c45
	}
	m.w16(m.ss, (m.bp+-10)&65535, 10)
L1c45:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-10)&65535), 10, 16)
	m.clock(16)
	if m.z || m.s != m.o {
		goto L1c50
	}
	m.w16(m.ss, (m.bp+-10)&65535, 1)
L1c50:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 13, 8)
	if m.z {
		goto L1c5a
	}
	goto L2438
L1c5a:
	m.clock(16)
	m.w16(m.ss, (m.bp+-24)&65535, 1)
	m.w16(m.ss, (m.bp+-26)&65535, 1)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (5530)&65535, m.ax)
L1c69:
	m.clock(16)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ds, (5558)&65535, m.ax)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ds, (5560)&65535, m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.clock(16)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r16(m.ds, (m.di+4494)&65535), 0, 16)
	if m.z || m.s != m.o {
		goto L1c69
	}
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r16(m.ds, (m.di+4494)&65535), 100, 16)
	if m.s == m.o {
		goto L1c69
	}
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-24)&65535, m.ax)
	m.ax = (m.r16(m.ds, (5560)&65535)) & 65535
	m.w16(m.ss, (m.bp+-26)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 16, 16)) & 65535
	m.w16(m.ss, (m.bp+-24)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-26)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 26, 16)) & 65535
	m.w16(m.ss, (m.bp+-26)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.w16(m.ss, (m.bp+-16)&65535, m.ax)
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-26)&65535)) & 65535
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 154, 16)) & 65535
	m.w16(m.ss, (m.bp+-28)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 18, 16)) & 65535
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-30)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.clock(16)
	m.push(m.dx)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-26)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.push(m.dx)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (137) & 65535
	m.bx = (m.alu(opXor, m.bx, m.bx, 16)) & 65535
	m.dx = (14592) & 65535
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
	m.clock(16)
	m.runtime(0x6ea, 0xc81)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
	m.runtime(0x6ea, 0xc81)
	m.push(m.dx)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-26)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.push(m.dx)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (136) & 65535
	m.clock(16)
	m.bx = (m.alu(opXor, m.bx, m.bx, 16)) & 65535
	m.dx = (14592) & 65535
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
	m.runtime(0x6ea, 0xc81)
	m.push(m.dx)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-28)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-24)&65535), 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.clock(16)
	m.di = (m.pop()) & 65535
	m.runtime(0x6ea, 0xc8d)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
	m.runtime(0x6ea, 0xc93)
	m.runtime(0x6ea, 0xcad)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.w16(m.ss, (m.bp+-6)&65535, 185)
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-14)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L1d9d
L1d9a:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L1d9d:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, 1)
	goto L1da9
L1da5:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L1da9:
	m.clock(16)
	m.ax = (100) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 100, 16)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (4) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-22)&65535), 16)) & 65535
	m.product(opImul, m.r16(m.ds, (5518)&65535))
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.runtime(0x688, 0x2f3)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (2) & 65535
	m.product(opIdiv, m.cx)
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-22)&65535), 16)) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 5, 16)
	if !m.z {
		goto L1da5
	}
	m.ax = (2) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 9, 16)) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (3) & 65535
	m.product(opIdiv, m.cx)
	m.dx = (m.ax) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 2, 16)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.dx, 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.pop()) & 65535
	m.bx = (m.pop()) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.push(m.ax)
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (3) & 65535
	m.product(opIdiv, m.cx)
	m.dx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-26)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 3, 16)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.dx, 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.pop()) & 65535
	m.bx = (m.pop()) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.clock(16)
	m.push(m.ax)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawPlayerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 3, 16)
	if m.z {
		goto L1e68
	}
	goto L1d9a
L1e68:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-6)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-26)&65535), 16)) & 65535
	m.ax = (m.alu(opOr, m.ax, m.ax, 16)) & 65535
	if !m.s {
		goto L1e74
	}
	m.ax = (m.alu(opSub, 0, m.ax, 16)) & 65535
L1e74:
	m.clock(16)
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-4)&65535), 16)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-24)&65535), 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (10619) & 65535
	m.si = (36700) & 65535
	m.di = (30146) & 65535
	m.runtime(0x6ea, 0xc8d)
	m.runtime(0x6ea, 0xcad)
	m.ax = (m.alu(opAdd, m.ax, 1, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, 0, 16)) & 65535
	m.w16(m.ss, (m.bp+-12)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.clock(16)
	m.push(m.dx)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-24)&65535), 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
	m.runtime(0x6ea, 0xc93)
	m.w16(m.ss, (m.bp+-36)&65535, m.ax)
	m.w16(m.ss, (m.bp+-34)&65535, m.bx)
	m.w16(m.ss, (m.bp+-32)&65535, m.dx)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.clock(16)
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.push(m.dx)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-6)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-26)&65535), 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
	m.runtime(0x6ea, 0xc93)
	m.w16(m.ss, (m.bp+-42)&65535, m.ax)
	m.w16(m.ss, (m.bp+-40)&65535, m.bx)
	m.clock(16)
	m.w16(m.ss, (m.bp+-38)&65535, m.dx)
	m.w16(m.ss, (m.bp+-14)&65535, 1)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.w16(m.ss, (m.bp+-44)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)
	if m.z || m.s != m.o {
		goto L1f05
	}
	goto L1fa4
L1f05:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L1f0d
L1f0a:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L1f0d:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5530)&65535)) & 65535
	m.cx = (m.alu(opXor, m.cx, m.cx, 16)) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.dx = (m.ax) & 65535
	m.ax = (200) & 65535
	m.ax = (m.alu(opSub, m.ax, m.dx, 16)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ds, (5518)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.clock(16)
	m.runtime(0x688, 0x2c6)
	m.w16(m.ds, (5530)&65535, m.step(m.r16(m.ds, (5530)&65535), 1, 16))
	m.ax = (22) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.r16(m.ss, (m.bp+-36)&65535)) & 65535
	m.si = (m.r16(m.ss, (m.bp+-34)&65535)) & 65535
	m.di = (m.r16(m.ss, (m.bp+-32)&65535)) & 65535
	m.runtime(0x6ea, 0xc8d)
	m.runtime(0x6ea, 0xcad)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.clock(16)
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.si = (m.r16(m.ss, (m.bp+-40)&65535)) & 65535
	m.di = (m.r16(m.ss, (m.bp+-38)&65535)) & 65535
	m.runtime(0x6ea, 0xc8d)
	m.runtime(0x6ea, 0xcad)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-26)&65535)) & 65535
	m.clock(16)
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 3, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, 0, 16)) & 65535
	m.push(m.ax)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawPlayerCrosshair()
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)
	if m.z {
		goto L1fa4
	}
	goto L1f0a
L1fa4:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L1fae
L1fab:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L1fae:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5530)&65535)) & 65535
	m.cx = (m.alu(opXor, m.cx, m.cx, 16)) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.dx = (m.ax) & 65535
	m.ax = (200) & 65535
	m.ax = (m.alu(opSub, m.ax, m.dx, 16)) & 65535
	m.clock(16)
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ds, (5518)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.w16(m.ds, (5530)&65535, m.step(m.r16(m.ds, (5530)&65535), 1, 16))
	m.ax = (24) & 65535
	m.push(m.ax)
	m.push(m.r16(m.ss, (m.bp+-4)&65535))
	m.ax = (m.r16(m.ss, (m.bp+-6)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 5, 16)) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawPlayerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 1, 16)
	if !m.z {
		goto L1fab
	}
	m.ax = (m.r16(m.ss, (m.bp+-30)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-6)&65535), 16)) & 65535
	m.ax = (m.alu(opOr, m.ax, m.ax, 16)) & 65535
	if !m.s {
		goto L1ffe
	}
	m.ax = (m.alu(opSub, 0, m.ax, 16)) & 65535
L1ffe:
	m.clock(16)
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-28)&65535), 16)) & 65535
	m.clock(16)
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-4)&65535), 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (10619) & 65535
	m.si = (36700) & 65535
	m.di = (30146) & 65535
	m.runtime(0x6ea, 0xc8d)
	m.runtime(0x6ea, 0xcad)
	m.ax = (m.alu(opAdd, m.ax, 1, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, 0, 16)) & 65535
	m.w16(m.ss, (m.bp+-12)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.push(m.dx)
	m.clock(16)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-28)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-4)&65535), 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
	m.runtime(0x6ea, 0xc93)
	m.w16(m.ss, (m.bp+-36)&65535, m.ax)
	m.w16(m.ss, (m.bp+-34)&65535, m.bx)
	m.w16(m.ss, (m.bp+-32)&65535, m.dx)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.clock(16)
	m.runtime(0x6ea, 0xca1)
	m.push(m.dx)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-30)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-6)&65535), 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
	m.runtime(0x6ea, 0xc93)
	m.w16(m.ss, (m.bp+-42)&65535, m.ax)
	m.w16(m.ss, (m.bp+-40)&65535, m.bx)
	m.w16(m.ss, (m.bp+-38)&65535, m.dx)
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.w16(m.ss, (m.bp+-44)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)
	if m.z || m.s != m.o {
		goto L2089
	}
	goto L2128
L2089:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L2091
L208e:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L2091:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5530)&65535)) & 65535
	m.cx = (m.alu(opXor, m.cx, m.cx, 16)) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.dx = (m.ax) & 65535
	m.ax = (200) & 65535
	m.ax = (m.alu(opSub, m.ax, m.dx, 16)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ds, (5518)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.w16(m.ds, (5530)&65535, m.step(m.r16(m.ds, (5530)&65535), 1, 16))
	m.ax = (23) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.r16(m.ss, (m.bp+-36)&65535)) & 65535
	m.si = (m.r16(m.ss, (m.bp+-34)&65535)) & 65535
	m.di = (m.r16(m.ss, (m.bp+-32)&65535)) & 65535
	m.runtime(0x6ea, 0xc8d)
	m.runtime(0x6ea, 0xcad)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.si = (m.r16(m.ss, (m.bp+-40)&65535)) & 65535
	m.di = (m.r16(m.ss, (m.bp+-38)&65535)) & 65535
	m.runtime(0x6ea, 0xc8d)
	m.runtime(0x6ea, 0xcad)
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-6)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.clock(16)
	m.ax = (m.alu(opAdd, m.ax, 3, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, 0, 16)) & 65535
	m.push(m.ax)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawPlayerCrosshair()
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)
	if m.z {
		goto L2128
	}
	goto L208e
L2128:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, 1)
	m.clock(16)
	goto L2134
L2130:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L2134:
	m.clock(16)
	m.ax = (2000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 1000, 16)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ds, (5518)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.clock(16)
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.runtime(0x688, 0x2f3)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (4) & 65535
	m.product(opIdiv, m.cx)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 10, 16)
	m.clock(16)
	if !m.z {
		goto L2130
	}
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r16(m.ds, (m.di+4782)&65535), 0, 16)
	if !m.z && m.s == m.o {
		goto L2197
	}
	goto L2350
L2197:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.clock(16)
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r16(m.ds, (m.di+4782)&65535), 100, 16)
	if m.s != m.o {
		goto L21b4
	}
	goto L234d
L21b4:
	m.clock(16)
	m.w16(m.ds, (4478)&65535, m.step(m.r16(m.ds, (4478)&65535), 1, 16))
	m.w16(m.ds, (4474)&65535, m.step(m.r16(m.ds, (4474)&65535), -1, 16))
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.clock(16)
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.ax = (m.r16(m.ds, (m.di+4782)&65535)) & 65535
	m.w16(m.ds, (5558)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.clock(16)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.ax = (m.r16(m.ds, (m.di+4782)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 100, 16)) & 65535
	m.bx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4782)&65535, m.bx)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.clock(16)
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.di = (m.alu(opAdd, m.di, 5350, 16)) & 65535
	m.w16(m.ss, (m.bp+-46)&65535, m.di)
	m.w16(m.ss, (m.bp+-44)&65535, m.ds)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+3)&65535, m.step(m.r8(m.es, (m.di+3)&65535), -1, 8))
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.alu(opSub, m.r8(m.es, (m.di+3)&65535), 0, 8)
	if m.z {
		goto L222f
	}
	goto L22c9
L222f:
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.alu(opSub, m.r8(m.es, (m.di+2)&65535), 0, 8)
	m.clock(16)
	if !m.z {
		goto L226e
	}
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+4)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 151, 16)) & 65535
	m.w16(m.ds, (5520)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+6)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 13, 16)) & 65535
	m.w16(m.ds, (5522)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+1)&65535)) & 255)
	m.clock(16)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.ax = (m.alu(opAdd, m.ax, 44, 16)) & 65535
	m.w16(m.ds, (5524)&65535, m.ax)
	goto L22a1
L226e:
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+4)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 152, 16)) & 65535
	m.w16(m.ds, (5520)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+6)&65535)) & 65535
	m.cx = (14) & 65535
	m.clock(16)
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 9, 16)) & 65535
	m.w16(m.ds, (5522)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+1)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.ax = (m.alu(opAdd, m.ax, 40, 16)) & 65535
	m.w16(m.ds, (5524)&65535, m.ax)
L22a1:
	m.clock(16)
	m.di = (m.r16(m.ds, (5524)&65535)) & 65535
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
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.r16(m.ds, (5520)&65535))
	m.push(m.r16(m.ds, (5522)&65535))
	m.ax = (1) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
L22c9:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L22d3
L22d0:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L22d3:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, 1)
	goto L22df
L22db:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L22df:
	m.clock(16)
	m.ax = (200) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 100, 16)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (4) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-22)&65535), 16)) & 65535
	m.product(opImul, m.r16(m.ds, (5518)&65535))
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.runtime(0x688, 0x29e)
	m.runtime(0x688, 0x2f3)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (2) & 65535
	m.product(opIdiv, m.cx)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 10, 16)
	if !m.z {
		goto L22db
	}
	m.ax = (26) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-28)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 2, 16)) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-30)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, 7, 16)) & 65535
	m.push(m.ax)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawPlayerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 3, 16)
	if !m.z {
		goto L22d0
	}
L234d:
	m.clock(16)
	goto L23e6
L2350:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L235a
L2357:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L235a:
	m.clock(16)
	m.ax = (25) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-28)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 2, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-30)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, 7, 16)) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawPlayerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 2, 16)
	if !m.z {
		goto L2357
	}
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L2385
L2382:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L2385:
	m.clock(16)
	m.ax = (19) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 159, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 22, 16)) & 65535
	m.push(m.ax)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.drawPlayerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 3, 16)
	if !m.z {
		goto L2382
	}
	m.temp = m.r16(m.ds, (5648)&65535)
	m.es = m.r16(m.ds, ((5648)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.clock(16)
	m.ax = (m.alu(opAdd, m.ax, 159, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 22, 16)) & 65535
	m.push(m.ax)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.w16(m.ds, (5526)&65535, 2)
L23e6:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, 1)
	m.clock(16)
	goto L23f2
L23ee:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L23f2:
	m.clock(16)
	m.ax = (1000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu(opAdd, m.ax, 100, 16)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.ax = (m.r16(m.ds, (5518)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.clock(16)
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.runtime(0x688, 0x2f3)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (4) & 65535
	m.product(opIdiv, m.cx)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 20, 16)
	m.clock(16)
	if !m.z {
		goto L23ee
	}
L2438:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 27, 8)
	if m.z {
		goto L2450
	}
	m.alu(opSub, m.r16(m.ds, (5526)&65535), 2, 16)
	if m.z {
		goto L2450
	}
	m.alu(opSub, m.r16(m.ds, (4474)&65535), 0, 16)
	if m.z {
		goto L2450
	}
	goto L1a82
L2450:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.w16(m.ds, (4458)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.clock(16)
	m.w16(m.ds, (4460)&65535, m.ax)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
