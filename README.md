# Jyobi-Project

## バリデーションチェック
バリデーションの種類はここに載っている！！！

👇👇👇👇👇👇👇👇👇👇👇👇👇👇👇👇

参考サイト：[ozzo-validation](https://github.com/go-ozzo/ozzo-validation)

👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆

### インストール
```go
  go get github.com/go-ozzo/ozzo-validation
```

### エラー
インストールしたとき`missing go.sum entry for module providing package`エラーでた
```go
  go mod tidy
```
で直った。
### コマンド詳細
go mod tidyコマンドでは、モジュール内の全てのパッケージのビルドタグの組み合わせを確認するので、安全に不要なモジュールを削除することができます。
一方で、gobuildやgotestでは、不要なパッケージの削除はされません。ビルド時に単一のパッケージのロードしか行われず、使用されていないパッケージを知ることができないためです。

らしい。

