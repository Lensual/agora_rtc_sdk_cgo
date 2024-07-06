package bridge

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "bridge/C_AudioFrameObserverBridge.h"

bool cgo_AudioFrameObserverBridge_onRecordAudioFrame(C_AudioFrameObserverBridge *this_, void *userData,
	char *channelId, struct C_AudioFrame *audioFrame);

bool cgo_AudioFrameObserverBridge_onPlaybackAudioFrame(C_AudioFrameObserverBridge *this_, void *userData,
	char *channelId, struct C_AudioFrame *audioFrame);

bool cgo_AudioFrameObserverBridge_onMixedAudioFrame(C_AudioFrameObserverBridge *this_, void *userData,
	char *channelId, struct C_AudioFrame *audioFrame);

bool cgo_AudioFrameObserverBridge_onEarMonitoringAudioFrame(C_AudioFrameObserverBridge *this_, void *userData,
	struct C_AudioFrame *audioFrame);

bool cgo_AudioFrameObserverBridge_onPlaybackAudioFrameBeforeMixing(C_AudioFrameObserverBridge *this_, void *userData,
	char *channelId, C_user_id_t userId, struct C_AudioFrame *audioFrame);

int cgo_AudioFrameObserverBridge_getObservedAudioFramePosition(C_AudioFrameObserverBridge *this_, void *userData);

struct C_AudioParams cgo_AudioFrameObserverBridge_getPlaybackAudioParams(C_AudioFrameObserverBridge *this_, void *userData);

struct C_AudioParams cgo_AudioFrameObserverBridge_getRecordAudioParams(C_AudioFrameObserverBridge *this_, void *userData);

struct C_AudioParams cgo_AudioFrameObserverBridge_getMixedAudioParams(C_AudioFrameObserverBridge *this_, void *userData);

struct C_AudioParams cgo_AudioFrameObserverBridge_getEarMonitoringAudioParams(C_AudioFrameObserverBridge *this_, void *userData);

*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/agora_rtc_sdk_cgo/pkg/agora"
)

type IAudioFrameObserverHandler interface {
	/**
	 * Occurs when the recorded audio frame is received.
	 * @param channelId The channel name
	 * @param audioFrame The reference to the audio frame: AudioFrame.
	 * @return
	 * - true: The recorded audio frame is valid and is encoded and sent.
	 * - false: The recorded audio frame is invalid and is not encoded or sent.
	 */
	OnRecordAudioFrame(channelId string, audioFrame *agora.AudioFrame) bool

	/**
	 * Occurs when the playback audio frame is received.
	 * @param channelId The channel name
	 * @param audioFrame The reference to the audio frame: AudioFrame.
	 * @return
	 * - true: The playback audio frame is valid and is encoded and sent.
	 * - false: The playback audio frame is invalid and is not encoded or sent.
	 */
	OnPlaybackAudioFrame(channelId string, audioFrame *agora.AudioFrame) bool

	/**
	 * Occurs when the mixed audio data is received.
	 * @param channelId The channel name
	 * @param audioFrame The reference to the audio frame: AudioFrame.
	 * @return
	 * - true: The mixed audio data is valid and is encoded and sent.
	 * - false: The mixed audio data is invalid and is not encoded or sent.
	 */
	OnMixedAudioFrame(channelId string, audioFrame *agora.AudioFrame) bool

	/**
	 * Occurs when the ear monitoring audio frame is received.
	 * @param audioFrame The reference to the audio frame: AudioFrame.
	 * @return
	 * - true: The ear monitoring audio data is valid and is encoded and sent.
	 * - false: The ear monitoring audio data is invalid and is not encoded or sent.
	 */
	OnEarMonitoringAudioFrame(audioFrame *agora.AudioFrame) bool

	/**
	 * Occurs when the before-mixing playback audio frame is received.
	 * @param channelId The channel name
	 * @param userId ID of the remote user.
	 * @param audioFrame The reference to the audio frame: AudioFrame.
	 * @return
	 * - true: The before-mixing playback audio frame is valid and is encoded and sent.
	 * - false: The before-mixing playback audio frame is invalid and is not encoded or sent.
	 */
	OnPlaybackAudioFrameBeforeMixing(channelId string, userId string, audioFrame *agora.AudioFrame) bool

	/**
	 * Sets the frame position for the audio observer.
	 * @return A bit mask that controls the frame position of the audio observer.
	 * @note - Use '|' (the OR operator) to observe multiple frame positions.
	 * <p>
	 * After you successfully register the audio observer, the SDK triggers this callback each time it receives a audio frame. You can determine which position to observe by setting the return value.
	 * The SDK provides 4 positions for observer. Each position corresponds to a callback function:
	 * - `AUDIO_FRAME_POSITION_PLAYBACK (1 << 0)`: The position for playback audio frame is received, which corresponds to the \ref onPlaybackFrame "onPlaybackFrame" callback.
	 * - `AUDIO_FRAME_POSITION_RECORD (1 << 1)`: The position for record audio frame is received, which corresponds to the \ref onRecordFrame "onRecordFrame" callback.
	 * - `AUDIO_FRAME_POSITION_MIXED (1 << 2)`: The position for mixed audio frame is received, which corresponds to the \ref onMixedFrame "onMixedFrame" callback.
	 * - `AUDIO_FRAME_POSITION_BEFORE_MIXING (1 << 3)`: The position for playback audio frame before mixing is received, which corresponds to the \ref onPlaybackFrameBeforeMixing "onPlaybackFrameBeforeMixing" callback.
	 *  @return The bit mask that controls the audio observation positions.
	 * See AUDIO_FRAME_POSITION.
	 */
	GetObservedAudioFramePosition() int

	/** Sets the audio playback format
	  **Note**:

	  - The SDK calculates the sample interval according to the `AudioParams`
	  you set in the return value of this callback and triggers the
	  `onPlaybackAudioFrame` callback at the calculated sample interval.
	  Sample interval (seconds) = `samplesPerCall`/(`sampleRate` × `channel`).
	  Ensure that the value of sample interval is equal to or greater than 0.01.

	  @return Sets the audio format. See AgoraAudioParams.
	*/
	GetPlaybackAudioParams() agora.AudioParams

	/** Sets the audio recording format
	  **Note**:
	  - The SDK calculates the sample interval according to the `AudioParams`
	  you set in the return value of this callback and triggers the
	  `onRecordAudioFrame` callback at the calculated sample interval.
	  Sample interval (seconds) = `samplesPerCall`/(`sampleRate` × `channel`).
	  Ensure that the value of sample interval is equal to or greater than 0.01.

	  @return Sets the audio format. See AgoraAudioParams.
	*/
	GetRecordAudioParams() agora.AudioParams

	/** Sets the audio mixing format
	  **Note**:
	  - The SDK calculates the sample interval according to the `AudioParams`
	  you set in the return value of this callback and triggers the
	  `onMixedAudioFrame` callback at the calculated sample interval.
	  Sample interval (seconds) = `samplesPerCall`/(`sampleRate` × `channel`).
	  Ensure that the value of sample interval is equal to or greater than 0.01.

	  @return Sets the audio format. See AgoraAudioParams.
	*/
	GetMixedAudioParams() agora.AudioParams

	/** Sets the ear monitoring audio format
	  **Note**:
	  - The SDK calculates the sample interval according to the `AudioParams`
	  you set in the return value of this callback and triggers the
	  `onEarMonitoringAudioFrame` callback at the calculated sample interval.
	  Sample interval (seconds) = `samplesPerCall`/(`sampleRate` × `channel`).
	  Ensure that the value of sample interval is equal to or greater than 0.01.

	  @return Sets the audio format. See AgoraAudioParams.
	*/
	GetEarMonitoringAudioParams() agora.AudioParams
}

