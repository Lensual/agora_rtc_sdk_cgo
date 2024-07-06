package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "C_AgoraMediaBase.h"
#include <stdlib.h>
*/
import "C"

// #region agora

// #region agora::rtc

type Uid_t C.C_uid_t
type Track_id_t C.C_track_id_t
type Conn_id_t C.C_conn_id_t
type Video_track_id_t C.C_video_track_id_t

// const INVALID_TRACK_ID = C.C_INVALID_TRACK_ID
// const DEFAULT_CONNECTION_ID = C.C_DEFAULT_CONNECTION_ID
// const DUMMY_CONNECTION_ID = C.C_DUMMY_CONNECTION_ID

const kAdmMaxDeviceNameSize = C.C_kAdmMaxDeviceNameSize
const kAdmMaxGuidSize = C.C_kAdmMaxGuidSize

/**
 * Audio routes.
 */
type AudioRoute C.enum_C_AudioRoute

const (
	/**
	 * -1: The default audio route.
	 */
	ROUTE_DEFAULT AudioRoute = C.ROUTE_DEFAULT
	/**
	 * The headset.
	 */
	ROUTE_HEADSET AudioRoute = C.ROUTE_HEADSET
	/**
	 * The earpiece.
	 */
	ROUTE_EARPIECE AudioRoute = C.ROUTE_EARPIECE
	/**
	 * The headset with no microphone.
	 */
	ROUTE_HEADSETNOMIC AudioRoute = C.ROUTE_HEADSETNOMIC
	/**
	 * The speakerphone.
	 */
	ROUTE_SPEAKERPHONE AudioRoute = C.ROUTE_SPEAKERPHONE
	/**
	 * The loudspeaker.
	 */
	ROUTE_LOUDSPEAKER AudioRoute = C.ROUTE_LOUDSPEAKER
	/**
	 * The Bluetooth headset.
	 */
	ROUTE_HEADSETBLUETOOTH AudioRoute = C.ROUTE_HEADSETBLUETOOTH
	/**
	 * The USB
	 */
	ROUTE_USB AudioRoute = C.ROUTE_USB
	/**
	 * The HDMI
	 */
	ROUTE_HDMI AudioRoute = C.ROUTE_HDMI
	/**
	 * The DISPLAYPORT
	 */
	ROUTE_DISPLAYPORT AudioRoute = C.ROUTE_DISPLAYPORT
	/**
	 * The AIRPLAY
	 */
	ROUTE_AIRPLAY AudioRoute = C.ROUTE_AIRPLAY
	/**
	 * The VIRTUAL SOUND CARD
	 */
	ROUTE_VIRTUAL AudioRoute = C.ROUTE_VIRTUAL
	/**
	 * The Continuity
	 */
	ROUTE_CONTINUITY AudioRoute = C.ROUTE_CONTINUITY
)

/**
 * Bytes per sample
 */
type BYTES_PER_SAMPLE C.enum_C_BYTES_PER_SAMPLE

const (
	/**
	 * two bytes per sample
	 */
	TWO_BYTES_PER_SAMPLE BYTES_PER_SAMPLE = C.TWO_BYTES_PER_SAMPLE
)

type AudioParameters C.struct_C_AudioParameters

// #region AudioParameters

func (this_ *AudioParameters) GetSampleRate() int {
	return int(this_.sample_rate)
}
func (this_ *AudioParameters) SetSampleRate(sample_rate int) {
	this_.sample_rate = C.int(sample_rate)
}

func (this_ *AudioParameters) GetChannels() uint {
	return uint(this_.channels)
}
func (this_ *AudioParameters) SetChannels(channels uint) {
	this_.channels = C.size_t(channels)
}

func (this_ *AudioParameters) GetFramesPerBuffer() uint {
	return uint(this_.frames_per_buffer)
}
func (this_ *AudioParameters) SetFramesPerBuffer(frames_per_buffer uint) {
	this_.frames_per_buffer = C.size_t(frames_per_buffer)
}

func NewAudioParameters() *AudioParameters {
	return (*AudioParameters)(C.C_AudioParameters_New())
}
func (this_ *AudioParameters) Delete() {
	C.C_AudioParameters_Delete((*C.struct_C_AudioParameters)(this_))
}

// #endregion AudioParameters

/**
 * The use mode of the audio data.
 */
type RAW_AUDIO_FRAME_OP_MODE_TYPE C.enum_C_RAW_AUDIO_FRAME_OP_MODE_TYPE

const (
	/** 0: Read-only mode: Users only read the data from `AudioFrame` without modifying anything.
	 * For example, when users acquire the data with the Agora SDK, then start the media push.
	 */
	RAW_AUDIO_FRAME_OP_MODE_READ_ONLY RAW_AUDIO_FRAME_OP_MODE_TYPE = C.RAW_AUDIO_FRAME_OP_MODE_READ_ONLY

	/** 2: Read and write mode: Users read the data from `AudioFrame`, modify it, and then play it.
	 * For example, when users have their own audio-effect processing module and perform some voice pre-processing, such as a voice change.
	 */
	RAW_AUDIO_FRAME_OP_MODE_READ_WRITE RAW_AUDIO_FRAME_OP_MODE_TYPE = C.RAW_AUDIO_FRAME_OP_MODE_READ_WRITE
)

/**
 * The mix policy of the track per connection.
 * For default, this connections mix policy is REMOTE | LOCAL(3).
 */
type TRACK_AUDIO_MIXED_POLICY_TYPE C.enum_C_TRACK_AUDIO_MIXED_POLICY_TYPE

const (
	/** 0: All remote track that created by remote UID would be mixed.
	 */
	TRACK_AUDIO_MIXED_LOCAL TRACK_AUDIO_MIXED_POLICY_TYPE = C.TRACK_AUDIO_MIXED_LOCAL
	/** 1: All local publish track would be mixed, microphone track, mpk track etc.
	 */
	TRACK_AUDIO_MIXED_REMOTE TRACK_AUDIO_MIXED_POLICY_TYPE = C.TRACK_AUDIO_MIXED_REMOTE
)

//   /**
//    * The struct of AudioDeviceInfo.
//    *
//    * @note
//    * This struct applies to Windows and macOS only.
//    */
//   struct C_AudioDeviceInfo
//   {
//     /**
//      * The name of the device. The maximum name size is 128 bytes. The default value is 0.
//      */
//     char deviceName[C_kAdmMaxDeviceNameSize];
//     /**
//      * The ID of the device. The maximum size is 128 bytes. The default value is 0.
//      */
//     char deviceId[C_kAdmMaxGuidSize];
//     /**
//      * Determines whether the current device is selected for audio capturing or playback.
//      * - true: Select the current device for audio capturing or playback.
//      * - false: (Default) Do not select the current device for audio capturing or playback.
//      */
//     bool isCurrentSelected;
//     /**
//      * Determines whether the current device is the audio playout device.
//      * - true: (Default) The current device is the playout device.
//      * - false: The current device is not the playout device.
//      */
//     bool isPlayoutDevice;

//     /**
//      * The routing of the device. The default value is ROUTE_DEFAULT.
//      */
//     enum C_AudioRoute routing;

//     // AudioDeviceInfo &operator=(const AudioDeviceInfo &other)
//     // {
//     //   isCurrentSelected = other.isCurrentSelected;
//     //   isPlayoutDevice = other.isPlayoutDevice;
//     //   routing = other.routing;
//     //   memcpy(deviceName, other.deviceName, sizeof(deviceName));
//     //   memcpy(deviceId, other.deviceId, sizeof(deviceId));
//     //   return *this;
//     // }
//   };
//   struct C_AudioDeviceInfo *C_AudioDeviceInfo_New();
//   void C_AudioDeviceInfo_Delete(struct C_AudioDeviceInfo *this_);

// #endregion agora::rtc

// #region agora::media

/**
 * The type of media device.
 */
type MEDIA_SOURCE_TYPE C.enum_C_MEDIA_SOURCE_TYPE

const (
	/**
	 * 0: The audio playback device.
	 */
	AUDIO_PLAYOUT_SOURCE MEDIA_SOURCE_TYPE = C.AUDIO_PLAYOUT_SOURCE
	/**
	 * 1: Microphone.
	 */
	AUDIO_RECORDING_SOURCE MEDIA_SOURCE_TYPE = C.AUDIO_RECORDING_SOURCE
	/**
	 * 2: Video captured by primary camera.
	 */
	PRIMARY_CAMERA_SOURCE MEDIA_SOURCE_TYPE = C.PRIMARY_CAMERA_SOURCE
	/**
	 * 3: Video captured by secondary camera.
	 */
	SECONDARY_CAMERA_SOURCE MEDIA_SOURCE_TYPE = C.SECONDARY_CAMERA_SOURCE
	/**
	 * 4: Video captured by primary screen capturer.
	 */
	PRIMARY_SCREEN_SOURCE MEDIA_SOURCE_TYPE = C.PRIMARY_SCREEN_SOURCE
	/**
	 * 5: Video captured by secondary screen capturer.
	 */
	SECONDARY_SCREEN_SOURCE MEDIA_SOURCE_TYPE = C.SECONDARY_SCREEN_SOURCE
	/**
	 * 6: Video captured by custom video source.
	 */
	CUSTOM_VIDEO_SOURCE MEDIA_SOURCE_TYPE = C.CUSTOM_VIDEO_SOURCE
	/**
	 * 7: Video for media player sharing.
	 */
	MEDIA_PLAYER_SOURCE MEDIA_SOURCE_TYPE = C.MEDIA_PLAYER_SOURCE
	/**
	 * 8: Video for png image.
	 */
	RTC_IMAGE_PNG_SOURCE MEDIA_SOURCE_TYPE = C.RTC_IMAGE_PNG_SOURCE
	/**
	 * 9: Video for jpeg image.
	 */
	RTC_IMAGE_JPEG_SOURCE MEDIA_SOURCE_TYPE = C.RTC_IMAGE_JPEG_SOURCE
	/**
	 * 10: Video for gif image.
	 */
	RTC_IMAGE_GIF_SOURCE MEDIA_SOURCE_TYPE = C.RTC_IMAGE_GIF_SOURCE
	/**
	 * 11: Remote video received from network.
	 */
	REMOTE_VIDEO_SOURCE MEDIA_SOURCE_TYPE = C.REMOTE_VIDEO_SOURCE
	/**
	 * 12: Video for transcoded.
	 */
	TRANSCODED_VIDEO_SOURCE MEDIA_SOURCE_TYPE = C.TRANSCODED_VIDEO_SOURCE
	/**
	 * 13: Video for voice drive.
	 */
	SPEECH_DRIVEN_VIDEO_SOURCE MEDIA_SOURCE_TYPE = C.SPEECH_DRIVEN_VIDEO_SOURCE
	/**
	 * 100: Internal Usage only.
	 */
	UNKNOWN_MEDIA_SOURCE MEDIA_SOURCE_TYPE = C.UNKNOWN_MEDIA_SOURCE
)

