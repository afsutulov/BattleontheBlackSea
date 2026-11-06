package main

// drawSprite preserves the original Pascal control flow.
func (m *gameState) drawSprite() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.push(m.bp)
	m.push(m.ds)
	m.ax = (m.r16(m.ss, (m.bp+6)&65535)) & 65535
	m.w8(m.cs, (23672)&65535, (m.ax & 255))
	m.dx = (m.dx & 65280) | ((m.ax >> 8) & 255)
	m.temp = m.r16(m.ss, (m.bp+16)&65535)
	m.ds = m.r16(m.ss, ((m.bp+16)&65535)+2)
	m.si = (m.temp) & 65535
	m.temp = m.r16(m.ss, (m.bp+12)&65535)
	m.es = m.r16(m.ss, ((m.bp+12)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.r16(m.ds, (m.si)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, false)) & 65535
	m.w16(m.cs, (23673)&65535, m.ax)
	m.clock(16)
	m.si = (m.alu(opAdd, m.si, 2, 16)) & 65535
	m.ax = (m.r16(m.ds, (m.si)&65535)) & 65535
	m.w16(m.cs, (23665)&65535, m.ax)
	m.si = (m.alu(opAdd, m.si, 4, 16)) & 65535
	m.bx = (m.r16(m.ss, (m.bp+8)&65535)) & 65535
	m.alu(opSub, (m.dx & 255), 0, 8)
	if m.z {
		goto L5cb0
	}
	m.ax = (m.shift(m.ax, 1, 16, false)) & 65535
	m.bx = (m.alu(opSub, m.bx, m.ax, 16)) & 65535
L5cb0:
	m.clock(16)
	m.ax = (m.bx) & 65535
	m.cx = (m.cx & 65280) | ((5) & 255)
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.cx = (m.cx & 65280) | ((m.alu(opSub, (m.cx & 255), 2, 8)) & 255)
	m.clock(16)
	m.bx = (m.shift(m.bx, (m.cx & 255), 16, true)) & 65535
	m.bx = (m.alu(opAdd, m.bx, m.ax, 16)) & 65535
	m.ax = (m.r16(m.ss, (m.bp+10)&65535)) & 65535
	m.alu(opSub, (m.dx & 255), 0, 8)
	if m.z {
		goto L5cca
	}
	m.ax = (m.alu(opSub, m.ax, m.r16(m.cs, (23673)&65535), 16)) & 65535
L5cca:
	m.clock(16)
	m.cx = (m.cx & 255) | (((m.ax & 255) & 255) << 8)
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, false)) & 65535
	m.bx = (m.alu(opAdd, m.bx, m.ax, 16)) & 65535
	m.di = (m.alu(opAdd, m.di, m.bx, 16)) & 65535
	m.cx = (m.cx & 255) | (((m.alu(opAnd, (m.cx >> 8), 7, 8)) & 255) << 8)
	m.w8(m.cs, (23671)&65535, (m.cx >> 8))
	m.w8(m.cs, (23675)&65535, (m.cx >> 8))
	m.clock(16)
	m.dx = (974) & 65535
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.videoOut(m.dx, (m.ax & 255))
	m.dx = (m.step(m.dx, 1, 16)) & 65535
	m.ax = (m.ax & 65280) | ((2) & 255)
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.dx = (m.step(m.dx, -1, 16)) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.dx = (m.step(m.dx, 1, 16)) & 65535
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.cs, (23667)&65535, m.ax)
L5cf4:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.w16(m.cs, (23673)&65535, m.ax)
	m.cx = (m.cx & 255) | (((m.r8(m.cs, (23675)&65535)) & 255) << 8)
	m.w8(m.cs, (23671)&65535, (m.cx >> 8))
	m.ax = (m.r16(m.ds, (m.si)&65535)) & 65535
	m.si = (m.alu(opAdd, m.si, 2, 16)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.si, 16)) & 65535
	m.w16(m.cs, (23669)&65535, m.ax)
	m.bx = (m.bx & 65280) | ((m.r8(m.ds, (m.si)&65535)) & 255)
	m.si = (m.step(m.si, 1, 16)) & 65535
L5d12:
	m.clock(16)
	m.alu(opSub, (m.bx & 255), 255, 8)
	if !m.z && m.s == m.o {
		goto L5d73
	}
	m.bx = (m.bx & 65280) | ((m.alu(opSub, (m.bx & 255), 128, 8)) & 255)
	m.bx = (m.bx & 255) | (((m.r8(m.ds, (m.si)&65535)) & 255) << 8)
L5d1c:
	m.clock(16)
	m.alu(opSub, (m.bx & 255), 7, 8)
	if !m.z && m.s == m.o {
		goto L5d28
	}
	m.cx = (m.cx & 65280) | ((m.bx & 255) & 255)
	m.cx = (m.cx & 65280) | ((m.step((m.cx & 255), -1, 8)) & 255)
	goto L5d2a
L5d28:
	m.clock(16)
	m.cx = (m.cx & 65280) | ((7) & 255)
