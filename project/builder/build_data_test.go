package builder

import (
	"testing"
)

func TestBuildFirstPlanProperly(t *testing.T) {
	plans := CreatePlans("../../resources/plans.csv")

	expectedPlan := Plan{
		Planid:     "74449NR9870320",
		State:      "GA",
		Metalfield: "Silver",
		Rate:       298.62,
		Ratearea:   7,
	}

	if plans[0] !=  expectedPlan {
		t.Errorf("expected %v, got %v", expectedPlan, plans[0])
	}
}

func TestBuildFirstZipProperly(t *testing.T) {
	zips := CreateZips("../../resources/zips.csv")

	expectedZip := Zip{
		Zipcode:     "36749",
		State:       "AL",
		Countrycode: "01001",
		Name:        "Autauga",
		Ratearea:    11,
	}

	if zips[0] !=  expectedZip {
		t.Errorf("expected %v, got %v", expectedZip, zips[0])
	}
}

func TestBuildFirstSLCSPProperly(t *testing.T) {
	slcsps := CreateSLCSPs("../../resources/slcsp.csv")

	expectedSLCSP:= SLCSP{
		Zipcode: "64148",
		Rate:    0,
	}

	if slcsps[0] !=  expectedSLCSP {
		t.Errorf("expected %v, got %v", expectedSLCSP, slcsps[0])
	}
}
