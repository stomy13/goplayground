#include "textflag.h"

TEXT main·add(SB), NOSPLIT|NOFRAME, $0
    MOVD   a+0(FP), R0
    MOVD   b+8(FP), R1
    ADD    R0, R1, R2
    MOVD   R2, ret+16(FP)
    RET
