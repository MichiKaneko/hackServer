# 概要
JWT認証のついたREST API

# サーバー起動方法
データベースへの接続が必要になります。

```
go run .
```

# リクエスト
[A] POST   /api/user/register
[B] POST   /api/token
[C] GET    /api/secured/ping

## [A] User作成
```rest
POST http://localhost:8080/api/user/register HTTP/1.1
content-type: application/json

{
"name": "Michi Kaneko",
"email": "mikaneko@go.com",
"password": "123465789"
}
```

## [B] トークン取得
```rest
POST http://localhost:8080/api/token HTTP/1.1
content-type: application/json

{
"email": "mikaneko@go.com",
"password": "123465789"
}
```

## [C] 認証付きリクエスト
```rest
GET http://{{host}}/api/secured/ping HTTP/1.1
content-type: application/json
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiTWljaGkgS2FuZWtvIiwiZW1haWwiOiJtaWthbmVrb0Bnby5jb20iLCJleHAiOjE2NzMyNTg3ODF9.WR9P2p8yvlI5gxbOVEE0C8Gh0w_80FKBuw-zCj3FXPs

```