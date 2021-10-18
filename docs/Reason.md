# Reason | Go言語を使う理由

[golang.org](https://golang.org/ref/spec#Introduction)

[Rob Pike, "Simplicity is Complicated", 2015](https://talks.golang.org/2015/simplicity-is-complicated.slide#1)

![image1](..//img//image1.png)

## どのような言語なのか

### 概要

Goは2009年にGoogleが開発したオープンソースプロジェクトのプログラミング言語であり、**大きなクラウドソフトウェアのための簡潔な手続き型言語**としてデザインされている。

- [厳密な型定義](http://go.shibu.jp/go_tutorial.html#id7)
- [関数 / メソッド](http://go.shibu.jp/effective_go.html#id12)
- [インターフェース](http://go.shibu.jp/effective_go.html#interfaces-and-other-types)
- パッケージ
- [並列処理](http://go.shibu.jp/effective_go.html#id26)
の全てが簡潔に使える。


Goの[簡潔さ](#簡潔さsimplicity)の特徴として以下のようなものがある。

- [ガベージコレクション](#ガベージコレクション)
- [ゴルーチン](ゴルーチン)
- [定数](#定数)
- [インターフェース](#インターフェース)
- [パッケージ](#パッケージ)
それぞれ、実際は複雑だが、簡潔に使えるのが特徴である。

### 生い立ち

Goは基本的な文法の大部分はC言語系から受け継いでいます。また、宣言やパッケージについては、Pascal、Modula、Oberonから大きな影響を受けています。加えて、並行処理については、Tony HoareのCSP (Communicating Sequential Process) に影響を受けている言語であるNewsqueak、Limboから同じアイディアを取り入れています。

しかしながら、全体としては新しい言語です。あらゆる点において、Goはプログラマが何をし、どうプログラムを作るのかということを考えられて設計されています。少なくとも私たち自身が行うプログラミングがより効率的に、より楽しくあるようにです。

### なぜC言語とこんなに文法が違うのか

宣言文を除くと、違いは大きくはありません。2つの希望によってこのようになっています。

まず最初に、文法は軽く感じるようにすべきで、義務的なキーワードがあまりにも多かったり、繰り返し書かなければならなかったり、呪文のようだったりというのは辞めたいと考えていました。

二つ目は、プログラミング言語はシンボルテーブルをパースしなくても簡単に分析できるように設計されました。この設計方針の影で、デバッガー、依存関係のアナライザ、自動ドキュメント抽出ツール、IDEプラグインなどのツールの開発がより簡単になります。C言語とその子孫の言語はこの点が困難で悪名をとどろかせていました。

### なぜガーベジコレクタを動作させるのか

システムプログラミングのおいて、定型的な作業の中でもっともソースコードを占めているものの一つがメモリ管理です。私たちは、プログラマーのオーバーヘッドを取り除くことが重要だと考えています。ここ数年のガーベジコレクタの技術の進歩により、システムプログラムを作る上で作業上のオーバーヘッドを十分に減らしつつ、実行時に大きな遅延の影響を受けることもなくなりました。現在の実装は単純なマーク・アンド・スイープのガーベジコレクタを使用していますが、現在置き換えの作業を行っています。

もう一つのポイントは、並列のマルチスレッドプログラミングの難しさの大部分を占めているのがメモリ管理だ、という点です。オブジェクトがスレッド間でやりとりされはじめると、そのオブジェクトを安全に開放できるかどうか保証するのが難しくなります。自動ガーベジコレクションを適用することで、並列プログラミングのコードを書くのが極めて簡単になります。もちろん、並列環境でのガーベジコレクションを実装するのは、それ自身、一つのチャレンジではあるのですが、この一回の努力は、全てのプログラム、すべての人を助けることになると考えています。

並列性の問題は置いておいても、ガーベジコレクタのおかげで、最終的にソースコード中のインタフェースはシンプルになっていきます。インタフェースのこちらと向こうで、メモリをどう管理するのか？というのを指定する必要がなくなるからです。

### なぜスレッドではなくgoroutineなのか

goroutineは並列処理を使い易くします。少し前からあたためていたそのアイデアとは、独立に実行する複数の関数(コルーチン)を、スレッドの集合に多重化することでした。ブロッキングするシステムコールを呼んだ場合などでコルーチンがブロックされる際に、ランタイムは同一スレッドにある他のコルーチンたちを別の実行可能なスレッドに自動的に移動して、それらがブロックされないようにします。プログラマはこの場面を見ることはありませんが、これこそが重要なのです。私たちがgoroutineと呼ぶこの仕組みは、非常に軽い処理にすることができます。実行時間の長いシステムコールで長時間費やさなければ、そのコストはスタック用のメモリ処理にかかるものより少し多い程度で済みます。

スタックを小さくするため、Go言語のランタイムはセグメント化されたスタックを使います。新しい出来立てのgoroutineでは数キロバイトが割り当てられますが、それはほとんどの場合で充分な大きさです。充分でない場合でも、ランタイムは追加でセグメントを自動的に割り当て(そして解放)します。関数呼び出しごとのオーバヘッドは、処理の軽い命令３つ分ぐらいが平均的なものです。同一アドレス空間で数十万規模のgoroutineを生成できるほど実用的なのです。もしgoroutineが単なるスレッドであったら、システムリソースはもっと小さな規模で枯渇してしまったでしょう。

### 簡潔さ(Simplicity)

Java, JavaScript (ECMAScript), Typescript, C#, C++, Hack (PHP)などの話を聞いていると、これらの言語は互いに特徴を取り入れ合っている。あたかも、1つの巨大な言語に収束しようとしている。

上にあげたような言語は、特徴を次々と追加することで、進化に完成に向かっている。それぞれ似たような言語になりながら、どんどんと複雑になっている。区別をしっかりとつけないまま肥大化しているだけである。

しかし、Goは違う。Goを初めて触った人から他言語の特徴がなぜ入っていないのかなどの質問をよく受けるが、Goは他の言語のようになろうと特徴を追加したりしない。特徴を追加していくことがGoを良くすることだとは思っていない、ただ肥大化させるだけである。そして、GoはGoとしての特徴を徐々に失っていき、面白くなくなっていく。

#### 読みやすさ(Readability)

言語に特徴が多すぎると、書かれているコードがそのように動く理由がわからなくなる。言語が複雑さを増すと、単純に理解しにくくなるのだ。特徴が増えると複雑になるだけだ、simplicity（簡潔な）方がいい。特徴が増えると読みにくい。readability（読みやすい）方がいい。

そして、読みやすいコードだと理解しやすい、取り組みやすい、修正しやすい、つまり信頼できるのである。

#### 表現力(Expressiveness)

表記が簡略であること（簡略記法）は表現力が高くあり得るが、読みやすいとは限らない。以下のコードは、とても強力なプリミティブを用いているが、どう動くかが良くわからない。

```
Consider APL: ⊃ 1 ω ∨ . ∧ 3 4 = +/ +⌿ 1 0 ‾1 ∘.θ 1 - ‾1 Φ″ ⊂ ω
```

一方で、冗長な書き方をしたところで読みやすくはなるが、コードの意図は不明瞭になる。

表現力を保ちつつも、簡略な表記をするべきである。

#### ガベージコレクション

これが最もうまく複雑さを見せずに簡潔に見せかけているものである。Goのコードを簡潔に保てる理由は、ガベージコレクションが存在するからである。コードを書く方は全く意識しなくて良い。

#### ゴルーチン

```go
go function(args)
```

ゴルーチンは、`g` `o` ` ` という3文字で構成される。とても簡潔である。ガベージコレクションと同じで、スタックのサイズやIDなど、プログラマーは全く気にしなくて良い。他の言語であったら提供されているであろうこれらの事をGoでは全く考えなくていいという所にGoの簡潔さが良く表れている。

#### 定数

Goは厳密な静的型付け言語だが、定数には型を持たせないことが可能であり、実際の数字のように扱うことができる。

```go
var nanosecond = time.Second/1e9
```

[more detailed .....](https://qiita.com/hkurokawa/items/a4d402d3182dff387674)

#### インターフェース

インターフェースは、Goの最大かつパワフルな特徴である。ライブラリをデザインするときに効果を発揮する。プログラマーとしてはとても簡潔に使えるが、裏側の複雑さがそれを可能にしている。

例えば、以下の例がそれを強く示している。

```go
type Reader interface {
    Read([]byte) (int, error)
}
```

```go
var r Reader = os.Stdin // Statically checked.
var x interface{} = os.Stdin // Statically checked.
r = x.(Reader) // Dynamically checked. Must be explicit here - design decision.
```

[more detailed .....](https://qiita.com/weloan/items/de3b1bcabd329ec61709)

#### パッケージ

パッケージは、プログラムやライブラリを構築するためのデザインである。

```go
package big

import "math/big"
```

これをデザインするのに長い時間を要した。コンポーネント化・スケーラビリティ・共有・データ隠蔽・分離を可能にした。また、プログラムのデザイン・シンタックス・ネーミング・ビルド・リンク・テストに大きな影響を与えている。

パッケージのパス"math/big"をパッケージ名bigから分離することで、go getのメカニズムを作り出すことができた。実装は複雑なものだが、自然に使うことができる。

## なぜ使うのか / なにができるのか

### 問題意識

ここ10年以上の間、メジャーなシステム言語と呼ばれるものは登場していません。しかし、この間コンピュータをめぐる景色は大きく変化しています。この変化には以下のような傾向があります:

- コンピュータはとてつもなく速くなりましたがが、ソフトウェア開発は速くなっていません。
- 依存関係の管理というものが、今日のソフトウェア開発において、大きな位置を占めていますが、C言語の伝統のヘッダファイルはクリーンな依存性分析と高速なコンパイルとは対極です。
- JavaやC++のような扱いにくい型システムに対する反乱が徐々に拡大しつつあります。PythonやJavaScriptのような、動的な型を持つ言語に対して、人々が押し寄せてきています。
- 人気のあるシステム言語の中には、ガーベジコレクションや並列計算のような基本概念をサポートしていないものがあります。
- マルチコアのコンピュータの出現によって、人々の間に、不安と混乱が発生しました。


### 解決策としてのGo

私たちは、並列をサポートし、ガベージコレクションを備えた、コンパイルの速い言語に挑戦することは価値があることだと信じています。上記の点に関して:

- 大きなGoのプログラムを一台のコンピュータで数秒で数秒以内にコンパイルすることが可能です。
- Goは、依存性の分析を簡単にし、C言語スタイルのインクルードファイルや、ライブラリの持つオーバーヘッドのほとんどを解消するようなソフトウェアの構成が行えるモデルを提供します。
- Goの型システムには階層構造がないため、型同士の関係を定義するに貴重な時間を使う必要はありません。また、Goは静的な型システムを持つ言語ですが、典型的なオブジェクト指向言語と比べると型を作成するのが気軽に行えるような言語になっています。
- Goは完全なガーベジコレクションを備え、言語レベルで並列実行とコミュニケーションの機能をサポートしています。

これらのデザインにより、**Goはマルチコアのマシン上で動作する、システムソフトウェアの構築を行える様々な機能を提供しています**。

Goができるより詳細な内容については[どのような言語なのか](#どのような言語なのか)を参照

## なにに使われているのか

### 活躍している分野

- 分散ネットワークサービス
- 並列処理の強みがあるGOはマイクロサービスやAPI・WEB開発によく使われています。
- クラウドネイティブ開発
- DockerやKubernetesなどのクラウドネイティブサービスはGO言語で開発されました。
- モバイルアプリケーション開発
- まだGoを使ったモバイルアプリのプロジェクトは少ないですが、Android・iosアプリケーションをGoで開発できます。

### 公開事例

- [メルカリ](https://www.mercari.com/jp/)

    - 「メルカリ アッテ」や「メルペイ」でGoを使用

    - [Goを「教育」で伝える。メルペイエンジニア2人のプログラミング言語談義](https://mercan.mercari.com/articles/15164/)

- [ピクシブ](https://www.pixiv.net/en/)
    
    - 広告配信サーバーでGoを使用

    - [Golang-ads-deliver](https://speakerdeck.com/catatsuy/golang-ads-deliver)

- [Gunosy](https://gunosy.com/)

    - 広告配信システムでGoを使用

    - [テックブログ](https://tech.gunosy.io/archive/category/Go)

- [ぐるなび](https://www.gnavi.co.jp/)

    - ユーザーの好みに合う飲食店の広告を表示する「属性バナー」という機能を実装するために使用

    - [ぐるなびにおけるGo言語の活用](https://developers.gnavi.co.jp/entry/golang-batch-tool-server)
