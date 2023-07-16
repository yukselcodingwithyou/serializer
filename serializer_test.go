package serializer

import (
	"reflect"
	"testing"
)

func Test_jsonSerializer_Serialize(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "serializes to json",
			args: args{value: struct {
				Name string `json:"name"`
			}{
				Name: "yuksel",
			}},
			want:    `{"name":"yuksel"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonSerializer{}
			got, err := j.Serialize(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Serialize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlSerializer_Serialize(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := yamlSerializer{}
			got, err := y.Serialize(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Serialize() got = %v, want %v", got, tt.want)
			}
		})
	}
}
