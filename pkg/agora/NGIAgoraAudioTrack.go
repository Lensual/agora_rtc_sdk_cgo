package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk-prefix/src/agora_rtc_sdk/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk-prefix/src/agora_rtc_sdk/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "C_NGIAgoraAudioTrack.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

// #ifndef C_NGI_AGORA_AUDIO_TRACK_H
// #define C_NGI_AGORA_AUDIO_TRACK_H

// #include "C_AgoraBase.h"

// #ifdef __cplusplus
// extern "C"
// {
// #endif // __cplusplus

// // // FIXME(Ender): use this class instead of AudioSendStream as local track
// #region agora

// #region agora::rtc

//   // class IAudioTrackStateObserver;
//   // class IAudioFilter;
//   // class IAudioSinkBase;
//   // class IMediaPacketReceiver;
//   // class IAudioEncodedFrameReceiver;
//   // /**
//   //  * Properties of audio frames expected by a sink.
//   //  *
//   //  * @note The source determines the audio frame to be sent to the sink based on a variety of factors,
//   //  * such as other sinks or the capability of the source.
//   //  *
//   //  */
//   // struct AudioSinkWants {
//   //   /** The sample rate of the audio frame to be sent to the sink. */
//   //   int samplesPerSec;

//   //   /** The number of audio channels of the audio frame to be sent to the sink. */
//   //   size_t channels;

//   //   AudioSinkWants() : samplesPerSec(0),
//   //                      channels(0) {}
//   //   AudioSinkWants(int sampleRate, size_t chs) : samplesPerSec(sampleRate),
//   //                                                channels(chs) {}
//   //   AudioSinkWants(int sampleRate, size_t chs, int trackNum) : samplesPerSec(sampleRate), channels(chs) {}
//   // };

//   /**
//    * The IAudioTrack class.
//    */
//   typedef void C_IAudioTrack;
//   // class IAudioTrack : public RefCountInterface {
//   //  public:

/**
 * The position of the audio filter in audio frame.
 */
type AudioFilterPosition C.enum_C_AudioFilterPosition

const (
	/**
	 * Work on the local playback
	 */
	RecordingLocalPlayback AudioFilterPosition = C.RecordingLocalPlayback
	/**
	 * Work on the post audio processing.
	 */
	PostAudioProcessing AudioFilterPosition = C.PostAudioProcessing
	/**
	 * Work on the remote audio before mixing.
	 */
	RemoteUserPlayback AudioFilterPosition = C.RemoteUserPlayback
	/**
	 * Work on the pcm source.
	 */
	PcmSource AudioFilterPosition = C.PcmSource
	/**
	 * Work on the sending branch of the pcm source.
	 */
	PcmSourceSending AudioFilterPosition = C.PcmSourceSending
	/**
	 * Work on the local playback branch of the pcm source.
	 */
	PcmSourceLocalPlayback AudioFilterPosition = C.PcmSourceLocalPlayback
	/**
	 * Work on the playback after remote-audio mix.
	 */
	RemoteMixedPlayback AudioFilterPosition = C.RemoteMixedPlayback
)

//   //  public:
//   //   /**
//   //    * Adjusts the playback volume.
//   //    * @param volume The playback volume. The value ranges between 0 and 100 (default).
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int adjustPlayoutVolume(int volume) = 0;

//   //   /**
//   //    * Gets the current playback volume.
//   //    * @param volume A pointer to the playback volume.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int getPlayoutVolume(int* volume) = 0;

//   //   /**
//   //    * Adds an audio filter.
//   //    *
//   //    * By adding an audio filter, you can apply various audio effects to the audio, for example, voice change.
//   //    * @param filter A pointer to the audio filter. See \ref agora::rtc::IAudioFilter "IAudioFilter".
//   //    * @param position The position of the audio filter. See \ref agora::rtc::IAudioTrack::AudioFilterPosition "AudioFilterPosition".
//   //    * @return
//   //    * - `true`: Success.
//   //    * - `false`: Failure.
//   //    */
//   //   virtual bool addAudioFilter(agora_refptr<IAudioFilter> filter, AudioFilterPosition position) = 0;
//   //   /**
//   //    * Removes the audio filter added by callling `addAudioFilter`.
//   //    *
//   //    * @param filter The pointer to the audio filter that you want to remove. See \ref agora::rtc::IAudioFilter "IAudioFilter".
//   //    * @param position The position of the audio filter. See #AudioFilterPosition.
//   //    * @return
//   //    * - `true`: Success.
//   //    * - `false`: Failure.
//   //    */
//   //   virtual bool removeAudioFilter(agora_refptr<IAudioFilter> filter, AudioFilterPosition position) = 0;

//   //   /**
//   //    * Enable / Disable specified audio filter
//   //    * @param id id of the filter
//   //    * @param enable enable / disable the filter with given id
//   //    * @param position The position of the audio filter. See #AudioFilterPosition.
//   //    * @return
//   //    * - 0: success
//   //    * - <0: failure
//   //    */
//   //   virtual int enableAudioFilter(const char* id, bool enable, AudioFilterPosition position) {
//   //     (void)id;
//   //     (void)enable;
//   //     (void)position;
//   //     return -1;
//   //   }

