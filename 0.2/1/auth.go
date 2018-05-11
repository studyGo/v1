/*

伪代码
original_secret = xxxx xxxx xxxx xxxx xxxx xxxx xxxx xxxx  
secret = BASE32_DECODE(TO_UPPERCASE(REMOVE_SPACES(original_secret)))  
input = CURRENT_UNIX_TIME() / 30  
hmac = SHA1(secret + SHA1(secret + input))  
four_bytes = hmac[LAST_BYTE(hmac):LAST_BYTE(hmac) + 4]  
large_integer = INT(four_bytes)  
small_integer = large_integer % 1,000,000 

author http://www.csdn.net/article/2014-09-23/2821808-Google-Authenticator

*/

package main

import (
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base32"
    "fmt"
    "os"
    "strings"
    "time"
)

func toBytes(value int64) []byte {
    var result []byte
    mask := int64(0xFF)
    shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
    for _, shift := range shifts {
        fmt.Println(byte((value>>shift)&mask))
        result = append(result, byte((value>>shift)&mask))
    }
    fmt.Println(result)
    return result
}

func toUint32(bytes []byte) uint32 {
    return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
        (uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func oneTimePassword(key []byte, value []byte) uint32 {
    hmacSha1 := hmac.New(sha1.New, key)
    hmacSha1.Write(value)
    hash := hmacSha1.Sum(nil)
    offset := hash[len(hash)-1] & 0x0F

    hashParts := hash[offset : offset+4]

    hashParts[0] = hashParts[0] & 0x7F

    number := toUint32(hashParts)

    pwd := number % 1000000

    return pwd
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "must specify key to use")
        os.Exit(1)
    }
    input := os.Args[1]

    inputNoSpaces := strings.Replace(input, " ", "", -1)
    inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
    key, err := base32.StdEncoding.DecodeString(inputNoSpacesUpper)
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }

    epochSeconds := time.Now().Unix()
    pwd := oneTimePassword(key, toBytes(epochSeconds/30))

    secondsRemaining := 30 - (epochSeconds % 30)
    fmt.Printf("%06d (%d second(s) remaining)\n", pwd, secondsRemaining)
}
