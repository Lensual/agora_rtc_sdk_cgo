package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "C_IAgoraParameter.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

// external key
/**
 * set the range of ports available for connection
 * @example "{\"rtc.udp_port_range\":[4500, 5000]}"
 */
const C_KEY_RTC_UDP_PORT_RANGE = C.C_KEY_RTC_UDP_PORT_RANGE

/**
 * set the list of ports available for connection
 * @example  "{\"rtc.udp_port_list\":[4501, 4502, 4503, 4504, 4505, 4506]}"
 */
const C_KEY_RTC_UDP_PORT_LIST = C.C_KEY_RTC_UDP_PORT_LIST

/**
 * get the fd of sending socket available for connection
 * note: set method is not supported.
 */
const C_KEY_RTC_UDP_SEND_FD = C.C_KEY_RTC_UDP_SEND_FD

/**
 * set the video encoder mode (hardware or software)
 */
const C_KEY_RTC_VIDEO_ENABLED_HW_ENCODER = C.C_KEY_RTC_VIDEO_ENABLED_HW_ENCODER
const C_KEY_RTC_VIDEO_HARDWARE_ENCODEING = C.C_KEY_RTC_VIDEO_HARDWARE_ENCODEING
const C_KEY_RTC_VIDEO_H264_HWENC = C.C_KEY_RTC_VIDEO_H264_HWENC

/**
 * set the hardware video encoder provider (nv for nvidia or qsv for intel)
 */
const C_KEY_RTC_VIDEO_HW_ENCODER_PROVIDER = C.C_KEY_RTC_VIDEO_HW_ENCODER_PROVIDER

/**
 * set the video decoder mode (hardware or software)
 */
const C_KEY_RTC_VIDEO_ENABLED_HW_DECODER = C.C_KEY_RTC_VIDEO_ENABLED_HW_DECODER
const C_KEY_RTC_VIDEO_HARDWARE_DECODING = C.C_KEY_RTC_VIDEO_HARDWARE_DECODING

/**
 * set the hardware video decoder provider (h264_cuvid(default) or h264_qsv)
 */
const C_KEY_RTC_VIDEO_HW_DECODER_PROVIDER = C.C_KEY_RTC_VIDEO_HW_DECODER_PROVIDER

/**
 * override the lua policy
 */
const C_KEY_RTC_VIDEO_OVERRIDE_SMALLVIDEO_NOT_USE_HWENC_POLICY = C.C_KEY_RTC_VIDEO_OVERRIDE_SMALLVIDEO_NOT_USE_HWENC_POLICY

/**
 * enable/disable video packet retransmission, enabled by default
 */
const C_KEY_RTC_VIDEO_RESEND = C.C_KEY_RTC_VIDEO_RESEND

/**
 * enable/disable audio packet retransmission, enabled by default
 */
const C_KEY_RTC_AUDIO_RESEND = C.C_KEY_RTC_AUDIO_RESEND

/**
 * set the bitrate ratio for video
 */
const C_KEY_RTC_VIDEO_BITRATE_ADJUST_RATIO = C.C_KEY_RTC_VIDEO_BITRATE_ADJUST_RATIO

/**
 * set the minbitrate / bitrate ratio for video
 */
const C_KEY_RTC_VIDEO_MINBITRATE_RATIO = C.C_KEY_RTC_VIDEO_MINBITRATE_RATIO

/**
 * set the degradation preference
 */
const C_KEY_RTC_VIDEO_DEGRADATION_PREFERENCE = C.C_KEY_RTC_VIDEO_DEGRADATION_PREFERENCE

/**
 * set the degradation fps down step
 */

const C_KEY_RTC_VIDEO_DEGRADATION_FPS_DOWN_STEP = C.C_KEY_RTC_VIDEO_DEGRADATION_FPS_DOWN_STEP

/**
 * set the degradation fps up step
 */
const C_KEY_RTC_VIDEO_DEGRADATION_FPS_UP_STEP = C.C_KEY_RTC_VIDEO_DEGRADATION_FPS_UP_STEP

/**
 * set the duration ms for connection lost callback
 */
const C_KEY_RTC_CONNECTION_LOST_PERIOD = C.C_KEY_RTC_CONNECTION_LOST_PERIOD

/**
 * set the local ip
 */
const C_KEY_RTC_LOCAL_IP = C.C_KEY_RTC_LOCAL_IP