//   //   /**
//   //    * set the properties of the specified audio filter
//   //    * @param id id of the filter
//   //    * @param key key of the property
//   //    * @param jsonValue json str value of the property
//   //    * @param position The position of the audio filter. See #AudioFilterPosition.
//   //    * @return
//   //    * - 0: success
//   //    * - <0: failure
//   //    */
//   //   virtual int setFilterProperty(const char* id, const char* key, const char* jsonValue, AudioFilterPosition position) {
//   //     (void)id;
//   //     (void)key;
//   //     (void)jsonValue;
//   //     (void)position;
//   //     return -1;
//   //   }

//   //   /**
//   //    * get the properties of the specified video filter
//   //    * @param id id of the filter
//   //    * @param key key of the property
//   //    * @param jsonValue json str value of the property
//   //    * @param bufSize max length of the json value buffer
//   //    * @param position The position of the audio filter. See #AudioFilterPosition.
//   //    * @return
//   //    * - 0: success
//   //    * - <0: failure
//   //    */
//   //   virtual int getFilterProperty(const char* id, const char* key, char* jsonValue, size_t bufSize, AudioFilterPosition position) {
//   //     (void)id;
//   //     (void)key;
//   //     (void)jsonValue;
//   //     (void)bufSize;
//   //     (void)position;
//   //     return -1;
//   //   }

//   //   /**
//   //    * Gets the audio filter by its name.
//   //    *
//   //    * @param name The name of the audio filter.
//   //    * @param position The position of the audio filter. See #AudioFilterPosition.
//   //    * @return
//   //    * - The pointer to the audio filter: Success.
//   //    * - A null pointer: Failure.
//   //    */
//   //   virtual agora_refptr<IAudioFilter> getAudioFilter(const char *name, AudioFilterPosition position) const = 0;

//   //   /**
//   //    * Adds an audio sink to get PCM data from the audio track.
//   //    *
//   //    * @param sink The pointer to the audio sink. See \ref agora::rtc::IAudioSinkBase "IAudioSinkBase".
//   //    * @param wants The properties an audio frame should have when it is delivered to the sink. See \ref agora::rtc::AudioSinkWants "AudioSinkWants".
//   //    * @return
//   //    * - `true`: Success.
//   //    * - `false`: Failure.
//   //    */
//   //   virtual bool addAudioSink(agora_refptr<IAudioSinkBase> sink, const AudioSinkWants& wants) = 0;

//   //   /**
//   //    * Removes an audio sink.
//   //    *
//   //    * @param sink The pointer to the audio sink to be removed. See \ref agora::rtc::IAudioSinkBase "IAudioSinkBase".
//   //    * @return
//   //    * - `true`: Success.
//   //    * - `false`: Failure.
//   //    */
//   //   virtual bool removeAudioSink(agora_refptr<IAudioSinkBase> sink) = 0;
//   // };

//   // /**
//   //  * The observer of the local audio track.
//   //  */
//   // class ILocalAudioTrackObserver {
//   //  public:
//   //   virtual ~ILocalAudioTrackObserver() {}

//   //   /**
//   //    * Occurs when the state of a local audio track changes.
//   //    *
//   //    * @param state The state of the local audio track.
//   //    * @param errorCode The error information for a state failure: \ref agora::rtc::LOCAL_AUDIO_STREAM_ERROR "LOCAL_AUDIO_STREAM_ERROR".
//   //    */
//   //   virtual void onLocalAudioTrackStateChanged(LOCAL_AUDIO_STREAM_STATE state,
//   //                                              LOCAL_AUDIO_STREAM_ERROR errorCode) = 0;
//   // };

/**
 * `ILocalAudioTrack` is the basic class for local audio tracks, providing main methods of local audio tracks.
 *
 * You can create a local audio track by calling one of the following methods:
 * - `createLocalAudioTrack`
 * - `createCustomAudioTrack`
 * - `createMediaPlayerAudioTrack`
 * @if (!Linux)
 * You can also use the APIs in the \ref agora::rtc::INGAudioDeviceManager "IAudioDeviceManager" class if multiple recording devices are available in the system.
 * @endif
 *
 * After creating local audio tracks, you can publish one or more local audio
 * tracks by calling \ref agora::rtc::ILocalUser::publishAudio "publishAudio".
 */
type ILocalAudioTrack C.C_ILocalAudioTrack

//   typedef C_IAudioTrack C_ILocalAudioTrack;
//   // class ILocalAudioTrack : public IAudioTrack {
//   //  public:
//   //   /**
//   //    * Statistics of a local audio track.
//   //    */
//   //   struct LocalAudioTrackStats {
//   //     /**
//   //      * The source ID of the local audio track.
//   //      */
//   //     uint32_t source_id;
//   //     /**
//   //      * The number of audio frames in the buffer.
//   //      *
//   //      * When sending PCM data, the PCM data is first stored in a buffer area.
//   //      * Then a thread gets audio frames from the buffer and sends PCM data.
//   //      */
//   //     uint32_t buffered_pcm_data_list_size;
//   //     /**
//   //      * The number of audio frames missed by the thread that gets PCM data from the buffer.
//   //      */
//   //     uint32_t missed_audio_frames;
//   //     /**
//   //      * The number of audio frames sent by the thread that gets PCM data from the buffer.
//   //      */
//   //     uint32_t sent_audio_frames;
//   //     /**
//   //      * The number of audio frames sent by the user.
//   //      */
//   //     uint32_t pushed_audio_frames;
//   //     /**
//   //      * The number of dropped audio frames caused by insufficient buffer size.
//   //      */
//   //     uint32_t dropped_audio_frames;
//   //     /**
//   //      * The number of playout audio frames.
//   //      */
//   //     uint32_t playout_audio_frames;
//   //     /**
//   //      * The type of audio effect.
//   //      */
//   //     uint32_t effect_type;
//   //     /**
//   //      * Whether the hardware ear monitor is enabled.
//   //      */
//   //     uint32_t hw_ear_monitor;
//   //     /**
//   //      * Whether the local audio track is enabled.
//   //      */
//   //     bool enabled;
//   //     /**
//   //      * The volume that ranges from 0 to 255.
//   //      */
//   //     uint32_t audio_volume;  // [0,255]

