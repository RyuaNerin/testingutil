package testingutil

import (
	"testing"
)

type CipherSize struct {
	Name string
	Size int
}

// Test All
func TA(
	t *testing.T,
	tests []CipherSize,
	do func(t *testing.T, bitSize int),
	skip bool,
) {
	if skip {
		t.Skip()
		return
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			do(t, test.Size)
		})
	}
}

// Bench All
func BA(
	b *testing.B,
	tests []CipherSize,
	do func(b *testing.B, bitSize int),
	skip bool,
) {
	if skip {
		b.Skip()
		return
	}

	for _, test := range tests {
		test := test
		b.Run(test.Name, func(b *testing.B) {
			do(b, test.Size)
		})
	}
}

// Block TestCase
type BlockTestCase struct {
	Reverse bool
	Key     string // hex
	Plain   string // hex
	Secure  string // hex
	IV      string // initialization vector, in hex

	KeyBytes    []byte
	PlainBytes  []byte
	SecureBytes []byte
	IVBytes     []byte
}

func (btc *BlockTestCase) parse() {
	if btc.KeyBytes == nil {
		btc.KeyBytes = h2b(btc.Key)
		if btc.Reverse {
			btc.KeyBytes = reverse(btc.KeyBytes)
		}
	}
	if btc.PlainBytes == nil {
		btc.PlainBytes = h2b(btc.Plain)
		if btc.Reverse {
			btc.PlainBytes = reverse(btc.PlainBytes)
		}
	}
	if btc.SecureBytes == nil {
		btc.SecureBytes = h2b(btc.Secure)
		if btc.Reverse {
			btc.SecureBytes = reverse(btc.SecureBytes)
		}
	}
	if btc.IVBytes == nil {
		btc.IVBytes = h2b(btc.IV)
		if btc.Reverse {
			btc.IVBytes = reverse(btc.IVBytes)
		}
	}
}

// Hash TestCase
type HashTestCase struct {
	Msg string // input, in hex
	MD  string // output, in hex

	MsgBytes []byte
	MDBytes  []byte
}

func (ht *HashTestCase) parse() {
	if ht.MsgBytes == nil {
		ht.MsgBytes = h2b(ht.Msg)
	}
	if ht.MDBytes == nil {
		ht.MDBytes = h2b(ht.MD)
	}
}
