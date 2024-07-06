package bridge

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "bridge/C_LocalUserObserverBridge.h"

void cgo_LocalUserObserverBridge_onAudioTrackPublishStart(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalAudioTrack *audioTrack);

void cgo_LocalUserObserverBridge_onAudioTrackPublishSuccess(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalAudioTrack *audioTrack);

void cgo_LocalUserObserverBridge_onAudioTrackUnpublished(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalAudioTrack *audioTrack);

void cgo_LocalUserObserverBridge_onAudioTrackPublicationFailure(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalAudioTrack *audioTrack,
	enum C_ERROR_CODE_TYPE error);

void cgo_LocalUserObserverBridge_onLocalAudioTrackStatistics(C_LocalUserObserverBridge *this_, void *userData,
	struct C_LocalAudioStats *stats);

void cgo_LocalUserObserverBridge_onRemoteAudioTrackStatistics(C_LocalUserObserverBridge *this_, void *userData,
	C_IRemoteAudioTrack *audioTrack,
	struct C_RemoteAudioTrackStats *stats);

void cgo_LocalUserObserverBridge_onUserAudioTrackSubscribed(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId,
	C_IRemoteAudioTrack *audioTrack);

void cgo_LocalUserObserverBridge_onUserAudioTrackStateChanged(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId,
	C_IRemoteAudioTrack *audioTrack,
	enum C_REMOTE_AUDIO_STATE state,
	enum C_REMOTE_AUDIO_STATE_REASON reason,
	int elapsed);

void cgo_LocalUserObserverBridge_onVideoTrackPublishStart(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalVideoTrack *videoTrack);

void cgo_LocalUserObserverBridge_onVideoTrackPublishSuccess(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalVideoTrack *videoTrack);

void cgo_LocalUserObserverBridge_onVideoTrackPublicationFailure(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalVideoTrack *videoTrack,
	enum C_ERROR_CODE_TYPE error);

void cgo_LocalUserObserverBridge_onVideoTrackUnpublished(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalVideoTrack *videoTrack);

void cgo_LocalUserObserverBridge_onLocalVideoTrackStateChanged(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalVideoTrack *videoTrack,
	enum C_LOCAL_VIDEO_STREAM_STATE state,
	enum C_LOCAL_VIDEO_STREAM_ERROR errorCode);

void cgo_LocalUserObserverBridge_onLocalVideoTrackStatistics(C_LocalUserObserverBridge *this_, void *userData,
	C_ILocalVideoTrack *videoTrack,
	struct C_LocalVideoTrackStats *stats);

void cgo_LocalUserObserverBridge_onUserVideoTrackSubscribed(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId, struct C_VideoTrackInfo trackInfo,
	C_IRemoteVideoTrack *videoTrack);

void cgo_LocalUserObserverBridge_onUserVideoTrackStateChanged(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId,
	C_IRemoteVideoTrack *videoTrack,
	enum C_REMOTE_VIDEO_STATE state,
	enum C_REMOTE_VIDEO_STATE_REASON reason,
	int elapsed);

void cgo_LocalUserObserverBridge_onFirstRemoteVideoFrameRendered(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId, int width, int height, int elapsed);

void cgo_LocalUserObserverBridge_onRemoteVideoTrackStatistics(C_LocalUserObserverBridge *this_, void *userData,
	C_IRemoteVideoTrack *videoTrack, struct C_RemoteVideoTrackStats *stats);

void cgo_LocalUserObserverBridge_onAudioVolumeIndication(C_LocalUserObserverBridge *this_, void *userData,
	struct C_AudioVolumeInformation *speakers, unsigned int speakerNumber,
	int totalVolume);

void cgo_LocalUserObserverBridge_onActiveSpeaker(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId);

void cgo_LocalUserObserverBridge_onAudioSubscribeStateChanged(C_LocalUserObserverBridge *this_, void *userData,
	char *channel, C_user_id_t userId, enum C_STREAM_SUBSCRIBE_STATE oldState, enum C_STREAM_SUBSCRIBE_STATE newState, int elapseSinceLastState);

void cgo_LocalUserObserverBridge_onVideoSubscribeStateChanged(C_LocalUserObserverBridge *this_, void *userData,
	char *channel, C_user_id_t userId, enum C_STREAM_SUBSCRIBE_STATE oldState, enum C_STREAM_SUBSCRIBE_STATE newState, int elapseSinceLastState);

void cgo_LocalUserObserverBridge_onAudioPublishStateChanged(C_LocalUserObserverBridge *this_, void *userData,
	char *channel, enum C_STREAM_PUBLISH_STATE oldState, enum C_STREAM_PUBLISH_STATE newState, int elapseSinceLastState);

void cgo_LocalUserObserverBridge_onVideoPublishStateChanged(C_LocalUserObserverBridge *this_, void *userData,
	char *channel, enum C_STREAM_PUBLISH_STATE oldState, enum C_STREAM_PUBLISH_STATE newState, int elapseSinceLastState);

void cgo_LocalUserObserverBridge_onRemoteSubscribeFallbackToAudioOnly(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId, bool isFallbackOrRecover);

void cgo_LocalUserObserverBridge_onFirstRemoteAudioFrame(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId, int elapsed);

void cgo_LocalUserObserverBridge_onFirstRemoteAudioDecoded(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId, int elapsed);

void cgo_LocalUserObserverBridge_onFirstRemoteVideoFrame(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId, int width, int height, int elapsed);

void cgo_LocalUserObserverBridge_onFirstRemoteVideoDecoded(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId, int width, int height, int elapsed);

