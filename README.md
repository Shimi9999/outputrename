# outputrename
CSVファイルのデータを元に複数のファイルを一括リネームします。  
CSVファイルのカラム1にリネーム前、カラム2にリネーム後を記述してください。

このプログラムは[bvcut](https://github.com/Shimi9999/bvcut)の出力ファイルを一括リネームするために作られています。

## Usage
```
outputrename <csvpath>
```

## Example
renamedata.csv
```
output_0.mp4,afterfilename0.mp4
output_1.mp4,afterfilename1.mp4
output_2.mp4,afterfilename2.mp4
output_3.mp4,afterfilename3.mp4
output_4.mp4,afterfilename4.mp4
output_5.mp4,afterfilename5.mp4
output_6.mp4,afterfilename6.mp4
output_7.mp4,afterfilename7.mp4
output_8.mp4,afterfilename8.mp4
output_9.mp4,afterfilename9.mp4
output_10.mp4,afterfilename10.mp4
```

実行
```
> outputrename renamedata.csv
Process completed.
```