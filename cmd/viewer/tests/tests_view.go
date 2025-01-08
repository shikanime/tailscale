// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by tailscale/cmd/viewer; DO NOT EDIT.

package tests

import (
	"encoding/json"
	"errors"
	"net/netip"

	"golang.org/x/exp/constraints"
	"tailscale.com/types/views"
)

//go:generate go run tailscale.com/cmd/cloner  -clonefunc=false -type=StructWithPtrs,StructWithoutPtrs,Map,StructWithSlices,OnlyGetClone,StructWithEmbedded,GenericIntStruct,GenericNoPtrsStruct,GenericCloneableStruct,StructWithContainers,StructWithTypeAliasFields,GenericTypeAliasStruct

// View returns a readonly view of StructWithPtrs.
func (p *StructWithPtrs) View() StructWithPtrsView {
	return StructWithPtrsView{ж: p}
}

// StructWithPtrsView provides a read-only view over StructWithPtrs.
//
// Its methods should only be called if `Valid()` returns true.
type StructWithPtrsView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *StructWithPtrs
}

// Valid reports whether underlying value is non-nil.
func (v StructWithPtrsView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v StructWithPtrsView) AsStruct() *StructWithPtrs {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v StructWithPtrsView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *StructWithPtrsView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x StructWithPtrs
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v StructWithPtrsView) Value() StructWithoutPtrsView { return v.ж.Value.View() }
func (v StructWithPtrsView) Int() views.ValuePointer[int] { return views.ValuePointerOf(v.ж.Int) }

func (v StructWithPtrsView) NoView() views.ValuePointer[StructWithNoView] {
	return views.ValuePointerOf(v.ж.NoView)
}

func (v StructWithPtrsView) NoCloneValue() *StructWithoutPtrs { return v.ж.NoCloneValue }
func (v StructWithPtrsView) String() string                   { return v.ж.String() }
func (v StructWithPtrsView) Equal(v2 StructWithPtrsView) bool { return v.ж.Equal(v2.ж) }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _StructWithPtrsViewNeedsRegeneration = StructWithPtrs(struct {
	Value        *StructWithoutPtrs
	Int          *int
	NoView       *StructWithNoView
	NoCloneValue *StructWithoutPtrs
}{})

// View returns a readonly view of StructWithoutPtrs.
func (p *StructWithoutPtrs) View() StructWithoutPtrsView {
	return StructWithoutPtrsView{ж: p}
}

// StructWithoutPtrsView provides a read-only view over StructWithoutPtrs.
//
// Its methods should only be called if `Valid()` returns true.
type StructWithoutPtrsView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *StructWithoutPtrs
}

// Valid reports whether underlying value is non-nil.
func (v StructWithoutPtrsView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v StructWithoutPtrsView) AsStruct() *StructWithoutPtrs {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v StructWithoutPtrsView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *StructWithoutPtrsView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x StructWithoutPtrs
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v StructWithoutPtrsView) Int() int          { return v.ж.Int }
func (v StructWithoutPtrsView) Pfx() netip.Prefix { return v.ж.Pfx }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _StructWithoutPtrsViewNeedsRegeneration = StructWithoutPtrs(struct {
	Int int
	Pfx netip.Prefix
}{})

// View returns a readonly view of Map.
func (p *Map) View() MapView {
	return MapView{ж: p}
}

// MapView provides a read-only view over Map.
//
// Its methods should only be called if `Valid()` returns true.
type MapView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *Map
}

// Valid reports whether underlying value is non-nil.
func (v MapView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v MapView) AsStruct() *Map {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v MapView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *MapView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x Map
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v MapView) Int() views.Map[string, int] { return views.MapOf(v.ж.Int) }

func (v MapView) SliceInt() views.MapSlice[string, int] { return views.MapSliceOf(v.ж.SliceInt) }

func (v MapView) StructPtrWithPtr() views.MapFn[string, *StructWithPtrs, StructWithPtrsView] {
	return views.MapFnOf(v.ж.StructPtrWithPtr, func(t *StructWithPtrs) StructWithPtrsView {
		return t.View()
	})
}

