pragma solidity ^0.4.24;

import "./SafeMath.sol";

// ----------------------------------------------------------------------------
// ERC Token Standard #20 Interface
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md
// ----------------------------------------------------------------------------
contract ERC20Interface {
	function totalSupply() public view returns (uint256);
	function balanceOf(address _owner) public view returns (uint256 balance);
	function allowance(address _owner, address _spender) public view returns (uint256 remaining);
	function transfer(address _to, uint256 _value) public returns (bool success);
	function transferFrom(address _from, address _to, uint _value) public returns (bool success);
	function approve(address _spender, uint256 _value) public returns (bool success);

	event Transfer(address indexed from, address indexed to, uint tokens);
	event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}


contract ReentrancyGuard {
	uint256 private guardCounter = 1;

	modifier noReentrant() {
		guardCounter += 1;
		uint256 localCounter = guardCounter;
		_;
		require(localCounter == guardCounter);
	}
}


interface tokenRecipient {
	function receiveApproval(address _from, uint256 _value, address _token, bytes _extraData) external;
}


contract ERC20Base is ERC20Interface , ReentrancyGuard{
	using SafeMath for uint256;

	string public name;
	string public symbol;
	uint8 public decimals = 18;
	uint256 public totalSupply;

	mapping(address => uint256) public balanceOf;
	mapping(address => mapping (address => uint256)) public allowance;

	constructor() public {
		//totalSupply = initialSupply * 10 ** uint256(decimals);
		uint256 initialSupply = 20000000000;
		totalSupply = initialSupply.mul(1 ether);
		balanceOf[msg.sender] = totalSupply;
		name = "ABCToken";
		symbol = "ABC";
	}

	function () payable public {
		revert();
	}

	function totalSupply() public view returns(uint256) {
		return totalSupply;
	}

	function balanceOf(address _owner) public view returns (uint256 balance) {
		return balanceOf[_owner];
	}

	function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
		return allowance[_owner][_spender];
	}

	function _transfer(address _from, address _to, uint256 _value) internal returns (bool success) {
		require(_to != 0x0);
		require(balanceOf[_from] >= _value);
		if (balanceOf[_to].add(_value) <= balanceOf[_to]) {
			revert();
		}

		uint256 previousBalances = balanceOf[_from].add(balanceOf[_to]);
		balanceOf[_from] = balanceOf[_from].sub(_value);
		balanceOf[_to] = balanceOf[_to].add(_value);
		emit Transfer(_from, _to, _value);
		assert(balanceOf[_from].add(balanceOf[_to]) == previousBalances);

		return true;
	}

	function transfer(address _to, uint256 _value) public returns (bool success) {
		return _transfer(msg.sender, _to, _value);
	}

	function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
		require(_value <= allowance[_from][msg.sender]);
		allowance[_from][msg.sender] = allowance[_from][msg.sender].sub(_value);
		return _transfer(_from, _to, _value);
	}

	function approve(address _spender, uint256 _value) public returns (bool success) {
		allowance[msg.sender][_spender] = _value;
		emit Approval(msg.sender, _spender, _value);
		return true;
	}

	function increaseApproval(address _spender, uint256 _addedValue) public returns (bool) {
		allowance[msg.sender][_spender] = (
		allowance[msg.sender][_spender].add(_addedValue));
		emit Approval(msg.sender, _spender, allowance[msg.sender][_spender]);
		return true;
	}

	function decreaseApproval(address _spender, uint256 _subtractedValue) public returns (bool) {
		uint256 oldValue = allowance[msg.sender][_spender];
		if (_subtractedValue >= oldValue) {
			allowance[msg.sender][_spender] = 0;
		} else {
			allowance[msg.sender][_spender] = oldValue.sub(_subtractedValue);
		}
		emit Approval(msg.sender, _spender, allowance[msg.sender][_spender]);
		return true;
	}

	function approveAndCall(address _spender, uint256 _value, bytes _extraData) noReentrant public returns (bool success) {
		tokenRecipient spender = tokenRecipient(_spender);
		if (approve(_spender, _value)) {
			spender.receiveApproval(msg.sender, _value, this, _extraData);
			return true;
		}
	}
}
