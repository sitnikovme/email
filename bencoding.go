package email

import (
	"bytes"
	"encoding/base64"
	"io"
)

// bencodeLen finds the len to encode (base on the one from 'mime' package)
func bencodeLen(s string) (last int) {
	len := 0
	r:= []rune(s)

	for i, b := range r {
		if (b < ' ' || b > '~') && b != '\t' {
			len = i + 1
		}
	}
	return len
}

// bencode encodes using B (base64) style of the mime encoding for headers.
func bencode(s string) string {
	blen := bencodeLen(s)
	if blen == 0 {
		return s
	}

	buf := &bytes.Buffer{}
	buf.WriteString("=?UTF-8?B?")

	enc := base64.NewEncoder(base64.StdEncoding, buf)
	r := []rune(s)
	io.WriteString(enc, string(r[:blen]))
	enc.Close()

	buf.WriteString("?=")

	buf.WriteString(string(r[blen:]))

	return buf.String()
}