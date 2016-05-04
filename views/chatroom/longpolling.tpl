<script type="text/javascript" src="/static/js/longpolling.js"></script>
<div class="row">
 	<div id="chat-column" class="span8 well">
	
	    <h3>用户名: <span id="uname">{{.UserName}}</span></h3>
		<div style="border:0px solid #ccc;">
            <ul id="chatbox"  class="message-container">
				<li>欢迎光临！</li>
            </ul>
        </div>
		<form class="form-horizontal">
		  <div class="form-group">		  
		    <div class="col-sm-10">
		      <textarea  id="sendbox" onkeydown="if(event.keyCode==13)return false;" required class="form-control" rows="3"></textarea>
		    </div>
		  </div>
		  <div class="form-group">
		    <div class="col-sm-10">
		        <button id="sendbtn" type="button" class="btn btn-info">发送</button>
		    </div>
		  </div>
		</form>	
	</div>
</div>