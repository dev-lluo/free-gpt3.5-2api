package common

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ErrorResponse(c *gin.Context, code int, msg interface{}, err interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"detail": struct {
			Code  int         `json:"code"`
			Msg   interface{} `json:"msg"`
			Error interface{} `json:"error"`
		}{
			Code:  code,
			Msg:   msg,
			Error: err,
		},
	})
	return
}

// GetTimestampSecond 获取当前时间戳 + 指定 秒
func GetTimestampSecond(second int) int64 {
	return time.Now().Add(time.Second * time.Duration(second)).Unix()
}

func ParseUrl(link string) *url.URL {
	if link == "" {
		return nil
	}
	u, err := url.Parse(link)
	if err != nil {
		return nil
	}
	return u
}

func GetOrigin(link string) string {
	u := ParseUrl(link)
	if u == nil {
		return ""
	}
	return u.Scheme + "://" + u.Host
}

func Struct2BytesBuffer(v interface{}) (*bytes.Buffer, error) {
	data := new(bytes.Buffer)
	err := json.NewEncoder(data).Encode(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Struct2Bytes(v interface{}) ([]byte, error) {
	// 创建一个jsonIter的Encoder
	configCompatibleWithStandardLibrary := jsoniter.ConfigCompatibleWithStandardLibrary
	// 将结构体转换为JSON文本并保持顺序
	bytes_, err := configCompatibleWithStandardLibrary.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes_, nil
}

func SplitAndAddBearer(authTokens string) []string {
	var authTokenList []string
	for _, v := range strings.Split(authTokens, ",") {
		authTokenList = append(authTokenList, "Bearer "+v)
	}
	return authTokenList
}

func RandomLanguage() string {
	// 初始化随机数生成器
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	// 语言列表
	//languages := []string{"af", "am", "ar-sa", "as", "az-Latn", "be", "bg", "bn-BD", "bn-IN", "bs", "ca", "ca-ES-valencia", "cs", "cy", "da", "de", "de-de", "el", "en-GB", "en-US", "es", "es-ES", "es-US", "es-MX", "et", "eu", "fa", "fi", "fil-Latn", "fr", "fr-FR", "fr-CA", "ga", "gd-Latn", "gl", "gu", "ha-Latn", "he", "hi", "hr", "hu", "hy", "id", "ig-Latn", "is", "it", "it-it", "ja", "ka", "kk", "km", "kn", "ko", "kok", "ku-Arab", "ky-Cyrl", "lb", "lt", "lv", "mi-Latn", "mk", "ml", "mn-Cyrl", "mr", "ms", "mt", "nb", "ne", "nl", "nl-BE", "nn", "nso", "or", "pa", "pa-Arab", "pl", "prs-Arab", "pt-BR", "pt-PT", "qut-Latn", "quz", "ro", "ru", "rw", "sd-Arab", "si", "sk", "sl", "sq", "sr-Cyrl-BA", "sr-Cyrl-RS", "sr-Latn-RS", "sv", "sw", "ta", "te", "tg-Cyrl", "th", "ti", "tk-Latn", "tn", "tr", "tt-Cyrl", "ug-Arab", "uk", "ur", "uz-Latn", "vi", "wo", "xh", "yo-Latn", "zh-Hans", "zh-Hant", "zu"}
	languages := []string{"en-US,en;q=0.9"}
	// 随机选择一个语言
	randomIndex := rng.Intn(len(languages))
	return languages[randomIndex]
}

// GetAbsPathAndGenerate 获取绝对路径并生成文件或文件夹
func GetAbsPathAndGenerate(path string, isFilePath bool, content string) string {
	// 获取绝对路径
	path = GetAbsPath(path)
	if isFilePath {
		//	判断文件是否存在
		if isExist := fileIsExistAndCreat(path, content); isExist {
			return path
		}
	} else {
		//	判断文件夹是否存在
		if isExist := dirIsExistAndMkdir(path, false); isExist {
			return path
		}
	}
	return path
}

// GetAbsPath 获取绝对路径
func GetAbsPath(path string) string {
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return ""
		}
		return absPath
	}
	return path
}

func dirIsExistAndMkdir(dirPath string, isFile bool) bool {
	// 判断路径是否存在
	_, err := os.Stat(dirPath)
	dir := dirPath
	if err != nil {
		if isFile {
			dir = filepath.Dir(dirPath)
		}
		// 创建路径
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

func fileIsExistAndCreat(filePath string, content string) bool {
	//判断文件是否存在
	_, err := os.Stat(filePath)
	if err != nil {
		// 判断文件夹是否存在
		if isExist := dirIsExistAndMkdir(filePath, true); !isExist {
			return false
		}
		// 创建文件
		_, err := os.Create(filePath)
		if err != nil {
			return false
		}
		if content != "" {
			//	写入content
			file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
			_, _ = file.Write([]byte(content))
			defer func(file *os.File) {
				_ = file.Close()
			}(file)
		}
	}
	return true
}
