package setting
import (
	"time"
)
type ServerSettings struct {
	RunMode			string
	HttpPort		string
	ReadTimeOut 	time.Duration
	WriteTimeout 	time.Duration
}

type AppSettings struct {
	DefaultPageSize int
	MaxPageSize 	int
	LogSavePath		string
	LogFileName		string
	LogFileExt		string
}

type DatabaseSettings struct {
	DBType			string
	Username		string
	Password		string
	Host			string
	DBName			string
	TablePrefix		string
	Charset			string
	ParseTime		bool
	MaxIdleConns	int
	MaxOpenConns	int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}