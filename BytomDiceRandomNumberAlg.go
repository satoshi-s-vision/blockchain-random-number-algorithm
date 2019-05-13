package main

import (
  "crypto/sha256"
  "encoding/binary"
  "encoding/hex"
  "fmt"
)

const (
  n = 624
)

func Roll(serverSeed, txID [32]byte, index uint32) uint32 {
  seed := getNumberSeed(serverSeed, txID, index)
  return getRandomNumber(seed)
}

func getNumberSeed(serverSeed, txID [32]byte, index uint32) [32]byte {
  var indexHex [4]byte
  binary.BigEndian.PutUint32(indexHex[:], index)

  data := append(serverSeed[:], txID[:]...)
  data = append(data, indexHex[:]...)

  return sha256.Sum256(data)
}

func getRandomNumber(seed [32]byte) uint32 {
  var seeds [8]uint32
  for i := 0; i < 8; i++ {
    seeds[i] = binary.BigEndian.Uint32(seed[i*4 : (i+1)*4])
  }
  return getMersenneTwister(seeds)
}

func getMersenneTwister(seeds [8]uint32) uint32 {
  x := [n]uint32{19650218}

  for i := uint32(1); i < n; i++ {
    x[i] = 1812433253*(x[i-1]^(x[i-1]>>30)) + i
  }

  i, j := uint32(1), uint32(0)
  for k := n; k > 0; k-- {
    x[i] = (x[i] ^ ((x[i-1] ^ (x[i-1] >> 30)) * 1664525) + seeds[j] + j)
    if i++; i >= n {
      x[0] = x[n-1]
      i = 1
    }
    if j++; j >= 8 {
      j = 0
    }
  }

  for j := uint32(0); j <= n-1; j++ {
    x[i] = x[i] ^ ((x[i-1] ^ (x[i-1] >> 30)) * 1566083941) - i
    if i++; i >= n {
      x[0] = x[n-1]
      i = 1
    }
  }

  y := 0x80000000 | (x[1] & 0x7fffffff)

  x[0] = x[397] ^ (y >> 1) ^ ((y & 1) * 0x9908b0df)

  z := x[0] ^ (x[0] >> 11)
  z ^= (z << 7) & 0x9d2c5680
  z ^= (z << 15) & 0xefc60000
  z ^= (z >> 18)
  return z >> 16
}

func main() {
  var serverSeed, txID [32]byte
  tempData, err := hex.DecodeString("c9656775fd0f9d40dcde80fd597cc7e863144cb4fccbfc1e47de1ff4849ab054")
  if err != nil {
    fmt.Println(err)
  }
  copy(serverSeed[:32], tempData)

  tempData, err = hex.DecodeString("b31e198874365efc77338439e16c9b92c9b0c75a49c82c61c4e4423b71139996")
  if err != nil {
    fmt.Println(err)
  }
  copy(txID[:32], tempData)

  fmt.Println("\nRoll res: ", Roll(serverSeed, txID, 0))
}
