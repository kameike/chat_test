// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "From the todo list tutorial on goswagger.io",
    "title": "Taimme-Chat",
    "version": "1.0.0"
  },
  "paths": {
    "/auth": {
      "post": {
        "description": "サインアップもしくはアクセストークンの更新を行います",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "ログイン",
        "parameters": [
          {
            "type": "string",
            "description": "他のユーザーから見えないユーザーを特定するハッシュ値です。パスワードのように扱われます。アクセストークンの取得に使用します。",
            "name": "authToken",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "apiサーバー等から払い出されるハッシュ値です。他のユーザーから見えても大丈夫で、推測が難しいものが望ましいです。これでユーザーは一意に特定されるのでuniqである必要もあります。",
            "name": "userHash",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/definitions/authInfo"
          },
          "403": {
            "$ref": "#/definitions/error"
          }
        }
      }
    },
    "/chatrooms": {
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "一覧が頑張るよ",
        "produces": [
          "application/json"
        ],
        "tags": [
          "chatRooms"
        ],
        "summary": "メッセージの一覧が取ってこれるよ",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/chatroomRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/chatroom"
              }
            }
          }
        }
      }
    },
    "/chatrooms/{id}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "一覧が頑張るよ",
        "tags": [
          "chatRooms"
        ],
        "summary": "メッセージの一覧が取ってこれるよ",
        "parameters": [
          {
            "type": "string",
            "description": "長いハッシュ値",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/definitions/chatroomFull"
          }
        }
      }
    },
    "/chatrooms/{id}/messages": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "ok",
        "tags": [
          "chatRooms"
        ],
        "summary": "メッセージの取得",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/chat"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "chatRooms"
        ],
        "summary": "メッセージの投稿",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/chatCreate"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/definitions/chatroomFull"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "長いハッシュ値",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "authInfo": {
      "description": "認証認可用のオブジェクト、x_chat_access_tokenのヘッダに入れて使用します。",
      "type": "object",
      "properties": {
        "accessToken": {
          "type": "string",
          "example": "HAB4cQxKTQkEj7rMdE6QQW391ffpVbQshya+R66OIhfu5drm"
        }
      }
    },
    "chat": {
      "description": "チャットに使われるやつ",
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "text": {
          "type": "string",
          "example": "よろしくお願いします"
        },
        "timestamp": {
          "type": "integer",
          "format": "int64"
        },
        "user": {
          "$ref": "#/definitions/user"
        }
      }
    },
    "chatCreate": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "chatroom": {
      "description": "チャットルームを取得するときに出てくるやつ",
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "participants": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/user"
          }
        },
        "peekedChat": {
          "description": "最大3件メッセージがあればpeekします。",
          "type": "array",
          "items": {
            "$ref": "#/definitions/chat"
          }
        },
        "unreads": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "unreadCount": {
                "type": "integer"
              },
              "userHash": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "chatroomFull": {
      "description": "チャットルームを完全に取得する際にでてくるやつ",
      "type": "object",
      "properties": {
        "chatroom": {
          "$ref": "#/definitions/chatroom"
        },
        "chats": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chat"
          }
        }
      }
    },
    "chatroomRequest": {
      "type": "object",
      "properties": {
        "request": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "errorMessage": {
          "description": "エラーメッセージ、中身をユーザーに見せて差し支えないものです。",
          "type": "string",
          "example": "権限がありません(code: 3)"
        }
      }
    },
    "user": {
      "description": "汎用的に出てくるユーザーオブジェクト",
      "type": "object",
      "properties": {
        "hash": {
          "type": "string",
          "example": "HAB4cQxKTQkEj7rMdE6QQW391ffpVbQshya+R66OIhfu5drm"
        },
        "name": {
          "type": "string",
          "example": "kameike"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "x_chat_access_token",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "From the todo list tutorial on goswagger.io",
    "title": "Taimme-Chat",
    "version": "1.0.0"
  },
  "paths": {
    "/auth": {
      "post": {
        "description": "サインアップもしくはアクセストークンの更新を行います",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "ログイン",
        "parameters": [
          {
            "type": "string",
            "description": "他のユーザーから見えないユーザーを特定するハッシュ値です。パスワードのように扱われます。アクセストークンの取得に使用します。",
            "name": "authToken",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "apiサーバー等から払い出されるハッシュ値です。他のユーザーから見えても大丈夫で、推測が難しいものが望ましいです。これでユーザーは一意に特定されるのでuniqである必要もあります。",
            "name": "userHash",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "認証認可用のオブジェクト、x_chat_access_tokenのヘッダに入れて使用します。"
          },
          "403": {}
        }
      }
    },
    "/chatrooms": {
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "一覧が頑張るよ",
        "produces": [
          "application/json"
        ],
        "tags": [
          "chatRooms"
        ],
        "summary": "メッセージの一覧が取ってこれるよ",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/chatroomRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/chatroom"
              }
            }
          }
        }
      }
    },
    "/chatrooms/{id}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "一覧が頑張るよ",
        "tags": [
          "chatRooms"
        ],
        "summary": "メッセージの一覧が取ってこれるよ",
        "parameters": [
          {
            "type": "string",
            "description": "長いハッシュ値",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "チャットルームを完全に取得する際にでてくるやつ"
          }
        }
      }
    },
    "/chatrooms/{id}/messages": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "ok",
        "tags": [
          "chatRooms"
        ],
        "summary": "メッセージの取得",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/chat"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "tags": [
          "chatRooms"
        ],
        "summary": "メッセージの投稿",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/chatCreate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "チャットルームを完全に取得する際にでてくるやつ"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "長いハッシュ値",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "authInfo": {
      "description": "認証認可用のオブジェクト、x_chat_access_tokenのヘッダに入れて使用します。",
      "type": "object",
      "properties": {
        "accessToken": {
          "type": "string",
          "example": "HAB4cQxKTQkEj7rMdE6QQW391ffpVbQshya+R66OIhfu5drm"
        }
      }
    },
    "chat": {
      "description": "チャットに使われるやつ",
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "text": {
          "type": "string",
          "example": "よろしくお願いします"
        },
        "timestamp": {
          "type": "integer",
          "format": "int64"
        },
        "user": {
          "$ref": "#/definitions/user"
        }
      }
    },
    "chatCreate": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "chatroom": {
      "description": "チャットルームを取得するときに出てくるやつ",
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "participants": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/user"
          }
        },
        "peekedChat": {
          "description": "最大3件メッセージがあればpeekします。",
          "type": "array",
          "items": {
            "$ref": "#/definitions/chat"
          }
        },
        "unreads": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "unreadCount": {
                "type": "integer"
              },
              "userHash": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "chatroomFull": {
      "description": "チャットルームを完全に取得する際にでてくるやつ",
      "type": "object",
      "properties": {
        "chatroom": {
          "$ref": "#/definitions/chatroom"
        },
        "chats": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chat"
          }
        }
      }
    },
    "chatroomRequest": {
      "type": "object",
      "properties": {
        "request": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "errorMessage": {
          "description": "エラーメッセージ、中身をユーザーに見せて差し支えないものです。",
          "type": "string",
          "example": "権限がありません(code: 3)"
        }
      }
    },
    "user": {
      "description": "汎用的に出てくるユーザーオブジェクト",
      "type": "object",
      "properties": {
        "hash": {
          "type": "string",
          "example": "HAB4cQxKTQkEj7rMdE6QQW391ffpVbQshya+R66OIhfu5drm"
        },
        "name": {
          "type": "string",
          "example": "kameike"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "x_chat_access_token",
      "in": "header"
    }
  }
}`))
}
