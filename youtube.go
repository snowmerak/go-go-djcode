package main

import (
	"github.com/kkdai/youtube/v2"
)

var youtubeClient = new(youtube.Client)

func GetStream(client *youtube.Client, url string) (string, error) {
	video, err := client.GetVideo(url)
	if err != nil {
		return "", err
	}
	stream, err := client.GetStreamURL(video, &video.Formats.Type("audio")[0])
	if err != nil {
		return "", err
	}
	return stream, nil
}
