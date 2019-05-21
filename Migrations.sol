pragma solidity ^0.4.24;

contract Migrations {
	address public owner;
	address public newOwner;

	address public manager;
	address public newManager;

	event TransferOwnership(address oldaddr, address newaddr);
	event TransferManager(address oldaddr, address newaddr);

	modifier onlyOwner() { require(msg.sender == owner); _; }
	modifier onlyManager() { require(msg.sender == manager); _; }
	modifier onlyAdmin() { require(msg.sender == owner || msg.sender == manager); _; }


	constructor() public {
		owner = msg.sender;
		manager = msg.sender;
	}

	function transferOwnership(address _newOwner) onlyOwner public {
		newOwner = _newOwner;
	}

	function transferManager(address _newManager) onlyAdmin public {
		newManager = _newManager;
	}

	function acceptOwnership() public {
		require(msg.sender == newOwner);
		address oldaddr = owner;
		owner = newOwner;
		newOwner = address(0);
		emit TransferOwnership(oldaddr, owner);
	}

	function acceptManager() public {
		require(msg.sender == newManager);
		address oldaddr = manager;
		manager = newManager;
		newManager = address(0);
		emit TransferManager(oldaddr, manager);
	}
}
