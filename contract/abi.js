const fs = require('fs');

console.log(__dirname)
const js = fs.readFileSync(__dirname + '/build/contracts/CerberusWallet.json')
const data = JSON.parse(js)
console.log(data.abi)

fs.writeFileSync("abi.json", JSON.stringify(data.abi))
