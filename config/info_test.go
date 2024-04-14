package config

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
)

func TestGetConf(t *testing.T) {
	os.Setenv(MYSQL_DATABASE, "unittest")
	os.Setenv(MYSQL_USER, "user")
	os.Setenv(MYSQL_PASS, "passwd")
	loc := getLocation()
	tests := []struct {
		name string
		want mysql.Config
	}{
		{
			name: "test ok",
			want: mysql.Config{
				Addr:   "database:3306",
				Net: "tcp",
				User:   "user",
				Passwd: "passwd",
				DBName: "unittest",
				Loc:    loc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConf() = %v, want %v", got, tt.want)
			}
		})
	}
	os.Unsetenv(MYSQL_DATABASE)
	os.Unsetenv(MYSQL_USER)
	os.Unsetenv(MYSQL_PASS)
}

func Test_getPort(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test ok",
			want: "4000",
		},
		{
			name: "test default ok",
			want: "3306",
		},
	}
	for i, tt := range tests {
		if i == 0 {
			os.Setenv(MYSQL_PORT, "4000")
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := getPort(); got != tt.want {
				t.Errorf("getPort() = %v, want %v", got, tt.want)
			}
		})
		if i == 0 {
			os.Unsetenv(MYSQL_PORT)
		}
	}
}

func Test_getHostname(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"test ok",
			"localhost",
		},
		{
			"test default ok",
			"database",
		},
	}
	for i, tt := range tests {
		if i == 0 {
			os.Setenv(MYSQL_HOST, "localhost")
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := getHostname(); got != tt.want {
				t.Errorf("getHostname() = %v, want %v", got, tt.want)
			}
		})
		if i == 0 {
			os.Unsetenv(MYSQL_HOST)
		}
	}
}

func Test_setUser(t *testing.T) {
	os.Setenv(MYSQL_USER, "user")
	conf := mysql.Config{}
	type confArgs struct {
		conf *mysql.Config
	}
	tests := []struct {
		name string
		args confArgs
	}{
		{
			name: "test ok",
			args: confArgs{conf: &conf},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUser(tt.args.conf)
		})
	}
	os.Unsetenv(MYSQL_USER)
}

func Test_setPasswd(t *testing.T) {
	os.Setenv(MYSQL_PASS, "password")
	conf := mysql.Config{}
	type args struct {
		conf *mysql.Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test ok",
			args: args{conf: &conf},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setPasswd(tt.args.conf)
		})
	}
	os.Unsetenv(MYSQL_PASS)
}

func Test_setDatabase(t *testing.T) {
	os.Setenv(MYSQL_DATABASE, "unittest")
	conf := mysql.Config{}
	type args struct {
		conf *mysql.Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test ok",
			args: args{conf: &conf},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setDatabase(tt.args.conf)
		})
	}
	os.Unsetenv(MYSQL_DATABASE)
}

func Test_getLocation(t *testing.T) {
	loc := time.UTC
	tests := []struct {
		name string
		want *time.Location
	}{
		{name: "test ok", want: loc},
		{name: "test default ok", want: getLocation()},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 0 {
				os.Setenv(MYSQL_LOC, "UTC")
			}
			if got := getLocation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLocation() = %v, want %v", got, tt.want)
			}
			if i == 0 {
				os.Unsetenv(MYSQL_LOC)
			}
		})
	}
}

func Test_keyfatal(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyfatal(tt.args.key)
		})
	}
}
