    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 广告位描述
 * @author Tinymeng
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'Description', width: 200, title: '广告位描述', align: 'center'}
            , {field: 'ItemName', width: 150, title: '所属站点', align: 'center'}
            , {field: 'CateName', width: 200, title: '所属栏目', align: 'center'}
            , {field: 'LocId', width: 120, title: '广告页面位置', align: 'center'}
            , {field: 'Platform', width: 100, title: '所属平台', align: 'center', templet(d) {
                    var cls = "";
                    if (d.Platform == 1) {
                        // PC网站
                        cls = "layui-btn-normal";
                    } else if (d.Platform == 2) {
                        // WAP手机站
                        cls = "layui-btn-danger";
                    } else if (d.Platform == 3) {
                        // 微信小程序
                        cls = "layui-btn-warm";
                    } else if (d.Platform == 4) {
                        // APP移动端
                        cls = "layui-btn-primary";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.PlatformName + '</span>';
                }
            }
            , {field: 'Sort', width: 100, title: '排序号', align: 'center'}
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.CreateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.UpdateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("广告位描述",750, 400);
    }
});
