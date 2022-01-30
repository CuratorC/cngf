package cache

import (
	"encoding/json"
	"github.com/curatorc/cngf/app"
	"github.com/curatorc/cngf/file"
	"github.com/curatorc/cngf/logger"
	"github.com/spf13/cast"
	"os"
	"time"
)

// FileStore 文件信息接口
type FileStore struct {
	Store    map[string]content
	FilePath string
}

type content struct {
	Value     string
	ExpiredAt time.Time
}

func NewFileStore() (fs *FileStore) {
	fs = &FileStore{}
	err := fs.IsAlive()
	logger.LogIf(err)
	return fs
}

func (s *FileStore) Set(key string, value string, expireTime time.Duration) {
	s.Store[key] = content{
		Value:     value,
		ExpiredAt: app.TimenowInTimezone().Add(expireTime),
	}
}

func (s *FileStore) Get(key string) (value string) {
	if content, ok := s.Store[key]; ok {
		if content.ExpiredAt.After(app.TimenowInTimezone()) {
			value = content.Value
		}
	}

	return
}

func (s *FileStore) Has(key string) bool {
	if _, ok := s.Store[key]; ok {
		return true
	}
	return false
}

func (s *FileStore) Forget(key string) {
	delete(s.Store, key)
}

func (s *FileStore) Forever(key string, value string) {
	s.Set(key, value, 99999999)
}

func (s *FileStore) Flush() {
	s.Store = make(map[string]content)
}

func (s *FileStore) Increment(key string) {
	if content, ok := s.Store[key]; ok {
		content.Value = cast.ToString(cast.ToInt64(content.Value) + 1)
	} else {
		Forever(key, "1")
	}
}

func (s *FileStore) Increments(key string, count int64) {
	if content, ok := s.Store[key]; ok {
		content.Value = cast.ToString(cast.ToInt64(content.Value) + count)
	} else {
		Forever(key, cast.ToString(count))
	}
}

func (s *FileStore) Decrement(key string) {
	if content, ok := s.Store[key]; ok {
		content.Value = cast.ToString(cast.ToInt64(content.Value) - 1)
	} else {
		Forever(key, "-1")
	}
}

func (s *FileStore) Decrements(key string, count int64) {
	if content, ok := s.Store[key]; ok {
		content.Value = cast.ToString(cast.ToInt64(content.Value) + count)
	} else {
		Forever(key, cast.ToString(count))
	}
}

func (s *FileStore) IsAlive() error {
	dir := "storage/cache/"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logger.ErrorString("make_model", "s.MkdirAll", err.Error())
		return err
	}
	s.Store = make(map[string]content)
	s.FilePath = dir + "cache.json"
	content := file.Get(s.FilePath)
	if content == nil {
		content = []byte("[]")
	}
	logger.DebugString("cache:isAlive", "content", string(content))
	err = json.Unmarshal(content, &s.Store)
	logger.LogIf(err)
	return err
}
