package dgkala

import (
	"reflect"
	"testing"
)

func Test_sendRequest(t *testing.T) {
	type args struct {
		address string
		headers map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test should send request and receive a 200 respose",
			args: args{
				address: "http://icanhazip.com/",
				headers: map[string]string{"ApplicationVersion": "1.3.2"},
			},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sendRequest(tt.args.address, tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.StatusCode, tt.want) {
				t.Errorf("sendRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncredibleOffers(t *testing.T) {
	tests := []struct {
		name     string
		wantType string
		wantErr  bool
	}{
		{
			name:     "Test should return a slice of incredible offers",
			wantType: "[]dgkala.IncredibleOffer",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IncredibleOffers()
			if (err != nil) != tt.wantErr {
				t.Errorf("IncredibleOffers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.TypeOf(got).String() != tt.wantType {
				t.Errorf("type of IncredibleOffers() = %v, want type of %v", reflect.TypeOf(got), tt.wantType)
			}
		})
	}
}
