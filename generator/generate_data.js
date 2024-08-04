import faker from 'faker-br';
import fs from 'fs';
import path from 'path';

const headers = 'Email,Name,Phone,CPF,Address\n';

const generateUser = () => {
    const email = faker.internet.email();
    const name = faker.name.findName();
    const phone = faker.phone.phoneNumber();
    const cpf = faker.br.cpf();
    const address = faker.address.streetAddress();
    return `${email},${name},${phone},${cpf},${address}\n`;
};

const generateSpecificUser = () => {
    const email = 'rene.epcrdz@gmail.com';
    const name = 'Renê Cardozo';
    const phone = '(11) 99999-9999';
    const cpf = '12345678909';
    const address = '225 Rua Jaceru';
    return `${email},${name},${phone},${cpf},${address}\n`;
};

const generateUsers = (numRecords, outputPath, specificUserRatio) => {
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

export default generateUsers;