//   //     LocalAudioTrackStats() : source_id(0),
//   //                              buffered_pcm_data_list_size(0),
//   //                              missed_audio_frames(0),
//   //                              sent_audio_frames(0),
//   //                              pushed_audio_frames(0),
//   //                              dropped_audio_frames(0),
//   //                              playout_audio_frames(0),
//   //                              effect_type(0),
//   //                              hw_ear_monitor(0),
//   //                              enabled(false),
//   //                              audio_volume(0) {}
//   //   };

//   //  public:
//   //   /**
//   //    * Enables or disables the local audio track.
//   //    *
//   //    * Once the local audio is enabled, the SDK allows for local audio capturing, processing, and encoding.
//   //    *
//   //    * @param enable Whether to enable the audio track:
//   //    * - `true`: Enable the local audio track.
//   //    * - `false`: Disable the local audio track.
//   //    */
//   //   virtual void setEnabled(bool enable) = 0;

//   //   /**
//   //    * Gets whether the local audio track is enabled.
//   //    * @return Whether the local audio track is enabled:
//   //    * - `true`: The local track is enabled.
//   //    * - `false`: The local track is disabled.
//   //    */
//   //   virtual bool isEnabled() const = 0;

//   //   /**
//   //    * Gets the state of the local audio.
//   //    * @return The state of the local audio: #LOCAL_AUDIO_STREAM_STATE: Success.
//   //    */
//   //   virtual LOCAL_AUDIO_STREAM_STATE getState() = 0;

//   //   /**
//   //    * Gets the statistics of the local audio track: LocalAudioTrackStats.
//   //    * @return The statistics of the local audio: LocalAudioTrackStats: Success.
//   //    */
//   //   virtual LocalAudioTrackStats GetStats() = 0;

//   //   /**
//   //    * Adjusts the audio volume for publishing.
//   //    *
//   //    * @param volume The volume for publishing. The value ranges between 0 and 100 (default).
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int adjustPublishVolume(int volume) = 0;

//   //   /**
//   //    * Gets the current volume for publishing.
//   //    * @param volume A pointer to the publishing volume.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int getPublishVolume(int* volume) = 0;

//   //   /**
//   //    * Enables or disables local playback.
//   //    * @param enable Whether to enable local playback:
//   //    * - `true`: Enable local playback.
//   //    * - `false`: Disable local playback.
//   //    * @param sync Whether to destroy local playback synchronously:
//   //    * - `true`: Destroy local playback synchronously.
//   //    * - `false`: Destroy local playback asynchronously.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int enableLocalPlayback(bool enable, bool sync = true) = 0;

//   //   /**
//   //    * Enables in-ear monitoring (for Android and iOS only).
//   //    *
//   //    * @param enabled Determines whether to enable in-ear monitoring.
//   //    * - true: Enable.
//   //    * - false: (Default) Disable.
//   //    * @param includeAudioFilters The type of the ear monitoring: #EAR_MONITORING_FILTER_TYPE
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int enableEarMonitor(bool enable, int includeAudioFilters) = 0;
//   //   /** Register an local audio track observer
//   //    *
//   //    * @param observer A pointer to the local audio track observer: \ref agora::rtc::ILocalAudioTrackObserver
//   //    * "ILocalAudioTrackObserver".
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int registerTrackObserver(ILocalAudioTrackObserver* observer) = 0;
//   //   /** Releases the local audio track observer
//   //    *
//   //    * @param observer A pointer to the local audio track observer: \ref agora::rtc::ILocalAudioTrackObserver
//   //    * "ILocalAudioTrackObserver".
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int unregisterTrackObserver(ILocalAudioTrackObserver* observer) = 0;

//   //  protected:
//   //   ~ILocalAudioTrack() {}
//   // };

/**
 * The statistics of a remote audio track.
 */
type RemoteAudioTrackStats C.struct_C_RemoteAudioTrackStats

// #region RemoteAudioTrackStats

/**
 * The user ID of the remote user who sends the audio track.
 */
func (this_ *RemoteAudioTrackStats) GetUid() uint {
	return uint(this_.uid)
}

/**
 * The user ID of the remote user who sends the audio track.
 */
func (this_ *RemoteAudioTrackStats) SetUid(uid uint) {
	this_.uid = C.C_uid_t(uid)
}

/**
 * The audio quality of the remote audio track: #QUALITY_TYPE.
 */
func (this_ *RemoteAudioTrackStats) GetQuality() int {
	return int(this_.quality)
}

/**
 * The audio quality of the remote audio track: #QUALITY_TYPE.
 */
func (this_ *RemoteAudioTrackStats) SetQuality(quality int) {
	this_.quality = C.int(quality)
}

/**
 * The network delay (ms) from the sender to the receiver.
 */
func (this_ *RemoteAudioTrackStats) GetNetworkTransportDelay() int {
	return int(this_.network_transport_delay)
}

