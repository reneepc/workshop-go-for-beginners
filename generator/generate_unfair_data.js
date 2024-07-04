const faker = require('faker-br');
const fs = require('fs');

const numRecords = 10000;
const repeatedEntries = 5000;
const outputPath = 'data/large_users_unfair.csv';

const headers = 'Email,Name,Telephone,CPF,Address\n';

// Function to generate a random user
const generateUser = () => {
    const email = faker.internet.email();
    const name = faker.name.findName();
    const telephone = faker.phone.phoneNumber();
    const cpf = faker.br.cpf();
    const address = faker.address.streetAddress();
    return `${email},${name},${telephone},${cpf},${address}\n`;
};

// Function to generate a specific user
const generateSpecificUser = () => {
    const email = 'rene.epcrdz@gmail.com';
    const name = 'Renê Cardozo';
    const telephone = '(11) 99999-9999';
    const cpf = '12345678909';
    const address = '302 Avenida Jurucê';
    return `${email},${name},${telephone},${cpf},${address}\n`;
};

fs.writeFileSync(outputPath, headers);

// Write repeated entries for the specific user
for (let i = 0; i < repeatedEntries; i++) {
    const user = generateSpecificUser();
    fs.appendFileSync(outputPath, user);
}

// Write random user entries
for (let i = 0; i < numRecords - repeatedEntries; i++) {
    const user = generateUser();
    fs.appendFileSync(outputPath, user);
}

console.log(`Generated ${numRecords} records in ${outputPath}`);
