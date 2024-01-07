package Script

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func ScrriptRunner(email, pwd string, steps map[string]string, Show bool) {
	// ctxx, cancel := context.WithCancel(context.Background())
	startTime := time.Now()
	defer fmt.Println("passed ", time.Since(startTime))
	// email, pwd, err := GetTheCredentials()
	// if err != nil {

	// 	log.Fatal(" Something Wrong with the Login The login.txt file !! ", err)
	// }

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", Show),
		chromedp.Flag("mute-audio", true),
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
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://app.ofppt-langues.ma/gw/api/saml/init?idp=https://sts.windows.net/dae54ad7-43df-47b7-ae86-4ac13ae567af/"),
		chromedp.WaitVisible(steps["enter_email"], chromedp.ByID),
		chromedp.SetValue(steps["enter_email"], email, chromedp.ByID),
		chromedp.Click("#idSIButton9", chromedp.ByID),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible(steps["enter_password"], chromedp.ByID),
		chromedp.SetValue(steps["enter_password"], pwd, chromedp.ByID),
		chromedp.WaitVisible(steps["sign_in"], chromedp.ByID),
		chromedp.Click(steps["sign_in"], chromedp.ByID),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible(steps["next"], chromedp.ByID),
		chromedp.Click(steps["next"], chromedp.ByID),
		chromedp.Sleep(10*time.Second),
		/// Fix: change the attr
		chromedp.WaitVisible(steps["B2"], chromedp.ByQuery),
		chromedp.Click(steps["B2"], chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible(steps["01-choice"], chromedp.ByQuery),
		chromedp.Click(steps["01-choice"], chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),

		chromedp.WaitVisible(steps["ch1-vid"], chromedp.ByQuery),
		chromedp.Click(steps["ch1-vid"], chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible(steps["play_btn"], chromedp.ByQuery),
		chromedp.Click(steps["play_btn"], chromedp.ByQuery),

		chromedp.WaitVisible(steps["after-video"], chromedp.ByQuery),
		chromedp.Click(steps["after-video"], chromedp.ByQuery),
		chromedp.Sleep(10*time.Second),
		chromedp.WaitVisible(steps["ch2-syn"], chromedp.ByQuery),
		chromedp.Click(steps["ch2-syn"], chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.Location(&res))
		fmt.Println(res)
		if err != nil {
			log.Fatal(err)
		}
		for i:=0 ;i<=34; i++{
			err = chromedp.Run(ctx,
				chromedp.WaitVisible(steps["continue"], chromedp.ByQuery),
				chromedp.Click(steps["continue"], chromedp.ByQuery),
				chromedp.Sleep(1*time.Second),
			)

			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println(" I got Here 1")
		err = chromedp.Run(ctx,
		chromedp.WaitVisible(steps["move_on"], chromedp.ByQuery),
		chromedp.Click(steps["move_on"], chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible(steps["ch2-exs1"], chromedp.ByQuery),
		chromedp.Click(steps["ch2-exs1"], chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		)
		if err != nil {
			log.Fatal(err)
		}


	fmt.Println("Current URL:", res)

}

// func GetTheCredentials() (e, p string, err error) {
// 	var email, pwd string

// 	file, err := os.Open("./login.txt")
// 	if err != nil {
// 		return "", "", err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	if scanner.Scan() {
// 		firstLine := scanner.Text()
// 		email = string(firstLine)
// 	} else {
// 		return "", "", scanner.Err()
// 	}

// 	_, err = file.Seek(0, 1)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	if scanner.Scan() {
// 		secondLine := scanner.Text()
// 		pwd = string(secondLine)
// 	} else {
// 		return "", "", scanner.Err()
// 	}

// 	return email, pwd, nil
// }

// func printElapsedTimePeriodically(ctx context.Context) {
// 	ticker := time.NewTicker(time.Hour)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			elapsed := time.Since(startTime)
// 			fmt.Printf("Script execution time: %s\n", elapsed.String())
// 		case <-ctx.Done():

// 			return
// 		}
// 	}
// }
