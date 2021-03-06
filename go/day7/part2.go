package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Node struct {
	name string
	weight int
	children map[string]Node
}

func buildNodes(input string) map[string]Node {
	nodes := map[string]Node{}
	for _,line := range strings.Split(input,"\n") {
		if line == "" { continue }
		node := Node{}
		node.children = map[string]Node{}
		parts := strings.Split(line," (")
		node.name = parts[0]
		next := strings.Split(parts[1],")")
		if val,err := strconv.Atoi(next[0]); err == nil {
			node.weight = val
		}
		nodes[node.name] = node
	}
	return nodes
}

func linkNodes(nodes map[string]Node, input string)string{
	hasParent := map[string]bool{}
	for _,line := range strings.Split(input,"\n") {
		if line == "" { continue }
		parts := strings.Split(line, " (")
		name := parts[0]
		node := nodes[name]
		children := strings.Split(parts[1], " -> ")
		if len(children) > 1 {
			for _,child := range strings.Split(children[1],", ") {
				hasParent[child]=true
				node.children[child] = nodes[child]
			}
		}
	}
	for name,_ := range nodes {
		if hasParent[name]==false {
			return name;
		}
	}
	panic("no root found")
}

func getWeight(nodes map[string]Node, node Node, depth int)int{
	weight := node.weight

	childWeights := map[int]int{}
	listWeights := []string{}
	for _,child := range node.children {
		childWeight := getWeight(nodes,child,depth+1)
		childWeights[childWeight]=childWeights[childWeight]+1
		listWeights=append(listWeights,child.name+":"+strconv.Itoa(childWeight))
		weight+=childWeight
	}

	if len(childWeights)>1 {

		//asumption that we can tell the odd weight because there will be more than one 'correct' weight
		targetWeight:=0
		oddWeight:=0
		for w,count := range childWeights {
			if count!=1 {
				targetWeight = w
			}else{
				oddWeight = w
			}
		}

		adjust:=targetWeight-oddWeight

		for _,child := range listWeights {
			parts := strings.Split(child,":")
			c := nodes[parts[0]]
			if w,err := strconv.Atoi(parts[1]); err == nil {
				if w==oddWeight {
					msg := "Node "+c.name+" has weight "+strconv.Itoa(c.weight)+" and should have weight "+strconv.Itoa(c.weight+adjust)
					panic(msg)
				}
			}
		}

	}

	return weight;
}

func main() {
	input:=testdata()
	nodes:=buildNodes(input)
	root:=linkNodes(nodes,input)
	getWeight(nodes,nodes[root],0)
	fmt.Println("Hello, playground, ",root,nodes)
}


func testdata() string {
data := `
pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)
`
return data
}
