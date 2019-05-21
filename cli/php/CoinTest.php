<?php
$SSL = $_POST['SSL'];
unset($_POST['SSL']);
$RPCTYPE = $_POST['RPCTYPE'];
unset($_POST['RPCTYPE']);
$URL = $_POST['DaemonURL'];
unset($_POST['DaemonURL']);
$USER = $_POST['USER'];
unset($_POST['USER']);
$PASS = $_POST['PASS'];
unset($_POST['PASS']);

function CurlSend($method, $params) {
    global $URL;
    global $SSL;
    global $USER;
    global $PASS;
    global $RPCTYPE;

    $id = 0;

    // performs the HTTP POST using curl
    $curl = curl_init();

    if($RPCTYPE=="NXT") {
        $paramstr = "";
        $paramlen = count($params);
        foreach ($params as $key => $value) {
            if(0 === --$paramlen) {
                $paramstr = $paramstr.$key."=".urlencode($value);
            }
            else {
                $paramstr = $paramstr.$key."=".urlencode($value)."&";
            }
        }
        if($paramstr == "") {
            $request = $URL."requestType=".$method;
        }
        else {
            $request = $URL."requestType=".$method."&".$paramstr;
        }
        curl_setopt($curl, CURLOPT_URL, $request);
    }
    if($RPCTYPE=="JSONRPC") {

        // only for ethereum
        if(!array_key_exists(0, $params)){
            $params = array($params);
        }

        $request = json_encode(array(
                    'json_rpc' => "2.0",
                    'method' => $method,
                    'params' => $params,
                    'id' => $id
                ), JSON_PRETTY_PRINT);
        curl_setopt($curl, CURLOPT_URL, $URL."/json_rpc");
    	curl_setopt($curl, CURLOPT_POSTFIELDS, $request);
    }

    if($RPCTYPE=="RESTAPI") {
        $request = json_encode(array($params));
        curl_setopt($curl, CURLOPT_URL, $URL."/".$method);
//    	curl_setopt($curl, CURLOPT_POSTFIELDS, json_encode($params));
//        print_r($params);
//        print_r(array($params));
//        print_r(json_encode(array($params)));
//        print_r($request);
    }

    if($USER != "" && $PASS != "") {
       curl_setopt($curl, CURLOPT_USERPWD, $USER.":".$PASS);
    }

    if($SSL=="true") {
        curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, FALSE);
        curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, FALSE);
    }

    curl_setopt($curl, CURLOPT_POST, TRUE);
    curl_setopt($curl, CURLOPT_HTTPHEADER, array('Content-Type: application/json'));
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);

    $response = curl_exec($curl);
    curl_close($curl);

    // process response
    if (!$response) {
        throw new Exception('Unable to connect to '.$URL, 0);
    }
    $response = json_decode($response,true);

    if (!is_null(isset($response['error'])?$response['error']:null)) {
        throw new Exception('Request error: '.print_r($response['error'],1),2);
    }

    return $response;
}


$funcname = $_POST['funcname'];
unset($_POST['funcname']);

$params = array();
foreach($_POST as $key => $value){
    if(strtolower($value)=="true") $value=true;
    if(strtolower($value)=="false") $value=true;
    $params[$key]=$value;
}
try {
    $RPCresult = CurlSend($funcname, $params);
} catch (Exception $e) {
    die($e->getMessage());
}
$result['success'] = true;
$result['data'] = json_encode($RPCresult, JSON_PRETTY_PRINT);
echo(json_encode($result));


/*
{
    "apiProxy": false,
    "correctInvalidFees": false,
    "ledgerTrimKeep": 30000,
    "maxAPIRecords": 100,
    "blockchainState": "DOWNLOADING",
    "currentMinRollbackHeight": 285305,
    "numberOfBlocks": 286106,
    "isTestnet": true,
    "includeExpiredPrunable": true,
    "isLightClient": false,
    "services": [
        "CORS"
    ],
    "requestProcessingTime": 0,
    "version": "2.0.14",
    "maxRollback": 800,
    "lastBlock": "15155525371215195387",
    "application": "Ardor",
    "isScanning": false,
    "isDownloading": false,
    "cumulativeDifficulty": "28937799196895973",
    "lastBlockchainFeederHeight": 286099,
    "maxPrunableLifetime": 7776000,
    "time": 16987723,
    "lastBlockchainFeeder": "192.168.56.1"
}
*/


/*
// 지갑 암호화?
$parameters = array("account"=>"ARDOR-RSLK-BEB2-EJRP-24RNR");
try {
    $result = $WalletRPC->function($parameters);
} catch (Exception $e) {
    die($e->getMessage());
}
echo("<p>: </p>");
echo("<pre>");
echo(json_encode($result, JSON_PRETTY_PRINT));
echo("</pre>");
*/
/*
*/


/*
// 주소 생성?
$parameters = array("account"=>"ARDOR-RSLK-BEB2-EJRP-24RNR");
try {
    $result = $WalletRPC->function($parameters);
} catch (Exception $e) {
    die($e->getMessage());
}
echo("<p>: </p>");
echo("<pre>");
echo(json_encode($result, JSON_PRETTY_PRINT));
echo("</pre>");
*/
/*
*/

