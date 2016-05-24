
<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="../../favicon.ico">

    <title>注册成为摇钱树会员</title>

    <!-- Bootstrap core CSS -->
    <link href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet">
	  <script src="./static/js/jquery-1.12.3.min.js"></script>
    <script src="./static/jquery-ui-1.11.4.custom/jquery-ui.min.js"></script>
    <link rel="stylesheet" href="./static/jquery-ui-1.11.4.custom/jquery-ui.min.css">
    <script src="./static/js/common.js"></script>

    <!-- Custom styles for this template -->
    <link href="./static/css/signin.css" rel="stylesheet">

    <!-- Just for debugging purposes. Don't actually copy these 2 lines! -->
    <!--[if lt IE 9]><script src="./static/js/ie8-responsive-file-warning.js"></script><![endif]-->
    <script src="./static/js/ie-emulation-modes-warning.js"></script>

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
      <script src="//cdn.bootcss.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="//cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>

  <body>

    <div class="container">

      <form class="form-signin" action="./register" method="post" id="register">
        <h2 class="form-signin-heading">加入摇钱木又寸</h2>
        <label for="inputEmail" class="sr-only">手机号码</label>
        <input name="username" id="inputEmail" class="form-control" placeholder="手机号码" required autofocus>
        <label for="inputPassword" class="sr-only">密码</label>
        <input type="password" name="password" id="inputPassword" class="form-control" placeholder="密码(8位以上)" required>
        <div>
          <div style="float:left;width:50%;">
            <label for="inputMobile" class="sr-only">手机验证码</label>
            <input type="text" name="vcode" id="inputMobile" class="form-control" placeholder="手机验证码" required>
          </div>
          <div style="float:right;"><input type="button" id="inputMobileButton" class="btn btn-lg btn-default" value="获取验证码"/></div>
        </div>

        <button class="btn btn-lg btn-primary btn-block" type="button" onclick="javascript:r_submit();">注册</button>

        <div style="padding-top:20px;">
          <p><a href="./index">>返回首页</a></p>
        </div>
      </form>
      <script>
        var wait=60;
        function vcode(o) {
          if (wait == 0) {
            o.removeAttribute("disabled");
            o.value="获取验证码";
            wait = 60;
          } else {
            o.setAttribute("disabled", true);
            o.value="重新发送(" + wait + ")";
            wait--;
            setTimeout(function() {
              vcode(o)
            }, 1000)
          }
        }
        document.getElementById("inputMobileButton").onclick=function(){
          var myreg = /^(((13[0-9]{1})|(15[0-9]{1})|(17[0-9]{1})|(18[0-9]{1}))+\d{8})$/;
          var account = $("#inputEmail").val()
          if(!myreg.test(account))
          {
              alert('请输入有效的手机号码！');
              return false;
          }

          if ( account != "" ) {
            $.getJSON("./register/"+ account +"/sms",function(result){
              //alert(result.status);
              if ( result.status == "1" ) {
                vcode(document.getElementById("inputMobileButton"));
              }else{
                alert("短信发送异常!请联系管理员");
              }
            });
          }else{
            alert("请您输入手机号码!");
          }
        }
      </script>

    </div> <!-- /container -->


    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <script src="./static/js/ie10-viewport-bug-workaround.js"></script>
  </body>
</html>
