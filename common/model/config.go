package model

type Configuration struct {
	AppName   string    `json:"appName" yaml:"appName" default:"gtool"`
	LogConfig LogConfig `json:"logConfig" yaml:"logConfig"`
	Port      string    `json:"port" yaml:"port" default:"9092"`
	DbConfig  DbConfig  `json:"dbConfig" yaml:"dbConfig"`
}

type DbConfig struct {
	DbType          string `json:"dbType" yaml:"dbType" default:"sqlite"`
	Dsn             string `json:"dsn" yaml:"dsn" default:"gtool.db"`
	Username        string `json:"username" yaml:"username"`
	Password        string `json:"password" yaml:"password"`
	SkipTransaction bool   `json:"skipTransaction" yaml:"skipTransaction" default:"false"`
}

type LogConfig struct {
	Folder   string `json:"folder" yaml:"folder" default:"./logs/"`
	Filename string `json:"filename" yaml:"filename" default:"app.log"`
	Level    string `json:"level" yaml:"level"  default:"info"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`
}
