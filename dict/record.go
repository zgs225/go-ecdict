package dict

import (
	"fmt"
)

// Record 词典记录
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

func (r Record) String() string {
	return fmt.Sprintf(`Record = {
		Word = "%s",
		Phonetic = "%s",
		Definition = "%s",
		Translation = "%s",
		Pos = "%s",
		Collins = "%s",
		Oxford = "%s",
		Tag = "%s",
		Bnc = %d,
		Frq = %d,
		Exchange = "%s",
}`, r.Word, r.Phonetic, r.Definition, r.Translation, r.Pos, r.Collins, r.Oxford, r.Tag, r.Bnc, r.Frq, r.Exchange)
}