/* raynear
// getBalance
$parameters = array("account"=>"ARDOR-RSLK-BEB2-EJRP-24RNR",
                    "chain"=>"ARDR");
try {
    $result = $WalletRPC->getBalance($parameters);
} catch (Exception $e) {
    die($e->getMessage());
}
echo("<p>getBalance : </p>");
echo("<pre>");
echo(json_encode($result, JSON_PRETTY_PRINT));
echo("</pre>");
*/
/*
{
    "unconfirmedBalanceNQT": "148800000000",
    "balanceNQT": "148800000000",
    "requestProcessingTime": 0
}
*/


/*
// 블록체인 정보확인 명령어와 같음
$parameters = array("account"=>"ARDOR-RSLK-BEB2-EJRP-24RNR");
try {
    $result = $WalletRPC->function($parameters);
} catch (Exception $e) {
    die($e->getMessage());
}
echo("<p>: </p>");
echo("<pre>");
echo(json_encode($result, JSON_PRETTY_PRINT));
echo("</pre>");
*/
/*
*/

/* raynear

// getAccountPublicKey
$parameters = array("account"=>"ARDOR-RSLK-BEB2-EJRP-24RNR");
try {
    $result = $WalletRPC->getAccountPublicKey($parameters);
} catch (Exception $e) {
    die($e->getMessage());
}
$AccountPubKey = $result["publicKey"];

$secret = "pillow%20freeze%20pathetic%20law%20mutter%20loose%20rip%20center%20street%20place%20threaten%20always";
$UserMessage = "UserMessageTest";
// 전송 명령어
$parameters = array("recipient"=>"ARDOR-4NM4-7NPP-989A-5XMCJ",
                    "chain"=>"ARDR",
                    "amountNQT"=>"1000000000",
                    "publicKey"=>$AccountPubKey,
                    "message"=>$UserMessage,
                    "secretPhrase"=>$secret,
                    "messageIsPrunable"=>"true");
try {
    $result = $WalletRPC->sendMoney($parameters);
} catch (Exception $e) {
    die($e->getMessage());
}
echo("<p>sendMoney </p>");
echo("<pre>");
echo(json_encode($result, JSON_PRETTY_PRINT));
echo("</pre>");
/*
{
    "minimumFeeFQT": "200000000",
    "signatureHash": "5783f293d70dc881a349b73902b625a7262b863a8183630ab1cfa43e132e57c0",
    "transactionJSON": {
        "senderPublicKey": "b461c113a745b0775181f9ab38a09a639d905d6567b53e2c8daab119e697b21c",
        "chain": 1,
        "signature": "4f82683f74065b3251c0f613c93e36f6b9624bfb5545a943c7d645b7ba8b7d090f5c092a474539eaa388b04e540f357e6cbd0276d743b809aeb6b3f4da32a6e7",
        "feeNQT": "200000000",
        "type": -2,
        "fullHash": "cf0676168d7e20be3d1cba509a5fa5e329069ee717f2c07bba4d907a448288ef",
        "version": 1,
        "phased": false,
        "ecBlockId": "12055439847609264363",
        "signatureHash": "5783f293d70dc881a349b73902b625a7262b863a8183630ab1cfa43e132e57c0",
        "attachment": {
            "version.FxtPayment": 0
        },
        "senderRS": "ARDOR-RSLK-BEB2-EJRP-24RNR",
        "subtype": 0,
        "amountNQT": "1000000000",
        "sender": "851992091579638353",
        "recipientRS": "ARDOR-4NM4-7NPP-989A-5XMCJ",
        "recipient": "4047184848099562082",
        "ecBlockHeight": 285385,
        "deadline": 15,
        "transaction": "13700089210893371087",
        "timestamp": 16988974,
        "height": 2147483647
    },
    "unsignedTransactionBytes": "01000000fe00012e3b03010f00b461c113a745b0775181f9ab38a09a639d905d6567b53e2c8daab119e697b21c6252512b2d7d2a3800ca9a3b0000000000c2eb0b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c95a0400ebe8d63fab864da700000000",
    "broadcasted": true,
    "requestProcessingTime": 10,
    "transactionBytes": "01000000fe00012e3b03010f00b461c113a745b0775181f9ab38a09a639d905d6567b53e2c8daab119e697b21c6252512b2d7d2a3800ca9a3b0000000000c2eb0b000000004f82683f74065b3251c0f613c93e36f6b9624bfb5545a943c7d645b7ba8b7d090f5c092a474539eaa388b04e540f357e6cbd0276d743b809aeb6b3f4da32a6e7c95a0400ebe8d63fab864da700000000",
    "fullHash": "cf0676168d7e20be3d1cba509a5fa5e329069ee717f2c07bba4d907a448288ef"
}
*/

