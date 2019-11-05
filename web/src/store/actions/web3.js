import getWeb3util from '../utils/getweb3';
import * as actionTypes from './actionTypes';
import CerberusWalletContract from '../../contracts/CerberusWallet.json';

export const getWeb3 = () => {
  return async (dispatch, getState) => {
    dispatch({type: actionTypes.WEB3_REQUEST});

    const web3 = await getWeb3util();
    const accounts = await web3.eth.getAccounts();
    const gasPrice = await web3.eth.getGasPrice();

    dispatch({
      type: actionTypes.WEB3_SUCCESS,
      payload: {web3, accounts, gasPrice},
    });
  };
};

export const deployContract = from => {
  return async (dispatch, getState) => {
    dispatch({type: actionTypes.CONTRACT_DEPLOY_REQUEST});
    const web3 = getState().web3.web3;

    const contract = new web3.eth.Contract(CerberusWalletContract.abi);
    const dContract = contract.deploy({
      data: CerberusWalletContract.bytecode,
      arguments: [
        '0x20fE562d797A42Dcb3399062AE9546cd06f63280', // Link Contract
        '0xc99B3D447826532722E41bc36e644ba3479E4365', // Alarm Oracle
        '0x2ebb1c1a4b1e4229adac24ee0b5f784f', // Alarm JobID
        '1000000000000000000', // Alarm Payment
        '0x83F00b902cbf06E316C95F51cbEeD9D2572a349a', // Cerberus Address
        '0x9a6e266667a64e83a6b3dcfc7e0fbd5f', // Cerberus JobID
        '1000000000000000000', // Cerberus Payment
      ],
    });

    const estimatedGas = await dContract.estimateGas();
    const gasPrice = await web3.eth.getGasPrice();

    dContract
      .send(
        {
          from,
          gas: Math.floor(estimatedGas * 1.5),
          gasPrice: gasPrice,
        },
        function(error, transactionHash) {
          console.log(transactionHash);
        },
      )
      .on('error', function(error) {
        console.log('Error', error);
      })
      .on('transactionHash', function(transactionHash) {
        console.log('THASH', transactionHash);
      })
      .on('receipt', function(receipt) {
        console.log(receipt.contractAddress); // contains the new contract address
        dispatch({
          type: actionTypes.CONTRACT_DEPLOY_SUCCESS,
          payload: receipt.contractAddress,
        });
      })
      .on('confirmation', function(confirmationNumber, receipt) {
        console.log(confirmationNumber);
      })
      .then(function(newContractInstance) {
        console.log(newContractInstance.options.address); // instance with the new contract address
      });
  };
};

export const depositMoney = (contractAddress, from, amount) => {
  console.log(contractAddress, amount);
  return async (dispatch, getState) => {
    dispatch({type: actionTypes.PAYMENT_REQUEST});
    console.log('Deposit money ', contractAddress, amount);
    const web3 = getState().web3.web3;

    const result = web3.eth.sendTransaction({
      from: from,
      to: contractAddress,
      value: amount,
    });
    dispatch({type: actionTypes.PAYMENT_SUCCESS});
    console.log(result);
  };
};

export const sendMoney = (contractAddress, from, to, amount) => {
  console.log(contractAddress, from, to, amount);
  return async (dispatch, getState) => {
    console.log('Send money ', contractAddress, to, amount);
    const web3 = getState().web3.web3;
    dispatch({type: actionTypes.PAYMENT_REQUEST});
    const contract = new web3.eth.Contract(
      CerberusWalletContract.abi,
      contractAddress,
    );
    const result = await contract.methods.sendMoney(to, amount).send({from});
    dispatch({type: actionTypes.PAYMENT_SUCCESS});
  };
};

export const depositLinks = (contract, amount) => {};