/**
 * The network delay (ms) from the sender to the receiver.
 */
func (this_ *RemoteAudioTrackStats) SetNetworkTransportDelay(network_transport_delay int) {
	this_.network_transport_delay = C.int(network_transport_delay)
}

/**
 * The delay (ms) from the receiver to the jitter buffer.
 */
func (this_ *RemoteAudioTrackStats) GetJitterBufferDelay() uint32 {
	return uint32(this_.jitter_buffer_delay)
}

/**
 * The delay (ms) from the receiver to the jitter buffer.
 */
func (this_ *RemoteAudioTrackStats) SetJitterBufferDelay(jitter_buffer_delay uint32) {
	this_.jitter_buffer_delay = C.uint32_t(jitter_buffer_delay)
}

/**
 * The audio frame loss rate in the reported interval.
 */
func (this_ *RemoteAudioTrackStats) GetAudioLossRate() int {
	return int(this_.audio_loss_rate)
}

/**
 * The audio frame loss rate in the reported interval.
 */
func (this_ *RemoteAudioTrackStats) SetAudioLossRate(audio_loss_rate int) {
	this_.audio_loss_rate = C.int(audio_loss_rate)
}

/**
 * The number of audio channels.
 */
func (this_ *RemoteAudioTrackStats) GetNumChannels() int {
	return int(this_.num_channels)
}

/**
 * The number of audio channels.
 */
func (this_ *RemoteAudioTrackStats) SetNumChannels(num_channels int) {
	this_.num_channels = C.int(num_channels)
}

/**
 * The sample rate (Hz) of the received audio track in the reported interval.
 */
func (this_ *RemoteAudioTrackStats) GetReceivedSampleRate() int {
	return int(this_.received_sample_rate)
}

/**
 * The sample rate (Hz) of the received audio track in the reported interval.
 */
func (this_ *RemoteAudioTrackStats) SetReceivedSampleRate(received_sample_rate int) {
	this_.received_sample_rate = C.int(received_sample_rate)
}

/**
 * The average bitrate (Kbps) of the received audio track in the reported interval.
 * */
func (this_ *RemoteAudioTrackStats) GetReceivedBitrate() int {
	return int(this_.received_bitrate)
}

/**
 * The average bitrate (Kbps) of the received audio track in the reported interval.
 * */
func (this_ *RemoteAudioTrackStats) SetReceivedBitrate(received_bitrate int) {
	this_.received_bitrate = C.int(received_bitrate)
}

/**
 * The total freeze time (ms) of the remote audio track after the remote user joins the channel.
 * In a session, audio freeze occurs when the audio frame loss rate reaches 4%.
 * The total audio freeze time = The audio freeze number × 2 seconds.
 */
func (this_ *RemoteAudioTrackStats) GetTotalFrozenTime() int {
	return int(this_.total_frozen_time)
}

/**
 * The total freeze time (ms) of the remote audio track after the remote user joins the channel.
 * In a session, audio freeze occurs when the audio frame loss rate reaches 4%.
 * The total audio freeze time = The audio freeze number × 2 seconds.
 */
func (this_ *RemoteAudioTrackStats) SetTotalFrozenTime(total_frozen_time int) {
	this_.total_frozen_time = C.int(total_frozen_time)
}

/**
 * The total audio freeze time as a percentage (%) of the total time when the audio is available.
 * */
func (this_ *RemoteAudioTrackStats) GetFrozenRate() int {
	return int(this_.frozen_rate)
}

/**
 * The total audio freeze time as a percentage (%) of the total time when the audio is available.
 * */
func (this_ *RemoteAudioTrackStats) SetFrozenRate(frozen_rate int) {
	this_.frozen_rate = C.int(frozen_rate)
}

/**
 * The number of audio bytes received.
 */
func (this_ *RemoteAudioTrackStats) GetReceivedBytes() int64 {
	return int64(this_.received_bytes)
}

/**
 * The number of audio bytes received.
 */
func (this_ *RemoteAudioTrackStats) SetReceivedBytes(received_bytes int64) {
	this_.received_bytes = C.int64_t(received_bytes)
}

/**
 * The average packet waiting time (ms) in the jitter buffer.
 */
func (this_ *RemoteAudioTrackStats) GetMeanWaitingTime() int {
	return int(this_.mean_waiting_time)
}

/**
 * The average packet waiting time (ms) in the jitter buffer.
 */
func (this_ *RemoteAudioTrackStats) SetMeanWaitingTime(mean_waiting_time int) {
	this_.mean_waiting_time = C.int(mean_waiting_time)
}

/**
 * The samples of expanded speech.
 */
func (this_ *RemoteAudioTrackStats) GetExpandedSpeechSamples() uint {
	return uint(this_.expanded_speech_samples)
}

/**
 * The samples of expanded speech.
 */
func (this_ *RemoteAudioTrackStats) SetExpandedSpeechSamples(expanded_speech_samples uint) {
	this_.expanded_speech_samples = C.size_t(expanded_speech_samples)
}

/**
 * The samples of expanded noise.
 */
func (this_ *RemoteAudioTrackStats) GetExpandedNoiseSamples() uint {
	return uint(this_.expanded_noise_samples)
}

/**
 * The samples of expanded noise.
 */
func (this_ *RemoteAudioTrackStats) SetExpandedNoiseSamples(expanded_noise_samples uint) {
	this_.expanded_noise_samples = C.size_t(expanded_noise_samples)
}

