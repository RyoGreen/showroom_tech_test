package youtube

import (
	"strings"
)

type YoutubeClient struct {
	service YoutubeApiService
}

func NewYoutubeClient(service YoutubeApiService) *YoutubeClient {
	return &YoutubeClient{service}
}

// YouTube Data APIの Search Listで使用するパラメーター
const (
	keyword              = "SHOWROOM" // 検索キーワード
	order                = "date"     // 取得順序(日付順)
	dataType             = "video"    // 取得するデータの種類
	maxResultsPerRequest = 50         // 1回のリクエストで取得する最大数
)

// 取得する動画IDの最大数
const maxVideoIDResults = 100

// 動画情報を取得する際に指定するpartパラメーター
var parts = []string{"snippet", "id"}

// 動画IDのリストを取得する
func (c *YoutubeClient) GetVideoIDs() ([]string, error) {
	var (
		videoIDs      []string
		nextPageToken string
	)

	// 100件取得するまで繰り返す
	for len(videoIDs) < maxVideoIDResults {
		response, err := c.service.FetchVideos(parts, keyword, order, dataType, nextPageToken, maxResultsPerRequest)
		if err != nil {
			return nil, err
		}
		for _, item := range response.Items {
			// タイトルに"SHOWROOM"が含まれている動画のみ追加する
			if strings.Contains(item.Snippet.Title, keyword) {
				videoIDs = append(videoIDs, item.Id.VideoId)
			}
			// 100件取得したらループを抜けて終了する
			if len(videoIDs) >= maxVideoIDResults {
				break
			}
		}

		nextPageToken = response.NextPageToken
		// 次のページがない場合はループを終了する
		if nextPageToken == "" {
			break
		}
	}

	return videoIDs, nil
}
