package bucket

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// declare bucket
type Bucket struct {
	Svc  *s3.S3
	Name string
	Url  string
}

var S3Bucket *Bucket

func Setup() {
	bucketHost := os.Getenv("BUCKET_HOST")
	bucketPort := os.Getenv("BUCKET_PORT")

	bucketName := os.Getenv("BUCKET_NAME")

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Endpoint:         aws.String("http://localhost:9000"),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials("admin", "password", ""),
	})

	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}

	log.Printf("S3 Session created")

	svc := s3.New(sess)
	_, err = svc.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Printf("Bucket %s does not exist, creating it...\n", bucketName)
		_, err = svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}

		log.Printf("Waiting for bucket %s to be created...\n", bucketName)
		err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			log.Fatalf("Error waiting for bucket to exist: %v", err)
		}

	}

	S3Bucket = &Bucket{
		Svc:  svc,
		Name: bucketName,
		Url:  fmt.Sprintf("http://%s:%s/%s", bucketHost, bucketPort, bucketName),
	}

	log.Printf("Bucket %s is ready to use.\n", bucketName)
}
