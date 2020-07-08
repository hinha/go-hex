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

func Test_compare_Password(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name     string
		args     args
		wantErr bool
		hashedPwd func() error
	} {
		{
			name: "hashedSecret too short to be a bcrypted password",
			args: args{
				password: "hashedpassword",
			},
			wantErr: true,
			hashedPwd: func() error {
				var ar args
				password := "TEST"
				return PasswordCompare([]byte(ar.password), []byte(password))
			},
		},
		{
			name: "Invalid user credentials",
			args: args{
				password: "hashedpassword",
			},
			wantErr: true,
			hashedPwd: func() error {
				var ar args
				generate := GeneratePasswordHash([]byte(ar.password))
				ar.password = "not password"
				ab := PasswordCompare([]byte(ar.password), []byte(generate))
				return ab
			},

		},
		{
			name: "success",
			args: args{
				password: "hashedpassword",
			},
			hashedPwd: func() error {
				var ar args
				generate := GeneratePasswordHash([]byte(ar.password))
				ab := PasswordCompare([]byte(ar.password), []byte(generate))
				return ab
			},

		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.hashedPwd()
			fmt.Println(err)
			if (err != nil) != tt.wantErr {
				t.Errorf("bcrypt.compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}