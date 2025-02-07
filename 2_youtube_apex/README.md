# テスト 2

# 日本人による直近３日以内の Apex Legends に関する動画の人気 TOP10 を出力するスクリプト

## 実行手順

- フォルダへの移動

```
cd path/to/your/2_youtube_apex
```

- `.sample.env`のファイル名を`.env`に変更し、YouTube Data API の API キーを設定する

```
YOUTUBE_API_KEY=YOUR_API_KEY
```

- Go がインストールされている環境で`go run .`コマンドを実行する

## コードについて

### TOP10 を出力するフロー

- YouTube Data API の [Search List](https://developers.google.com/youtube/v3/docs/search/list?hl=ja) を利用して、直近 3 日以内の`Apex Legends`に関する動画情報を視聴回数順に取得する
- Search List で取得した動画情報を元に YouTube Data API の[Channels List](https://developers.google.com/youtube/v3/docs/channels/list?hl=ja) を利用して、日本人による動画であるか判定する
- 日本人の動画の中から 10 件、ランキング形式でタイトル、チャンネル名、投稿日、URL を出力する

### 期間指定について(直近３日以内の動画データを取得)

Search List API に `publishedAfter` パラメーターを追加し、実行時刻の 3 日前の 00 時 00 分 00 秒以降のデータを取得

例: 実行日時が 2025 年 2 月 7 日の場合、`publishedAfter` には 2025 年 2 月 4 日の 00 時 00 分 00 秒を指定し、その日時以降のデータを取得する

### 人気順について

視聴回数を基準に人気順とする

Search List API を使用して動画データを取得する際、`order` パラメーターに `viewCount` を指定することで、視聴回数順に動画データを取得する

### 日本人による動画の判定方法

Search List API にて取得した動画情報に含まれる Channel ID から Channels List を利用し、チャンネルの居住国が日本の場合は、日本人による動画であると判定する

### テストについて

テスト実行時に YouTube Data API の実際の呼び出しを避けるため、`YoutubeApiService` interface を作成し、本番環境では実際の API サービスを使用し、テスト環境では Mock を使用する

## 今後の改修

- api key などの機密情報は secret 管理ツール使う
- 人気順について、視聴回数以外の指標（いいね数、チャンネル登録者数など）も考慮して精度を高くする
- エラーログは簡易的な出力になっているので、エラーをラップし、トレースをつけることで発生箇所や原因をより明確にする

## 実行結果

実行日時: 2025 年 2 月 7 日

```
第1位
 チャンネル名: ぽこに / Poconi
 動画タイトル: 【APEX】やられたらブチギレ確定ｗｗｗ敵パスを三連続リスキルした結果がこちら#apex #apexlegends #fyp #おすすめ #tiktok #ゲーム #shorts
 投稿日: 2025-02-04 17:50:00
 URL: https://www.youtube.com/watch?v=DqAkaFdyy4U

第2位
 チャンネル名: TIE Ru
 動画タイトル: 【最新情報】シーズン24でアッシュ/オルター大幅強化予定、レジェンドBANシステム、新大会『ALGS OEPN』 | Apex Legends
 投稿日: 2025-02-04 07:00:27
 URL: https://www.youtube.com/watch?v=UB8LKP4V8fM

第3位
 チャンネル名: もつくFPS
 動画タイトル: 意外と知られてないエッジスライディングを解説 | apexlegends  #apex #shorts
 投稿日: 2025-02-04 07:00:06
 URL: https://www.youtube.com/watch?v=AWq9r1MQGWE

第4位
 チャンネル名: 渋谷ハル
 動画タイトル: 【APEX LEGENDS】ALGSでモチベが上がったぼぶきなハルのフルパランク【渋谷ハル】
 投稿日: 2025-02-04 00:31:58
 URL: https://www.youtube.com/watch?v=f34YQv-uhGk

第5位
 チャンネル名: tttcheekyttt
 動画タイトル: ESCL G1スクリム助っ人【Apex Legends】
 投稿日: 2025-02-06 23:44:27
 URL: https://www.youtube.com/watch?v=pnrKB8lewIA

第6位
 チャンネル名: tttcheekyttt
 動画タイトル: プレデターランク w/ SangJoon、MiaKさん【Apex Legends】
 投稿日: 2025-02-05 03:18:44
 URL: https://www.youtube.com/watch?v=YJYEKHvOaj0

第7位
 チャンネル名: じーぷろ【APEX解説】
 動画タイトル: フレンドとの1v1で使えるキャラコン集【APEX LEGENDS】#shorts #エーペックス #apex
 投稿日: 2025-02-04 17:00:09
 URL: https://www.youtube.com/watch?v=ChS7600x5Zg

第8位
 チャンネル名: tttcheekyttt
 動画タイトル: プレデターランク w/ LEO様、さつきんぐ【Apex Legends】
 投稿日: 2025-02-04 02:34:01
 URL: https://www.youtube.com/watch?v=vKjJrUrbUSY

第9位
 チャンネル名: まつげさん(生声)
 動画タイトル: 未だにプレマス帯で6.5％使われてるやつ【Apex　Legends】#shorts
 投稿日: 2025-02-05 16:07:53
 URL: https://www.youtube.com/watch?v=da44yRSypgs

第10位
 チャンネル名: 渋谷ハル
 動画タイトル: 【APEX LEGENDS】APEXブーム到来ってマジですか！？ w/ 紫宮るな 白那しずく【渋谷ハル】
 投稿日: 2025-02-04 23:35:37
 URL: https://www.youtube.com/watch?v=c_GeH0A0qUI
```
