# POMODORK-BACKEND
## 環境変数
|Key|値|内容|備考|
|---|---|---|---|
|DB_USER|`pomodork_dev`|DBのユーザ名|-|
|DB_PASSWORD|`developer`|DBのパスワード|-|
|DB_ADDRESS|`127.0.0.1`|DBのアドレス-|
|DB_PORT|`3306`|DBのユーザ名|-|
|DB_NAME|`pomodork_db`|DBのユーザ名|-|
|DB_TIMEZONE|`Asia/Tokyo`|DBのタイムゾーン|-|
|MAX_OPEN_CONNECTION|`25`|最大コネクション数<br>データベース側の接続数上限よりも少ない数にする|-|
|MAX_IDLE_CONNECTION|`25`|最大アイドルコネクション数<br>データベース側の接続数上限よりも少ない数にする|-|
|MAX_LIFE_TIME|`5000`|コネクションの生存時間(ms)<br>確率したコネクションの最大生存時間|-|
