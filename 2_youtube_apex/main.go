package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"youtube_apex/youtube"

	"github.com/joho/godotenv"
)

func main() {
	var ctx = context.Background()

	// APIキーをenvファイルから取得
	// 本来はsecret管理ツールから取得する
	if err := godotenv.Load(); err != nil {
		slog.Error(err.Error())
		return
	}
	apiKey := os.Getenv("YOUTUBE_API_KEY")

	service, err := youtube.NewYoutubeApiService(ctx, apiKey)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	client := youtube.NewYoutubeClient(service)

	videoInfos, err := client.GetVideoInfos()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for i, info := range videoInfos {
		fmt.Printf("第%d位\n チャンネル名: %s\n 動画タイトル: %s\n 投稿日: %s\n URL: https://www.youtube.com/watch?v=%s\n\n", i+1, info.ChannelTitle, info.Title, info.PublishedAt, info.VideoID)
	}
}
