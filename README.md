# AirMeet WebAPI

API一覧
- [イベントの登録](#register_event)
- [イベントの削除](#remove_event)
- [イベントの情報取得](#event_info)
- [イベントへのユーザの登録](#register_user)
- [イベント参加者の取得](#participants)
- [ユーザ登録の削除](#remove_user)

## <a id ="register_event">イベントの登録</a>
親機はイベント名、会場名、説明文、子機に入力させる必須項目などを設定し、サーバに送る
### リクエストURL
http://******/register_event

メソッド:POST

### リクエストボディ
|キー|必須|説明|
|:--|:--:|:--|
|event_name|○|イベント名|
|room_name||会場名|
|description||説明文|
|items|○|ユーザに入力させる必須項目<br>","区切り|


### レスポンスボディ(JSON形式)
#### 成功時
```json
{
    "result":{
      "major": 数字,
      "message": "なんかメッセージ",
    },
    "code": 200
}
```

#### 失敗時(クライアントエラー)
```json
{
    "message": "なんかメッセージ",
    "code": 400
}
```

#### 失敗時(サーバエラー)
```json
{
    "message": "なんかメッセージ",
    "code": 500
}
```

## <a id ="remove_event">イベントの削除</a>
### リクエストURL
http://******/remove_event

メソッド:POST

### リクエストボディ
|キー|必須|説明|
|:--|:--:|:--|
|major|○|親機のmajorの値|

### レスポンスボディ(JSON形式)
#### 成功時
```json
{

    "result": "なんか成功って感じのメッセージ",
    "code": 200
}
```

#### 失敗時(クライアントエラー)
```json
{
    "message": "なんかメッセージ",
    "code": 400
}
```

#### 失敗時(サーバエラー)
```json
{
    "message": "なんかメッセージ",
    "code": 500
}
```


## <a id ="event_info">イベントの情報取得
### リクエストURL
http://******/event_info

メソッド:GET

### リクエストクエリパラメータ
|キー|必須|説明|
|:--|:--:|:--|
|major|○|親機からiBeaconで取得したmajorの値|

### レスポンスボディ(JSON形式)
#### 成功時
```json
{
    "result": {
      "event":{
        "event_name": "イベント名",
        "room_name": "ルーム名",
        "description": "説明文",
        "major": ユニークな数列,
        "items": [
            "親機側が設定した必須項目（配列形式）",
            "親機側が設定した必須項目（配列形式）"
        ],
      },
      "count": 現在の参加者数,
      "message": "メッセージ",
    },
    "code": 200
}
```

#### 失敗時(クライアントエラー)
```json
{
    "message": "なんかメッセージ",
    "code": 400
}
```

#### 失敗時(サーバエラー)
```json
{
    "message": "なんかメッセージ",
    "code": 500
}
```


## <a id ="register_user">イベントへのユーザの登録</a>

### リクエストURL
http://******/users

メソッド:POST

### リクエストボディ (multipart/form-data形式)
|キー|必須|説明|
|:--|:--:|:--|
|major|○|親機からiBeaconで取得したmajorの値|
|name|○|名前|
|profile||プロフィール|
|items|○|項目|
|image||プロフィール画像|
|image_header||ヘッダー画像|

### レスポンスボディ(JSON形式)
#### 成功時
```json
{
    "result": {
      "id": "自分のid",
      "message": "なんか成功って感じのメッセージ",
    },
    "code": 200
}
```

#### 失敗時(クライアントエラー)
```json
{
    "message":"なんかメッセージ",
    "code":400
}
```

#### 失敗時(サーバエラー)
```json
{
    "message":"なんかメッセージ",
    "code":500
}
```

## <a id ="participants">イベント参加者の取得</a>
### リクエストURL
http://******/participants

メソッド:GET

### リクエストクエリパラメータ
|キー|必須|説明|
|:--|:--:|:--|
|major|○|親機からiBeaconで取得したmajorの値|
|id|○|自分のID|


### レスポンスボディ(JSON形式)
#### 成功時
```json
{
    "result":{
      "participants": [  
          {
              "id": "参加している他のユーザのid",
              "name": "参加している他のユーザの名前",
              "profile": "プロフィール",
              "items": {
                  "": "",
                  "": "",
              },
              "image": "プロフィール画像のURL",
              "image_header": "ヘッダー画像のURL"
          },
          {
              "id": "参加している他のユーザのid",
              "name": "参加している他のユーザ",
              "profile": "プロフィール",
              "items": {
                  "": "",
                  "": "",
              },
              "image": "プロフィール画像のURL",
              "image_header": "ヘッダー画像のURL"
          }
      ],
      "count" : "イベントの参加者数（自分除く）",
      "message": "成功メッセージ",
    },
    "code": 200
}
```

#### 失敗時(クライアントエラー)
```json
{
    "message":"なんかメッセージ",
    "code":400
}
```

#### 失敗時(サーバエラー)
```json
{
    "message":"なんかメッセージ",
    "code":500
}
```


## <a id ="remove_user">ユーザ登録の削除</a>
### リクエストURL
http://******/remove_user

メソッド:POST

### リクエストボディ
|キー|必須|説明|
|:--|:--:|:--|
|id|○|削除したいユーザのID|

### レスポンスボディ(JSON形式)
#### 成功時
```json
{
    "result":{
      "id": "削除したユーザのID",
      "message": "なんか成功って感じのメッセージ",
    },
    "code": 200
}
```

#### 失敗時(クライアントエラー)
```json
{
    "message": "なんかメッセージ",
    "code": 400
}
```

#### 失敗時(サーバエラー)
```json
{
    "message": "なんかメッセージ",
    "code": 500
}
```
