    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 个人中心
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['form', 'element', 'admin', 'func'], function () {
    var $ = layui.jquery;
    var form = layui.form;
    var element = layui.element;
    var admin = layui.admin;
    var func = layui.func;

    /* 选择头像 */
    $('#userInfoHead').click(function () {
        layer.msg("头像裁剪完善中");
        return false;
        // admin.cropImg({
        //     imgSrc: $('#userInfoHead>img').attr('src'),
        //     onCrop: function (res) {
        //         $('#userInfoHead>img').attr('src', res);
        //         parent.layui.jquery('.layui-layout-admin>.layui-header .layui-nav img.layui-nav-img').attr('src', res);
        //     }
        // });
    });

    /* 监听表单提交 */
    form.on('submit(userInfoSubmit)', function (data) {
        console.log(data.field)
        // 切记采用FormData表单提交
        var formData = new FormData();
        // formData.append("avatar", data.field.avatar);
        formData.append("realname", data.field.realname);
        formData.append("nickname", data.field.nickname);
        formData.append("gender", data.field.gender);
        formData.append("mobile", data.field.mobile);
        formData.append("email", data.field.email);
        formData.append("address", data.field.address);
        console.log(formData)
        func.ajaxPost("/userInfo", formData);
        return false;
    });

});
