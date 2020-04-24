package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {
    fmt.Println(reverse("  ab cd ef   gg  h   "))
    fmt.Println(reverse("  Hello world!"))
    fmt.Println(reverse("a good   example"))
}

func reverse(str string) string {
    str = strings.Trim(str, " ")
    re := regexp.MustCompile("\\s+")
    str = re.ReplaceAllString(str, " ")
    arr := strings.Split(str, " ")
    ret := ""
    for i := len(arr) - 1; i >= 0; i -- {
        ret += arr[i] + " "
    }
    return ret[0:len(ret) - 1]
}

