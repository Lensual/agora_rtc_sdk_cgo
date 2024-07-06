package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "C_IAgoraService.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

// #region agora

//TODO

// #region agora::rtc

/**
 * The audio encoder configuration.
 */
type AudioEncoderConfiguration C.struct_C_AudioEncoderConfiguration

// #region AudioEncoderConfiguration

// /**
//  * The audio profile: #AUDIO_PROFILE_TYPE
//  */
// func (this_ *AudioEncoderConfiguration) GetAudioProfile() AUDIO_PROFILE_TYPE {
// 	return AUDIO_PROFILE_TYPE(this_.audioProfile)
// }

// /**
//  * The audio profile: #AUDIO_PROFILE_TYPE
//  */
// func (this_ *AudioEncoderConfiguration) SetAudioProfile(audioProfile AUDIO_PROFILE_TYPE) {
// 	this_.audioProfile = C.enum_AUDIO_PROFILE_TYPE(audioProfile)
// }

// func NewAudioEncoderConfiguration() *AudioEncoderConfiguration {
// 	return (*AudioEncoderConfiguration)(C.C_AudioEncoderConfiguration_New())
// }
// func (this_ *AudioEncoderConfiguration) Delete() {
// 	C.C_AudioEncoderConfiguration_Delete((*C.struct_C_AudioEncoderConfiguration)(this_))
// }

// #endregion AudioEncoderConfiguration

// #endregion agora::rtc

// #region agora::rtm

// #endregion agora::rtm

// #region agora::base

type AgoraServiceConfiguration C.struct_C_AgoraServiceConfiguration

// #region AgoraServiceConfiguration

/**
 * Whether to enable the audio processing module.
 * - `true`: (Default) Enable the audio processing module.
 * - `false`: Disable the audio processing module. If you disable the audio processing module, you cannot create audio tracks.
 */
func (this_ *AgoraServiceConfiguration) GetEnableAudioProcessor() bool {
	return bool(this_.enableAudioProcessor)
}

/**
 * Whether to enable the audio processing module.
 * - `true`: (Default) Enable the audio processing module.
 * - `false`: Disable the audio processing module. If you disable the audio processing module, you cannot create audio tracks.
 */
func (this_ *AgoraServiceConfiguration) SetEnableAudioProcessor(enableAudioProcessor bool) {
	this_.enableAudioProcessor = C.bool(enableAudioProcessor)
}

/**
 * Whether to enable the audio device module. The function of the audio device module is to manage audio devices,
 * such as recording and playing audio.
 * - `true`: (Default) Enable the audio device module. Audio recording and playback is available.
 * - `false`: Disable the audio device module. Audio recording and playback is unavailable.
 *
 * @note
 * If you set `enableAudioDevice` as `false` and set `enableAudioProcessor` as `true`, you cannot use audio devices,
 * but you can push PCM audio data.
 */
func (this_ *AgoraServiceConfiguration) GetEnableAudioDevice() bool {
	return bool(this_.enableAudioDevice)
}

/**
 * Whether to enable the audio device module. The function of the audio device module is to manage audio devices,
 * such as recording and playing audio.
 * - `true`: (Default) Enable the audio device module. Audio recording and playback is available.
 * - `false`: Disable the audio device module. Audio recording and playback is unavailable.
 *
 * @note
 * If you set `enableAudioDevice` as `false` and set `enableAudioProcessor` as `true`, you cannot use audio devices,
 * but you can push PCM audio data.
 */
func (this_ *AgoraServiceConfiguration) SetEnableAudioDevice(enableAudioDevice bool) {
	this_.enableAudioDevice = C.bool(enableAudioDevice)
}

/**
 * Whether to enable video.
 * - `true`: Enable video.
 * - `false`: (Default) Disable video.
 */
func (this_ *AgoraServiceConfiguration) GetEnableVideo() bool {
	return bool(this_.enableVideo)
}

/**
 * Whether to enable video.
 * - `true`: Enable video.
 * - `false`: (Default) Disable video.
 */
func (this_ *AgoraServiceConfiguration) SetEnableVideo(enableVideo bool) {
	this_.enableVideo = C.bool(enableVideo)
}

/**
 * The user context.
 * - For Windows, it is the handle of the window that loads the video. Specify this value to support plugging or unplugging the video devices while the host is powered on.
 * - For Android, it is the context of activity.
 */
func (this_ *AgoraServiceConfiguration) GetContext() unsafe.Pointer {
	return this_.context
}

/**
 * The user context.
 * - For Windows, it is the handle of the window that loads the video. Specify this value to support plugging or unplugging the video devices while the host is powered on.
 * - For Android, it is the context of activity.
 */