/** Definition of contentinspect
 */
const MAX_CONTENT_INSPECT_MODULE_COUNT = C.C_MAX_CONTENT_INSPECT_MODULE_COUNT

type CONTENT_INSPECT_RESULT C.enum_C_CONTENT_INSPECT_RESULT

const (
	CONTENT_INSPECT_NEUTRAL CONTENT_INSPECT_RESULT = C.CONTENT_INSPECT_NEUTRAL
	CONTENT_INSPECT_SEXY    CONTENT_INSPECT_RESULT = C.CONTENT_INSPECT_SEXY
	CONTENT_INSPECT_PORN    CONTENT_INSPECT_RESULT = C.CONTENT_INSPECT_PORN
)

type CONTENT_INSPECT_TYPE C.enum_C_CONTENT_INSPECT_TYPE

const (
	/**
	 * (Default) content inspect type invalid
	 */
	CONTENT_INSPECT_INVALID CONTENT_INSPECT_TYPE = C.CONTENT_INSPECT_INVALID
	/**
	 * @deprecated
	 * Content inspect type moderation
	 */
	CONTENT_INSPECT_MODERATION CONTENT_INSPECT_TYPE = C.CONTENT_INSPECT_MODERATION
	/**
	 * Content inspect type supervise
	 */
	CONTENT_INSPECT_SUPERVISION CONTENT_INSPECT_TYPE = C.CONTENT_INSPECT_SUPERVISION
	/**
	 * Content inspect type image moderation
	 */
	CONTENT_INSPECT_IMAGE_MODERATION CONTENT_INSPECT_TYPE = C.CONTENT_INSPECT_IMAGE_MODERATION
)

//   struct C_ContentInspectModule
//   {
//     /**
//      * The content inspect module type.
//      */
//     enum C_CONTENT_INSPECT_TYPE type;
//     /**The content inspect frequency, default is 0 second.
//      * the frequency <= 0 is invalid.
//      */
//     unsigned int interval;
//   };
//   struct C_ContentInspectModule *C_ContentInspectModule_New();
//   void C_ContentInspectModule_Delete(struct C_ContentInspectModule *this_);

//   /** Definition of ContentInspectConfig.
//    */
//   struct C_ContentInspectConfig
//   {
//     const char *extraInfo;
//     /**
//      * The specific server configuration for image moderation. Please contact technical support.
//      */
//     const char *serverConfig;
//     /**The content inspect modules, max length of modules is 32.
//      * the content(snapshot of send video stream, image) can be used to max of 32 types functions.
//      */
//     struct C_ContentInspectModule modules[C_MAX_CONTENT_INSPECT_MODULE_COUNT];
//     /**The content inspect module count.
//      */
//     int moduleCount;
//     //  ContentInspectConfig& operator=(const ContentInspectConfig& rth)
//     // {
//     //       extraInfo = rth.extraInfo;
//     //       serverConfig = rth.serverConfig;
//     //       moduleCount = rth.moduleCount;
//     // 	memcpy(&modules, &rth.modules,  MAX_CONTENT_INSPECT_MODULE_COUNT * sizeof(ContentInspectModule));
//     // 	return *this;
//     // }
//   };
//   struct C_ContentInspectConfig *C_ContentInspectConfig_New();
//   void C_ContentInspectConfig_Delete(struct C_ContentInspectConfig *this_);

// #region agora::media::base

//   typedef void *C_view_t;

//   typedef const char *C_user_id_t;

//   // static const uint8_t kMaxCodecNameLength = 50;

//   /**
//    * The definition of the PacketOptions struct, which contains infomation of the packet
//    * in the RTP (Real-time Transport Protocal) header.
//    */
//   struct C_PacketOptions
//   {
//     /**
//      * The timestamp of the packet.
//      */
//     uint32_t timestamp;
//     // Audio level indication.
//     uint8_t audioLevelIndication;
//   };
//   struct C_PacketOptions *C_PacketOptions_New();
//   void C_PacketOptions_Delete(struct C_PacketOptions *this_);

//   /**
//    * The detailed information of the incoming audio encoded frame.
//    */
//   struct C_AudioEncodedFrameInfo
//   {
//     /**
//      * The send time of the packet.
//      */
//     uint64_t sendTs;
//     /**
//      * The codec of the packet.
//      */
//     uint8_t codec;
//   };
//   struct C_AudioEncodedFrameInfo *C_AudioEncodedFrameInfo_New();
//   void C_AudioEncodedFrameInfo_Delete(struct C_AudioEncodedFrameInfo *this_);

//   /**
//    * The detailed information of the incoming audio frame in the PCM format.
//    */
//   struct C_AudioPcmFrame
//   {
//     /**
//      * The buffer size of the PCM audio frame.
//      */
//     C_OPTIONAL_ENUM_SIZE_T{
//         // Stereo, 32 kHz, 60 ms (2 * 32 * 60)
//         /**
//          * The max number of the samples of the data.
//          *
//          * When the number of audio channel is two, the sample rate is 32 kHZ,
//          * the buffer length of the data is 60 ms, the number of the samples of the data is 3840 (2 x 32 x 60).
//          */
//         kMaxDataSizeSamples = 3840,
//         /** The max number of the bytes of the data. */
//         kMaxDataSizeBytes = kMaxDataSizeSamples * sizeof(int16_t),
//     };

//     /** The timestamp (ms) of the audio frame.
//      */
//     int64_t capture_timestamp;
//     /** The number of samples per channel.
//      */
//     size_t samples_per_channel_;
//     /** The sample rate (Hz) of the audio data.
//      */
//     int sample_rate_hz_;
//     /** The channel number.
//      */
//     size_t num_channels_;
//     /** The number of bytes per sample.
//      */
//     enum C_BYTES_PER_SAMPLE bytes_per_sample;
//     /** The audio frame data. */
//     int16_t data_[kMaxDataSizeSamples];

//     // AudioPcmFrame& operator=(const AudioPcmFrame& src) {
//     //   if(this == &src) {
//     //     return *this;
//     //   }

//     //   this->capture_timestamp = src.capture_timestamp;
//     //   this->samples_per_channel_ = src.samples_per_channel_;
//     //   this->sample_rate_hz_ = src.sample_rate_hz_;
//     //   this->bytes_per_sample = src.bytes_per_sample;
//     //   this->num_channels_ = src.num_channels_;

//     //   size_t length = src.samples_per_channel_ * src.num_channels_;
//     //   if (length > kMaxDataSizeSamples) {
//     //     length = kMaxDataSizeSamples;
//     //   }

//     //   memcpy(this->data_, src.data_, length * sizeof(int16_t));

//     //   return *this;
//     // }
//   };
//   struct C_AudioPcmFrame *C_AudioPcmFrame_New();
//   void C_AudioPcmFrame_Delete(struct C_AudioPcmFrame *this_);

/** Audio dual-mono output mode
 */
type AUDIO_DUAL_MONO_MODE C.enum_C_AUDIO_DUAL_MONO_MODE

const (
	/**< ChanLOut=ChanLin, ChanRout=ChanRin */
	AUDIO_DUAL_MONO_STEREO AUDIO_DUAL_MONO_MODE = C.AUDIO_DUAL_MONO_STEREO
	/**< ChanLOut=ChanRout=ChanLin */
	AUDIO_DUAL_MONO_L AUDIO_DUAL_MONO_MODE = C.AUDIO_DUAL_MONO_L
	/**< ChanLOut=ChanRout=ChanRin */
	AUDIO_DUAL_MONO_R AUDIO_DUAL_MONO_MODE = C.AUDIO_DUAL_MONO_R
	/**< ChanLout=ChanRout=(ChanLin+ChanRin)/2 */
	AUDIO_DUAL_MONO_MIX AUDIO_DUAL_MONO_MODE = C.AUDIO_DUAL_MONO_MIX
)

/**
 * Video pixel formats.
 */
type VIDEO_PIXEL_FORMAT C.enum_C_VIDEO_PIXEL_FORMAT

