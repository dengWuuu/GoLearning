package repository

import (
	"sync"
)

type Topic struct {
	Id         int64  `json_test:"id"`
	Title      string `json_test:"title"`
	Content    string `json_test:"content"`
	CreateTime int64  `json_test:"create_time"`
}
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
