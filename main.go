package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"bufio"

	"github.com/chromedp/chromedp"
)

func main() {
	email, pwd := GetTheCredentials()
	// Create a logger to output the logs to a file
	logFile, err := os.Create("chromedp_log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// logger := log.New(logFile, "", log.LstdFlags)

	// Enable logging for chromedp and run in non-headless mode
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("start-maximized", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("no-sandbox", true),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var res string
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://app.ofppt-langues.ma/gw/api/saml/init?idp=https://sts.windows.net/dae54ad7-43df-47b7-ae86-4ac13ae567af/"),
		chromedp.WaitVisible(steps["enter_email"], chromedp.ByID),
		chromedp.SetValue(steps["enter_email"], email, chromedp.ByID),
		chromedp.Click("#idSIButton9", chromedp.ByID),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible(steps["enter_password"], chromedp.ByID),
		chromedp.SetValue(steps["enter_password"], pwd, chromedp.ByID),
		chromedp.Click(steps["sign_in"], chromedp.ByID),
		chromedp.Sleep(5*time.Second), // Adjust as needed
		chromedp.Location(&res),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current URL:", res)
}

func GetTheCredentials() (email, pwd string) {

	file, err := os.Open("./login.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		firstLine := scanner.Text()
		email = firstLine
	} else {
		return
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error rewinding file:", err)
		return
	}

	if scanner.Scan() {
		secondLine := scanner.Text()
		fmt.Println("Second Line:", secondLine)
		pwd = secondLine
	} else {
		fmt.Println("Error reading the second line:", scanner.Err())
		return
	}
	return email, pwd
}

var steps = map[string]string{
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