void cgo_LocalUserObserverBridge_onVideoSizeChanged(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId, int width, int height, int rotation);

void cgo_LocalUserObserverBridge_onUserInfoUpdated(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t userId, enum C_USER_MEDIA_INFO msg, bool val);

void cgo_LocalUserObserverBridge_onIntraRequestReceived(C_LocalUserObserverBridge *this_, void *userData);

void cgo_LocalUserObserverBridge_onStreamMessage(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t user_id, int streamId, char *data, size_t length);

void cgo_LocalUserObserverBridge_onUserStateChanged(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t user_id, uint32_t state);

void cgo_LocalUserObserverBridge_onVideoRenderingTracingResult(C_LocalUserObserverBridge *this_, void *userData,
	C_user_id_t user_id, enum C_MEDIA_TRACE_EVENT currentState, struct C_VideoRenderingTracingInfo tracingInfo);

*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/agora_rtc_sdk_cgo/pkg/agora"
)

type ILocalUserObserverHandler interface {
	OnAudioTrackPublishStart(audioTrack *agora.ILocalAudioTrack)

	/**
	 * Occurs when the first packet of the local audio track is sent, indicating that the local audio track
	 * is successfully published.
	 *
	 * @param audioTrack The pointer to \ref agora::rtc::ILocalAudioTrack "ILocalAudioTrack".
	 */
	OnAudioTrackPublishSuccess(audioTrack *agora.ILocalAudioTrack)

	OnAudioTrackUnpublished(audioTrack *agora.ILocalAudioTrack)

	/**
	 * Occurs when a local audio track fails to be published.
	 *
	 * @param audioTrack The pointer to `ILocalAudioTrack`.
	 * @param error The error information: #ERROR_CODE_TYPE.
	 */
	OnAudioTrackPublicationFailure(audioTrack *agora.ILocalAudioTrack,
		_error agora.ERROR_CODE_TYPE)

	/**
	 * Reports the statistics of a local audio track.
	 *
	 * @param stats The reference to the statistics of the local audio track.
	 */
	OnLocalAudioTrackStatistics(stats *agora.LocalAudioStats)

	/**
	 * Reports the statistics of a remote audio track.
	 *
	 * @param audioTrack The pointer to `IRemoteAudioTrack`.
	 * @param stats The statistics of the remote audio track.
	 */
	OnRemoteAudioTrackStatistics(audioTrack *agora.IRemoteAudioTrack,
		stats *agora.RemoteAudioTrackStats)

	/**
	 * Occurs when the first remote audio frame is received.
	 *
	 * This callback indicates that the local user has subscribed to a specified remote audio track, and the first
	 * frame of this audio track has been received.
	 *
	 * @param userId The ID of the remote user whose audio frame is received.
	 * @param audioTrack The pointer to `IRemoteAudioTrack`.
	 */
	OnUserAudioTrackSubscribed(userId string,
		audioTrack *agora.IRemoteAudioTrack)

	/**
	 * Occurs when the state of a remote audio track changes.
	 *
	 * @param userId The ID of the remote user whose audio track state has changed.
	 * @param audioTrack The pointer to `IRemoteAudioTrack`.
	 * @param state The current state of the audio track.
	 * @param reason The reason for the state change.
	 * @param elapsed The time (ms) since the user connects to an Agora channel.
	 */
	OnUserAudioTrackStateChanged(userId string,
		audioTrack *agora.IRemoteAudioTrack,
		state agora.REMOTE_AUDIO_STATE,
		reason agora.REMOTE_AUDIO_STATE_REASON,
		elapsed int)

	OnVideoTrackPublishStart(videoTrack *agora.ILocalVideoTrack)

	/**
	 * Occurs when the first packet of a local video track is sent, indicating that the local video track
	 * is successfully published.
	 * @param videoTrack The pointer to `ILocalVideoTrack`.
	 */
	OnVideoTrackPublishSuccess(videoTrack *agora.ILocalVideoTrack)

	/**
	 * Occurs when a local video track fails to be published.
	 *
	 * @param videoTrack The pointer to `ILocalVideoTrack`.
	 * @param error The error information: #ERROR_CODE_TYPE.
	 */
	OnVideoTrackPublicationFailure(videoTrack *agora.ILocalVideoTrack,
		_error agora.ERROR_CODE_TYPE)

	OnVideoTrackUnpublished(videoTrack *agora.ILocalVideoTrack)

	/**
	 * Occurs when the state of a local video track changes.
	 * @note
	 * When you receive error from this callback, the corresponding track is in error state.
	 * To make the track recover from error state, we highly recommend that you disable the track and
	 * try re-enabling it again.
	 *
	 * @param videoTrack The pointer to `ILocalVideoTrack`.
	 * @param state The state of the local video track.
	 * @param errorCode The error information.
	 */
	OnLocalVideoTrackStateChanged(videoTrack *agora.ILocalVideoTrack,
		state agora.LOCAL_VIDEO_STREAM_STATE,
		errorCode agora.LOCAL_VIDEO_STREAM_ERROR)

	/**
	 * Reports the statistics of a local video track.
	 *
	 * @param videoTrack The pointer to `ILocalVideoTrack`.
	 * @param stats The statistics of the local video track.
	 */
	OnLocalVideoTrackStatistics(videoTrack *agora.ILocalVideoTrack,
		stats *agora.LocalVideoTrackStats)

	/**
	 * Occurs when the first remote video frame is received.
	 *
	 * This callback indicates that the local user has subscribed to a specified remote video track, and
	 * the first frame of this video track has been received.
	 *
	 * @param userId The ID of the remote user whose video frame is received.
	 * @param trackInfo The information of the remote video track.
	 * @param videoTrack The pointer to `IRemoteVideoTrack`.
	 */
	OnUserVideoTrackSubscribed(userId string, trackInfo agora.VideoTrackInfo,
		videoTrack *agora.IRemoteVideoTrack)

	/**
	 * Occurs when the state of a remote video track changes.
	 *
	 * @param userId the ID of the remote user whose video track state has changed.
	 * @param videoTrack The pointer to `IRemoteVideoTrack`.
	 * @param state The current state of the video track.
	 * @param reason The reason for the state change.
	 * @param elapsed The time (ms) since the user connects to an Agora channel.
	 */
	OnUserVideoTrackStateChanged(userId string,
		videoTrack *agora.IRemoteVideoTrack,
		state agora.REMOTE_VIDEO_STATE,
		reason agora.REMOTE_VIDEO_STATE_REASON,
		elapsed int)

	/**
	 * Occurs when the first remote video frame is rendered.
	 *
	 * @param userId the ID of the remote user.
	 * @param width Width (px) of the video frame.
	 * @param height Height (px) of the video stream.
	 * @param elapsed The time (ms) since the user connects to an Agora channel.
	 */
	OnFirstRemoteVideoFrameRendered(userId string, width int, height int, elapsed int)

	/**
	 * Reports the statistics of a remote video track.
	 *
	 * @param videoTrack The pointer to `IRemoteVideoTrack`.
	 * @param stats The statistics of the remote video track.
	 */
	OnRemoteVideoTrackStatistics(videoTrack *agora.IRemoteVideoTrack, stats *agora.RemoteVideoTrackStats)

	/**
	 * Reports which users are speaking, the speakers' volumes, and whether the local user is speaking.
	 *
	 * This callback reports the IDs and volumes of the loudest speakers at the moment in the channel,
	 * and whether the local user is speaking.
	 *
	 * You can set the time interval of this callback using \ref ILocalUser::setAudioVolumeIndicationParameters "setAudioVolumeIndicationParameters".
	 *
	 * The SDK triggers two `onAudioVolumeIndication` callbacks at one time, one reporting the
	 * volume information of the local user and the other reporting the volume information of all remote users.
	 *
	 * @note
	 * - To detect whether the local user is speaking, set `report_vad` in `enableAudioVolumeIndication` to `true`.
	 *
	 * @param speakers The pointer to \ref agora::rtc::AudioVolumeInformation "AudioVolumeInformation", which is an array containing the user ID and volume information for each speaker.
	 * - In the local user's callback, this array contains the following members:
	 *   - `uid`, which is always `0`
	 *   - `volume`, which reports the sum of the voice volume and the audio-mixing volume of the local user
	 *   - `vad`, which reports whether the local user is speaking
	 * - In the remote users' callback, this array contains the following members:
	 *   - `uid`, which is the UID of each remote speaker
	 *   - `volume`, which reports the sum of the voice volume and the audio-mixing volume of each remote speaker.
	 *   - `vad`, which is always `0`
	 * An empty `speakers` array in the callback indicates that no remote user is speaking at the moment.
	 * @param speakerNumber The total number of the speakers.
	 * @param totalVolume The total volume after audio mixing. The value ranges between 0 (lowest volume) and 255 (highest volume).
	 * - In the local user's callback, `totalVolume` is the sum of the voice volume and the audio-mixing volume of the local user.
	 * - In the remote speakers' callback, `totalVolume` is the sum of the voice volume and the audio-mixing volume of all remote speakers.
	 */
	OnAudioVolumeIndication(speakers *agora.AudioVolumeInformation, speakerNumber uint,
		totalVolume int)

	/**
	 * Occurs when an active speaker is detected.
	 *
	 * You can add relative functions on your app, for example, the active speaker, once detected,
	 * will have the portrait zoomed in.
	 *
	 * @note
	 * - The active speaker means the user ID of the speaker who speaks at the highest volume during a
	 * certain period of time.
	 *
	 * @param userId The ID of the active speaker. A `userId` of `0` means the local user.
	 */
	OnActiveSpeaker(userId string)

	/**
	 * Occurs when the audio subscribe state changed.
	 *
	 * @param channel The channel name of user joined.
	 * @param userId The remote string user ID that is subscribed to.
	 * @param oldState The old state of the audio stream subscribe : #STREAM_SUBSCRIBE_STATE.
	 * @param newState The new state of the audio stream subscribe : #STREAM_SUBSCRIBE_STATE.
	 * @param elapseSinceLastState The time elapsed (ms) from the old state to the new state.
	 */
	OnAudioSubscribeStateChanged(channel string, userId string, oldState agora.STREAM_SUBSCRIBE_STATE, newState agora.STREAM_SUBSCRIBE_STATE, elapseSinceLastState int)

	/**
	 * Occurs when the video subscribe state changed.
	 *
	 * @param channel The channel name of user joined.
	 * @param userId The remote string user ID that is subscribed to.
	 * @param oldState The old state of the video stream subscribe : #STREAM_SUBSCRIBE_STATE.
	 * @param newState The new state of the video stream subscribe : #STREAM_SUBSCRIBE_STATE.
	 * @param elapseSinceLastState The time elapsed (ms) from the old state to the new state.
	 */
	OnVideoSubscribeStateChanged(channel string, userId string, oldState agora.STREAM_SUBSCRIBE_STATE, newState agora.STREAM_SUBSCRIBE_STATE, elapseSinceLastState int)

	/**
	 * Occurs when the audio publish state changed.
	 *
	 * @param channel The channel name of user joined.
	 * @param oldState The old state of the audio stream publish : #STREAM_PUBLISH_STATE.
	 * @param newState The new state of the audio stream publish : #STREAM_PUBLISH_STATE.
	 * @param elapseSinceLastState The time elapsed (ms) from the old state to the new state.
	 */
	OnAudioPublishStateChanged(channel string, oldState agora.STREAM_PUBLISH_STATE, newState agora.STREAM_PUBLISH_STATE, elapseSinceLastState int)

	/**
	 * Occurs when the video publish state changed.
	 *
	 * @param channel The channel name of user joined.
	 * @param oldState The old state of the video stream publish : #STREAM_PUBLISH_STATE.
	 * @param newState The new state of the video stream publish : #STREAM_PUBLISH_STATE.
	 * @param elapseSinceLastState The time elapsed (ms) from the old state to the new state.
	 */
	OnVideoPublishStateChanged(channel string, oldState agora.STREAM_PUBLISH_STATE, newState agora.STREAM_PUBLISH_STATE, elapseSinceLastState int)

	/** Occurs when the first remote audio frame is received.
	 *
	 * @param userId ID of the remote user.
	 * @param isFallbackOrRecover Whether the remotely subscribed media stream
	 * falls back to audio-only or switches back to the video:
	 * - true: The remotely subscribed media stream falls back to audio-only
	 * due to poor network conditions.
	 * - false: The remotely subscribed media stream switches back to the
	 * video stream after the network conditions improved.
	 **/
	OnRemoteSubscribeFallbackToAudioOnly(userId string, isFallbackOrRecover bool)

	/** Occurs when the first remote audio frame is received.
	 *
	 * @param userId ID of the remote user.
	 * @param elapsed The time (ms) since the user connects to an Agora channel.
	 **/
	OnFirstRemoteAudioFrame(userId string, elapsed int)

	/**
	 * Occurs when the SDK decodes the first remote audio frame for playback.
	 *
	 * @param uid User ID of the remote user sending the audio stream.
	 * @param elapsed The time (ms) since the user connects to an Agora channel.
	 */
	OnFirstRemoteAudioDecoded(userId string, elapsed int)

	/**
	 * Occurs when the first remote video frame is rendered.
	 * The SDK triggers this callback when the first frame of the remote video is displayed in the user's video window. The application can get the time elapsed from a user joining the channel until the first video frame is displayed.
	 *
	 * @param userId ID of the remote user.
	 * @param width Width (px) of the video frame.
	 * @param height Height (px) of the video stream.
	 * @param elapsed Time elapsed (ms) from the local user calling the \ref IRtcEngine::joinChannel "joinChannel" method until the SDK triggers this callback.
	 */
	OnFirstRemoteVideoFrame(userId string, width int, height int, elapsed int)

	/**
	 * Occurs when the SDK decodes the first remote video frame for playback.
	 *
	 * @param userId ID of the remote user.
	 * @param width Width (px) of the video stream.
	 * @param height Height (px) of the video stream.
	 * @param elapsed The time (ms) since the user connects to an Agora channel.
	 */
	OnFirstRemoteVideoDecoded(userId string, width int, height int, elapsed int)

	/**
	 * The local or remote video size or rotation changed.
	 *
	 * @param uid User ID of the user whose video size or rotation has changed.
	 * @param width Width (pixels) of the video stream.
	 * @param height Height (pixels) of the video stream.
	 * @param rotation Rotation [0 to 360).
	 */
	OnVideoSizeChanged(userId string, width int, height int, rotation int)

	/**
	 * Occurs when the user media information is updated.
	 *
	 *
	 * @param userId The ID of the user.
	 * @param msg The media information of the user. See #USER_MEDIA_INFO.
	 * @param val Whether the user is muted.
	 */
	OnUserInfoUpdated(userId string, msg agora.USER_MEDIA_INFO, val bool)

	/**
	 * Occurs when the intra request is received from a remote user.
	 *
	 * The method notifies the local user to encode a key frame.
	 *
	 */
	OnIntraRequestReceived()

	/**
	 * datastream from this connection.
	 */
	OnStreamMessage(user_id string, streamId int, data []byte, length uint)

	/**
	 * Occurs when the remote user state is updated.
	 * @param uid The uid of the remote user.
	 * @param state The remote user state.Just & #REMOTE_USER_STATE
	 */
	OnUserStateChanged(user_id string, state uint32)

	OnVideoRenderingTracingResult(user_id string, currentState agora.MEDIA_TRACE_EVENT, tracingInfo agora.VideoRenderingTracingInfo)
}

