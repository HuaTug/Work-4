package cache

import (
	"fmt"
	"strconv"
)

const Rankey = "ranking"

func TaskViewKey(id uint) string {
	fmt.Print("view:task:%s", strconv.Itoa(int(id)))
	return fmt.Sprintf("view:task:%s", strconv.Itoa(int(id)))
}
