library(randomForest)
library(caret)

## Load the iris dataset
#data("iris")
iris <- read.csv("datasets/iris_headers.csv")

## Set seed for reproducibility
set.seed(123)

## Split the data so that we use 70% of it for training
train_index <- createDataPartition(y=iris$Species, p=0.7, list=FALSE)

## Subset the data
train_data <- iris[train_index, ]
test_data <- iris[-train_index, ]

## Convert Species to categorical variable
train_data$Species <- as.factor(train_data$Species)
test_data$Species <- as.factor(test_data$Species)

# Create a Random Forest model using 5-fold cross-validation
ctrl <- trainControl(method = "cv", number = 5)
model <- train(Species ~ ., data=train_data, method = "rf", trControl = ctrl, metric="Accuracy", ntree = 100)

# Print the cross-validated results
print(model)

# Make predictions on the test set
predictions <- predict(model, newdata = test_data[, -5])

# Confusion matrix
cm <- confusionMatrix(predictions, test_data$Species)
print(cm)

accuracy <- mean(predictions == test_data$Species)*100
cat('Accuracy on testing data: ', round(accuracy, 2), '%',  sep='')