const (
	/**
	 * 0: Default format.
	 */
	VIDEO_PIXEL_DEFAULT VIDEO_PIXEL_FORMAT = C.VIDEO_PIXEL_DEFAULT
	/**
	 * 1: I420.
	 */
	VIDEO_PIXEL_I420 VIDEO_PIXEL_FORMAT = C.VIDEO_PIXEL_I420
	/**
	 * 2: BGRA.
	 */
	VIDEO_PIXEL_BGRA VIDEO_PIXEL_FORMAT = C.VIDEO_PIXEL_BGRA
	/**
	 * 3: NV21.
	 */
	VIDEO_PIXEL_NV21 VIDEO_PIXEL_FORMAT = C.VIDEO_PIXEL_NV21
	/**
	 * 4: RGBA.
	 */
	VIDEO_PIXEL_RGBA VIDEO_PIXEL_FORMAT = C.VIDEO_PIXEL_RGBA
	/**
	 * 8: NV12.
	 */
	VIDEO_PIXEL_NV12 VIDEO_PIXEL_FORMAT = C.VIDEO_PIXEL_NV12
	/**
	 * 10: GL_TEXTURE_2D
	 */
	VIDEO_TEXTURE_2D VIDEO_PIXEL_FORMAT = C.VIDEO_TEXTURE_2D
	/**
	 * 11: GL_TEXTURE_OES
	 */
	VIDEO_TEXTURE_OES VIDEO_PIXEL_FORMAT = C.VIDEO_TEXTURE_OES
	/*
	   12: pixel format for iOS CVPixelBuffer NV12
	*/
	VIDEO_CVPIXEL_NV12 VIDEO_PIXEL_FORMAT = C.VIDEO_CVPIXEL_NV12
	/*
	   13: pixel format for iOS CVPixelBuffer I420
	*/
	VIDEO_CVPIXEL_I420 VIDEO_PIXEL_FORMAT = C.VIDEO_CVPIXEL_I420
	/*
	   14: pixel format for iOS CVPixelBuffer BGRA
	*/
	VIDEO_CVPIXEL_BGRA VIDEO_PIXEL_FORMAT = C.VIDEO_CVPIXEL_BGRA
	/**
	 * 16: I422.
	 */
	VIDEO_PIXEL_I422 VIDEO_PIXEL_FORMAT = C.VIDEO_PIXEL_I422
	/**
	 * 17: ID3D11Texture2D, only support DXGI_FORMAT_B8G8R8A8_UNORM, DXGI_FORMAT_B8G8R8A8_TYPELESS, DXGI_FORMAT_NV12 texture format
	 */
	VIDEO_TEXTURE_ID3D11TEXTURE2D VIDEO_PIXEL_FORMAT = C.VIDEO_TEXTURE_ID3D11TEXTURE2D
	/**
	 * 18: I010. 10bit I420 data.
	 */
	VIDEO_PIXEL_I010 VIDEO_PIXEL_FORMAT = C.VIDEO_PIXEL_I010
)

/**
 * The video display mode.
 */
type RENDER_MODE_TYPE C.enum_C_RENDER_MODE_TYPE

const (
	/**
	 * 1: Uniformly scale the video until it fills the visible boundaries
	 * (cropped). One dimension of the video may have clipped contents.
	 */
	RENDER_MODE_HIDDEN RENDER_MODE_TYPE = C.RENDER_MODE_HIDDEN
	/**
	 * 2: Uniformly scale the video until one of its dimension fits the boundary
	 * (zoomed to fit). Areas that are not filled due to the disparity in the
	 * aspect ratio will be filled with black.
	 */
	RENDER_MODE_FIT RENDER_MODE_TYPE = C.RENDER_MODE_FIT
	/**
	 * @deprecated
	 * 3: This mode is deprecated.
	 */
	RENDER_MODE_ADAPTIVE RENDER_MODE_TYPE = C.RENDER_MODE_ADAPTIVE
)

/**
 * The camera video source type
 */
type CAMERA_VIDEO_SOURCE_TYPE C.enum_C_CAMERA_VIDEO_SOURCE_TYPE

const (
	/**
	 * 0: the video frame comes from the front camera
	 */
	CAMERA_SOURCE_FRONT CAMERA_VIDEO_SOURCE_TYPE = C.CAMERA_SOURCE_FRONT
	/**
	 * 1: the video frame comes from the back camera
	 */
	CAMERA_SOURCE_BACK CAMERA_VIDEO_SOURCE_TYPE = C.CAMERA_SOURCE_BACK
	/**
	 * 1: the video frame source is unsepcified
	 */
	VIDEO_SOURCE_UNSPECIFIED CAMERA_VIDEO_SOURCE_TYPE = C.VIDEO_SOURCE_UNSPECIFIED
)

// /**
//   - The IVideoFrameMetaInfo class.
//   - This interface provides access to metadata information.
//     */
//
// typedef void C_IVideoFrameMetaInfo;
// // class IVideoFrameMetaInfo {
// //   public:
type META_INFO_KEY C.enum_C_META_INFO_KEY

const (
	KEY_FACE_CAPTURE META_INFO_KEY = C.KEY_FACE_CAPTURE
)

//   //     virtual ~IVideoFrameMetaInfo() {};
//   //     virtual const char* getMetaInfoStr(META_INFO_KEY key) const = 0;
//   // };

//   // /**
//   //  * The definition of the ExternalVideoFrame struct.
//   //  */
//   // struct ExternalVideoFrame {
//   //   ExternalVideoFrame()
//   //       : type(VIDEO_BUFFER_RAW_DATA),
//   //         format(VIDEO_PIXEL_DEFAULT),
//   //         buffer(NULL),
//   //         stride(0),
//   //         height(0),
//   //         cropLeft(0),
//   //         cropTop(0),
//   //         cropRight(0),
//   //         cropBottom(0),
//   //         rotation(0),
//   //         timestamp(0),
//   //         eglContext(NULL),
//   //         eglType(EGL_CONTEXT10),
//   //         textureId(0),
//   //         metadata_buffer(NULL),
//   //         metadata_size(0),
//   //         alphaBuffer(NULL),
//   //         fillAlphaBuffer(false),
//   //         d3d11_texture_2d(NULL),
//   //         texture_slice_index(0){}

/**
 * The EGL context type.
 */
type EGL_CONTEXT_TYPE C.enum_C_EGL_CONTEXT_TYPE

const (
	/**
	 * 0: When using the OpenGL interface (javax.microedition.khronos.egl.*) defined by Khronos
	 */
	EGL_CONTEXT10 EGL_CONTEXT_TYPE = C.EGL_CONTEXT10
	/**
	 * 0: When using the OpenGL interface (android.opengl.*) defined by Android
	 */
	EGL_CONTEXT14 EGL_CONTEXT_TYPE = C.EGL_CONTEXT14
)

/**
 * Video buffer types.
 */
type VIDEO_BUFFER_TYPE C.enum_C_VIDEO_BUFFER_TYPE

const (
	/**
	 * 1: Raw data.
	 */
	VIDEO_BUFFER_RAW_DATA VIDEO_BUFFER_TYPE = C.VIDEO_BUFFER_RAW_DATA
	/**
	 * 2: The same as VIDEO_BUFFER_RAW_DATA.
	 */
	VIDEO_BUFFER_ARRAY VIDEO_BUFFER_TYPE = C.VIDEO_BUFFER_ARRAY
	/**
	 * 3: The video buffer in the format of texture.
	 */
	VIDEO_BUFFER_TEXTURE VIDEO_BUFFER_TYPE = C.VIDEO_BUFFER_TEXTURE
)

//   //   /**
//   //    * The buffer type: #VIDEO_BUFFER_TYPE.
//   //    */
//   //   VIDEO_BUFFER_TYPE type;
//   //   /**
//   //    * The pixel format: #VIDEO_PIXEL_FORMAT
//   //    */
//   //   VIDEO_PIXEL_FORMAT format;
//   //   /**
//   //    * The video buffer.
//   //    */
//   //   void* buffer;
//   //   /**
//   //    * The line spacing of the incoming video frame (px). For
//   //    * texture, it is the width of the texture.
//   //    */
//   //   int stride;
//   //   /**
//   //    * The height of the incoming video frame.
//   //    */
//   //   int height;
//   //   /**
//   //    * [Raw data related parameter] The number of pixels trimmed from the left. The default value is
//   //    * 0.
//   //    */
//   //   int cropLeft;
//   //   /**
//   //    * [Raw data related parameter] The number of pixels trimmed from the top. The default value is
//   //    * 0.
//   //    */
//   //   int cropTop;
//   //   /**
//   //    * [Raw data related parameter] The number of pixels trimmed from the right. The default value is
//   //    * 0.
//   //    */
//   //   int cropRight;
//   //   /**
//   //    * [Raw data related parameter] The number of pixels trimmed from the bottom. The default value
//   //    * is 0.
//   //    */
//   //   int cropBottom;
//   //   /**
//   //    * [Raw data related parameter] The clockwise rotation information of the video frame. You can set the
//   //    * rotation angle as 0, 90, 180, or 270. The default value is 0.
//   //    */
//   //   int rotation;
//   //   /**
//   //    * The timestamp (ms) of the incoming video frame. An incorrect timestamp results in a frame loss or
//   //    * unsynchronized audio and video.
//   //    *
//   //    * Please refer to getAgoraCurrentMonotonicTimeInMs or getCurrentMonotonicTimeInMs
//   //    * to determine how to fill this filed.
//   //    */
//   //   long long timestamp;
//   //   /**
//   //    * [Texture-related parameter]
//   //    * When using the OpenGL interface (javax.microedition.khronos.egl.*) defined by Khronos, set EGLContext to this field.
//   //    * When using the OpenGL interface (android.opengl.*) defined by Android, set EGLContext to this field.
//   //    */
//   //   void *eglContext;
//   //   /**
//   //    * [Texture related parameter] Texture ID used by the video frame.
//   //    */
//   //   EGL_CONTEXT_TYPE eglType;
//   //   /**
//   //    * [Texture related parameter] Incoming 4 &times; 4 transformational matrix. The typical value is a unit matrix.
//   //    */
//   //   int textureId;
//   //   /**
//   //    * [Texture related parameter] Incoming 4 &times; 4 transformational matrix. The typical value is a unit matrix.
//   //    */
//   //   float matrix[16];
//   //   /**
//   //    * [Texture related parameter] The MetaData buffer.
//   //    *  The default value is NULL
//   //    */
//   //   uint8_t* metadata_buffer;
//   //   /**
//   //    * [Texture related parameter] The MetaData size.
//   //    *  The default value is 0
//   //    */
//   //   int metadata_size;
//   //   /**
//   //    *  Indicates the alpha channel of current frame, which is consistent with the dimension of the video frame.
//   //    *  The value range of each pixel is [0,255], where 0 represents the background; 255 represents the foreground.
//   //    *  The default value is NULL.
//   //    *  @technical preview
//   //    */
//   //   uint8_t* alphaBuffer;
//   //   /**
//   //    *  Extract alphaBuffer from bgra or rgba data. Set it true if you do not explicitly specify the alphabuffer.
//   //    *  The default value is false
//   //    *  @technical preview
//   //    */
//   //   bool fillAlphaBuffer;

