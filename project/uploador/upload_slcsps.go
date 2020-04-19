package uploador

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"tmorris/project/builder"
)

func CreateNewSLCSPFileWithRates(slcsps []builder.SLCSP, fileName string) {
	csvfile, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)
	_ = csvwriter.Write([]string{"zipcode", "rate"})

	for _, slcsp := range slcsps {
		_ = csvwriter.Write([]string{slcsp.Zipcode, fmt.Sprintf("%.2f", slcsp.Rate)})
	}

	csvwriter.Flush()

	_ = csvfile.Close()
}