type LocalUserObserverBridge struct {
	handler ILocalUserObserverHandler
	cBridge *C.C_LocalUserObserverBridge
}

func (b *LocalUserObserverBridge) ToAgoraEventHandler() *agora.ILocalUserObserver {
	return (*agora.ILocalUserObserver)(b.cBridge)
}

func (b *LocalUserObserverBridge) Delete() {
	C.C_LocalUserObserverBridge_Delete(unsafe.Pointer(b.cBridge))
	b.handler = nil
	b.cBridge = nil
}

func NewLocalUserObserverBridge(handler ILocalUserObserverHandler) *LocalUserObserverBridge {
	b := LocalUserObserverBridge{}
	userData := unsafe.Pointer(&b)
	b.cBridge = (*C.C_LocalUserObserverBridge)(C.C_LocalUserObserverBridge_New(
		C.C_LocalUserObserverBridge_Callbacks{
			onAudioTrackPublishStart:             C.C_LocalUserObserverBridge_onAudioTrackPublishStart(C.cgo_LocalUserObserverBridge_onAudioTrackPublishStart),
			onAudioTrackPublishSuccess:           C.C_LocalUserObserverBridge_onAudioTrackPublishSuccess(C.cgo_LocalUserObserverBridge_onAudioTrackPublishSuccess),
			onAudioTrackUnpublished:              C.C_LocalUserObserverBridge_onAudioTrackUnpublished(C.cgo_LocalUserObserverBridge_onAudioTrackUnpublished),
			onAudioTrackPublicationFailure:       C.C_LocalUserObserverBridge_onAudioTrackPublicationFailure(C.cgo_LocalUserObserverBridge_onAudioTrackPublicationFailure),
			onLocalAudioTrackStatistics:          C.C_LocalUserObserverBridge_onLocalAudioTrackStatistics(C.cgo_LocalUserObserverBridge_onLocalAudioTrackStatistics),
			onRemoteAudioTrackStatistics:         C.C_LocalUserObserverBridge_onRemoteAudioTrackStatistics(C.cgo_LocalUserObserverBridge_onRemoteAudioTrackStatistics),
			onUserAudioTrackSubscribed:           C.C_LocalUserObserverBridge_onUserAudioTrackSubscribed(C.cgo_LocalUserObserverBridge_onUserAudioTrackSubscribed),
			onUserAudioTrackStateChanged:         C.C_LocalUserObserverBridge_onUserAudioTrackStateChanged(C.cgo_LocalUserObserverBridge_onUserAudioTrackStateChanged),
			onVideoTrackPublishStart:             C.C_LocalUserObserverBridge_onVideoTrackPublishStart(C.cgo_LocalUserObserverBridge_onVideoTrackPublishStart),
			onVideoTrackPublishSuccess:           C.C_LocalUserObserverBridge_onVideoTrackPublishSuccess(C.cgo_LocalUserObserverBridge_onVideoTrackPublishSuccess),
			onVideoTrackPublicationFailure:       C.C_LocalUserObserverBridge_onVideoTrackPublicationFailure(C.cgo_LocalUserObserverBridge_onVideoTrackPublicationFailure),
			onVideoTrackUnpublished:              C.C_LocalUserObserverBridge_onVideoTrackUnpublished(C.cgo_LocalUserObserverBridge_onVideoTrackUnpublished),
			onLocalVideoTrackStateChanged:        C.C_LocalUserObserverBridge_onLocalVideoTrackStateChanged(C.cgo_LocalUserObserverBridge_onLocalVideoTrackStateChanged),
			onLocalVideoTrackStatistics:          C.C_LocalUserObserverBridge_onLocalVideoTrackStatistics(C.cgo_LocalUserObserverBridge_onLocalVideoTrackStatistics),
			onUserVideoTrackSubscribed:           C.C_LocalUserObserverBridge_onUserVideoTrackSubscribed(C.cgo_LocalUserObserverBridge_onUserVideoTrackSubscribed),
			onUserVideoTrackStateChanged:         C.C_LocalUserObserverBridge_onUserVideoTrackStateChanged(C.cgo_LocalUserObserverBridge_onUserVideoTrackStateChanged),
			onFirstRemoteVideoFrameRendered:      C.C_LocalUserObserverBridge_onFirstRemoteVideoFrameRendered(C.cgo_LocalUserObserverBridge_onFirstRemoteVideoFrameRendered),
			onRemoteVideoTrackStatistics:         C.C_LocalUserObserverBridge_onRemoteVideoTrackStatistics(C.cgo_LocalUserObserverBridge_onRemoteVideoTrackStatistics),
			onAudioVolumeIndication:              C.C_LocalUserObserverBridge_onAudioVolumeIndication(C.cgo_LocalUserObserverBridge_onAudioVolumeIndication),
			onActiveSpeaker:                      C.C_LocalUserObserverBridge_onActiveSpeaker(C.cgo_LocalUserObserverBridge_onActiveSpeaker),
			onAudioSubscribeStateChanged:         C.C_LocalUserObserverBridge_onAudioSubscribeStateChanged(C.cgo_LocalUserObserverBridge_onAudioSubscribeStateChanged),
			onVideoSubscribeStateChanged:         C.C_LocalUserObserverBridge_onVideoSubscribeStateChanged(C.cgo_LocalUserObserverBridge_onVideoSubscribeStateChanged),
			onAudioPublishStateChanged:           C.C_LocalUserObserverBridge_onAudioPublishStateChanged(C.cgo_LocalUserObserverBridge_onAudioPublishStateChanged),
			onVideoPublishStateChanged:           C.C_LocalUserObserverBridge_onVideoPublishStateChanged(C.cgo_LocalUserObserverBridge_onVideoPublishStateChanged),
			onRemoteSubscribeFallbackToAudioOnly: C.C_LocalUserObserverBridge_onRemoteSubscribeFallbackToAudioOnly(C.cgo_LocalUserObserverBridge_onRemoteSubscribeFallbackToAudioOnly),
			onFirstRemoteAudioFrame:              C.C_LocalUserObserverBridge_onFirstRemoteAudioFrame(C.cgo_LocalUserObserverBridge_onFirstRemoteAudioFrame),
			onFirstRemoteAudioDecoded:            C.C_LocalUserObserverBridge_onFirstRemoteAudioDecoded(C.cgo_LocalUserObserverBridge_onFirstRemoteAudioDecoded),
			onFirstRemoteVideoFrame:              C.C_LocalUserObserverBridge_onFirstRemoteVideoFrame(C.cgo_LocalUserObserverBridge_onFirstRemoteVideoFrame),
			onFirstRemoteVideoDecoded:            C.C_LocalUserObserverBridge_onFirstRemoteVideoDecoded(C.cgo_LocalUserObserverBridge_onFirstRemoteVideoDecoded),
			onVideoSizeChanged:                   C.C_LocalUserObserverBridge_onVideoSizeChanged(C.cgo_LocalUserObserverBridge_onVideoSizeChanged),
			onUserInfoUpdated:                    C.C_LocalUserObserverBridge_onUserInfoUpdated(C.cgo_LocalUserObserverBridge_onUserInfoUpdated),
			onIntraRequestReceived:               C.C_LocalUserObserverBridge_onIntraRequestReceived(C.cgo_LocalUserObserverBridge_onIntraRequestReceived),
			onStreamMessage:                      C.C_LocalUserObserverBridge_onStreamMessage(C.cgo_LocalUserObserverBridge_onStreamMessage),
			onUserStateChanged:                   C.C_LocalUserObserverBridge_onUserStateChanged(C.cgo_LocalUserObserverBridge_onUserStateChanged),
			onVideoRenderingTracingResult:        C.C_LocalUserObserverBridge_onVideoRenderingTracingResult(C.cgo_LocalUserObserverBridge_onVideoRenderingTracingResult),
		},
		userData,
	))
	b.handler = handler
	return &b
}

