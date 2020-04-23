package easy

import "strconv"

func groupAnagrams(strs []string) [][]string {
	allCombine := make(map[string][]string)

	for _,str:= range strs{
		key := getKey(str)
		allCombine[key] = append(allCombine[key],str)
	}

	i := 0
	result := make([][]string,len(allCombine))
	for v := range allCombine {
		result[i] = append(result[i],allCombine[v]...)
		i++
	}
	return result
}

func getKey(str string) (key string) {
	keyMap := make([]int,26)
	for _,s:=range str{
		keyMap[s-'a']++
	}
	for _,v := range keyMap {
		key += strconv.Itoa(v)
	}
	return
}
