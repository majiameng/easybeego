    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 城市管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
              {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true}
            , {field: 'Name', width: 200, title: '城市名称', align: 'left'}
            , {field: 'Level', width: 100, title: '城市级别', align: 'center', templet(d) {
                var cls = "";
                var levelStr = ""
                if (d.Level == 1) {
                    // 省份
                    cls = "layui-btn-normal";
                    levelStr = "省份"
                } else if (d.Level == 2) {
                    // 市区
                    cls = "layui-btn-danger";
                    levelStr = "市区"
                } else if (d.Level == 3) {
                    // 区县
                    cls = "layui-btn-warm";
                    levelStr = "区县"
                }
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+levelStr+'</span>';
            }}
            , {field: 'Citycode', width: 150, title: '城市编号（区号）', align: 'center'}
            , {field: 'PAdcode', width: 150, title: '父级地理编号', align: 'center'}
            , {field: 'Adcode', width: 150, title: '地理编号', align: 'center'}
            , {field: 'Sort', width: 100, title: '排序号', align: 'center'}
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.CreateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.UpdateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {width: 230, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.treetable(cols, "tableList");

        //【设置弹框】
        func.setWin("城市",750, 400);

    }
});
