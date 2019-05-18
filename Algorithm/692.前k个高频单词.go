import "sort"

/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */
func topKFrequent(words []string, k int) []string {
	m := make(map[string]int, 0)
	for _, word := range words {
		val := m[word]
		m[word] = val + 1
	}
	entries := make([]Entry, len(m))
	i := 0
	for k, v := range m {
		entries[i] = Entry{
			Word:      k,
			Frequency: v,
		}
		i++
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Frequency < entries[j].Frequency {
			return false
		} else if entries[i].Frequency > entries[j].Frequency {
			return true
		} else {
			return entries[i].Word < entries[j].Word
		}
	})
	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = entries[i].Word
	}
	return result

}

type Entry struct {
	Word      string
	Frequency int
}

