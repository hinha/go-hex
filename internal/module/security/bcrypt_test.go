package security

import (
	"fmt"
	"testing"
)

func Test_hash_Password(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name     string
		args     args
		wantErr bool
		hashedPwd func() string
	} {
		{
			name: "success",
			args: args{
				email:    "clyf@email.com",
				password: "hashedpassword",
			},
			hashedPwd: func() string {
				generate := GeneratePasswordHash([]byte("hashedpassword"))
				return generate
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isHashed := tt.hashedPwd()

			fmt.Println(tt.wantErr)
			if (isHashed == "") != tt.wantErr {
				t.Errorf("bcrypt.Generate()  wantErr %v", tt.wantErr)
				return
			}
		})
	}
}
