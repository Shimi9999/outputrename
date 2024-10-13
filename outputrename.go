// CSVを元に複数のファイルを一括リネームする
// カラム1:リネーム前, カラム2:リネーム後

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: outputrename <csvpath>")
		os.Exit(1)
	}

	csvpath := os.Args[1]

	file, err := os.Open(csvpath)
	if err != nil {
		fmt.Println("File open error:", err.Error())
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	var line []string
	charsToReplace := [][]string{{"\\", "￥"}, {"/", "／"}, {":", "："}, {"*", "＊"}, {"?", "？"},
		{"\"", "”"}, {"<", "＜"}, {">", "＞"}, {"|", "｜"}} // ファイル名に使用できない文字を置き換える

	for {
		line, err = reader.Read()
		if err == io.EOF {
			fmt.Println("Process completed.")
			break
		} else if err != nil {
			fmt.Println("Read line error:", err.Error())
		} else {
			for _, char := range charsToReplace {
				line[1] = strings.Replace(line[1], char[0], char[1], -1)
			}
			err = os.Rename(line[0], line[1])
			if err != nil {
				fmt.Println("Rename error:", err.Error())
			}
		}
	}
}
