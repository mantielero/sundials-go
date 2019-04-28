// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sun, 28 Apr 2019 00:38:31 CEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package nvector

/*
#cgo LDFLAGS: -lsundials_nvecserial -L/usr/lib
#include <stdio.h>
#include <sundials/sundials_types.h>
#include <sundials/sundials_nvector.h>
#include <sundials/sundials_fnvector.h>
#include <nvector/nvector_serial.h>
#include <sundials/sundials_math.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocNVectorContentSerialMemory allocates memory for type C.N_VectorContent_Serial in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNVectorContentSerialMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNVectorContentSerialValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNVectorContentSerialValue = unsafe.Sizeof([1]C.N_VectorContent_Serial{})

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *NVectorContentSerial) Ref() *C.N_VectorContent_Serial {
	if x == nil {
		return nil
	}
	return x.ref469bbb76
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *NVectorContentSerial) Free() {
	if x != nil && x.allocs469bbb76 != nil {
		x.allocs469bbb76.(*cgoAllocMap).Free()
		x.ref469bbb76 = nil
	}
}

// NewNVectorContentSerialRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewNVectorContentSerialRef(ref unsafe.Pointer) *NVectorContentSerial {
	if ref == nil {
		return nil
	}
	obj := new(NVectorContentSerial)
	obj.ref469bbb76 = (*C.N_VectorContent_Serial)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *NVectorContentSerial) PassRef() (*C.N_VectorContent_Serial, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref469bbb76 != nil {
		return x.ref469bbb76, nil
	}
	mem469bbb76 := allocNVectorContentSerialMemory(1)
	ref469bbb76 := (*C.N_VectorContent_Serial)(mem469bbb76)
	allocs469bbb76 := new(cgoAllocMap)
	allocs469bbb76.Add(mem469bbb76)

	var clength_allocs *cgoAllocMap
	ref469bbb76.length, clength_allocs = (C.sunindextype)(x.Length), cgoAllocsUnknown
	allocs469bbb76.Borrow(clength_allocs)

	var cown_data_allocs *cgoAllocMap
	ref469bbb76.own_data, cown_data_allocs = (C.int)(x.OwnData), cgoAllocsUnknown
	allocs469bbb76.Borrow(cown_data_allocs)

	var cdata_allocs *cgoAllocMap
	ref469bbb76.data, cdata_allocs = (*C.realtype)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Data)).Data)), cgoAllocsUnknown
	allocs469bbb76.Borrow(cdata_allocs)

	x.ref469bbb76 = ref469bbb76
	x.allocs469bbb76 = allocs469bbb76
	return ref469bbb76, allocs469bbb76

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x NVectorContentSerial) PassValue() (C.N_VectorContent_Serial, *cgoAllocMap) {
	if x.ref469bbb76 != nil {
		return *x.ref469bbb76, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *NVectorContentSerial) Deref() {
	if x.ref469bbb76 == nil {
		return
	}
	x.Length = (Sunindextype)(x.ref469bbb76.length)
	x.OwnData = (int32)(x.ref469bbb76.own_data)
	hxfc4425b := (*sliceHeader)(unsafe.Pointer(&x.Data))
	hxfc4425b.Data = unsafe.Pointer(x.ref469bbb76.data)
	hxfc4425b.Cap = 0x7fffffff
	// hxfc4425b.Len = ?

}

// allocFILEMemory allocates memory for type C.FILE in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFILEMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFILEValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfFILEValue = unsafe.Sizeof([1]C.FILE{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *FILE) Ref() *C.FILE {
	if x == nil {
		return nil
	}
	return x.refba0adba4
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *FILE) Free() {
	if x != nil && x.allocsba0adba4 != nil {
		x.allocsba0adba4.(*cgoAllocMap).Free()
		x.refba0adba4 = nil
	}
}

// NewFILERef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewFILERef(ref unsafe.Pointer) *FILE {
	if ref == nil {
		return nil
	}
	obj := new(FILE)
	obj.refba0adba4 = (*C.FILE)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *FILE) PassRef() (*C.FILE, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refba0adba4 != nil {
		return x.refba0adba4, nil
	}
	memba0adba4 := allocFILEMemory(1)
	refba0adba4 := (*C.FILE)(memba0adba4)
	allocsba0adba4 := new(cgoAllocMap)
	allocsba0adba4.Add(memba0adba4)

	var c_flags_allocs *cgoAllocMap
	refba0adba4._flags, c_flags_allocs = (C.int)(x.Flags), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_flags_allocs)

	var c_IO_read_ptr_allocs *cgoAllocMap
	refba0adba4._IO_read_ptr, c_IO_read_ptr_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOReadPtr)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_read_ptr_allocs)

	var c_IO_read_end_allocs *cgoAllocMap
	refba0adba4._IO_read_end, c_IO_read_end_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOReadEnd)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_read_end_allocs)

	var c_IO_read_base_allocs *cgoAllocMap
	refba0adba4._IO_read_base, c_IO_read_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOReadBase)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_read_base_allocs)

	var c_IO_write_base_allocs *cgoAllocMap
	refba0adba4._IO_write_base, c_IO_write_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOWriteBase)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_write_base_allocs)

	var c_IO_write_ptr_allocs *cgoAllocMap
	refba0adba4._IO_write_ptr, c_IO_write_ptr_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOWritePtr)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_write_ptr_allocs)

	var c_IO_write_end_allocs *cgoAllocMap
	refba0adba4._IO_write_end, c_IO_write_end_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOWriteEnd)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_write_end_allocs)

	var c_IO_buf_base_allocs *cgoAllocMap
	refba0adba4._IO_buf_base, c_IO_buf_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOBufBase)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_buf_base_allocs)

	var c_IO_buf_end_allocs *cgoAllocMap
	refba0adba4._IO_buf_end, c_IO_buf_end_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOBufEnd)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_buf_end_allocs)

	var c_IO_save_base_allocs *cgoAllocMap
	refba0adba4._IO_save_base, c_IO_save_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOSaveBase)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_save_base_allocs)

	var c_IO_backup_base_allocs *cgoAllocMap
	refba0adba4._IO_backup_base, c_IO_backup_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOBackupBase)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_backup_base_allocs)

	var c_IO_save_end_allocs *cgoAllocMap
	refba0adba4._IO_save_end, c_IO_save_end_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.IOSaveEnd)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_save_end_allocs)

	var c_fileno_allocs *cgoAllocMap
	refba0adba4._fileno, c_fileno_allocs = (C.int)(x.Fileno), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_fileno_allocs)

	var c_flags2_allocs *cgoAllocMap
	refba0adba4._flags2, c_flags2_allocs = (C.int)(x.Flags2), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_flags2_allocs)

	var c_old_offset_allocs *cgoAllocMap
	refba0adba4._old_offset, c_old_offset_allocs = (C.__off_t)(x.OldOffset), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_old_offset_allocs)

	var c_cur_column_allocs *cgoAllocMap
	refba0adba4._cur_column, c_cur_column_allocs = (C.ushort)(x.CurColumn), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_cur_column_allocs)

	var c_vtable_offset_allocs *cgoAllocMap
	refba0adba4._vtable_offset, c_vtable_offset_allocs = (C.char)(x.VtableOffset), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_vtable_offset_allocs)

	var c_shortbuf_allocs *cgoAllocMap
	refba0adba4._shortbuf, c_shortbuf_allocs = *(*[1]C.char)(unsafe.Pointer(&x.Shortbuf)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_shortbuf_allocs)

	var c_lock_allocs *cgoAllocMap
	refba0adba4._lock, c_lock_allocs = *(**C._IO_lock_t)(unsafe.Pointer(&x.Lock)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_lock_allocs)

	var c_offset_allocs *cgoAllocMap
	refba0adba4._offset, c_offset_allocs = (C.__off64_t)(x.Offset), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_offset_allocs)

	var c_freeres_buf_allocs *cgoAllocMap
	refba0adba4._freeres_buf, c_freeres_buf_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.FreeresBuf)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_freeres_buf_allocs)

	var c__pad5_allocs *cgoAllocMap
	refba0adba4.__pad5, c__pad5_allocs = (C.size_t)(x._Pad5), cgoAllocsUnknown
	allocsba0adba4.Borrow(c__pad5_allocs)

	var c_mode_allocs *cgoAllocMap
	refba0adba4._mode, c_mode_allocs = (C.int)(x.Mode), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_mode_allocs)

	var c_unused2_allocs *cgoAllocMap
	refba0adba4._unused2, c_unused2_allocs = *(*[20]C.char)(unsafe.Pointer(&x.Unused2)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_unused2_allocs)

	x.refba0adba4 = refba0adba4
	x.allocsba0adba4 = allocsba0adba4
	return refba0adba4, allocsba0adba4

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x FILE) PassValue() (C.FILE, *cgoAllocMap) {
	if x.refba0adba4 != nil {
		return *x.refba0adba4, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *FILE) Deref() {
	if x.refba0adba4 == nil {
		return
	}
	x.Flags = (int32)(x.refba0adba4._flags)
	hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&x.IOReadPtr))
	hxf95e7c8.Data = unsafe.Pointer(x.refba0adba4._IO_read_ptr)
	hxf95e7c8.Cap = 0x7fffffff
	// hxf95e7c8.Len = ?

	hxff2234b := (*sliceHeader)(unsafe.Pointer(&x.IOReadEnd))
	hxff2234b.Data = unsafe.Pointer(x.refba0adba4._IO_read_end)
	hxff2234b.Cap = 0x7fffffff
	// hxff2234b.Len = ?

	hxff73280 := (*sliceHeader)(unsafe.Pointer(&x.IOReadBase))
	hxff73280.Data = unsafe.Pointer(x.refba0adba4._IO_read_base)
	hxff73280.Cap = 0x7fffffff
	// hxff73280.Len = ?

	hxfa9955c := (*sliceHeader)(unsafe.Pointer(&x.IOWriteBase))
	hxfa9955c.Data = unsafe.Pointer(x.refba0adba4._IO_write_base)
	hxfa9955c.Cap = 0x7fffffff
	// hxfa9955c.Len = ?

	hxfa3f05c := (*sliceHeader)(unsafe.Pointer(&x.IOWritePtr))
	hxfa3f05c.Data = unsafe.Pointer(x.refba0adba4._IO_write_ptr)
	hxfa3f05c.Cap = 0x7fffffff
	// hxfa3f05c.Len = ?

	hxf0d18b7 := (*sliceHeader)(unsafe.Pointer(&x.IOWriteEnd))
	hxf0d18b7.Data = unsafe.Pointer(x.refba0adba4._IO_write_end)
	hxf0d18b7.Cap = 0x7fffffff
	// hxf0d18b7.Len = ?

	hxf2fab0d := (*sliceHeader)(unsafe.Pointer(&x.IOBufBase))
	hxf2fab0d.Data = unsafe.Pointer(x.refba0adba4._IO_buf_base)
	hxf2fab0d.Cap = 0x7fffffff
	// hxf2fab0d.Len = ?

	hxf69fe70 := (*sliceHeader)(unsafe.Pointer(&x.IOBufEnd))
	hxf69fe70.Data = unsafe.Pointer(x.refba0adba4._IO_buf_end)
	hxf69fe70.Cap = 0x7fffffff
	// hxf69fe70.Len = ?

	hxf65bf54 := (*sliceHeader)(unsafe.Pointer(&x.IOSaveBase))
	hxf65bf54.Data = unsafe.Pointer(x.refba0adba4._IO_save_base)
	hxf65bf54.Cap = 0x7fffffff
	// hxf65bf54.Len = ?

	hxf3b8dbd := (*sliceHeader)(unsafe.Pointer(&x.IOBackupBase))
	hxf3b8dbd.Data = unsafe.Pointer(x.refba0adba4._IO_backup_base)
	hxf3b8dbd.Cap = 0x7fffffff
	// hxf3b8dbd.Len = ?

	hxf7a6dff := (*sliceHeader)(unsafe.Pointer(&x.IOSaveEnd))
	hxf7a6dff.Data = unsafe.Pointer(x.refba0adba4._IO_save_end)
	hxf7a6dff.Cap = 0x7fffffff
	// hxf7a6dff.Len = ?

	x.Fileno = (int32)(x.refba0adba4._fileno)
	x.Flags2 = (int32)(x.refba0adba4._flags2)
	x.OldOffset = (int)(x.refba0adba4._old_offset)
	x.CurColumn = (uint16)(x.refba0adba4._cur_column)
	x.VtableOffset = (byte)(x.refba0adba4._vtable_offset)
	x.Shortbuf = *(*[1]byte)(unsafe.Pointer(&x.refba0adba4._shortbuf))
	x.Lock = (unsafe.Pointer)(unsafe.Pointer(x.refba0adba4._lock))
	x.Offset = (int)(x.refba0adba4._offset)
	x.FreeresBuf = (unsafe.Pointer)(unsafe.Pointer(x.refba0adba4._freeres_buf))
	x._Pad5 = (uint)(x.refba0adba4.__pad5)
	x.Mode = (int32)(x.refba0adba4._mode)
	x.Unused2 = *(*[20]byte)(unsafe.Pointer(&x.refba0adba4._unused2))
}

// allocNVectorOpsMemory allocates memory for type C.N_Vector_Ops in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNVectorOpsMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNVectorOpsValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNVectorOpsValue = unsafe.Sizeof([1]C.N_Vector_Ops{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *NVectorOps) Ref() *C.N_Vector_Ops {
	if x == nil {
		return nil
	}
	return x.ref25b9efb8
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *NVectorOps) Free() {
	if x != nil && x.allocs25b9efb8 != nil {
		x.allocs25b9efb8.(*cgoAllocMap).Free()
		x.ref25b9efb8 = nil
	}
}

// NewNVectorOpsRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewNVectorOpsRef(ref unsafe.Pointer) *NVectorOps {
	if ref == nil {
		return nil
	}
	obj := new(NVectorOps)
	obj.ref25b9efb8 = (*C.N_Vector_Ops)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *NVectorOps) PassRef() (*C.N_Vector_Ops, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref25b9efb8 != nil {
		return x.ref25b9efb8, nil
	}
	mem25b9efb8 := allocNVectorOpsMemory(1)
	ref25b9efb8 := (*C.N_Vector_Ops)(mem25b9efb8)
	allocs25b9efb8 := new(cgoAllocMap)
	allocs25b9efb8.Add(mem25b9efb8)

	var cnvgetvectorid_allocs *cgoAllocMap
	ref25b9efb8.nvgetvectorid, cnvgetvectorid_allocs = x.Nvgetvectorid.PassRef()
	allocs25b9efb8.Borrow(cnvgetvectorid_allocs)

	var cnvclone_allocs *cgoAllocMap
	ref25b9efb8.nvclone, cnvclone_allocs = x.Nvclone.PassRef()
	allocs25b9efb8.Borrow(cnvclone_allocs)

	var cnvcloneempty_allocs *cgoAllocMap
	ref25b9efb8.nvcloneempty, cnvcloneempty_allocs = x.Nvcloneempty.PassRef()
	allocs25b9efb8.Borrow(cnvcloneempty_allocs)

	var cnvdestroy_allocs *cgoAllocMap
	ref25b9efb8.nvdestroy, cnvdestroy_allocs = x.Nvdestroy.PassRef()
	allocs25b9efb8.Borrow(cnvdestroy_allocs)

	var cnvspace_allocs *cgoAllocMap
	ref25b9efb8.nvspace, cnvspace_allocs = x.Nvspace.PassRef()
	allocs25b9efb8.Borrow(cnvspace_allocs)

	var cnvgetarraypointer_allocs *cgoAllocMap
	ref25b9efb8.nvgetarraypointer, cnvgetarraypointer_allocs = x.Nvgetarraypointer.PassRef()
	allocs25b9efb8.Borrow(cnvgetarraypointer_allocs)

	var cnvsetarraypointer_allocs *cgoAllocMap
	ref25b9efb8.nvsetarraypointer, cnvsetarraypointer_allocs = x.Nvsetarraypointer.PassRef()
	allocs25b9efb8.Borrow(cnvsetarraypointer_allocs)

	var cnvlinearsum_allocs *cgoAllocMap
	ref25b9efb8.nvlinearsum, cnvlinearsum_allocs = x.Nvlinearsum.PassRef()
	allocs25b9efb8.Borrow(cnvlinearsum_allocs)

	var cnvconst_allocs *cgoAllocMap
	ref25b9efb8.nvconst, cnvconst_allocs = x.Nvconst.PassRef()
	allocs25b9efb8.Borrow(cnvconst_allocs)

	var cnvprod_allocs *cgoAllocMap
	ref25b9efb8.nvprod, cnvprod_allocs = x.Nvprod.PassRef()
	allocs25b9efb8.Borrow(cnvprod_allocs)

	var cnvdiv_allocs *cgoAllocMap
	ref25b9efb8.nvdiv, cnvdiv_allocs = x.Nvdiv.PassRef()
	allocs25b9efb8.Borrow(cnvdiv_allocs)

	var cnvscale_allocs *cgoAllocMap
	ref25b9efb8.nvscale, cnvscale_allocs = x.Nvscale.PassRef()
	allocs25b9efb8.Borrow(cnvscale_allocs)

	var cnvabs_allocs *cgoAllocMap
	ref25b9efb8.nvabs, cnvabs_allocs = x.Nvabs.PassRef()
	allocs25b9efb8.Borrow(cnvabs_allocs)

	var cnvinv_allocs *cgoAllocMap
	ref25b9efb8.nvinv, cnvinv_allocs = x.Nvinv.PassRef()
	allocs25b9efb8.Borrow(cnvinv_allocs)

	var cnvaddconst_allocs *cgoAllocMap
	ref25b9efb8.nvaddconst, cnvaddconst_allocs = x.Nvaddconst.PassRef()
	allocs25b9efb8.Borrow(cnvaddconst_allocs)

	var cnvdotprod_allocs *cgoAllocMap
	ref25b9efb8.nvdotprod, cnvdotprod_allocs = x.Nvdotprod.PassRef()
	allocs25b9efb8.Borrow(cnvdotprod_allocs)

	var cnvmaxnorm_allocs *cgoAllocMap
	ref25b9efb8.nvmaxnorm, cnvmaxnorm_allocs = x.Nvmaxnorm.PassRef()
	allocs25b9efb8.Borrow(cnvmaxnorm_allocs)

	var cnvwrmsnorm_allocs *cgoAllocMap
	ref25b9efb8.nvwrmsnorm, cnvwrmsnorm_allocs = x.Nvwrmsnorm.PassRef()
	allocs25b9efb8.Borrow(cnvwrmsnorm_allocs)

	var cnvwrmsnormmask_allocs *cgoAllocMap
	ref25b9efb8.nvwrmsnormmask, cnvwrmsnormmask_allocs = x.Nvwrmsnormmask.PassRef()
	allocs25b9efb8.Borrow(cnvwrmsnormmask_allocs)

	var cnvmin_allocs *cgoAllocMap
	ref25b9efb8.nvmin, cnvmin_allocs = x.Nvmin.PassRef()
	allocs25b9efb8.Borrow(cnvmin_allocs)

	var cnvwl2norm_allocs *cgoAllocMap
	ref25b9efb8.nvwl2norm, cnvwl2norm_allocs = x.Nvwl2norm.PassRef()
	allocs25b9efb8.Borrow(cnvwl2norm_allocs)

	var cnvl1norm_allocs *cgoAllocMap
	ref25b9efb8.nvl1norm, cnvl1norm_allocs = x.Nvl1norm.PassRef()
	allocs25b9efb8.Borrow(cnvl1norm_allocs)

	var cnvcompare_allocs *cgoAllocMap
	ref25b9efb8.nvcompare, cnvcompare_allocs = x.Nvcompare.PassRef()
	allocs25b9efb8.Borrow(cnvcompare_allocs)

	var cnvinvtest_allocs *cgoAllocMap
	ref25b9efb8.nvinvtest, cnvinvtest_allocs = x.Nvinvtest.PassRef()
	allocs25b9efb8.Borrow(cnvinvtest_allocs)

	var cnvconstrmask_allocs *cgoAllocMap
	ref25b9efb8.nvconstrmask, cnvconstrmask_allocs = x.Nvconstrmask.PassRef()
	allocs25b9efb8.Borrow(cnvconstrmask_allocs)

	var cnvminquotient_allocs *cgoAllocMap
	ref25b9efb8.nvminquotient, cnvminquotient_allocs = x.Nvminquotient.PassRef()
	allocs25b9efb8.Borrow(cnvminquotient_allocs)

	var cnvlinearcombination_allocs *cgoAllocMap
	ref25b9efb8.nvlinearcombination, cnvlinearcombination_allocs = x.Nvlinearcombination.PassRef()
	allocs25b9efb8.Borrow(cnvlinearcombination_allocs)

	var cnvscaleaddmulti_allocs *cgoAllocMap
	ref25b9efb8.nvscaleaddmulti, cnvscaleaddmulti_allocs = x.Nvscaleaddmulti.PassRef()
	allocs25b9efb8.Borrow(cnvscaleaddmulti_allocs)

	var cnvdotprodmulti_allocs *cgoAllocMap
	ref25b9efb8.nvdotprodmulti, cnvdotprodmulti_allocs = x.Nvdotprodmulti.PassRef()
	allocs25b9efb8.Borrow(cnvdotprodmulti_allocs)

	var cnvlinearsumvectorarray_allocs *cgoAllocMap
	ref25b9efb8.nvlinearsumvectorarray, cnvlinearsumvectorarray_allocs = x.Nvlinearsumvectorarray.PassRef()
	allocs25b9efb8.Borrow(cnvlinearsumvectorarray_allocs)

	var cnvscalevectorarray_allocs *cgoAllocMap
	ref25b9efb8.nvscalevectorarray, cnvscalevectorarray_allocs = x.Nvscalevectorarray.PassRef()
	allocs25b9efb8.Borrow(cnvscalevectorarray_allocs)

	var cnvconstvectorarray_allocs *cgoAllocMap
	ref25b9efb8.nvconstvectorarray, cnvconstvectorarray_allocs = x.Nvconstvectorarray.PassRef()
	allocs25b9efb8.Borrow(cnvconstvectorarray_allocs)

	var cnvwrmsnormvectorarray_allocs *cgoAllocMap
	ref25b9efb8.nvwrmsnormvectorarray, cnvwrmsnormvectorarray_allocs = x.Nvwrmsnormvectorarray.PassRef()
	allocs25b9efb8.Borrow(cnvwrmsnormvectorarray_allocs)

	var cnvwrmsnormmaskvectorarray_allocs *cgoAllocMap
	ref25b9efb8.nvwrmsnormmaskvectorarray, cnvwrmsnormmaskvectorarray_allocs = x.Nvwrmsnormmaskvectorarray.PassRef()
	allocs25b9efb8.Borrow(cnvwrmsnormmaskvectorarray_allocs)

	var cnvscaleaddmultivectorarray_allocs *cgoAllocMap
	ref25b9efb8.nvscaleaddmultivectorarray, cnvscaleaddmultivectorarray_allocs = x.Nvscaleaddmultivectorarray.PassRef()
	allocs25b9efb8.Borrow(cnvscaleaddmultivectorarray_allocs)

	var cnvlinearcombinationvectorarray_allocs *cgoAllocMap
	ref25b9efb8.nvlinearcombinationvectorarray, cnvlinearcombinationvectorarray_allocs = x.Nvlinearcombinationvectorarray.PassRef()
	allocs25b9efb8.Borrow(cnvlinearcombinationvectorarray_allocs)

	x.ref25b9efb8 = ref25b9efb8
	x.allocs25b9efb8 = allocs25b9efb8
	return ref25b9efb8, allocs25b9efb8

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x NVectorOps) PassValue() (C.N_Vector_Ops, *cgoAllocMap) {
	if x.ref25b9efb8 != nil {
		return *x.ref25b9efb8, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *NVectorOps) Deref() {
	if x.ref25b9efb8 == nil {
		return
	}
	x.Nvgetvectorid = NewNVectorIDRef(unsafe.Pointer(x.ref25b9efb8.nvgetvectorid))
	x.Nvclone = NewNVectorRef(unsafe.Pointer(x.ref25b9efb8.nvclone))
	x.Nvcloneempty = NewNVectorRef(unsafe.Pointer(x.ref25b9efb8.nvcloneempty))
	x.Nvdestroy = NewRef(unsafe.Pointer(x.ref25b9efb8.nvdestroy))
	x.Nvspace = NewRef(unsafe.Pointer(x.ref25b9efb8.nvspace))
	x.Nvgetarraypointer = NewRealtypeRef(unsafe.Pointer(x.ref25b9efb8.nvgetarraypointer))
	x.Nvsetarraypointer = NewRef(unsafe.Pointer(x.ref25b9efb8.nvsetarraypointer))
	x.Nvlinearsum = NewRef(unsafe.Pointer(x.ref25b9efb8.nvlinearsum))
	x.Nvconst = NewRef(unsafe.Pointer(x.ref25b9efb8.nvconst))
	x.Nvprod = NewRef(unsafe.Pointer(x.ref25b9efb8.nvprod))
	x.Nvdiv = NewRef(unsafe.Pointer(x.ref25b9efb8.nvdiv))
	x.Nvscale = NewRef(unsafe.Pointer(x.ref25b9efb8.nvscale))
	x.Nvabs = NewRef(unsafe.Pointer(x.ref25b9efb8.nvabs))
	x.Nvinv = NewRef(unsafe.Pointer(x.ref25b9efb8.nvinv))
	x.Nvaddconst = NewRef(unsafe.Pointer(x.ref25b9efb8.nvaddconst))
	x.Nvdotprod = NewRealtypeRef(unsafe.Pointer(x.ref25b9efb8.nvdotprod))
	x.Nvmaxnorm = NewRealtypeRef(unsafe.Pointer(x.ref25b9efb8.nvmaxnorm))
	x.Nvwrmsnorm = NewRealtypeRef(unsafe.Pointer(x.ref25b9efb8.nvwrmsnorm))
	x.Nvwrmsnormmask = NewRealtypeRef(unsafe.Pointer(x.ref25b9efb8.nvwrmsnormmask))
	x.Nvmin = NewRealtypeRef(unsafe.Pointer(x.ref25b9efb8.nvmin))
	x.Nvwl2norm = NewRealtypeRef(unsafe.Pointer(x.ref25b9efb8.nvwl2norm))
	x.Nvl1norm = NewRealtypeRef(unsafe.Pointer(x.ref25b9efb8.nvl1norm))
	x.Nvcompare = NewRef(unsafe.Pointer(x.ref25b9efb8.nvcompare))
	x.Nvinvtest = NewRef(unsafe.Pointer(x.ref25b9efb8.nvinvtest))
	x.Nvconstrmask = NewRef(unsafe.Pointer(x.ref25b9efb8.nvconstrmask))
	x.Nvminquotient = NewRealtypeRef(unsafe.Pointer(x.ref25b9efb8.nvminquotient))
	x.Nvlinearcombination = NewRef(unsafe.Pointer(x.ref25b9efb8.nvlinearcombination))
	x.Nvscaleaddmulti = NewRef(unsafe.Pointer(x.ref25b9efb8.nvscaleaddmulti))
	x.Nvdotprodmulti = NewRef(unsafe.Pointer(x.ref25b9efb8.nvdotprodmulti))
	x.Nvlinearsumvectorarray = NewRef(unsafe.Pointer(x.ref25b9efb8.nvlinearsumvectorarray))
	x.Nvscalevectorarray = NewRef(unsafe.Pointer(x.ref25b9efb8.nvscalevectorarray))
	x.Nvconstvectorarray = NewRef(unsafe.Pointer(x.ref25b9efb8.nvconstvectorarray))
	x.Nvwrmsnormvectorarray = NewRef(unsafe.Pointer(x.ref25b9efb8.nvwrmsnormvectorarray))
	x.Nvwrmsnormmaskvectorarray = NewRef(unsafe.Pointer(x.ref25b9efb8.nvwrmsnormmaskvectorarray))
	x.Nvscaleaddmultivectorarray = NewRef(unsafe.Pointer(x.ref25b9efb8.nvscaleaddmultivectorarray))
	x.Nvlinearcombinationvectorarray = NewRef(unsafe.Pointer(x.ref25b9efb8.nvlinearcombinationvectorarray))
}

// allocNVectorMemory allocates memory for type C.N_Vector in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNVectorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNVectorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNVectorValue = unsafe.Sizeof([1]C.N_Vector{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *NVector) Ref() *C.N_Vector {
	if x == nil {
		return nil
	}
	return x.refdb950bdc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *NVector) Free() {
	if x != nil && x.allocsdb950bdc != nil {
		x.allocsdb950bdc.(*cgoAllocMap).Free()
		x.refdb950bdc = nil
	}
}

// NewNVectorRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewNVectorRef(ref unsafe.Pointer) *NVector {
	if ref == nil {
		return nil
	}
	obj := new(NVector)
	obj.refdb950bdc = (*C.N_Vector)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *NVector) PassRef() (*C.N_Vector, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refdb950bdc != nil {
		return x.refdb950bdc, nil
	}
	memdb950bdc := allocNVectorMemory(1)
	refdb950bdc := (*C.N_Vector)(memdb950bdc)
	allocsdb950bdc := new(cgoAllocMap)
	allocsdb950bdc.Add(memdb950bdc)

	var ccontent_allocs *cgoAllocMap
	refdb950bdc.content, ccontent_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.Content)), cgoAllocsUnknown
	allocsdb950bdc.Borrow(ccontent_allocs)

	var cops_allocs *cgoAllocMap
	refdb950bdc.ops, cops_allocs = x.Ops.PassValue()
	allocsdb950bdc.Borrow(cops_allocs)

	x.refdb950bdc = refdb950bdc
	x.allocsdb950bdc = allocsdb950bdc
	return refdb950bdc, allocsdb950bdc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x NVector) PassValue() (C.N_Vector, *cgoAllocMap) {
	if x.refdb950bdc != nil {
		return *x.refdb950bdc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *NVector) Deref() {
	if x.refdb950bdc == nil {
		return
	}
	x.Content = (unsafe.Pointer)(unsafe.Pointer(x.refdb950bdc.content))
	x.Ops = *NewGenericNVectorOpsRef(unsafe.Pointer(&x.refdb950bdc.ops))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *NVectorS) Ref() **C.N_Vector {
	if x == nil {
		return nil
	}
	return x.ref86072daf
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *NVectorS) Free() {
	if x != nil && x.allocs86072daf != nil {
		x.allocs86072daf.(*cgoAllocMap).Free()
		x.ref86072daf = nil
	}
}

// NewNVectorSRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewNVectorSRef(ref unsafe.Pointer) *NVectorS {
	if ref == nil {
		return nil
	}
	obj := new(NVectorS)
	obj.ref86072daf = (**C.N_Vector)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *NVectorS) PassRef() (**C.N_Vector, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref86072daf != nil {
		return x.ref86072daf, nil
	}
	mem86072daf := allocNVectorMemory(1)
	ref86072daf := (*C.N_Vector)(mem86072daf)
	allocs86072daf := new(cgoAllocMap)
	allocs86072daf.Add(mem86072daf)

	var ccontent_allocs *cgoAllocMap
	ref86072daf.content, ccontent_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.Content)), cgoAllocsUnknown
	allocs86072daf.Borrow(ccontent_allocs)

	var cops_allocs *cgoAllocMap
	ref86072daf.ops, cops_allocs = x.Ops.PassValue()
	allocs86072daf.Borrow(cops_allocs)

	x.ref86072daf = ref86072daf
	x.allocs86072daf = allocs86072daf
	return ref86072daf, allocs86072daf

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x NVectorS) PassValue() (*C.N_Vector, *cgoAllocMap) {
	if x.ref86072daf != nil {
		return *x.ref86072daf, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *NVectorS) Deref() {
	if x.ref86072daf == nil {
		return
	}
	x.Content = (unsafe.Pointer)(unsafe.Pointer(x.ref86072daf.content))
	x.Ops = *NewGenericNVectorOpsRef(unsafe.Pointer(&x.ref86072daf.ops))
}

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackArgSNVector transforms a sliced Go data structure into plain C format.
func unpackArgSNVector(x []NVector) (unpacked *C.N_Vector, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.N_Vector) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocNVectorMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: mem0,
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.N_Vector)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.N_Vector)(h.Data)
	return
}

// packSNVector reads sliced Go data structure out from plain C format.
func packSNVector(v []NVector, ptr0 *C.N_Vector) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfNVectorValue]C.N_Vector)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewNVectorRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSFILE transforms a sliced Go data structure into plain C format.
func unpackArgSFILE(x []FILE) (unpacked *C.FILE, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.FILE) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocFILEMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: mem0,
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.FILE)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.FILE)(h.Data)
	return
}

// packSFILE reads sliced Go data structure out from plain C format.
func packSFILE(v []FILE, ptr0 *C.FILE) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfFILEValue]C.FILE)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewFILERef(unsafe.Pointer(&ptr1))
	}
}

// allocPNVectorMemory allocates memory for type *C.N_Vector in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPNVectorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPNVectorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPNVectorValue = unsafe.Sizeof([1]*C.N_Vector{})

// unpackArgSSNVector transforms a sliced Go data structure into plain C format.
func unpackArgSSNVector(x [][]NVector) (unpacked **C.N_Vector, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.N_Vector) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPNVectorMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: mem0,
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.N_Vector)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocNVectorMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: mem1,
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.N_Vector)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.N_Vector)(h.Data)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.N_Vector)(h.Data)
	return
}

// packSSNVector reads sliced Go data structure out from plain C format.
func packSSNVector(v [][]NVector, ptr0 **C.N_Vector) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.N_Vector)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfNVectorValue]C.N_Vector)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *NewNVectorRef(unsafe.Pointer(&ptr2))
		}
	}
}
