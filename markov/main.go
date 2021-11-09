package markov

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mb-14/gomarkov"
)

type MarkovGenerator struct {
}

func NewGenerator() *MarkovGenerator {
	return &MarkovGenerator{}
}

func (s *MarkovGenerator) buildModel(model TrainingDatabase, order int) *gomarkov.Chain {
	chain := gomarkov.NewChain(order)
	for _, data := range s.getDataset(model) {
		chain.Add(split(data))
	}
	return chain
}

func split(str string) []string {
	return strings.Split(str, "")
}

func (s *MarkovGenerator) getDataset(model TrainingDatabase) []string {
	file, _ := os.Open("./models/" + model.String())
	scanner := bufio.NewScanner(file)
	var list []string
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}

func (s *MarkovGenerator) loadModel(model MarkovModel) (*gomarkov.Chain, error) {
	var chain gomarkov.Chain
	data, err := ioutil.ReadFile("./models/" + model.String())
	if err != nil {
		return &chain, err
	}
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return &chain, err
	}
	return &chain, nil
}

func (s *MarkovGenerator) saveModel(model MarkovModel, chain *gomarkov.Chain) {
	jsonObj, _ := json.Marshal(chain)
	err := ioutil.WriteFile("./models/"+model.String(), jsonObj, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *MarkovGenerator) Train(database TrainingDatabase, chainsize int) {
	chain := s.buildModel(database, chainsize)
	s.saveModel(database.ToMarkovModel(), chain)
}

func (s *MarkovGenerator) Generate(model MarkovModel) string {
	chain, err := s.loadModel(model)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	order := chain.Order
	tokens := make([]string, 0)
	for i := 0; i < order; i++ {
		tokens = append(tokens, gomarkov.StartToken)
	}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - order):])
		tokens = append(tokens, next)
	}
	name := strings.Join(tokens[order:len(tokens)-1], "")

	return name
}
