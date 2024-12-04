package bucket

import (
	"fmt"
	"log"
	"net"
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

  minioHost:= os.Getenv("MINIO_HOST")
  minioPort:= os.Getenv("MINIO_PORT")
	minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
  minioSecretKey := os.Getenv("MINIO_SECRET_KEY")

 if bucketHost == "" {
		ip, err := getMachineIP()
		if err != nil {
			log.Fatalf("failed to get machine IP address: %v", err)
		}

  log.Printf("bucket host: %s", ip)
		bucketHost = ip 
	}



	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
    Endpoint:         aws.String(fmt.Sprintf("http://%s:%s", minioHost, minioPort)),
		S3ForcePathStyle: aws.Bool(true),


		Credentials:      credentials.NewStaticCredentials(minioAccessKey, minioSecretKey, ""),
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

func getMachineIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	// Loop over all interfaces to find a valid one with an IPv4 address
	for _, iface := range interfaces {
		// Skip down interfaces
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			// Check for IPv4 addresses (ignore IPv6)
			if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no valid IPv4 address found")
}
