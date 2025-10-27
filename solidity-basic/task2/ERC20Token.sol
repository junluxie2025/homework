// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8;

/**
任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
合约包含以下标准 ERC20 功能：
balanceOf：查询账户余额。
transfer：转账。
approve 和 transferFrom：授权和代扣转账。
使用 event 记录转账和授权操作。
提供 mint 函数，允许合约所有者增发代币。
提示：
使用 mapping 存储账户余额和授权信息。
使用 event 定义 Transfer 和 Approval 事件。
部署到sepolia 测试网，导入到自己的钱包
 */
 contract ERC20Token{

    event Transfer(address indexed from,address indexed to,uint256 value);
    event Approval(
        address indexed owner,//合约所有者
        address indexed spender,//授权用户
        uint256 value
    );

    string private _name;
    string private _symbol;
    uint8 private _decimals;//小数位数
    uint256 private _totalSupply;
    address private _self;//合约所有者
    //账户余额
    mapping(address account=>uint256) private _balance;
    //授权记录表
    mapping(address account=>mapping(address spender=>uint256)) private _allowances;
    //XToken XTK
    constructor(string memory name_,string memory symbol_){
        _name = name_;
        _symbol = symbol_;
        _decimals = 18;
        _totalSupply = 100000;
        _self = msg.sender;
        //铸造初始代币给部署合约的
        mint(_self,_totalSupply*10**_decimals);

    }

    function name() public view returns(string memory){
        return _name;
    }

    
    function symbol() public view returns(string memory){
        return _symbol;
    }

    
    function decimals() public view returns(uint8){
        return _decimals;
    }

    
    function totalSupply() public view returns(uint256){
        return _totalSupply;
    }

    //查询账户额度
    function balanceOf(address account) external view returns(uint256){
        return _balance[account];
    }
    //查询授权额度
    function allowance(address account,address spender) external view returns(uint256){
        return _allowances[account][spender];
    }
    //转账
    function transfer(address to,uint256 value) external returns(bool){
        return _transfer(msg.sender,to,value);
    }

    function _transfer(address from,address to,uint256 value) internal returns(bool){
        require(from!=address(0),"ERC20Token: transfer to the zero address");
        require(to!=address(0),"ERC20Token: transfer from the zero address");
        uint256 fromBalance = _balance[from];
        require(fromBalance >value,"ERC20Token: transfer amount exceeds balance");
        _balance[from] = fromBalance - value;
        _balance[to] += value;
        emit Transfer(from,to,value);
        return true;
    }

    //授权
    function approve(address spender,uint256 value) external returns(bool){
        return _approve(msg.sender,spender,value);
    }

    function _approve(address owner,address spender,uint256 amount) internal returns(bool){
        require(owner!=address(0),"ERC20Token: approve owner the zero address");
        require(spender!=address(0),"ERC20Token: approve spender the zero address");
        require(_balance[owner]>amount,"ERC20Token: approve amount exceeds balance");
        //记录授权 owner允许spender花 value的代币
        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
        return true;
    }

    //代扣转账 from合约所有者  to接受代币者
    function transferFrom(address from,address to,uint256 value) external returns(bool){
        bool success = _approve(from, msg.sender,_allowances[from][msg.sender]-value);
        if (success){
             success = _transfer(from, to, value);
        }
        return success;
        
    }

    modifier onlyOwner(){
        require(msg.sender==_self,"ERC20Token: only owner allowed");
        _;
    }

    function mint(address account,uint256 amount) internal onlyOwner{
        require(account!=address(0),"ERC20Token: mint to the zero address");
        //增加总发行量
        _totalSupply += amount;
        //给目标地址价钱
        _balance[account] += amount;
        //从0地址转过来代表新发行
        emit Transfer(address(0), account, amount);
    }
 }

