package main

import (
	"tmorris/project/builder"
	"tmorris/project/calculator"
	"tmorris/project/uploador"
)

func main() {
	slcsps := builder.CreateSLCSPs("../resources/slcsp.csv")
	plans := builder.CreatePlans("../resources/plans.csv")
	zips := builder.CreateZips("../resources/zips.csv")

	updatedSlcsps := calculator.UpdateSlcspsWithRates(slcsps, plans, zips)

	uploador.CreateNewSLCSPFileWithRates(updatedSlcsps, "../resources/new_slcsp.csv")
}
