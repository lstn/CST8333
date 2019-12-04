// CST8333 Cheese Directory App - Unit Tests - Lucas Estienne
package main

import (
	"testing"
	"reflect"
)

// test to verify that our "loadData" function loads the first record from the dataset properly
func TestLoadData(t *testing.T) {

	// create a record with the proper data
	firstRecord := Record {
		CheeseId: 226,
		CheeseName: "Sieur de Duplessis (Le)",
		ManufacturerName: "Fromages la faim de loup",
		ManufacturerProvCode: "NB",
		ManufacturingType: "Farmstead",
		WebSite: "N/A",
		FatContentPercent: 24.2,
		MoisturePercent: 47,
		Particularities: "N/A",
		Flavour: "Sharp, lactic",
		Characteristics: "Uncooked",
		Ripening: "9 Months",
		Organic: false,
		CategoryType: "Firm Cheese",
		MilkType: "Ewe",
		MilkTreatmentType: "Raw Milk",
		RindType: "Washed Rind",
		LastUpdateDate: "2016-02-03",
	}

	// load data and get the first record
	records := loadData("data/canadianCheeseDirectory.csv", 5)
	loadedFirstRecord := records[0]

	// check if loaded first record and our test record are equal
    if !reflect.DeepEqual(firstRecord, loadedFirstRecord) {
       t.Errorf("Loaded First Record was incorrect, \n got: \n%+v\n, want: \n%+v\n", firstRecord, loadedFirstRecord)
    }
}