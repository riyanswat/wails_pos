// (Almost) all the code in this file has been copied from:
// https://github.com/kkdai/youtube/blob/master/example_test.go
package backend

import (
	// "fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func Downloader(videoID string) {

	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	videoID := "JzzjOXQ-G8I"
// 	Downloader(videoID)
// }
