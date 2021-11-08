package mocker

import (
	"fmt"
	"strings"

	"github.com/cjlapao/mocker-go/markov"
)

var (
	jobTitleSectionOne = []string{
		"Lead", "Senior", "Direct", "Corporate", "Dynamic", "Future", "Product", "National", "Regional", "District",
		"Central", "Global", "Relational", "Customer", "Investor", "Dynamic", "International", "Legacy", "Forward",
		"Interactive", "Internal", "Human", "Chief", "Principal", "Quick",
	}
	jobTitleSectionTwo = []string{
		"Solutions", "Program", "Brand", "Security", "Research", "Marketing", "Directives", "Implementation", "Integration",
		"Functionality", "Response", "Paradigm", "Tactics", "Identity", "Markets", "Group", "Resonance", "Applications",
		"Optimization", "Operations", "Infrastructure", "Intranet", "Communications", "Web", "Branding", "Quality",
		"Assurance", "Impact", "Mobility", "Ideation", "Data", "Creative", "Configuration", "Accountability", "Interactions",
		"Factors", "Usability", "Metrics", "Team",
	}
	jobTitleSectionThree = []string{
		"Supervisor", "Associate", "Executive", "Liaison", "Officer", "Manager", "Engineer", "Specialist", "Director", "Coordinator",
		"Administrator", "Architect", "Analyst", "Designer", "Planner", "Synergist", "Orchestrator", "Technician", "Developer",
		"Producer", "Consultant", "Assistant", "Facilitator", "Agent", "Representative", "Strategist",
	}
	companyPrefix = []string{
		"Solutions", "Capital", "Capital Partners", "Partners", "Bank", "Holdings", "and Son", "and Sons", "& Son", "& Sons",
		"and Co", "& Co",
	}
	companyDenomination = []string{
		"LLC", "LTD", "SA", "GmbH", "Inc.",
	}
)

type Company struct {
	Mocker *Mocker
}

func testMarkov() {
	//Create a chain of order 2
	chain := markov.NewMarkovChain(2)

	//Feed in training data
	chain.Add(strings.Split("I want a cheese burger", " "))
	chain.Add(strings.Split("I want a chilled sprite", " "))
	chain.Add(strings.Split("I want to go to the movies", " "))

	//Get transition probability of a sequence
	// prob, _ := chain.TransitionProbability("a", []string{"I", "want"})
	// fmt.Println(prob)
	//Output: 0.6666666666666666

	//You can even generate new text based on an initial seed
	chain.Add(strings.Split("Mother should I build the wall?", " "))
	chain.Add(strings.Split("Mother should I run for President?", " "))
	chain.Add(strings.Split("Mother should I trust the government?", " "))
	next, _ := chain.Generate([]string{"should", "I"})
	fmt.Println(next)

	//The chain is JSON serializable
	// jsonObj, _ := json.Marshal(chain)
	// err := ioutil.WriteFile("model.json", jsonObj, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func (c Company) Name() string {
	// datasource := data.GetCompanyNameDatasource()
	// sp.Logger.Info(datasource.Names[46])
	// chain := markov.NewChain(2)
	// chain.Build(datasource.Names)
	// generatedName := chain.Generate(3)
	// sp.Logger.Info(generatedName)
	// mchain := markov.NewMarkovChain(3)
	// mchain.Train(datasource.Names)
	// order := mchain.Order
	// tokens := make([]string, 0)
	// for i := 0; i < order; i++ {
	// 	tokens = append(tokens, markov.StartToken)
	// }
	// for tokens[len(tokens)-1] != markov.EndToken {
	// 	next, _ := mchain.Generate(tokens[(len(tokens) - order):])
	// 	tokens = append(tokens, next)
	// }
	// markov.GetPokeme()
	// // mchain.SaveModel()
	// aName := strings.Join(tokens[order:len(tokens)-1], " ")
	// // fmt.Println(strings.Join(tokens[order:len(tokens)-1], " "))
	// hasDenomination := c.Mocker.Boolean().ChanceOfBool(20)
	// hasPrefix := c.Mocker.Boolean().ChanceOfBool(38)
	// name := fmt.Sprintf("%v", aName)
	// if hasPrefix {
	// 	name = fmt.Sprintf("%v %v", name, c.Mocker.RandomStringElement(companyPrefix))
	// }
	// if hasDenomination {
	// 	name = fmt.Sprintf("%v, %v", name, c.Mocker.RandomStringElement(companyDenomination))
	// }

	name := testerPokemon(false)

	return name
}

func (c Company) JobTitle() string {
	titleSize := c.Mocker.IntBetween(1, 3)

	switch titleSize {
	case 1:
		return c.Mocker.RandomStringElement(jobTitleSectionThree)
	case 2:
		return fmt.Sprintf("%v %v", c.Mocker.RandomStringElement(jobTitleSectionTwo), c.Mocker.RandomStringElement(jobTitleSectionThree))
	case 3:
		return fmt.Sprintf("%v %v %v", c.Mocker.RandomStringElement(jobTitleSectionOne), c.Mocker.RandomStringElement(jobTitleSectionTwo), c.Mocker.RandomStringElement(jobTitleSectionThree))
	default:
		return fmt.Sprintf("%v %v %v", c.Mocker.RandomStringElement(jobTitleSectionOne), c.Mocker.RandomStringElement(jobTitleSectionTwo), c.Mocker.RandomStringElement(jobTitleSectionThree))
	}
}
