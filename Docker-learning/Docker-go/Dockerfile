# Chọn image Golang chính thức
FROM golang:1.24-alpine AS builder

# Set thư mục làm việc trong container
WORKDIR /app

# Sao chép các file Go và module
COPY go.mod go.sum ./
RUN go mod tidy

# Sao chép mã nguồn
COPY . .

# Biên dịch ứng dụng Go
RUN go build -o myapp .

# Tạo image cuối cùng (chạy ứng dụng)
FROM alpine:latest

# Cài đặt thư viện cần thiết
RUN apk --no-cache add ca-certificates

# Sao chép ứng dụng từ container builder vào container chạy ứng dụng
COPY --from=builder /app/myapp /myapp

# Định nghĩa lệnh chạy container
ENTRYPOINT ["/myapp"]