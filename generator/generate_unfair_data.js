const faker = require('faker-br');
const fs = require('fs');
const path = require('path');

const numRecords = 10000;
const specificUserRatio = 0.5; // Approximate ratio for the specific user
const outputPath = 'testdata/large_users_unfair.csv';

const headers = 'Email,Name,Telephone,CPF,Address\n';

const generateUser = () => {
    const email = faker.internet.email();
    const name = faker.name.findName();
    const telephone = faker.phone.phoneNumber();
    const cpf = faker.br.cpf();
    const address = faker.address.streetAddress();
    return `${email},${name},${telephone},${cpf},${address}\n`;
};

const generateSpecificUser = () => {
    const email = 'rene.epcrdz@gmail.com';
    const name = 'Renê Cardozo';
    const telephone = '(11) 99999-9999';
    const cpf = '12345678909';
    const address = '302 Avenida Jurucê';
    return `${email},${name},${telephone},${cpf},${address}\n`;
};

const generateUnfairData = () => {
    const outputDir = path.dirname(outputPath);
    if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir);
    }

    fs.writeFileSync(outputPath, headers);

    let specificUserCount = 0;
    let randomUserCount = 0;

    for (let i = 0; i < numRecords; i++) {
        const useSpecificUser = Math.random() < specificUserRatio && specificUserCount < numRecords * specificUserRatio;
        const user = useSpecificUser ? generateSpecificUser() : generateUser();

        fs.appendFileSync(outputPath, user);

        if (useSpecificUser) {
            specificUserCount++;
        } else {
            randomUserCount++;
        }
    }

    console.log(`Generated ${numRecords} records in ${outputPath}`);
    console.log(`Specific User Records: ${specificUserCount}`);
    console.log(`Random User Records: ${randomUserCount}`);
};

module.exports = generateUnfairData;