func (this_ *AgoraServiceConfiguration) SetContext(context unsafe.Pointer) {
	this_.context = context
}

/**
 * The App ID of your project.
 */
func (this_ *AgoraServiceConfiguration) GetAppId() string {
	return C.GoString(this_.appId)
}

/**
 * The App ID of your project.
 */
func (this_ *AgoraServiceConfiguration) SetAppId(appId string) {
	this_.appId = C.CString(appId)
}

/**
 * The supported area code, default is AREA_CODE_GLOB
 */
func (this_ *AgoraServiceConfiguration) GetAreaCode() uint {
	return uint(this_.areaCode)
}

/**
 * The supported area code, default is AREA_CODE_GLOB
 */
func (this_ *AgoraServiceConfiguration) SetAreaCode(areaCode uint) {
	this_.areaCode = C.uint(areaCode)
}

//  /** The channel profile. For details, see \ref agora::CHANNEL_PROFILE_TYPE "CHANNEL_PROFILE_TYPE". The default channel profile is `CHANNEL_PROFILE_LIVE_BROADCASTING`.
//   */
//  enum C_CHANNEL_PROFILE_TYPE channelProfile;

//  /**
//   * The license used for verification when connectting channel. Charge according to the license
//   */
//  const char *license;

//  /**
//   * The audio scenario. See \ref agora::rtc::AUDIO_SCENARIO_TYPE "AUDIO_SCENARIO_TYPE". The default value is `AUDIO_SCENARIO_DEFAULT`.
//   */
//  enum C_AUDIO_SCENARIO_TYPE audioScenario;
//  /**
//   * The config for custumer set log path, log size and log level.
//   */
//  struct C_LogConfig logConfig;

/**
 * Whether to enable string uid.
 */
func (this_ *AgoraServiceConfiguration) GetUseStringUid() bool {
	return bool(this_.useStringUid)
}

/**
 * Whether to enable string uid.
 */
func (this_ *AgoraServiceConfiguration) SetUseStringUid(useStringUid bool) {
	this_.useStringUid = C.bool(useStringUid)
}

//  /**
//   * The service observer.
//   */
//  // C_IServiceObserver *serviceObserver;
//  void *serviceObserver; // HACK

/**
 * Thread priority for SDK common threads
 */
func (this_ *AgoraServiceConfiguration) GetThreadPriority() THREAD_PRIORITY_TYPE {
	return THREAD_PRIORITY_TYPE(this_.threadPriority)
}

/**
 * Thread priority for SDK common threads
 */
func (this_ *AgoraServiceConfiguration) SetThreadPriority(threadPriority THREAD_PRIORITY_TYPE) {
	this_.threadPriority = C.enum_C_THREAD_PRIORITY_TYPE(threadPriority)
}

/**
 * Whether use egl context in current thread as sdk‘s root egl context
 * which shared by all egl related modules. eg. camera capture, video renderer.
 * @note
 * This property applies to Android only.
 */
func (this_ *AgoraServiceConfiguration) GetUseExternalEglContext() bool {
	return bool(this_.useExternalEglContext)
}

/**
 * Whether use egl context in current thread as sdk‘s root egl context
 * which shared by all egl related modules. eg. camera capture, video renderer.
 * @note
 * This property applies to Android only.
 */
func (this_ *AgoraServiceConfiguration) SetUseExternalEglContext(useExternalEglContext bool) {
	this_.useExternalEglContext = C.bool(useExternalEglContext)
}

/**
 * Determines whether to enable domain limit.
 * - `true`: only connect to servers that already parsed by DNS
 * - `false`: (Default) connect to servers with no limit
 */
func (this_ *AgoraServiceConfiguration) GetDomainLimit() bool {
	return bool(this_.domainLimit)
}

/**
 * Determines whether to enable domain limit.
 * - `true`: only connect to servers that already parsed by DNS
 * - `false`: (Default) connect to servers with no limit
 */
func (this_ *AgoraServiceConfiguration) SetDomainLimit(domainLimit bool) {
	this_.domainLimit = C.bool(domainLimit)
}

func NewAgoraServiceConfiguration() *AgoraServiceConfiguration {
	return (*AgoraServiceConfiguration)(C.C_AgoraServiceConfiguration_New())
}
func (this_ *AgoraServiceConfiguration) Delete() {
	C.C_AgoraServiceConfiguration_Delete((*C.struct_C_AgoraServiceConfiguration)(this_))
}

// #endregion AgoraServiceConfiguration

// #region IServiceObserver

type IServiceObserver C.C_IServiceObserver

func (this_ *IServiceObserver) Delete() {
	C.C_IServiceObserver_Delete(unsafe.Pointer(this_))
}

