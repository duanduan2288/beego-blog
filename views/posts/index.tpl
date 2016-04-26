<div class="row">
	<div class="col-md-8">
		<div class="alert alert-danger" style="display:none;" role="alert">{{.flash.error}}</div>
		<table class="table table-striped">
			<tr>
				<td>标题</td>
				<td>创建时间</td>
				<td>状态</td>
				<td>操作</td>
			</tr>
			{{range $key,$value:=.list}}
			<tr>
				<td><a href='{{urlfor "PostsController.Detail" ":id" $value.id}}'>{{$value.title}}</a></td>
				<td>{{$value.create_time}}</td>
				<td>{{$value.status}}</td>
				<td>
					<a href='{{urlfor "PostsController.Edit" ":id" $value.id}}'>编辑</a>
					{{if "正常" == $value.status}}
						<a href='{{urlfor "PostsController.Delete" ":id" $value.id}}'>删除</a>
					{{else}}
						<a href='{{urlfor "PostsController.Delete" ":id" $value.id}}'>回复</a>
					{{endif}}
				</td>
			</tr>
			{{end}}
		</table>
	</div>
</div>