package ga

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type FitnessFunc func(interface{}) float64
type DecodeFunc func(*BitArray) interface{}

type tChromoSome struct {
	bits    *BitArray
	fitness float64
}

type tChromoSomeArray []tChromoSome

type BasicBitsGaResult struct {
	Bits    *BitArray
	Fitness float64
	Decoded interface{}
}

type BasicBitsGa struct {
	length      int
	count       int
	population  tChromoSomeArray
	fitnessFunc FitnessFunc
	decodeFunc  DecodeFunc
}

func NewBasicBitsGa(fitnessFunc FitnessFunc, decodeFunc DecodeFunc, length int, count int) *BasicBitsGa {
	ga := new(BasicBitsGa)
	ga.fitnessFunc = fitnessFunc
	ga.decodeFunc = decodeFunc
	ga.length = length
	ga.count = count
	ga.generatePopulation()
	return ga
}

func (ga *BasicBitsGa) Evolve(retainRate, randomSelectRate, mutationRate float64) {
	parents := ga.selection(retainRate, randomSelectRate)
	ga.crossover(parents)
	ga.mutation(mutationRate)
}

func (ga *BasicBitsGa) EvolveN(num int, retainRate, randomSelectRate, mutationRate float64) {
	for i := 0; i < num; i++ {
		ga.Evolve(retainRate, randomSelectRate, mutationRate)
	}
}

func (ga *BasicBitsGa) generatePopulation() {
	count := ga.count
	length := ga.length

	ga.population = make([]tChromoSome, count)

	for i := 0; i < count; i++ {
		ba := NewBitArray(length)
		ba.FillRandBits()
		ga.population[i].bits = ba
		ga.population[i].fitness = ga.calcFitness(ba)
	}
}

func (ga *BasicBitsGa) calcFitness(bits *BitArray) float64 {
	decoded := ga.decodeFunc(bits)
	return ga.fitnessFunc(decoded)
}

func (ga *BasicBitsGa) selection(retainRate, randomSelectRate float64) int {
	count := ga.count

	ga.population.Sort()

	retainLen := int(math.Floor(retainRate * float64(count)))

	parentsNum := retainLen

	for i := retainLen; i < count; i++ {
		if rand.Float64() < randomSelectRate {
			ga.population[parentsNum] = ga.population[i]
			parentsNum++
		}
	}

	return parentsNum
}

func (ga *BasicBitsGa) crossover(parentsNum int) {
	parents := ga.population
	children := ga.population[parentsNum:]
	childrenNum := len(children)

	for i := 0; i < childrenNum; {
		male := rand.Intn(parentsNum)
		female := rand.Intn(parentsNum)
		if male != female {
			childBits := ga.bearChild(parents[male].bits, parents[female].bits)
			children[i].bits = childBits
			children[i].fitness = ga.calcFitness(childBits)
			i++
		}
	}
}

func (ga *BasicBitsGa) bearChild(male, female *BitArray) *BitArray {
	crossPos := rand.Intn(ga.length)
	return male.CrossAt(crossPos, female)
}

func (ga *BasicBitsGa) mutation(rate float64) {
	bitsLen := ga.length
	population := ga.population

	for i, n := 0, len(population); i < n; i++ {
		if rand.Float64() < rate {
			population[i].bits.Flip(rand.Intn(bitsLen))
			population[i].fitness = ga.calcFitness(population[i].bits)
		}
	}
}

func (ga *BasicBitsGa) Result() *BasicBitsGaResult {
	res := new(BasicBitsGaResult)
	res.Fitness = math.Inf(-1)

	for _, x := range ga.population {
		if x.fitness > res.Fitness {
			res.Fitness = x.fitness
			res.Bits = x.bits
		}
	}

	if res.Bits == nil {
		return nil
	}

	res.Decoded = ga.decodeFunc(res.Bits)
	return res
}

func (ga *BasicBitsGa) ResultAsString() string {
	r := ga.Result()
	if r == nil {
		return ""
	}

	return r.String()
}

// Len is the number of elements in the collection.
func (a tChromoSomeArray) Len() int {
	return len(a)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (a tChromoSomeArray) Less(i, j int) bool {
	return a[i].fitness > a[j].fitness
}

// Swap swaps the elements with indexes i and j.
func (a tChromoSomeArray) Swap(i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

// Sort the chromosome array
func (a tChromoSomeArray) Sort() {
	sort.Sort(a)
}

func (r *BasicBitsGaResult) String() string {
	return fmt.Sprintf("%f -- %s -- %v ", r.Fitness, r.Bits.String(), r.Decoded)
}
