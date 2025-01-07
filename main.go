package main

import (
	"fmt"
	"image-pdf-tool/internal/convert"
	"image-pdf-tool/internal/merge"
	"os"
	"path/filepath"
	"strings"
)

const (
	ImagesToPdfsDir    = "files/images_to_pdfs"
	FilesMergeToPdfDir = "files/files_merge_to_pdf"
	OutputPdfDir       = "files/pdf"
)

func main() {
	// 初始化文件夹
	initializeDirectories()

	// 命令行界面
	fmt.Println("欢迎使用 IMAGE-PDF 工具！")
	fmt.Println("请选择功能：")
	fmt.Println("1. 将图片转换为对应的PDF")
	fmt.Println("2. 合并图片和PDF，为一个PDF")
	fmt.Println("请输入选项（1 或 2）：")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		convertImagesToPDF()
	case 2:
		mergeFilesToPDF()
	case 0:
		fmt.Println("退出程序")
		return
	default:
		fmt.Println("无效选择，请重新运行程序")
	}
}

func convertImagesToPDF() {

	fmt.Printf("请将待转换的图片放入 %s 文件夹中\n", ImagesToPdfsDir)
	fmt.Println("生成的 PDF 文件会保存在", OutputPdfDir, "文件夹中")
	fmt.Println("按 Y 开始转换，其他键退出")

	var confirm string
	fmt.Scan(&confirm)
	if strings.ToLower(confirm) != "y" {
		fmt.Println("退出程序")
		return
	}

	files, err := os.ReadDir(ImagesToPdfsDir)
	if err != nil {
		fmt.Println("读取文件夹失败：", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("文件夹中没有文件")
		return
	}

	fmt.Printf("检测到 %d 张图片\n", len(files))
	if err := convert.ConvertImagesToPDF(ImagesToPdfsDir, OutputPdfDir); err != nil {
		fmt.Println("转换失败：", err)
		return
	}
	fmt.Println("转换完成, 请查看", OutputPdfDir, "文件夹")
}

func mergeFilesToPDF() {
	fmt.Printf("请将待合并的图片和 PDF 文件放入 %s 文件夹中\n", FilesMergeToPdfDir)
	fmt.Println("生成的 PDF 文件会保存在", OutputPdfDir, "中")
	fmt.Println("按 Y 开始合并，其他键退出")

	var confirm string
	fmt.Scan(&confirm)
	if strings.ToLower(confirm) != "y" {
		fmt.Println("操作取消")
		return
	}

	outputFile := filepath.Join(OutputPdfDir, "merged.pdf")
	if err := merge.MergeFilesToPDF(FilesMergeToPdfDir, outputFile); err != nil {
		fmt.Println("合并失败：", err)
		return
	}
	fmt.Println("合并完成, 请查看", OutputPdfDir, "文件夹")
}

// 初始化文件夹
func initializeDirectories() {
	dirs := []string{ImagesToPdfsDir, FilesMergeToPdfDir, OutputPdfDir}
	for _, dir := range dirs {
		initializeDirectory(dir)
	}
}

func initializeDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("创建文件夹失败：%s\n", dir)
			return
		}
		fmt.Printf("已创建文件夹：%s\n", dir)
	}
}
