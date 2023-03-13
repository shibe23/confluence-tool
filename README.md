# Confluence Tools
ユーザー定義ファイルを元にConfluenceのページを作成します

# Setup

`~/.zshrc` などに下記の値を環境変数として登録します。

```shell
export CONFLUENCE_DOMAIN='your domain'
export CONFLUENCE_USER='your email'
export CONFLUENCE_TOKEN='your API token'
```

# How to Use
## jsonファイルの準備
実行用のユーザー定義ファイルを作成します。

```shell
{
  "keys": [
    {
      "title": "2023-03-01 MTG議事録 - ${value} さん", // ${value}と書かれた部分が置き変わる
      "value": "田中 光一",
      "space": "~1111111111", // Confluenceのスペース
      "ancestor": "1234567890", // 親ページのID
      "templateID": "1122334455" // テンプレートとなるページのID
    },
    {
      "title": "2023-03-01 MTG議事録 - ${value} さん",
      "value": "鈴木 亜希",
      "space": "~1111111111",
      "ancestor": "1234567890",
      "templateID": "1122334455"
    }
  ]
}
```

## 実行

```shell
./confluence-tool create --path="./rules/example.json"
```