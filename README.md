# 🥷 Ninja School Online (NSO) Server - Go Implementation

> **Một backend server dành cho game NSO được viết lại bằng Golang.**  
> ⚠️ Dự án chỉ mang tính chất học tập và nghiên cứu kỹ thuật, **không nhằm mục đích thương mại hoặc phá hoại**.

---

## 🧠 Mục tiêu

- Mô phỏng lại backend server game Ninja School Online
- Tái tạo luồng login, session, nhân vật, map, skill...
- Kiến trúc theo hướng **clean architecture**, dễ mở rộng và bảo trì
- Viết bằng Golang, hiệu suất cao, tối ưu đa kết nối

---

## 📁 Cấu trúc thư mục

```
nso-server/
│
├── cmd/                    # Entry point (main.go)
│
├── internal/               # Code nội bộ, chia module rõ ràng
│   ├── app/                # Bootstrap, khởi tạo app, seed, migrate
│   ├── config/             # Load config từ .env và config.json
│   ├── infra/              # Kết nối DB, logger, các tiện ích hạ tầng
│   ├── lang/               # Hệ thống đa ngôn ngữ (vi, en, ...)
│   ├── model/              # Các model GORM (Account, Character, ...)
│   ├── net/                # Server TCP, session, router
│   │   ├── handler/        # Xử lý logic theo command
│   │   │   ├── not_login/  # Xử lý login, register, info ban đầu
│   │   │   ├── sub_command/# Các command sau đăng nhập
│   │   └── ...
│   └── proto/              # Đọc/ghi message theo giao thức NSO
│
├── data/                   # Dữ liệu seed (json)
├── logs/                   # Log server
├── .env                    # Thông tin kết nối DB
├── .gitattributes          # Kiểm soát EOL (LF/CRLF)
├── .editorconfig           # Quy tắc định dạng code
├── go.mod / go.sum         # Module Golang
└── README.md               # Tài liệu này
```

---

## ⚠️ Miễn trừ trách nhiệm

- **Dự án không có liên quan đến TeaMobi hoặc bất kỳ tổ chức nào sở hữu NSO.**
- Mục đích của project là để học hỏi về:  
  Golang, TCP server, game protocol, kiến trúc phần mềm...
- **Không sử dụng vào mục đích thương mại, phá hoại hoặc cheat.**

---

## 🧪 Tiến trình hiện tại

- [x] Handshake + mã hóa XOR  
- [x] Đăng nhập, đăng ký tài khoản  
- [x] Quản lý session theo user  
- [ ] Danh sách nhân vật  
- [ ] Logic chọn nhân vật, vào map  
- [ ] Di chuyển + đồng bộ mob

---

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

---

## 📜 License

This project is licensed under the [Creative Commons Attribution-NonCommercial 4.0 International (CC BY-NC 4.0)](https://creativecommons.org/licenses/by-nc/4.0/) license.

> This means you're free to use, share, and modify the code for **non-commercial purposes**, as long as you give appropriate credit.

[![License: CC BY-NC 4.0](https://img.shields.io/badge/License-BY--NC%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc/4.0/)
