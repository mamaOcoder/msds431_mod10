package main

import (
	"fmt"
	"testing"
)

func TestRandomForestCrossValidation(t *testing.T) {
	// Specify the path to a test dataset
	testDatasetPath := "datasets/iris_headers.csv"

	// Set hyperparameters and cross-validation folds for testing using results from R
	numTrees := 100
	numFeatures := 2
	numFolds := 5

	// Run the random forest with cross-validation
	// Run the random forest with cross-validation
	accuracy := RandomForestCrossValidation(testDatasetPath, numTrees, numFeatures, numFolds)

	fmt.Println(accuracy)

	// Expected accuracy (adjust as needed)
	expectedAccuracy := 0.93

	// Check if the obtained accuracy is close to the expected value
	if !isCloseEnough(accuracy, expectedAccuracy, 0.02) {
		t.Fatalf("Test failed. Obtained accuracy (%.2f) is not close to expected accuracy (%.2f)", accuracy, expectedAccuracy)
	}
}

func isCloseEnough(a, b, tolerance float64) bool {
	return a >= b-tolerance && a <= b+tolerance
}
