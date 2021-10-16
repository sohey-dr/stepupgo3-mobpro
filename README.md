# go vetできるWebアプリ

## 課題

* 任意のモジュールにgo vetを実行し結果を表示する
* リリースされているすべてのバージョンを対象とする
* モジュールパスまたはインポートパスをURLで指定する
  * 例：https://example.com/github.com/tenntenn/greeting/v2
* インポートパスが指定された場合もモジュール全体を対象とする
* 多重に処理は実行しない
  * 1つのモジュールの各バージョンに対して1回しかgo vetを実行しない
  * go vetを実行中の場合などは処理中であることを示す
