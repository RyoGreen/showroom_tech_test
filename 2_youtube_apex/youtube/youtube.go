package youtube

import (
	"log/slog"
	"os"
	"time"

	"google.golang.org/api/youtube/v3"
)

var jst *time.Location

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

type YoutubeClient struct {
	service YoutubeApiService
}

func NewYoutubeClient(service YoutubeApiService) *YoutubeClient {
	return &YoutubeClient{service}
}

// 動画情報の構造体
type VideoInfo struct {
	VideoID      string
	Title        string
	ChannelTitle string
	PublishedAt  string
}

// YouTube Data APIのSearch Listで使用するパラメーター
const (
	query                = "Apex Legends" // 検索キーワード
	order                = "viewCount"    // 取得順序(視聴回数順)
	dataType             = "video"        // 取得するデータの種類
	maxResultsPerRequest = 20             // 1回のリクエストで取得するデータの最大数
)

// 動画情報を取得する際に指定するpartパラメーター
var videoParts = []string{"snippet", "id"}

// 取得する動画情報の最大数
const maxVideoInfoResults = 10

// 指定された条件に基づいて動画情報を取得する
func (c *YoutubeClient) GetVideoInfos() ([]*VideoInfo, error) {
	var (
		videoInfos    []*VideoInfo
		nextPageToken string
	)

	// 現在時刻から3日前の日付を計算する
	now := time.Now()
	publishedAfter := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, jst).AddDate(0, 0, -3).Format(time.RFC3339)

	// 10件取得するまでループする
	for len(videoInfos) < maxVideoInfoResults {
		response, err := c.service.FetchVideos(videoParts, query, order, dataType, publishedAfter, nextPageToken, maxResultsPerRequest)
		if err != nil {
			return nil, err
		}

		for _, item := range response.Items {
			videoInfo, isValid, err := c.filterValidVideo(item)
			if err != nil {
				return nil, err
			}
			if isValid {
				videoInfos = append(videoInfos, videoInfo)
			}
			// 10件取得した場合、ループを抜ける
			if len(videoInfos) >= maxVideoInfoResults {
				break
			}
		}

		nextPageToken = response.NextPageToken
		// 次のページがない場合、ループを抜ける
		if nextPageToken == "" {
			break
		}
	}

	return videoInfos, nil
}

// 日本のチャンネルでない場合や動画の公開日（PublishedAt）が不正な値の場合は無効と判断する
// 有効な場合は動画情報も返す
func (c *YoutubeClient) filterValidVideo(item *youtube.SearchResult) (videoInfo *VideoInfo, isValid bool, err error) {
	publishedAtUTC, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
	// 不正な値の場合はログを出力し、falseを返す
	if err != nil {
		slog.Warn("Failed to parse PublishedAt", "error", err)
		return nil, false, nil
	}
	isJapaneseChannel, err := c.isJapaneseChannel(item.Snippet.ChannelId)
	if err != nil {
		return nil, false, err
	}

	// 日本人のチャンネルでない場合は、falseを返す
	if !isJapaneseChannel {
		return nil, false, nil
	}

	return &VideoInfo{
		VideoID:      item.Id.VideoId,
		Title:        item.Snippet.Title,
		ChannelTitle: item.Snippet.ChannelTitle,
		PublishedAt:  publishedAtUTC.In(jst).Format("2006-01-02 15:04:05"),
	}, true, nil
}

// YouTube Data APIの Channel Listで使用するパラメーター
// チャンネル情報を取得する際に指定するpartパラメーター
var channelParts = []string{"snippet"}

// 国コード
const countryName = "JP"

// 指定されたチャンネルIDが日本人のチャンネルかを判定する
func (c *YoutubeClient) isJapaneseChannel(channelID string) (bool, error) {
	channelResponse, err := c.service.FetchChannelInfo(channelParts, channelID)
	if err != nil {
		return false, err
	}
	// チャンネルのItemが取得できなかった場合はfalseを返す
	if len(channelResponse.Items) == 0 {
		slog.Info("channel not found", "channelID", channelID)
		return false, nil
	}
	// チャンネルの居住地が日本(JP)であればtrueとする
	if channelResponse.Items[0].Snippet.Country == countryName {
		return true, nil
	}

	return false, nil
}
