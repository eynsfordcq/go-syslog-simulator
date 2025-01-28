package utils

import (
	"fmt"
	"math/rand"
	"strconv"
)

func RandomElement[T any](list []T) T {
	return list[rand.Intn(len(list))]
}

func RandomHost(hostname, domainName string) string {
	randomHost := strconv.Itoa(rand.Intn(10) + 1)
	return fmt.Sprintf("%s%s%s", hostname, randomHost, domainName)
}

func RandomPid() int {
	return rand.Intn(9500) + 500
}
