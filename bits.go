package main

import (
    "bufio"
    "log"
    "math"
)

type BitReader struct {
    bufio_reader *bufio.Reader
    offset int
    curr_byte byte
}

type BitWriter struct {
    bufio_writer *bufio.Writer
    offset int
    value byte
}

func (bitreader *BitReader) readBit() (bit bool, err error) {
    var e error
    if bitreader.offset == 0 {
        bitreader.offset = 8
        bitreader.curr_byte, e = bitreader.bufio_reader.ReadByte()
        next_bit := bitreader.curr_byte & (1 << uint(bitreader.offset-1)) > 0
        bitreader.offset--
        return next_bit, e
    } else {
        next_bit := bitreader.curr_byte & (1 << uint(bitreader.offset-1)) > 0
        bitreader.offset--
        return next_bit, e
    }
}

func (bitreader *BitReader) ReadBits(n int) (value int) {
    v := 0
    if n > 64 {
        log.Fatal("Cannot read more than 64 bits at once")
    }
    for i:=0; i<n; i++ {
        next_bit, err := bitreader.readBit()
        if err != nil {
            log.Fatal("Error reading bit")
        }
        if next_bit {
            v = 2*v + 1
        } else {
            v = 2*v
        }
    }
    return v
}

func (bitwriter *BitWriter) writeBit(bit bool) (err error) {
    var e error
    if bit {
        bitwriter.value = bitwriter.value<<1 + 1
    } else {
        bitwriter.value = bitwriter.value<<1
    }
    bitwriter.offset++
    if bitwriter.offset == 8 {
        e = bitwriter.bufio_writer.WriteByte(bitwriter.value)
        e = bitwriter.bufio_writer.Flush()
        bitwriter.value = byte(0x00)
        bitwriter.offset = 0
    }
    return e
}

func (bitwriter *BitWriter) WriteUint(n uint, num_bits int) {
    var e error
    if n >= uint(math.Pow(2,float64(num_bits))) {
        log.Fatal("Need more bits")
    }
    for i:=num_bits; i>0; i-- {
        if n >= uint(math.Pow(2,float64(i-1))) {
            e = bitwriter.writeBit(true)
            n -= uint(math.Pow(2,float64(i-1)))
        } else {
            e = bitwriter.writeBit(false)
        }
        if e != nil {
            log.Fatal("Error writing bits")
        }
    }
}

func (bitwriter *BitWriter) FinishByte() {
    for bitwriter.offset != 0 {
        bitwriter.writeBit(false)
    }
}