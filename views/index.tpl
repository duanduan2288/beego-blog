<div class="container" style="width: 500px;">
    <h3>设置信息</h3>
    <br>
    <form action="/join" method="post" class="form-horizontal">
        <div class="form-group">
            <label class="col-md-3 control-label">用户名: </label>
            <div class="col-md-5">
                  {{.username}}
            </div>
        </div>

        <div class="form-group">
            <label class="col-md-3 control-label">技术: </label>
            <div class="col-md-5">
                <select class="form-control" name="tech">
                    <option value="longpolling">长轮询</option>
                    <option value="websocket">WebSocket</option>
                </select>
            </div>
        </div>

        <div class="form-group">
            <div class="col-sm-offset-3 col-sm-10">
                <button type="submit" class="btn btn-info">进入聊天室</button>
            </div>
        </div>
    </form>
</div>