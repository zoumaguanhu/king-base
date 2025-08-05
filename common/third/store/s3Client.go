package store

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/constants"
	"net/http"
	"path/filepath"
	"time"
)

type S3Client struct {
	config *S3Config
	client *s3.Client
}

func NewS3Client(c *S3Config) *S3Client {
	s := &S3Client{config: c}
	client, err := s.getR2Client()
	if err != nil {
		logx.Infof("NewS3Client config:%+v err:%v", c, err)
		return nil
	}
	s.client = client
	return s
}

type S3Config struct {
	Url             string `json:"url"`
	AccountId       string `json:"accountId"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BucketName      string `json:"bucketName"`
	PublicURL       string `json:"publicURL"`
}

func (c *S3Client) getR2Client() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.config.AccessKeyId, c.config.AccessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		logx.Errorf("getR2Client err:%v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf(c.config.Url, c.config.AccountId))
	})
	return client, nil
}

// 生成唯一文件名
func (c *S3Client) GenerateUniqueKey(originalName string) string {
	ext := filepath.Ext(originalName)
	//name := strings.TrimSuffix(originalName, ext)
	// 这里可以使用更复杂的逻辑生成唯一文件名，如UUID
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}

// 生成文件访问URL
func (c *S3Client) GenerateFileURL(objectKey string) string {
	// 根据你的R2配置返回正确的URL
	// 1. 如果使用自定义域名
	// return "https://images.yourdomain.com/" + objectKey

	// 2. 如果使用R2公共访问URL
	return c.config.PublicURL + "/" + objectKey
}

// 处理文件上传请求
func (c *S3Client) UploadHandler(w http.ResponseWriter, r *http.Request, imgField string) (string, error) {
	// 解析multipart表单
	err := r.ParseMultipartForm(1 << 20) // 限制最大文件大小为1MB
	if err != nil {
		logx.Errorf("upload err:%v", err)
		return constants.EMPTY_STRING, err
	}

	// 获取上传的文件
	file, handler, err := r.FormFile(imgField)
	if err != nil {
		logx.Errorf("upload err:%v", err)
		return constants.EMPTY_STRING, err
	}
	defer file.Close() // 确保文件流最终关闭

	// 生成唯一的文件名（避免覆盖）
	objectKey := c.GenerateUniqueKey(handler.Filename)

	// 直接将文件流上传到R2，无需临时文件
	ct := handler.Header.Get("Content-Type")
	_, err = c.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        &c.config.BucketName,
		Key:           &objectKey,
		Body:          file,                            // 直接使用HTTP请求中的文件流
		ContentLength: &handler.Size,                   // 从上传信息中获取文件大小
		ContentType:   &ct,                             // 设置正确的MIME类型
		ACL:           types.ObjectCannedACLPublicRead, // 根据需要设置ACL
	})

	if err != nil {
		logx.Errorf("upload err:%v", err)
		return constants.EMPTY_STRING, err
	}

	// 生成文件访问URL（假设已配置自定义域名或公共访问）
	fileURL := c.GenerateFileURL(objectKey)

	// 返回成功响应
	return fileURL, nil
}
