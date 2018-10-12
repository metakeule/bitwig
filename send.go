package bitwig

import (
	"fmt"

	"github.com/metakeule/osc"
)

func PreRoll0() (msg osc.Message) {
	msg.Address = "/preroll"
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func PreRoll1() (msg osc.Message) {
	msg.Address = "/preroll"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func PreRoll2() (msg osc.Message) {
	msg.Address = "/preroll"
	msg.Arguments = []osc.Argument{osc.Int(2)}
	return
}

func PreRoll4() (msg osc.Message) {
	msg.Address = "/preroll"
	msg.Arguments = []osc.Argument{osc.Int(4)}
	return
}

func Undo() (msg osc.Message) {
	msg.Address = "/undo"
	msg.Arguments = []osc.Argument{}
	return
}

func Redo() (msg osc.Message) {
	msg.Address = "/redo"
	msg.Arguments = []osc.Argument{}
	return
}

func TempoRaw(bpm uint32) (msg osc.Message) {
	msg.Address = "/tempo/raw"
	msg.Arguments = []osc.Argument{osc.Int(bpm)}
	return
}

func TempoTap() (msg osc.Message) {
	msg.Address = "/tempo/tap"
	msg.Arguments = []osc.Argument{}
	return
}

func PositionNext() (msg osc.Message) {
	msg.Address = "/position/+"
	msg.Arguments = []osc.Argument{}
	return
}

func PositionPrevious() (msg osc.Message) {
	msg.Address = "/position/-"
	msg.Arguments = []osc.Argument{}
	return
}

func PositionFastForward() (msg osc.Message) {
	msg.Address = "/position/++"
	msg.Arguments = []osc.Argument{}
	return
}

func PositionRewind() (msg osc.Message) {
	msg.Address = "/position/--"
	msg.Arguments = []osc.Argument{}
	return
}

func ProjectNext() (msg osc.Message) {
	msg.Address = "/project/+"
	msg.Arguments = []osc.Argument{}
	return
}

func ProjectPrevious() (msg osc.Message) {
	msg.Address = "/project/-"
	msg.Arguments = []osc.Argument{}
	return
}

func ProjectEngineOn() (msg osc.Message) {
	msg.Address = "/project/engine"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func ProjectEngineOff() (msg osc.Message) {
	msg.Address = "/project/engine"
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func Click() (msg osc.Message) {
	msg.Address = "/click"
	msg.Arguments = []osc.Argument{}
	return
}

func Stop() (msg osc.Message) {

	msg.Address = "/stop"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func Play() (msg osc.Message) {

	msg.Address = "/play"
	//msg.Arguments = []osc.Argument{osc.Int(1)}
	msg.Arguments = []osc.Argument{}

	return
}

func Restart() (msg osc.Message) {
	msg.Address = "/restart"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func Repeat() (msg osc.Message) {
	msg.Address = "/repeat"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func PunchIn() (msg osc.Message) {
	msg.Address = "/punchIn"
	msg.Arguments = []osc.Argument{}
	return
}

func PunchOut() (msg osc.Message) {
	msg.Address = "/punchOut"
	msg.Arguments = []osc.Argument{}
	return
}

func Record() (msg osc.Message) {
	msg.Address = "/record"
	msg.Arguments = []osc.Argument{}
	return
}

func Overdub() (msg osc.Message) {
	msg.Address = "/overdub"
	msg.Arguments = []osc.Argument{}
	return
}

func OverdubLauncher() (msg osc.Message) {
	msg.Address = "/overdub/launcher"
	msg.Arguments = []osc.Argument{}
	return
}

func Crossfade(n int8) (msg osc.Message) {
	msg.Address = "/crossfade"
	msg.Arguments = []osc.Argument{osc.Int(n)}
	return
}

func Autowrite() (msg osc.Message) {
	msg.Address = "/autowrite"
	msg.Arguments = []osc.Argument{}
	return
}

func AutomationWriteModeLatch(s string) (msg osc.Message) {
	msg.Address = "/automationWriteMode"
	msg.Arguments = []osc.Argument{osc.String("latch")}
	return
}

func AutomationWriteModeTouch(s string) (msg osc.Message) {
	msg.Address = "/automationWriteMode"
	msg.Arguments = []osc.Argument{osc.String("touch")}
	return
}

func AutomationWriteModeWrite(s string) (msg osc.Message) {
	msg.Address = "/automationWriteMode"
	msg.Arguments = []osc.Argument{osc.String("write")}
	return
}

func LayoutArrange(s string) (msg osc.Message) {
	msg.Address = "/layout/arrange"
	msg.Arguments = []osc.Argument{}
	return
}

func LayoutMix(s string) (msg osc.Message) {
	msg.Address = "/layout/mix"
	msg.Arguments = []osc.Argument{}
	return
}

func LayoutEdit(s string) (msg osc.Message) {
	msg.Address = "/layout/edit"
	msg.Arguments = []osc.Argument{}
	return
}

func PanelNoteEditor(s string) (msg osc.Message) {
	msg.Address = "/panel/noteEditor"
	msg.Arguments = []osc.Argument{}
	return
}

func PanelAutomationEditor(s string) (msg osc.Message) {
	msg.Address = "/panel/automationEditor"
	msg.Arguments = []osc.Argument{}
	return
}

func PanelDevices(s string) (msg osc.Message) {
	msg.Address = "/panel/devices"
	msg.Arguments = []osc.Argument{}
	return
}

func PanelMixer(s string) (msg osc.Message) {
	msg.Address = "/panel/mixer"
	msg.Arguments = []osc.Argument{}
	return
}

func PanelFullscreen(s string) (msg osc.Message) {
	msg.Address = "/panel/fullscreen"
	msg.Arguments = []osc.Argument{}
	return
}

func ArrangerCueMarkerVisibility(s string) (msg osc.Message) {
	msg.Address = "/arranger/cueMarkerVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func ArrangerPlaybackFollow(s string) (msg osc.Message) {
	msg.Address = "/arranger/playbackFollow"
	msg.Arguments = []osc.Argument{}
	return
}

func ArrangerTrackRowHeight(s string) (msg osc.Message) {
	msg.Address = "/arranger/trackRowHeight"
	msg.Arguments = []osc.Argument{}
	return
}

func ArrangerClipLauncherSectionVisibility(s string) (msg osc.Message) {
	msg.Address = "/arranger/clipLauncherSectionVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func ArrangerTimeLineVisibility(s string) (msg osc.Message) {
	msg.Address = "/arranger/timeLineVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func ArrangerIOSectionVisibility(s string) (msg osc.Message) {
	msg.Address = "/arranger/ioSectionVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func ArrangerEffectTracksVisibility(s string) (msg osc.Message) {
	msg.Address = "/arranger/effectTracksVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func MixerClipLauncherSectionVisibility(s string) (msg osc.Message) {
	msg.Address = "/mixer/clipLauncherSectionVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func MixerCrossFadeSectionVisibility(s string) (msg osc.Message) {
	msg.Address = "/mixer/crossFadeSectionVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func MixerDeviceSectionVisibility(s string) (msg osc.Message) {
	msg.Address = "/mixer/deviceSectionVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func MixerSendsSectionVisibility(s string) (msg osc.Message) {
	msg.Address = "/mixer/sendsSectionVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func MixerIOSectionVisibility(s string) (msg osc.Message) {
	msg.Address = "/mixer/ioSectionVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func MixerMeterSectionVisibility(s string) (msg osc.Message) {
	msg.Address = "/mixer/meterSectionVisibility"
	msg.Arguments = []osc.Argument{}
	return
}

func TrackBankNext() (msg osc.Message) {
	msg.Address = "/track/bank/+"
	msg.Arguments = []osc.Argument{}
	return
}

func TrackBankPrevious() (msg osc.Message) {
	msg.Address = "/track/bank/-"
	msg.Arguments = []osc.Argument{}
	return
}

// Page = 8
func TrackBankPageNext() (msg osc.Message) {
	msg.Address = "/track/bank/page/+"
	msg.Arguments = []osc.Argument{}
	return
}

func TrackBankPagePrevious() (msg osc.Message) {
	msg.Address = "/track/bank/page/-"
	msg.Arguments = []osc.Argument{}
	return
}

// track is 1-8; red, green, blue between 0 and 1
func TrackColor(track uint8, red, green, blue float32) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d", track)
	msg.Arguments = []osc.Argument{osc.Float(red), osc.Float(green), osc.Float(blue)}
	return
}

func TrackNext() (msg osc.Message) {
	msg.Address = "/track/+"
	msg.Arguments = []osc.Argument{}
	return
}

func TrackPrevious() (msg osc.Message) {
	msg.Address = "/track/-"
	msg.Arguments = []osc.Argument{}
	return
}

// Enable VU-Meter notifications
func TrackVUOn() (msg osc.Message) {
	msg.Address = "/track/vu"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

// Disable VU-Meter notifications
func TrackVUOff() (msg osc.Message) {
	msg.Address = "/track/vu"
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

// Toggles between the Audio/Instrument and Effect track bank
func TrackToggleBank() (msg osc.Message) {
	msg.Address = "/track/toggleBank"
	msg.Arguments = []osc.Argument{}
	return
}

func TrackAddAudio() (msg osc.Message) {
	msg.Address = "/track/add/audio"
	msg.Arguments = []osc.Argument{}
	return
}

func TrackAddEffect() (msg osc.Message) {
	msg.Address = "/track/add/effect"
	msg.Arguments = []osc.Argument{}
	return
}

func TrackAddInstrument() (msg osc.Message) {
	msg.Address = "/track/add/instrument"
	msg.Arguments = []osc.Argument{}
	return
}

func TrackActivatedOn(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/activated", trackno)
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackActivatedOff(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/activated", trackno)
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func TrackSelect(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/select", trackno)
	msg.Arguments = []osc.Argument{}
	return
}
func TrackVolume(trackno uint8, vol int8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/volume", trackno)
	msg.Arguments = []osc.Argument{osc.Int(vol)}
	return
}
func TrackVolumeIndicateOn(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/volume/indicate", trackno)
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackVolumeIndicateOff(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/volume/indicate", trackno)
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func TrackPan(trackno uint8, pan int8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/pan", trackno)
	msg.Arguments = []osc.Argument{osc.Int(pan)}
	return
}

func TrackPanIndicateOn(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/pan/indicate", trackno)
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackPanIndicateOff(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/pan/indicate", trackno)
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func TrackMute(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/mute", trackno)
	msg.Arguments = []osc.Argument{}
	return
}

func TrackSolo(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/solo", trackno)
	msg.Arguments = []osc.Argument{}
	return
}
func TrackRecArm(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/recarm", trackno)
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackRecUnarm(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/recarm", trackno)
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func TrackMonitor(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/monitor", trackno)
	msg.Arguments = []osc.Argument{}
	return
}
func TrackMonitorAuto(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/monitor/auto", trackno)
	msg.Arguments = []osc.Argument{}
	return
}
func TrackcCrossfadeModeA(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/crossfadeMode/A", trackno)
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackcCrossfadeModeB(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/crossfadeMode/B", trackno)
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackcCrossfadeModeAB(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/crossfadeMode/AB", trackno)
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackSendVolume(trackno, send uint8, vol int8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/send/%d/volume", trackno, send)
	msg.Arguments = []osc.Argument{osc.Int(vol)}
	return
}

func TrackSendVolumeIndicateOn(trackno uint8, send uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/send/%d/volume/indicate", trackno, send)
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackSendVolumeIndicateOff(trackno uint8, send uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/send/%d/volume/indicate", trackno, send)
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func TrackClipLaunch(track, scene uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%v/clip/%v/launch", track, scene)
	msg.Arguments = []osc.Argument{}
	return
}

func TrackClipRecord(track, scene uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%v/clip/%v/record", track, scene)
	msg.Arguments = []osc.Argument{}
	return
}

func TrackClipSelect(track, scene uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%v/clip/%v/select", track, scene)
	msg.Arguments = []osc.Argument{}
	return
}

func TrackClipStop(track uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%v/clip/stop", track)
	msg.Arguments = []osc.Argument{}
	return
}

func TrackClipReturnToArrangement(track uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%v/clip/returntoarrangement", track)
	msg.Arguments = []osc.Argument{}
	return
}

func TrackStop() (msg osc.Message) {
	msg.Address = "/track/stop"
	return
}

func TrackIndicateVolumeOn() (msg osc.Message) {
	msg.Address = "/track/indicate/volume"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackIndicateVolumeOff() (msg osc.Message) {
	msg.Address = "/track/indicate/volume"
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func TrackIndicatePanOn() (msg osc.Message) {
	msg.Address = "/track/indicate/pan"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackIndicatePanOff() (msg osc.Message) {
	msg.Address = "/track/indicate/pan"
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func TrackIndicateSendOn(send uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/indicate/send/%d", send)
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func TrackIndicateSendOff(send uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/indicate/send/%d", send)
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func MasterIndicateVolumeOn() (msg osc.Message) {
	msg.Address = "/master/indicate/volume"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func MasterIndicateVolumeOff() (msg osc.Message) {
	msg.Address = "/master/indicate/volume"
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func MasterIndicatePanOn() (msg osc.Message) {
	msg.Address = "/master/indicate/pan"
	msg.Arguments = []osc.Argument{osc.Int(1)}
	return
}

func MasterIndicatePanOff() (msg osc.Message) {
	msg.Address = "/master/indicate/pan"
	msg.Arguments = []osc.Argument{osc.Int(0)}
	return
}

func SceneLaunch(num uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/scene/%v/launch", num)
	return
}

func SceneNext() (msg osc.Message) {
	msg.Address = "/scene/+"
	msg.Arguments = []osc.Argument{}
	return
}

func ScenePrevious() (msg osc.Message) {
	msg.Address = "/scene/-"
	msg.Arguments = []osc.Argument{}
	return
}

// Bank = 8
func SceneBankNext() (msg osc.Message) {
	msg.Address = "/scene/bank/+"
	msg.Arguments = []osc.Argument{}
	return
}

func SceneBankPrevious() (msg osc.Message) {
	msg.Address = "/scene/bank/-"
	msg.Arguments = []osc.Argument{}
	return
}

// Group actions
func TrackEnter(trackno uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/track/%d/enter", trackno)
	msg.Arguments = []osc.Argument{}
	return
}

func TrackParent() (msg osc.Message) {
	msg.Address = "/track/parent"
	msg.Arguments = []osc.Argument{}
	return
}

func VkbMIDINote(channel uint8, note int8, velocity int8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/vkb_midi/%d/note/%d", channel, note)
	msg.Arguments = []osc.Argument{osc.Int(velocity)}
	return
}

func VkbMIDIOctaveUp(channel uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/vkb_midi/%d/note/+", channel)
	msg.Arguments = []osc.Argument{}
	return
}

func VkbMIDIOctaveDown(channel uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/vkb_midi/%d/note/-", channel)
	msg.Arguments = []osc.Argument{}
	return
}

func VkbMIDIDrumNote(channel uint8, note int8, velocity int8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/vkb_midi/%d/drum/%d", channel, note)
	msg.Arguments = []osc.Argument{osc.Int(velocity)}
	return
}

func VkbMIDIDrumOctaveUp(channel uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/vkb_midi/%d/drum/+", channel)
	msg.Arguments = []osc.Argument{}
	return
}

func VkbMIDIDrumOctaveDown(channel uint8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/vkb_midi/%d/drum/-", channel)
	msg.Arguments = []osc.Argument{}
	return
}

func VkbMIDIControlChange(channel uint8, cc int8, val int8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/vkb_midi/%d/cc/%d", channel, cc)
	msg.Arguments = []osc.Argument{osc.Int(val)}
	return
}

func VkbMIDIAfterTouch(channel uint8, note int8, pressure int8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/vkb_midi/%d/aftertouch/%d", channel, note)
	msg.Arguments = []osc.Argument{osc.Int(pressure)}
	return
}

// (No-Bend:64)
func VkbMIDIPitchBend(channel uint8, pitch int8) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/vkb_midi/%d/pitchbend", channel)
	msg.Arguments = []osc.Argument{osc.Int(pitch)}
	return
}

// val = 0 -> no fixed velocity, otherwise: set velocity to the val
func VkbMIDIFixedVelocity(val uint8) (msg osc.Message) {
	msg.Address = "/vkb_midi/velocity"
	msg.Arguments = []osc.Argument{osc.Int(val)}
	return
}

// Executes one action from the action list (e.g. Save). Replace spaces in action-ids with a dash (e.g. Select-All).
func Action(action string) (msg osc.Message) {
	msg.Address = fmt.Sprintf("/action/%s", action)
	msg.Arguments = []osc.Argument{}
	return
}

// Flushes all values to the clients
func Refresh() (msg osc.Message) {
	msg.Address = "/refresh"
	msg.Arguments = []osc.Argument{}
	return
}

/*
 Browser Actions
• /browser/preset           Activates the browser to browse for presets of the currently selected device
• /browser/device           Activates the browser to insert a device after the currently selected device
• /browser/device/after (same as /browser/device)
• /browser/device/before    Activates the browser to insert a device before the currently selected device
• /browser/commit     Commits the current selection in the browser
• /browser/cancel     Cancels the current browser session
• /browser/filter/{1-6}/{+,-}   The columns are as follows: 1: Location, 2: Favorites, 3: Creator, 4: Tags, 5: Devices, 6: Category
• /browser/preset/{+,-}
*/

type browser struct{}

var Browser browser

// Activates the browser to browse for presets of the currently selected device
func (browser) Preset() (msg osc.Message) {
	msg.Address = "/browser/preset"
	msg.Arguments = []osc.Argument{}
	return
}

// Activates the browser to insert a device after the currently selected device
func (browser) DeviceAfter() (msg osc.Message) {
	msg.Address = "/browser/device"
	msg.Arguments = []osc.Argument{}
	return
}

// Activates the browser to insert a device before the currently selected device
func (browser) DeviceBefore() (msg osc.Message) {
	msg.Address = "/browser/device/before"
	msg.Arguments = []osc.Argument{}
	return
}

// Commits the current selection in the browser
func (browser) Commit() (msg osc.Message) {
	msg.Address = "/browser/commit"
	msg.Arguments = []osc.Argument{}
	return
}

// Cancels the current browser session
func (browser) Cancel() (msg osc.Message) {
	msg.Address = "/browser/cancel"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterLocationNext() (msg osc.Message) {
	msg.Address = "/browser/filter/1/+"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterLocationPrevious() (msg osc.Message) {
	msg.Address = "/browser/filter/1/-"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterFavoritesNext() (msg osc.Message) {
	msg.Address = "/browser/filter/2/+"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterFavoritesPrevious() (msg osc.Message) {
	msg.Address = "/browser/filter/2/-"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterCreatorNext() (msg osc.Message) {
	msg.Address = "/browser/filter/3/+"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterCreatorPrevious() (msg osc.Message) {
	msg.Address = "/browser/filter/3/-"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterTagsNext() (msg osc.Message) {
	msg.Address = "/browser/filter/4/+"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterTagsPrevious() (msg osc.Message) {
	msg.Address = "/browser/filter/4/-"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterDevicesNext() (msg osc.Message) {
	msg.Address = "/browser/filter/5/+"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterDevicesPrevious() (msg osc.Message) {
	msg.Address = "/browser/filter/5/-"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterCategoryNext() (msg osc.Message) {
	msg.Address = "/browser/filter/6/+"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) FilterCategoryPrevious() (msg osc.Message) {
	msg.Address = "/browser/filter/6/-"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) PresetNext() (msg osc.Message) {
	msg.Address = "/browser/preset/+"
	msg.Arguments = []osc.Argument{}
	return
}

func (browser) PresetPrevious() (msg osc.Message) {
	msg.Address = "/browser/preset/-"
	msg.Arguments = []osc.Argument{}
	return
}

/*
not implemented yet:

 Cursor Device / Primary Device
• /device/{+,-}

• /device/window        Displays the window for VST plugins

• /device/bypass

• /device/param/{+,-}

• /device/param/{1-8}/value {0-127}

• /device/param/{1-8}/indicate {0,1}

• /device/indicate/param {0,1}

• /device/layer/{1-8}/selected

• /device/layer/{1-8}/volume {0-127}

• /device/layer/{1-8}/pan {0-127}

• /device/layer/{1-8}/mute {1,0,-}

• /device/layer/{1-8}/solo {1,0,-}

• /device/layer/{1-8}/send/{1-8}/volume {0-127}

• /device/layer/{1-8}/enter

• /device/layer/parent

• /device/layer/{+,-}

• /device/layer/page/{+,-}

• Same commands apply for the primary device but use /primary/... instead of /device/...

*/
