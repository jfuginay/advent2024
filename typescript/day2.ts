import { readFileSync } from 'fs';

function parseReport(line: string): number[] {
    return line.split(' ').map(Number);
}

function isSafeReport(report: number[]): boolean {
    let increasing = true;
    let decreasing = true;

    for (let i = 0; i < report.length - 1; i++) {
        const diff = report[i + 1] - report[i];
        if (diff <= 0 || diff > 3) {
            increasing = false;
        }
        if (diff >= 0 || -diff > 3) {
            decreasing = false;
        }
        if (!increasing && !decreasing) {
            break;
        }
    }

    return increasing || decreasing;
}

function isSafeWithDampener(report: number[]): boolean {
    if (isSafeReport(report)) {
        return true;
    }

    for (let i = 0; i < report.length; i++) {
        const dampened = [...report.slice(0, i), ...report.slice(i + 1)];
        if (isSafeReport(dampened)) {
            return true;
        }
    }

    return false;
}

function solveBothParts(data: string[]): [number, number] {
    let part1Count = 0;
    let part2Count = 0;

    for (const line of data) {
        const report = parseReport(line);
        if (isSafeReport(report)) {
            part1Count++;
        }
        if (isSafeWithDampener(report)) {
            part2Count++;
        }
    }

    return [part1Count, part2Count];
}

function main() {
    try {
        // Read input file
        const inputPath = '/Users/iesouskurios/dev/advent2024/inputs/dayTwoInput.txt';
        const input = readFileSync(inputPath, 'utf-8')
            .trim()
            .split('\n');

        const [part1, part2] = solveBothParts(input);
        console.log(`Part 1: ${part1} safe reports`);
        console.log(`Part 2: ${part2} safe reports with dampener`);
    } catch (error) {
        console.error('Error:', error);
    }
}

main();
