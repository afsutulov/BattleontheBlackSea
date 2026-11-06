package main

// computerTurn preserves the original Pascal control flow.
func (m *gameState) computerTurn() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (46) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu(opSub, m.sp, 46, 16)) & 65535
L263a:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-14)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (5530)&65535, m.ax)
	m.runtime(0x688, 0x2fa)
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.alu(opOr, (m.ax & 255), (m.ax & 255), 8)) & 255)
	if m.z {
		goto L2675
	}
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 0, 8)
	if !m.z {
		goto L2675
	}
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 59, 8)
	if !m.z {
		goto L2675
	}
	m.ax = (1) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ds, (5518)&65535), 16)) & 65535
	m.w16(m.ds, (5518)&65535, m.ax)
L2675:
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (4466)&65535), 0, 16)
	if !m.z {
		goto L26b1
	}
L267c:
	m.clock(16)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ss, (m.bp+-10)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.cx = (m.ax) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r8(m.ds, (m.di+5070)&65535), 0, 8)
	if !m.z {
		goto L267c
	}
	goto L2778
L26b1:
	m.clock(16)
	m.ax = (m.r16(m.ds, (4464)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ds, (4462)&65535)) & 65535
	m.dx = (24) & 65535
	m.clock(16)
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.ax = (m.r16(m.ds, (m.di+4494)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (100) & 65535
	m.product(opIdiv, m.cx)
	m.temp = m.dx
	m.dx = (m.ax) & 65535
	m.ax = (m.temp) & 65535
	m.w16(m.ds, (5558)&65535, m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.alu(opSub, m.r8(m.ds, (m.di+5352)&65535), 0, 8)
	if !m.z {
		goto L2720
	}
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.clock(16)
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.push(m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di+5354)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.pop()) & 65535
	m.clock(16)
	m.bx = (m.pop()) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di+5356)&65535)) & 65535
	m.w16(m.ss, (m.bp+-10)&65535, m.ax)
	goto L275d
L2720:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.push(m.dx)
	m.push(m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di+5356)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.pop()) & 65535
	m.bx = (m.pop()) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.clock(16)
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.w16(m.ss, (m.bp+-10)&65535, m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di+5354)&65535)) & 65535
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
L275d:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.clock(16)
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r8(m.ds, (m.di+5070)&65535), 0, 8)
	if m.z {
		goto L2778
	}
	goto L26b1
L2778:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
	m.w16(m.ss, (m.bp+-24)&65535, 294)
	m.clock(16)
	m.w16(m.ss, (m.bp+-26)&65535, 20)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.w16(m.ss, (m.bp+-16)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-26)&65535)) & 65535
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opSub, m.ax, 3, 16)) & 65535
	m.w16(m.ss, (m.bp+-28)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.alu(opAdd, m.ax, 18, 16)) & 65535
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-30)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
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
	m.clock(16)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
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
	m.clock(16)
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (136) & 65535
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
	m.clock(16)
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
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
	m.clock(16)
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.w16(m.ss, (m.bp+-6)&65535, 185)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-14)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L2856
L2853:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L2856:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, 1)
	goto L2862
L285e:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L2862:
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
	m.clock(16)
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
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
	m.clock(16)
	m.alu(opSub, m.r16(m.ds, (5558)&65535), 5, 16)
	if !m.z {
		goto L285e
	}
	m.ax = (2) & 65535
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
	m.clock(16)
	m.cx = (3) & 65535
	m.product(opIdiv, m.cx)
	m.dx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, 2, 16)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.dx, 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (m.pop()) & 65535
	m.bx = (m.pop()) & 65535
	m.ax = (m.alu(opSub, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opSbb, m.dx, m.bx, 16)) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.clock(16)
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
	m.clock(16)
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawComputerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 3, 16)
	if m.z {
		goto L2921
	}
	goto L2853
L2921:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-24)&65535), 16)) & 65535
	m.ax = (m.alu(opOr, m.ax, m.ax, 16)) & 65535
	m.clock(16)
	if !m.s {
		goto L292d
	}
	m.ax = (m.alu(opSub, 0, m.ax, 16)) & 65535
L292d:
	m.clock(16)
	m.dx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-6)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-26)&65535), 16)) & 65535
	m.ax = (m.alu(opOr, m.ax, m.ax, 16)) & 65535
	if !m.s {
		goto L293b
	}
	m.ax = (m.alu(opSub, 0, m.ax, 16)) & 65535
