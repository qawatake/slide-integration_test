---
marp: true
---

# integration test書いてみた

---

# usecaseのテストで疲弊していませんか？

## つらみ

- mockまみれ😣
  - testを書くにはまず実装を読まねばならない。
  - メソッド名は？引数は？
- repo.FindXXX/ListXXXの回数なんてチェックしたくない😣
  - 参照の回数なんて仕様とは関係ない。
- ちょっとリファクタリングしただけで壊れる😣
  - adapterの処理をusecaseに持っていったり、モジュール間で処理を移しただけでテストごと作り直し
- DBとのやりとりの過程なんて知りたくない。知りたいのは最終的な状態😣

---

# integration testを書いてみよう

> xx

「単体テストの使い方/考え方」

- ここではintegration testの単位は1プロセスとします。
  - Railsでいう、request spec的な。
  - batchやworkerも対象。

---

# こんなintegration testだったらうれしい

- 可能な限り実際に動くプログラムとの差分が小さい。
  - integration testの意義の一つは、多くの依存先を結合してちゃんと動くことを確認すること。
  - mockするのはできる限り低いレイヤにしたい。
- DIが楽ちん。
  - 本家のDIに修正が入っても、integration testの修正は軽微。

---

# ディレクトリ構成

```
.
|-- integration_test
    |-- greeting: greetingバッチ
    |   |-- wire.go: wire.ProviderSetを定義
    |   |-- injector.go: wire.Buildするところ
    |   |-- wire_gen.go: wireの生成コード
    |   `-- notification_test.go: テスト
    `-- server
        |-- wire.go
        |-- injector.go
        |-- wire_gen.go
        |-- http/external/hoge/hoge_test.go
        |-- http/internal/hoge/hoge_test.go
        `-- http/private/fuga/fuga_test.go
```

- 1つのmain.goに対して1つのディレクトリ。
- HTTPについてはパスとディレクトリを一致させるとわかりやすそう？

---

# まずは本家（プロダクションコード）のwire.go

```go
// usecase層やadapter層は一つのprovider setにまとめてしまう。
var CoreSet = wire.NewSet(
  usecase.New,
  wire.Bind(new(usecase.Greeter), new(*greeting.Client)),
  greeting.Set,
  ...
)

// mockするconfig層だけprovider setを分ける。
var configSet = wire.NewSet(
  config.NewGreetingConfig,
)
```

---

# 本家のinjector.go

config層はprovider set `configSet`から渡される。

```go
func New(
  context.Context,
) (*usecase.Greeting, error) {
  panic(wire.Build(
    CoreSet,
    configSet, // <- ここ！
  ))
}
```

---

# integration test用のinjector.go

integration testではinjector functionの引数としてconfig層を渡す。

```go
func New(
  context.Context,
  config.Greeting, // <- ここ！
) (*app.Notification, error) {
  panic(wire.Build(
    app.CoreSet,
  ))
}
```

---

# greeting_test.go

```go
func TestGreeting(t *testing.T) {
  // DI
  ts := httpstub.NewServer(t)
  cfg := config.GreetingConfig{
    URL: ts.URL,
  }
  u, _ := greeting.New(ctx, cfg)

  // arrange
  ts.Path("/hello").Method(http.MethodPost).Response(http.StatusOK, nil)

  // act
  ctx = context.Background()
  u.Greet(ctx)

 // assert
  reqs := ts.Requests()
  b, _ := io.ReadAll(reqs[0].Body)
  got := string(b)
  want := "Hello, World!"
  if got != want {
    t.Errorf("got %v, want %s", got, want)
 }
}
```

---

# （再掲）こんなintegration testだったらうれしい

- 可能な限り実際に動くプログラムとの差分が小さい。
- DIが楽ちん。

---

# 可能な限り実際に動くプログラムとの差分が小さい？？

---

# DIが楽ちん？？

---

# どうmockする？？

使っているライブラリ

- [httpstub](https://github.com/k1LoW/httpstub): HTTPサーバをモック。
- [smtptest](https://github.com/k1LoW/smtptest): メール送信をモック。
- [matryer/moq](https://github.com/matryer/moq): AWS SDKとgRPCサーバのmockに使っているよ。
  - localstackを使うという意見もある。
  - grpcstubは.protoファイルが必要。テストのためにそれを引っ張ってくるかは微妙なところ。。（バージョンの食い違いが発生したら嫌だ。）

---

# どうmockする？？

使っているmockライブラリの共通点

- テストの準備（arrange）フェーズと検証（assert）フェースを分離しやすい。
- gomockはarrangeとassertの分離が厳しい。（`DoAndReturn`はスタブの準備と検証が混ざっている。。）

---

# integration test書いてみたくなりましたか？？
# ご清聴ありがとうございましたmm
