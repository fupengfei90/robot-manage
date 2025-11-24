package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// 生成bcrypt密码哈希的工具脚本
// 使用方法: go run sql/generate_password.go <password>
func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run sql/generate_password.go <password>")
		fmt.Println("示例: go run sql/generate_password.go admin123456")
		os.Exit(1)
	}

	password := os.Args[1]
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("生成密码哈希失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("密码: %s\n", password)
	fmt.Printf("bcrypt哈希值: %s\n", string(hash))
	fmt.Printf("\nSQL插入语句:\n")
	fmt.Printf("INSERT INTO `manage_users` (`username`, `password`, `email`, `role`, `status`) VALUES\n")
	fmt.Printf("  ('your_username', '%s', 'your_email@example.com', 'user', 1);\n", string(hash))
}

