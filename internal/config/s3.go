package config

type S3 struct {
	Endpoint          string
	Region            string
	Bucket            string // 所有图片都存储在这个bucket下，不同组织的图片存储在不同的目录下
	AutoCreateBucket  bool
	AccessID          string
	AccessKey         string
	PathStyle         bool
	FilePreviewExpire int // 文件预览时间，单位秒
}
