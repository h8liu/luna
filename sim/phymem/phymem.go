package phymem

import (
	. "github.com/h8liu/luna/sim/consts"
	. "github.com/h8liu/luna/sim/phypage"
)

// Physical memory
type PhyMemory struct {
	npage uint32
	pages map[uint32]*PhyPage
}

func NewPhyMemory(size uint32) *PhyMemory {
	m := new(PhyMemory)
	if size%PageSize != 0 {
		panic("memory size not aligned to page size")
	}

	m.npage = size / PageSize
	m.pages = make(map[uint32]*PhyPage)

	return m
}

func pageAddr(paddr uint32) (uint32, uint32) {
	return paddr >> PageNbit, paddr & PageMask
}

func (m *PhyMemory) page(index uint32) *PhyPage {
	if index >= m.npage {
		panic("physical memory address out of range")
	}
	p := m.pages[index]
	if p == nil {
		p = NewPhyPage()
		m.pages[index] = p
	}
	return p
}

func (m *PhyMemory) pageOff(paddr uint32) (*PhyPage, uint32) {
	index, off := pageAddr(paddr)
	return m.page(index), off
}

func (m *PhyMemory) ReadU8(paddr uint32) uint32 {
	p, off := m.pageOff(paddr)
	return p.ReadU8(off)
}

func (m *PhyMemory) WriteU8(paddr, v uint32) {
	p, off := m.pageOff(paddr)
	p.WriteU8(off, v)
}

func (m *PhyMemory) ReadU32(paddr uint32) uint32 {
	p, off := m.pageOff(paddr)
	return p.ReadU32(off)
}

func (m *PhyMemory) WriteU32(paddr, v uint32) {
	p, off := m.pageOff(paddr)
	p.WriteU32(off, v)
}