/**
 * The timestamps since last report.
 */
func (this_ *RemoteAudioTrackStats) GetTimestampsSinceLastReport() uint32 {
	return uint32(this_.timestamps_since_last_report)
}

/**
 * The timestamps since last report.
 */
func (this_ *RemoteAudioTrackStats) SetTimestampsSinceLastReport(timestamps_since_last_report uint32) {
	this_.timestamps_since_last_report = C.uint32_t(timestamps_since_last_report)
}

/**
 * The minimum sequence number.
 */
func (this_ *RemoteAudioTrackStats) GetMinSequenceNumber() uint16 {
	return uint16(this_.min_sequence_number)
}

/**
 * The minimum sequence number.
 */
func (this_ *RemoteAudioTrackStats) SetMinSequenceNumber(min_sequence_number uint16) {
	this_.min_sequence_number = C.uint16_t(min_sequence_number)
}

/**
 * The maximum sequence number.
 */
func (this_ *RemoteAudioTrackStats) GetMaxSequenceNumber() uint16 {
	return uint16(this_.max_sequence_number)
}

/**
 * The maximum sequence number.
 */
func (this_ *RemoteAudioTrackStats) SetMaxSequenceNumber(max_sequence_number uint16) {
	this_.max_sequence_number = C.uint16_t(max_sequence_number)
}

/**
 * The audio energy.
 */
func (this_ *RemoteAudioTrackStats) GetAudioLevel() int32 {
	return int32(this_.audio_level)
}

/**
 * The audio energy.
 */
func (this_ *RemoteAudioTrackStats) SetAudioLevel(audio_level int32) {
	this_.audio_level = C.int32_t(audio_level)
}

/**
 * audio downlink average process time
 */
func (this_ *RemoteAudioTrackStats) GetDownlinkProcessTimeMs() uint32 {
	return uint32(this_.downlink_process_time_ms)
}

/**
 * audio downlink average process time
 */
func (this_ *RemoteAudioTrackStats) SetDownlinkProcessTimeMs(downlink_process_time_ms uint32) {
	this_.downlink_process_time_ms = C.uint32_t(downlink_process_time_ms)
}

/**
 * audio neteq loss because of expired
 */
func (this_ *RemoteAudioTrackStats) GetPacketExpiredLoss() uint32 {
	return uint32(this_.packet_expired_loss)
}

/**
 * audio neteq loss because of expired
 */
func (this_ *RemoteAudioTrackStats) SetPacketExpiredLoss(packet_expired_loss uint32) {
	this_.packet_expired_loss = C.uint32_t(packet_expired_loss)
}

/**
 * audio neteq packet arrival expired time ms
 */
func (this_ *RemoteAudioTrackStats) GetPacketMaxExpiredMs() uint32 {
	return uint32(this_.packet_max_expired_ms)
}

/**
 * audio neteq packet arrival expired time ms
 */
func (this_ *RemoteAudioTrackStats) SetPacketMaxExpiredMs(packet_max_expired_ms uint32) {
	this_.packet_max_expired_ms = C.uint32_t(packet_max_expired_ms)
}

/**
 * audio neteq jitter peak num in two second
 */
func (this_ *RemoteAudioTrackStats) GetBurstPeakNum() uint32 {
	return uint32(this_.burst_peak_num)
}

/**
 * audio neteq jitter peak num in two second
 */
func (this_ *RemoteAudioTrackStats) SetBurstPeakNum(burst_peak_num uint32) {
	this_.burst_peak_num = C.uint32_t(burst_peak_num)
}

/**
 * audio neteq jitter calc by burst opti feature
 */
func (this_ *RemoteAudioTrackStats) GetBurstJitter() uint32 {
	return uint32(this_.burst_jitter)
}

/**
 * audio neteq jitter calc by burst opti feature
 */
func (this_ *RemoteAudioTrackStats) SetBurstJitter(burst_jitter uint32) {
	this_.burst_jitter = C.uint32_t(burst_jitter)
}

/**
 * audio base target level
 */
func (this_ *RemoteAudioTrackStats) GetTargetLevelBaseMs() uint32 {
	return uint32(this_.target_level_base_ms)
}

/**
 * audio base target level
 */
func (this_ *RemoteAudioTrackStats) SetTargetLevelBaseMs(target_level_base_ms uint32) {
	this_.target_level_base_ms = C.uint32_t(target_level_base_ms)
}

/**
 * audio average target level
 */
func (this_ *RemoteAudioTrackStats) GetTargetLevelPreferedMs() uint32 {
	return uint32(this_.target_level_prefered_ms)
}

/**
 * audio average target level
 */
func (this_ *RemoteAudioTrackStats) SetTargetLevelPreferedMs(target_level_prefered_ms uint32) {
	this_.target_level_prefered_ms = C.uint32_t(target_level_prefered_ms)
}

/**
 * audio average accelerate ratio in 2s
 */
func (this_ *RemoteAudioTrackStats) GetAccelerateRate() uint16 {
	return uint16(this_.accelerate_rate)
}

/**
 * audio average accelerate ratio in 2s
 */
func (this_ *RemoteAudioTrackStats) SetAccelerateRate(accelerate_rate uint16) {
	this_.accelerate_rate = C.uint16_t(accelerate_rate)
}

/**
 * audio average preemptive expand ratio in 2s
 */
func (this_ *RemoteAudioTrackStats) GetPreemptiveExpandRate() uint16 {
	return uint16(this_.preemptive_expand_rate)
}

