package error

import (
	"fmt"
	"testing"
)

func TestInternalServerError(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Internal-Server-Err", args: args{msg: "Internal Server Err"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InternalServerError(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("InternalServerError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotFound(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Internal-Server-Err", args: args{msg: "Not Found"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NotFound(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("NotFound() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRespError_Error(t *testing.T) {
	type fields struct {
		Code    int
		Message string
		CodeMsg string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Error-Custom", fields: fields{
			Code:    400,
			Message: "Error Custom",
			CodeMsg: "Error Custom",
		}, want: fmt.Sprintf("%d: %s", 400, "Error Custom")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RespError{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				CodeMsg: tt.fields.CodeMsg,
			}
			if got := r.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