//export cgo_LocalUserObserverBridge_onAudioTrackPublishStart
func cgo_LocalUserObserverBridge_onAudioTrackPublishStart(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	audioTrack *C.C_ILocalAudioTrack) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnAudioTrackPublishStart((*agora.ILocalAudioTrack)(audioTrack))
}

//export cgo_LocalUserObserverBridge_onAudioTrackPublishSuccess
func cgo_LocalUserObserverBridge_onAudioTrackPublishSuccess(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	audioTrack *C.C_ILocalAudioTrack) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnAudioTrackPublishSuccess((*agora.ILocalAudioTrack)(audioTrack))
}

//export cgo_LocalUserObserverBridge_onAudioTrackUnpublished
func cgo_LocalUserObserverBridge_onAudioTrackUnpublished(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	audioTrack *C.C_ILocalAudioTrack) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnAudioTrackUnpublished((*agora.ILocalAudioTrack)(audioTrack))
}

//export cgo_LocalUserObserverBridge_onAudioTrackPublicationFailure
func cgo_LocalUserObserverBridge_onAudioTrackPublicationFailure(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	audioTrack *C.C_ILocalAudioTrack,
	_error C.enum_C_ERROR_CODE_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnAudioTrackPublicationFailure((*agora.ILocalAudioTrack)(audioTrack), agora.ERROR_CODE_TYPE(_error))
}

