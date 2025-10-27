pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

//tokenURI
contract ERC821Token is ERC721URIStorage,Ownable{
    
    uint public MAX_APES = 100;

    constructor(address initialOwner) ERC721("XNFT","XNFT") Ownable(initialOwner){

    }
    
    function mintNFT(address to,string memory tokenURI,uint256 tokenId) public onlyOwner{
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
    }
}