package store

import "dict/client/model"

// 定义单词存储方法
type WordStore interface {
	// 存储单词
	Save(word *model.Word) int

	// 查询单词
	Load(word string) *model.Word
}
