package repository

import (
	"sync"
)

type Post struct {
	Id         int64  `json_test:"id"`
	ParentId   int64  `json_test:"parent_id"`
	Content    string `json_test:"content"`
	CreateTime int64  `json_test:"create_time"`
}
type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}
func (*PostDao) QueryPostsByParentId(parentId int64) []*Post {
	return postIndexMap[parentId]
}
