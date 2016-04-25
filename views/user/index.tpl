<div class="row">
	<div class="col-md-8">
		<table class="table table-striped">
		<tr>
			<td>用户邮箱</td>
			<td>手机</td>
			<td>创建时间</td>
			<td>创建IP</td>
		</tr>
		{{range $key,$value:=.list}}
		<tr>
			<td>{{$value.email}}</td>
			<td>{{$value.mobile}}</td>
			<td>{{$value.create_time}}</td>
			<td>{{$value.create_ip}}</td>
		</tr>
		{{end}}
		</table>
	</div>
</div>