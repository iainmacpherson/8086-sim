package datastream

import "testing"

func TestDataStreamCreate(t *testing.T) {
	t.Run("create from existing file", func(t *testing.T) {
		filepath := "test_data/data_file"
		ds := DataStreamCreateFromFile(filepath)
		if ds == nil {
			t.Error("Should be able to create DataStream from existing file")
		}
	})

	t.Run("create from nonexistant file", func(t *testing.T) {
		// ignores the expected panic so the test can continue
		defer func() { _ = recover() }()
		filepath := "test_data/no_data_file"
		DataStreamCreateFromFile(filepath)
		t.Error("Should panic when attempting to create from nonexistant file")
	})

	t.Run("create from data", func(t *testing.T) {
		test_data := []byte{0b10101010, 0xDE}
		ds := DataStreamCreateFromData(test_data)
		if ds == nil {
			t.Error("Should be able to create DataStream from provided data")
		}
	})
}

func TestTryPopBits(t *testing.T) {
	test_data := []byte{0b10101010, 0xDE}
	ds := DataStreamCreateFromData(test_data)

	var tests = []struct {
		name               string
		input_bits_to_read uint
		expected_data      uint8
		expected_bits_read uint
	}{
		{"read full byte", 8, test_data[0], 8},
		{"read half last byte", 4, test_data[1] >> 4, 4},
		{"attempt over read", 8, test_data[1] & 0xF, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, bits_read := ds.TryPopBits(test.input_bits_to_read)
			if data != test.expected_data && bits_read != test.expected_bits_read {
				t.Errorf("got data:0x%x bits_read:%d, expected data:0x%x bits_read:%d",
					data, bits_read, test.expected_data, test.expected_bits_read)
			}
		})
	}
}

func TestTryPopByte(t *testing.T) {
	test_data := []byte{0xDE}
	ds := DataStreamCreateFromData(test_data)

	var tests = []struct {
		name             string
		expected_data    uint8
		expected_success bool
	}{
		{"read byte", 0xDE, true},
		{"read byte when empty", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, success := ds.TryPopByte()
			if data != test.expected_data && success != test.expected_success {
				t.Errorf("got data:0x%x success:%t, expected data:0x%x success:%t",
					data, success, test.expected_data, test.expected_success)
			}
		})
	}
}

func TestMixBitAndByteReads(t *testing.T) {
	test_data := []byte{0b10101010, 0xDE}
	ds := DataStreamCreateFromData(test_data)

	t.Run("read single bit", func(t *testing.T) {
		expected_data := test_data[0] >> 7
		bits_to_read := uint(1)
		data, bits_read := ds.TryPopBits(bits_to_read)
		if data != expected_data && bits_read != bits_to_read {
			t.Errorf("got data:0x%x bits_read:%d, expected data:0x%x bits_read:%d",
				data, bits_read, expected_data, bits_to_read)
		}
	})

	t.Run("unaligned byte read", func(t *testing.T) {
		// ignores the expected panic so the test can continue
		defer func() { _ = recover() }()
		ds.TryPopByte()
		t.Error("Unaligned byte read should panic")
	})

	t.Run("read remaining bits", func(t *testing.T) {
		expected_data := test_data[0] & 0b01111111
		bits_to_read := uint(7)
		data, bits_read := ds.TryPopBits(bits_to_read)
		if data != expected_data && bits_read != bits_to_read {
			t.Errorf("got data:0x%x bits_read:%d, expected data:0x%x bits_read:%d",
				data, bits_read, expected_data, bits_to_read)
		}
	})

	t.Run("read byte", func(t *testing.T) {
		expected_data := test_data[1]
		data, success := ds.TryPopByte()
		if data != expected_data && success {
			t.Errorf("got data:0x%x success:%t, expected data:0x%x success:%t",
				data, success, expected_data, true)
		}
	})
}

func TestPopBitsAndByte(t *testing.T) {
	test_data := []byte{0xDE, 0xAD}
	ds := DataStreamCreateFromData(test_data)
	t.Run("pop bits success", func(t *testing.T) {
		expected_data := test_data[0]
		data := ds.PopBits(8)
		if data != expected_data {
			t.Errorf("got data:0x%x, expected data:0x%x", data, expected_data)
		}
	})
	t.Run("pop byte success", func(t *testing.T) {
		expected_data := test_data[1]
		data := ds.PopByte()
		if data != expected_data {
			t.Errorf("got data:0x%x, expected data:0x%x", data, expected_data)
		}
	})
	t.Run("pop bit fail", func(t *testing.T) {
		defer func() { _ = recover() }()
		ds.PopBits(8)
		t.Errorf("should panic on failure to read bits")
	})
	t.Run("pop byte fail", func(t *testing.T) {
		defer func() { _ = recover() }()
		ds.PopByte()
		t.Errorf("should panic on failure to read byte")
	})

}

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		name            string
		expected_result bool
		test_data       []byte
	}{
		{"contains data", false, []byte{0xDE}},
		{"no data", true, []byte{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ds := DataStreamCreateFromData(test.test_data)
			result := ds.IsEmpty()
			if result != test.expected_result {
				t.Errorf("got result:%t, expected result:%t",
					result, test.expected_result)
			}
		})
	}
}
