// func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
// order defines the bit-packing order 0-===== litWidth is the word size in bits,
// in gif file corresponds to the pixel depth, typically 8.
package main

import (
	"compress/lzw"
	"io")

  


type blockReader struct {
	r 		reader  // Input source: implements io.Read and io.ByteReader.
	slice   []byte // Buffer of unread data.
	tmp     [256]byte //Storage for slice 
}

func main(){
	deblockingReader := &blockReader{r: imageFile}
	lzwr := lzw.NewReader(&blockReader{r: d.r}, lzw.LSB, int(litWidth))
	if _, err = io.ReadFull(lzwr, m.Pix); err != nil {
		break
	}
}

func (b *blockReader) Read(p []byte) (int, os.Error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(b.slice) == 0 {
		blockLen, err := b.r.ReadByte()
		if err != nil {
			return 0, err 
		}
		if blockLen == 0 {
			return 0, os.EOF
		}
		b.slice = b.tmp[0:blockLen]
		if _, err = io.ReadFull(b.r, b.slice); err != nil {
			return 0,nil
		}

	}
	n := copy(p, b.slice)
	b.slice = b.slice[n:]
	return n, nil
}