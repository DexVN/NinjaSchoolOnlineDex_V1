# 🥷 Ninja School Online (NSO) Server - Go Implementation

> **Một backend server dành cho game NSO được viết lại bằng Golang.**  
> 📚 *Tài liệu và mã nguồn này chỉ mang tính chất học tập và nghiên cứu.*

---

## 📌 Giới thiệu

Đây là một dự án mô phỏng lại server của **Ninja School Online**, một trò chơi nhập vai nổi tiếng tại Việt Nam, được viết lại bằng ngôn ngữ **Go (Golang)** nhằm phục vụ mục đích:

- Học tập kiến trúc server game multiplayer
- Hiểu rõ về hệ thống message/tcp protocol trong client NSO
- Thực hành kỹ năng lập trình backend chịu tải cao

> ⚠️ **Lưu ý**: Đây **không phải** là server chính thức hoặc được phát hành bởi nhà phát hành game gốc. Dự án không khuyến khích sử dụng cho mục đích thương mại.

---

## 🏗 Kiến trúc chính

- TCP server theo kiểu custom protocol (bắt chước server NSO thật)
- Hệ thống phân tích và định tuyến lệnh (Command-based handler)
- Xử lý đăng nhập, đăng ký, session, handshake, mã hóa XOR key
- Tích hợp với PostgreSQL để lưu trữ dữ liệu
- Giao tiếp Binary-based message (tự viết encoder/decoder)

---

## 📂 Cấu trúc thư mục

```
.
├── cmd/                        # Điểm khởi động (entrypoint)
├── data/                       # Dữ liệu JSON để seed DB
│   ├── n_class.json
│   └── skill_option_template.json
├── internal/
│   ├── app/                    # Khởi động app, seed, migrate
│   ├── config/                 # Đọc .env và cấu hình
│   ├── infra/                  # Kết nối PostgreSQL (gorm)
│   ├── model/                  # Khai báo các bảng DB: Account, Character, Server...
│   ├── net/
│   │   ├── handler/            # Xử lý từng Cmd, SubCmd (client gửi lên)
│   │   ├── session.go          # Đọc/gửi message TCP với XOR
│   │   └── server.go           # Khởi tạo và lắng nghe socket
│   ├── proto/                  # Định nghĩa Reader/Writer để decode/encode binary message
│   └── utils/                  # Tiện ích chung
├── scripts/                   # Shell script (migrate DB, tool...)
├── test/                      # File test local
├── .env                       # Cấu hình môi trường
├── go.mod / go.sum            # Go module config
└── main.go                    # Điểm chạy chính của server
```

## 📥 Cài đặt & Chạy

### Yêu cầu

- Go 1.24+
- PostgreSQL 17+

### 1. Clone và cấu hình

```bash
git clone https://github.com/yourname/nso-server.git
cd nso-server
cp .env.example .env
```

### 2. Chạy server

```bash
go run main.go
```

---

## 🧪 Seed dữ liệu

Dữ liệu sẽ được seed tự động khi:

```go
AppEnv = development
```

---

## 🔐 Giao tiếp với Client

- TCP socket với XOR key `E5XXYY...`
- Nhận message dạng `opcode + length + payload`
- Gửi `cmd=-27` (handshake) để bắt đầu phiên
- Xử lý `subCommand` login: -127, register: -126...

---

## 📫 Liên hệ / Đóng góp

> Nếu bạn quan tâm hoặc muốn đóng góp:
> - Issue/Pull Request
> - Email: khanhld.developer@gmail.com
