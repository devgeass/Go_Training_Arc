package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("len(\"devgeass\") =", len("devgeass"))
	fmt.Println("ToLower(\"DEVGEASS\") =", strings.ToLower("DEVGEASS"))
	fmt.Println("ToUpper(\"devgeass\") =", strings.ToUpper("devgeass"))
	fmt.Println("ToTitle(\"devGeass\") =", strings.ToTitle("devGeass"))
	fmt.Println("Contains(\"devgeass\",\"ss\") =", strings.Contains("devgeass", "ss"))
	fmt.Println("Count(\"devgeass\",\"e\") =", strings.Count("devgeass", "e"))
	fmt.Println("Compare(\"devg\",\"geass\") =", strings.Compare("dev", "geass"))
	fmt.Println("HasPrefix(\"devgeass\",\"d\") =", strings.HasPrefix("devgeass", "d"))
	fmt.Println("HasSuffix(\"devgeass\",\"d\") =", strings.HasSuffix("devgeass", "d"))
	fmt.Println("Fields(\"Follow devgeass on Instagram!\") =", strings.Fields("Follow devgeass on Instagram!"))
	fmt.Println("Split(\"Follow devgeass on Instagram!\",\" \") =", strings.Split("Follow devgeass on Instagram!", " "))
	fmt.Println("Split(\"Follow,devgeass,on,Instagram!\",\",\") =", strings.Split("Follow,devgeass,on,Instagram!", ","))
	fmt.Println("Split(\"devgeass\",\"\") =", strings.Split("devgeass", ""))
	fmt.Println("Index(\"devgeass\",\"e\") =", strings.Index("devgeass", "e"))
	fmt.Println("LastIndex(\"devgeass\",\"e\") =", strings.LastIndex("devgeass", "e"))
	fmt.Println("Join([]string{\"Follow\", \"devgeass!\"},\" \") =", strings.Join([]string{"Follow", "devgeass!"}, " "))
	fmt.Println("Repeat(\"abcb\",3) =", strings.Repeat("abcb", 3))
	fmt.Println("Replace(\"no pain no gain\", \"no\", \"yes\", 1) =", strings.Replace("no pain no gain", "no", "yes", 1))
	fmt.Println("ReplaceAll(\"no pain no gain\", \"no\", \"yes\") =", strings.ReplaceAll("no pain no gain", "no", "yes"))
	fmt.Println("Map(func IncrementCharByOne(char rune) rune,\"abcd\") =", strings.Map(func(char rune) rune {
		char += 1
		return char
	}, "abcd"))

	fmt.Println()
	fmt.Println("Format() =", time.Now().Format(time.ANSIC))
	fmt.Println("Day() =", time.Now().Day())
	fmt.Println("Month() =", time.Now().Month())
	fmt.Println("Year() =", time.Now().Year())
	fmt.Println("Date(2022, time.April, 1, 9, 25, 45, 45, 20, time.UTC) =", time.Date(2022, time.April, 1, 9, 25, 45, 20, time.UTC))
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Printing after 1s timeout!")
}