/**
 * audio average preemptive expand ratio in 2s
 */
func (this_ *RemoteAudioTrackStats) SetPreemptiveExpandRate(preemptive_expand_rate uint16) {
	this_.preemptive_expand_rate = C.uint16_t(preemptive_expand_rate)
}

/**
 *  The count of 80 ms frozen in 2 seconds
 */
func (this_ *RemoteAudioTrackStats) GetFrozenCount80Ms() uint16 {
	return uint16(this_.frozen_count_80_ms)
}

/**
 *  The count of 80 ms frozen in 2 seconds
 */
func (this_ *RemoteAudioTrackStats) SetFrozenCount80Ms(frozen_count_80_ms uint16) {
	this_.frozen_count_80_ms = C.uint16_t(frozen_count_80_ms)
}

/**
 *  The time of 80 ms frozen in 2 seconds
 */
func (this_ *RemoteAudioTrackStats) GetFrozenTime80Ms() uint16 {
	return uint16(this_.frozen_time_80_ms)
}

/**
 *  The time of 80 ms frozen in 2 seconds
 */
func (this_ *RemoteAudioTrackStats) SetFrozenTime80Ms(frozen_time_80_ms uint16) {
	this_.frozen_time_80_ms = C.uint16_t(frozen_time_80_ms)
}

/**
 *  The count of 200 ms frozen in 2 seconds
 */
func (this_ *RemoteAudioTrackStats) GetFrozenCount200Ms() uint16 {
	return uint16(this_.frozen_count_200_ms)
}

/**
 *  The count of 200 ms frozen in 2 seconds
 */
func (this_ *RemoteAudioTrackStats) SetFrozenCount200Ms(frozen_count_200_ms uint16) {
	this_.frozen_count_200_ms = C.uint16_t(frozen_count_200_ms)
}

/**
 *  The time of 200 ms frozen in 2 seconds
 */
func (this_ *RemoteAudioTrackStats) GetFrozenTime200Ms() uint16 {
	return uint16(this_.frozen_time_200_ms)
}

/**
 *  The time of 200 ms frozen in 2 seconds
 */
func (this_ *RemoteAudioTrackStats) SetFrozenTime200Ms(frozen_time_200_ms uint16) {
	this_.frozen_time_200_ms = C.uint16_t(frozen_time_200_ms)
}

/**
 *  The estimate delay
 */
func (this_ *RemoteAudioTrackStats) GetDelayEstimateMs() uint32 {
	return uint32(this_.delay_estimate_ms)
}

/**
 *  The estimate delay
 */
func (this_ *RemoteAudioTrackStats) SetDelayEstimateMs(delay_estimate_ms uint32) {
	this_.delay_estimate_ms = C.uint32_t(delay_estimate_ms)
}

/**
 *  The MOS value
 */
func (this_ *RemoteAudioTrackStats) GetMosValue() uint32 {
	return uint32(this_.mos_value)
}

/**
 *  The MOS value
 */
func (this_ *RemoteAudioTrackStats) SetMosValue(mos_value uint32) {
	this_.mos_value = C.uint32_t(mos_value)
}

/**
 * If the packet loss concealment (PLC) occurs for N consecutive times, freeze is considered as PLC occurring for M consecutive times.
 * freeze cnt = (n_plc - n) / m
 */
func (this_ *RemoteAudioTrackStats) GetFrozenRateByCustomPlcCount() uint32 {
	return uint32(this_.frozen_rate_by_custom_plc_count)
}

/**
 * If the packet loss concealment (PLC) occurs for N consecutive times, freeze is considered as PLC occurring for M consecutive times.
 * freeze cnt = (n_plc - n) / m
 */
func (this_ *RemoteAudioTrackStats) SetFrozenRateByCustomPlcCount(frozen_rate_by_custom_plc_count uint32) {
	this_.frozen_rate_by_custom_plc_count = C.uint32_t(frozen_rate_by_custom_plc_count)
}

/**
 * The number of audio packet loss concealment
 */
func (this_ *RemoteAudioTrackStats) GetPlcCount() uint32 {
	return uint32(this_.plc_count)
}

/**
 * The number of audio packet loss concealment
 */
func (this_ *RemoteAudioTrackStats) SetPlcCount(plc_count uint32) {
	this_.plc_count = C.uint32_t(plc_count)
}

/**
 *  Duration of inbandfec
 */
func (this_ *RemoteAudioTrackStats) GetFecDecodeMs() int32 {
	return int32(this_.fec_decode_ms)
}

/**
 *  Duration of inbandfec
 */
func (this_ *RemoteAudioTrackStats) SetFecDecodeMs(fec_decode_ms int32) {
	this_.fec_decode_ms = C.int32_t(fec_decode_ms)
}

/**
 * The total time (ms) when the remote user neither stops sending the audio
 * stream nor disables the audio module after joining the channel.
 */
func (this_ *RemoteAudioTrackStats) GetTotalActiveTime() uint64 {
	return uint64(this_.total_active_time)
}

/**
 * The total time (ms) when the remote user neither stops sending the audio
 * stream nor disables the audio module after joining the channel.
 */
func (this_ *RemoteAudioTrackStats) SetTotalActiveTime(total_active_time uint64) {
	this_.total_active_time = C.uint64_t(total_active_time)
}

/**
 * The total publish duration (ms) of the remote audio stream.
 */
