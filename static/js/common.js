function r_submit(){
  var username = $("#inputEmail").val()
  var password = $("#inputPassword").val()
  var vcode = $("#inputMobile").val()

  if (username == "") {
    alert("请填写用户名");
    return false;
  }

  if (password == "" || password.length < 8) {
    alert("请填写密码或密码长度不足8位");
    return false;
  }

  if (vcode =="") {
    alert("请获取验证码");
    return false;
  }

  $("#register").submit();
}
