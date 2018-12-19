package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

type PairList []Pair

type Pair struct {
	k string
	v []string
}

func NewMapSorter(m map[string][]string) PairList {
	ms := make(PairList, 0, len(m))
	for k, v := range m {

		ms = append(ms, Pair{k, v})

	}
	return ms
}

func (ms PairList) Len() int {
	return len(ms)
}

func (ms PairList) Less(i, j int) bool {
	return len(ms[i].v) > len(ms[j].v) //从大到小排序
}

func (ms PairList) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func minIndexOf(arr []int) int {
	minIndex := 0
	minValue := arr[0]
	for i := range arr {
		if arr[i] < minValue {
			minIndex = i
			minValue = arr[i]
		}
	}
	return minIndex
}

func countUrl(arr []string) int {
	var tmpUrl string
	urlCount := 0
	for _, line := range arr {
		lineSplit := strings.Split(line, "\t")
		if lineSplit[1] != tmpUrl {
			urlCount ++
		}
		tmpUrl = lineSplit[1]
	}
	return urlCount
}

func main() {
	args := os.Args
	var userInput string
	if args == nil || len(args) < 2 {
		fmt.Print("请将文件拖动到这个程序上，不要直接双击打开。按回车关闭此程序。")
		fmt.Scanln(&userInput)
		return
	}
	filePath := args[1]
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	lines := strings.Split(string(fd), "\n")
	var nameLineMap map[string][]string = map[string][]string{}
	for _, line := range lines {
		lineSplit := strings.Split(line, "\t")
		if len(lineSplit) != 6 {
			continue
		}
		tmpName := lineSplit[2]
		nameLineMap[tmpName] = append(nameLineMap[tmpName], line)
	}
	nameLinePairSorted := NewMapSorter(nameLineMap)
	sort.Stable(nameLinePairSorted)
	entityCount := 0
	for i := range nameLinePairSorted {
		entityCount += countUrl(nameLinePairSorted[i].v)
	}
	fmt.Printf("总共有%d个名称，%d个实体，请输入分割份数。", len(nameLinePairSorted), entityCount)
	splitNum := 2
	crossNum := 1
	fmt.Scanln(&userInput)
	splitNum, err = strconv.Atoi(userInput)
	fmt.Print("请输入交叉验证份数(默认是1，即不做交叉验证)，建议可以使用3，让每个人标注 3/人数 的数据，这样最后可以做交叉验证标注效果。")
	fmt.Scanln(&userInput)
	crossNum, err = strconv.Atoi(userInput)
	fmt.Printf("交叉%d倍，分成%d份。大约每份包含%d个实体。", crossNum, splitNum, entityCount/splitNum*crossNum)
	var groupStr = make([][]string, splitNum)
	var groupLineCount = make([]int, splitNum)
	for _, pair := range nameLinePairSorted {
		minIndex := minIndexOf(groupLineCount)
		for _, line := range pair.v {
			groupStr[minIndex] = append(groupStr[minIndex], line)
		}
		groupLineCount[minIndex] += countUrl(pair.v)
	}
	var fileObj *os.File
	for i := range groupStr {
		fileObj, err = os.OpenFile(strings.Join([]string{"./split_", ".txt"}, strconv.Itoa(i)), syscall.O_RDWR|syscall.O_CREAT|syscall.O_APPEND, 0666)
		for j := 0; j < crossNum; j++ {
			for _, line := range groupStr[(i+j)%splitNum] {
				fileObj.WriteString(line + "\n")
			}
		}
		fileObj.Close()
	}
}