/**
 * Reports the permission error.
 * @param permission {@link PERMISSION}
 */
func (this_ *IServiceObserver) onPermissionError(permissionType PERMISSION_TYPE) {
	C.C_IServiceObserver_onPermissionError(unsafe.Pointer(this_),
		C.enum_C_PERMISSION_TYPE(permissionType),
	)
}

/**
 * Reports the audio device error.
 * @param error {@link ERROR_CODE_TYPE}
 */
func (this_ *IServiceObserver) onAudioDeviceError(_error ERROR_CODE_TYPE, description string) {
	cDescription := C.CString(description)
	C.C_IServiceObserver_onAudioDeviceError(unsafe.Pointer(this_),
		C.enum_C_ERROR_CODE_TYPE(_error),
		cDescription,
	)
	C.free(unsafe.Pointer(cDescription))
}

/**
 * Reports the config fetch result.
 *
 * @param code The error code of fetching config.
 *  - 0(ERR_OK): Success.
 *  - 10(ERR_TIMEDOUT): Fetching config is timed out.
 * @param configType The type of fetching config.
 *  - 1(CONFIG_FETCH_TYPE_INITIALIZE): Fetch config when initializing RtcEngine without channel info.
 *  - 2(CONFIG_FETCH_TYPE_JOIN_CHANNEL): Fetch config when joining channel with channel info, such as channel name and uid.
 * @param configContent The config fetched from server.
 */
func (this_ *IServiceObserver) onFetchConfigResult(code int, configType CONFIG_FETCH_TYPE, configContent string) {
	cConfigContent := C.CString(configContent)
	C.C_IServiceObserver_onFetchConfigResult(unsafe.Pointer(this_),
		C.int(code),
		C.enum_C_CONFIG_FETCH_TYPE(configType),
		cConfigContent,
	)
	C.free(unsafe.Pointer(cConfigContent))

}

// #endregion IServiceObserver

/**
 * The IAgoraService class.
 *
 * `IAgoraService` is the entry point of Agora low-level APIs. Use this interface to
 * create access points to Agora interfaces, including RTC connections and media tracks.
 *
 * You can create an `IAgoraService` object by calling \ref agora::base::IAgoraService::createAgoraService "createAgoraService".
 *
 * You can configure the `IAgoraService` object for different user scenarios on the global level by using `AgoraServiceConfiguration`.
 */
type IAgoraService C.C_IAgoraService

// #region IAgoraService

/**
 * Initializes the \ref agora::base::IAgoraService "AgoraService" object.
 *
 * @param config The configuration of the initialization. For details, see \ref agora::base::AgoraServiceConfiguration "AgoraServiceConfiguration".
 * @return
 * - 0: Success.
 * - < 0: Failure.
 *   - `ERR_INVALID_ARGUMENT`, if `context` in `AgoraServiceConfiguration` is not provided for
 * Android.
 *   - `ERR_INIT_NET_ENGINE`, if the network engine cannot be initialized. On Windows, the error occurs mostly because the connection to the local port is disabled by the firewall. In this case, turn off the firewall and then turn it on again.
 */
func (this_ *IAgoraService) Initialize(config *AgoraServiceConfiguration) int {
	return int(C.C_IAgoraService_initialize(unsafe.Pointer(this_), (*C.struct_C_AgoraServiceConfiguration)(config)))
}

//   /**
//    * Flush log & cache before exit
//    */
//    void C_IAgoraService_atExit(C_IAgoraService *this_);

/**
* Releases the \ref agora::base::IAgoraService "AgoraService" object.
*
* @return
* - 0: Success.
* - < 0: Failure.
 */
func (this_ *IAgoraService) Release() int {
	return int(C.C_IAgoraService_release(unsafe.Pointer(this_)))
}

//    /**
// 	* Configures the preset audio scenario.
// 	*
// 	* @param scenario The preset audio scenario: \ref agora::rtc::AUDIO_SCENARIO_TYPE
// 	* "AUDIO_SCENARIO_TYPE".
// 	* @return
// 	* - 0: Success.
// 	* - < 0: Failure.
// 	*/
//    int C_IAgoraService_setAudioSessionPreset(C_IAgoraService *this_, enum C_AUDIO_SCENARIO_TYPE scenario);

//    /**
// 	* Customizes the audio session configuration.
// 	*
// 	* @param config The reference to the audio session configuration: \ref agora::base::AudioSessionConfiguration "AudioSessionConfiguration".
// 	* @return
// 	* - 0: Success.
// 	* - < 0: Failure.
// 	*/
//    int C_IAgoraService_setAudioSessionConfiguration(C_IAgoraService *this_, const struct C_AudioSessionConfiguration *config);

