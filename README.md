# bits
My interpretation of a bit reader/writer for Golang

####BitReader functions:
ReadBits(n int): Reads the next n bits and returns the result stored in an integer.

####BitWriter functions:
WriteUint(n uint, num_bits int): Write a number of bits equal to num_bits from the value represented by n.  If there are not enough bits available to represent n, an error will be thrown.  Bits are queued up in chunks of 8 and then written out as an entire byte when the whole byte is filled.  At the very end, FinishByte() should be called in order to fill up the rest of the last byte and ensure everything is written out.

FinishByte(): Fill in the remaining bits in the current bit queue with 0 and writes the byte out.