L5d2a:
	m.clock(16)
	m.ax = (32640) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, false)) & 65535
	m.ax = (m.alu(opAnd, m.ax, 255, 16)) & 65535
	m.cx = (m.cx & 255) | (((m.cx & 255) & 255) << 8)
	m.clock(16)
	m.cx = (m.cx & 65280) | ((m.r8(m.cs, (23671)&65535)) & 255)
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, false)) & 65535
	m.alu(opSub, (m.bx >> 8), m.r8(m.cs, (23672)&65535), 8)
	if m.z {
		goto L5d49
	}
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.es, (m.di)&65535, (m.bx >> 8))
L5d49:
	m.clock(16)
	m.ax = (7) & 65535
	m.ax = (m.ax & 65280) | ((m.alu(opSub, (m.ax & 255), (m.cx & 255), 8)) & 255)
	m.alu(opSub, (m.cx >> 8), (m.ax & 255), 8)
	if m.z || m.s != m.o {
		goto L5d54
	}
	m.cx = (m.cx & 255) | (((m.ax & 255) & 255) << 8)
L5d54:
	m.clock(16)
	m.cx = (m.cx & 255) | (((m.step((m.cx >> 8), 1, 8)) & 255) << 8)
	m.cx = (m.cx & 65280) | ((m.alu(opAdd, (m.cx & 255), (m.cx >> 8), 8)) & 255)
	m.alu(opSub, (m.cx & 255), 7, 8)
	if m.z || m.s != m.o {
		goto L5d66
	}
	m.cx = (m.cx & 65280) | ((m.alu(opSub, (m.cx & 255), 8, 8)) & 255)
	m.di = (m.step(m.di, 1, 16)) & 65535
	m.w16(m.cs, (23673)&65535, m.step(m.r16(m.cs, (23673)&65535), 1, 16))
L5d66:
	m.clock(16)
	m.w8(m.cs, (23671)&65535, (m.cx & 255))
	m.bx = (m.bx & 65280) | ((m.alu(opSub, (m.bx & 255), (m.cx >> 8), 8)) & 255)
	if !m.z && m.s == m.o {
		goto L5d1c
	}
	m.si = (m.step(m.si, 1, 16)) & 65535
	goto L5daa
L5d73:
	m.clock(16)
	m.cx = (m.cx & 255) | (((0) & 255) << 8)
	m.cx = (m.cx & 65280) | ((m.r8(m.cs, (23671)&65535)) & 255)
L5d7a:
	m.clock(16)
	m.bx = (m.bx & 255) | (((m.r8(m.ds, (m.si)&65535)) & 255) << 8)
	m.si = (m.step(m.si, 1, 16)) & 65535
	m.ax = (128) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, false)) & 65535
	m.alu(opSub, (m.bx >> 8), m.r8(m.cs, (23672)&65535), 8)
	if m.z {
		goto L5d90
	}
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((m.r8(m.es, (m.di)&65535)) & 255)
	m.w8(m.es, (m.di)&65535, (m.bx >> 8))
L5d90:
	m.clock(16)
	m.cx = (m.cx & 65280) | ((m.step((m.cx & 255), 1, 8)) & 255)
	m.alu(opSub, (m.cx & 255), 7, 8)
	if m.z || m.s != m.o {
		goto L5d9f
	}
	m.di = (m.step(m.di, 1, 16)) & 65535
	m.w16(m.cs, (23673)&65535, m.step(m.r16(m.cs, (23673)&65535), 1, 16))
	m.cx = (m.cx & 65280) | ((0) & 255)
L5d9f:
	m.clock(16)
	m.cx = (m.cx & 255) | (((m.step((m.cx >> 8), 1, 8)) & 255) << 8)
	m.alu(opSub, (m.bx & 255), (m.cx >> 8), 8)
	if !m.z {
		goto L5d7a
	}
	m.w8(m.cs, (23671)&65535, (m.cx & 255))
L5daa:
	m.clock(16)
	m.bx = (m.bx & 65280) | ((m.r8(m.ds, (m.si)&65535)) & 255)
	m.si = (m.step(m.si, 1, 16)) & 65535
	m.clock(16)
	m.ax = (m.r16(m.cs, (23669)&65535)) & 65535
	m.alu(opSub, m.ax, m.si, 16)
	if m.z {
		goto L5db8
	}
	goto L5d12
L5db8:
	m.clock(16)
	m.ax = (m.r16(m.cs, (23673)&65535)) & 65535
	m.di = (m.alu(opSub, m.di, m.r16(m.cs, (23673)&65535), 16)) & 65535
	m.di = (m.alu(opAdd, m.di, 40, 16)) & 65535
	m.ax = (m.r16(m.cs, (23667)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.cs, (23667)&65535, m.ax)
	m.alu(opSub, m.ax, m.r16(m.cs, (23665)&65535), 16)
	if m.s == m.o {
		goto L5dd7
	}
	goto L5cf4
L5dd7:
	m.clock(16)
	m.ds = (m.pop()) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 18) & 65535
	return
}
