# DI考察

依存関係をバラすために

- 構造体を生成するコンストラクタをモジュールごとに実装
- 初期化のタイミングでそれらを呼び出す（DI）

こうすることで

- モジュールの分離（変更に強いコード）
- 単体テストでモックが書ける

の恩恵が受けられる

## 参考記事
https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/
