package merge

import (
	"os"
	"testing"
)

func TestMergeFilesToPDF(t *testing.T) {
	inputDir := "../../files/files_merge_to_pdf"
	outputFile := "../../files/pdf/merged.pdf"

	_ = os.MkdirAll(inputDir, os.ModePerm)

	err := MergeFilesToPDF(inputDir, outputFile)
	if err != nil {
		t.Fatalf("文件合并失败: %v", err)
	}
}
