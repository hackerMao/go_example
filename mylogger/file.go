package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

/**
往文件里写日志
*/

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileName string
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

var chanMaxSize int = 50000

func NewFileLogger(level, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, chanMaxSize),
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	go fl.writeLogBackground()
	return fl
}

func (f *FileLogger) initFile() error {
	logFileName := path.Join(f.filePath, f.fileName)
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file fieled, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile("error_"+logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open error log file fieled, err:%v\n", err)
		return err
	}
	f.fileObj = file
	f.errFileObj = errFileObj

	return nil
}

func (f *FileLogger) CLose() {
	defer f.fileObj.Close()
	defer f.errFileObj.Close()
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info fieled, err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitLog(file *os.File) (*os.File, error) {
	// 备份并rename log文件
	timestamp := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("open log file fieled, err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak.%s", logName, timestamp)
	// 关闭旧日志文件
	_ = file.Close()
	_ = os.Rename(logName, newLogName)
	// 打开新日志文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file fieled, err:%v\n", err)
		return nil, err
	}
	return fileObj, nil
}

func (f *FileLogger) writeLogBackground() {
	for {
		// 判断是否需要切割文件
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitLog(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}

		select {
		case logBody := <-f.logChan:
			// 拼凑日志
			logMsg := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n",
				logBody.timestamp,
				getLevelString(logBody.level),
				logBody.fileName,
				logBody.funcName,
				logBody.line,
				logBody.msg)
			_, _ = fmt.Fprintf(f.fileObj, logMsg)

			// 如果要记录的日志大于等于ERROR级别，还要往err日志文件中记录
			if logBody.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitLog(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				_, _ = fmt.Fprintf(f.errFileObj, logMsg)
			}
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (f *FileLogger) log(lv LogLevel, format string, arg ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, arg...)
		timeFmt := time.Now().Format("2006-01-02 15:04:05")
		funcName, filename, lineNo := getInfo(3)

		logStruct := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  filename,
			timestamp: timeFmt,
			line:      lineNo,
		}
		select {
		case f.logChan <- logStruct:
		default:
			// 把日志丢掉保证业务不阻塞
		}
	}
}

func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.Level
}

func (f *FileLogger) Debug(format string, arg ...interface{}) {
	f.log(DEBUG, format, arg...)
}

func (f *FileLogger) Trace(format string, arg ...interface{}) {
	f.log(TRACE, format, arg...)
}

func (f *FileLogger) Info(format string, arg ...interface{}) {
	f.log(INFO, format, arg...)
}

func (f *FileLogger) Warn(format string, arg ...interface{}) {
	f.log(WARN, format, arg...)
}

func (f *FileLogger) Error(format string, arg ...interface{}) {
	f.log(ERROR, format, arg...)
}

func (f *FileLogger) Fatal(format string, arg ...interface{}) {
	f.log(FATAL, format, arg...)
}