type AudioFrameObserverBridge struct {
	handler IAudioFrameObserverHandler
	cBridge *C.C_AudioFrameObserverBridge
}

func (b *AudioFrameObserverBridge) ToAgoraEventHandler() *agora.IAudioFrameObserverBase {
	return (*agora.IAudioFrameObserverBase)(b.cBridge)
}

func (b *AudioFrameObserverBridge) Delete() {
	C.C_AudioFrameObserverBridge_Delete(unsafe.Pointer(b.cBridge))
	b.handler = nil
	b.cBridge = nil
}

func NewAudioFrameObserverBridge(handler IAudioFrameObserverHandler) *AudioFrameObserverBridge {
	b := AudioFrameObserverBridge{}
	userData := unsafe.Pointer(&b)
	b.cBridge = (*C.C_AudioFrameObserverBridge)(C.C_AudioFrameObserverBridge_New(
		C.C_AudioFrameObserverBridge_Callbacks{
			onRecordAudioFrame:               C.C_AudioFrameObserverBridge_onRecordAudioFrame(C.cgo_AudioFrameObserverBridge_onRecordAudioFrame),
			onPlaybackAudioFrame:             C.C_AudioFrameObserverBridge_onPlaybackAudioFrame(C.cgo_AudioFrameObserverBridge_onPlaybackAudioFrame),
			onMixedAudioFrame:                C.C_AudioFrameObserverBridge_onMixedAudioFrame(C.cgo_AudioFrameObserverBridge_onMixedAudioFrame),
			onEarMonitoringAudioFrame:        C.C_AudioFrameObserverBridge_onEarMonitoringAudioFrame(C.cgo_AudioFrameObserverBridge_onEarMonitoringAudioFrame),
			onPlaybackAudioFrameBeforeMixing: C.C_AudioFrameObserverBridge_onPlaybackAudioFrameBeforeMixing(C.cgo_AudioFrameObserverBridge_onPlaybackAudioFrameBeforeMixing),
			getObservedAudioFramePosition:    C.C_AudioFrameObserverBridge_getObservedAudioFramePosition(C.cgo_AudioFrameObserverBridge_getObservedAudioFramePosition),
			getPlaybackAudioParams:           C.C_AudioFrameObserverBridge_getPlaybackAudioParams(C.cgo_AudioFrameObserverBridge_getPlaybackAudioParams),
			getRecordAudioParams:             C.C_AudioFrameObserverBridge_getRecordAudioParams(C.cgo_AudioFrameObserverBridge_getRecordAudioParams),
			getMixedAudioParams:              C.C_AudioFrameObserverBridge_getMixedAudioParams(C.cgo_AudioFrameObserverBridge_getMixedAudioParams),
			getEarMonitoringAudioParams:      C.C_AudioFrameObserverBridge_getEarMonitoringAudioParams(C.cgo_AudioFrameObserverBridge_getEarMonitoringAudioParams),
		},
		userData,
	))
	b.handler = handler
	return &b
}

