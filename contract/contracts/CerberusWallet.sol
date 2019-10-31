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

  event NewPaymentRegistered(bytes32 request, address to, uint256 amount);
  event MoneyCome(uint256 value);
  event MoneySent(address to, uint256 amount);

  uint256 public data;
  uint256 public oraclePayment;

  address public alarmOracle;
  bytes32 public alarmJobId;
  uint256 public alarmPayment;

  address public cerberusOracle;
  bytes32 public cerberusJobId;
  uint256 public cerberusPayment;

  uint8 delay;

  struct PaymentRequest {
    address recipient;
    uint256 amount;
  }

  mapping(bytes32 => PaymentRequest) public orders;
  mapping(bytes32 => bytes32) public ordersMapping;

  /**
   * @notice Deploy the contract with a specified address for the LINK
   * and Oracle contract addresses
   * @dev Sets the storage for the specified addresses
   * @param _link The address of the LINK token contract
   */
  constructor(address _link,
              address _alarmOracle,
              bytes32 _alarmJobId,
              uint256 _alarmPayment,
              address _cerberusOracle,
              bytes32 _cerberusJobId,
              uint256 _cerberusPayment
            ) public {

    alarmOracle = _alarmOracle;
    alarmJobId = _alarmJobId;
    alarmPayment = _alarmPayment;
    cerberusOracle = _cerberusOracle;
    cerberusJobId = _cerberusJobId;
    cerberusPayment = _cerberusPayment;

    if (_link == address(0)) {
      setPublicChainlinkToken();
    } else {
      setChainlinkToken(_link);
    }
  }

  function () payable onlyOwner {
    emit MoneyCome(msg.value);
  }

  function sendMoney(address _to, uint256 _amount) {
    Chainlink.Request memory req = buildChainlinkRequest(alarmJobId, this, this.fulfillConfirmationRequest.selector);
    req.addUint("until", now + 1 minutes);
    bytes32 reqID = sendChainlinkRequestTo(alarmOracle, req, alarmPayment);
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
    Chainlink.Request memory req = buildChainlinkRequest(cerberusJobId, this, this.fulfillPaymentRequest.selector);
    req.add("id", bytes32ToStr(_requestId));
    bytes32 reqID = sendChainlinkRequestTo(cerberusOracle, req, cerberusPayment);
    ordersMapping[reqID] = _requestId;
  }

  function bytes32ToStr(bytes32 _bytes32) public pure returns (string) {

    bytes memory bytesArray = new bytes(64);
    for (uint256 i; i < 32; i++) {
      bytesArray[2*i] = bytes1(uint8(_bytes32[i] & 240 ) / 16 + 65);
      bytesArray[2*i+1] = bytes1(uint8(_bytes32[i] & 15) + 65);

    }
    return string(bytesArray);
  }

  function fulfillPaymentRequest(bytes32 _requestId, bool _data)
  public
  recordChainlinkFulfillment(_requestId)
  {
    if(_data) {
      PaymentRequest request = orders[ordersMapping[_requestId]];
      address to = request.recipient;
      uint256 amount = request.amount;

      emit MoneySent(to, amount);
    }
  }

  function getAlarmOracle() public view returns (address)  {
    return alarmOracle;
  }

  function getAlarmPayment() public view returns (uint256)  {
    return alarmPayment;
  }
  function getAlarmJobId() public view returns (bytes32)  {
    return alarmJobId;
  }



}
