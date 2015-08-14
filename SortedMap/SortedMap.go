package SortedMap

import "sort"

type SortedMap struct {
	Map  map[string]int
	Keys []string
}

func (sm *SortedMap) Len() int {
	return len(sm.Map)
}

func (sm *SortedMap) Less(i, j int) bool {
	return sm.Map[sm.Keys[i]] > sm.Map[sm.Keys[j]]
}

func (sm *SortedMap) Swap(i, j int) {
	sm.Keys[i], sm.Keys[j] = sm.Keys[j], sm.Keys[i]
}

func SortedKeys(m map[string]int) []string {
	sm := new(SortedMap)
	sm.Map = m
	sm.Keys = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.Keys[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.Keys
}
