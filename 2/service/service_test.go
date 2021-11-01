package service

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		pagination string
		searchword string
	}
	tests := []struct {
		name    string
		args    args
		want    Data
		wantErr bool
	}{
		{
			name: "Get Success",
			args: args{
				pagination: "2",
				searchword: "batman",
			},
			wantErr: false,
			want:    Data{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.pagination, tt.args.searchword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