func (this_ *RemoteAudioTrackStats) GetPublishDuration() uint64 {
	return uint64(this_.publish_duration)
}

/**
 * The total publish duration (ms) of the remote audio stream.
 */
func (this_ *RemoteAudioTrackStats) SetPublishDuration(publish_duration uint64) {
	this_.publish_duration = C.uint64_t(publish_duration)
}

func (this_ *RemoteAudioTrackStats) GetE2eDelayMs() int32 {
	return int32(this_.e2e_delay_ms)
}
func (this_ *RemoteAudioTrackStats) SetE2eDelayMs(e2e_delay_ms int32) {
	this_.e2e_delay_ms = C.int32_t(e2e_delay_ms)
}

/**
 * Quality of experience (QoE) of the local user when receiving a remote audio stream. See #EXPERIENCE_QUALITY_TYPE.
 */
func (this_ *RemoteAudioTrackStats) GetEoeQuality() int {
	return int(this_.qoe_quality)
}

/**
 * Quality of experience (QoE) of the local user when receiving a remote audio stream. See #EXPERIENCE_QUALITY_TYPE.
 */
func (this_ *RemoteAudioTrackStats) SetEoeQuality(qoe_quality int) {
	this_.qoe_quality = C.int(qoe_quality)
}

/**
 * The reason for poor QoE of the local user when receiving a remote audio stream. See #EXPERIENCE_POOR_REASON.
 */
func (this_ *RemoteAudioTrackStats) GetQualityChangedReason() int32 {
	return int32(this_.quality_changed_reason)
}

/**
 * The reason for poor QoE of the local user when receiving a remote audio stream. See #EXPERIENCE_POOR_REASON.
 */
func (this_ *RemoteAudioTrackStats) SetQualityChangedReason(quality_changed_reason int32) {
	this_.quality_changed_reason = C.int32_t(quality_changed_reason)
}

/**
 * The type of downlink audio effect.
 */
func (this_ *RemoteAudioTrackStats) GetDownlinkEffectType() int32 {
	return int32(this_.downlink_effect_type)
}

/**
 * The type of downlink audio effect.
 */
func (this_ *RemoteAudioTrackStats) SetDownlinkEffectType(downlink_effect_type int32) {
	this_.downlink_effect_type = C.int32_t(downlink_effect_type)
}

func NewRemoteAudioTrackStats() *RemoteAudioTrackStats {
	return (*RemoteAudioTrackStats)(C.C_RemoteAudioTrackStats_New())
}
func (this_ *RemoteAudioTrackStats) Delete() {
	C.C_RemoteAudioTrackStats_Delete((*C.struct_C_RemoteAudioTrackStats)(this_))
}

// #endregion RemoteAudioTrackStats

/**
 * The IRemoteAudioTrack class.
 */
type IRemoteAudioTrack C.C_IRemoteAudioTrack

// #region IRemoteAudioTrack

func (this_ *IRemoteAudioTrack) Delete(stats *RemoteAudioTrackStats) {
	C.C_IRemoteAudioTrack_Delete(unsafe.Pointer(this_))
}

/**
 * Gets the statistics of the remote audio track.
 * @param stats A reference to the statistics of the remote audio track: RemoteAudioTrackStats.
 * @return
 * - `true`: Success.
 * - `false`: Failure.
 */
func (this_ *IRemoteAudioTrack) GetStatistics(stats *RemoteAudioTrackStats) bool {
	return bool(C.C_IRemoteAudioTrack_getStatistics(unsafe.Pointer(this_), (*C.struct_C_RemoteAudioTrackStats)(stats)))
}

/**
 * Gets the state of the remote audio.
 * @return The state of the remote audio: #REMOTE_AUDIO_STATE.
 */
func (this_ *IRemoteAudioTrack) GetState() REMOTE_AUDIO_STATE {
	return REMOTE_AUDIO_STATE(C.C_IRemoteAudioTrack_getState(unsafe.Pointer(this_)))
}