//    /**
// 	* Gets the audio session configuration.
// 	*
// 	* @param [out] config The pointer to the audio session configuration: \ref agora::base::AudioSessionConfiguration "AudioSessionConfiguration".
// 	* @return
// 	* - 0: Success.
// 	* - < 0: Failure.
// 	*/
//    int C_IAgoraService_getAudioSessionConfiguration(C_IAgoraService *this_, struct C_AudioSessionConfiguration *config);

//    /**
// 	* Sets the path and size of the SDK log files.
// 	*
// 	* The SDK records all the log data during the SDK runtime in two log files,
// 	* each with a default size of 512 KB. If you set `fileSize` as 1024 KB,
// 	* the SDK outputs log files with a maximum size of 2 MB. If the total size
// 	* of the log files exceeds the set value, the new output log
// 	* overwrites the old output log.
// 	*
// 	* @note
// 	* To ensure that the output log is complete, call this method immediately after calling
// 	* \ref agora::base::IAgoraService::initialize "initialize".
// 	*
// 	* @param filePath The pointer to the log file. Ensure that the directory of the log file exists and is writable.
// 	* @param fileSize The size of the SDK log file size (KB).
// 	* @return
// 	* - 0: Success.
// 	* - < 0: Failure.
// 	*/
//    int C_IAgoraService_setLogFile(C_IAgoraService *this_, const char *filePath, unsigned int fileSize);
//    /**
// 	* Sets the SDK log output filter.
// 	*
// 	* The log level follows the sequence of OFF, CRITICAL, ERROR, WARNING, INFO, and DEBUG.
// 	*
// 	* Select a level to output the logs in and above the selected level.
// 	* For example, if you set the log level to WARNING, you can see the logs in the levels of CRITICAL, ERROR, and WARNING.
// 	*
// 	* @param filters The log output filter.
// 	* - `LOG_LEVEL_NONE (0x0000)`: Do not output any log file.
// 	* - `LOG_LEVEL_INFO (0x0001)`: (Recommended) Output log files of the INFO level.
// 	* - `LOG_LEVEL_WARN (0x0002)`: Output log files of the WARN level.
// 	* - `LOG_LEVEL_ERROR (0x0004)`: Output log files of the ERROR level.
// 	* - `LOG_LEVEL_FATAL (0x0008)`: Output log files of the FATAL level.
// 	* @return
// 	* - 0: Success.
// 	* - < 0: Failure.
// 	*/
//    int C_IAgoraService_setLogFilter(C_IAgoraService *this_, unsigned int filters);

/**
* Creates an \ref agora::rtc::IRtcConnection "RtcConnection" object and returns the pointer.
*
* @param cfg The reference to the RTC connection configuration: \ref agora::rtc::RtcConnectionConfiguration "RtcConnectionConfiguration".
* @return
* - The pointer to \ref rtc::IRtcConnection "IRtcConnection": Success.
* - A null pointer: Failure.
 */
func (this_ *IAgoraService) CreateRtcConnection(cfg *RtcConnectionConfiguration) *IRtcConnection {
	return (*IRtcConnection)(C.C_IAgoraService_createRtcConnection(unsafe.Pointer(this_), (*C.struct_C_RtcConnectionConfiguration)(cfg)))
}

//    // agora_refptr<rtc::IRtmpConnection>
//    // C_agora_refptr C_IAgoraService_createRtmpConnection(C_IAgoraService *this_,
//    //                                                     const struct C_RtmpConnectionConfiguration *cfg);

//    /**
// 	* Creates a local audio track object and returns the pointer.
// 	*
// 	* By default, the audio track is created from the selected audio input device, such as
// 	* the built-in microphone on a mobile device.
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	* - `INVALID_STATE`, if `enableAudioProcessor` in \ref agora::base::AgoraServiceConfiguration "AgoraServiceConfiguration" is set as `false`.
// 	*/
//    C_agora_refptr C_IAgoraService_createLocalAudioTrack(C_IAgoraService *this_);

//    /**
// 	* Creates a local mixed audio track object and returns the pointer.
// 	*
// 	* By default, the audio track is created from mix source, which could mixed target track.
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createLocalMixedAudioTrack(C_IAgoraService *this_);

//    /**
// 	* Creates a local audio track object with a PCM data sender and returns the pointer.
// 	*
// 	* Once created, this track can be used to send PCM audio data.
// 	*
// 	* @param audioSource The pointer to the PCM audio data sender: \ref agora::rtc::IAudioPcmDataSender "IAudioPcmDataSender".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	* - `INVALID_STATE`, if `enableAudioProcessor` in \ref agora::base::AgoraServiceConfiguration "AgoraServiceConfiguration" is set as `false`.
// 	*/
//    // C_agora_refptr C_IAgoraService_createCustomAudioTrack_pcmSender(C_IAgoraService *this_,
//    //                                                       C_agora_refptr audioSource);

