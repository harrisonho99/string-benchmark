package concat

import (
	"bytes"
	"concat_string/builder"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

type check_allocation map[uintptr]bool

func UseConcat(n int, isShowResult bool) {

	var checker = check_allocation{}
	var str = ""
	p_string_header := (*reflect.StringHeader)(unsafe.Pointer(&str))

	for i := 0; i < n; i++ {
		str += strconv.Itoa(i)
		checker[p_string_header.Data] = true
	}
	if isShowResult {
		fmt.Printf("reusult :%q\n", str)
	}
	fmt.Println("---Normal concat---")
	fmt.Printf("num of allocation: %d\n", len(checker))
}

func UseStringBuilder(n int, isShowResult bool) {
	var checker = check_allocation{}
	var b builder.Builder
	b.WriteString("")
	p_Builder := b.GetStringPointer()
	p_string_header := (*reflect.StringHeader)(unsafe.Pointer(p_Builder))
	for i := 0; i < n; i++ {
		b.WriteString(strconv.Itoa(i))
		checker[(*p_string_header).Data] = true
	}
	if isShowResult {
		fmt.Printf("reusult :%q\n", *p_Builder)
	}
	fmt.Println("---String Builder concat---")
	fmt.Printf("num of allocation: %d\n", len(checker))
}

type readOp int8
type Buffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.
}

func UseBytesBuffer(n int, isShowResult bool) {
	var checker = check_allocation{}
	var b = bytes.Buffer{}
	b.WriteString("")
	p_Buffer := (*Buffer)(unsafe.Pointer(&b))
	p_buf := &p_Buffer.buf
	p_buf_header := (*reflect.SliceHeader)(unsafe.Pointer(p_buf))

	for i := 0; i < n; i++ {
		b.WriteString(strconv.Itoa(i))
		checker[(*p_buf_header).Data] = true
	}
	if isShowResult {
		fmt.Printf("reusult :%q\n", *p_buf)
	}
	fmt.Println("---Bytes Buffer concat---")
	fmt.Printf("num of allocation: %d\n", len(checker))
}