//export cgo_LocalUserObserverBridge_onLocalAudioTrackStatistics
func cgo_LocalUserObserverBridge_onLocalAudioTrackStatistics(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	stats *C.struct_C_LocalAudioStats) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnLocalAudioTrackStatistics((*agora.LocalAudioStats)(unsafe.Pointer(stats)))
}

//export cgo_LocalUserObserverBridge_onRemoteAudioTrackStatistics
func cgo_LocalUserObserverBridge_onRemoteAudioTrackStatistics(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	audioTrack *C.C_IRemoteAudioTrack,
	stats *C.struct_C_RemoteAudioTrackStats) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnRemoteAudioTrackStatistics(
		(*agora.IRemoteAudioTrack)(audioTrack),
		(*agora.RemoteAudioTrackStats)(unsafe.Pointer(stats)),
	)
}

//export cgo_LocalUserObserverBridge_onUserAudioTrackSubscribed
func cgo_LocalUserObserverBridge_onUserAudioTrackSubscribed(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t,
	audioTrack *C.C_IRemoteAudioTrack) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnUserAudioTrackSubscribed(
		C.GoString(userId),
		(*agora.IRemoteAudioTrack)(audioTrack),
	)
}

//export cgo_LocalUserObserverBridge_onUserAudioTrackStateChanged
func cgo_LocalUserObserverBridge_onUserAudioTrackStateChanged(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t,
	audioTrack *C.C_IRemoteAudioTrack,
	state C.enum_C_REMOTE_AUDIO_STATE,
	reason C.enum_C_REMOTE_AUDIO_STATE_REASON,
	elapsed C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnUserAudioTrackStateChanged(
		C.GoString(userId),
		(*agora.IRemoteAudioTrack)(audioTrack),
		agora.REMOTE_AUDIO_STATE(state),
		agora.REMOTE_AUDIO_STATE_REASON(reason),
		int(elapsed),
	)
}

