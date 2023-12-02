# Week 10 Assignment: Modern Applied Statistics With Go

## Project Summary
This week we were asked to select a computer intensive statistical methods that was listed in Gelman and Vehtari's (2021) article,[“What Are the Most Important Statistical Ideas of the Past 50 Years?"](https://arxiv.org/abs/2012.00174). The section that stood out to my was the "overparameterized models." The idea of overfitting data is a major cause of concern when building machine learning models. Methods such as cross-fold validation and hyperparameter tuning are important to help reduce overfitting of training data. 

For this assignment we are going to run a simple random forest classifier and use cross-fold validation to avoiding overfitting in the model. Cross-validation provides a mechanism to assess a model's performance on multiple subsets of the data. It is a computer intensive method as it involves partitioning the dataset into subsets, training the model on some of these subsets, and evaluating its performance on the remaining subsets. It assess how well a predictive model will generalize to an independent dataset and helps detect potential issues such as overfitting.

We will be using the well-known iris dataset. The iris dataset is popular in the field of machine learning and statistics. In my experience, most textbooks for data science utilize the iris dataset for educational purposes. It is a built-in dataset in R and is also available in the [GoLearn](https://github.com/sjwhitworth/golearn/blob/master/examples/datasets/iris.csv) package. For consistency, I will download the copy from and use that for both codebases.

### R randomForest and caret
The randomForest package in R is used for building random forest models and the caret package is used to perform cross-validation. These packages are simple, well-documented and easy to use. In addition to performing cross-validation, the caret package automatically performs parameter tuning for "mtry" (number of features to consider at each split) as part of the cross-validation process.

### GoLearn
The [GoLearn](https://github.com/sjwhitworth/golearn) package implements a variety of machine learning models in Go, as well as a cross-validation evaluation. The cross-validation package automatically creates training/test views for each fold and trains the classifier on the training data while evaluating on the test data. I did not find this package to be necessarily well documented and I needed to read-through the actual code to understand what was going on. I do not find it as intuitive as the R packages used.

In order to get comparable results to the R output, I needed to manually adjust the features parameter. The R carat package used to build the random forest automatically performs parameter tuning for features. Initially I set it to the result that was returned from R (2), however this performed poorly for the Go code, so needed to manually adjust. Finally setting it to 4 returned a similar result.

## Files
### *main.go*
This file contains the RandomForestCrossValidation() function which takes in 4 parameters: input filename, number of trees for random forest model, number of features for model and number of folds for cross-validation. The main() function assigns values for the parameters and calls RandomForestCrossValidation.

### *main_test.go*
This file contains a test for the RandomForestCrossValidation() function. It uses the parameters from the R code, 100 trees, 2 features and 5 folds. The test will fail as the accuracy is not close to 93% which was the accuracy returned in the R code.

### *rf_classifierR.R*
This file contains the R code for building a random forest classifier using cross-validation.

## Results
### Benchmarking
Processing times were computed using the time command. The Go code ran in half the time as the R code, however, I note that the R code started by trying install/load additional packages when running from the terminal. This did not happen when running the code directly in VSCode.

> time ./msds431_mod10
```
0.84s user 
0.07s system 
85% cpu 
1.067 total
```
> time Rscript rf_classifierR.R
```
2.04s user 
0.19s system 
98% cpu 
2.252 total
```

## Conclusion
In general, I have been disappointed with the GoLearn pakage. I did not find the methods for performing cross-validation, in particular, very intuitive. It does not seem to follow the functionality of similar R and Python packages and it is not well documented. The few examples online that use the package don't really seem to follow the standard practice for building and testing ML models etiher. For example, they do not split the data into training and testing sets to actually test the model on an unseen subset of data. 

A reason that people would want to use Go for computer intensive statistical methods is its concurrency capabilities. However, GoLearn's cross-validation package does not even utilize concurrency. 

Overall, machine learning in Go is not mature enough to be used in a production setting. Unless the firm is ready to build their own models from scratch, I would recommend they stick to R or Python for machine learning tasks at this time. It may be beneficial for the firm to consider incorporating Go into their overall pipeline, for example using Go for the data engineering side of projects such as pre-processing of large datastreams and then use Python or R to build the actual model. There may be possibilities to deploy the resulting trained model back into Go. For example, I came across an interesting article, [TensorFlow in Go](https://medium.com/@exyzzy/tensorflow-in-go-9387fc130f50), which discusses using the Go language bindings to the TensorFlow C/C++ API. Essentially, this enables a user to build TensorFlow models in Python and then incoporate those models in their Go-backed production environment.

## References

Whitenack, Daniel. "Machine Learning With Go". Packt Publishing, 2017. 
[O'Reilley link](https://learning.oreilly.com/library/view/machine-learning-with/9781785882104/7e948a5b-542e-4d66-ae40-5b2964afc905.xhtml)