package youtube

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
)

func TestGetVideoInfos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockYoutubeApiService(ctrl)

	c := NewYoutubeClient(mockService)

	t.Run("正常系: 'SHOWROOM'がタイトルに含まれる最大100件の動画IDを取得する", func(t *testing.T) {
		expected := []string{}
		for i := 1; i <= 100; i++ {
			expected = append(expected, fmt.Sprintf("video%d", i))
		}
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "SHOWROOM", "date", "video", gomock.Any(), int64(50)).Return(&youtube.SearchListResponse{Items: generateMockSearchResult(50, 1, true), NextPageToken: "next_token"}, nil)
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "SHOWROOM", "date", "video", "next_token", int64(50)).Return(&youtube.SearchListResponse{Items: generateMockSearchResult(50, 51, true), NextPageToken: ""}, nil)
		got, err := c.GetVideoIDs()
		assert.NoError(t, err)
		assert.Len(t, got, 100)
		assert.Equal(t, expected, got)
	})
	t.Run("正常系: 'SHOWROOM'がタイトルに含まれる動画IDを100件取得後、nextPageToken があっても追加取得しない", func(t *testing.T) {
		expected := []string{}
		for i := 1; i <= 100; i++ {
			expected = append(expected, fmt.Sprintf("video%d", i))
		}

		mockService.EXPECT().FetchVideos(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), "", int64(50)).Return(&youtube.SearchListResponse{Items: generateMockSearchResult(50, 1, true), NextPageToken: "next_token"}, nil)
		mockService.EXPECT().FetchVideos(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), "next_token", int64(50)).Return(&youtube.SearchListResponse{Items: generateMockSearchResult(50, 51, true), NextPageToken: "another_token"}, nil)
		got, err := c.GetVideoIDs()
		assert.NoError(t, err)
		assert.Len(t, got, 100)
		assert.Equal(t, expected, got)
	})
	t.Run("正常系: 30件, 50件, 50件の順で'SHOWROOM'がタイトルに含まれるデータが返された場合でも、100件で打ち切る", func(t *testing.T) {
		expected := []string{}
		// 最初の50件、'SHOWROOM'キーワードが含まれているのは最初の30件
		// そのため、expectedにはvideo1からvideo30までのIDを追加
		for i := 1; i <= 30; i++ {
			expected = append(expected, fmt.Sprintf("video%d", i))
		}

		// 次の50件、すべてのタイトルに'SHOWROOM'が含まれているので、video51からvideo100まで追加
		for i := 51; i <= 80; i++ {
			expected = append(expected, fmt.Sprintf("video%d", i))
		}

		// 最後の50件、すべてのタイトルに'SHOWROOM'が含まれているので、video101からvideo120まで追加
		for i := 81; i <= 120; i++ {
			expected = append(expected, fmt.Sprintf("video%d", i))
		}

		// 最初の50件、'SHOWROOM'キーワードが含まれているのは30件
		mockService.EXPECT().FetchVideos(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), "", int64(50)).Return(&youtube.SearchListResponse{Items: append(generateMockSearchResult(30, 1, true), generateMockSearchResult(20, 31, false)...), NextPageToken: "next_token_1"}, nil)

		// すべてのタイトルに'SHOWROOM'が含まれている
		mockService.EXPECT().FetchVideos(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), "next_token_1", int64(50)).Return(&youtube.SearchListResponse{Items: generateMockSearchResult(50, 51, true), NextPageToken: "next_token_2"}, nil)

		// すべてのタイトルに'SHOWROOM'が含まれている
		mockService.EXPECT().FetchVideos(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), "next_token_2", int64(50)).Return(&youtube.SearchListResponse{Items: generateMockSearchResult(50, 101, true), NextPageToken: ""}, nil)

		got, err := c.GetVideoIDs()
		assert.NoError(t, err)
		assert.Len(t, got, 100)
		assert.Equal(t, expected, got)
	})

	t.Run("正常系: 'SHOWROOM'がタイトルに含まれていない場合は除外して動画IDを取得できる", func(t *testing.T) {
		expected := []string{"video1", "video2"}
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "SHOWROOM", "date", "video", gomock.Any(), int64(50)).Return(
			&youtube.SearchListResponse{
				Items: []*youtube.SearchResult{
					{Id: &youtube.ResourceId{VideoId: "video1"}, Snippet: &youtube.SearchResultSnippet{Title: "SHOWROOM 配信"}},
					{Id: &youtube.ResourceId{VideoId: "video2"}, Snippet: &youtube.SearchResultSnippet{Title: "SHOWROOM video"}},
					{Id: &youtube.ResourceId{VideoId: "video3"}, Snippet: &youtube.SearchResultSnippet{Title: "Other 配信"}},
				},
				NextPageToken: "",
			}, nil,
		)
		got, err := c.GetVideoIDs()
		assert.NoError(t, err)
		assert.Equal(t, got, expected)
	})
	t.Run("正常系: 小文字の'showroom'がタイトルに含まれる場合は除外して動画IDを取得できる", func(t *testing.T) {
		expected := []string{"video1"}
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "SHOWROOM", "date", "video", gomock.Any(), int64(50)).Return(
			&youtube.SearchListResponse{
				Items: []*youtube.SearchResult{
					{Id: &youtube.ResourceId{VideoId: "video1"}, Snippet: &youtube.SearchResultSnippet{Title: "SHOWROOM 配信"}},
					{Id: &youtube.ResourceId{VideoId: "video2"}, Snippet: &youtube.SearchResultSnippet{Title: "showroom 配信"}},
				},
				NextPageToken: "",
			}, nil,
		)
		got, err := c.GetVideoIDs()
		assert.NoError(t, err)
		assert.Equal(t, got, expected)
	})
	t.Run("異常系: Google Data APIエラー時に適切に処理する", func(t *testing.T) {
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "SHOWROOM", "date", "video", gomock.Any(), int64(50)).Return(nil, &googleapi.Error{})
		got, err := c.GetVideoIDs()
		assert.Nil(t, got)
		assert.Error(t, err)
	})
}

func generateMockSearchResult(count, startID int, isShowroom bool) []*youtube.SearchResult {
	var items []*youtube.SearchResult
	for i := 0; i < count; i++ {
		videoID := fmt.Sprintf("video%d", startID+i)
		title := "SHOWROOM 配信"
		if !isShowroom {
			title = "Other 配信"
		}
		items = append(items, &youtube.SearchResult{
			Id:      &youtube.ResourceId{VideoId: videoID},
			Snippet: &youtube.SearchResultSnippet{Title: title},
		})
	}
	return items
}