/* secretphrase 안 넣으면 sign과 broadcast를 따로 해야 됨.
if ($result["broadcasted"]!=true) {

    print_r("<p>transaction not broadcasted make it sure</p>");
    // sign
    $parameters = array("unsignedTransactionJSON"=>json_encode($result["transactionJSON"]),//"unsignedTransactionBytes"=>$result["unsignedTransactionBytes"],
                        "secretPhrase"=>$secret);
    try {
        $result = $WalletRPC->signTransaction($parameters);
    } catch (Exception $e) {
        die($e->getMessage());
    }
    echo("<p>signTransaction </p>");
    echo("<pre>");
    echo(json_encode($result, JSON_PRETTY_PRINT));
    echo("</pre>");

    // broadcast
    $parameters = array("transactionJSON"=>json_encode($result["transactionJSON"]));//"transactionBytes"=>$result["transactionBytes"]);
    try {
        $result = $WalletRPC->broadcastTransaction($parameters);
    } catch (Exception $e) {
        die($e->getMessage());
    }
    echo("<p>broadcastTransaction </p>");
    echo("<pre>");
    echo(json_encode($result, JSON_PRETTY_PRINT));
    echo("</pre>");
}
*/

/* raynear
// 트랜젝션 리스트 조회
$parameters = array("chain"=>"ARDR", "account"=>"ARDOR-RSLK-BEB2-EJRP-24RNR");
try {
    $result = $WalletRPC->getBlockchainTransactions($parameters);
} catch (Exception $e) {
    die($e->getMessage());
}
echo("<p>getBlockchainTransactions: </p>");
echo("<pre>");
echo(json_encode($result, JSON_PRETTY_PRINT));
echo("</pre>");
/*
{
    "requestProcessingTime": 0,
    "transactions": [
        {
            "signature": "d1c1a454b0a72f6ac76ac1bd9ae23a2b73f31c3edcc99e3a18e9bc142fbb4908e957d4e0a7004883745ca1f96252c86674a96bfc8d5cde5643c049f15e81b2d6",
            "transactionIndex": 0,
            "type": -2,
            "phased": false,
            "ecBlockId": "15869792346936013547",
            "signatureHash": "9cf74b301259e27d8e3679e423cbb3ee9e0015b12b5833390a6ac9614554e434",
            "attachment": {
                "version.FxtPayment": 0
            },
            "senderRS": "ARDOR-VS8T-QYQS-SK9H-5BYBB",
            "subtype": 0,
            "amountNQT": "150000000000",
            "recipientRS": "ARDOR-RSLK-BEB2-EJRP-24RNR",
            "block": "3782359136430544025",
            "blockTimestamp": 16817659,
            "deadline": 1440,
            "timestamp": 16817011,
            "height": 283305,
            "senderPublicKey": "7cc7a8404d9c69b860035bce8785f162dae2076eb21d55349aab7b856de91f53",
            "chain": 1,
            "feeNQT": "1100000000",
            "confirmations": 2800,
            "fullHash": "fd7efc6f3a2b7c112f768131563de2914385c960858a28400eb263e36c5e0a75",
            "version": 1,
            "sender": "3794223001810886873",
            "recipient": "851992091579638353",
            "ecBlockHeight": 282571,
            "transaction": "1259929525743812349"
        }
    ]
}
*/


/* raynear
// 단일 트랜젝션 상세보기
$parameters = array("fullHash"=>"fd7efc6f3a2b7c112f768131563de2914385c960858a28400eb263e36c5e0a75",
                    "chain"=>"ARDR");
try {
    $result = $WalletRPC->getTransaction($parameters);
} catch (Exception $e) {
    die($e->getMessage());
}
echo("<p>getTransaction: </p>");
echo("<pre>");
echo(json_encode($result, JSON_PRETTY_PRINT));
echo("</pre>");
/*
{
    "signature": "d1c1a454b0a72f6ac76ac1bd9ae23a2b73f31c3edcc99e3a18e9bc142fbb4908e957d4e0a7004883745ca1f96252c86674a96bfc8d5cde5643c049f15e81b2d6",
    "transactionIndex": 0,
    "type": -2,
    "phased": false,
    "ecBlockId": "15869792346936013547",
    "signatureHash": "9cf74b301259e27d8e3679e423cbb3ee9e0015b12b5833390a6ac9614554e434",
    "attachment": {
        "version.FxtPayment": 0
    },
    "senderRS": "ARDOR-VS8T-QYQS-SK9H-5BYBB",
    "subtype": 0,
    "amountNQT": "150000000000",
    "recipientRS": "ARDOR-RSLK-BEB2-EJRP-24RNR",
    "block": "3782359136430544025",
    "blockTimestamp": 16817659,
    "deadline": 1440,
    "timestamp": 16817011,
    "height": 283305,
    "senderPublicKey": "7cc7a8404d9c69b860035bce8785f162dae2076eb21d55349aab7b856de91f53",
    "chain": 1,
    "feeNQT": "1100000000",
    "requestProcessingTime": 0,
    "confirmations": 2800,
    "fullHash": "fd7efc6f3a2b7c112f768131563de2914385c960858a28400eb263e36c5e0a75",
    "version": 1,
    "sender": "3794223001810886873",
    "recipient": "851992091579638353",
    "ecBlockHeight": 282571,
    "transaction": "1259929525743812349"
}
*/
?>
