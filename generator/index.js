import generateUsers from './generate_data.js';

function runGenerators() {
    try {
        generateUsers(10, 'testdata/users_001.csv', 0);
        generateUsers(10000, 'testdata/users_002.csv', 0);
        generateUsers(10000, 'testdata/users_003.csv', 0.4);
    } catch (error) {
        console.error(`Error running generators: ${error}`);
    }
}

runGenerators();
