package main

import (
	"fmt"
	"github.com/marguerite/linux-darkmode/sunrise"
	"time"
)

func main() {
	longitude := 123.4312
	latitude := 41.7806
	t := time.Now()
	rise := sunrise.SunRise(t, -1*longitude, latitude)
	fmt.Println(rise)
	fmt.Printf("%.7f\n", rise)
	set := sunrise.SunSet(t, -1*longitude, latitude)
	fmt.Printf("%.7f\n", set)

}
