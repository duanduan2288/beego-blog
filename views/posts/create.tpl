<script type="text/javascript" charset="utf-8" src="/static/js/ueditor/ueditor.config.js"></script>
<script type="text/javascript" charset="utf-8" src="/static/js/ueditor/ueditor.all.min.js"> </script>
<script type="text/javascript" charset="utf-8" src="/static/js/ueditor/lang/zh-cn/zh-cn.js"></script>
<script type="text/javascript" charset="utf-8" src="/static/js/jquery-validation/dist/jquery.validate.min.js"></script>
<script type="text/javascript" charset="utf-8" src="/static/js/jquery-validation/dist/jquery.validate.bootstrap.js"></script>
<script type="text/javascript" charset="utf-8" src="/static/js/jquery-validation/dist/additional-methods.min.js"></script>
<div class="row">
	<div class="col-md-8">
		<form action="http://127.0.0.1:8080/posts/create" method="post" id="form_posts">
			<h3 class="form-title">添加文章</h3>
			<div class="alert alert-danger" style="display:none;" role="alert">{{.flash.error}}</div>
			<div class="form-group">
			   <label for="title">标题</label>
			   <input type="text" class="form-control" name="title" id="title" placeholder="请输入标题">
			</div>
			<div class="form-group">
			   <label for="content">内容</label>
			    <script id="editor" type="text/plain" style="width:1024px;height:500px;"></script>
			</div>
			<input type="hidden" id="content" name="content" value="" />
			<button type="button" id="send_form" class="btn btn-default">保存</button>
		</form>
	</div>
</div>
<script type="text/javascript">

    //实例化编辑器
    var ue = UE.getEditor('editor');
	
	$(function(){
		
		$('#form_posts').validate({
	        errorElement: 'span', //default input error message container
	        errorClass: 'help-block', // default input error message class
	        focusInvalid: false, // do not focus the last invalid input
	        rules: {
	            title : {
	                required:true
	            },
	            content : {
	                required:true
	            }	    
	        },
	
	        messages: {
	            title : {
	                required:'请输入标题'
	            },
	            content : {
	                required:'请输入内容'
	            }
	        },
	
	        invalidHandler: function (event, validator) { //display error alert on form submit
	            $('.alert-danger', $('#form_posts')).show();
	        },
	
	        highlight: function (element) { // hightlight error inputs
	            $(element)
	                .closest('.form-group').addClass('has-error'); // set error class to the control group
	        },
	
	        success: function (label) {
	            label.closest('.form-group').removeClass('has-error');
	            label.remove();
	        },
	
	        errorPlacement: function (error, element) {
	
	            error.insertAfter(element.closest('.form-control'));
	        },
	
	        submitHandler: function (form) {
	            form.submit();
	        }
    	});
		
		$("#send_form").click(function(){
			
			var content = UE.getEditor('editor').getContent();
			$('#content').val(content);
			
			if($('#form_posts').validate().form()){
				$('#form_posts').submit();
			}
		});
	});
	
</script>