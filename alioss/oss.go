package alioss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/url"
)

const (
	IMAGE_SIZE_SMALL  = 180
	IMAGE_SIZE_MEDIAN = 360
	IMAGE_SIZE_LARGE  = 720
	expiredInSec      = 604800 //7 days
	IMAGE_FORMAT_PNG  = "png"
	IMAGE_FORMAT_JPEG = "jpeg"
)

type OssClient struct {
	Provider    string
	Endpoint    string
	AccessKey   string
	SecretKey   string
	ImageBucket string
}

var Client *OssClient

func InitialOssClient(Provider, ImageBucket, AccessKey, SecretKey, Endpoint string) {
	Client = &OssClient{Provider: Provider, ImageBucket: ImageBucket, AccessKey: AccessKey, SecretKey: SecretKey, Endpoint: Endpoint}
}

func (client OssClient) getClient() (*oss.Client, error) {
	ossClient, err := oss.New(client.Endpoint, client.AccessKey, client.SecretKey)
	if err != nil {
		return nil, err
	}
	return ossClient, nil
}

type Photo struct {
	FullPath string `json:"fullPath"`

	URL string `json:"url"`
}

func (client OssClient) GetDefaultImageURL(url string) (*Photo, error) {
	return client.GetSizeImageURL(url, IMAGE_SIZE_MEDIAN)
}

func (client OssClient) GetSizeImageURL(fullpath string, size int) (*Photo, error) {
	ossClient, err := client.getClient()
	bucket, err := ossClient.Bucket(client.ImageBucket)
	if err != nil {
		return nil, nil
	}
	process := fmt.Sprintf("image/resize,m_lfit,l_%d/format,%s", size, IMAGE_FORMAT_JPEG)
	signedURL, err := bucket.SignURL(fullpath, oss.HTTPGet, expiredInSec, oss.Process(process))
	if err != nil {
		return nil, nil
	}
	imageURL, err := url.PathUnescape(signedURL)
	if err != nil {
		return nil, nil
	}
	return &Photo{
		FullPath: fullpath,
		URL:      imageURL,
	}, nil
}
