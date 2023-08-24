    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 会员等级
 * @author Tinymeng
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //【声明变量】
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'Name', width: 200, title: '分组名称', align: 'center'}
            , {field: 'Sort', width: 120, title: '排序', align: 'center'}
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.CreateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.UpdateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("会员等级", 450, 250);
    }
});
