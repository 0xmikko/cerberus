pragma solidity 0.4.24;

import "chainlink/contracts/ChainlinkClient.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

/**
 * @title CerberusWallet is an example contract which requests data from
 * the Chainlink network
 * @dev This contract is designed to work on multiple networks, including
 * local test networks
 */
contract CerberusWallet is ChainlinkClient, Ownable {
  event NewPaymentRegistered(bytes32 request, address to, int256 amount);
  event GetPermission(bytes32 id);
  event Code(int256 code);
  event MoneyCome(int256 value);

  uint256 public data;
  uint256 oraclePayment;

  address alarmOracle;
  bytes32 alarmJobId;

  address cerberusOracle;
  bytes32 cerberusJobId;

  struct PaymentRequest {
    address recipient;
    int256 amount;
  }

  mapping(bytes32=>PaymentRequest) public orders;

  /**
   * @notice Deploy the contract with a specified address for the LINK
   * and Oracle contract addresses
   * @dev Sets the storage for the specified addresses
   * @param _link The address of the LINK token contract
   */
  constructor(address _link) public {
    if (_link == address(0)) {
      setPublicChainlinkToken();
    } else {
      setChainlinkToken(_link);
    }
  }

  function () payable onlyOwner {
    emit MoneyCome(msg.value);
  }

  function makeConfirmationRequest(uint256 _payment, address _to, int256 _amount) {
    Chainlink.Request memory req = buildChainlinkRequest(alarmJobId, this, this.fulfillConfirmationRequest.selector);
    req.addUint("until", now + 1 minutes);
    bytes32 reqID = sendChainlinkRequestTo(alarmOracle, req, _payment);
    orders[reqID] = PaymentRequest(_to, _amount);
    emit NewPaymentRegistered(reqID, _to, _amount);
  }

  /**
   * @notice The fulfill method from requests created by this contract
   * @dev The recordChainlinkFulfillment protects this function from being called
   * by anyone other than the oracle address that the request was sent to
   * @param _requestId The ID that was generated for the request
   * @param _data The answer provided by the oracle
   */
  function fulfillConfirmationRequest(bytes32 _requestId, uint256 _data)
  public
  recordChainlinkFulfillment(_requestId)
  {
    emit GetPermission(_requestId);
  }

  function makePaymentRequest(bytes32 _requestID) {
    Chainlink.Request memory req = buildChainlinkRequest(cerberusJobId, this, this.fulfillTimeRequest.selector);
    req.addUint("until", now + 1 minutes);
    bytes32 reqID = sendChainlinkRequestTo(cerberusOracle, req, _payment);
    orders[reqID] = PaymentRequest(_to, _amount);
    emit NewPaymentRegistered(reqID, _to, _amount);
  }









}
