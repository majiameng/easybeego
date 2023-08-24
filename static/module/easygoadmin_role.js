    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 角色管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func', 'admin', 'zTree'], function () {

    //声明变量
    var func = layui.func
        , admin = layui.admin
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'Name', width: 200, title: '角色名称', align: 'center'}
            , {field: 'Status', width: 100, title: '状态', align: 'center', templet: function (d) {
                return  '<input type="checkbox" name="status" value="' + d.Id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" '+(d.Status==1 ? 'checked' : '')+'>';
            }}
            , {field: 'Sort', width: 100, title: '排序', align: 'center'}
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.CreateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.UpdateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {fixed: 'right', width: 250, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList", function (layEvent, data) {
            if (layEvent === 'permission') {
                admin.open({
                    title: '角色权限分配',
                    btn: ['保存', '取消'],
                    content: '<ul id="roleAuthTree" class="ztree"></ul>',
                    success: function (layero, dIndex) {
                        var loadIndex = layer.load(2);
                        $.get('/rolemenu/index?roleId=' + data.Id, function (res) {
                            layer.close(loadIndex);
                            if (0 === res.code) {
                                $.fn.zTree.init($('#roleAuthTree'), {
                                    check: {enable: true},
                                    data: {simpleData: {enable: true}}
                                }, res.data);
                            } else {
                                layer.msg(res.msg, {icon: 2});
                            }
                        }, 'json');
                        // 超出一定高度滚动
                        $(layero).children('.layui-layer-content').css({'max-height': '300px', 'overflow': 'auto'});
                    },
                    yes: function (dIndex) {
                        var insTree = $.fn.zTree.getZTreeObj('roleAuthTree');
                        var checkedRows = insTree.getCheckedNodes(true);
                        var ids = [];
                        for (var i = 0; i < checkedRows.length; i++) {
                            ids.push(checkedRows[i].id);
                        }
                        // 切记采用FormData表单提交
                        var formData = new FormData();
                        formData.append("roleId", data.Id);
                        formData.append("menuIds", ids.join(','));
                        func.ajaxPost("/rolemenu/save", formData, function (res, success) {
                            // 关闭窗体
                            layer.close(dIndex);
                        });
                    }
                });
            }
        });

        //【设置弹框】
        func.setWin("角色", 500, 350);

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