//    /**
// 	* Creates a local audio track object with a PCM data sender and returns the pointer.
// 	* The source is not intended to be mixed with other source.
// 	*
// 	* Once created, this track can be used to send PCM audio data.
// 	*
// 	* @param audioSource The pointer to the PCM audio data sender: \ref agora::rtc::IAudioPcmDataSender "IAudioPcmDataSender".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	* - `INVALID_STATE`, if `enableAudioProcessor` in \ref agora::base::AgoraServiceConfiguration "AgoraServiceConfiguration" is set as `false`.
// 	*/
//    C_agora_refptr C_IAgoraService_createDirectCustomAudioTrack(C_IAgoraService *this_,
// 															   C_agora_refptr audioSource);

//    /**
// 	* Creates a local audio track object with a PCM data sender and returns the pointer.
// 	*
// 	* Once created, this track can be used to send PCM audio data.
// 	*
// 	* @param audioSource The pointer to the PCM audio data sender: \ref agora::rtc::IAudioPcmDataSender "IAudioPcmDataSender".
// 	* @param enableAec Whether enable audio echo cancellation for PCM audio data.
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	* - `INVALID_STATE`, if `enableAudioProcessor` in \ref agora::base::AgoraServiceConfiguration "AgoraServiceConfiguration" is set as `false`.
// 	*/
//    C_agora_refptr C_IAgoraService_createCustomAudioTrack_pcmSender_aec(C_IAgoraService *this_,
// 																	C_agora_refptr audioSource, bool enableAec);

//    /**
// 	* Creates a local audio track object with a audio mixer source and returns the pointer.
// 	*
// 	* Once created, this track can be used to send PCM audio data.
// 	*
// 	* @param audioSource The pointer to the audio mixer source : \ref agora::rtc::IRemoteAudioMixerSource "IRemoteAudioMixerSource".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack", if the method call succeeds.
// 	* - A null pointer, if the method call fails.
// 	* - `INVALID_STATE`, if `enableAudioProcessor` in `AgoraServiceConfiguration` is set as `false`.
// 	*/
//    C_agora_refptr C_IAgoraService_createCustomAudioTrack_mixerSource(C_IAgoraService *this_,
// 																	 C_agora_refptr audioSource);

//    /**
// 	* Creates a local audio track object with an encoded audio frame sender and returns the pointer.
// 	*
// 	* Once created, this track can be used to send encoded audio frames, such as Opus frames.
// 	*
// 	* @param audioSource The pointer to the encoded audio frame sender: \ref agora::rtc::IAudioEncodedFrameSender "IAudioEncoderFrameSender".
// 	* @param mixMode The mixing mode of the encoded audio in the channel: #TMixMode.
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	*   - `INVALID_STATE`, if `enableAudioProcessor` is set as `false` in \ref agora::base::AgoraServiceConfiguration "AgoraServiceConfiguration".
// 	*/
//    C_agora_refptr C_IAgoraService_createCustomAudioTrack_frameSender(C_IAgoraService *this_,
// 																	 C_agora_refptr audioSource, enum C_TMixMode mixMode);
//    /// @cond (!Linux)
//    /**
// 	* Creates a local audio track object with a media packet sender and returns the pointer.
// 	*
// 	* Once created, this track can be used to send audio packets, such as customized UDP/RTP packets.
// 	*
// 	* @param source The pointer to the media packet sender: \ref agora::rtc::IMediaPacketSender "IMediaPacketSender".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	* - `INVALID_STATE`, if `enableAudioProcessor` is set as `false` in \ref agora::base::AgoraServiceConfiguration "AgoraServiceConfiguration".
// 	*/
//    C_agora_refptr C_IAgoraService_createCustomAudioTrack_packetSender(C_IAgoraService *this_,
// 																	  C_agora_refptr source);
//    /// @endcond
//    /**
// 	* Creates a local audio track object with an IMediaPlayerSource object and returns the pointer.
// 	*
// 	* Once created, this track can be used to send PCM audio data decoded from a media player.
// 	*
// 	* @param audioSource The pointer to the player source. See \ref agora::rtc::IMediaPlayerSource "IMediaPlayerSource".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	* - `INVALID_STATE`, if `enableAudioProcessor` is set as `false` in \ref agora::base::AgoraServiceConfiguration "AgoraServiceConfiguration".
// 	*/
//    C_agora_refptr C_IAgoraService_createMediaPlayerAudioTrack(C_IAgoraService *this_,
// 															  C_agora_refptr audioSource);

