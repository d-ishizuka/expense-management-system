# Expense Management System

このプロジェクトは、GraphQL、gRPC、OpenAPIを使用した経費精算システムのGoバックエンドです。以下に、プロジェクトの構成と使用方法を示します。

## プロジェクト構成

```
expense-management-system
├── cmd
│   └── server
│       └── main.go          // アプリケーションのエントリーポイント
├── configs
│   ├── app.yaml             // アプリケーションの設定ファイル
│   └── database.yaml        // データベース接続の設定ファイル
├── internal
│   ├── api
│   │   ├── graphql          // GraphQL関連の実装
│   │   ├── grpc             // gRPC関連の実装
│   │   └── rest             // REST API関連の実装
│   ├── db                  // データベース関連の実装
│   ├── domain              // ドメインロジック
│   └── utils               // ユーティリティ関数
├── pkg
│   ├── logger               // ロギング機能
│   └── middleware           // ミドルウェア
├── scripts
│   └── db                  // データベース関連のスクリプト
├── api
│   └── openapi             // OpenAPI仕様
├── go.mod                  // Goモジュールの依存関係
├── go.sum                  // Goモジュールの依存関係のバージョン
├── Makefile                // ビルドや実行に関するコマンド
└── README.md               // プロジェクトのドキュメント
```

## 使用方法

1. **依存関係のインストール**
   ```
   go mod tidy
   ```

2. **データベースのマイグレーション**
   マイグレーションスクリプトを実行して、データベースをセットアップします。

3. **アプリケーションの起動**
   ```
   go run cmd/server/main.go
   ```

4. **APIの利用**
   - GraphQLエンドポイント: `/graphql`
   - gRPCエンドポイント: `/grpc`
   - REST APIエンドポイント: `/api`

## ライセンス

このプロジェクトはMITライセンスの下で提供されています。# expense-management-system
