package group

type Score struct {
	Id int
	Name string
	Class string
	Score int
}

var scores []Score

func Group(source []Score,
		keySelector func(s Score)interface{}) map[interface{}][]Score{
	returnMap := make(map[interface{}][]Score)
	for _, s := range source{
		key:=keySelector(s)
		returnMap[key]=append(returnMap[key],s)
	}
	return returnMap
}

func SumScores(gs map[interface{}][]Score) map[interface{}]int{
	return Aggregate(gs, getSum)
}

func AvgScores(gs map[interface{}][]Score) map[interface{}]int {
	return Aggregate(gs, getAvg)
}
func MaxScores(gs map[interface{}][]Score) map[interface{}]int {
	getMax := func(s []Score) int {
		max := 0
		for _, score := range s {
			if max < score.Score {
				max = score.Score
			}
		}
		return max
	}
	return Aggregate(gs, getMax)
}

func Aggregate(gs map[interface{}][]Score, aggregator func(s []Score)int) map[interface{}]int {
	returnMap := make(map[interface{}]int)
	for k, sSlice := range gs {
		avg := aggregator(sSlice)
		returnMap[k] = avg
	}
	return returnMap
}


func getSum(sSlice []Score) int {
	sum := 0
	for _, s := range sSlice {
		sum += s.Score
	}
	return sum
}

func getAvg(sSlice []Score) int {
	sum := 0
	for _, s := range sSlice {
		sum += s.Score
	}
	avg := sum / len(sSlice)
	return avg
}

func GroupByClass(source []Score) map[interface{}][]Score{
	return Group(source,
		func(s Score)interface{}{
		return s.Class
	})
}