L293b:
	m.clock(16)
	m.ax = (m.alu(opAdd, m.ax, m.dx, 16)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.clock(16)
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
	m.push(m.bx)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.clock(16)
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
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.push(m.dx)
	m.push(m.bx)
	m.clock(16)
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
	m.w16(m.ss, (m.bp+-38)&65535, m.dx)
	m.w16(m.ss, (m.bp+-14)&65535, 1)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-44)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)
	if m.z || m.s != m.o {
		goto L29c8
	}
	goto L2a68
L29c8:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L29d0
L29cd:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L29d0:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5530)&65535)) & 65535
	m.clock(16)
	m.cx = (2) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.dx = (m.ax) & 65535
	m.ax = (400) & 65535
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
	m.ax = (34) & 65535
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
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
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
	m.ax = (m.r16(m.ss, (m.bp+-26)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.ax = (m.alu(opAdd, m.ax, m.cx, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, m.bx, 16)) & 65535
	m.clock(16)
	m.ax = (m.alu(opSub, m.ax, 5, 16)) & 65535
	m.dx = (m.alu(opSbb, m.dx, 0, 16)) & 65535
	m.push(m.ax)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawComputerCrosshair()
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)
	if m.z {
		goto L2a68
	}
	goto L29cd
L2a68:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L2a72
L2a6f:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L2a72:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5530)&65535)) & 65535
	m.cx = (2) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.dx = (m.ax) & 65535
	m.ax = (400) & 65535
	m.ax = (m.alu(opSub, m.ax, m.dx, 16)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.clock(16)
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
	m.ax = (36) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 2, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-6)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 3, 16)) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawComputerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 1, 16)
	if !m.z {
		goto L2a6f
	}
	m.ax = (m.r16(m.ss, (m.bp+-28)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-4)&65535), 16)) & 65535
	m.ax = (m.alu(opOr, m.ax, m.ax, 16)) & 65535
	if !m.s {
		goto L2ac7
	}
	m.ax = (m.alu(opSub, 0, m.ax, 16)) & 65535
L2ac7:
	m.clock(16)
	m.dx = (m.ax) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-30)&65535)) & 65535
	m.ax = (m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-6)&65535), 16)) & 65535
	m.ax = (m.alu(opOr, m.ax, m.ax, 16)) & 65535
	if !m.s {
		goto L2ad5
	}
	m.ax = (m.alu(opSub, 0, m.ax, 16)) & 65535
L2ad5:
	m.clock(16)
	m.ax = (m.alu(opAdd, m.ax, m.dx, 16)) & 65535
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
	m.clock(16)
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
	m.clock(16)
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
	m.clock(16)
	m.cx = (m.pop()) & 65535
	m.si = (m.pop()) & 65535
	m.di = (m.pop()) & 65535
	m.runtime(0x6ea, 0xc93)
	m.w16(m.ss, (m.bp+-42)&65535, m.ax)
	m.w16(m.ss, (m.bp+-40)&65535, m.bx)
	m.w16(m.ss, (m.bp+-38)&65535, m.dx)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.w16(m.ss, (m.bp+-44)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)
	if m.z || m.s != m.o {
		goto L2b5c
	}
	goto L2c02
L2b5c:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L2b64
L2b61:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L2b64:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5530)&65535)) & 65535
	m.cx = (2) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.dx = (m.ax) & 65535
	m.ax = (400) & 65535
	m.ax = (m.alu(opSub, m.ax, m.dx, 16)) & 65535
	m.dx = (m.alu(opXor, m.dx, m.dx, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.bx = (m.dx) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ds, (5518)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.w16(m.ds, (5530)&65535, m.step(m.r16(m.ds, (5530)&65535), 1, 16))
	m.ax = (35) & 65535
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
	m.clock(16)
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
	m.ax = (m.alu(opAdd, m.ax, 4, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, 0, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.runtime(0x6ea, 0xca1)
	m.cx = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.si = (m.r16(m.ss, (m.bp+-40)&65535)) & 65535
	m.clock(16)
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
	m.ax = (m.alu(opAdd, m.ax, 1, 16)) & 65535
	m.dx = (m.alu(opAdc, m.dx, 0, 16)) & 65535
	m.push(m.ax)
	m.ax = (11) & 65535
	m.push(m.ax)
	m.push(m.bp)
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.drawComputerCrosshair()
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)
	if m.z {
		goto L2c02
	}
	goto L2b61
L2c02:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, 1)
	goto L2c0e
L2c0a:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L2c0e:
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
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.runtime(0x688, 0x2f3)
	m.clock(16)
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
	if !m.z {
		goto L2c0a
	}
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.clock(16)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r16(m.ds, (m.di+4494)&65535), 0, 16)
	if !m.z && m.s == m.o {
		goto L2c71
	}
	goto L2ebb
L2c71:
	m.clock(16)
	m.w16(m.ds, (4466)&65535, 1)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.w16(m.ds, (4462)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.w16(m.ds, (4464)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r16(m.ds, (m.di+4494)&65535), 100, 16)
	if m.s != m.o {
		goto L2ca0
	}
	goto L2eb8
L2ca0:
	m.clock(16)
	m.w16(m.ds, (4480)&65535, m.step(m.r16(m.ds, (4480)&65535), 1, 16))
	m.w16(m.ds, (4472)&65535, m.step(m.r16(m.ds, (4472)&65535), -1, 16))
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.ax = (m.r16(m.ds, (m.di+4494)&65535)) & 65535
	m.w16(m.ds, (5558)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ds, (m.di+4494)&65535)) & 65535
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
	m.w16(m.ds, (m.di+4494)&65535, m.bx)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.clock(16)
	m.di = (m.alu(opAdd, m.di, 5350, 16)) & 65535
	m.w16(m.ss, (m.bp+-46)&65535, m.di)
	m.w16(m.ss, (m.bp+-44)&65535, m.ds)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.clock(16)
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.clock(16)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+3)&65535, m.step(m.r8(m.es, (m.di+3)&65535), -1, 8))
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.alu(opSub, m.r8(m.es, (m.di+3)&65535), 0, 8)
	if m.z {
		goto L2d7b
	}
	goto L2e34
