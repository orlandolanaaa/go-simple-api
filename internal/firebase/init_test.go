package firebase

import (
	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	"context"
	"testing"
)

func TestFirebase_NewService(t *testing.T) {
	type fields struct {
		Ctx             context.Context
		Storage         *cloud.Client
		FireStoreClient *firestore.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Firebase
		wantErr bool
	}{
		{
			name:    "init-firebase",
			fields:  fields{},
			args:    args{ctx: context.Background()},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Firebase{
				Ctx:             tt.fields.Ctx,
				Storage:         tt.fields.Storage,
				FireStoreClient: tt.fields.FireStoreClient,
			}
			_, err := s.NewService(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