/**
 * set the network interface
 */
const C_KEY_RTC_NETWORK_INTERFACE = C.C_KEY_RTC_NETWORK_INTERFACE

/**
 * set the video codec type, such as "H264", "JPEG"
 */
const C_KEY_RTC_VIDEO_CODEC_TYPE = C.C_KEY_RTC_VIDEO_CODEC_TYPE
const C_KEY_RTC_VIDEO_MINOR_STREAM_CODEC_TYPE = C.C_KEY_RTC_VIDEO_MINOR_STREAM_CODEC_TYPE
const C_KEY_RTC_VIDEO_CODEC_INDEX = C.C_KEY_RTC_VIDEO_CODEC_INDEX

/**
 * only use average QP for quality scaling
 */
const C_KEY_RTC_VIDEO_QUALITY_SCALE_ONLY_ON_AVERAGE_QP = C.C_KEY_RTC_VIDEO_QUALITY_SCALE_ONLY_ON_AVERAGE_QP

/**
 * low bound for quality scaling
 */
const C_KEY_RTC_VIDEO_H264_QP_THRESHOLD_LOW = C.C_KEY_RTC_VIDEO_H264_QP_THRESHOLD_LOW

/**
 * high bound for quality scaling
 */
const C_KEY_RTC_VIDEO_H264_QP_THRESHOLD_HIGH = C.C_KEY_RTC_VIDEO_H264_QP_THRESHOLD_HIGH

// #region agora

// #region agora::util

//   // template <class T>
//   // class CopyableAutoPtr;
//   typedef struct C_AutoPtr C_CopyableAutoPtr;

//   // class C_IString;
//   // typedef CopyableAutoPtr<IString> AString;
//   typedef C_CopyableAutoPtr C_AString;

// #endregion agora::util

// #region agora::base

type IAgoraParameter C.C_IAgoraParameter

// #region IAgoraParameter

/**
 * release the resource
 */
func (this_ *IAgoraParameter) Release() {
	C.C_IAgoraParameter_release(unsafe.Pointer(this_))
}

//   /**
//    * set bool value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_setBool(C_IAgoraParameter *this_, const char *key, bool value);

//   /**
//    * set int value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_setInt(C_IAgoraParameter *this_, const char *key, int value);

//   /**
//    * set unsigned int value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_setUInt(C_IAgoraParameter *this_, const char *key, unsigned int value);

//   /**
//    * set double value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_setNumber(C_IAgoraParameter *this_, const char *key, double value);

//   /**
//    * set string value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_setString(C_IAgoraParameter *this_, const char *key, const char *value);

//   /**
//    * set object value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_setObject(C_IAgoraParameter *this_, const char *key, const char *value);

//   /**
//    * set array value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_setArray(C_IAgoraParameter *this_, const char *key, const char *value);
//   /**
//    * get bool value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in, out] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_getBool(C_IAgoraParameter *this_, const char *key, bool *value);

//   /**
//    * get int value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in, out] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_getInt(C_IAgoraParameter *this_, const char *key, int *value);

//   /**
//    * get unsigned int value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in, out] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_getUInt(C_IAgoraParameter *this_, const char *key, unsigned int *value);

//   /**
//    * get double value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in, out] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_getNumber(C_IAgoraParameter *this_, const char *key, double *value);

//   /**
//    * get string value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in, out] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_getString(C_IAgoraParameter *this_, const char *key, C_AString *value);

//   /**
//    * get a child object value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in, out] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_getObject(C_IAgoraParameter *this_, const char *key, C_AString *value);

//   /**
//    * get array value of the json
//    * @param [in] key
//    *        the key name
//    * @param [in, out] value
//    *        the value
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_getArray(C_IAgoraParameter *this_, const char *key, const char *args, C_AString *value);

//   /**
//    * set parameters of the sdk or engine
//    * @param [in] parameters
//    *        the parameters
//    * @return return 0 if success or an error code
//    */
//   int C_IAgoraParameter_setParameters(C_IAgoraParameter *this_, const char *parameters);

//   int C_IAgoraParameter_convertPath(C_IAgoraParameter *this_, const char *filePath, C_AString *value);

//   void C_IAgoraParameter_Delete(C_IAgoraParameter *this_);

// #endregion agora::C_IAgoraParameter

// #endregion agora::base

// #endregion agora
