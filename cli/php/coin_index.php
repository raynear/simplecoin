<?php

$CoinJsonStr = file_get_contents('ethereum.json');
$json = json_decode($CoinJsonStr, true);

echo("
<html>
<head>
    <title>".$json['CoinName']." Test</title>
    <style>
        .button_large {
            border:1x solid #333333;
            background-Color:#E4E4E4;
            font:12px;
            font-weight:bold;
            color:#444444;
            width:130;height:30;
        }
        .box {
            border:1px solid #333333;
            backgroud-color:lightgrey;
            padding: 10px;
        }
    </style>
    <script src='http://code.jquery.com/jquery-latest.min.js'></script>
    <script>
        $(document).ready(function() {
            var URL = 'CoinTest.php';

            $('input#RunButton').click(function(event){
                var paramLen = $(event.target).parents('div').children('#param').length;

                var params = '\"DaemonURL\":\"".$json['URL']."\",';
                params = params+'\"SSL\":\"".$json['SSL']."\",';
                params = params+'\"RPCTYPE\":\"".$json['RPCTYPE']."\",';
                params = params+'\"USER\":\"".$json['USER']."\",';
                params = params+'\"PASS\":\"".$json['PASS']."\",';
                params = params+'\"funcname\":\"'+$(event.target).parents('div').attr('id') + '\"';
                if(paramLen > 0){
                    params = params+\",\";
                }

                for(var i=0 ; i<paramLen ; i++) {
                    var name = $(event.target).parents('div').children('#param').eq(i).attr('name');
                    var value = $(event.target).parents('div').children('#param').eq(i).val();
                    params = params + '\"' + name  + '\":' + JSON.stringify(value);
                    if(paramLen > i+1){
                        params = params + ',';
                    }
                }
                params = '{'+params+'}';
//                alert(params);
                params = JSON.parse(params);

                $.ajax({
                    type:'post',
                    url:URL,
                    dataType:'json',
                    data: params,
                    success:function(data) {
                        $(event.target).parents('div').children('pre').children('#resultDiv').html(data['data']);
                    },
                    error:function(xhr, status, error) {
                        alert('code:'+xhr.status+'\\n'+'message:'+xhr.responseText+'\\n'+'error:'+error);
                    }
                });
            });
        })
    </script>
</head>
<body>
");

// list안에 json 있는 경우 처리 아직 안됨
// 그 부분 처리 하고 나면 문제 없을것으로 예상됨

foreach($json['functions'] as $aFunction) {
    foreach($aFunction as $key => $value) {
        if($key == 'function') {
            echo("<div class='box' id='".$value."'>\n");
            echo("\t<font size='3' color='dimgrey'><b>".$value."</b></font><br>\n");
        }
        if($key == 'parameters') {
            foreach($value as $pName => $pValue) {
                if(gettype($pValue) == "boolean") {
                    $pValue = $pValue ? "true" : "false";
                }
                echo("\t<label>".$pName."</label><br>\n");
                echo("\t<input type='text' size='60' id='param' name='".$pName."' value='".$pValue."'/><br>\n");
            }
        }
    }
    echo("\t<input type='button' id='RunButton' class='button_large' value='Run'/>\n");
    echo("\t<pre><div id='resultDiv'></div></pre>\n");
    echo("</div>\n<br>\n");
}

echo("
</body>
</html>
");
?>
