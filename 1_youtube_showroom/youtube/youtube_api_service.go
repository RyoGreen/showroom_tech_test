//go:generate mockgen -source=youtube_api_service.go -destination=mock_youtube_api_service.go -package=youtube
package youtube

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeApiService interface {
	// 動画検索結果を取得する
	FetchVideos(parts []string, query, order, dataType, nextPageToken string, maxResults int64) (*youtube.SearchListResponse, error)
}

type YoutubeApiServiceImpl struct {
	service *youtube.Service
}

// YouTube Data APIのサービスを初期化し、YoutubeApiServiceImpl を生成
func NewYoutubeApiService(ctx context.Context, apiKey string) (*YoutubeApiServiceImpl, error) {
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &YoutubeApiServiceImpl{service}, nil
}

func (c *YoutubeApiServiceImpl) FetchVideos(parts []string, query, order, dataType, nextPageToken string, maxResults int64) (*youtube.SearchListResponse, error) {
	return c.service.Search.List(parts).Q(query).Order(order).Type(dataType).PageToken(nextPageToken).MaxResults(maxResults).Do()
}
