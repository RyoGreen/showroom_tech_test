package youtube

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
)

func TestGetVideoInfos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockYoutubeApiService(ctrl)

	c := NewYoutubeClient(mockService)

	t.Run("正常系: 動画情報を10件取得できる", func(t *testing.T) {
		expected := []*VideoInfo{
			{VideoID: "video_id_1", Title: "title_1", ChannelTitle: "channel_title_1", PublishedAt: "2025-01-01 01:00:00"},
			{VideoID: "video_id_2", Title: "title_2", ChannelTitle: "channel_title_2", PublishedAt: "2025-01-01 02:00:00"},
			{VideoID: "video_id_3", Title: "title_3", ChannelTitle: "channel_title_3", PublishedAt: "2025-01-01 03:00:00"},
			{VideoID: "video_id_4", Title: "title_4", ChannelTitle: "channel_title_4", PublishedAt: "2025-01-01 04:00:00"},
			{VideoID: "video_id_5", Title: "title_5", ChannelTitle: "channel_title_5", PublishedAt: "2025-01-01 05:00:00"},
			{VideoID: "video_id_6", Title: "title_6", ChannelTitle: "channel_title_6", PublishedAt: "2025-01-01 06:00:00"},
			{VideoID: "video_id_7", Title: "title_7", ChannelTitle: "channel_title_7", PublishedAt: "2025-01-01 07:00:00"},
			{VideoID: "video_id_8", Title: "title_8", ChannelTitle: "channel_title_8", PublishedAt: "2025-01-01 08:00:00"},
			{VideoID: "video_id_9", Title: "title_9", ChannelTitle: "channel_title_9", PublishedAt: "2025-01-01 09:00:00"},
			{VideoID: "video_id_10", Title: "title_10", ChannelTitle: "channel_title_10", PublishedAt: "2025-01-01 10:00:00"},
		}
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "Apex Legends", "viewCount", "video", gomock.Any(), gomock.Any(), int64(20)).Return(&youtube.SearchListResponse{Items: mockYoutubeSearchResult[:10]}, nil)
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(mockChannelListResponse, nil).Times(10)
		got, err := c.GetVideoInfos()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("正常系: 取得した動画情報が10件以上でも10件のみ返す", func(t *testing.T) {
		expected := []*VideoInfo{
			{VideoID: "video_id_1", Title: "title_1", ChannelTitle: "channel_title_1", PublishedAt: "2025-01-01 01:00:00"},
			{VideoID: "video_id_2", Title: "title_2", ChannelTitle: "channel_title_2", PublishedAt: "2025-01-01 02:00:00"},
			{VideoID: "video_id_3", Title: "title_3", ChannelTitle: "channel_title_3", PublishedAt: "2025-01-01 03:00:00"},
			{VideoID: "video_id_4", Title: "title_4", ChannelTitle: "channel_title_4", PublishedAt: "2025-01-01 04:00:00"},
			{VideoID: "video_id_5", Title: "title_5", ChannelTitle: "channel_title_5", PublishedAt: "2025-01-01 05:00:00"},
			{VideoID: "video_id_6", Title: "title_6", ChannelTitle: "channel_title_6", PublishedAt: "2025-01-01 06:00:00"},
			{VideoID: "video_id_7", Title: "title_7", ChannelTitle: "channel_title_7", PublishedAt: "2025-01-01 07:00:00"},
			{VideoID: "video_id_8", Title: "title_8", ChannelTitle: "channel_title_8", PublishedAt: "2025-01-01 08:00:00"},
			{VideoID: "video_id_9", Title: "title_9", ChannelTitle: "channel_title_9", PublishedAt: "2025-01-01 09:00:00"},
			{VideoID: "video_id_10", Title: "title_10", ChannelTitle: "channel_title_10", PublishedAt: "2025-01-01 10:00:00"},
		}
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "Apex Legends", "viewCount", "video", gomock.Any(), gomock.Any(), int64(20)).Return(&youtube.SearchListResponse{Items: mockYoutubeSearchResult[:14]}, nil)
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(mockChannelListResponse, nil).Times(10)
		got, err := c.GetVideoInfos()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
	t.Run("正常系: 複数回のAPIの呼び出しで正しく取得できる", func(t *testing.T) {
		expected := []*VideoInfo{
			{VideoID: "video_id_1", Title: "title_1", ChannelTitle: "channel_title_1", PublishedAt: "2025-01-01 01:00:00"},
			{VideoID: "video_id_2", Title: "title_2", ChannelTitle: "channel_title_2", PublishedAt: "2025-01-01 02:00:00"},
			{VideoID: "video_id_3", Title: "title_3", ChannelTitle: "channel_title_3", PublishedAt: "2025-01-01 03:00:00"},
			{VideoID: "video_id_4", Title: "title_4", ChannelTitle: "channel_title_4", PublishedAt: "2025-01-01 04:00:00"},
			{VideoID: "video_id_5", Title: "title_5", ChannelTitle: "channel_title_5", PublishedAt: "2025-01-01 05:00:00"},
			{VideoID: "video_id_6", Title: "title_6", ChannelTitle: "channel_title_6", PublishedAt: "2025-01-01 06:00:00"},
			{VideoID: "video_id_7", Title: "title_7", ChannelTitle: "channel_title_7", PublishedAt: "2025-01-01 07:00:00"},
			{VideoID: "video_id_8", Title: "title_8", ChannelTitle: "channel_title_8", PublishedAt: "2025-01-01 08:00:00"},
			{VideoID: "video_id_9", Title: "title_9", ChannelTitle: "channel_title_9", PublishedAt: "2025-01-01 09:00:00"},
			{VideoID: "video_id_10", Title: "title_10", ChannelTitle: "channel_title_10", PublishedAt: "2025-01-01 10:00:00"},
		}
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "Apex Legends", "viewCount", "video", gomock.Any(), gomock.Any(), int64(20)).Return(&youtube.SearchListResponse{Items: mockYoutubeSearchResult[:5], NextPageToken: "token"}, nil)
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "Apex Legends", "viewCount", "video", gomock.Any(), gomock.Any(), int64(20)).Return(&youtube.SearchListResponse{Items: mockYoutubeSearchResult[5:14]}, nil)
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(mockChannelListResponse, nil).Times(10)
		got, err := c.GetVideoInfos()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
	t.Run("正常系: 日本人のチャンネルのみを返す", func(t *testing.T) {
		expected := []*VideoInfo{
			{VideoID: "video_id_1", Title: "title_1", ChannelTitle: "channel_title_1", PublishedAt: "2025-01-01 01:00:00"},
			{VideoID: "video_id_2", Title: "title_2", ChannelTitle: "channel_title_2", PublishedAt: "2025-01-01 02:00:00"},
			{VideoID: "video_id_3", Title: "title_3", ChannelTitle: "channel_title_3", PublishedAt: "2025-01-01 03:00:00"},
			{VideoID: "video_id_4", Title: "title_4", ChannelTitle: "channel_title_4", PublishedAt: "2025-01-01 04:00:00"},
			{VideoID: "video_id_5", Title: "title_5", ChannelTitle: "channel_title_5", PublishedAt: "2025-01-01 05:00:00"},
			{VideoID: "video_id_11", Title: "title_11", ChannelTitle: "channel_title_11", PublishedAt: "2025-01-01 11:00:00"},
			{VideoID: "video_id_12", Title: "title_12", ChannelTitle: "channel_title_12", PublishedAt: "2025-01-01 12:00:00"},
			{VideoID: "video_id_13", Title: "title_13", ChannelTitle: "channel_title_13", PublishedAt: "2025-01-01 13:00:00"},
			{VideoID: "video_id_14", Title: "title_14", ChannelTitle: "channel_title_14", PublishedAt: "2025-01-01 14:00:00"},
			{VideoID: "video_id_15", Title: "title_15", ChannelTitle: "channel_title_15", PublishedAt: "2025-01-01 15:00:00"},
		}
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "Apex Legends", "viewCount", "video", gomock.Any(), gomock.Any(), int64(20)).Return(&youtube.SearchListResponse{Items: mockYoutubeSearchResult[:9], NextPageToken: "token"}, nil)
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(&youtube.ChannelListResponse{Items: []*youtube.Channel{{Snippet: &youtube.ChannelSnippet{Country: "JP"}}}}, nil).Times(5)
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(&youtube.ChannelListResponse{Items: []*youtube.Channel{{Snippet: &youtube.ChannelSnippet{Country: "USA"}}}}, nil).Times(5)
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "Apex Legends", "viewCount", "video", gomock.Any(), gomock.Any(), int64(20)).Return(&youtube.SearchListResponse{Items: mockYoutubeSearchResult[9:15]}, nil)
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(&youtube.ChannelListResponse{Items: []*youtube.Channel{{Snippet: &youtube.ChannelSnippet{Country: "JP"}}}}, nil).Times(5)
		got, err := c.GetVideoInfos()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
	t.Run("正常系: NextPage Tokenがなかった場合、正常にループが終了する", func(t *testing.T) {
		expected := []*VideoInfo{
			{VideoID: "video_id_1", Title: "title_1", ChannelTitle: "channel_title_1", PublishedAt: "2025-01-01 01:00:00"},
			{VideoID: "video_id_2", Title: "title_2", ChannelTitle: "channel_title_2", PublishedAt: "2025-01-01 02:00:00"},
			{VideoID: "video_id_3", Title: "title_3", ChannelTitle: "channel_title_3", PublishedAt: "2025-01-01 03:00:00"},
			{VideoID: "video_id_4", Title: "title_4", ChannelTitle: "channel_title_4", PublishedAt: "2025-01-01 04:00:00"},
			{VideoID: "video_id_5", Title: "title_5", ChannelTitle: "channel_title_5", PublishedAt: "2025-01-01 05:00:00"},
		}
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "Apex Legends", "viewCount", "video", gomock.Any(), gomock.Any(), int64(20)).Return(&youtube.SearchListResponse{Items: mockYoutubeSearchResult[:5], NextPageToken: ""}, nil)
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(mockChannelListResponse, nil).Times(5)
		got, err := c.GetVideoInfos()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
	t.Run("異常系: Google Data APIエラー時に適切に処理する", func(t *testing.T) {
		mockService.EXPECT().FetchVideos([]string{"snippet", "id"}, "Apex Legends", "viewCount", "video", gomock.Any(), gomock.Any(), int64(20)).Return(nil, &googleapi.Error{})
		got, err := c.GetVideoInfos()
		assert.Nil(t, got)
		assert.Error(t, err)
	})
}

func TestFilterValidVideo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := NewMockYoutubeApiService(ctrl)

	c := NewYoutubeClient(mockService)
	t.Run("正常系: 動画情報を取得できる", func(t *testing.T) {
		expected := &VideoInfo{
			VideoID:      "video_id_1",
			Title:        "title_1",
			ChannelTitle: "channel_title_1",
			PublishedAt:  "2025-01-01 01:00:00",
		}
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(&youtube.ChannelListResponse{Items: []*youtube.Channel{{Snippet: &youtube.ChannelSnippet{Country: "JP"}}}}, nil).Times(1)
		got, isValid, err := c.filterValidVideo(&youtube.SearchResult{Snippet: &youtube.SearchResultSnippet{ChannelId: "channel_1", Title: "title_1", ChannelTitle: "channel_title_1", PublishedAt: "2024-12-31T16:00:00Z"}, Id: &youtube.ResourceId{VideoId: "video_id_1"}})
		assert.NoError(t, err)
		assert.True(t, isValid)
		assert.Equal(t, expected, got)
	})
	t.Run("正常系: 日本以外のチャンネルはスキップされる", func(t *testing.T) {
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(&youtube.ChannelListResponse{Items: []*youtube.Channel{{Snippet: &youtube.ChannelSnippet{Country: "USA"}}}}, nil).Times(1)
		got, isValid, err := c.filterValidVideo(&youtube.SearchResult{Snippet: &youtube.SearchResultSnippet{ChannelId: "channel_1", Title: "title_1", ChannelTitle: "channel_title_1", PublishedAt: "2024-12-31T16:00:00Z"}, Id: &youtube.ResourceId{VideoId: "video_id_1"}})
		assert.NoError(t, err)
		assert.False(t, isValid)
		assert.Nil(t, got)
	})
	t.Run("正常系: 公開日は不正の場合はスキップされる", func(t *testing.T) {
		got, isValid, err := c.filterValidVideo(&youtube.SearchResult{Snippet: &youtube.SearchResultSnippet{ChannelId: "channel_1", Title: "title_1", ChannelTitle: "channel_title_1", PublishedAt: ""}, Id: &youtube.ResourceId{VideoId: "video_id_1"}})
		assert.NoError(t, err)
		assert.False(t, isValid)
		assert.Nil(t, got)
	})
	t.Run("異常系: Youtube Data APIにてエラー", func(t *testing.T) {
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(nil, &googleapi.Error{}).Times(1)
		got, isValid, err := c.filterValidVideo(&youtube.SearchResult{Snippet: &youtube.SearchResultSnippet{ChannelId: "channel_1", Title: "title_1", ChannelTitle: "channel_title_1", PublishedAt: "2024-12-31T16:00:00Z"}, Id: &youtube.ResourceId{VideoId: "video_id_1"}})
		assert.Error(t, err)
		assert.False(t, isValid)
		assert.Nil(t, got)
	})
}

func TestIsJapaneseChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockYoutubeApiService(ctrl)

	c := NewYoutubeClient(mockService)

	t.Run("正常系: チャンネルの居住国が'日本'の場合は日本人のチャンネルであると判定する", func(t *testing.T) {
		expected := true
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(&youtube.ChannelListResponse{Items: []*youtube.Channel{{Snippet: &youtube.ChannelSnippet{Country: "JP"}}}}, nil).Times(1)
		got, err := c.isJapaneseChannel(gomock.Any().String())
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
	t.Run("正常系: チャンネルの居住国が'日本'以外の場合は日本人のチャンネルでないと判定する", func(t *testing.T) {
		expected := false
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(&youtube.ChannelListResponse{Items: []*youtube.Channel{{Snippet: &youtube.ChannelSnippet{Country: "USA"}}}}, nil).Times(1)
		got, err := c.isJapaneseChannel(gomock.Any().String())
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
	t.Run("正常系: チャンネル情報が取得できない場合はfalseを返す", func(t *testing.T) {
		expected := false
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(&youtube.ChannelListResponse{Items: []*youtube.Channel{}}, nil).Times(1)
		got, err := c.isJapaneseChannel(gomock.Any().String())
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
	t.Run("異常系: Google Data APIエラー時はfalseを返す", func(t *testing.T) {
		expected := false
		mockService.EXPECT().FetchChannelInfo([]string{"snippet"}, gomock.Any()).Return(nil, &googleapi.Error{}).Times(1)
		got, err := c.isJapaneseChannel(gomock.Any().String())
		assert.Error(t, err)
		assert.Equal(t, expected, got)
	})
}

var mockYoutubeSearchResult = []*youtube.SearchResult{
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_1",
			Title:        "title_1",
			ChannelTitle: "channel_title_1",
			PublishedAt:  "2024-12-31T16:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_1"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_2",
			Title:        "title_2",
			ChannelTitle: "channel_title_2",
			PublishedAt:  "2024-12-31T17:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_2"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_3",
			Title:        "title_3",
			ChannelTitle: "channel_title_3",
			PublishedAt:  "2024-12-31T18:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_3"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_4",
			Title:        "title_4",
			ChannelTitle: "channel_title_4",
			PublishedAt:  "2024-12-31T19:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_4"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_5",
			Title:        "title_5",
			ChannelTitle: "channel_title_5",
			PublishedAt:  "2024-12-31T20:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_5"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_6",
			Title:        "title_6",
			ChannelTitle: "channel_title_6",
			PublishedAt:  "2024-12-31T21:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_6"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_7",
			Title:        "title_7",
			ChannelTitle: "channel_title_7",
			PublishedAt:  "2024-12-31T22:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_7"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_8",
			Title:        "title_8",
			ChannelTitle: "channel_title_8",
			PublishedAt:  "2024-12-31T23:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_8"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_9",
			Title:        "title_9",
			ChannelTitle: "channel_title_9",
			PublishedAt:  "2025-01-01T00:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_9"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_10",
			Title:        "title_10",
			ChannelTitle: "channel_title_10",
			PublishedAt:  "2025-01-01T01:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_10"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_11",
			Title:        "title_11",
			ChannelTitle: "channel_title_11",
			PublishedAt:  "2025-01-01T02:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_11"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_12",
			Title:        "title_12",
			ChannelTitle: "channel_title_12",
			PublishedAt:  "2025-01-01T03:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_12"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_13",
			Title:        "title_13",
			ChannelTitle: "channel_title_13",
			PublishedAt:  "2025-01-01T04:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_13"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_14",
			Title:        "title_14",
			ChannelTitle: "channel_title_14",
			PublishedAt:  "2025-01-01T05:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_14"},
	},
	{
		Snippet: &youtube.SearchResultSnippet{
			ChannelId:    "channel_15",
			Title:        "title_15",
			ChannelTitle: "channel_title_15",
			PublishedAt:  "2025-01-01T06:00:00Z",
		},
		Id: &youtube.ResourceId{VideoId: "video_id_15"},
	},
}

var mockChannelListResponse = &youtube.ChannelListResponse{
	Items: []*youtube.Channel{{Snippet: &youtube.ChannelSnippet{Country: "JP"}}},
}
