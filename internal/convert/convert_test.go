package convert

import (
	"os"
	"testing"
)

func TestConvertImagesToPDF(t *testing.T) {
	inputDir := "../../files/images_to_pdfs"
	outputDir := "../../files/pdf"

	_ = os.MkdirAll(inputDir, os.ModePerm)
	_ = os.MkdirAll(outputDir, os.ModePerm)

	err := ConvertImagesToPDF(inputDir, outputDir)
	if err != nil {
		t.Fatalf("图片转 PDF 失败: %v", err)
	}
}
