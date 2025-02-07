# テスト 1

# Youtube にて、動画タイトルに 'SHOWROOM'' が含まれている の動画 URL を最新順に 100 件抽出する

## 実行手順

- フォルダへの移動

```
cd path/to/your/1_youtube_showroom
```

- `.sample.env`のファイル名を`.env`に変更し、YouTube Data API の API キーを設定する

```
YOUTUBE_API_KEY=YOUR_API_KEY
```

- Go がインストールされている環境で`go run .`コマンドを実行する

## コードについて

### api key の管理について

本来は api key などの機密情報は secret 管理ツールを使うが、今回は`.env`ファイルから取得する

### Youtube Data API を用いた動画 URL の抽出について

YouTube Data API の [Search List](https://developers.google.com/youtube/v3/docs/search/list?hl=ja) を利用して、キーワード`SHOWROOM`に関連する動画情報を取得する

- 最新順のデータを取得するため、パラメーターの`order`に`date`を追加して、作成日時順に取得する
- 一回の API リクエストで取得できるデータの上限が 50 件のため、100 件取得するまで API のリクエストを複数回実行する
- キーワード`SHOWROOM`に関連する動画検索を行うため、コード側でタイトルに`SHOWROOM`が含まれているのかの判断をする
  - タイトルに大文字の `SHOWROOM` が含まれている場合のみ対象とする
- Search List API で取得した video_id を利用し、`https://www.youtube.com/watch?v=VIDEO_ID` の形式で動画 URL を 100 件出力する

### テストについて

テスト実行時に YouTube Data API の実際の呼び出しを避けるため、`YoutubeApiService`interface を作成し、本番環境では実際の API サービスを使用し、テスト環境では Mock を使用する

## 実行結果

実行日時: 2025 年 2 月 7 日

```
No.1 https://www.youtube.com/watch?v=hgZa8Q_zN5w
No.2 https://www.youtube.com/watch?v=fTl1L8pGHXE
No.3 https://www.youtube.com/watch?v=ujb0d40J8g8
No.4 https://www.youtube.com/watch?v=MuFUgV7aJOM
No.5 https://www.youtube.com/watch?v=hIW3PbsOSmU
No.6 https://www.youtube.com/watch?v=LWSaJI7_dqk
No.7 https://www.youtube.com/watch?v=5ug9cCz3wwQ
No.8 https://www.youtube.com/watch?v=N_TdCPN9yto
No.9 https://www.youtube.com/watch?v=7MDq77xBmSs
No.10 https://www.youtube.com/watch?v=UG7S3E6r8l4
No.11 https://www.youtube.com/watch?v=Gr_4xNVY4rY
No.12 https://www.youtube.com/watch?v=VErb-NmOVgQ
No.13 https://www.youtube.com/watch?v=o3gfOBWqaUw
No.14 https://www.youtube.com/watch?v=MjZe3z0Q04Y
No.15 https://www.youtube.com/watch?v=TQ4Gw2RYDQo
No.16 https://www.youtube.com/watch?v=QXvWy-5B0bM
No.17 https://www.youtube.com/watch?v=DmfMbl5bcnE
No.18 https://www.youtube.com/watch?v=hd_0X_RpDy4
No.19 https://www.youtube.com/watch?v=E0S7lrRY2aI
No.20 https://www.youtube.com/watch?v=n0koYcYWQzQ
No.21 https://www.youtube.com/watch?v=nSLH80cizbg
No.22 https://www.youtube.com/watch?v=9OfPvQhJqzA
No.23 https://www.youtube.com/watch?v=Sz6EkoBqACw
No.24 https://www.youtube.com/watch?v=UBhVn7WQV8w
No.25 https://www.youtube.com/watch?v=F8cXM1eIyHo
No.26 https://www.youtube.com/watch?v=b6zR4QP2jnI
No.27 https://www.youtube.com/watch?v=Quu4a2GT6gY
No.28 https://www.youtube.com/watch?v=-YWonGrZUDs
No.29 https://www.youtube.com/watch?v=n-Yo5gbEhpw
No.30 https://www.youtube.com/watch?v=4PXwdX0RkC4
No.31 https://www.youtube.com/watch?v=5lTZbG6Ft_4
No.32 https://www.youtube.com/watch?v=irSvf3Wrvl4
No.33 https://www.youtube.com/watch?v=8gqZQV6dtro
No.34 https://www.youtube.com/watch?v=sN0_jiMKyW4
No.35 https://www.youtube.com/watch?v=epCvGKFYN0U
No.36 https://www.youtube.com/watch?v=4F_j7LIALVo
No.37 https://www.youtube.com/watch?v=KQ9UaQ1dzGM
No.38 https://www.youtube.com/watch?v=KsV7luRIIIs
No.39 https://www.youtube.com/watch?v=3Zhd58sDtPc
No.40 https://www.youtube.com/watch?v=5-x5rdO3LhI
No.41 https://www.youtube.com/watch?v=0OKmVdHPP6o
No.42 https://www.youtube.com/watch?v=27hDfV1aua4
No.43 https://www.youtube.com/watch?v=9XsjNVUxM-E
No.44 https://www.youtube.com/watch?v=LRKBhzTJdmE
No.45 https://www.youtube.com/watch?v=_JEgoswV2uY
No.46 https://www.youtube.com/watch?v=ddtCqLjzfjQ
No.47 https://www.youtube.com/watch?v=r5ElF0tkQzw
No.48 https://www.youtube.com/watch?v=tlvkbcvvPZY
No.49 https://www.youtube.com/watch?v=Pl6T4khSdvk
No.50 https://www.youtube.com/watch?v=ypglSZ5dtBQ
No.51 https://www.youtube.com/watch?v=FCSgfiV4A14
No.52 https://www.youtube.com/watch?v=j-FC2j8K3AM
No.53 https://www.youtube.com/watch?v=5cIYt2pD7ow
No.54 https://www.youtube.com/watch?v=5AyX0RnUtd0
No.55 https://www.youtube.com/watch?v=g3b8V9sKKgc
No.56 https://www.youtube.com/watch?v=fnEA-OuLjy4
No.57 https://www.youtube.com/watch?v=PeWsy_NK11U
No.58 https://www.youtube.com/watch?v=NtyAyvE1dHI
No.59 https://www.youtube.com/watch?v=7KlYDyEJIFs
No.60 https://www.youtube.com/watch?v=AyE5Oyyys1o
No.61 https://www.youtube.com/watch?v=SQgKatBJFWM
No.62 https://www.youtube.com/watch?v=9Y-cDPwg-Ng
No.63 https://www.youtube.com/watch?v=j3NkwtHP44U
No.64 https://www.youtube.com/watch?v=fwp4kzO4RAk
No.65 https://www.youtube.com/watch?v=G7VTzV8FhuI
No.66 https://www.youtube.com/watch?v=ObGr8M9Kr-4
No.67 https://www.youtube.com/watch?v=Q3UCgG0siNs
No.68 https://www.youtube.com/watch?v=swftBR-a328
No.69 https://www.youtube.com/watch?v=kzhGV2DB9xY
No.70 https://www.youtube.com/watch?v=t0LcZt-tmHs
No.71 https://www.youtube.com/watch?v=8zhjNLW0Y5Y
No.72 https://www.youtube.com/watch?v=87LNGeSoVFw
No.73 https://www.youtube.com/watch?v=VAfdOTgV5_U
No.74 https://www.youtube.com/watch?v=7Heed3pGFbk
No.75 https://www.youtube.com/watch?v=9Ok_McLzJjA
No.76 https://www.youtube.com/watch?v=W3aCVRY48_4
No.77 https://www.youtube.com/watch?v=CjvpEiVcb8Y
No.78 https://www.youtube.com/watch?v=u3dxhwHKsHM
No.79 https://www.youtube.com/watch?v=ejDPg-ZvIlI
No.80 https://www.youtube.com/watch?v=NiC8XtZTaFA
No.81 https://www.youtube.com/watch?v=WtHrC8I5SEo
No.82 https://www.youtube.com/watch?v=qYsHIU41H30
No.83 https://www.youtube.com/watch?v=lk9G-S3Fu5Y
No.84 https://www.youtube.com/watch?v=xds2IdVAkEE
No.85 https://www.youtube.com/watch?v=E01tFMK5WlM
No.86 https://www.youtube.com/watch?v=CLCagvdKSSA
No.87 https://www.youtube.com/watch?v=tD4lWFOEf80
No.88 https://www.youtube.com/watch?v=59iMo6L2bt4
No.89 https://www.youtube.com/watch?v=FV3MJo_A-r0
No.90 https://www.youtube.com/watch?v=n3ylkj7TiiU
No.91 https://www.youtube.com/watch?v=kA4IXFGAR9A
No.92 https://www.youtube.com/watch?v=WyPS5kgyWDQ
No.93 https://www.youtube.com/watch?v=tuzLcVIl1uA
No.94 https://www.youtube.com/watch?v=4RIYwvSGfkI
No.95 https://www.youtube.com/watch?v=nw_H8P2tZ94
No.96 https://www.youtube.com/watch?v=zhogZnfxwko
No.97 https://www.youtube.com/watch?v=W7PGOy13Dio
No.98 https://www.youtube.com/watch?v=-UZ9TqQL4IQ
No.99 https://www.youtube.com/watch?v=e-xK5qI81U8
No.100 https://www.youtube.com/watch?v=yxXWnlNYv_M
```