L2d7b:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (4466)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.alu(opSub, m.r8(m.es, (m.di+2)&65535), 0, 8)
	if m.z {
		goto L2d94
	}
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.alu(opSub, m.r8(m.es, (m.di+1)&65535), 1, 8)
	if !m.z {
		goto L2dda
	}
L2d94:
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+6)&65535)) & 65535
	m.cx = (m.ax) & 65535
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+4)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+6)&65535)) & 65535
	m.cx = (m.ax) & 65535
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+1)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.es, (m.di+4)&65535), 16)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
L2dda:
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.alu(opSub, m.r8(m.es, (m.di+2)&65535), 1, 8)
	if m.z {
		goto L2dee
	}
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.clock(16)
	m.alu(opSub, m.r8(m.es, (m.di+1)&65535), 1, 8)
	if !m.z {
		goto L2e34
	}
L2dee:
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+6)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.cx = (m.ax) & 65535
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+4)&65535)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di+1)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.es, (m.di+6)&65535), 16)) & 65535
	m.cx = (m.ax) & 65535
	m.temp = m.r16(m.ss, (m.bp+-46)&65535)
	m.es = m.r16(m.ss, ((m.bp+-46)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.es, (m.di+4)&65535)) & 65535
	m.dx = (12) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w8(m.ds, (m.di+5070)&65535, 1)
L2e34:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L2e3e
L2e3b:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L2e3e:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, 1)
	goto L2e4a
L2e46:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L2e4a:
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
		goto L2e46
	}
	m.ax = (33) & 65535
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
	m.drawComputerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 3, 16)
	if !m.z {
		goto L2e3b
	}
L2eb8:
	m.clock(16)
	goto L2fa3
L2ebb:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L2ec5
L2ec2:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L2ec5:
	m.clock(16)
	m.ax = (32) & 65535
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
	m.drawComputerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 2, 16)
	if !m.z {
		goto L2ec2
	}
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	goto L2ef0
L2eed:
	m.clock(16)
	m.w16(m.ss, (m.bp+-22)&65535, m.step(m.r16(m.ss, (m.bp+-22)&65535), 1, 16))
L2ef0:
	m.clock(16)
	m.ax = (19) & 65535
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.cx = (14) & 65535
	m.product(opImul, m.cx)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
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
	m.drawComputerCrosshair()
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 3, 16)
	if !m.z {
		goto L2eed
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
	m.ax = (m.step(m.ax, 1, 16)) & 65535
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
	m.w16(m.ds, (5558)&65535, 1)
	goto L2f53
L2f4f:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L2f53:
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
	m.runtime(0x6ea, 0x261)
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.ax = (2) & 65535
	m.clock(16)
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
	if !m.z {
		goto L2f4f
	}
	m.push(m.cs)
	m.push(0)
	m.drawPlayerShips()
	m.w16(m.ds, (5526)&65535, 1)
L2fa3:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.w16(m.ds, (4462)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.w16(m.ds, (4464)&65535, m.ax)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 27, 8)
	if m.z {
		goto L2fc7
	}
	m.alu(opSub, m.r16(m.ds, (5526)&65535), 1, 16)
	if m.z {
		goto L2fc7
	}
	m.alu(opSub, m.r16(m.ds, (4472)&65535), 0, 16)
	if m.z {
		goto L2fc7
	}
	goto L263a
L2fc7:
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.clock(16)
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
