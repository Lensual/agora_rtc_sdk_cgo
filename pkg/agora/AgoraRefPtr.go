package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk-prefix/src/agora_rtc_sdk/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk-prefix/src/agora_rtc_sdk/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "C_AgoraRefPtr.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type AgoraRefPtr C.C_agora_refptr

func (this_ AgoraRefPtr) Get() unsafe.Pointer {
	return this_.ptr_
}