//   //   /**
//   //    * [Windows Texture related parameter] The pointer of ID3D11Texture2D used by the video frame.
//   //    */
//   //   void *d3d11_texture_2d;

//   //   /**
//   //    * [Windows Texture related parameter] The index of ID3D11Texture2D array used by the video frame.
//   //    */
//   //   int texture_slice_index;
//   // };

//   // /**
//   //  * The definition of the VideoFrame struct.
//   //  */
//   // struct VideoFrame {
//   //   VideoFrame():
//   //   type(VIDEO_PIXEL_DEFAULT),
//   //   width(0),
//   //   height(0),
//   //   yStride(0),
//   //   uStride(0),
//   //   vStride(0),
//   //   yBuffer(NULL),
//   //   uBuffer(NULL),
//   //   vBuffer(NULL),
//   //   rotation(0),
//   //   renderTimeMs(0),
//   //   avsync_type(0),
//   //   metadata_buffer(NULL),
//   //   metadata_size(0),
//   //   sharedContext(0),
//   //   textureId(0),
//   //   d3d11Texture2d(NULL),
//   //   alphaBuffer(NULL),
//   //   pixelBuffer(NULL),
//   //   metaInfo(NULL){
//   //     memset(matrix, 0, sizeof(matrix));
//   //   }
//   //   /**
//   //    * The video pixel format: #VIDEO_PIXEL_FORMAT.
//   //    */
//   //   VIDEO_PIXEL_FORMAT type;
//   //   /**
//   //    * The width of the video frame.
//   //    */
//   //   int width;
//   //   /**
//   //    * The height of the video frame.
//   //    */
//   //   int height;
//   //   /**
//   //    * The line span of Y buffer in the YUV data.
//   //    */
//   //   int yStride;
//   //   /**
//   //    * The line span of U buffer in the YUV data.
//   //    */
//   //   int uStride;
//   //   /**
//   //    * The line span of V buffer in the YUV data.
//   //    */
//   //   int vStride;
//   //   /**
//   //    * The pointer to the Y buffer in the YUV data.
//   //    */
//   //   uint8_t* yBuffer;
//   //   /**
//   //    * The pointer to the U buffer in the YUV data.
//   //    */
//   //   uint8_t* uBuffer;
//   //   /**
//   //    * The pointer to the V buffer in the YUV data.
//   //    */
//   //   uint8_t* vBuffer;
//   //   /**
//   //    * The clockwise rotation information of this frame. You can set it as 0, 90, 180 or 270.
//   //    */
//   //   int rotation;
//   //   /**
//   //    * The timestamp to render the video stream. Use this parameter for audio-video synchronization when
//   //    * rendering the video.
//   //    *
//   //    * @note This parameter is for rendering the video, not capturing the video.
//   //    */
//   //   int64_t renderTimeMs;
//   //   /**
//   //    * The type of audio-video synchronization.
//   //    */
//   //   int avsync_type;
//   //   /**
//   //    * [Texture related parameter] The MetaData buffer.
//   //    *  The default value is NULL
//   //    */
//   //   uint8_t* metadata_buffer;
//   //   /**
//   //    * [Texture related parameter] The MetaData size.
//   //    *  The default value is 0
//   //    */
//   //   int metadata_size;
//   //   /**
//   //    * [Texture related parameter], egl context.
//   //    */
//   //   void* sharedContext;
//   //   /**
//   //    * [Texture related parameter], Texture ID used by the video frame.
//   //    */
//   //   int textureId;
//   //   /**
//   //    * [Texture related parameter] The pointer of ID3D11Texture2D used by the video frame,for Windows only.
//   //    */
//   //   void* d3d11Texture2d;
//   //   /**
//   //    * [Texture related parameter], Incoming 4 &times; 4 transformational matrix.
//   //    */
//   //   float matrix[16];
//   //   /**
//   //    *  Indicates the alpha channel of current frame, which is consistent with the dimension of the video frame.
//   //    *  The value range of each pixel is [0,255], where 0 represents the background; 255 represents the foreground.
//   //    *  The default value is NULL.
//   //    *  @technical preview
//   //    */
//   //   uint8_t* alphaBuffer;
//   //   /**
//   //    *The type of CVPixelBufferRef, for iOS and macOS only.
//   //    */
//   //   void* pixelBuffer;
//   //   /**
//   //    *  The pointer to IVideoFrameMetaInfo, which is the interface to get metainfo contents from VideoFrame.
//   //    */
//   //   IVideoFrameMetaInfo* metaInfo;
//   // };

//   /**
//    * The IVideoFrameObserver class.
//    */
//   typedef void C_IVideoFrameObserver;
//   // class IVideoFrameObserver {
//   //  public:
//   //   /**
//   //    * Occurs each time the player receives a video frame.
//   //    *
//   //    * After registering the video frame observer,
//   //    * the callback occurs each time the player receives a video frame to report the detailed information of the video frame.
//   //    * @param frame The detailed information of the video frame. See {@link VideoFrame}.
//   //    */
//   //   virtual void onFrame(const VideoFrame* frame) = 0;
//   //   virtual ~IVideoFrameObserver() {}
//   //   virtual bool isExternal() { return true; }
//   //   virtual VIDEO_PIXEL_FORMAT getVideoFormatPreference() { return VIDEO_PIXEL_DEFAULT; }
//   // };

type MEDIA_PLAYER_SOURCE_TYPE C.enum_C_MEDIA_PLAYER_SOURCE_TYPE

const (
	/**
	 * The real type of media player when use MEDIA_PLAYER_SOURCE_DEFAULT is decided by the
	 * type of SDK package. It is full feature media player in full-featured SDK, or simple
	 * media player in others.
	 */
	MEDIA_PLAYER_SOURCE_DEFAULT MEDIA_PLAYER_SOURCE_TYPE = C.MEDIA_PLAYER_SOURCE_DEFAULT
	/**
	 * Full featured media player is designed to support more codecs and media format, which
	 * requires more package size than simple player. If you need this player enabled, you
	 * might need to download a full-featured SDK.
	 */
	MEDIA_PLAYER_SOURCE_FULL_FEATURED MEDIA_PLAYER_SOURCE_TYPE = C.MEDIA_PLAYER_SOURCE_FULL_FEATURED
	/**
	 * Simple media player with limit codec supported, which requires minimal package size
	 * requirement and is enabled by default
	 */
	MEDIA_PLAYER_SOURCE_SIMPLE MEDIA_PLAYER_SOURCE_TYPE = C.MEDIA_PLAYER_SOURCE_SIMPLE
)

type VIDEO_MODULE_POSITION C.enum_C_VIDEO_MODULE_POSITION

const (
	POSITION_POST_CAPTURER VIDEO_MODULE_POSITION = C.POSITION_POST_CAPTURER
	POSITION_PRE_RENDERER  VIDEO_MODULE_POSITION = C.POSITION_PRE_RENDERER
	POSITION_PRE_ENCODER   VIDEO_MODULE_POSITION = C.POSITION_PRE_ENCODER
)

// #endregion agora::media::base

//   // /**
//   //  * The audio frame observer.
//   //  */
//   // class IAudioPcmFrameSink {
//   //  public:
//   //   /**
//   //    * Occurs when each time the player receives an audio frame.
//   //    *
//   //    * After registering the audio frame observer,
//   //    * the callback occurs when each time the player receives an audio frame,
//   //    * reporting the detailed information of the audio frame.
//   //    * @param frame The detailed information of the audio frame. See {@link AudioPcmFrame}.
//   //    */
//   //   virtual void onFrame(agora::media::base::AudioPcmFrame* frame) = 0;
//   //   virtual ~IAudioPcmFrameSink() {}
//   // };

/**
 * The IAudioFrameObserverBase class.
 */
type IAudioFrameObserverBase C.C_IAudioFrameObserverBase

// #region IAudioFrameObserverBase

/**
 * Audio frame types.
 */
type AUDIO_FRAME_TYPE C.enum_C_AUDIO_FRAME_TYPE

const (
	/**
	 * 0: 16-bit PCM.
	 */
	FRAME_TYPE_PCM16 AUDIO_FRAME_TYPE = C.FRAME_TYPE_PCM16
)

//   // enum { MAX_HANDLE_TIME_CNT = 10 };

/**
 * The definition of the AudioFrame struct.
 */
type AudioFrame C.struct_C_AudioFrame

// #region AudioFrame

/**
 * The audio frame type: #AUDIO_FRAME_TYPE.
 */
func (this_ *AudioFrame) GetType() AUDIO_FRAME_TYPE {
	return AUDIO_FRAME_TYPE(this_._type)
}

/**
 * The audio frame type: #AUDIO_FRAME_TYPE.
 */
func (this_ *AudioFrame) SetType(_type AUDIO_FRAME_TYPE) {
	this_._type = C.enum_C_AUDIO_FRAME_TYPE(_type)
}

/**
 * The number of samples per channel in this frame.
 */