/**
 * Registers an `IMediaPacketReceiver` object.
 *
 * You need to implement the `IMediaPacketReceiver` class in this method. Once you successfully register
 * the media packet receiver, the SDK triggers the `onMediaPacketReceived` callback when it receives an
 * audio packet.
 *
 * @param packetReceiver The pointer to the `IMediaPacketReceiver` object.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRemoteAudioTrack) RegisterMediaPacketReceiver(packetReceiver *IMediaPacketReceiver) int {
	return int(C.C_IRemoteAudioTrack_registerMediaPacketReceiver(unsafe.Pointer(this_),
		unsafe.Pointer(packetReceiver),
	))
}

/**
 * Releases the `IMediaPacketReceiver` object.
 * @param packetReceiver The pointer to the `IMediaPacketReceiver` object.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRemoteAudioTrack) UnregisterMediaPacketReceiver(packetReceiver *IMediaPacketReceiver) int {
	return int(C.C_IRemoteAudioTrack_unregisterMediaPacketReceiver(unsafe.Pointer(this_),
		unsafe.Pointer(packetReceiver),
	))
}

/**
 * Registers an `IAudioEncodedFrameReceiver` object.
 *
 * You need to implement the `IAudioEncodedFrameReceiver` class in this method. Once you successfully register
 * the media packet receiver, the SDK triggers the `onEncodedAudioFrameReceived` callback when it receives an
 * audio packet.
 *
 * @param packetReceiver The pointer to the `IAudioEncodedFrameReceiver` object.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRemoteAudioTrack) RegisterAudioEncodedFrameReceiver(packetReceiver *IAudioEncodedFrameReceiver) int {
	return int(C.C_IRemoteAudioTrack_registerAudioEncodedFrameReceiver(unsafe.Pointer(this_),
		unsafe.Pointer(packetReceiver),
	))
}

/**
 * Releases the `IAudioEncodedFrameReceiver` object.
 * @param packetReceiver The pointer to the `IAudioEncodedFrameReceiver` object.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRemoteAudioTrack) UnregisterAudioEncodedFrameReceiver(packetReceiver *IAudioEncodedFrameReceiver) int {
	return int(C.C_IRemoteAudioTrack_unregisterAudioEncodedFrameReceiver(unsafe.Pointer(this_),
		unsafe.Pointer(packetReceiver),
	))
}

/*
* Sets the sound position and gain

	@param pan The sound position of the remote user. The value ranges from -1.0 to 1.0:
	- 0.0: the remote sound comes from the front.
	- -1.0: the remote sound comes from the left.
	- 1.0: the remote sound comes from the right.
	@param gain Gain of the remote user. The value ranges from 0.0 to 100.0. The default value is 100.0 (the original gain of the remote user). The smaller the value, the less the gain.

	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *IRemoteAudioTrack) SetRemoteVoicePosition(pan float32, gain float32) int {
	return int(C.C_IRemoteAudioTrack_setRemoteVoicePosition(unsafe.Pointer(this_),
		C.float(pan),
		C.float(gain),
	))
}

/*
* Sets the volume of each audio decoded channel

	@param decoded_index The channel index of the remote user. The value ranges from 0 to 100:
	@param volume The channel index of the remote user. The value ranges from 0 to 100.
	- 0: mute the channel.
	- 100: keep the origin volume of the channel.

	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *IRemoteAudioTrack) AdjustDecodedAudioVolume(decoded_index int, volume int) int {
	return int(C.C_IRemoteAudioTrack_adjustDecodedAudioVolume(unsafe.Pointer(this_),
		C.int(decoded_index),
		C.int(volume),
	))
}

/*
* mute remote stream from timestamp

	@note
	- unmuteRemoteFromTimestamp should be called after muteRemoteFromTimestamp, othewise this stream will be muted all time

	@param timestamp The rtp timestamp of start mute
	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *IRemoteAudioTrack) MuteRemoteFromTimestamp(timestamp uint32) int {
	return int(C.C_IRemoteAudioTrack_muteRemoteFromTimestamp(unsafe.Pointer(this_),
		C.uint32_t(timestamp),
	))
}

/*
* unmute remote stream from timestamp

	@note
	- unmuteRemoteFromTimestamp should be called after muteRemoteFromTimestamp, othewise this stream will be muted all time

	@param timestamp The rtp timestamp of start unmute
	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *IRemoteAudioTrack) UnmuteRemoteFromTimestamp(timestamp uint32) int {
	return int(C.C_IRemoteAudioTrack_unmuteRemoteFromTimestamp(unsafe.Pointer(this_),
		C.uint32_t(timestamp),
	))
}

/*
* set percentage of audio acceleration during poor network

	@note
	- The relationship between this percentage and the degree of audio acceleration is non-linear and varies with different audio material.

	@param percentage The percentage of audio acceleration. The value ranges from 0 to 100. The higher the
	* percentage, the faster the acceleration. The default value is 100 (no change to the acceleration):
	- 0: disable audio acceleration.
	- > 0: enable audio acceleration.
	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *IRemoteAudioTrack) AdjustAudioAcceleration(percentage int) int {
	return int(C.C_IRemoteAudioTrack_adjustAudioAcceleration(unsafe.Pointer(this_),
		C.int(percentage),
	))
}

/*
* set percentage of audio deceleration during poor network

	@note
	- The relationship between this percentage and the degree of audio deceleration is non-linear and varies with different audio material.

	@param percentage The percentage of audio deceleration. The value ranges from 0 to 100. The higher the
	* percentage, the faster the deceleration. The default value is 100 (no change to the deceleration):
	- 0: disable audio deceleration.
	- > 0: enable audio deceleration.
	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *IRemoteAudioTrack) AdjustAudioDeceleration(percentage int) int {
	return int(C.C_IRemoteAudioTrack_adjustAudioDeceleration(unsafe.Pointer(this_),
		C.int(percentage),
	))
}

/*
* enable spatial audio

	@param enabled enable/disable spatial audio:
	- true: enable spatial audio.
	- false: disable spatial audio.
	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *IRemoteAudioTrack) EnableSpatialAudio(percentage bool) int {
	return int(C.C_IRemoteAudioTrack_enableSpatialAudio(unsafe.Pointer(this_),
		C.bool(percentage),
	))
}

/*
* Sets remote user parameters for spatial audio

	@param params spatial audio parameters

	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *IRemoteAudioTrack) SetRemoteUserSpatialAudioParams(params *SpatialAudioParams) int {
	return int(C.C_IRemoteAudioTrack_setRemoteUserSpatialAudioParams(unsafe.Pointer(this_),
		(*C.struct_C_SpatialAudioParams)(params),
	))
}

// #endregion IRemoteAudioTrack

// #endregion agora::rtc

// #endregion agora

// #ifdef __cplusplus
// }
// #endif // __cplusplus

// #endif // C_NGI_AGORA_AUDIO_TRACK_H
