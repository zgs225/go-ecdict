go-ecdict
===

基于[ECDICT](https://github.com/skywind3000/ECDICT)库的数据实现单次的查询功能。

目前实现的查询方式有：

+ `SimpleDict` 使用完全常驻在内存中的有序列表作为索引，用二分法查找。索引比较大，需要 `70M ~ 80M`
  的内存

## 使用方式

``` go
import "github.com/zgs225/go-ecdict/dict"

var dictFile string // 从 ECDICT 下载 ecdict.csv 文件的路径
dict, err := dict.NewSimpleDict(dictFile)
record, err := dict.Match("hello")
records, err := dict.Like("hell)
```

### 查询结果结构体

``` go
type Record struct {
	Word        string // 单词名称
	Phonetic    string // 音标，以英语英标为主
	Definition  string // 单词释义（英文），每行一个释义
	Translation string // 单词释义（中文），每行一个释义
	Pos         string // 词语位置，用 "/" 分割不同位置
	Collins     string // 柯林斯星级
	Oxford      string // 是否是牛津三千核心词汇
	Tag         string // 字符串标签：zk/中考，gk/高考，cet4/四级 等等标签，空格分割
	Bnc         int    // 英国国家语料库词频顺序
	Frq         int    // 当代语料库词频顺序
	Exchange    string // 时态复数等变换，使用 "/" 分割不同项目，见后面表格
	Detail      string // json 扩展信息，字典形式保存例句（待添加）
	Audio       string // 读音音频 url （待添加）
}
```

## Benchmark

| 名称 | 次数 | 平均耗时 | 平均内存占用 | 平均内存分配 |
|--- |--- |--- |--- |--- |
|BenchmarkSimpleDict_Match-4 |116396 |8676 ns/op |5154 B/op |14 allocs/o |
|BenchmarkSimpleDict_Like_4Characters-8|1488|699324 ns/op|568970 B/op|1583 allocs/op |