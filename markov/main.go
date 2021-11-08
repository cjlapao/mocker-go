package markov

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	StartToken = "$"
	EndToken   = "^"
)

type chainJSON struct {
	Order    int                 `json:"int"`
	SpoolMap map[string]int      `json:"spool_map"`
	FreqMat  map[int]sparseArray `json:"freq_mat"`
}

//MarshalJSON ...
func (chain MarkovChain) MarshalJSON() ([]byte, error) {
	obj := chainJSON{
		chain.Order,
		chain.statePool.stringMap,
		chain.frequencyMat,
	}
	return json.Marshal(obj)
}

//UnmarshalJSON ...
func (chain *MarkovChain) UnmarshalJSON(b []byte) error {
	var obj chainJSON
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}
	chain.Order = obj.Order
	intMap := make(map[int]string)
	for k, v := range obj.SpoolMap {
		intMap[v] = k
	}
	chain.statePool = &spool{
		stringMap: obj.SpoolMap,
		intMap:    intMap,
	}
	chain.frequencyMat = obj.FreqMat
	chain.lock = new(sync.RWMutex)
	return nil
}

type Prefix []string

type Chain struct {
	chain     map[string][]string
	prefixLen int
}

type MarkovChain struct {
	Order        int
	statePool    *spool
	frequencyMat map[int]sparseArray
	lock         *sync.RWMutex
}

func (p Prefix) String() string {
	return strings.Join(p, " ")
}

func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Build(data []string) {
	p := make(Prefix, c.prefixLen)
	for _, word := range data {
		key := p.String()
		c.chain[key] = append(c.chain[key], word)
		p.Shift(word)
	}
}

func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}

	return strings.Join(words, " ")
}

// ######################

func NewMarkovChain(order int) *MarkovChain {
	chain := MarkovChain{Order: order}
	chain.statePool = &spool{
		stringMap: make(map[string]int),
		intMap:    make(map[int]string),
	}
	chain.frequencyMat = make(map[int]sparseArray, 0)
	chain.lock = new(sync.RWMutex)
	return &chain
}

func (chain *MarkovChain) Add(input []string) {
	startTokens := array(StartToken, chain.Order)
	endtokens := array(EndToken, chain.Order)
	tokens := make([]string, 0)
	tokens = append(tokens, startTokens...)
	tokens = append(tokens, input...)
	tokens = append(tokens, endtokens...)
	pairs := MakePairs(tokens, chain.Order)
	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]
		currentIndex := chain.statePool.add(pair.CurrentState.key())
		nextIndex := chain.statePool.add(pair.NextState)
		chain.lock.Lock()
		if chain.frequencyMat[currentIndex] == nil {
			chain.frequencyMat[currentIndex] = make(sparseArray, 0)
		}
		chain.frequencyMat[currentIndex][nextIndex]++
		chain.lock.Unlock()
	}
}

func (chain *MarkovChain) TransitionProbability(next string, current NGram) (float64, error) {
	if len(current) != chain.Order {
		return 0, errors.New("N-gram length does not match chain order")
	}
	currentIndex, currentExists := chain.statePool.get(current.key())
	nextIndex, nextExists := chain.statePool.get(next)
	if !currentExists || !nextExists {
		return 0, nil
	}
	arr := chain.frequencyMat[currentIndex]
	sum := float64(arr.sum())
	freq := float64(arr[nextIndex])
	return freq / sum, nil
}

func (chain *MarkovChain) Train(input []string) {
	for _, name := range input {

		removeSymbols := regexp.MustCompile("[?,:\\,,(,),#]")
		chain.Add(strings.Split(removeSymbols.ReplaceAllString(name, ""), " "))
	}
}

func (chain *MarkovChain) Generate(current NGram) (string, error) {
	if len(current) != chain.Order {
		return "", errors.New("N-gram length does not match chain order")
	}
	if current[len(current)-1] == EndToken {
		// Dont generate anything after the end token
		return "", nil
	}
	currentIndex, currentExists := chain.statePool.get(current.key())
	if !currentExists {
		return "", fmt.Errorf("Unknown ngram %v", current)
	}
	arr := chain.frequencyMat[currentIndex]
	sum := arr.sum()
	randN := rand.Intn(sum)
	for i, freq := range arr {
		randN -= freq
		if randN <= 0 {
			return chain.statePool.intMap[i], nil
		}
	}
	return "", nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (chain *MarkovChain) SaveModel() {
	jsonObj, _ := json.Marshal(chain)
	err := ioutil.WriteFile("model.json", jsonObj, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
