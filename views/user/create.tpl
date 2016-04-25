<div class="row">
	<div class="col-md-8">
		<form action="http://127.0.0.1:8080/create" method="post">
			<h3 class="form-title">添加用户</h3>
			<div class="alert alert-danger" role="alert">{{.flash.error}}</div>
			<div class="form-group">
			   <label for="username">邮箱</label>
			   <input type="text" class="form-control" name="username" placeholder="请输入邮箱">
			</div>
			<div class="form-group">
			   <label for="password">密码</label>
			   <input type="password" class="form-control" name="password" placeholder="请输入密码">
			</div>
			<div class="form-group">
			   <label for="rpassword">确认密码</label>
			   <input type="password" class="form-control" name="rpassword" placeholder="请确认密码">
			</div>
			<div class="form-group">
			   <label for="mobile">手机</label>
			   <input type="text" class="form-control" name="mobile" placeholder="请输入手机">
			</div>
			<button type="submit" class="btn btn-default">保存</button>
		</form>
	</div>
</div>