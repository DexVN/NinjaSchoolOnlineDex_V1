package net

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close() // Đảm bảo đóng kết nối khi hàm kết thúc
	fmt.Println("New client connected:", conn.RemoteAddr())

	buf := make([]byte, 1024) // Tạo bộ đệm để đọc dữ liệu
	for {
		n, err := conn.Read(buf) // Đọc dữ liệu từ kết nối
		if err != nil {
			fmt.Println("❌ Disconnected:", conn.RemoteAddr())
			return
		}
		fmt.Printf("📨 Received %d bytes: %x\n", n, buf[:n])
	}
}
