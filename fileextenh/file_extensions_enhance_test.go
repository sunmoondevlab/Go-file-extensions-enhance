package fileextenh

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleBase_withExtA() {
	f := Base("testdata/exists.test.a", true)
	fmt.Println(f)
	// Output: exists.test.a
}

func ExampleBase_withoutExtA() {
	f := Base("testdata/exists.test.a", false)
	fmt.Println(f)
	// Output: exists.test
}

func ExampleBase_withExtTarGz() {
	f := Base("testdata/exists.test.tar.gz", true)
	fmt.Println(f)
	// Output: exists.test.tar.gz
}

func ExampleBase_withoutExtTarGz() {
	f := Base("testdata/exists.test.tar.gz", false)
	fmt.Println(f)
	// Output: exists.test

}

func ExampleExt_a() {
	f := Ext("testdata/exists.test.a")
	fmt.Println(f)
	// Output: .a
}

func ExampleExt_gz() {
	f := Ext("testdata/exists.test.gz")
	fmt.Println(f)
	// Output: .gz
}

func ExampleExt_tarGz() {
	f := Ext("testdata/exists.test.tar.gz")
	fmt.Println(f)
	// Output: .tar.gz
}

func ExampleMatchFileExt_matchA() {
	f := MatchFileExt("testdata/exists.test.a", ".a")
	fmt.Println(f)
	// Output: true
}

func ExampleMatchFileExt_matchTarGz() {
	f := MatchFileExt("testdata/exists.test.tar.gz", ".tar.gz")
	fmt.Println(f)
	// Output: true
}

func ExampleMatchFileExt_notMatchTarGz() {
	f := MatchFileExt("testdata/exists.test.tar.gz", ".tgz")
	fmt.Println(f)
	// Output: false
}

func TestBase(t *testing.T) {
	type args struct {
		fname   string
		withExt bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a with ext",
			args: args{
				fname:   "testdata/exists.test.a",
				withExt: true,
			},
			want: "exists.test.a",
		},
		{
			name: "gz with ext",
			args: args{
				fname:   "testdata/exists.test.gz",
				withExt: true,
			},
			want: "exists.test.gz",
		},
		{
			name: "tar.gz with ext",
			args: args{
				fname:   "testdata/exists.test.tar.gz",
				withExt: true,
			},
			want: "exists.test.tar.gz",
		},
		{
			name: "a without ext",
			args: args{
				fname:   "testdata/exists.test.a",
				withExt: false,
			},
			want: "exists.test",
		},
		{
			name: "gz without ext",
			args: args{
				fname:   "testdata/exists.test.gz",
				withExt: false,
			},
			want: "exists.test",
		},
		{
			name: "tar.gz without ext",
			args: args{
				fname:   "testdata/exists.test.tar.gz",
				withExt: false,
			},
			want: "exists.test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base(tt.args.fname, tt.args.withExt); got != tt.want {
				t.Errorf("Base() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExt(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a",
			args: args{
				fname: "testdata/exists.test.a",
			},
			want: ".a",
		},
		{
			name: "gz",
			args: args{
				fname: "testdata/exists.test.gz",
			},
			want: ".gz",
		},
		{
			name: "tar.gz",
			args: args{
				fname: "testdata/exists.test.tar.gz",
			},
			want: ".tar.gz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ext(tt.args.fname); got != tt.want {
				t.Errorf("GetExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchFileExt(t *testing.T) {
	type args struct {
		fname string
		ext   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "match ext .a",
			args: args{
				fname: "testdata/exists.test.a",
				ext:   ".a",
			},
			want: true,
		},
		{
			name: "match ext .tar.gz",
			args: args{
				fname: "testdata/exists.test.tar.gz",
				ext:   ".tar.gz",
			},
			want: true,
		},
		{
			name: "not match ext .tar.gz",
			args: args{
				fname: "testdata/exists.test.tar.gz",
				ext:   ".tgz",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchFileExt(tt.args.fname, tt.args.ext); got != tt.want {
				t.Errorf("MatchFileExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDblExts(t *testing.T) {
	type args struct {
		ext string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: ".gz",
			args: args{
				ext: ".gz",
			},
			want: []string{".tar.gz"},
		},
		{
			name: ".bz2",
			args: args{
				ext: ".bz2",
			},
			want: []string{".tar.bz2"},
		},
		{
			name: ".Z",
			args: args{
				ext: ".Z",
			},
			want: []string{".tar.Z"},
		},
		{
			name: ".xz",
			args: args{
				ext: ".xz",
			},
			want: []string{".tar.xz"},
		},
		{
			name: ".lz",
			args: args{
				ext: ".lz",
			},
			want: []string{".tar.lz"},
		},
		{
			name: ".lzma",
			args: args{
				ext: ".lzma",
			},
			want: []string{".tar.lzma"},
		},
		{
			name: ".lzo",
			args: args{
				ext: ".lzo",
			},
			want: []string{".tar.lzo"},
		},
		{
			name: ".zst",
			args: args{
				ext: ".zst",
			},
			want: []string{".tar.tzst"},
		},
		{
			name: ".a",
			args: args{
				ext: ".a",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDblExts(tt.args.ext); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDblExts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPossibleDoubleExt(t *testing.T) {
	type args struct {
		ext string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: ".gz",
			args: args{
				ext: ".gz",
			},
			want: true,
		},
		{
			name: ".bz",
			args: args{
				ext: ".bz2",
			},
			want: true,
		},
		{
			name: ".Z",
			args: args{
				ext: ".Z",
			},
			want: true,
		},
		{
			name: ".xz",
			args: args{
				ext: ".xz",
			},
			want: true,
		},
		{
			name: ".lz",
			args: args{
				ext: ".lz",
			},
			want: true,
		},
		{
			name: ".lzma",
			args: args{
				ext: ".lzma",
			},
			want: true,
		},
		{
			name: ".lzo",
			args: args{
				ext: ".lzo",
			},
			want: true,
		},
		{
			name: ".zst",
			args: args{
				ext: ".zst",
			},
			want: true,
		},
		{
			name: ".a",
			args: args{
				ext: ".a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleDoubleExt(tt.args.ext); got != tt.want {
				t.Errorf("PossibleDoubleExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
