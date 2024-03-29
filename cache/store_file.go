package cache

import (
	"encoding/json"
	"github.com/curatorc/cngf/app"
	"github.com/curatorc/cngf/file"
	"github.com/curatorc/cngf/logger"
	"github.com/curatorc/cngf/timer"
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
	Value     string     `json:"value,omitempty"`
	ExpiredAt timer.Time `json:"expired_at"`
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
		ExpiredAt: app.Now().Add(expireTime),
	}
	str, err := json.Marshal(s.Store)
	logger.LogIf(err)
	err = file.Put(str, s.FilePath)
	logger.LogIf(err)
}

func (s *FileStore) Get(key string) (value string) {
	if content, ok := s.Store[key]; ok {
		if content.ExpiredAt.After(app.Now()) {
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
	err = json.Unmarshal(file.Get(s.FilePath, "{}"), &s.Store)
	logger.LogIf(err)
	return err
}
