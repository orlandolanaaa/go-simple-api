package firebase

import (
	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	"context"
	"errors"
	"google.golang.org/api/option"
	"log"
)

// Firebase: struct
type Firebase struct {
	Ctx             context.Context
	Storage         *cloud.Client
	FireStoreClient *firestore.Client
}

func (s *Firebase) NewService(ctx context.Context) (*Firebase, error) {
	s.Ctx = ctx

	opt := option.WithServiceAccountFile("internal/firebase/account-key/be-entry-task-firebase-adminsdk-jcmxl-758397b6f8.json")
	var err error
	s.Storage, err = cloud.NewClient(s.Ctx, opt)

	if err != nil {
		log.Fatalf("Error creating GCS client: %v\n", err)
		return s, errors.New("Error creating GCS client: %v\n")
	}

	return s, nil
}
