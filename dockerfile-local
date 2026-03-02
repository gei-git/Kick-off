# ====================== Builder Stage ======================
FROM golang:1.23-alpine AS builder

# 设置 Go 模块代理（新加坡/中国用户必备，加速下载）
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# 先复制 go.mod 和 go.sum，利用 Docker 层缓存（关键！）
COPY go.mod go.sum ./
RUN go mod download

# 再复制源代码
COPY . .

# 编译静态二进制（生产推荐）
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /todo-api ./cmd/api

# ====================== Runtime Stage ======================
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /root/
COPY --from=builder /todo-api .

EXPOSE 4567

CMD ["./todo-api"]