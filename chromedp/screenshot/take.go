package screenshot

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func FullScreen(filename string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
		if err != nil {
			return err
		}
		buf, err := page.CaptureScreenshot().
			WithQuality(100).
			WithClip(&page.Viewport{
				X:      contentSize.X,
				Y:      contentSize.Y,
				Width:  contentSize.Width,
				Height: contentSize.Height,
				Scale:  1,
			}).Do(ctx)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
			log.Panic(err)
		}

		return nil
	})
}
