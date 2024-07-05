const faker = require('faker-br');
const fs = require('fs');
const path = require('path');

const numRecords = 10000;
const outputPath = 'testdata/large_users.csv';

const headers = 'Email,Name,Telephone,CPF,Address\n';

const generateUser = () => {
    const email = faker.internet.email();
    const name = faker.name.findName();
    const telephone = faker.phone.phoneNumber();
    const cpf = faker.br.cpf();
    const address = faker.address.streetAddress();
    return `${email},${name},${telephone},${cpf},${address}\n`;
};

const generateFairData = () => {
    const outputDir = path.dirname(outputPath);
    if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir);
    }

    fs.writeFileSync(outputPath, headers);

    for (let i = 0; i < numRecords; i++) {
        const user = generateUser();
        fs.appendFileSync(outputPath, user);
    }

    console.log(`Generated ${numRecords} records in ${outputPath}`);
}

module.exports = generateFairData;