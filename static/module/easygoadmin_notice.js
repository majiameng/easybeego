    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 通知公告
 * @author Tinymeng
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
              {type: 'checkbox', fixed: 'left'}
            , {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'Title', width: 350, title: '通知标题', align: 'center'}
            , {field: 'Source', width: 100, title: '通知来源', align: 'center', templet(d) {
                var cls = "";
                if (d.Source == 1) {
                    cls = "layui-btn-normal";
                } else if (d.Source == 2) {
                    cls = "layui-btn-primary"
                }
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.SourceName+'</span>';
            }}
            , {field: 'IsTop', width: 100, title: '是否置顶', align: 'center', templet(d) {
                if (d.IsTop == 1) {
                    // 已置顶
                    return '<span class="layui-btn layui-btn-primary layui-btn-xs">已置顶</span>';
                } else {
                    // 未置顶
                    return '<span class="layui-btn layui-btn-danger layui-btn-xs">未置顶</span>';
                }
            }}
            , {field: 'Status', width: 100, title: '发布状态', align: 'center', templet(d) {
                if (d.Status == 1) {
                    // 草稿箱
                    return '<span class="layui-btn layui-btn-normal layui-btn-xs">已发布</span>';
                } else if (d.Status == 2) {
                    // 立即发布
                    return '<span class="layui-btn layui-btn-danger layui-btn-xs">待发布</span>';
                }
            }}
            , {field: 'Browse', width: 100, title: '阅读量', align: 'center'}
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.CreateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.UpdateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("通知公告");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
