package s3

type s3Interface interface {
}

type s3Repository struct {
}

func NewS3Repository() s3Interface {
	return &s3Repository{}
}
