# アーキテクチャ
このレポジトリでは，レイヤードアーキテクチャにポートを追加して，domainレイヤとinfrastructureレイヤの依存関係を逆転させたアーキテクチャを採用します．

## パッケージ構成

```bash
.
├── cmd
│   └── main.go # エントリーポイント
├── docs # ドキュメント
└── pkg
    ├── application # entityを操作してビジネスロジックを実行します．
    ├── domain
    │   ├── entity # ドメインモデルやドメインロジックを実装します．
    │   ├── repository # 永続化に関するポート(interface)を定義します．
    │   └── service # ドメインモデルに責務を持たせるべきではない処理を実装します．
    ├── infrastructure # repositoryに対する実装を行います．
    └── presentation
        ├── controller # 入出力を行います．ドメインロジックや技術的な実装は置きません．
        └── router.go # ルーティングの設定を行います．
```

## レイヤの依存関係
```
presentation
↓
application
↓
domain
↑
infrastructure
```