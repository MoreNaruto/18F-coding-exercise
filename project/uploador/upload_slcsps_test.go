package uploador

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"testing"
	"tmorris/project/builder"
)

func TestCrateNewSLCSPFileWithRates(t *testing.T) {
	expectedRates := []float64{234.23, 103.23, 21.30, 0.00}
	slcsps := buildSLCSPs()
	CreateNewSLCSPFileWithRates(slcsps, "test.csv")

	reader := ReadCSV("test.csv")

	for i := 0; i < len(expectedRates); i++{
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		rate, _ := strconv.ParseFloat(line[1], 64)

		if expectedRates[i] != rate {
			t.Errorf("expected %v, got %v", expectedRates[i], rate)
		}
	}

	t.Cleanup(func() {
		var err = os.Remove("test.csv")
		if err != nil {
			fmt.Println(err.Error())
		}
	})
}

func buildSLCSPs() []builder.SLCSP {
	return []builder.SLCSP{
		{
			Zipcode: "36749",
			Rate:    234.23,
		},
		{
			Zipcode: "36703",
			Rate:    103.234,
		},
		{
			Zipcode: "36343",
			Rate:    21.3,
		},
		{
			Zipcode: "35585",
			Rate:    0,
		},
	}
}

func ReadCSV(fileName string) *csv.Reader {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	_, _ = reader.Read()
	return reader
}