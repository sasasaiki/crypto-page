package eth

import (
	"testing"
)

func TestGetTransactionList(t *testing.T) {
	type args struct {
		id   string
		page string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ちゃんと返る",
			args: args{
				id:   "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a",
				page: "1",
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetTransactionList(tt.args.id, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransactionList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetBalance(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *transaction
		wantErr bool
	}{
		{
			name: "ちゃんと返る",
			args: args{
				id: "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetBalance(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetBalance() = %v, want %v", got, tt.want)
			// }
		})
	}
}
