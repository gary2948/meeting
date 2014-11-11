package main

import (
	. "commonPackage"
	"hash/crc32"
	"io"
	"os"

	. "github.com/qiniu/api/conf"
	. "github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
)

var (
	bucket    string
	upString  = "hello qiniu world"
	policy    rs.PutPolicy
	localFile = "io_api.go"
	key1      = "test_put_1"
	key2      = "test_put_2"
	key3      = "test_put_3"
	extra     = []*PutExtra{
		&PutExtra{
			MimeType: "text/plain",
			CheckCrc: 0,
		},
		&PutExtra{
			MimeType: "text/plain",
			CheckCrc: 1,
		},
		&PutExtra{
			MimeType: "text/plain",
			CheckCrc: 2,
		},
		nil,
	}
)

func init() {
	bucket = BucketName
	if ACCESS_KEY == "" || SECRET_KEY == "" {
		panic("require test env")
	}
	Println(ACCESS_KEY)
	Println(SECRET_KEY)

	policy.Scope = bucket
}

func crc32File(file string) uint32 {

	//it is so simple that do not check any err!!
	f, _ := os.Open(file)
	defer f.Close()
	info, _ := f.Stat()
	h := crc32.NewIEEE()
	buf := make([]byte, info.Size())
	io.ReadFull(f, buf)
	h.Write(buf)
	return h.Sum32()
}

func crc32String(s string) uint32 {

	h := crc32.NewIEEE()
	h.Write([]byte(s))
	return h.Sum32()
}

func testPutFileWithoutKey(localFile string) (key, hash string) {

	ret := new(PutRet)
	err := PutFile(nil, ret, UptokenHelper(), "test.pdf", localFile, nil)
	CheckError(err)
	//for _, v := range extra {
	//	if v != nil {
	//		v.Crc32 = crc32File(localFile)
	//	}

	//	err := PutFileWithoutKey(nil, ret, UptokenHelper(), localFile, v)

	//	if err != nil {
	//		panic(err)
	//	}
	//}

	return ret.Key, ret.Hash
}

func main() {
	Println("test upload qiniu")
	//k, h := testPutFileWithoutKey(`c:\test1.pdf`)
	//Println("key:" + k)
	//Println("hash:" + h)
	durl := DownloadUrl(`test.pdf`)
	Println(durl)

}