func (v MapView) StructPtrWithoutPtr() views.MapFn[string, *StructWithoutPtrs, StructWithoutPtrsView] {
	return views.MapFnOf(v.ж.StructPtrWithoutPtr, func(t *StructWithoutPtrs) StructWithoutPtrsView {
		return t.View()
	})
}

func (v MapView) StructWithoutPtr() views.Map[string, StructWithoutPtrs] {
	return views.MapOf(v.ж.StructWithoutPtr)
}

func (v MapView) SlicesWithPtrs() views.MapFn[string, []*StructWithPtrs, views.SliceView[*StructWithPtrs, StructWithPtrsView]] {
	return views.MapFnOf(v.ж.SlicesWithPtrs, func(t []*StructWithPtrs) views.SliceView[*StructWithPtrs, StructWithPtrsView] {
		return views.SliceOfViews[*StructWithPtrs, StructWithPtrsView](t)
	})
}

func (v MapView) SlicesWithoutPtrs() views.MapFn[string, []*StructWithoutPtrs, views.SliceView[*StructWithoutPtrs, StructWithoutPtrsView]] {
	return views.MapFnOf(v.ж.SlicesWithoutPtrs, func(t []*StructWithoutPtrs) views.SliceView[*StructWithoutPtrs, StructWithoutPtrsView] {
		return views.SliceOfViews[*StructWithoutPtrs, StructWithoutPtrsView](t)
	})
}

func (v MapView) StructWithoutPtrKey() views.Map[StructWithoutPtrs, int] {
	return views.MapOf(v.ж.StructWithoutPtrKey)
}

func (v MapView) StructWithPtr() views.MapFn[string, StructWithPtrs, StructWithPtrsView] {
	return views.MapFnOf(v.ж.StructWithPtr, func(t StructWithPtrs) StructWithPtrsView {
		return t.View()
	})
}
func (v MapView) SliceIntPtr() map[string][]*int           { panic("unsupported") }
func (v MapView) PointerKey() map[*string]int              { panic("unsupported") }
func (v MapView) StructWithPtrKey() map[StructWithPtrs]int { panic("unsupported") }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _MapViewNeedsRegeneration = Map(struct {
	Int                 map[string]int
	SliceInt            map[string][]int
	StructPtrWithPtr    map[string]*StructWithPtrs
	StructPtrWithoutPtr map[string]*StructWithoutPtrs
	StructWithoutPtr    map[string]StructWithoutPtrs
	SlicesWithPtrs      map[string][]*StructWithPtrs
	SlicesWithoutPtrs   map[string][]*StructWithoutPtrs
	StructWithoutPtrKey map[StructWithoutPtrs]int
	StructWithPtr       map[string]StructWithPtrs
	SliceIntPtr         map[string][]*int
	PointerKey          map[*string]int
	StructWithPtrKey    map[StructWithPtrs]int
}{})

// View returns a readonly view of StructWithSlices.
func (p *StructWithSlices) View() StructWithSlicesView {
	return StructWithSlicesView{ж: p}
}

// StructWithSlicesView provides a read-only view over StructWithSlices.
//
// Its methods should only be called if `Valid()` returns true.
type StructWithSlicesView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *StructWithSlices
}

// Valid reports whether underlying value is non-nil.
func (v StructWithSlicesView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v StructWithSlicesView) AsStruct() *StructWithSlices {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v StructWithSlicesView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *StructWithSlicesView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x StructWithSlices
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v StructWithSlicesView) Values() views.Slice[StructWithoutPtrs] {
	return views.SliceOf(v.ж.Values)
}
func (v StructWithSlicesView) ValuePointers() views.SliceView[*StructWithoutPtrs, StructWithoutPtrsView] {
	return views.SliceOfViews[*StructWithoutPtrs, StructWithoutPtrsView](v.ж.ValuePointers)
}
func (v StructWithSlicesView) StructPointers() views.SliceView[*StructWithPtrs, StructWithPtrsView] {
	return views.SliceOfViews[*StructWithPtrs, StructWithPtrsView](v.ж.StructPointers)
}
func (v StructWithSlicesView) Slice() views.Slice[string] { return views.SliceOf(v.ж.Slice) }
func (v StructWithSlicesView) Prefixes() views.Slice[netip.Prefix] {
	return views.SliceOf(v.ж.Prefixes)
}
func (v StructWithSlicesView) Data() views.ByteSlice[[]byte] { return views.ByteSliceOf(v.ж.Data) }
func (v StructWithSlicesView) Structs() StructWithPtrs       { panic("unsupported") }
func (v StructWithSlicesView) Ints() *int                    { panic("unsupported") }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _StructWithSlicesViewNeedsRegeneration = StructWithSlices(struct {
	Values         []StructWithoutPtrs
	ValuePointers  []*StructWithoutPtrs
	StructPointers []*StructWithPtrs
	Slice          []string
	Prefixes       []netip.Prefix
	Data           []byte
	Structs        []StructWithPtrs
	Ints           []*int
}{})

