package ptx

const KERNMULC = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_0000400c_00000000-9_kernmulc.cpp3.i"
	.file	2 "/home/arne/src/nimble-cube/gpu/ptx/kernmulc.cu"

.visible .entry kernmulC(
	.param .u64 kernmulC_param_0,
	.param .u64 kernmulC_param_1,
	.param .u64 kernmulC_param_2,
	.param .u64 kernmulC_param_3,
	.param .u64 kernmulC_param_4,
	.param .u64 kernmulC_param_5,
	.param .u64 kernmulC_param_6,
	.param .u64 kernmulC_param_7,
	.param .u64 kernmulC_param_8,
	.param .u64 kernmulC_param_9,
	.param .u64 kernmulC_param_10,
	.param .u64 kernmulC_param_11,
	.param .u32 kernmulC_param_12
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<41>;
	.reg .f32 	%f<82>;
	.reg .s64 	%rd<51>;


	ld.param.u64 	%rd10, [kernmulC_param_0];
	ld.param.u64 	%rd11, [kernmulC_param_1];
	ld.param.u64 	%rd12, [kernmulC_param_2];
	ld.param.u64 	%rd13, [kernmulC_param_3];
	ld.param.u64 	%rd14, [kernmulC_param_4];
	ld.param.u64 	%rd15, [kernmulC_param_5];
	ld.param.u64 	%rd16, [kernmulC_param_6];
	ld.param.u64 	%rd17, [kernmulC_param_7];
	ld.param.u64 	%rd18, [kernmulC_param_8];
	ld.param.u64 	%rd19, [kernmulC_param_9];
	ld.param.u64 	%rd20, [kernmulC_param_10];
	ld.param.u64 	%rd21, [kernmulC_param_11];
	ld.param.u32 	%r2, [kernmulC_param_12];
	cvta.to.global.u64 	%rd1, %rd21;
	cvta.to.global.u64 	%rd2, %rd20;
	cvta.to.global.u64 	%rd3, %rd19;
	cvta.to.global.u64 	%rd4, %rd18;
	cvta.to.global.u64 	%rd5, %rd17;
	cvta.to.global.u64 	%rd6, %rd16;
	cvta.to.global.u64 	%rd7, %rd15;
	cvta.to.global.u64 	%rd8, %rd14;
	cvta.to.global.u64 	%rd9, %rd13;
	.loc 2 16 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r9, %r6, %r7, %r8;
	.loc 2 17 1
	shl.b32 	%r1, %r9, 1;
	.loc 2 19 1
	setp.ge.s32 	%p1, %r9, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd22, %rd10;
	.loc 2 20 1
	mul.wide.s32 	%rd23, %r1, 4;
	add.s64 	%rd24, %rd22, %rd23;
	.loc 2 21 1
	add.s32 	%r10, %r1, 1;
	mul.wide.s32 	%rd25, %r10, 4;
	add.s64 	%rd26, %rd22, %rd25;
	cvta.to.global.u64 	%rd27, %rd11;
	.loc 2 23 1
	add.s64 	%rd28, %rd27, %rd23;
	.loc 2 24 1
	add.s64 	%rd29, %rd27, %rd25;
	cvta.to.global.u64 	%rd30, %rd12;
	.loc 2 26 1
	add.s64 	%rd31, %rd30, %rd23;
	.loc 2 27 1
	add.s64 	%rd32, %rd30, %rd25;
	.loc 2 29 1
	add.s64 	%rd33, %rd9, %rd23;
	.loc 2 30 1
	add.s64 	%rd34, %rd8, %rd23;
	ld.global.f32 	%f1, [%rd34];
	.loc 2 31 1
	add.s64 	%rd35, %rd7, %rd23;
	ld.global.f32 	%f2, [%rd35];
	.loc 2 32 1
	add.s64 	%rd36, %rd9, %rd25;
	.loc 2 33 1
	add.s64 	%rd37, %rd8, %rd25;
	ld.global.f32 	%f3, [%rd37];
	.loc 2 34 1
	add.s64 	%rd38, %rd7, %rd25;
	ld.global.f32 	%f4, [%rd38];
	.loc 2 36 1
	add.s64 	%rd39, %rd6, %rd23;
	ld.global.f32 	%f5, [%rd39];
	.loc 2 37 1
	add.s64 	%rd40, %rd5, %rd23;
	.loc 2 38 1
	add.s64 	%rd41, %rd4, %rd23;
	.loc 2 39 1
	add.s64 	%rd42, %rd6, %rd25;
	ld.global.f32 	%f6, [%rd42];
	.loc 2 40 1
	add.s64 	%rd43, %rd5, %rd25;
	.loc 2 41 1
	add.s64 	%rd44, %rd4, %rd25;
	.loc 2 43 1
	add.s64 	%rd45, %rd3, %rd23;
	ld.global.f32 	%f7, [%rd45];
	.loc 2 44 1
	add.s64 	%rd46, %rd2, %rd23;
	ld.global.f32 	%f8, [%rd46];
	.loc 2 45 1
	add.s64 	%rd47, %rd1, %rd23;
	ld.global.f32 	%f9, [%rd47];
	.loc 2 46 1
	add.s64 	%rd48, %rd3, %rd25;
	ld.global.f32 	%f10, [%rd48];
	.loc 2 47 1
	add.s64 	%rd49, %rd2, %rd25;
	ld.global.f32 	%f11, [%rd49];
	.loc 2 48 1
	add.s64 	%rd50, %rd1, %rd25;
	ld.global.f32 	%f12, [%rd50];
	.loc 2 29 1
	ld.global.f32 	%f13, [%rd33];
	.loc 2 20 1
	ld.global.f32 	%f14, [%rd24];
	.loc 2 50 1
	mul.ftz.f32 	%f15, %f14, %f13;
	.loc 2 32 1
	ld.global.f32 	%f16, [%rd36];
	.loc 2 21 1
	ld.global.f32 	%f17, [%rd26];
	.loc 2 50 1
	mul.ftz.f32 	%f18, %f17, %f16;
	sub.ftz.f32 	%f19, %f15, %f18;
	.loc 2 38 1
	ld.global.f32 	%f20, [%rd41];
	.loc 2 23 1
	ld.global.f32 	%f21, [%rd28];
	.loc 2 50 1
	mul.ftz.f32 	%f22, %f21, %f20;
	.loc 2 41 1
	ld.global.f32 	%f23, [%rd44];
	.loc 2 24 1
	ld.global.f32 	%f24, [%rd29];
	.loc 2 50 1
	mul.ftz.f32 	%f25, %f24, %f23;
	sub.ftz.f32 	%f26, %f22, %f25;
	add.ftz.f32 	%f27, %f19, %f26;
	.loc 2 37 1
	ld.global.f32 	%f28, [%rd40];
	.loc 2 26 1
	ld.global.f32 	%f29, [%rd31];
	.loc 2 50 1
	mul.ftz.f32 	%f30, %f29, %f28;
	.loc 2 40 1
	ld.global.f32 	%f31, [%rd43];
	.loc 2 27 1
	ld.global.f32 	%f32, [%rd32];
	.loc 2 50 1
	mul.ftz.f32 	%f33, %f32, %f31;
	sub.ftz.f32 	%f34, %f30, %f33;
	add.ftz.f32 	%f35, %f27, %f34;
	st.global.f32 	[%rd24], %f35;
	.loc 2 51 1
	mul.ftz.f32 	%f36, %f17, %f13;
	fma.rn.ftz.f32 	%f37, %f14, %f16, %f36;
	mul.ftz.f32 	%f38, %f24, %f20;
	fma.rn.ftz.f32 	%f39, %f21, %f23, %f38;
	add.ftz.f32 	%f40, %f37, %f39;
	mul.ftz.f32 	%f41, %f32, %f28;
	fma.rn.ftz.f32 	%f42, %f29, %f31, %f41;
	add.ftz.f32 	%f43, %f40, %f42;
	st.global.f32 	[%rd26], %f43;
	.loc 2 53 1
	mul.ftz.f32 	%f44, %f14, %f9;
	mul.ftz.f32 	%f45, %f17, %f12;
	sub.ftz.f32 	%f46, %f44, %f45;
	mul.ftz.f32 	%f47, %f21, %f1;
	mul.ftz.f32 	%f48, %f24, %f3;
	sub.ftz.f32 	%f49, %f47, %f48;
	add.ftz.f32 	%f50, %f46, %f49;
	mul.ftz.f32 	%f51, %f29, %f5;
	mul.ftz.f32 	%f52, %f32, %f6;
	sub.ftz.f32 	%f53, %f51, %f52;
	add.ftz.f32 	%f54, %f50, %f53;
	st.global.f32 	[%rd28], %f54;
	.loc 2 54 1
	mul.ftz.f32 	%f55, %f17, %f9;
	fma.rn.ftz.f32 	%f56, %f14, %f12, %f55;
	mul.ftz.f32 	%f57, %f24, %f1;
	fma.rn.ftz.f32 	%f58, %f21, %f3, %f57;
	add.ftz.f32 	%f59, %f56, %f58;
	mul.ftz.f32 	%f60, %f32, %f5;
	fma.rn.ftz.f32 	%f61, %f29, %f6, %f60;
	add.ftz.f32 	%f62, %f59, %f61;
	st.global.f32 	[%rd29], %f62;
	.loc 2 56 1
	mul.ftz.f32 	%f63, %f14, %f8;
	mul.ftz.f32 	%f64, %f17, %f11;
	sub.ftz.f32 	%f65, %f63, %f64;
	mul.ftz.f32 	%f66, %f21, %f7;
	mul.ftz.f32 	%f67, %f24, %f10;
	sub.ftz.f32 	%f68, %f66, %f67;
	add.ftz.f32 	%f69, %f65, %f68;
	mul.ftz.f32 	%f70, %f29, %f2;
	mul.ftz.f32 	%f71, %f32, %f4;
	sub.ftz.f32 	%f72, %f70, %f71;
	add.ftz.f32 	%f73, %f69, %f72;
	st.global.f32 	[%rd31], %f73;
	.loc 2 57 1
	mul.ftz.f32 	%f74, %f17, %f8;
	fma.rn.ftz.f32 	%f75, %f14, %f11, %f74;
	mul.ftz.f32 	%f76, %f24, %f7;
	fma.rn.ftz.f32 	%f77, %f21, %f10, %f76;
	add.ftz.f32 	%f78, %f75, %f77;
	mul.ftz.f32 	%f79, %f32, %f2;
	fma.rn.ftz.f32 	%f80, %f29, %f4, %f79;
	add.ftz.f32 	%f81, %f78, %f80;
	st.global.f32 	[%rd32], %f81;

BB0_2:
	.loc 2 59 2
	ret;
}


`
