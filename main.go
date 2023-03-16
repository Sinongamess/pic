package main

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
)


func main() {
    // 定义包含分类和文件路径的 map
    categories := make(map[string][]string)

    // 遍历文件夹
    err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // 排除文件夹本身
        if path == "." {
            return nil
        }

        // 获取文件分类
        category := filepath.Dir(path)

        // 将文件路径添加到分类数组中
        categories[category] = append(categories[category], path)

        return nil
    })

    if err != nil {
        panic(err)
    }

    // 将 map 转换为 JSON 格式
    categoriesJSON, err := json.MarshalIndent(categories, "", "    ")
    if err != nil {
        panic(err)
    }

    // 创建 JSON 文件
    file, err := os.Create("tree.json")
    if err != nil {
        panic(err)
    }

    defer file.Close()

    // 写入 JSON 数据
    _, err = file.Write(categoriesJSON)
    if err != nil {
        panic(err)
    }

    fmt.Println("JSON 文件已保存")
}