//export cgo_AudioFrameObserverBridge_onRecordAudioFrame
func cgo_AudioFrameObserverBridge_onRecordAudioFrame(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer,
	channelId *C.char, audioFrame *C.struct_C_AudioFrame) C.bool {
	if userData == nil {
		return true
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	return C.bool(bridge.handler.OnRecordAudioFrame(
		C.GoString(channelId),
		(*agora.AudioFrame)(unsafe.Pointer(audioFrame)),
	))
}

//export cgo_AudioFrameObserverBridge_onPlaybackAudioFrame
func cgo_AudioFrameObserverBridge_onPlaybackAudioFrame(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer,
	channelId *C.char, audioFrame *C.struct_C_AudioFrame) C.bool {
	if userData == nil {
		return true
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	return C.bool(bridge.handler.OnPlaybackAudioFrame(
		C.GoString(channelId),
		(*agora.AudioFrame)(unsafe.Pointer(audioFrame)),
	))
}

//export cgo_AudioFrameObserverBridge_onMixedAudioFrame
func cgo_AudioFrameObserverBridge_onMixedAudioFrame(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer,
	channelId *C.char, audioFrame *C.struct_C_AudioFrame) C.bool {
	if userData == nil {
		return true
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	return C.bool(bridge.handler.OnMixedAudioFrame(
		C.GoString(channelId),
		(*agora.AudioFrame)(unsafe.Pointer(audioFrame)),
	))
}

//export cgo_AudioFrameObserverBridge_onEarMonitoringAudioFrame
func cgo_AudioFrameObserverBridge_onEarMonitoringAudioFrame(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer,
	audioFrame *C.struct_C_AudioFrame) C.bool {
	if userData == nil {
		return true
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	return C.bool(bridge.handler.OnEarMonitoringAudioFrame(
		(*agora.AudioFrame)(unsafe.Pointer(audioFrame)),
	))
}

//export cgo_AudioFrameObserverBridge_onPlaybackAudioFrameBeforeMixing
func cgo_AudioFrameObserverBridge_onPlaybackAudioFrameBeforeMixing(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer,
	channelId *C.char, userId C.C_user_id_t, audioFrame *C.struct_C_AudioFrame) C.bool {
	if userData == nil {
		return true
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	return C.bool(bridge.handler.OnPlaybackAudioFrameBeforeMixing(
		C.GoString(channelId),
		C.GoString(userId),
		(*agora.AudioFrame)(unsafe.Pointer(audioFrame)),
	))
}

//export cgo_AudioFrameObserverBridge_getObservedAudioFramePosition
func cgo_AudioFrameObserverBridge_getObservedAudioFramePosition(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer) C.int {
	if userData == nil {
		return 0
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	return C.int(bridge.handler.GetObservedAudioFramePosition())
}

//export cgo_AudioFrameObserverBridge_getPlaybackAudioParams
func cgo_AudioFrameObserverBridge_getPlaybackAudioParams(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer) C.struct_C_AudioParams {
	if userData == nil {
		return C.struct_C_AudioParams{}
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	ret := bridge.handler.GetPlaybackAudioParams()
	return *(*C.struct_C_AudioParams)(unsafe.Pointer(&ret))
}

//export cgo_AudioFrameObserverBridge_getRecordAudioParams
func cgo_AudioFrameObserverBridge_getRecordAudioParams(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer) C.struct_C_AudioParams {
	if userData == nil {
		return C.struct_C_AudioParams{}
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	ret := bridge.handler.GetRecordAudioParams()
	return *(*C.struct_C_AudioParams)(unsafe.Pointer(&ret))
}

//export cgo_AudioFrameObserverBridge_getMixedAudioParams
func cgo_AudioFrameObserverBridge_getMixedAudioParams(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer) C.struct_C_AudioParams {
	if userData == nil {
		return C.struct_C_AudioParams{}
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	ret := bridge.handler.GetMixedAudioParams()
	return *(*C.struct_C_AudioParams)(unsafe.Pointer(&ret))
}

//export cgo_AudioFrameObserverBridge_getEarMonitoringAudioParams
func cgo_AudioFrameObserverBridge_getEarMonitoringAudioParams(_ *C.C_AudioFrameObserverBridge, userData unsafe.Pointer) C.struct_C_AudioParams {
	if userData == nil {
		return C.struct_C_AudioParams{}
	}

	bridge := (*AudioFrameObserverBridge)(userData)
	ret := bridge.handler.GetEarMonitoringAudioParams()
	return *(*C.struct_C_AudioParams)(unsafe.Pointer(&ret))
}
