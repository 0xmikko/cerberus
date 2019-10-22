const abi = require("ethereumjs-abi");
const util = require("ethereumjs-util");
const BN = require("bn.js");

const assertActionThrows = action => (
	Promise
		.resolve()
		.then(action)
		.catch(error => {
			assert(error, "Expected an error to be raised");
			assert(error.message, "Expected an error to be raised");
			return error.message;
		})
		.then(errorMessage => {
			assert(errorMessage, "Expected an error to be raised");
			const invalidOpcode = errorMessage.includes("invalid opcode");
			const reverted = errorMessage.includes(
				"VM Exception while processing transaction: revert");
			assert(invalidOpcode || reverted,
				"expected following error message to include \"invalid JUMP\" or " +
			`"revert": "${errorMessage}"`);
		// see https://github.com/ethereumjs/testrpc/issues/39
		// for why the "invalid JUMP" is the throw related error when using TestRPC
		})
);

const toHexWithoutPrefix = arg => {
	if (arg instanceof Buffer || arg instanceof BN) {
		return arg.toString("hex");
	} else if (arg instanceof Uint8Array) {
		return Array.prototype.reduce.call(arg, (a, v) => a + v.toString("16").padStart(2, "0"), "");
	} else {
		return Buffer.from(arg, "ascii").toString("hex");
	}
};

const toHex = value => {
	return Ox(toHexWithoutPrefix(value));
};

const Ox = value => ("0x" !== value.slice(0, 2)) ? `0x${value}` : value;
const startMapBuffer = Buffer.from([0xBF]);
const endMapBuffer = Buffer.from([0xFF]);

const autoAddMapDelimiters = (data) => {
	let buffer = data;

	if (buffer[0] >> 5 !== 5) {
		buffer = Buffer.concat([startMapBuffer, buffer, endMapBuffer], buffer.length + 2);
	}

	return buffer;
};

const decodeRunRequest = log => {
	const runABI = util.toBuffer(log.data);
	const types = ["address", "bytes32", "uint256", "address", "bytes4", "uint256", "uint256", "bytes"];
	const [
		requester,
		requestId,
		payment,
		callbackAddress,
		callbackFunc,
		expiration,
		version,
		data
	] = abi.rawDecode(types, runABI);

	return {
		topic: log.topics[0],
		jobId: log.topics[1],
		requester: Ox(requester),
		id: toHex(requestId),
		payment: toHex(payment),
		callbackAddr: Ox(callbackAddress),
		callbackFunc: toHex(callbackFunc),
		expiration: toHex(expiration),
		dataVersion: version,
		data: autoAddMapDelimiters(data)
	};
};

const fulfillOracleRequest = async (oracle, request, response, options) => {
	if (!options) options = {};

	return oracle.fulfillOracleRequest(
		request.id,
		request.payment,
		request.callbackAddr,
		request.callbackFunc,
		request.expiration,
		response,
		options);
};

const increaseTime5Minutes = async () => {
	await web3.currentProvider.send({
		jsonrpc: "2.0",
		method: "evm_increaseTime",
		params: [300],
		id: 0
	}, (error, result) => {
		if (error) {
			// eslint-disable-next-line no-console
			console.log(`Error during helpers.increaseTime5Minutes! ${error}`);
			throw error;
		}
	});
};

const increaseBlocks = async (blocks) => {
	for (let i = 0; i < blocks; i++) {
		await web3.currentProvider.send({
			id: 0,
			jsonrpc: "2.0",
			method: "evm_mine",
		}, (error, result) => {
			if (error) {
				// eslint-disable-next-line no-console
				console.log(`Error during helpers.increaseBlocks! ${error}`);
				throw error;
			}
		});
	}
};

exports.assertActionThrows = assertActionThrows;
exports.decodeRunRequest = decodeRunRequest;
exports.fulfillOracleRequest = fulfillOracleRequest;
exports.increaseTime5Minutes = increaseTime5Minutes;
exports.increaseBlocks = increaseBlocks;