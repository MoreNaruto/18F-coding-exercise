package calculator

import (
	"testing"
	"tmorris/project/builder"
)

func TestCalculatingSlcsps(t *testing.T) {
	expectedRates := []float64{348.73, 348.73, 266.62, 0}

	plans := buildPlans()
	zips := buildZips()
	slcsp := buildSLCSPs()

	updatedSlcsps := UpdateSlcspsWithRates(slcsp, plans, zips)

	for i := 0; i < len(updatedSlcsps); i++ {
		if expectedRates[i] != updatedSlcsps[i].Rate {
			t.Errorf("expected %v, got %v", expectedRates[i], updatedSlcsps[i].Rate)
		}
	}

}

func buildSLCSPs() []builder.SLCSP {
	return []builder.SLCSP{
		{
			Zipcode: "36749",
			Rate:    0,
		},
		{
			Zipcode: "36703",
			Rate:    0,
		},
		{
			Zipcode: "36343",
			Rate:    0,
		},
		{
			Zipcode: "35585",
			Rate:    0,
		},
	}
}

func buildZips() []builder.Zip {
	return []builder.Zip{
		{
			Zipcode:     "36749",
			State:       "AL",
			Countrycode: "01001",
			Name:        "Autauga",
			Ratearea:    11,
		},
		{
			Zipcode:     "36703",
			State:       "AL",
			Countrycode: "01001",
			Name:        "Autauga",
			Ratearea:    11,
		},
		{
			Zipcode:     "36343",
			State:       "AL",
			Countrycode: "01069",
			Name:        "Houston",
			Ratearea:    6,
		},
		{
			Zipcode:     "35585",
			State:       "AL",
			Countrycode: "01059",
			Name:        "Franklin",
			Ratearea:    13,
		},
		{
			Zipcode:     "36477",
			State:       "AL",
			Countrycode: "01061",
			Name:        "Geneva",
			Ratearea:    6,
		},
	}
}

func buildPlans() []builder.Plan {
	return []builder.Plan{
		{
			Planid:     "74449NR9870320z",
			State:      "GA",
			Metalfield: "Silver",
			Rate:       298.62,
			Ratearea:   7,
		},
		{
			Planid:     "26325VH2723968",
			State:      "FL",
			Metalfield: "Silver",
			Rate:       421.43,
			Ratearea:   60,
		},
		{
			Planid:     "09846WB8636633",
			State:      "IL",
			Metalfield: "Gold",
			Rate:       361.69,
			Ratearea:   5,
		},
		{
			Planid:     "78590JQ3204809",
			State:      "SC",
			Metalfield: "Gold",
			Rate:       379.97,
			Ratearea:   2,
		},
		{
			Planid:     "21163YD2193896",
			State:      "AR",
			Metalfield: "Gold",
			Rate:       359.75,
			Ratearea:   2,
		},
		{
			Planid:     "36956VI7724021",
			State:      "TX",
			Metalfield: "Silver",
			Rate:       310.12,
			Ratearea:   11,
		},
		{
			Planid:     "39287DQ4265604",
			State:      "FL",
			Metalfield: "Platinum",
			Rate:       420.79,
			Ratearea:   18,
		},
		{
			Planid:     "12654WJ5898785",
			State:      "WI",
			Metalfield: "Bronze",
			Rate:       332.97,
			Ratearea:   11,
		},
		{
			Planid:     "97573AL3146346",
			State:      "SC",
			Metalfield: "Gold",
			Rate:       338.17,
			Ratearea:   23,
		},
		{
			Planid:     "97573AL3146346",
			State:      "SC",
			Metalfield: "Gold",
			Rate:       338.17,
			Ratearea:   23,
		},
		{
			Planid:     "48279KR8882106",
			State:      "WV",
			Metalfield: "Gold",
			Rate:       356.81,
			Ratearea:   9,
		},
		{
			Planid:     "72516MB1294942",
			State:      "PA",
			Metalfield: "Gold",
			Rate:       290.53,
			Ratearea:   3,
		},
		{
			Planid:     "47733MN6034814",
			State:      "FL",
			Metalfield: "Gold",
			Rate:       426.97,
			Ratearea:   17,
		},
		{
			Planid:     "72190DW3284384",
			State:      "AZ",
			Metalfield: "Silver",
			Rate:       353.37,
			Ratearea:   11,
		},
		{
			Planid:     "39103VX4715005",
			State:      "WI",
			Metalfield: "Bronze",
			Rate:       294.39,
			Ratearea:   14,
		},
		{
			Planid:     "80895CM6393587",
			State:      "NV",
			Metalfield: "Platinum",
			Rate:       325.2669285,
			Ratearea:   1,
		},
		{
			Planid:     "18378UN5835046",
			State:      "TX",
			Metalfield: "Gold",
			Rate:       249.89,
			Ratearea:   15,
		},
		{
			Planid:     "04175ZU5314839",
			State:      "IL",
			Metalfield: "Gold",
			Rate:       350.99,
			Ratearea:   11,
		},
		{
			Planid:     "22521KG6880321",
			State:      "GA",
			Metalfield: "Silver",
			Rate:       285.76,
			Ratearea:   13,
		},
		{
			Planid:     "52890CP9247848",
			State:      "PA",
			Metalfield: "Silver",
			Rate:       264.46,
			Ratearea:   5,
		},
		{
			Planid:     "71371IL9623602",
			State:      "WI",
			Metalfield: "Silver",
			Rate:       348.73,
			Ratearea:   11,
		},
		{
			Planid:     "82206LV3597031",
			State:      "MI",
			Metalfield: "Silver",
			Rate:       321.99,
			Ratearea:   6,
		},
		{
			Planid:     "25775GT9423594",
			State:      "MT",
			Metalfield: "Silver",
			Rate:       262.93,
			Ratearea:   6,
		},
		{
			Planid:     "65128VB5535475",
			State:      "SC",
			Metalfield: "Silver",
			Rate:       266.62,
			Ratearea:   6,
		},
		{
			Planid:     "94970PT4045197",
			State:      "SC",
			Metalfield: "Gold",
			Rate:       377.73,
			Ratearea:   6,
		},
		{
			Planid:     "66347PV5658737",
			State:      "MI",
			Metalfield: "Silver",
			Rate:       361.0751527,
			Ratearea:   6,
		},
	}
}
