package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"youtube_showroom/youtube"

	"github.com/joho/godotenv"
)

func main() {
	var ctx = context.Background()

	// APIキーをenvファイルから取得
	// 本来はsecret管理ツールから取得する
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		slog.Error("YOUTUBE_API_KEY is not set")
		return
	}
	service, err := youtube.NewYoutubeApiService(ctx, apiKey)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	client := youtube.NewYoutubeClient(service)

	videoIDs, err := client.GetVideoIDs()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for i, videoID := range videoIDs {
		fmt.Printf("No.%v https://www.youtube.com/watch?v=%s\n", i+1, videoID)
	}
}
