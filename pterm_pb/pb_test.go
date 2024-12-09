package pterm_pb

import "testing"

func TestNewPB(t *testing.T) {
	pb := NewPB(9, "test")
	pb.Close()
}
