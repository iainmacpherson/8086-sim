package main

import (
	"8086-sim/datastream"
	"testing"
)

func TestDisassembleMov(t *testing.T) {
	var tests = []struct {
		name                 string
		input                []byte
		expected_disassembly []string
	}{
		{
			"single mov to from reg",
			[]byte{0x89, 0xD9},
			[]string{"mov cx, bx"},
		},
		{
			"many mov to from reg",
			[]byte{0x89, 0xD9, 0x88, 0xE5, 0x89, 0xDA, 0x89, 0xDE, 0x89, 0xFB, 0x88, 0xC8, 0x88, 0xED, 0x89, 0xC3, 0x89, 0xF3, 0x89, 0xFC, 0x89, 0xC5},
			[]string{"mov cx, bx", "mov ch, ah", "mov dx, bx", "mov si, bx", "mov bx, di", "mov al, cl", "mov ch, ch", "mov bx, ax", "mov bx, si", "mov sp, di", "mov bp, ax"},
		},
		{
			"complex mov with address calcuations",
			[]byte{0x89, 0xDE, 0x88, 0xC6, 0xB9, 0x0C, 0x00, 0xB9, 0xF4, 0xFF, 0xBA, 0x6C, 0x0F, 0xBA, 0x94, 0xF0, 0x8A, 0x00, 0x8B, 0x1B, 0x8B, 0x56, 0x00, 0x8A, 0x60, 0x04, 0x8A, 0x80, 0x87, 0x13, 0x89, 0x09, 0x88, 0x0A, 0x88, 0x6E, 0x00},
			[]string{"mov si, bx", "mov dh, al", "mov cx, 12", "mov cx, 65524", "mov dx, 3948", "mov dx, 61588", "mov al, [bx + si]", "mov bx, [bp + di]", "mov dx, [bp + 0]", "mov ah, [bx + si + 4]", "mov al, [bx + si + 4999]", "mov [bx + di], cx", "mov [bp + si], cl", "mov [bp + 0], ch"},
		},
		{
			"more complex mov plus accumulator use",
			[]byte{0x8B, 0x41, 0xDB, 0x89, 0x8C, 0xD4, 0xFE, 0x8B, 0x57, 0xE0, 0xC6, 0x03, 0x07, 0xC7, 0x85, 0x85, 0x03, 0x5B, 0x01, 0x8B, 0x2E, 0x05, 0x00, 0x8B, 0x1E, 0x82, 0x0D, 0xA1, 0xFB, 0x09, 0xA1, 0x10, 0x00, 0xA3, 0xFA, 0x09, 0xA3, 0x0F, 0x00},
			[]string{"mov ax, [bx + di + 219]", "mov [si + 65236], cx", "mov dx, [bx + 224]", "mov [bp + di], 7", "mov [di + 901], 347", "mov bp, [5]", "mov bx, [3458]", "mov ax, [2555]", "mov ax, [16]", "mov [2554], ax", "mov [15], ax"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ds := datastream.DataStreamCreateFromData(test.input)
			decoded_instructions := Decode(ds)
			disassembly := Disassemble(decoded_instructions)
			for i := range disassembly {
				if disassembly[i] != test.expected_disassembly[i] {
					t.Errorf("got disassembly '%s', expected '%s'", disassembly[i], test.expected_disassembly[i])
				}
			}
		})
	}

}