func (this_ *AudioFrame) GetSamplesPerChannel() int {
	return int(this_.samplesPerChannel)
}

/**
 * The number of samples per channel in this frame.
 */
func (this_ *AudioFrame) SetSamplesPerChannel(samplesPerChannel int) {
	this_.samplesPerChannel = C.int(samplesPerChannel)
}

/**
 * The number of bytes per sample: #BYTES_PER_SAMPLE
 */
func (this_ *AudioFrame) GetBytesPerSample() BYTES_PER_SAMPLE {
	return BYTES_PER_SAMPLE(this_.bytesPerSample)
}

/**
 * The number of bytes per sample: #BYTES_PER_SAMPLE
 */
func (this_ *AudioFrame) SetBytesPerSample(bytesPerSample BYTES_PER_SAMPLE) {
	this_.bytesPerSample = C.enum_C_BYTES_PER_SAMPLE(bytesPerSample)
}

/**
 * The number of audio channels (data is interleaved, if stereo).
 * - 1: Mono.
 * - 2: Stereo.
 */
func (this_ *AudioFrame) GetChannels() int {
	return int(this_.channels)
}

/**
 * The number of audio channels (data is interleaved, if stereo).
 * - 1: Mono.
 * - 2: Stereo.
 */
func (this_ *AudioFrame) SetChannels(channels int) {
	this_.channels = C.int(channels)
}

/**
 * The sample rate
 */
func (this_ *AudioFrame) GetSamplesPerSec() int {
	return int(this_.samplesPerSec)
}

/**
 * The sample rate
 */
func (this_ *AudioFrame) SetSamplesPerSec(samplesPerSec int) {
	this_.samplesPerSec = C.int(samplesPerSec)
}

/**
 * The data buffer of the audio frame. When the audio frame uses a stereo channel, the data
 * buffer is interleaved.
 *
 * Buffer data size: buffer = samplesPerChannel × channels × bytesPerSample.
 */
func (this_ *AudioFrame) GetBuffer() []byte {
	size := this_.samplesPerChannel * this_.channels * C.int(this_.bytesPerSample)
	return C.GoBytes(this_.buffer, size)
}

/**
 * The data buffer of the audio frame. When the audio frame uses a stereo channel, the data
 * buffer is interleaved.
 *
 * Buffer data size: buffer = samplesPerChannel × channels × bytesPerSample.
 */
func (this_ *AudioFrame) SetBuffer(buffer []byte) {
	this_.buffer = C.CBytes(buffer)
}

/**
 * The timestamp to render the audio data.
 *
 * You can use this timestamp to restore the order of the captured audio frame, and synchronize
 * audio and video frames in video scenarios, including scenarios where external video sources
 * are used.
 */
func (this_ *AudioFrame) GetRenderTimeMs() int64 {
	return int64(this_.renderTimeMs)
}

/**
 * The timestamp to render the audio data.
 *
 * You can use this timestamp to restore the order of the captured audio frame, and synchronize
 * audio and video frames in video scenarios, including scenarios where external video sources
 * are used.
 */
func (this_ *AudioFrame) SetRenderTimeMs(renderTimeMs int64) {
	this_.renderTimeMs = C.int64_t(renderTimeMs)
}

/**
 * A reserved parameter.
 */
func (this_ *AudioFrame) GetAvsyncType() int {
	return int(this_.avsync_type)
}

/**
 * A reserved parameter.
 */
func (this_ *AudioFrame) SetAvsyncType(avsync_type int) {
	this_.avsync_type = C.int(avsync_type)
}

/**
 * A reserved parameter.
 *
 * You can use this presentationMs parameter to indicate the presenation milisecond timestamp,
 * this will then filled into audio4 extension part, the remote side could use this pts in av
 * sync process with video frame.
 */
func (this_ *AudioFrame) GetPresentationMs() int64 {
	return int64(this_.presentationMs)
}

/**
 * A reserved parameter.
 *
 * You can use this presentationMs parameter to indicate the presenation milisecond timestamp,
 * this will then filled into audio4 extension part, the remote side could use this pts in av
 * sync process with video frame.
 */
func (this_ *AudioFrame) SetPresentationMs(presentationMs int64) {
	this_.presentationMs = C.int64_t(presentationMs)
}

/**
 * The number of the audio track.
 */
func (this_ *AudioFrame) GetAudioTrackNumber() int {
	return int(this_.audioTrackNumber)
}

/**
 * The number of the audio track.
 */
func (this_ *AudioFrame) SetAudioTrackNumber(audioTrackNumber int) {
	this_.audioTrackNumber = C.int(audioTrackNumber)
}

/**
 * RTP timestamp of the first sample in the AudioFrame
 */
func (this_ *AudioFrame) GetRtpTimestamp() uint32 {
	return uint32(this_.rtpTimestamp)
}

/**
 * RTP timestamp of the first sample in the AudioFrame
 */
func (this_ *AudioFrame) SetRtpTimestamp(rtpTimestamp uint32) {
	this_.rtpTimestamp = C.uint32_t(rtpTimestamp)
}

func NewAudioFrame() *AudioFrame {
	return (*AudioFrame)(C.C_AudioFrame_New())
}
func (this_ *AudioFrame) Delete() {
	C.C_AudioFrame_Delete((*C.struct_C_AudioFrame)(this_))
}

// #endregion AudioFrame

type AUDIO_FRAME_POSITION C.enum_C_AUDIO_FRAME_POSITION

const (
	AUDIO_FRAME_POSITION_NONE AUDIO_FRAME_POSITION = C.AUDIO_FRAME_POSITION_NONE
	/** The position for observing the playback audio of all remote users after mixing
	 */
	AUDIO_FRAME_POSITION_PLAYBACK AUDIO_FRAME_POSITION = C.AUDIO_FRAME_POSITION_PLAYBACK
	/** The position for observing the recorded audio of the local user
	 */
	AUDIO_FRAME_POSITION_RECORD AUDIO_FRAME_POSITION = C.AUDIO_FRAME_POSITION_RECORD
	/** The position for observing the mixed audio of the local user and all remote users
	 */
	AUDIO_FRAME_POSITION_MIXED AUDIO_FRAME_POSITION = C.AUDIO_FRAME_POSITION_MIXED
	/** The position for observing the audio of a single remote user before mixing
	 */
	AUDIO_FRAME_POSITION_BEFORE_MIXING AUDIO_FRAME_POSITION = C.AUDIO_FRAME_POSITION_BEFORE_MIXING
	/** The position for observing the ear monitoring audio of the local user
	 */
	AUDIO_FRAME_POSITION_EAR_MONITORING AUDIO_FRAME_POSITION = C.AUDIO_FRAME_POSITION_EAR_MONITORING
)

type AudioParams C.struct_C_AudioParams

//   struct C_AudioParams
//   {
//     /** The audio sample rate (Hz), which can be set as one of the following values:

//      - `8000`
//      - `16000` (Default)
//      - `32000`
//      - `44100 `
//      - `48000`
//      */
//     int sample_rate;

//     /* The number of audio channels, which can be set as either of the following values:

//      - `1`: Mono (Default)
//      - `2`: Stereo
//      */
//     int channels;

//     /* The use mode of the audio data. See AgoraAudioRawFrameOperationMode.
//      */
//     enum C_RAW_AUDIO_FRAME_OP_MODE_TYPE mode;

//     /** The number of samples. For example, set it as 1024 for RTMP or RTMPS
//      streaming.
//      */
//     int samples_per_call;
//   };

func NewAudioParams() *AudioParams {
	return (*AudioParams)(C.C_AudioParams_New())
}
func (this_ *AudioParams) Delete() {
	C.C_AudioParams_Delete((*C.struct_C_AudioParams)(this_))
}

//   void C_IAudioFrameObserverBase_Delete(C_IAudioFrameObserverBase *this_);

//   /**
//    * Occurs when the recorded audio frame is received.
//    * @param channelId The channel name
//    * @param audioFrame The reference to the audio frame: AudioFrame.
//    * @return
//    * - true: The recorded audio frame is valid and is encoded and sent.
//    * - false: The recorded audio frame is invalid and is not encoded or sent.
//    */
//   bool C_IAudioFrameObserverBase_onRecordAudioFrame(C_IAudioFrameObserverBase *this_, const char *channelId, struct C_AudioFrame *audioFrame);
//   /**
//    * Occurs when the playback audio frame is received.
//    * @param channelId The channel name
//    * @param audioFrame The reference to the audio frame: AudioFrame.
//    * @return
//    * - true: The playback audio frame is valid and is encoded and sent.
//    * - false: The playback audio frame is invalid and is not encoded or sent.
//    */
//   bool C_IAudioFrameObserverBase_onPlaybackAudioFrame(C_IAudioFrameObserverBase *this_, const char *channelId, struct C_AudioFrame *audioFrame);
//   /**
//    * Occurs when the mixed audio data is received.
//    * @param channelId The channel name
//    * @param audioFrame The reference to the audio frame: AudioFrame.
//    * @return
//    * - true: The mixed audio data is valid and is encoded and sent.
//    * - false: The mixed audio data is invalid and is not encoded or sent.
//    */
//   bool C_IAudioFrameObserverBase_onMixedAudioFrame(C_IAudioFrameObserverBase *this_, const char *channelId, struct C_AudioFrame *audioFrame);
//   /**
//    * Occurs when the ear monitoring audio frame is received.
//    * @param audioFrame The reference to the audio frame: AudioFrame.
//    * @return
//    * - true: The ear monitoring audio data is valid and is encoded and sent.
//    * - false: The ear monitoring audio data is invalid and is not encoded or sent.
//    */
//   bool C_IAudioFrameObserverBase_onEarMonitoringAudioFrame(C_IAudioFrameObserverBase *this_, struct C_AudioFrame *audioFrame);
//   /**
//    * Occurs when the before-mixing playback audio frame is received.
//    * @param channelId The channel name
//    * @param userId ID of the remote user.
//    * @param audioFrame The reference to the audio frame: AudioFrame.
//    * @return
//    * - true: The before-mixing playback audio frame is valid and is encoded and sent.
//    * - false: The before-mixing playback audio frame is invalid and is not encoded or sent.
//    */
//   bool C_IAudioFrameObserverBase_onPlaybackAudioFrameBeforeMixing(C_IAudioFrameObserverBase *this_, const char *channelId, C_user_id_t userId, struct C_AudioFrame *audioFrame);