// View returns a readonly view of StructWithEmbedded.
func (p *StructWithEmbedded) View() StructWithEmbeddedView {
	return StructWithEmbeddedView{ж: p}
}

// StructWithEmbeddedView provides a read-only view over StructWithEmbedded.
//
// Its methods should only be called if `Valid()` returns true.
type StructWithEmbeddedView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *StructWithEmbedded
}

// Valid reports whether underlying value is non-nil.
func (v StructWithEmbeddedView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v StructWithEmbeddedView) AsStruct() *StructWithEmbedded {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v StructWithEmbeddedView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *StructWithEmbeddedView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x StructWithEmbedded
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v StructWithEmbeddedView) A() StructWithPtrsView { return v.ж.A.View() }
func (v StructWithEmbeddedView) StructWithSlices() StructWithSlicesView {
	return v.ж.StructWithSlices.View()
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _StructWithEmbeddedViewNeedsRegeneration = StructWithEmbedded(struct {
	A *StructWithPtrs
	StructWithSlices
}{})

// View returns a readonly view of GenericIntStruct.
func (p *GenericIntStruct[T]) View() GenericIntStructView[T] {
	return GenericIntStructView[T]{ж: p}
}

// GenericIntStructView[T] provides a read-only view over GenericIntStruct[T].
//
// Its methods should only be called if `Valid()` returns true.
type GenericIntStructView[T constraints.Integer] struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *GenericIntStruct[T]
}

// Valid reports whether underlying value is non-nil.
func (v GenericIntStructView[T]) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v GenericIntStructView[T]) AsStruct() *GenericIntStruct[T] {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v GenericIntStructView[T]) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *GenericIntStructView[T]) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x GenericIntStruct[T]
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v GenericIntStructView[T]) Value() T { return v.ж.Value }
func (v GenericIntStructView[T]) Pointer() views.ValuePointer[T] {
	return views.ValuePointerOf(v.ж.Pointer)
}

func (v GenericIntStructView[T]) Slice() views.Slice[T] { return views.SliceOf(v.ж.Slice) }

func (v GenericIntStructView[T]) Map() views.Map[string, T]  { return views.MapOf(v.ж.Map) }
func (v GenericIntStructView[T]) PtrSlice() *T               { panic("unsupported") }
func (v GenericIntStructView[T]) PtrKeyMap() map[*T]string   { panic("unsupported") }
func (v GenericIntStructView[T]) PtrValueMap() map[string]*T { panic("unsupported") }
func (v GenericIntStructView[T]) SliceMap() map[string][]T   { panic("unsupported") }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
func _GenericIntStructViewNeedsRegeneration[T constraints.Integer](GenericIntStruct[T]) {
	_GenericIntStructViewNeedsRegeneration(struct {
		Value       T
		Pointer     *T
		Slice       []T
		Map         map[string]T
		PtrSlice    []*T
		PtrKeyMap   map[*T]string `json:"-"`
		PtrValueMap map[string]*T
		SliceMap    map[string][]T
	}{})
}

// View returns a readonly view of GenericNoPtrsStruct.
func (p *GenericNoPtrsStruct[T]) View() GenericNoPtrsStructView[T] {
	return GenericNoPtrsStructView[T]{ж: p}
}

// GenericNoPtrsStructView[T] provides a read-only view over GenericNoPtrsStruct[T].
//
// Its methods should only be called if `Valid()` returns true.
type GenericNoPtrsStructView[T StructWithoutPtrs | netip.Prefix | BasicType] struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *GenericNoPtrsStruct[T]
}

