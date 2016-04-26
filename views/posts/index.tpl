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
					{{if eq $value.status "正常" }}
						<a href="javascript:;" class="delete" data-id="{{$value.id}}" data-value="删除">删除</a>
					{{else}}
						<a href="javascript:;" class="delete" data-id="{{$value.id}}" data-value="正常">恢复</a>
					{{end}}
				</td>
			</tr>
			{{end}}
		</table>
	</div>
</div>
<script type="text/javascript">
	$(function(){	  
		$('.table').on("click",".delete",function(){
			var id = $(this).data("id");
			var status = $(this).data("value");
			$.ajax({
				type:"post",
				url:"http://127.0.0.1:8080/posts/delete",
				data:{"id":id,"status":status},
				dataType:"json",
				success:function(data){
					alert(data.message);
					window.location.reload();
				},
				error:function(){
					alert("删除失败");
				}
			});
		});
		
	});
	
</script>