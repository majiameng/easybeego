    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 网站配置
 * @author Tinymeng
 * @since 2021/7/26
 */
layui.use(['func'], function () {
    var $ = layui.jquery;
    var form = layui.form;
    var func = layui.func;

    /**
     * 提交表单
     */
    form.on('submit(submitForm2)', function (data) {
        // 网络请求
        var loadIndex = layer.load(2);
        $.ajax({
            type: "POST",
            url: '/configweb/index',
            data: JSON.stringify(data.field),
            contentType: "application/json",
            dataType: "json",
            beforeSend: function () {
                // TODO...
            },
            success: function (res) {
                layer.close(loadIndex);
                if (res.code == 0) {
                    //0.5秒后关闭
                    layer.msg(res.msg, {icon: 1, time: 500});
                } else {
                    layer.msg(res.msg, {icon: 5});
                    return false;
                }
            },
            error: function () {
                layer.msg("AJAX请求异常");
            }
        });
        return false;
    });
});