// Valid reports whether underlying value is non-nil.
func (v GenericNoPtrsStructView[T]) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v GenericNoPtrsStructView[T]) AsStruct() *GenericNoPtrsStruct[T] {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v GenericNoPtrsStructView[T]) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *GenericNoPtrsStructView[T]) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x GenericNoPtrsStruct[T]
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v GenericNoPtrsStructView[T]) Value() T { return v.ж.Value }
func (v GenericNoPtrsStructView[T]) Pointer() views.ValuePointer[T] {
	return views.ValuePointerOf(v.ж.Pointer)
}

func (v GenericNoPtrsStructView[T]) Slice() views.Slice[T] { return views.SliceOf(v.ж.Slice) }

func (v GenericNoPtrsStructView[T]) Map() views.Map[string, T]  { return views.MapOf(v.ж.Map) }
func (v GenericNoPtrsStructView[T]) PtrSlice() *T               { panic("unsupported") }
func (v GenericNoPtrsStructView[T]) PtrKeyMap() map[*T]string   { panic("unsupported") }
func (v GenericNoPtrsStructView[T]) PtrValueMap() map[string]*T { panic("unsupported") }
func (v GenericNoPtrsStructView[T]) SliceMap() map[string][]T   { panic("unsupported") }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
func _GenericNoPtrsStructViewNeedsRegeneration[T StructWithoutPtrs | netip.Prefix | BasicType](GenericNoPtrsStruct[T]) {
	_GenericNoPtrsStructViewNeedsRegeneration(struct {
		Value       T
		Pointer     *T
		Slice       []T
		Map         map[string]T
		PtrSlice    []*T
		PtrKeyMap   map[*T]string `json:"-"`
		PtrValueMap map[string]*T
		SliceMap    map[string][]T
	}{})
}

// View returns a readonly view of GenericCloneableStruct.
func (p *GenericCloneableStruct[T, V]) View() GenericCloneableStructView[T, V] {
	return GenericCloneableStructView[T, V]{ж: p}
}

// GenericCloneableStructView[T, V] provides a read-only view over GenericCloneableStruct[T, V].
//
// Its methods should only be called if `Valid()` returns true.
type GenericCloneableStructView[T views.ViewCloner[T, V], V views.StructView[T]] struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *GenericCloneableStruct[T, V]
}

// Valid reports whether underlying value is non-nil.
func (v GenericCloneableStructView[T, V]) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v GenericCloneableStructView[T, V]) AsStruct() *GenericCloneableStruct[T, V] {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v GenericCloneableStructView[T, V]) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *GenericCloneableStructView[T, V]) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x GenericCloneableStruct[T, V]
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v GenericCloneableStructView[T, V]) Value() V { return v.ж.Value.View() }
func (v GenericCloneableStructView[T, V]) Slice() views.SliceView[T, V] {
	return views.SliceOfViews[T, V](v.ж.Slice)
}

func (v GenericCloneableStructView[T, V]) Map() views.MapFn[string, T, V] {
	return views.MapFnOf(v.ж.Map, func(t T) V {
		return t.View()
	})
}
func (v GenericCloneableStructView[T, V]) Pointer() map[string]T      { panic("unsupported") }
func (v GenericCloneableStructView[T, V]) PtrSlice() *T               { panic("unsupported") }
func (v GenericCloneableStructView[T, V]) PtrKeyMap() map[*T]string   { panic("unsupported") }
func (v GenericCloneableStructView[T, V]) PtrValueMap() map[string]*T { panic("unsupported") }
func (v GenericCloneableStructView[T, V]) SliceMap() map[string][]T   { panic("unsupported") }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
func _GenericCloneableStructViewNeedsRegeneration[T views.ViewCloner[T, V], V views.StructView[T]](GenericCloneableStruct[T, V]) {
	_GenericCloneableStructViewNeedsRegeneration(struct {
		Value       T
		Slice       []T
		Map         map[string]T
		Pointer     *T
		PtrSlice    []*T
		PtrKeyMap   map[*T]string `json:"-"`
		PtrValueMap map[string]*T
		SliceMap    map[string][]T
	}{})
}

// View returns a readonly view of StructWithContainers.
func (p *StructWithContainers) View() StructWithContainersView {
	return StructWithContainersView{ж: p}
}

