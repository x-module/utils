/**
* Created by GoLand
* @file image.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/20 11:30 上午
* @desc image.go
 */

package utils

import (
	"bytes"
	"encoding/base64"
	"github.com/disintegration/imaging"
	"io/ioutil"
	"os"
)

type ImageUtil struct {
	width     int
	height    int
	sourceImg string
	targetImg string
}

func (i *ImageUtil) SetImage(sourceImg string, targetImg string) *ImageUtil {
	i.sourceImg = sourceImg
	i.targetImg = targetImg
	return i
}
func (i *ImageUtil) SetSize(width int, height int) *ImageUtil {
	i.height = height
	i.width = width
	return i
}
func (i *ImageUtil) Resize() error {
	imgData, err := ioutil.ReadFile(i.sourceImg)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		return err
	}
	image = imaging.Resize(image, i.width, i.height, imaging.Lanczos)
	err = imaging.Save(image, i.targetImg)
	return err
}

// TransStrToImage base64 字符串转图片
func TransStrToImage(sourceString string, imageName string) error {
	dist, err := base64.StdEncoding.DecodeString(sourceString)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(imageName, dist, os.ModePerm)
	return err
}

func GetImgBase64String(image string) (string, error) {
	bit, err := ioutil.ReadFile(image)
	if err != nil {
		return "", err
	}
	imgStr := base64.StdEncoding.EncodeToString(bit)
	return imgStr, nil
}

// Average 球平局数
func Average(xs []float64) (avg float64) {
	sum := 0.0
	if len(xs) == 0 {
		avg = 0
	} else {
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return
}
