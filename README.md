# dddGameOnlineshop
DDDしてみる．ゲームのオンラインショップをモデリングしてみる

## 概要

「ドメイン駆動設計入門 ボトムアップでわかる！ドメイン駆動設計の基本」(https://www.shoeisha.co.jp/book/detail/9784798150727)
を読んだので，DDDを右も左もわからないけど早速アウトプットしてみる．
まずはできるだけシンプルに．軽量DDDになってもよい

題材は「ゲームのオンラインショップ」
ニンテンドーE ShopやPS store的なもの

モデリングができたらGoでコード書いてみる

## 題材
「ゲームのオンラインショップ」

ユーザはオンラインでゲームを購入することができる．購入したゲームはダウンロードして遊ぶことができる

## モデリング

### ユースケース
- ゲーム検索　タイトル，ジャンル，メーカから検索できる
- ゲーム詳細　ゲームの詳細を確認できる
- ゲーム購入
- ゲームの登録 ゲーム情報を追加する
- ゲームのダウンロード(購入されたゲームをDLできる)
- ユーザ登録
- ユーザログイン
- ユーザ情報修正


### ドメインモデル
- ゲーム -> ライフサイクルがらある．一意な値で識別されるため，エンティティとして定義する
  - タイトル -> 文字数などの制限があるため値オブジェクトとして定義
  - ジャンル -> ホラーやアクションなど，特定の文字列を持つため，値オブジェクトとする
  - 発売日 -> 発売日に何かロジックがあるかどうか検討
  - 価格 -> 制限あるため値オブジェクトとして持つ
  - ゲームメーカーIDのリスト -> ゲームメーカーエンティティのIDだけを持つ.
  - 対象年齢 -> https://www.jp.playstation.com/cero/ を見るとロジックを持ちそうなので，値オブジェクトにした方がよい？？
  
- ゲームメーカー
  - 名前
  - 創立日
  - 開発ゲーム一覧(ゲームIDのリスト)

- ユーザ

### ユースケース


## アーキテクチャ
実装はClean Architectureを意識

![CleanArchitecture](https://user-images.githubusercontent.com/25469830/79748569-a35ebb80-8348-11ea-9c1e-d9647baf77b9.jpg)

抽象に依存
