package convert

import (
	"fmt"
	"image"
	_ "image/jpeg" // 支持 JPEG
	_ "image/png"  // 支持 PNG
	"os"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf/v2"
)

// ConvertImagesToPDF 批量将图片转换为 PDF
func ConvertImagesToPDF(inputDir, outputDir string) error {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("无法读取目录 %s: %v", inputDir, err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			fmt.Printf("跳过不支持的文件类型: %s\n", file.Name())
			continue
		}

		// 创建 PDF
		err := convertImageToPDF(filepath.Join(inputDir, file.Name()), filepath.Join(outputDir, strings.TrimSuffix(file.Name(), ext)+".pdf"))
		if err != nil {
			fmt.Printf("转换失败: %s, 错误: %v\n", file.Name(), err)
			continue
		}
		fmt.Printf("已成功转换: %s\n", file.Name())
	}
	return nil
}

// convertImageToPDF 将单张图片转换为 PDF
func convertImageToPDF(imagePath, outputPath string) error {
	imgFile, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("无法打开图片文件 %s: %v", imagePath, err)
	}
	defer imgFile.Close()

	_, format, err := image.Decode(imgFile)
	if err != nil {
		return fmt.Errorf("图片解码失败: %v", err)
	}

	_, err = imgFile.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("无法重置文件指针: %v", err)
	}

	supportedFormats := map[string]string{
		"jpeg": "JPG",
		"png":  "PNG",
	}
	imageType, ok := supportedFormats[format]
	if !ok {
		return fmt.Errorf("不支持的图片格式: %s", format)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	opts := gofpdf.ImageOptions{
		ImageType: imageType,
	}

	// width, height := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())
	// pageWidth, pageHeight := 210.0, 297.0 // A4 尺寸

	// // 根据比例缩放
	// scale := min(pageWidth/width, pageHeight/height)
	// width, height = width*scale, height*scale

	// x, y := (pageWidth-width)/2, (pageHeight-height)/2
	x, y, width, height := 0.0, 0.0, 0.0, 0.0
	pdf.ImageOptions(imagePath, x, y, width, height, false, opts, 0, "")

	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		return fmt.Errorf("PDF 保存失败: %v", err)
	}
	return nil
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
