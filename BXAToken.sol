pragma solidity ^0.4.24;

import "./Migrations.sol";
import "./ERC20Base.sol";

contract BXAToken is Migrations, ERC20Base {
	bool public isTokenLocked;
	bool public isUseFreeze;
	struct Frozen {
		bool from;
		uint256 amount;
	}
	mapping (address => Frozen) public frozenAccount;

	event FrozenFunds(address target, bool freezeFrom, uint256 freezeAmount);

	constructor()
		ERC20Base()
		onlyOwner()
		public
	{
		uint256 initialSupply = 20000000000;
		isUseFreeze = true;
		totalSupply = initialSupply.mul(1 ether);
		isTokenLocked = false;
		symbol = "BXA";
		name = "BXA";
		balanceOf[msg.sender] = totalSupply;
		emit Transfer(address(0), msg.sender, totalSupply);
	}

	modifier tokenLock() {
		require(isTokenLocked == false);
		_;
	}

	function setLockToken(bool _lock) onlyOwner public {
		isTokenLocked = _lock;
	}

	function setUseFreeze(bool _useOrNot) onlyAdmin public {
		isUseFreeze = _useOrNot;
	}

	function freezeFrom(address target, bool fromFreeze) onlyAdmin public {
		frozenAccount[target].from = fromFreeze;
		emit FrozenFunds(target, fromFreeze, 0);
	}

	function freezeAmount(address target, uint256 amountFreeze) onlyAdmin public {
		frozenAccount[target].amount = amountFreeze;
		emit FrozenFunds(target, false, amountFreeze);
	}

	function freezeAccount(
		address target,
		bool fromFreeze,
		uint256 amountFreeze
	) onlyAdmin public {
		require(isUseFreeze);
		frozenAccount[target].from = fromFreeze;
		frozenAccount[target].amount = amountFreeze;
		emit FrozenFunds(target, fromFreeze, amountFreeze);
	}

	function isFrozen(address target) public view returns(bool, uint256) {
		return (frozenAccount[target].from, frozenAccount[target].amount);
	}

	function _transfer(address _from, address _to, uint256 _value) tokenLock internal returns(bool success) {
		require(balanceOf[_from] >= _value);

		if (balanceOf[_to].add(_value) <= balanceOf[_to]) {
			revert();
		}

		if (isUseFreeze == true) {
			require(frozenAccount[_from].from == false);

			if(balanceOf[_from].sub(_value) < frozenAccount[_from].amount) {
				revert();
			}
		}

		if (_to == address(0)) {
			require(msg.sender == owner);
			totalSupply = totalSupply.sub(_value);
		}
		balanceOf[_from] = balanceOf[_from].sub(_value);
		balanceOf[_to] = balanceOf[_to].add(_value);
		emit Transfer(_from, _to, _value);

		return true;
	}

	function totalBurn() public view returns(uint256) {
		return balanceOf[address(0)];
	}
}
