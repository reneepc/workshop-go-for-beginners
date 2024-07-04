const faker = require('faker-br');
const fs = require('fs');

const numRecords = 10000;
const outputPath = 'data/large_users.csv';

const headers = 'Email,Name,Telephone,CPF,Address\n';

const generateUser = () => {
    const email = faker.internet.email();
    const name = faker.name.findName();
    const telephone = faker.phone.phoneNumber();
    const cpf = faker.br.cpf();
    const address = faker.address.streetAddress();
    return `${email},${name},${telephone},${cpf},${address}\n`;
};

fs.writeFileSync(outputPath, headers);

for (let i = 0; i < numRecords; i++) {
    const user = generateUser();
    fs.appendFileSync(outputPath, user);
}

console.log(`Generated ${numRecords} records in ${outputPath}`);