//   /**
//    * Sets the frame position for the audio observer.
//    * @return A bit mask that controls the frame position of the audio observer.
//    * @note - Use '|' (the OR operator) to observe multiple frame positions.
//    * <p>
//    * After you successfully register the audio observer, the SDK triggers this callback each time it receives a audio frame. You can determine which position to observe by setting the return value.
//    * The SDK provides 4 positions for observer. Each position corresponds to a callback function:
//    * - `AUDIO_FRAME_POSITION_PLAYBACK (1 << 0)`: The position for playback audio frame is received, which corresponds to the \ref onPlaybackFrame "onPlaybackFrame" callback.
//    * - `AUDIO_FRAME_POSITION_RECORD (1 << 1)`: The position for record audio frame is received, which corresponds to the \ref onRecordFrame "onRecordFrame" callback.
//    * - `AUDIO_FRAME_POSITION_MIXED (1 << 2)`: The position for mixed audio frame is received, which corresponds to the \ref onMixedFrame "onMixedFrame" callback.
//    * - `AUDIO_FRAME_POSITION_BEFORE_MIXING (1 << 3)`: The position for playback audio frame before mixing is received, which corresponds to the \ref onPlaybackFrameBeforeMixing "onPlaybackFrameBeforeMixing" callback.
//    *  @return The bit mask that controls the audio observation positions.
//    * See AUDIO_FRAME_POSITION.
//    */
//   int C_IAudioFrameObserverBase_getObservedAudioFramePosition(C_IAudioFrameObserverBase *this_);

//   /** Sets the audio playback format
//    **Note**:

//    - The SDK calculates the sample interval according to the `AudioParams`
//    you set in the return value of this callback and triggers the
//    `onPlaybackAudioFrame` callback at the calculated sample interval.
//    Sample interval (seconds) = `samplesPerCall`/(`sampleRate` × `channel`).
//    Ensure that the value of sample interval is equal to or greater than 0.01.

//    @return Sets the audio format. See AgoraAudioParams.
//    */
//   struct C_AudioParams C_IAudioFrameObserverBase_getPlaybackAudioParams(C_IAudioFrameObserverBase *this_);

//   /** Sets the audio recording format
//    **Note**:
//    - The SDK calculates the sample interval according to the `AudioParams`
//    you set in the return value of this callback and triggers the
//    `onRecordAudioFrame` callback at the calculated sample interval.
//    Sample interval (seconds) = `samplesPerCall`/(`sampleRate` × `channel`).
//    Ensure that the value of sample interval is equal to or greater than 0.01.

//    @return Sets the audio format. See AgoraAudioParams.
//    */
//   struct C_AudioParams C_IAudioFrameObserverBase_getRecordAudioParams(C_IAudioFrameObserverBase *this_);

//   /** Sets the audio mixing format
//    **Note**:
//    - The SDK calculates the sample interval according to the `AudioParams`
//    you set in the return value of this callback and triggers the
//    `onMixedAudioFrame` callback at the calculated sample interval.
//    Sample interval (seconds) = `samplesPerCall`/(`sampleRate` × `channel`).
//    Ensure that the value of sample interval is equal to or greater than 0.01.

//    @return Sets the audio format. See AgoraAudioParams.
//    */
//   struct C_AudioParams C_IAudioFrameObserverBase_getMixedAudioParams(C_IAudioFrameObserverBase *this_);

//   /** Sets the ear monitoring audio format
//    **Note**:
//    - The SDK calculates the sample interval according to the `AudioParams`
//    you set in the return value of this callback and triggers the
//    `onEarMonitoringAudioFrame` callback at the calculated sample interval.
//    Sample interval (seconds) = `samplesPerCall`/(`sampleRate` × `channel`).
//    Ensure that the value of sample interval is equal to or greater than 0.01.

//    @return Sets the audio format. See AgoraAudioParams.
//    */
//   struct C_AudioParams C_IAudioFrameObserverBase_getEarMonitoringAudioParams(C_IAudioFrameObserverBase *this_);

// #endregion IAudioFrameObserverBase

/**
 * The IAudioFrameObserver class.
 */
type IAudioFrameObserver C.C_IAudioFrameObserver

// #region IAudioFrameObserver
//   // using IAudioFrameObserverBase::onPlaybackAudioFrameBeforeMixing;
//   /**
//    * Occurs when the before-mixing playback audio frame is received.
//    * @param channelId The channel name
//    * @param uid ID of the remote user.
//    * @param audioFrame The reference to the audio frame: AudioFrame.
//    * @return
//    * - true: The before-mixing playback audio frame is valid and is encoded and sent.
//    * - false: The before-mixing playback audio frame is invalid and is not encoded or sent.
//    */
//   bool C_IAudioFrameObserver_onPlaybackAudioFrameBeforeMixing(C_IAudioFrameObserver *this_, const char *channelId, C_uid_t uid, struct C_AudioFrame *audioFrame);
// #endregion IAudioFrameObserver

//   // struct AudioSpectrumData {
//   //   /**
//   //    * The audio spectrum data of audio.
//   //    */
//   //   const float *audioSpectrumData;
//   //   /**
//   //    * The data length of audio spectrum data.
//   //    */
//   //   int dataLength;

//   //   AudioSpectrumData() : audioSpectrumData(NULL), dataLength(0) {}
//   //   AudioSpectrumData(const float *data, int length) :
//   //     audioSpectrumData(data), dataLength(length) {}
//   // };

//   // struct UserAudioSpectrumInfo  {
//   //   /**
//   //    * User ID of the speaker.
//   //    */
//   //   agora::rtc::uid_t uid;
//   //   /**
//   //    * The audio spectrum data of audio.
//   //    */
//   //   struct AudioSpectrumData spectrumData;

//   //   UserAudioSpectrumInfo () : uid(0), spectrumData() {}
//   //   UserAudioSpectrumInfo(agora::rtc::uid_t _uid, const float *data, int length) :
//   //     uid(_uid) { spectrumData.audioSpectrumData = data; spectrumData.dataLength = length; }
//   // };

//   /**
//    * The IAudioSpectrumObserver class.
//    */
//   typedef void C_IAudioSpectrumObserver;
//   // class IAudioSpectrumObserver {
//   // public:
//   //   virtual ~IAudioSpectrumObserver() {}

//   //   /**
//   //    * Reports the audio spectrum of local audio.
//   //    *
//   //    * This callback reports the audio spectrum data of the local audio at the moment
//   //    * in the channel.
//   //    *
//   //    * You can set the time interval of this callback using \ref ILocalUser::enableAudioSpectrumMonitor "enableAudioSpectrumMonitor".
//   //    *
//   //    * @param data The audio spectrum data of local audio.
//   //    * - true: Processed.
//   //    * - false: Not processed.
//   //    */
//   //   virtual bool onLocalAudioSpectrum(const AudioSpectrumData& data) = 0;

//   //   /**
//   //    * Reports the audio spectrum of remote user.
//   //    *
//   //    * This callback reports the IDs and audio spectrum data of the loudest speakers at the moment
//   //    * in the channel.
//   //    *
//   //    * You can set the time interval of this callback using \ref ILocalUser::enableAudioSpectrumMonitor "enableAudioSpectrumMonitor".
//   //    *
//   //    * @param spectrums The pointer to \ref agora::media::UserAudioSpectrumInfo "UserAudioSpectrumInfo", which is an array containing
//   //    * the user ID and audio spectrum data for each speaker.
//   //    * - This array contains the following members:
//   //    *   - `uid`, which is the UID of each remote speaker
//   //    *   - `spectrumData`, which reports the audio spectrum of each remote speaker.
//   //    * @param spectrumNumber The array length of the spectrums.
//   //    * - true: Processed.
//   //    * - false: Not processed.
//   //    */
//   //   virtual bool onRemoteAudioSpectrum(const UserAudioSpectrumInfo * spectrums, unsigned int spectrumNumber) = 0;
//   // };

//   /**
//    * The IVideoEncodedFrameObserver class.
//    */
//   typedef void C_IVideoEncodedFrameObserver;
//   // class IVideoEncodedFrameObserver {
//   //  public:
//   //   /**
//   //    * Occurs each time the SDK receives an encoded video image.
//   //    * @param uid The user id of remote user.
//   //    * @param imageBuffer The pointer to the video image buffer.
//   //    * @param length The data length of the video image.
//   //    * @param videoEncodedFrameInfo The information of the encoded video frame: EncodedVideoFrameInfo.
//   //    * @return Determines whether to accept encoded video image.
//   //    * - true: Accept.
//   //    * - false: Do not accept.
//   //    */
//   //   virtual bool onEncodedVideoFrameReceived(rtc::uid_t uid, const uint8_t* imageBuffer, size_t length,
//   //                                            const rtc::EncodedVideoFrameInfo& videoEncodedFrameInfo) = 0;

//   //   virtual ~IVideoEncodedFrameObserver() {}
//   // };

// /**
//   - The IVideoFrameObserver class.
//     */
//
// typedef void C_IVideoFrameObserver;
// // class IVideoFrameObserver {
// //  public:
// //   typedef media::base::VideoFrame VideoFrame;
// //   /**
// //    * The process mode of the video frame:
// //    */
type VIDEO_FRAME_PROCESS_MODE C.enum_C_VIDEO_FRAME_PROCESS_MODE

