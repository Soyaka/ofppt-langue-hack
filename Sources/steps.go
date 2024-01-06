
package sources

var Steps = map[string]string{
	"enter_email":    "#i0116",
	"next_password":  "#idSIButton9",
	"enter_password": "#i0118",
	"sign_in":        "#idSIButton9",
	"cancel":         "#cancelLink",
	"other_account":  "#otherTileText",
	"next":           "#idSIButton9",
	"submit_again":   "#idSIButton9",
	"yes":            "#idSIButton9",
	"select_lesson":  ".lesson-card-container-is-active > div:nth-child(1)",
	"select_video":   ".lesson-menu-activities-list > li:nth-child(1) > altissia-activity-overview-card:nth-child(1) > a:nth-child(1)",
	"video_element":  ".media-player-container",
	"mute":           "#app-main-content > altissia-app-container > div > main > altissia-app-video-activity > div > div > altissia-media-player > div > div > plyr > div > div.plyr__controls > div.plyr__controls__item.plyr__volume > button",
	"video_id":       ".plyr__video-wrapper > video:nth-child(1)",
	"play_btn":       "button.plyr__controls__item:nth-child(1)",
	"video_dur":      "div.plyr__controls__item:nth-child(3)",
	"continue":       ".center-content",
}
