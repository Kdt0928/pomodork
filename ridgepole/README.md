# ridgepole

## フォルダ構成
```
ridgepole
├── Dockerfile
├── README.md
├── config
│   └── database.yml ・・・DBへの接続情報を記載したYML
└── schemas
    └── Schemafile   ・・・テーブル定義ファイル
```

## コマンド
```bash
$ ridgepole -c [YMLファイル] -f [テーブル定義ファイル] オプション
```
### オプション
|オプション|内容|
|---|---|
|`--apply`|テーブル定義ファイルの内容をDBに適用する|
|`--dry-run`|実行計画を作成する|
