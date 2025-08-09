FROM golang:1.24.6

WORKDIR /src

COPY ./app/go.mod ./app/go.mod
COPY ./app/go.sum ./app/go.sum
WORKDIR /src/app
# github.com/[あなたの名前]/appモジュールの依存関係を全てコンテナ内にダウンロード
RUN go mod download
# srcディレクトリ配下にローカルのソースコードをコピー
COPY ./ ./

WORKDIR /src/app

# 必要なツールをインストール

RUN go install github.com/air-verse/air@latest
RUN go install go.uber.org/mock/mockgen@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# airを起動
CMD ["air"]