{{template "./common/header.html" .}}

      <!-- Jumbotron -->
      <div class="jumbotron" style="text-align:left;">
        <h2>易通贷_成熟投资人选择的安全投资平台</h1>
        <p class="lead text-danger">返利配比：2%</p>
        <p class="lead">P2P理财,P2P网贷,投资理财,网络理财,网络投资,小额投资理财,如何投资理财,个人理财,家庭理财,易通贷</p>
        <p>易通贷，中国领先的P2P网络投资理财平台，注册资金1个亿。投资无风险，本息有保障，门槛低，收益高，100元也能做理财。是个人小额投资、家庭理财的理想选择。</p>
        <p class="lead text-danger">需投资1个月及以上的标，新手标1个月标也有效，投资金额0-50000元。</p>
        <p><a class="btn btn-lg btn-success" href="javascript:void(0);" id="yitong" role="button" target="_blank">立即参与赚取积分</a></p>
      </div>

      <div class="jumbotron" style="text-align:left;">
        <h2>财路通豆蔓理财官网-投资理财_中国首家保险公司承保的互联网金融理财平台</h1>
        <p class="lead text-danger">返利配比：2%</p>
        <p class="lead">财路通豆蔓理财,财路通豆蔓理财官网,P2P网贷,投资理财,网络信贷</p>
        <p>财路通豆蔓理财官网 荣获中国电子商务协会优秀信用AAA评级,中国首家保险公司承保的网络信贷平台。提供安全的网络信贷,P2P网贷,投资理财,小额贷款,短期贷款等服务。财路通豆蔓理财互联网理财超低门槛,超高收益！</p>
        <p class="lead text-danger">需投资1个月及以上的标，新手标1个月标也有效，投资金额0-50000元。</p>
        <p><a class="btn btn-lg btn-success" href="javascript:void(0);" id="cailutong" role="button" target="_blank">立即参与赚取积分</a></p>
      </div>

      <div id="dialog" title="摇钱木又寸提示您" style="display:none;">为了更好的确认您的收益，请登陆平台后继续操作.</div>

      <script>
      //$( "#dialog" ).dialog({ autoOpen: false });
      $( "#yitong" ).click(function() {
        //http://ad.etongdai.com/cpa/enter/zhonghulian?target_url=https://app.etongdai.com?friendId=MjMxOTI0OQ==
        {{if eq .session ""}}
        $( "#dialog" ).dialog({
          modal:true,
          click: function() {
              $( this ).dialog( "open" );
          },
          buttons: { "没有账号?注册": function()
                    {
                      $(this).dialog("close");
                      window.location.href = "./register";
                    } ,
                    "登陆": function()
                    {
                      $(this).dialog("close");
                      window.location.href = "./login";
                    }
                  }

        });
        {{else}}
        window.location.href = "http://app.etongdai.com/register/landingPage7?friendId=MjMxOTI0OQ==";
        {{end}}
      });


      //
      $( "#cailutong" ).click(function() {
        //https://www.zlot.cn/register/xlwcfspread?hmsr=zhonghulian&hmpl=&hmcu=&hmkw=&hmci=
        {{if eq .session ""}}
        $( "#dialog" ).dialog({
          modal:true,
          click: function() {
              $( this ).dialog( "open" );
          },
          buttons: { "没有账号?注册": function()
                    {
                      $(this).dialog("close");
                      window.location.href = "./register";
                    } ,
                    "登陆": function()
                    {
                      $(this).dialog("close");
                      window.location.href = "./login";
                    }
                  }

        });
        {{else}}
        window.location.href = "https://www.zlot.cn/register/xlwcfspread?hmsr=zhonghulian&hmpl=&hmcu=&hmkw=&hmci=";
        {{end}}
      });
      </script>


      <!-- Example row of columns -->
      <div class="row">
        <div class="col-lg-4">
          <h2>敬请期待</h2>
          <p class="text-danger">随时有更新，惊喜等着你</p>
          <p>期待....</p>
          <p><a class="btn btn-primary" target="_blank" href="https://www.zlot.cn/register/xlwcfspread?hmsr=zhonghulian&hmpl=&hmcu=&hmkw=&hmci=" role="button">开始赚钱 &raquo;</a></p>
        </div>
        <div class="col-lg-4">
          <h2>敬请期待</h2>
          <p class="text-danger">随时有更新，惊喜等着你</p>
          <p>期待....</p>
          <p><a class="btn btn-primary" target="_blank" href="https://www.zlot.cn/register/xlwcfspread?hmsr=zhonghulian&hmpl=&hmcu=&hmkw=&hmci=" role="button">开始赚钱 &raquo;</a></p>
        </div>
        <div class="col-lg-4">
          <h2>敬请期待</h2>
          <p class="text-danger">随时有更新，惊喜等着你</p>
          <p>期待....</p>
          <p><a class="btn btn-primary" target="_blank" href="https://www.zlot.cn/register/xlwcfspread?hmsr=zhonghulian&hmpl=&hmcu=&hmkw=&hmci=" role="button">开始赚钱 &raquo;</a></p>
        </div>
      </div>
      <!--
      <div class="row">
        <div class="col-lg-4">
          <h2>Safari bug warning!</h2>
          <p class="text-danger">As of v8.0, Safari exhibits a bug in which resizing your browser horizontally causes rendering errors in the justified nav that are cleared upon refreshing.</p>
          <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui. </p>
          <p><a class="btn btn-primary" href="#" role="button">View details &raquo;</a></p>
        </div>
        <div class="col-lg-4">
          <h2>Heading</h2>
          <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui. </p>
          <p><a class="btn btn-primary" href="#" role="button">View details &raquo;</a></p>
       </div>
        <div class="col-lg-4">
          <h2>Heading</h2>
          <p>Donec sed odio dui. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Vestibulum id ligula porta felis euismod semper. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa.</p>
          <p><a class="btn btn-primary" href="#" role="button">View details &raquo;</a></p>
        </div>
      </div>
      -->
{{template "./common/footer.html" .}}
