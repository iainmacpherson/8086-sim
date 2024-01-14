package datastream

import (
	"8086-sim/logger"
	"errors"
	"os"
)

const NAME = "DATASTREAM"

type DataStream struct {
	data   []byte
	offset uint
}

func DataStreamCreateFromFile(filepath string) *DataStream {
	data, err := os.ReadFile(filepath)
	if err != nil {
		logger.LogfErr(NAME, "Error reading input program. Please ensure that the correct path was provided. Use -h for help.\n\n")
		panic(err)
	}

	return &DataStream{
		data:   data,
		offset: 0,
	}
}

func DataStreamCreateFromData(data []byte) *DataStream {
	return &DataStream{
		data:   data,
		offset: 0,
	}
}

func (ds *DataStream) TryPopBits(numBits uint) (uint8, uint) {
	var result uint8
	var startOffset uint = ds.offset
	for i := uint(0); i < numBits; i++ {
		if startOffset+i < uint(len(ds.data)*8) {
			byteOffset := ds.offset / 8
			bitOffset := 7 - ds.offset%8
			bitValue := ds.data[byteOffset] >> bitOffset
			result = (result << 1) | (bitValue & 0b1)
			ds.offset++
		} else {
			// reached the end of data
			break
		}
	}
	return result, ds.offset - startOffset
}

func (ds *DataStream) TryPopByte() (byte, bool) {
	// check byte aligned
	if ds.offset%8 != 0 {
		logger.LogErr(NAME, "Attempting to read a byte from data stream but not byte aligned.")
		panic(errors.New("UnalignedRead"))
	}

	var read_byte byte = 0
	success := false
	byteOffset := ds.offset / 8
	if byteOffset < uint(len(ds.data)) {
		read_byte = ds.data[byteOffset]
		ds.offset += 8
		success = true
	}

	return read_byte, success
}

func (ds *DataStream) PopBits(nr_bits uint) uint8 {
	value, nr_read := ds.TryPopBits(nr_bits)
	if nr_read != nr_bits {
		logger.LogErr(NAME, "Failed to read all required bits")
		panic(errors.New("DataStreamFailedRead"))
	}
	return value
}

func (ds *DataStream) PopByte() uint8 {
	value, success := ds.TryPopByte()
	if !success {
		logger.LogErr(NAME, "Failed to read required byte from stream")
		panic(errors.New("DataStreamFailedRead"))
	}
	return value
}

func (ds *DataStream) IsEmpty() bool {
	var result bool = true
	if ds.offset < uint(len(ds.data)*8) {
		result = false
	}
	return result
}