// StructWithContainersView provides a read-only view over StructWithContainers.
//
// Its methods should only be called if `Valid()` returns true.
type StructWithContainersView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *StructWithContainers
}

// Valid reports whether underlying value is non-nil.
func (v StructWithContainersView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v StructWithContainersView) AsStruct() *StructWithContainers {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v StructWithContainersView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *StructWithContainersView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x StructWithContainers
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v StructWithContainersView) IntContainer() Container[int] { return v.ж.IntContainer }
func (v StructWithContainersView) CloneableContainer() ContainerView[*StructWithPtrs, StructWithPtrsView] {
	return ContainerViewOf(&v.ж.CloneableContainer)
}
func (v StructWithContainersView) BasicGenericContainer() Container[GenericBasicStruct[int]] {
	return v.ж.BasicGenericContainer
}
func (v StructWithContainersView) CloneableGenericContainer() ContainerView[*GenericNoPtrsStruct[int], GenericNoPtrsStructView[int]] {
	return ContainerViewOf(&v.ж.CloneableGenericContainer)
}
func (v StructWithContainersView) CloneableMap() MapContainerView[int, *StructWithPtrs, StructWithPtrsView] {
	return MapContainerViewOf(&v.ж.CloneableMap)
}
func (v StructWithContainersView) CloneableGenericMap() MapContainerView[int, *GenericNoPtrsStruct[int], GenericNoPtrsStructView[int]] {
	return MapContainerViewOf(&v.ж.CloneableGenericMap)
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _StructWithContainersViewNeedsRegeneration = StructWithContainers(struct {
	IntContainer              Container[int]
	CloneableContainer        Container[*StructWithPtrs]
	BasicGenericContainer     Container[GenericBasicStruct[int]]
	CloneableGenericContainer Container[*GenericNoPtrsStruct[int]]
	CloneableMap              MapContainer[int, *StructWithPtrs]
	CloneableGenericMap       MapContainer[int, *GenericNoPtrsStruct[int]]
}{})

// View returns a readonly view of StructWithTypeAliasFields.
func (p *StructWithTypeAliasFields) View() StructWithTypeAliasFieldsView {
	return StructWithTypeAliasFieldsView{ж: p}
}

// StructWithTypeAliasFieldsView provides a read-only view over StructWithTypeAliasFields.
//
// Its methods should only be called if `Valid()` returns true.
type StructWithTypeAliasFieldsView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *StructWithTypeAliasFields
}

// Valid reports whether underlying value is non-nil.
func (v StructWithTypeAliasFieldsView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v StructWithTypeAliasFieldsView) AsStruct() *StructWithTypeAliasFields {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v StructWithTypeAliasFieldsView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *StructWithTypeAliasFieldsView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x StructWithTypeAliasFields
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v StructWithTypeAliasFieldsView) WithPtr() StructWithPtrsAliasView   { return v.ж.WithPtr.View() }
func (v StructWithTypeAliasFieldsView) WithoutPtr() StructWithoutPtrsAlias { return v.ж.WithoutPtr }
func (v StructWithTypeAliasFieldsView) WithPtrByPtr() StructWithPtrsAliasView {
	return v.ж.WithPtrByPtr.View()
}
func (v StructWithTypeAliasFieldsView) WithoutPtrByPtr() StructWithoutPtrsAliasView {
	return v.ж.WithoutPtrByPtr.View()
}
func (v StructWithTypeAliasFieldsView) SliceWithPtrs() views.SliceView[*StructWithPtrsAlias, StructWithPtrsAliasView] {
	return views.SliceOfViews[*StructWithPtrsAlias, StructWithPtrsAliasView](v.ж.SliceWithPtrs)
}
func (v StructWithTypeAliasFieldsView) SliceWithoutPtrs() views.SliceView[*StructWithoutPtrsAlias, StructWithoutPtrsAliasView] {
	return views.SliceOfViews[*StructWithoutPtrsAlias, StructWithoutPtrsAliasView](v.ж.SliceWithoutPtrs)
}

