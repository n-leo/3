package ptx

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"code.google.com/p/mx3/streams"
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var madd2_code cu.Function

type madd2_args struct {
	arg_dst  cu.DevicePtr
	arg_src1 cu.DevicePtr
	arg_fac1 float32
	arg_src2 cu.DevicePtr
	arg_fac2 float32
	arg_N    int
	argptr   [6]unsafe.Pointer
}

// CUDA kernel wrapper for madd2.
// The kernel is launched in a separate stream so that it can be parallel with memcpys etc.
// The stream is synchronized before this call returns.
func K_madd2(dst cu.DevicePtr, src1 cu.DevicePtr, fac1 float32, src2 cu.DevicePtr, fac2 float32, N int, gridDim, blockDim cu.Dim3) {
	if madd2_code == 0 {
		madd2_code = cu.ModuleLoadData(madd2_ptx).GetFunction("madd2")
	}

	var a madd2_args

	a.arg_dst = dst
	a.argptr[0] = unsafe.Pointer(&a.arg_dst)
	a.arg_src1 = src1
	a.argptr[1] = unsafe.Pointer(&a.arg_src1)
	a.arg_fac1 = fac1
	a.argptr[2] = unsafe.Pointer(&a.arg_fac1)
	a.arg_src2 = src2
	a.argptr[3] = unsafe.Pointer(&a.arg_src2)
	a.arg_fac2 = fac2
	a.argptr[4] = unsafe.Pointer(&a.arg_fac2)
	a.arg_N = N
	a.argptr[5] = unsafe.Pointer(&a.arg_N)

	args := a.argptr[:]
	str := streams.Get()
	cu.LaunchKernel(madd2_code, gridDim.X, gridDim.Y, gridDim.Z, blockDim.X, blockDim.Y, blockDim.Z, 0, str, args)
	streams.SyncAndRecycle(str)
}

const madd2_ptx = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry madd2(
	.param .u64 madd2_param_0,
	.param .u64 madd2_param_1,
	.param .f32 madd2_param_2,
	.param .u64 madd2_param_3,
	.param .f32 madd2_param_4,
	.param .u32 madd2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<12>;
	.reg .f32 	%f<7>;
	.reg .s64 	%rd<11>;


	ld.param.u64 	%rd4, [madd2_param_0];
	ld.param.u64 	%rd5, [madd2_param_1];
	ld.param.f32 	%f1, [madd2_param_2];
	ld.param.u64 	%rd6, [madd2_param_3];
	ld.param.f32 	%f2, [madd2_param_4];
	ld.param.u32 	%r2, [madd2_param_5];
	cvta.to.global.u64 	%rd1, %rd4;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 2 4 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 5 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	.loc 2 6 1
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd3, %rd7;
	ld.global.f32 	%f3, [%rd8];
	add.s64 	%rd9, %rd2, %rd7;
	ld.global.f32 	%f4, [%rd9];
	mul.f32 	%f5, %f4, %f2;
	fma.rn.f32 	%f6, %f3, %f1, %f5;
	add.s64 	%rd10, %rd1, %rd7;
	st.global.f32 	[%rd10], %f6;

BB0_2:
	.loc 2 8 2
	ret;
}


`
