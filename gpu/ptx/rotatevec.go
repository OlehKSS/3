package ptx

const ROTATEVEC = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_000040a2_00000000-9_rotatevec.cpp3.i"
	.file	2 "/home/arne/src/nimble-cube/gpu/ptx/rotatevec.cu"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"
	.file	4 "/usr/local/cuda/bin/../include/device_functions.h"

.visible .entry rotatevec(
	.param .u64 rotatevec_param_0,
	.param .u64 rotatevec_param_1,
	.param .u64 rotatevec_param_2,
	.param .u64 rotatevec_param_3,
	.param .u64 rotatevec_param_4,
	.param .u64 rotatevec_param_5,
	.param .f32 rotatevec_param_6,
	.param .u32 rotatevec_param_7
)
{
	.reg .pred 	%p<3>;
	.reg .s32 	%r<18>;
	.reg .f32 	%f<19>;
	.reg .s64 	%rd<20>;


	ld.param.u64 	%rd7, [rotatevec_param_0];
	ld.param.u64 	%rd8, [rotatevec_param_1];
	ld.param.u64 	%rd9, [rotatevec_param_2];
	ld.param.u64 	%rd10, [rotatevec_param_3];
	ld.param.u64 	%rd11, [rotatevec_param_4];
	ld.param.u64 	%rd12, [rotatevec_param_5];
	ld.param.f32 	%f1, [rotatevec_param_6];
	ld.param.u32 	%r2, [rotatevec_param_7];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd11;
	cvta.to.global.u64 	%rd4, %rd8;
	cvta.to.global.u64 	%rd5, %rd10;
	cvta.to.global.u64 	%rd6, %rd7;
	.loc 2 6 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 7 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	.loc 2 9 1
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd6, %rd13;
	add.s64 	%rd15, %rd5, %rd13;
	ld.global.f32 	%f2, [%rd15];
	ld.global.f32 	%f3, [%rd14];
	fma.rn.ftz.f32 	%f4, %f2, %f1, %f3;
	.loc 2 10 1
	add.s64 	%rd16, %rd4, %rd13;
	add.s64 	%rd17, %rd3, %rd13;
	ld.global.f32 	%f5, [%rd17];
	ld.global.f32 	%f6, [%rd16];
	fma.rn.ftz.f32 	%f7, %f5, %f1, %f6;
	.loc 2 11 1
	add.s64 	%rd18, %rd2, %rd13;
	add.s64 	%rd19, %rd1, %rd13;
	ld.global.f32 	%f8, [%rd19];
	ld.global.f32 	%f9, [%rd18];
	fma.rn.ftz.f32 	%f10, %f8, %f1, %f9;
	.loc 2 13 1
	mul.ftz.f32 	%f11, %f7, %f7;
	fma.rn.ftz.f32 	%f12, %f4, %f4, %f11;
	fma.rn.ftz.f32 	%f13, %f10, %f10, %f12;
	.loc 3 996 5
	sqrt.approx.ftz.f32 	%f14, %f13;
	.loc 2 14 1
	setp.eq.ftz.f32 	%p2, %f14, 0f00000000;
	selp.f32 	%f15, 0f3F800000, %f14, %p2;
	.loc 3 749 5
	div.approx.ftz.f32 	%f16, %f4, %f15;
	.loc 2 16 1
	st.global.f32 	[%rd14], %f16;
	.loc 3 749 5
	div.approx.ftz.f32 	%f17, %f7, %f15;
	.loc 2 17 1
	st.global.f32 	[%rd16], %f17;
	.loc 3 749 5
	div.approx.ftz.f32 	%f18, %f10, %f15;
	.loc 2 18 1
	st.global.f32 	[%rd18], %f18;

BB0_2:
	.loc 2 20 2
	ret;
}


`