const (
	/**
	 * Read-only mode.
	 *
	 * In this mode, you do not modify the video frame. The video frame observer is a renderer.
	 */
	PROCESS_MODE_READ_ONLY VIDEO_FRAME_PROCESS_MODE = C.PROCESS_MODE_READ_ONLY // Observer works as a pure renderer and will not modify the original frame.
	/**
	 * Read and write mode.
	 *
	 * In this mode, you modify the video frame. The video frame observer is a video filter.
	 */
	PROCESS_MODE_READ_WRITE VIDEO_FRAME_PROCESS_MODE = C.PROCESS_MODE_READ_WRITE // Observer works as a filter that will process the video frame and affect the following frame processing in SDK.
)

//   //  public:
//   //   virtual ~IVideoFrameObserver() {}

//   //   /**
//   //    * Occurs each time the SDK receives a video frame captured by the local camera.
//   //    *
//   //    * After you successfully register the video frame observer, the SDK triggers this callback each time
//   //    * a video frame is received. In this callback, you can get the video data captured by the local
//   //    * camera. You can then pre-process the data according to your scenarios.
//   //    *
//   //    * After pre-processing, you can send the processed video data back to the SDK by setting the
//   //    * `videoFrame` parameter in this callback.
//   //    *
//   //    * @note
//   //    * - If you get the video data in RGBA color encoding format, Agora does not support using this callback to send the processed data in RGBA color encoding format back to the SDK.
//   //    * - The video data that this callback gets has not been pre-processed, such as watermarking, cropping content, rotating, or image enhancement.
//   //    *
//   //    * @param videoFrame A pointer to the video frame: VideoFrame
//   //    * @param sourceType source type of video frame. See #VIDEO_SOURCE_TYPE.
//   //    * @return Determines whether to ignore the current video frame if the pre-processing fails:
//   //    * - true: Do not ignore.
//   //    * - false: Ignore, in which case this method does not sent the current video frame to the SDK.
//   //   */
//   //   virtual bool onCaptureVideoFrame(agora::rtc::VIDEO_SOURCE_TYPE sourceType, VideoFrame& videoFrame) = 0;

//   //   /**
//   //    * Occurs each time the SDK receives a video frame before encoding.
//   //    *
//   //    * After you successfully register the video frame observer, the SDK triggers this callback each time
//   //    * when it receives a video frame. In this callback, you can get the video data before encoding. You can then
//   //    * process the data according to your particular scenarios.
//   //    *
//   //    * After processing, you can send the processed video data back to the SDK by setting the
//   //    * `videoFrame` parameter in this callback.
//   //    *
//   //    * @note
//   //    * - To get the video data captured from the second screen before encoding, you need to set (1 << 2) as a frame position through `getObservedFramePosition`.
//   //    * - The video data that this callback gets has been pre-processed, such as watermarking, cropping content, rotating, or image enhancement.
//   //    * - This callback does not support sending processed RGBA video data back to the SDK.
//   //    *
//   //    * @param videoFrame A pointer to the video frame: VideoFrame
//   //    * @param sourceType source type of video frame. See #VIDEO_SOURCE_TYPE.
//   //    * @return Determines whether to ignore the current video frame if the pre-processing fails:
//   //    * - true: Do not ignore.
//   //    * - false: Ignore, in which case this method does not sent the current video frame to the SDK.
//   //    */
//   //   virtual bool onPreEncodeVideoFrame(agora::rtc::VIDEO_SOURCE_TYPE sourceType, VideoFrame& videoFrame) = 0;

//   //   /**
//   //    * Occurs each time the SDK receives a video frame decoded by the MediaPlayer.
//   //    *
//   //    * After you successfully register the video frame observer, the SDK triggers this callback each
//   //    * time a video frame is decoded. In this callback, you can get the video data decoded by the
//   //    * MediaPlayer. You can then pre-process the data according to your scenarios.
//   //    *
//   //    * After pre-processing, you can send the processed video data back to the SDK by setting the
//   //    * `videoFrame` parameter in this callback.
//   //    *
//   //    * @note
//   //    * - This callback will not be affected by the return values of \ref getVideoFrameProcessMode "getVideoFrameProcessMode", \ref getRotationApplied "getRotationApplied", \ref getMirrorApplied "getMirrorApplied", \ref getObservedFramePosition "getObservedFramePosition".
//   //    * - On Android, this callback is not affected by the return value of \ref getVideoFormatPreference "getVideoFormatPreference"
//   //    *
//   //    * @param videoFrame A pointer to the video frame: VideoFrame
//   //    * @param mediaPlayerId ID of the mediaPlayer.
//   //    * @return Determines whether to ignore the current video frame if the pre-processing fails:
//   //    * - true: Do not ignore.
//   //    * - false: Ignore, in which case this method does not sent the current video frame to the SDK.
//   //    */
//   //   virtual bool onMediaPlayerVideoFrame(VideoFrame& videoFrame, int mediaPlayerId) = 0;

//   //   /**
//   //    * Occurs each time the SDK receives a video frame sent by the remote user.
//   //    *
//   //    * After you successfully register the video frame observer, the SDK triggers this callback each time a
//   //    * video frame is received. In this callback, you can get the video data sent by the remote user. You
//   //    * can then post-process the data according to your scenarios.
//   //    *
//   //    * After post-processing, you can send the processed data back to the SDK by setting the `videoFrame`
//   //    * parameter in this callback.
//   //    *
//   //    * @note This callback does not support sending processed RGBA video data back to the SDK.
//   //    *
//   //    * @param channelId The channel name
//   //    * @param remoteUid ID of the remote user who sends the current video frame.
//   //    * @param videoFrame A pointer to the video frame: VideoFrame
//   //    * @return Determines whether to ignore the current video frame if the post-processing fails:
//   //    * - true: Do not ignore.
//   //    * - false: Ignore, in which case this method does not sent the current video frame to the SDK.
//   //    */
//   //   virtual bool onRenderVideoFrame(const char* channelId, rtc::uid_t remoteUid, VideoFrame& videoFrame) = 0;

//   //   virtual bool onTranscodedVideoFrame(VideoFrame& videoFrame) = 0;

//   //   /**
//   //    * Occurs each time the SDK receives a video frame and prompts you to set the process mode of the video frame.
//   //    *
//   //    * After you successfully register the video frame observer, the SDK triggers this callback each time it receives
//   //    * a video frame. You need to set your preferred process mode in the return value of this callback.
//   //    * @return VIDEO_FRAME_PROCESS_MODE.
//   //    */
//   //   virtual VIDEO_FRAME_PROCESS_MODE getVideoFrameProcessMode() {
//   //     return PROCESS_MODE_READ_ONLY;
//   //   }

//   //   /**
//   //    * Sets the format of the raw video data output by the SDK.
//   //    *
//   //    * If you want to get raw video data in a color encoding format other than YUV 420, register this callback when
//   //    * calling `registerVideoFrameObserver`. After you successfully register the video frame observer, the SDK triggers
//   //    * this callback each time it receives a video frame. You need to set your preferred video data in the return value
//   //    * of this callback.
//   //    *
//   //    * @note If you want the video captured by the sender to be the original format, set the original video data format
//   //    * to VIDEO_PIXEL_DEFAULT in the return value. On different platforms, the original video pixel format is also
//   //    * different, for the actual video pixel format, see `VideoFrame`.
//   //    *
//   //    * @return Sets the video format. See VIDEO_PIXEL_FORMAT.
//   //    */
//   //   virtual base::VIDEO_PIXEL_FORMAT getVideoFormatPreference() { return base::VIDEO_PIXEL_DEFAULT; }

//   //   /**
//   //    * Occurs each time the SDK receives a video frame, and prompts you whether to rotate the captured video.
//   //    *
//   //    * If you want to rotate the captured video according to the rotation member in the `VideoFrame` class, register this
//   //    * callback by calling `registerVideoFrameObserver`. After you successfully register the video frame observer, the
//   //    * SDK triggers this callback each time it receives a video frame. You need to set whether to rotate the video frame
//   //    * in the return value of this callback.
//   //    *
//   //    * @note This function only supports video data in RGBA or YUV420.
//   //    *
//   //    * @return Determines whether to rotate.
//   //    * - `true`: Rotate the captured video.
//   //    * - `false`: (Default) Do not rotate the captured video.
//   //    */
//   //   virtual bool getRotationApplied() { return false; }

//   //   /**
//   //    * Occurs each time the SDK receives a video frame and prompts you whether or not to mirror the captured video.
//   //    *
//   //    * If the video data you want to obtain is a mirror image of the original video, you need to register this callback
//   //    * when calling `registerVideoFrameObserver`. After you successfully register the video frame observer, the SDK
//   //    * triggers this callback each time it receives a video frame. You need to set whether or not to mirror the video
//   //    * frame in the return value of this callback.
//   //    *
//   //    * @note This function only supports video data in RGBA and YUV420 formats.
//   //    *
//   //    * @return Determines whether to mirror.
//   //    * - `true`: Mirror the captured video.
//   //    * - `false`: (Default) Do not mirror the captured video.
//   //    */
//   //   virtual bool getMirrorApplied() { return false; }

