package main

import (
	"fmt"
	"strings"
	"sort"
)

func splitWords(line string, ch chan string){
	for _,word := range strings.Split(line, " ") {
		ch<-word
	}
	close(ch)
}

func normaliseWord(word string)string{
	s:=strings.Split(word,"")
	sort.Strings(s)
	return strings.Join(s,"")
}

func doDay4(input string){
	validPhrases := 0

	for _, line := range strings.Split(input, "\n") {
		ch := make(chan string)
		if line == "" { continue }
		go splitWords(line,ch)
		seen := make(map[string]int)
		for word := range ch {
			sortWord := normaliseWord(word)
			seen[sortWord] = seen[sortWord]+1
		}
		phraseValid := true
		for _,count := range seen {
			if count!=1 {
				phraseValid=false
				break
			}
		}
		if phraseValid {
			validPhrases++
		}
	}
	fmt.Println(validPhrases)	
}


func main() {
	input := `
bdwdjjo avricm cjbmj ran lmfsom ivsof
mxonybc fndyzzi gmdp gdfyoi inrvhr kpuueel wdpga vkq
bneh ylltsc vhryov lsd hmruxy ebnh pdln vdprrky
fumay zbccai qymavw zwoove hqpd rcxyvy
bcuo khhkkro mpt dxrebym qwum zqp lhmbma esmr qiyomu
`	

doDay4(input)
}
