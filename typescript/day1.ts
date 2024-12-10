import * as fs from "fs";
import * as path from "path";

// Define the relative path to the input file
const inputFilePath = path.join(__dirname, "../inputs/dayOneInput.txt");

// Function to read and process the input file
const processDay1Input = (filePath: string): number => {
    try {
        // Read the file content
        const data = fs.readFileSync(filePath, "utf-8");

        // Split the file into rows (lines)
        const rows = data.trim().split("\n");

        // Initialize variables
        let grandTotal = 0;
        const list1: number[] = [];
        const list2: number[] = [];

        // Process each row
        for (const row of rows) {
            // Split the row into two halves
            const middle = Math.round(row.length / 2);
            const firstHalf = parseInt(row.slice(0, middle), 10);
            const secondHalf = parseInt(row.slice(middle), 10);

            // Add to the respective lists
            list1.push(secondHalf);
            list2.push(firstHalf);
        }

        // Sort the lists
        list1.sort((a, b) => a - b);
        list2.sort((a, b) => a - b);

        // Calculate the grand total
        for (let i = 0; i < list1.length; i++) {
            const sumTotal = list1[i] - list2[i];
            grandTotal += Math.abs(sumTotal);
        }

        return grandTotal;
    } catch (error) {
        console.error("Error processing input file:", error);
        return 0;
    }
};

// Main execution
(() => {
    console.log("Processing Day 1 input...");

    // Process the input file and calculate the result
    const result = processDay1Input(inputFilePath);

    // Output the result
    console.log("Answer for Day 1 Part 1 with your input file is " + result);
})();
