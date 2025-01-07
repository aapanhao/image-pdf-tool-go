package merge

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jung-kurt/gofpdf/v2"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergeFilesToPDF(inputDir, outputFile string) error {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("无法读取目录 %s: %v", inputDir, err)
	}

	if len(files) == 0 {
		fmt.Printf("目录 %s 中没有文件\n", inputDir)
		return nil
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	tempFiles := []string{}
	allFiles := []string{}
	defer func() {
		for _, tempFile := range tempFiles {
			_ = os.Remove(tempFile)
		}
	}()

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(inputDir, file.Name())
		ext := strings.ToLower(filepath.Ext(file.Name()))

		switch ext {
		case ".jpg", ".jpeg", ".png":
			tempPDF := filePath + ".pdf"
			err := converImageToPDF(filePath, tempPDF)
			if err != nil {
				fmt.Printf("添加图片失败: %s, 错误: %v\n", file.Name(), err)
				continue
			}
			tempFiles = append(tempFiles, tempPDF)
			allFiles = append(allFiles, tempPDF)
		case ".pdf":
			allFiles = append(allFiles, filePath)
		default:
			fmt.Printf("跳过不支持的文件类型: %s\n", file.Name())
		}
	}

	err = api.MergeCreateFile(allFiles, outputFile, false, nil)
	if err != nil {
		return fmt.Errorf("合并 PDF 失败: %v", err)
	}
	fmt.Printf("已成功合并 %d 个文件到 %s\n", len(tempFiles), outputFile)
	return nil
}

func converImageToPDF(imagePath string, outputPath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	opts := gofpdf.ImageOptions{
		ImageType: strings.ToUpper(strings.TrimPrefix(filepath.Ext(imagePath), ".")),
	}

	img, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("无法打开图片文件 %s: %v", imagePath, err)
	}
	defer img.Close()

	file, _, err := image.Decode(img)
	if err != nil {
		return fmt.Errorf("图片解码失败: %v", err)
	}

	width, height := float64(file.Bounds().Dx()), float64(file.Bounds().Dy())
	pageWidth, pageHeight := 210.0, 297.0
	scale := min(pageWidth/width, pageHeight/height)
	width, height = width*scale, height*scale
	x, y := (pageWidth-width)/2, (pageHeight-height)/2

	pdf.ImageOptions(imagePath, x, y, width, height, false, opts, 0, "")
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		return fmt.Errorf("保存 PDF 文件失败: %v", err)
	}

	return nil

}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
