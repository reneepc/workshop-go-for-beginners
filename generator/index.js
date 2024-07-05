const generateFairData = require('./generate_fair_data');
const generateUnfairData = require('./generate_unfair_data');

function runGenerators() {
    try {
        generateFairData();
        generateUnfairData();
    } catch (error) {
        console.error(`Error running generators: ${error}`);
    }
}

runGenerators();
