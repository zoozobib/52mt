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

    <title>独乐乐不如众乐乐 - 摇钱木又寸</title>

    <!-- Bootstrap core CSS -->
    <link href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet">

    <script src="/static/js/jquery-1.12.3.min.js"></script>
    <script src="/static/jquery-ui-1.11.4.custom/jquery-ui.min.js"></script>
    <link rel="stylesheet" href="/static/jquery-ui-1.11.4.custom/jquery-ui.min.css">

    <!-- Custom styles for this template -->
    <link href="/static/css/dashboard.css" rel="stylesheet">

    <!-- Just for debugging purposes. Don't actually copy these 2 lines! -->
    <!--[if lt IE 9]><script src="/static/js/ie8-responsive-file-warning.js"></script><![endif]-->
    <script src="/static/js/ie-emulation-modes-warning.js"></script>

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
      <script src="//cdn.bootcss.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="//cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>

  <body>

    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container-fluid">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="#">摇钱木又寸</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav navbar-right">
            {{if eq .session ""}}
            <li><a href="/register">注册</a></li>
            <li><a href="/login">登录</a></li>
            {{else}}
            <li><a href="/index"><返回首页></a></li>
            <li><a href="/user">{{.session}}&nbsp;&nbsp;欢迎您回来</a></li>
            <li><a href="/user">用户中心</a></li>
            <li><a href="/logout">注销</a></li>
            {{end}}
            <li><a href="#">帮助</a></li>
          </ul>
        </div>
      </div>
    </nav>

    <div class="container-fluid">
      <div class="row">
        <div class="col-sm-3 col-md-2 sidebar">
          <ul class="nav nav-sidebar">
            <li {{if eq .function "index" ""}}class="active"{{end}}><a href="/user/index">收益预览 <span class="sr-only">(current)</span></a></li>
            <li {{if eq .function "money"}}class="active"{{end}}><a href="/user/money">财务管理</a></li>
            <li {{if eq .function "profile"}}class="active"{{end}}><a href="/user/profile">个人资料管理</a></li>
            <li {{if eq .function "invest"}}class="active"{{end}}><a href="/user/invest">投资记录</a></li>
          </ul>
          <ul class="nav nav-sidebar">
            <li><a style="color:#ccc" href="javascript:void(0);">内部产品(即将开放)</a></li>
          </ul>
          <ul class="nav nav-sidebar">
            <li><a style="color:#ccc" href="javascript:void(0);">积分商城(即将开放)</a></li>
          </ul>
        </div>
        <div class="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
          {{if eq .function "index" ""}}
            {{template "./common/user_index.html" .}}
          {{else if eq .function "money"}}
            {{template "./common/user_money.html" .}}
          {{else if eq .function "profile"}}
            {{template "./common/user_profile.html" .}}
          {{else if eq .function "invest"}}
            {{if eq .action "add"}}
              {{template "./common/user_invest_add.html" .}}
            {{else}}
              {{template "./common/user_invest.html" .}}
            {{end}}
          {{end}}
        </div>
      </div>
    </div>

    <!-- Bootstrap core JavaScript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
    <script src="//cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <script src="/static/js/ie10-viewport-bug-workaround.js"></script>
  </body>
</html>
