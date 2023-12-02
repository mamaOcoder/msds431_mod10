package main

import (
	"fmt"
	"log"
	"math"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
)

func RandomForestCrossValidation(dataPath string, numTrees, numFeatures, numFolds int) float64 {
	// Load the dataset
	iris, err := base.ParseCSVToInstances(dataPath, true)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a Random Forest classifier
	rf := ensemble.NewRandomForest(numTrees, numFeatures)

	// Perform cross-validation
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(iris, rf, numFolds)
	if err != nil {
		log.Fatal(err)
	}

	// Print out confusion matrices
	for i, cm := range cv {
		fmt.Printf("Confusion Matrix - Fold %d:\n", i+1)
		fmt.Println(evaluation.ShowConfusionMatrix(cm))
	}

	// Get the mean, variance, and standard deviation of the accuracy for cross-validation
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// Output the cross-validation metrics to standard out
	fmt.Printf("\nAccuracy\n%.2f (+/- %.2f)\n\n", mean, stdev*2)

	return mean
}

func main() {
	// Specify the path to the dataset
	dataPath := "datasets/iris_headers.csv"

	// Set hyperparameters and cross-validation folds
	// The GoLearn package doesn't perform any hyperparameter tuning for features, so set manually
	// For better performance (and to better simulate R results), consider implementing feature tuning
	numTrees := 100
	numFeatures := 4
	numFolds := 5

	// Run the random forest with cross-validation
	RandomForestCrossValidation(dataPath, numTrees, numFeatures, numFolds)
}