//    /**
// 	* Creates a local audio track object with an IMediaStreamingSource object and returns the pointer.
// 	*
// 	* Once created, this track can be used to send encoded audio data which demuxed from a media streaming.
// 	*
// 	* @param streamingSource The pointer to the streaming source. See \ref agora::rtc::IMediaStreamingSource "IMediaStreamingSource".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	* - `INVALID_STATE`, if `enableAudioProcessor` is set as `false` in \ref agora::base::AgoraServiceConfiguration "AgoraServiceConfiguration".
// 	*/
//    C_agora_refptr C_IAgoraService_createMediaStreamingAudioTrack(C_IAgoraService *this_,
// 																 C_agora_refptr streamingSource);

//    /**
// 	* Creates a local audio track object with the recording device source and returns the pointer.
// 	*
// 	* Once created, this track can be used to send audio data got from a recording device.
// 	* @param audioSource The pointer to the recording device source. See \ref agora::rtc::IRecordingDeviceSource "IRecordingDeviceSource".
// 	* @param enableAec Whether enable audio echo cancellation for loopback recording. If loopback
// 	*                  recording device is a virtual sound card, it should be false, or it should be true.
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createRecordingDeviceAudioTrack(C_IAgoraService *this_,
// 																  C_agora_refptr audioSource, bool enableAec);

//    /**
// 	* Creates an audio device manager object and returns the pointer.
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::INGAudioDeviceManager "INGAudioDeviceManager": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createAudioDeviceManager(C_IAgoraService *this_);

//    /**
// 	* Creates a media node factory object and returns the pointer.
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::IMediaNodeFactory "IMediaNodeFactory": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createMediaNodeFactory(C_IAgoraService *this_);

//    /**
// 	* Creates a local video track object with a camera capturer and returns the pointer.
// 	*
// 	* Once created, this track can be used to send video data captured by the camera.
// 	*
// 	* @param videoSource The pointer to the camera capturer: \ref agora::rtc::ICameraCapturer "ICameraCapturer".
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::ILocalVideoTrack "ILocalVideoTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createCameraVideoTrack(C_IAgoraService *this_,
// 														 C_agora_refptr videoSource, const char *id);

//    /**
// 	* Creates a local video track object with a screen capturer and returns the pointer.
// 	*
// 	* Once created, this track can be used to send video data for screen sharing.
// 	*
// 	* @param videoSource The pointer to the screen capturer: \ref agora::rtc::IScreenCapturer "IScreenCapturer".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalVideoTrack "ILocalVideoTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createScreenVideoTrack(C_IAgoraService *this_,
// 														 C_agora_refptr videoSource, const char *id);

//    /**
// 	* Creates a local video track object with a video mixer and returns the pointer.
// 	*
// 	* Once created, this track can be used to send video data processed by the video mixer.
// 	*
// 	* @param videoSource The pointer to the video mixer. See \ref agora::rtc::IVideoMixerSource "IVideoMixerSource".
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::ILocalVideoTrack "ILocalVideoTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createMixedVideoTrack(C_IAgoraService *this_, C_agora_refptr videoSource,
// 														const char *id);

//    /**
// 	* Creates a local video track object with a video frame transceiver and returns the pointer.
// 	*
// 	* Once created, this track can be used to send video data processed by the transceiver.
// 	*
// 	* @param transceiver The pointer to the video transceiver. See \ref agora::rtc::IVideoFrameTransceiver "IVideoFrameTransceiver".
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::ILocalVideoTrack "ILocalVideoTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createTranscodedVideoTrack(C_IAgoraService *this_, C_agora_refptr transceiver,
// 															 const char *id);

//    /// @cond (!RTSA)
//    /**
// 	* Creates a local video track object with a customized video source and returns the pointer.
// 	*
// 	* Once created, this track can be used to send video data from a customized source.
// 	*
// 	* @param videoSource The pointer to the customized video frame sender: \ref agora::rtc::IVideoFrameSender "IVideoFrameSender".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalVideoTrack "ILocalVideoTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createCustomVideoTrack_frameSender(C_IAgoraService *this_,
// 														 C_agora_refptr videoSource, const char *id);
//    /// @endcond

//    /**
// 	* Creates a local video track object with an encoded video image sender and returns the pointer.
// 	*
// 	* Once created, this track can be used to send encoded video images, such as H.264 or VP8 frames.
// 	*
// 	* @param videoSource The pointer to the encoded video frame sender. See \ref agora::rtc::IVideoEncodedImageSender "IVideoEncodedImageSender".
// 	* @param options The configuration for creating video encoded image track.
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::ILocalVideoTrack "ILocalVideoTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createCustomVideoTrack_encoded(C_IAgoraService *this_,
// 																 C_agora_refptr videoSource,
// 																 const struct C_SenderOptions *options,
// 																 const char *id);

//  #if defined(__ANDROID__) || (defined(TARGET_OS_IPHONE) && TARGET_OS_IPHONE)
//    /**
// 	* Creates a local video track object with a screen capture source extension and returns the pointer.
// 	*
// 	* Once created, this track can be used to work with the screen capture extension.
// 	*
// 	* @param screen The pointer to the screen capture source.
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::ILocalVideoTrack "ILocalVideoTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createScreenCaptureVideoTrack(C_IAgoraService *this_,
// 																C_agora_refptr screen);

//    /**
// 	* Creates a local audio track object with a screen capture source extension and returns the pointer.
// 	*
// 	* Once created, this track can be used to work with the screen capture extension.
// 	*
// 	* @param screen The pointer to the screen capture source.
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createScreenCaptureAudioTrack(C_IAgoraService *this_,
// 																C_agora_refptr screen);
//  #else
//  /**
//   * Creates a local video track object with a screen capture source extension and returns the pointer.
//   *
//   * Once created, this track can be used to work with the screen capture extension.
//   *
//   * @param screen The pointer to the screen capture source.
//   *
//   * @return
//   * - The pointer to \ref rtc::ILocalVideoTrack "ILocalVideoTrack": Success.
//   * - A null pointer: Failure.
//   */
//  C_agora_refptr C_IAgoraService_createScreenCaptureVideoTrack(C_IAgoraService *this_,
// 															  C_agora_refptr screen, const char *id);
//  #endif

//    /// @cond (!Linux)
//    /**
// 	* Creates a local video track object with a media packet sender and returns the pointer.
// 	*
// 	* Once created, this track can be used to send video packets, such as customized UDP/RTP packets.
// 	*
// 	* @param source The pointer to the media packet sender: \ref agora::rtc::IMediaPacketSender "IMediaPacketSender".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalVideoTrack "ILocalVideoTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createCustomVideoTrack_packetSender(C_IAgoraService *this_,
// 														 C_agora_refptr source, const char *id);
//    /// @endcond
//    /**
// 	* Creates a local video track object with an IMediaPlayerSource object and returns the pointer.
// 	*
// 	* Once created, this track can be used to send YUV frames decoded from a player.
// 	*
// 	* @param videoSource The pointer to the player source. See \ref agora::rtc::IMediaPlayerSource "IMediaPlayerSource".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createMediaPlayerVideoTrack(C_IAgoraService *this_,
// 															  C_agora_refptr videoSource, const char *id);

//    /**
// 	* Creates a local video track object with an IMediaStreamingSource object and returns the pointer.
// 	*
// 	* Once created, this track can be used to send H264 frames which demuxed from a streaming.
// 	*
// 	* @param streamingSource The pointer to the player source. See \ref agora::rtc::IMediaStreamingSource "IMediaStreamingSource".
// 	* @return
// 	* - The pointer to \ref rtc::ILocalAudioTrack "ILocalAudioTrack": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createMediaStreamingVideoTrack(C_IAgoraService *this_,
// 																 C_agora_refptr streamingSource, const char *id);

//    /**
// 	* Creates an RTMP streaming service object and returns the pointer.
// 	*
// 	* @param rtcConnection The pointer to \ref rtc::IRtcConnection "IRtcConnection".
// 	* @param appId The App ID of the live streaming service.
// 	* @return
// 	* - The pointer to \ref rtc::IRtmpStreamingService "IRtmpStreamingService": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_createRtmpStreamingService(C_IAgoraService *this_,
// 															 C_agora_refptr rtcConnection, const char *appId);

//    /**
// 	* Creates an Media Relay service object and returns the pointer.
// 	*
// 	* @param rtcConnection The pointer to \ref rtc::IRtcConnection "IRtcConnection".
// 	* @param appId The App ID of the media relay service.
// 	* @return
// 	* - The pointer to \ref rtc::IMediaRelayService "IMediaRelayService", if the method call
// 	* succeeds.
// 	* - A null pointer, if the method call fails.
// 	*/
//    C_agora_refptr C_IAgoraService_createMediaRelayService(C_IAgoraService *this_,
// 														  C_agora_refptr rtcConnection, const char *appId);

//    /**
// 	* Creates an file upload object and returns the pointer.
// 	*
// 	* @param rtcConnection The pointer to \ref rtc::IRtcConnection "IRtcConnection".
// 	* @param appId The App ID of the media relay service.
// 	* @return
// 	* - The pointer to \ref rtc::IFileUploaderService "IFileUploaderService", if the method call
// 	* succeeds.
// 	* - A null pointer, if the method call fails.
// 	*/
//    C_agora_refptr C_IAgoraService_createFileUploadService(C_IAgoraService *this_,
// 														  C_agora_refptr rtcConnection, const char *appId);

//    // /**
//    //  * Creates an RTM servive object and returns the pointer.
//    //  *
//    //  * @return
//    //  * - The pointer to \ref rtm::IRtmService "IRtmService": Success.
//    //  * - A null pointer: Failure.
//    //  */
//    // C_IRtmService *C_IAgoraService_createRtmService();

//    int C_IAgoraService_addExtensionObserver(C_IAgoraService *this_, C_agora_refptr observer);

//    int C_IAgoraService_removeExtensionObserver(C_IAgoraService *this_, C_agora_refptr observer);

//    // /**
//    //  * Creates an audio device manager and returns the pointer.
//    //  *
//    //  * @return
//    //  * - The pointer to \ref rtc::IAudioDeviceManager "IAudioDeviceManager": Success.
//    //  * - A null pointer: Failure.
//    //  */
//    // C_agora_refptr C_IAgoraService_createAudioDeviceManagerComponent(
//    //     C_IAudioDeviceManagerObserver *observer);

//    // /**
//    //  * Creates an data channel and returns the pointer.
//    //  *
//    //  * @return
//    //  * - The pointer to \ref rtc::ILocalDataChannel "ILocalDataChannel": Success.
//    //  * - A null pointer: Failure.
//    //  */
//    // C_agora_refptr C_IAgoraService_createLocalDataChannel(const C_DataChannelConfig *config);

//    /**
// 	* @brief Get the ID of the registered extension
// 	*
// 	* @param provider_name The pointer to provider name string (null-terminated)
// 	* @param extension_name The pointer to extension name string (null-terminated)
// 	* @return
// 	*  - Pointer to the extension id string (null-terminated). The pointer will be valid during service's lifetime
// 	*/
//    const char *C_IAgoraService_getExtensionId(C_IAgoraService *this_, const char *provider_name, const char *extension_name);

//  #if defined(_WIN32) || defined(__linux__) || defined(__ANDROID__)
//    /**
// 	* @brief load the dynamic library of the extension
// 	*
// 	* @param path path of the extension library
// 	* @param unload_after_use unload the library when engine release
// 	* @return int
// 	*/
//    int C_IAgoraService_loadExtensionProvider(C_IAgoraService *this_, const char *path, bool unload_after_use);
//  #endif
//    /**
// 	* Enable extension.
// 	* If the extension is enabled, the track loads the extension automatically.
// 	*
// 	* @param provider_name name for provider, e.g. agora.io.
// 	* @param extension_name name for extension, e.g. agora.beauty.
// 	* @param track_id id for the track, OPTIONAL_NULLPTR means effective on all tracks
// 	* @param auto_enable_on_track if the extension is automatically open on track.
// 	*
// 	* @return
// 	* - 0: Success.
// 	* - < 0: Failure.
// 	*/
//    int C_IAgoraService_enableExtension(C_IAgoraService *this_,
// 									   const char *provider_name, const char *extension_name, const char *track_id,
// 									   bool auto_enable_on_track);
//    /**
// 	* Disable extension.
// 	*
// 	* @param provider_name name for provider, e.g. agora.io.
// 	* @param extension_name name for extension, e.g. agora.beauty.
// 	* @param track_id id for the track, OPTIONAL_NULLPTR means effective on all tracks
// 	*
// 	* @return
// 	* - 0: Success.
// 	* - < 0: Failure.
// 	*/
//    int C_IAgoraService_disableExtension(C_IAgoraService *this_,
// 										const char *provider_name, const char *extension_name, const char *track_id);

//    /**
// 	* Get the \ref agora::rtc::IConfigCenter "IConfigCenter" object and return the pointer.
// 	*
// 	* @return
// 	* - The pointer to \ref rtc::IConfigCenter "IConfigCenter": Success.
// 	* - A null pointer: Failure.
// 	*/
//    C_agora_refptr C_IAgoraService_getConfigCenter(C_IAgoraService *this_);

// #endregion IAgoraService

// #endregion agora::base

// #endregion agora

/** \addtogroup createAgoraService
  @{
*/
/**
 * Creates an \ref agora::base::IAgoraService "IAgoraService" object and returns the pointer.
 *
 * @return
 * - The pointer to \ref agora::base::IAgoraService "IAgoraService": Success.
 * - A null pointer: Failure.
 */
func CreateAgoraService() *IAgoraService {
	return (*IAgoraService)(C.C_createAgoraService())
}
