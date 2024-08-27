package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Print("あなたの名前を入力してください: ")
    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    if name == "" {
        fmt.Println("こんにちは、匿名さん！")
    } else {
        fmt.Printf("こんにちは、%sさん！\n", name)
    }
}