//   //   /**
//   //    * Sets the frame position for the video observer.
//   //    *
//   //    * After you successfully register the video observer, the SDK triggers this callback each time it receives
//   //    * a video frame. You can determine which position to observe by setting the return value. The SDK provides
//   //    * 3 positions for observer. Each position corresponds to a callback function:
//   //    *
//   //    * POSITION_POST_CAPTURER(1 << 0): The position after capturing the video data, which corresponds to the onCaptureVideoFrame callback.
//   //    * POSITION_PRE_RENDERER(1 << 1): The position before receiving the remote video data, which corresponds to the onRenderVideoFrame callback.
//   //    * POSITION_PRE_ENCODER(1 << 2): The position before encoding the video data, which corresponds to the onPreEncodeVideoFrame callback.
//   //    *
//   //    * To observe multiple frame positions, use '|' (the OR operator).
//   //    * This callback observes POSITION_POST_CAPTURER(1 << 0) and POSITION_PRE_RENDERER(1 << 1) by default.
//   //    * To conserve the system consumption, you can reduce the number of frame positions that you want to observe.
//   //    *
//   //    * @return A bit mask that controls the frame position of the video observer: VIDEO_OBSERVER_POSITION.
//   //    */
//   //   virtual uint32_t getObservedFramePosition() {
//   //     return base::POSITION_POST_CAPTURER | base::POSITION_PRE_RENDERER;
//   //   }

//   //   /**
//   //    * Indicate if the observer is for internal use.
//   //    * Note: Never override this function
//   //    * @return
//   //    * - true: the observer is for external use
//   //    * - false: the observer is for internal use
//   //    */
//   //   virtual bool isExternal() { return true; }
//   // };

/**
 * The external video source type.
 */
type EXTERNAL_VIDEO_SOURCE_TYPE C.enum_C_EXTERNAL_VIDEO_SOURCE_TYPE

const (
	/**
	 * 0: non-encoded video frame.
	 */
	VIDEO_FRAME EXTERNAL_VIDEO_SOURCE_TYPE = C.VIDEO_FRAME
	/**
	 * 1: encoded video frame.
	 */
	ENCODED_VIDEO_FRAME EXTERNAL_VIDEO_SOURCE_TYPE = C.ENCODED_VIDEO_FRAME
)

/**
 * The format of the recording file.
 *
 * @since v3.5.2
 */
type MediaRecorderContainerFormat C.enum_C_MediaRecorderContainerFormat

const (
	/**
	 * 1: (Default) MP4.
	 */
	FORMAT_MP4 MediaRecorderContainerFormat = C.FORMAT_MP4
)

/**
 * The recording content.
 *
 * @since v3.5.2
 */
type MediaRecorderStreamType C.enum_C_MediaRecorderStreamType

const (
	/**
	 * Only audio.
	 */
	STREAM_TYPE_AUDIO MediaRecorderStreamType = C.STREAM_TYPE_AUDIO
	/**
	 * Only video.
	 */
	STREAM_TYPE_VIDEO MediaRecorderStreamType = C.STREAM_TYPE_VIDEO
	/**
	 * (Default) Audio and video.
	 */
	STREAM_TYPE_BOTH MediaRecorderStreamType = C.STREAM_TYPE_BOTH
)

/**
 * The current recording state.
 *
 * @since v3.5.2
 */
type RecorderState C.enum_C_RecorderState

const (
	/**
	 * -1: An error occurs during the recording. See RecorderErrorCode for the reason.
	 */
	RECORDER_STATE_ERROR RecorderState = C.RECORDER_STATE_ERROR
	/**
	 * 2: The audio and video recording is started.
	 */
	RECORDER_STATE_START RecorderState = C.RECORDER_STATE_START
	/**
	 * 3: The audio and video recording is stopped.
	 */
	RECORDER_STATE_STOP RecorderState = C.RECORDER_STATE_STOP
)

/**
 * The reason for the state change
 *
 * @since v3.5.2
 */
type RecorderErrorCode C.enum_C_RecorderErrorCode

const (
	/**
	 * 0: No error occurs.
	 */
	RECORDER_ERROR_NONE RecorderErrorCode = C.RECORDER_ERROR_NONE
	/**
	 * 1: The SDK fails to write the recorded data to a file.
	 */
	RECORDER_ERROR_WRITE_FAILED RecorderErrorCode = C.RECORDER_ERROR_WRITE_FAILED
	/**
	 * 2: The SDK does not detect audio and video streams to be recorded, or audio and video streams are interrupted for more than five seconds during recording.
	 */
	RECORDER_ERROR_NO_STREAM RecorderErrorCode = C.RECORDER_ERROR_NO_STREAM
	/**
	 * 3: The recording duration exceeds the upper limit.
	 */
	RECORDER_ERROR_OVER_MAX_DURATION RecorderErrorCode = C.RECORDER_ERROR_OVER_MAX_DURATION
	/**
	 * 4: The recording configuration changes.
	 */
	RECORDER_ERROR_CONFIG_CHANGED RecorderErrorCode = C.RECORDER_ERROR_CONFIG_CHANGED
)

//   // /**
//   //  * Configurations for the local audio and video recording.
//   //  *
//   //  * @since v3.5.2
//   //  */
//   // struct MediaRecorderConfiguration {
//   //   /**
//   //    * The absolute path (including the filename extensions) of the recording file.
//   //    * For example, `C:\Users\<user_name>\AppData\Local\Agora\<process_name>\example.mp4` on Windows,
//   //    * `/App Sandbox/Library/Caches/example.mp4` on iOS, `/Library/Logs/example.mp4` on macOS, and
//   //    * `/storage/emulated/0/Android/data/<package name>/files/example.mp4` on Android.
//   //    *
//   //    * @note Ensure that the specified path exists and is writable.
//   //    */
//   //   const char* storagePath;
//   //   /**
//   //    * The format of the recording file. See \ref agora::rtc::MediaRecorderContainerFormat "MediaRecorderContainerFormat".
//   //    */
//   //   MediaRecorderContainerFormat containerFormat;
//   //   /**
//   //    * The recording content. See \ref agora::rtc::MediaRecorderStreamType "MediaRecorderStreamType".
//   //    */
//   //   MediaRecorderStreamType streamType;
//   //   /**
//   //    * The maximum recording duration, in milliseconds. The default value is 120000.
//   //    */
//   //   int maxDurationMs;
//   //   /**
//   //    * The interval (ms) of updating the recording information. The value range is
//   //    * [1000,10000]. Based on the set value of `recorderInfoUpdateInterval`, the
//   //    * SDK triggers the \ref IMediaRecorderObserver::onRecorderInfoUpdated "onRecorderInfoUpdated"
//   //    * callback to report the updated recording information.
//   //    */
//   //   int recorderInfoUpdateInterval;

//   //   MediaRecorderConfiguration() : storagePath(NULL), containerFormat(FORMAT_MP4), streamType(STREAM_TYPE_BOTH), maxDurationMs(120000), recorderInfoUpdateInterval(0) {}
//   //   MediaRecorderConfiguration(const char* path, MediaRecorderContainerFormat format, MediaRecorderStreamType type, int duration, int interval) : storagePath(path), containerFormat(format), streamType(type), maxDurationMs(duration), recorderInfoUpdateInterval(interval) {}
//   // };

//   typedef void C_IFaceInfoObserver;
//   // class IFaceInfoObserver {
//   // public:
//   //    /**
//   //     * Occurs when the face info is received.
//   //     * @param outFaceInfo The output face info.
//   //     * @return
//   //     * - true: The face info is valid.
//   //     * - false: The face info is invalid.
//   //    */
//   //    virtual bool onFaceInfo(const char* outFaceInfo) = 0;

//   //    virtual ~IFaceInfoObserver() {}
//   // };

//   // /**
//   //  * Information for the recording file.
//   //  *
//   //  * @since v3.5.2
//   //  */
//   // struct RecorderInfo {
//   //   /**
//   //    * The absolute path of the recording file.
//   //    */
//   //   const char* fileName;
//   //   /**
//   //    * The recording duration, in milliseconds.
//   //    */
//   //   unsigned int durationMs;
//   //   /**
//   //    * The size in bytes of the recording file.
//   //    */
//   //   unsigned int fileSize;

//   //   RecorderInfo() : fileName(NULL), durationMs(0), fileSize(0) {}
//   //   RecorderInfo(const char* name, unsigned int dur, unsigned int size) : fileName(name), durationMs(dur), fileSize(size) {}
//   // };

//   typedef void C_IMediaRecorderObserver;
//   // class IMediaRecorderObserver {
//   //  public:
//   //   /**
//   //    * Occurs when the recording state changes.
//   //    *
//   //    * @since v4.0.0
//   //    *
//   //    * When the local audio and video recording state changes, the SDK triggers this callback to report the current
//   //    * recording state and the reason for the change.
//   //    *
//   //    * @param channelId The channel name.
//   //    * @param uid ID of the user.
//   //    * @param state The current recording state. See \ref agora::media::RecorderState "RecorderState".
//   //    * @param error The reason for the state change. See \ref agora::media::RecorderErrorCode "RecorderErrorCode".
//   //    */
//   //   virtual void onRecorderStateChanged(const char* channelId, rtc::uid_t uid, RecorderState state, RecorderErrorCode error) = 0;
//   //   /**
//   //    * Occurs when the recording information is updated.
//   //    *
//   //    * @since v4.0.0
//   //    *
//   //    * After you successfully register this callback and enable the local audio and video recording, the SDK periodically triggers
//   //    * the `onRecorderInfoUpdated` callback based on the set value of `recorderInfoUpdateInterval`. This callback reports the
//   //    * filename, duration, and size of the current recording file.
//   //    *
//   //    * @param channelId The channel name.
//   //    * @param uid ID of the user.
//   //    * @param info Information about the recording file. See \ref agora::media::RecorderInfo "RecorderInfo".
//   //    *
//   //    */
//   //   virtual void onRecorderInfoUpdated(const char* channelId, rtc::uid_t uid, const RecorderInfo& info) = 0;
//   //   virtual ~IMediaRecorderObserver() {}
//   // };

// #endregion agora::media

// #endregion agora

// #ifdef __cplusplus
// }
// #endif // __cplusplus

// #endif // C_AGORA_MEDIA_BASE_H
