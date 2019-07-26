package alioss

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Provider    = "aliyun"
	ImageBucket = "shopintar-staging"
	Endpoint    = "https://oss-ap-southeast-5.aliyuncs.com"
	AccessKey   = "LTAIbW13QrwtWOmP"
	//AccessKey   = "accessKey"
	SecretKey = "kD8uT53d3jzT5bzmSvzxXNvtVx0gFE"
)

func init() {
	InitialOssClient(Provider, ImageBucket, AccessKey, SecretKey, Endpoint)
}

func TestGetImageURL(t *testing.T) {
	//objectName := "photos/shopintar-1534588244535-8c24f062d6eb4915.jpeg"
	objectName := "photos/bl/2808bd726a018935bf794c1ba4ece2d3.png"
	image, err := Client.GetSizeImageURL(objectName, 720)
	assert.Nil(t, err)
	fmt.Println(image.URL)
	assert.NotNil(t, image)
	//https://shopintar-staging.oss-ap-southeast-5.aliyuncs.com/photos/shopintar-1534588244535-8c24f062d6eb4915.jpeg?Expires=1561132800&OSSAccessKeyId=LTAIbW13QrwtWOmP&Signature=REjLNSgKLmaiPRUAevflnDWmZuU%3D&x-oss-process=image%2Fresize%2Cm_lfit%2Cl_720%2Fformat%2Cjpeg
}
