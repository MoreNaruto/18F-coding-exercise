package calculator

import (
	"sort"
	"tmorris/project/builder"
)

func UpdateSlcspsWithRates(initialSlcsps []builder.SLCSP, plans []builder.Plan, zips []builder.Zip) []builder.SLCSP {
	silverPlans := filterPlans(plans, isSilverMetalLevel)
	ratesByArea := getRatesBasedOnArea(silverPlans)
	slcspsZips := retrieveZipsOnlyInSlcsps(initialSlcsps, zips)

	return calculateSlcsps(initialSlcsps, ratesByArea, slcspsZips)
}

func calculateSlcsps(slcsps []builder.SLCSP, ratesByArea map[int64][]float64, zips map[string]builder.Zip) (updatedSlcsps []builder.SLCSP) {
	for _, slcsp := range slcsps {
		zip := zips[slcsp.Zipcode]
		rates := ratesByArea[zip.Ratearea]

		if len(rates) > 1 {
			sort.Float64s(rates)
			slcsp.Rate = rates[1]
		}
		updatedSlcsps = append(updatedSlcsps, slcsp)
	}
	return
}

func retrieveZipsOnlyInSlcsps(slcsps []builder.SLCSP, zips []builder.Zip) map[string]builder.Zip {
	slcspsZips := make(map[string]builder.Zip)

	for _, slcsp := range slcsps {
		slcspsZips[slcsp.Zipcode] = builder.Zip{}
	}

	for _, zip := range zips {
		if contains(slcspsZips, zip.Zipcode) {
			slcspsZips[zip.Zipcode] = zip
		}
	}
	return slcspsZips
}

func contains(zips map[string]builder.Zip, string string) bool {
	for key, _ := range zips {
		if key == string {
			return true
		}
	}
	return false
}

func filterPlans(plans []builder.Plan, filter func(plan builder.Plan) bool) (filteredPlans []builder.Plan) {
	for _, plan := range plans {
		if filter(plan) {
			filteredPlans = append(filteredPlans, plan)
		}
	}
	return
}

func isSilverMetalLevel(plan builder.Plan) bool {
	return plan.Metalfield == "Silver"
}

func getRatesBasedOnArea(plans []builder.Plan) map[int64][]float64 {
	rateAreaRates := make(map[int64][]float64)
	for _, plan := range plans {
		currentRateArea := plan.Ratearea
		if _, ok := rateAreaRates[currentRateArea]; ok {
			rateAreaRates[currentRateArea] = append(rateAreaRates[currentRateArea], plan.Rate)
		} else {
			rateAreaRates[currentRateArea] = []float64{plan.Rate}
		}
	}

	return rateAreaRates
}
