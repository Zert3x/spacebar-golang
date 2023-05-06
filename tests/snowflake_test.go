package tests

import (
	"fmt"
	"math/big"
	"spacebar-server/types"
	"testing"
)

func TestSnowflakeToBinary(t *testing.T) {
	id := types.Snowflake("1012248609196093511")
	binaryResult := "111000001100001110101011111100110111100001000011000001000111"

	bin, err := id.ToBinary()
	if err != nil {
		t.Fatalf("Encountered error: %s", err)
	}

	if bin != binaryResult {
		t.Fatal("Results did not match!")
	}
}

func TestBinaryToSnowflake(t *testing.T) {
	binary := "111000001100001110101011111100110111100001000011000001000111"
	idResult := types.Snowflake("1012248609196093511")

	snow, err := types.SnowflakeFromBinary(binary)
	if err != nil {
		t.Fatalf("Encountered error: %s", err)
	}

	if snow != idResult {
		t.Fatal("Results did not match!")
	}
}

func TestGenerateSnowflake(t *testing.T) {
	time := new(big.Int).Lsh(big.NewInt(1683333212865-types.EPOCH), 22)
	worker := big.NewInt((0 % 31) << 17)
	process := big.NewInt((1 % 31) << 12)
	orRes := new(big.Int)
	orRes.Or(time, worker)
	orRes.Or(orRes, process)
	orRes.Or(orRes, big.NewInt(0))
	target := "1104204269050925056"
	finalStr := orRes.String()
	if target != finalStr {
		t.Fatal("Snowflakes did not match!")
	}
}

func TestDeconstructSnowflake(t *testing.T) {
	snow := types.GenerateSnowflake(1, 0)
	decon, err := snow.Deconstruct()
	if err != nil {
		t.Fatalf("Failed to deconstruct: %s", err)
	}

	fmt.Println(decon)
}
