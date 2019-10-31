let MyContract = artifacts.require('CerberusWallet')
let LinkToken = artifacts.require('LinkToken')
let Oracle = artifacts.require('Oracle')

module.exports = (deployer, network) => {
  // Local (development) networks need their own deployment of the LINK
  // token and the Oracle contract
  if (!network.startsWith('live')) {
    deployer.deploy(LinkToken).then(() => {
      return deployer.deploy(Oracle, LinkToken.address).then(() => {
        return deployer.deploy(MyContract,
            "0x20fE562d797A42Dcb3399062AE9546cd06f63280",                       // Link Contract
            "0xc99B3D447826532722E41bc36e644ba3479E4365",                       // Alarm Oracle
            "0x2ebb1c1a4b1e4229adac24ee0b5f784f",     // Alarm JobID
            "1000000000000000000",                  // Alarm Payment
            "0x83F00b902cbf06E316C95F51cbEeD9D2572a349a",                       // Cerberus Address
            "0xd4b02e3a2c354111911739c5dd3264a9",     // Cerberus JobID
            "1000000000000000000") //LinkToken.address)
      })
    })
  } else {
    // For live networks, use the 0 address to allow the ChainlinkRegistry
    // contract automatically retrieve the correct address for you
    deployer.deploy(MyContract, '0x0000000000000000000000000000000000000000')
  }
}
