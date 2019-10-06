# Laravel使って出てきたエラーたち
## まず
### Laravelのログ出るとこ
- デフォルトはstorage/logs/laravel-yyyy-mm-dd.log

## エラー集

### No application encryption key has been specified. {"exception":"[object]
- 暗号化用のキーが設定されていない

### The only supported ciphers are AES-128-CBC and AES-256-CBC with the correct key lengths.
- 暗号化用のキーが設定されているんだけど暗号化で規定されているルールに則ってない
  - ランダムの32文字で構成(master key)

### docker上のMysqlに接続できん
- docker-composeファイルで指定したcontainer＿nameでホスト指定してやる
- user別で追加した場合は.Dockerfileでgrantを一連で追加する必要あり

