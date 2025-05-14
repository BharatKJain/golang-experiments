package main

import (
	"encoding/json"
	"fmt"
)

type response2 struct {
	Page   int
	Fruits []string
}

// Advanced Json structs
type Medication struct {
	Name     string `json:"name"`
	Dose     string `json:"dose"`
	Strength string `json:"strength"`
}

type ClassName struct {
	AssociatedDrug  []Medication `json:"associatedDrug"`
	AssociatedDrug2 []Medication `json:"associatedDrug#2"`
}

type MedicationsClass struct {
	ClassName []ClassName `json:"className"`
}

type MedicationProblem struct {
	Medications []struct {
		MedicationsClasses []MedicationsClass `json:"medicationsClasses"`
	} `json:"medications"`
}

type Problem struct {
	Diabetes []MedicationProblem `json:"Diabetes"`
}

type Problems struct {
	Problems []Problem `json:"problems"`
}

func main() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	jsonData := `{
		"problems": [{
			"Diabetes":[{
				"medications":[{
					"medicationsClasses":[{
						"className":[{
							"associatedDrug":[{
								"name":"asprin",
								"dose":"",
								"strength":"500 mg"
							}],
							"associatedDrug#2":[{
								"name":"somethingElse",
								"dose":"",
								"strength":"500 mg"
							}]
						}]
					}]
				}]
			}]
		}]}`

	var problems Problems
	err := json.Unmarshal([]byte(jsonData), &problems)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Name: %s", problems.Problems[0].Diabetes[0].Medications[0].MedicationsClasses[0].ClassName[0].AssociatedDrug[0].Name)

}
