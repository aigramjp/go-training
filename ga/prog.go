package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	npop := 50
	ndiv := 10
	ngen := 100
	nselect := 3

	pop := getInitialSamplePoint(npop, ndiv)
	evaluations(pop)
	bestPrint(pop, 0)

	for i := 0; i < ngen; i++ {
		pop = update(pop, nselect)
		evaluations(pop)
		bestPrint(pop, i+1)
	}
}

func bestPrint(pop []*Individual, imark int) {
	ind := best(pop)
	fmt.Printf("%d: %T %v\n", imark, ind, ind)
}

func update(pop []*Individual, nselect int) []*Individual {
	npop := len(pop)
	nu := 0.1
	mRate := 0.1
	mpRate := 0.1
	xpRate := 0.1

	pop2 := make([]*Individual, npop)
	for i := 0; i < npop; i++ {
		ind1 := selection(pop, nselect)
		ind2 := selection(pop, nselect)
		pop2[i] = crossover(ind1, ind2, nu, xpRate)
		if getRandom(0, 1) < mRate {
			mutation(pop2[i], mpRate)
		}
	}
	return pop2
}

func mutation(ind *Individual, rate float64) {
	for i := 0; i < len(ind.gene); i++ {
		if getRandom(0, 1) < rate {
			ind.gene[i] = getRandom(0, 1)
		}
	}
}

func evaluations(pop []*Individual) {
	for _, ind := range pop {
		evaluation(ind)
		//		fmt.Printf("%d: %T %v\n", i, ind, ind)
	}
}

func best(pop []*Individual) *Individual {
	npop := len(pop)
	iselected := 0

	for i := 1; i < npop; i++ {
		if pop[i].value > pop[iselected].value {
			iselected = i
		}
	}

	return pop[iselected]
}

func selection(pop []*Individual, nselect int) *Individual {
	npop := len(pop)

	iselected := -1
	for i := 0; i < nselect; i++ {
		idx := rand.Intn(npop)

		if i != 0 {
			if pop[idx].value > pop[iselected].value {
				iselected = idx
			}
		} else {
			iselected = idx
		}

	}

	return pop[iselected]
}

func evaluation(ind *Individual) {
	ind.value = 0
	for _, v := range ind.gene {
		ind.value += v
	}
}

type Individual struct {
	gene  []float64
	value float64
}

func crossover(ind1 *Individual, ind2 *Individual, nu float64, rate float64) *Individual {
	ndiv1 := len(ind1.gene)
	ndiv2 := len(ind2.gene)

	if ndiv1 != ndiv2 {
		fmt.Println("err")
	}

	li := make([]float64, ndiv1)
	for i := 0; i < ndiv1; i++ {
		if getRandom(0, 1) < rate {
			li[i] = blendX(ind1.gene[i], ind2.gene[i], nu)
		} else {
			li[i] = ind1.gene[i]
		}
	}

	return NewAssign(li)
}

func blendX(v1 float64, v2 float64, nu float64) float64 {
	minv := math.Min(v1, v2)
	maxv := math.Max(v1, v2)
	lenv := (maxv - minv) * nu

	minv = math.Max(minv-lenv, 0)
	maxv = math.Min(maxv+lenv, 1)

	return getRandom(minv, maxv)
}

func NewRandom(ndiv int) *Individual {
	return &Individual{getRandomList(ndiv), 0}
}

func NewAssign(gene []float64) *Individual {
	return &Individual{gene, 0}
}

func getInitialSamplePoint(npop int, ndiv int) []*Individual {
	pop := make([]*Individual, npop)
	for i := 0; i < npop; i++ {
		pop[i] = NewRandom(ndiv)
	}
	return pop
}

func getRandomList(n int) []float64 {
	li := make([]float64, n)
	for i := 0; i < n; i++ {
		li[i] = getRandom(0, 1)
	}
	return li
}

func getRandom(rmin float64, rmax float64) (r float64) {
	r = rmin + rand.Float64()*(rmax-rmin)
	return
}