//export cgo_LocalUserObserverBridge_onVideoTrackPublishStart
func cgo_LocalUserObserverBridge_onVideoTrackPublishStart(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	videoTrack *C.C_ILocalVideoTrack) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnVideoTrackPublishStart(
		(*agora.ILocalVideoTrack)(videoTrack),
	)
}

//export cgo_LocalUserObserverBridge_onVideoTrackPublishSuccess
func cgo_LocalUserObserverBridge_onVideoTrackPublishSuccess(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	videoTrack *C.C_ILocalVideoTrack) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnVideoTrackPublishSuccess(
		(*agora.ILocalVideoTrack)(videoTrack),
	)
}

//export cgo_LocalUserObserverBridge_onVideoTrackPublicationFailure
func cgo_LocalUserObserverBridge_onVideoTrackPublicationFailure(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	videoTrack *C.C_ILocalVideoTrack,
	_error C.enum_C_ERROR_CODE_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnVideoTrackPublicationFailure(
		(*agora.ILocalVideoTrack)(videoTrack),
		agora.ERROR_CODE_TYPE(_error),
	)
}

//export cgo_LocalUserObserverBridge_onVideoTrackUnpublished
func cgo_LocalUserObserverBridge_onVideoTrackUnpublished(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	videoTrack *C.C_ILocalVideoTrack) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnVideoTrackUnpublished(
		(*agora.ILocalVideoTrack)(videoTrack),
	)
}

