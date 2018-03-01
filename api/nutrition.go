package main

type nutrition struct {
	Calories    float64 `json:"calories"`
	Carbs       float64 `json:"carbs"`
	Protein     float64 `json:"protein"`
	Fat         float64 `json:"fat"`
	Cholestorol float64 `json:"cholestorol"`
}

// Appends the nutrition stats
func (n *nutrition) append(new nutrition) {
	n.Calories += new.Calories
	n.Carbs += new.Carbs
	n.Protein += new.Protein
	n.Fat += new.Fat
	n.Cholestorol += new.Cholestorol
}

// Truncates the precision of the nutrition fields to 2 digits
func (n *nutrition) truncate() {
	n.Calories = truncatePrecision(n.Calories)
	n.Carbs = truncatePrecision(n.Carbs)
	n.Protein = truncatePrecision(n.Protein)
	n.Fat = truncatePrecision(n.Fat)
	n.Cholestorol = truncatePrecision(n.Cholestorol)
}

func truncatePrecision(f float64) float64 {
	tmp := int(f * 100)
	last := int(f*1000) - tmp*10
	if last >= 5 {
		tmp++
	}
	return float64(tmp) / 100
}
