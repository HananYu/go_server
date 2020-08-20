package utils

import (
	"fmt"
	"gin/config"
	"github.com/nfnt/resize"
	uuid "github.com/satori/go.uuid"
	"image"
	"image/gif"
	jpeg2 "image/jpeg"
	"image/png"
	"os"
	"strings"
)

//返回内容里面的所有图片地址
func GetImageByString(content string, list []string) []string {
	if len(content) <= config.CommonZero {
		return list
	}
	if !strings.Contains(content, config.CommonImageIndex) {
		return list
	}
	i := strings.Index(content, config.CommonImageIndex)
	content = content[(i + len(config.CommonImageIndex) + config.CommonOne):]
	j := strings.Index(content, config.CommonImageLast)
	img := content[0 : j-config.CommonOne]
	fmt.Println("图片地址----------" + img)
	list = append(list, img)
	content = content[j:]
	//fmt.Println("content内容----------"+content)
	return GetImageByString(content, list)
}

//传入完整的图片路径，压缩图片直接返回访问地址，删除原来图片
func ChangeImage(url string) string {
	file, err := os.Open(url)
	if err != nil {
		return config.CommonNull
	}
	img, nametype, er := image.Decode(file)
	if er != nil {
		//表示不是jpg类型，或者是不可识别类型
		i := strings.LastIndex(file.Name(), ".") //含有sub字段的位置
		nametype = file.Name()[i+config.CommonOne:]
		if nametype == config.CommonPng {
			img, er = png.Decode(file)
		} else if nametype == config.CommonGif {
			img, er = gif.Decode(file)
		} else {
			return config.CommonNull
		}
	}
	name := strings.ReplaceAll(uuid.Must(uuid.NewV4(), err).String(), "-", config.CommonNull) + "." + nametype
	newFile, _ := os.Create(config.SavePathUrl + name)
	_ = jpeg2.Encode(newFile, img, &jpeg2.Options{Quality: 10})
	_ = file.Close()
	err = os.Remove(url)
	if err != nil {
		//删除失败
		fmt.Println("删除失败")
		_ = newFile.Close()
		_ = os.Remove(config.SavePathUrl + name)
		return config.CommonNull
	}
	_ = newFile.Close()
	return config.ServiceUrl + name
}

//压缩图片并返回新的图片地址，修改图片尺寸，保留原来图片
func ChangeImageBySize(url string) string {
	file, err := os.Open(url)
	defer file.Close()
	if err != nil {
		return config.CommonNull
	}
	img, nametype, er := image.Decode(file)
	if er != nil {
		//表示不是jpg类型，或者是不可识别类型
		i := strings.LastIndex(file.Name(), ".") //含有sub字段的位置
		nametype = file.Name()[i+config.CommonOne:]
		if nametype == config.CommonPng {
			img, er = png.Decode(file)
		} else if nametype == config.CommonGif {
			img, er = gif.Decode(file)
		} else {
			return config.CommonNull
		}
	}
	name := strings.ReplaceAll(uuid.Must(uuid.NewV4(), err).String(), "-", config.CommonNull) + "." + nametype
	//高度设置为0，则高度会根据原来宽高比例来设置
	m := resize.Resize(200, 0, img, resize.NearestNeighbor)
	newFile, _ := os.Create(config.SavePathUrl + name)
	defer newFile.Close()
	_ = jpeg2.Encode(newFile, m, &jpeg2.Options{Quality: 10})
	return config.ServiceUrl + name
}
