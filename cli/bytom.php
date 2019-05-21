<?php
$URL = "http://192.168.56.104";
$PORT = "9888";
$USER = "raynear";
$PASS = "30edccbc99670ab4796e7dbcf34f2b29cf4f410dc605458c8da2bf100a873257";

function CurlSend($method, $params) {
    $curl = curl_init();
    curl_setopt($curl, CURLOPT_URL, $URL.":".$PORT."/".$method);
    curl_setopt($curl, CURLOPT_USERPWD, $USER.":".$PASS);
    curl_setopt($curl, CURLOPT_POST, TRUE);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($curl, CURLOPT_POSTFIELDS, $params);

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

function main() {
    $resp = CurlSend("net-info", "");
    print_r($resp);
    $resp = CurlSend("list-balances", "");
    print_r($resp);
}

?>