//export cgo_LocalUserObserverBridge_onLocalVideoTrackStateChanged
func cgo_LocalUserObserverBridge_onLocalVideoTrackStateChanged(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	videoTrack *C.C_ILocalVideoTrack,
	state C.enum_C_LOCAL_VIDEO_STREAM_STATE,
	errorCode C.enum_C_LOCAL_VIDEO_STREAM_ERROR) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnLocalVideoTrackStateChanged(
		(*agora.ILocalVideoTrack)(videoTrack),
		agora.LOCAL_VIDEO_STREAM_STATE(state),
		agora.LOCAL_VIDEO_STREAM_ERROR(errorCode),
	)
}

//export cgo_LocalUserObserverBridge_onLocalVideoTrackStatistics
func cgo_LocalUserObserverBridge_onLocalVideoTrackStatistics(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	videoTrack *C.C_ILocalVideoTrack,
	stats *C.struct_C_LocalVideoTrackStats) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnLocalVideoTrackStatistics(
		(*agora.ILocalVideoTrack)(videoTrack),
		(*agora.LocalVideoTrackStats)(unsafe.Pointer(stats)),
	)
}

//export cgo_LocalUserObserverBridge_onUserVideoTrackSubscribed
func cgo_LocalUserObserverBridge_onUserVideoTrackSubscribed(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, trackInfo C.struct_C_VideoTrackInfo,
	videoTrack *C.C_IRemoteVideoTrack) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnUserVideoTrackSubscribed(
		C.GoString(userId),
		*(*agora.VideoTrackInfo)(unsafe.Pointer(&trackInfo)),
		(*agora.IRemoteVideoTrack)(videoTrack),
	)
}

//export cgo_LocalUserObserverBridge_onUserVideoTrackStateChanged
func cgo_LocalUserObserverBridge_onUserVideoTrackStateChanged(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t,
	videoTrack *C.C_IRemoteVideoTrack,
	state C.enum_C_REMOTE_VIDEO_STATE,
	reason C.enum_C_REMOTE_VIDEO_STATE_REASON,
	elapsed C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnUserVideoTrackStateChanged(
		C.GoString(userId),
		(*agora.IRemoteVideoTrack)(videoTrack),
		agora.REMOTE_VIDEO_STATE(state),
		agora.REMOTE_VIDEO_STATE_REASON(reason),
		int(elapsed),
	)
}

//export cgo_LocalUserObserverBridge_onFirstRemoteVideoFrameRendered
func cgo_LocalUserObserverBridge_onFirstRemoteVideoFrameRendered(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, width C.int, height C.int, elapsed C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnFirstRemoteVideoFrameRendered(
		C.GoString(userId),
		int(width),
		int(height),
		int(elapsed),
	)
}

//export cgo_LocalUserObserverBridge_onRemoteVideoTrackStatistics
func cgo_LocalUserObserverBridge_onRemoteVideoTrackStatistics(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	videoTrack *C.C_IRemoteVideoTrack, stats *C.struct_C_RemoteVideoTrackStats) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnRemoteVideoTrackStatistics(
		(*agora.IRemoteVideoTrack)(videoTrack),
		(*agora.RemoteVideoTrackStats)(unsafe.Pointer(stats)),
	)
}

//export cgo_LocalUserObserverBridge_onAudioVolumeIndication
func cgo_LocalUserObserverBridge_onAudioVolumeIndication(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	speakers *C.struct_C_AudioVolumeInformation, speakerNumber C.uint,
	totalVolume C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnAudioVolumeIndication(
		(*agora.AudioVolumeInformation)(unsafe.Pointer(speakers)),
		uint(speakerNumber),
		int(totalVolume),
	)
}

//export cgo_LocalUserObserverBridge_onActiveSpeaker
func cgo_LocalUserObserverBridge_onActiveSpeaker(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnActiveSpeaker(
		C.GoString(userId),
	)
}

//export cgo_LocalUserObserverBridge_onAudioSubscribeStateChanged
func cgo_LocalUserObserverBridge_onAudioSubscribeStateChanged(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	channel *C.char, userId C.C_user_id_t, oldState C.enum_C_STREAM_SUBSCRIBE_STATE, newState C.enum_C_STREAM_SUBSCRIBE_STATE, elapseSinceLastState C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnAudioSubscribeStateChanged(
		C.GoString(channel),
		C.GoString(userId),
		agora.STREAM_SUBSCRIBE_STATE(oldState),
		agora.STREAM_SUBSCRIBE_STATE(newState),
		int(elapseSinceLastState),
	)
}

