# 🥷 Ninja School Online (NSO) Server - Go Implementation

> **Một backend server dành cho game NSO được viết lại bằng Golang.**  
---

## 🚀 Tính năng chính

- ✅ Server socket dạng `TCP`
- ✅ Gửi và nhận `Message XOR`
- ✅ Quản lý session người chơi
- ✅ Hỗ trợ đa server
- ✅ Auto migrate và seed database

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
