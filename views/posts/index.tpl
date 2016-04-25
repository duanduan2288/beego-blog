<div class="row">
	<div class="col-md-8">
		<table class="table table-striped">
		<tr>
			<td>标题</td>
			<td>创建时间</td>
			<td>操作</td>
		</tr>
		{{range $key,$value:=.list}}
		<tr>
			<td><a href='{{urlfor "PostsController.Detail" ":id" $value.id}}'>{{$value.title}}</a></td>
			<td>{{$value.create_time}}</td>
			<td>删除</td>
		</tr>
		{{end}}
		</table>
	</div>
</div>