//export cgo_LocalUserObserverBridge_onVideoSubscribeStateChanged
func cgo_LocalUserObserverBridge_onVideoSubscribeStateChanged(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	channel *C.char, userId C.C_user_id_t, oldState C.enum_C_STREAM_SUBSCRIBE_STATE, newState C.enum_C_STREAM_SUBSCRIBE_STATE, elapseSinceLastState C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnVideoSubscribeStateChanged(
		C.GoString(channel),
		C.GoString(userId),
		agora.STREAM_SUBSCRIBE_STATE(oldState),
		agora.STREAM_SUBSCRIBE_STATE(newState),
		int(elapseSinceLastState),
	)
}

//export cgo_LocalUserObserverBridge_onAudioPublishStateChanged
func cgo_LocalUserObserverBridge_onAudioPublishStateChanged(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	channel *C.char, oldState C.enum_C_STREAM_PUBLISH_STATE, newState C.enum_C_STREAM_PUBLISH_STATE, elapseSinceLastState C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnAudioPublishStateChanged(
		C.GoString(channel),
		agora.STREAM_PUBLISH_STATE(oldState),
		agora.STREAM_PUBLISH_STATE(newState),
		int(elapseSinceLastState),
	)
}

//export cgo_LocalUserObserverBridge_onVideoPublishStateChanged
func cgo_LocalUserObserverBridge_onVideoPublishStateChanged(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	channel *C.char, oldState C.enum_C_STREAM_PUBLISH_STATE, newState C.enum_C_STREAM_PUBLISH_STATE, elapseSinceLastState C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnVideoPublishStateChanged(
		C.GoString(channel),
		agora.STREAM_PUBLISH_STATE(oldState),
		agora.STREAM_PUBLISH_STATE(newState),
		int(elapseSinceLastState),
	)
}

//export cgo_LocalUserObserverBridge_onRemoteSubscribeFallbackToAudioOnly
func cgo_LocalUserObserverBridge_onRemoteSubscribeFallbackToAudioOnly(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, isFallbackOrRecover C.bool) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnRemoteSubscribeFallbackToAudioOnly(
		C.GoString(userId),
		bool(isFallbackOrRecover),
	)
}

//export cgo_LocalUserObserverBridge_onFirstRemoteAudioFrame
func cgo_LocalUserObserverBridge_onFirstRemoteAudioFrame(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, elapsed C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnFirstRemoteAudioFrame(
		C.GoString(userId),
		int(elapsed),
	)
}

//export cgo_LocalUserObserverBridge_onFirstRemoteAudioDecoded
func cgo_LocalUserObserverBridge_onFirstRemoteAudioDecoded(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, elapsed C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnFirstRemoteAudioDecoded(
		C.GoString(userId),
		int(elapsed),
	)
}

//export cgo_LocalUserObserverBridge_onFirstRemoteVideoFrame
func cgo_LocalUserObserverBridge_onFirstRemoteVideoFrame(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, width C.int, height C.int, elapsed C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnFirstRemoteVideoFrame(
		C.GoString(userId),
		int(width),
		int(height),
		int(elapsed),
	)
}

//export cgo_LocalUserObserverBridge_onFirstRemoteVideoDecoded
func cgo_LocalUserObserverBridge_onFirstRemoteVideoDecoded(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, width C.int, height C.int, elapsed C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnFirstRemoteVideoDecoded(
		C.GoString(userId),
		int(width),
		int(height),
		int(elapsed),
	)
}

//export cgo_LocalUserObserverBridge_onVideoSizeChanged
func cgo_LocalUserObserverBridge_onVideoSizeChanged(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, width C.int, height C.int, rotation C.int) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnVideoSizeChanged(
		C.GoString(userId),
		int(width),
		int(height),
		int(rotation),
	)
}

//export cgo_LocalUserObserverBridge_onUserInfoUpdated
func cgo_LocalUserObserverBridge_onUserInfoUpdated(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, msg C.enum_C_USER_MEDIA_INFO, val C.bool) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnUserInfoUpdated(
		C.GoString(userId),
		agora.USER_MEDIA_INFO(msg),
		bool(val),
	)
}

//export cgo_LocalUserObserverBridge_onIntraRequestReceived
func cgo_LocalUserObserverBridge_onIntraRequestReceived(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnIntraRequestReceived()
}

//export cgo_LocalUserObserverBridge_onStreamMessage
func cgo_LocalUserObserverBridge_onStreamMessage(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	user_id C.C_user_id_t, streamId C.int, data *C.char, length C.size_t) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnStreamMessage(
		C.GoString(user_id),
		int(streamId),
		C.GoBytes(unsafe.Pointer(data), C.int(length)),
		uint(length),
	)
}

//export cgo_LocalUserObserverBridge_onUserStateChanged
func cgo_LocalUserObserverBridge_onUserStateChanged(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	user_id C.C_user_id_t, state C.uint32_t) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnUserStateChanged(
		C.GoString(user_id),
		uint32(state),
	)
}

//export cgo_LocalUserObserverBridge_onVideoRenderingTracingResult
func cgo_LocalUserObserverBridge_onVideoRenderingTracingResult(_ *C.C_LocalUserObserverBridge, userData unsafe.Pointer,
	user_id C.C_user_id_t, currentState C.enum_C_MEDIA_TRACE_EVENT, tracingInfo C.struct_C_VideoRenderingTracingInfo) {
	if userData == nil {
		return
	}

	bridge := (*LocalUserObserverBridge)(userData)
	bridge.handler.OnVideoRenderingTracingResult(
		C.GoString(user_id),
		agora.MEDIA_TRACE_EVENT(currentState),
		*(*agora.VideoRenderingTracingInfo)(unsafe.Pointer(&tracingInfo)),
	)
}