func (v StructWithTypeAliasFieldsView) MapWithPtrs() views.MapFn[string, *StructWithPtrsAlias, StructWithPtrsAliasView] {
	return views.MapFnOf(v.ж.MapWithPtrs, func(t *StructWithPtrsAlias) StructWithPtrsAliasView {
		return t.View()
	})
}

func (v StructWithTypeAliasFieldsView) MapWithoutPtrs() views.MapFn[string, *StructWithoutPtrsAlias, StructWithoutPtrsAliasView] {
	return views.MapFnOf(v.ж.MapWithoutPtrs, func(t *StructWithoutPtrsAlias) StructWithoutPtrsAliasView {
		return t.View()
	})
}

func (v StructWithTypeAliasFieldsView) MapOfSlicesWithPtrs() views.MapFn[string, []*StructWithPtrsAlias, views.SliceView[*StructWithPtrsAlias, StructWithPtrsAliasView]] {
	return views.MapFnOf(v.ж.MapOfSlicesWithPtrs, func(t []*StructWithPtrsAlias) views.SliceView[*StructWithPtrsAlias, StructWithPtrsAliasView] {
		return views.SliceOfViews[*StructWithPtrsAlias, StructWithPtrsAliasView](t)
	})
}

func (v StructWithTypeAliasFieldsView) MapOfSlicesWithoutPtrs() views.MapFn[string, []*StructWithoutPtrsAlias, views.SliceView[*StructWithoutPtrsAlias, StructWithoutPtrsAliasView]] {
	return views.MapFnOf(v.ж.MapOfSlicesWithoutPtrs, func(t []*StructWithoutPtrsAlias) views.SliceView[*StructWithoutPtrsAlias, StructWithoutPtrsAliasView] {
		return views.SliceOfViews[*StructWithoutPtrsAlias, StructWithoutPtrsAliasView](t)
	})
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _StructWithTypeAliasFieldsViewNeedsRegeneration = StructWithTypeAliasFields(struct {
	WithPtr                StructWithPtrsAlias
	WithoutPtr             StructWithoutPtrsAlias
	WithPtrByPtr           *StructWithPtrsAlias
	WithoutPtrByPtr        *StructWithoutPtrsAlias
	SliceWithPtrs          []*StructWithPtrsAlias
	SliceWithoutPtrs       []*StructWithoutPtrsAlias
	MapWithPtrs            map[string]*StructWithPtrsAlias
	MapWithoutPtrs         map[string]*StructWithoutPtrsAlias
	MapOfSlicesWithPtrs    map[string][]*StructWithPtrsAlias
	MapOfSlicesWithoutPtrs map[string][]*StructWithoutPtrsAlias
}{})

// View returns a readonly view of GenericTypeAliasStruct.
func (p *GenericTypeAliasStruct[T, T2, V2]) View() GenericTypeAliasStructView[T, T2, V2] {
	return GenericTypeAliasStructView[T, T2, V2]{ж: p}
}

// GenericTypeAliasStructView[T, T2, V2] provides a read-only view over GenericTypeAliasStruct[T, T2, V2].
//
// Its methods should only be called if `Valid()` returns true.
type GenericTypeAliasStructView[T integer, T2 views.ViewCloner[T2, V2], V2 views.StructView[T2]] struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *GenericTypeAliasStruct[T, T2, V2]
}

// Valid reports whether underlying value is non-nil.
func (v GenericTypeAliasStructView[T, T2, V2]) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v GenericTypeAliasStructView[T, T2, V2]) AsStruct() *GenericTypeAliasStruct[T, T2, V2] {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v GenericTypeAliasStructView[T, T2, V2]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.ж)
}

func (v *GenericTypeAliasStructView[T, T2, V2]) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x GenericTypeAliasStruct[T, T2, V2]
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v GenericTypeAliasStructView[T, T2, V2]) NonCloneable() T { return v.ж.NonCloneable }
func (v GenericTypeAliasStructView[T, T2, V2]) Cloneable() V2   { return v.ж.Cloneable.View() }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
func _GenericTypeAliasStructViewNeedsRegeneration[T integer, T2 views.ViewCloner[T2, V2], V2 views.StructView[T2]](GenericTypeAliasStruct[T, T2, V2]) {
	_GenericTypeAliasStructViewNeedsRegeneration(struct {
		NonCloneable T
		Cloneable    T2
	}{})
}
