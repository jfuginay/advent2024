"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var fs = require("fs");
var path = require("path");
// Define the relative path to the input file
var inputFilePath = path.join(__dirname, "../inputs/dayOneInput.txt");
// Function to read and process the input file
var processDay1Input = function (filePath) {
    try {
        // Read the file content
        var data = fs.readFileSync(filePath, "utf-8");
        // Split the file into rows (lines)
        var rows = data.trim().split("\n");
        // Initialize variables
        var grandTotal = 0;
        var list1 = [];
        var list2 = [];
        // Process each row
        for (var _i = 0, rows_1 = rows; _i < rows_1.length; _i++) {
            var row = rows_1[_i];
            // Split the row into two halves
            var middle = Math.round(row.length / 2);
            var firstHalf = parseInt(row.slice(0, middle), 10);
            var secondHalf = parseInt(row.slice(middle), 10);
            // Add to the respective lists
            list1.push(secondHalf);
            list2.push(firstHalf);
        }
        // Sort the lists
        list1.sort(function (a, b) { return a - b; });
        list2.sort(function (a, b) { return a - b; });
        // Calculate the grand total
        for (var i = 0; i < list1.length; i++) {
            var sumTotal = list1[i] - list2[i];
            grandTotal += Math.abs(sumTotal);
        }
        return grandTotal;
    }
    catch (error) {
        console.error("Error processing input file:", error);
        return 0;
    }
};
// Main execution
(function () {
    console.log("Processing Day 1 input...");
    // Process the input file and calculate the result
    var result = processDay1Input(inputFilePath);
    // Output the result
    console.log("Answer for Day 1 Part 1 with your input file is " + result);
})();
