package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// allergen -> list of possible ingredients
	data := make(map[string]map[string]struct{})

	var allIngredients []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		splt := strings.Split(line, " (contains ")
		ingredients := strings.Fields(splt[0])
		allIngredients = append(allIngredients, ingredients...)
		allergens := strings.Split(strings.TrimSuffix(splt[1], ")"), ", ")
		for _, a := range allergens {
			data[a] = intersection(data[a], ingredients)
		}
	}
	fmt.Println(data)
	var i int
	allergicIngredients := make(map[string]string, len(data))
	for len(data) > 0 {
		i++
		for allgn, ings := range data {
			if len(ings) == 1 {
				var ing string
				for ing = range ings {
				}
				allergicIngredients[ing] = allgn
				delete(data, allgn)
				for _, ings := range data {
					delete(ings, ing)
				}
			}
		}
		fmt.Println("Pass", i, "found", len(allergicIngredients))
	}
	fmt.Println(allergicIngredients)
	var count int
	for _, ing := range allIngredients {
		if _, ok := allergicIngredients[ing]; !ok {
			count++
		}
	}
	fmt.Println("Count of hypoallergenic ingredients:", count)

	type pair struct {
		ing      string
		allergen string
	}
	slc := make([]pair, 0, len(allergicIngredients))
	for ing, allgn := range allergicIngredients {
		slc = append(slc, pair{ing, allgn})
	}
	sort.Slice(slc, func(i, j int) bool {
		return slc[i].allergen < slc[j].allergen
	})
	var buf strings.Builder
	for i, p := range slc {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(p.ing)
	}
	fmt.Println(buf.String())
}

func intersection(a map[string]struct{}, b []string) map[string]struct{} {
	z := make(map[string]struct{})
	for _, bb := range b {
		if _, ok := a[bb]; a == nil || ok {
			z[bb] = struct{}{}
		}
	}
	return z
}

/*
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)

One must contain dairy: [ trh fvjkl sbzzf mxmxvkd ]
One must contain soy: [ sqjhc fvjkl ]
One must contain fish: [ sqjhc mxmxvkd sbzzf ]

One must contain dairy: [ mxmxvkd kfcds sqjhc nhms ]
One must contain fish: [ mxmxvkd kfcds sqjhc nhms ]

If sqjhc is soy

One must contain dairy: [ trh fvjkl sbzzf mxmxvkd ]
One must contain soy: [ sqjhc ] fvjkl
One must contain fish: [ mxmxvkd sbzzf ] sqjhc

One must contain dairy: [ mxmxvkd kfcds nhms ] sqjhc
One must contain fish: [ mxmxvkd kfcds nhms ] sqjhc

if mxmxvkd is fish

One must contain dairy: [ trh fvjkl sbzzf ] mxmxvkd ]
One must contain soy: [ sqjhc ] fvjkl
One must contain fish: [ mxmxvkd ] sbzzf ] sqjhc

One must contain dairy: [ kfcds nhms ] mxmxvkd ] sqjhc
One must contain fish: [ mxmxvkd ] kfcds nhms ] sqjhc

if kfcds is dairy

One must contain dairy: [ trh fvjkl sbzzf ] mxmxvkd ]
One must contain soy: [ sqjhc ] fvjkl
One must contain fish: [ mxmxvkd ] sbzzf ] sqjhc

One must contain dairy: [ kfcds ] nhms ] mxmxvkd ] sqjhc
One must contain fish: [ mxmxvkd ] kfcds nhms ] sqjhc

*/
/*
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)

One must contain dairy: [ trh fvjkl sbzzf mxmxvkd ]
One must contain dairy: [ mxmxvkd kfcds sqjhc nhms ]
One must contain soy: [ sqjhc fvjkl ]
One must contain fish: [ sqjhc mxmxvkd sbzzf ]
One must contain fish: [ mxmxvkd kfcds sqjhc nhms ]

Union:
Dairy: mxmxvkd
Fish: sqjhc
Soy: fvjkl
*/
