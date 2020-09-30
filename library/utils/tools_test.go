/**
 * @Author   KevinGunn <cnphp@hotmail.com>
 *
 * @Description //TODO
 * @Version: 1.0.0
 * @Date     2020/9/29
 */

package utils

import (
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

func TestDecryptCBC(t *testing.T) {
	//type args struct {
	//	plainText string
	//	publicKey string
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want string
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := DecryptCBC(tt.args.plainText, tt.args.publicKey); got != tt.want {
	//			t.Errorf("DecryptCBC() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	str1 := DecryptCBC("IJ1xz+Wve+ZONVMFfXJQMw==", AdminCbcPublicKey)
	t.Log(str1)
}

func TestEncryptCBC(t *testing.T) {
	//type args struct {
	//	plainText string
	//	publicKey string
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want string
	//}{
	//	// TODO: Add test cases.
	//
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := EncryptCBC(tt.args.plainText, tt.args.publicKey); got != tt.want {
	//			t.Errorf("EncryptCBC() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	str1 := EncryptCBC("sealms#Admin", AdminCbcPublicKey)
	t.Log(str1)
}

func TestGetCityByIp(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCityByIp(tt.args.ip); got != tt.want {
				t.Errorf("GetCityByIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetClientIp(t *testing.T) {
	type args struct {
		r *ghttp.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClientIp(tt.args.r); got != tt.want {
				t.Errorf("GetClientIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDomain(t *testing.T) {
	type args struct {
		r *ghttp.Request
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDomain(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDomain() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFilesPath(t *testing.T) {
	type args struct {
		fileUrl string
	}
	tests := []struct {
		name     string
		args     args
		wantPath string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, err := GetFilesPath(tt.args.fileUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilesPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPath != tt.wantPath {
				t.Errorf("GetFilesPath() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

func TestGetHourDiffer(t *testing.T) {
	type args struct {
		startTime string
		endTime   string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHourDiffer(tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("GetHourDiffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocalIP(t *testing.T) {
	tests := []struct {
		name    string
		wantIp  string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIp, err := GetLocalIP()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocalIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIp != tt.wantIp {
				t.Errorf("GetLocalIP() gotIp = %v, want %v", gotIp, tt.wantIp)
			}
		})
	}
}

func TestGetRealFilesUrl(t *testing.T) {
	type args struct {
		r    *ghttp.Request
		path string
	}
	tests := []struct {
		name         string
		args         args
		wantRealPath string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRealPath, err := GetRealFilesUrl(tt.args.r, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRealFilesUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRealPath != tt.wantRealPath {
				t.Errorf("GetRealFilesUrl() gotRealPath = %v, want %v", gotRealPath, tt.wantRealPath)
			}
		})
	}
}

func TestStrToTimestamp(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToTimestamp(tt.args.dateStr); got != tt.want {
				t.Errorf("StrToTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeStampToDate(t *testing.T) {
	type args struct {
		timeStamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeStampToDate(tt.args.timeStamp); got != tt.want {
				t.Errorf("TimeStampToDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeStampToDateTime(t *testing.T) {
	type args struct {
		timeStamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeStampToDateTime(tt.args.timeStamp); got != tt.want {
				t.Errorf("TimeStampToDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}