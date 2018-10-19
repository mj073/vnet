// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=pg -id buffer_type_pool -d PoolType=buffer_type_pool -d Type=buffer_type -d Data=elts github.com/platinasystems/elib/pool.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pg

import (
	"github.com/platinasystems/elib"
)

type buffer_type_pool struct {
	elib.Pool
	elts []buffer_type
}

func (p *buffer_type_pool) GetIndex() (i uint) {
	l := uint(len(p.elts))
	i = p.Pool.GetIndex(l)
	if i >= l {
		p.Validate(i)
	}
	return i
}

func (p *buffer_type_pool) PutIndex(i uint) (ok bool) {
	return p.Pool.PutIndex(i)
}

func (p *buffer_type_pool) IsFree(i uint) (v bool) {
	v = i >= uint(len(p.elts))
	if !v {
		v = p.Pool.IsFree(i)
	}
	return
}

func (p *buffer_type_pool) Resize(n uint) {
	c := uint(cap(p.elts))
	l := uint(len(p.elts) + int(n))
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]buffer_type, l, c)
		copy(q, p.elts)
		p.elts = q
	}
	p.elts = p.elts[:l]
}

func (p *buffer_type_pool) Validate(i uint) {
	c := uint(cap(p.elts))
	l := uint(i) + 1
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]buffer_type, l, c)
		copy(q, p.elts)
		p.elts = q
	}
	if l > uint(len(p.elts)) {
		p.elts = p.elts[:l]
	}
}

func (p *buffer_type_pool) Elts() uint {
	return uint(len(p.elts)) - p.FreeLen()
}

func (p *buffer_type_pool) Len() uint {
	return uint(len(p.elts))
}

func (p *buffer_type_pool) Foreach(f func(x buffer_type)) {
	for i := range p.elts {
		if !p.Pool.IsFree(uint(i)) {
			f(p.elts[i])
		}
	}
}

func (p *buffer_type_pool) ForeachIndex(f func(i uint)) {
	for i := range p.elts {
		if !p.Pool.IsFree(uint(i)) {
			f(uint(i))
		}
	}
}

func (p *buffer_type_pool) Reset() {
	p.Pool.Reset()
	if len(p.elts) > 0 {
		p.elts = p.elts[:0]
	}
}
