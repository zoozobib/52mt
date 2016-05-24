
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

    <title>登录摇钱木又寸</title>

    <!-- Bootstrap core CSS -->
    <link href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet">

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

      <form class="form-signin" action="./login" method="post">
        <h2 class="form-signin-heading">欢迎回来，登录摇钱木又寸</h2>
        <label for="inputEmail" class="sr-only">邮件地址</label>
        <input name="username" id="inputEmail" class="form-control" placeholder="邮件地址" required autofocus>
        <label for="inputPassword" class="sr-only">密码</label>
        <input type="password" name="password" id="inputPassword" class="form-control" placeholder="密码" required>
		<!--
        <div class="checkbox">
          <label>
            <input type="checkbox" value="remember-me"> 下次免输入
          </label>
        </div>
		-->
        <button class="btn btn-lg btn-primary btn-block" type="submit">登录</button>
        <div style="padding-top:20px;">
        	<p style="float:left;"><a href="./index">>返回首页</a></a></p>
			<p style="float:right;"><a href="./register">还没有账号？立即加入！</a></p>
        </div>
      </form>

    </div> <!-- /container -->


    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <script src="./static/js/ie10-viewport-bug-workaround.js"></script>
  </body>
</html>
