package types

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

const EPOCH = 1420070400000

var INCREMENT = new(big.Int)

type Snowflake string

type DeconstructedSnowflake struct {
	Timestamp int64
	WorkerID  int64
	ProcessID int64
	Increment int64
	Binary    string
}

func SnowflakeFromBinary(num string) (Snowflake, error) {
	dec := ""

	for len(num) > 50 {
		highSlice := num[:len(num)-32]
		high, _ := strconv.ParseInt(highSlice, 2, 64)
		low, _ := strconv.ParseInt(strconv.FormatInt(high%10, 2)+num[len(num)-32:], 2, 64)

		dec = strconv.FormatInt(low%10, 10) + dec
		second := fmt.Sprintf("%032b", low/10)
		num = strconv.FormatInt(high/10, 2) + second
	}

	numInt, _ := strconv.ParseInt(num, 2, 64)
	for numInt > 0 {
		dec = strconv.FormatInt(numInt%10, 10) + dec
		numInt = numInt / 10
	}

	return Snowflake(dec), nil
}

func (s Snowflake) ToBinary() (string, error) {
	bin := ""
	high, err := strconv.Atoi(string(s[:len(s)-10]))
	if err != nil {
		return "", err
	}
	low, err := strconv.Atoi(string(s[len(s)-10:]))
	if err != nil {
		return "", err
	}

	for low > 0 || high > 0 {
		bin = strconv.Itoa(low&1) + bin
		low = int(math.Floor(float64(low / 2)))
		if high > 0 {
			low += 5000000000 * (high % 2)
			high = int(math.Floor(float64(high / 2)))
		}
	}
	return bin, nil
}

func (s Snowflake) Deconstruct() (*DeconstructedSnowflake, error) {
	binary, err := s.ToBinary()
	if err != nil {
		return nil, err
	}

	binary = fmt.Sprintf("%064s", binary)

	ts, err := strconv.ParseInt(binary[:42], 2, 64)
	if err != nil {
		return nil, err
	}
	wId, err := strconv.ParseInt(binary[42:47], 2, 64)
	if err != nil {
		return nil, err
	}
	pId, err := strconv.ParseInt(binary[47:52], 2, 64)
	if err != nil {
		return nil, err
	}
	inc, err := strconv.ParseInt(binary[52:64], 2, 64)
	if err != nil {
		return nil, err
	}

	return &DeconstructedSnowflake{
		Timestamp: ts + EPOCH,
		WorkerID:  wId,
		ProcessID: pId,
		Increment: inc,
		Binary:    binary,
	}, nil
}

func GenerateSnowflake(processId int64, workerId int64) Snowflake {
	timeBig := new(big.Int).Lsh(big.NewInt((time.Now().UnixNano()/int64(time.Millisecond))-EPOCH), 22)
	worker := big.NewInt((workerId % 31) << 17)
	process := big.NewInt((processId % 31) << 12)
	orRes := new(big.Int)
	orRes.Or(timeBig, worker)
	orRes.Or(orRes, process)
	orRes.Or(orRes, INCREMENT)
	INCREMENT.Add(INCREMENT, big.NewInt(1))
	return Snowflake(orRes.String())
}

func (dsf *DeconstructedSnowflake) Date() time.Time {
	return time.UnixMilli(dsf.Timestamp)
}
