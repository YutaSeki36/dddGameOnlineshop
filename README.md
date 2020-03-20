# dddGameOnlineshop
DDDしてみる．ゲームのオンラインショップをモデリングしてみる

## 概要

「ドメイン駆動設計入門 ボトムアップでわかる！ドメイン駆動設計の基本」(https://www.shoeisha.co.jp/book/detail/9784798150727)
を読んだので，DDDを右も左もわからないけど早速アウトプットしてみる．
まずはできるだけシンプルに．

題材は「ゲームのオンラインショップ」
ニンテンドーE ShopやPS store的なもの

モデイングができたらGoでコード書いてみる

## 題材
「ゲームのオンラインショップ」

ユーザはオンラインでゲームを購入することができる．購入したゲームはダウンロードして遊ぶことができる

## モデリング

### ユースケース
- ゲーム検索　タイトル，ジャンル，メーカから検索できる
- ゲーム詳細　ゲームの詳細を確認できる
- ゲーム購入
- ゲームのダウンロード(購入されたゲームをDLできる)
- ユーザ登録
- ユーザログイン
- ユーザ情報修正


### ドメインモデル
- ゲーム -> ライフサイクルがらある．一意な値で識別されるため，エンティティとして定義する
  - タイトル -> 文字数などの制限があるため値オブジェクトとして定義
  - ジャンル -> ホラーやアクションなど，特定の文字列を持つため，値オブジェクトとする
  - 発売日 -> 発売日に何かロジックがあるかどうか検討
  - ゲームメーカーIDのリスト -> ゲームメーカーエンティティのIDだけを持つ.
  
- ゲームメーカー
  - 名前
  - 創立日
  - 開発ゲーム一覧(ゲームIDのリスト)

- ユーザ
