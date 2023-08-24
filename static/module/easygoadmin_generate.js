    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 代码生成器
 * @author 半城风雨
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
            , {field: 'Name', width: 150, title: '表名', align: 'center'}
            , {field: 'Engine', width: 100, title: '引擎', align: 'center'}
            , {field: 'Version', width: 100, title: '版本', align: 'center'}
            , {field: 'Collation', width: 180, title: '编码', align: 'center'}
            , {field: 'Rows', width: 100, title: '记录数', align: 'center'}
            , {field: 'DataLength', width: 100, title: '大小', align: 'center'}
            , {field: 'AutoIncrement', width: 100, title: '自增索引', align: 'center'}
            , {field: 'Comment', width: 150, title: '表备注', align: 'center'}
            , {
                field: '', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '未备份';
                }
            }
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.CreateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.UpdateTime*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {fixed: 'right', width: 100, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList", function (layEvent, data) {
            console.log(data)
            if (layEvent === 'generate') {
                // 一键生成模块
                layer.confirm('您确定要生成表【' + data.Name + '】对应的模块吗？生成后将会覆盖已有的模块文件！', {
                    icon: 3,
                    skin: 'layer-ext-moon',
                    btn: ['确认', '取消'] //按钮
                }, function (index) {
                    // layer.msg("一键生成模块");
                    // 切记采用FormData表单提交
                    var formData = new FormData();
                    formData.append("name", data.Name);
                    formData.append("comment",  data.Comment);
                    // 调用内部方法
                    var url = cUrl + "/generate";
                    func.ajaxPost(url, formData, function (data, flag) {
                        // 关闭弹窗
                        layer.close(index);
                    }, '模块文件生成中。。。');
                });
            }
        });

        // 批量生成
        $("#batchGenerate").click(function () {
            // 选择数据
            var data = func.getCheckData();
            console.log(data)
            var item = [];
            var tables = []
            $.each(data, function (key, val) {
                console.log(key)
                console.log(val["Name"]+"|"+val["Comment"])
                item.push(val["Name"]+"|"+val["Comment"]);
                tables.push(val["Name"])
            })
            console.log(item.join(","))
            if (data.length == 0) {
                layer.msg("请选择数据表", {icon: 5});
                return false;
            }
            layer.confirm('您确定要生成表【' + tables.join(",") + '】对应的模块吗？生成后将会覆盖已有的模块文件！', {
                icon: 3,
                skin: 'layer-ext-moon',
                btn: ['确认', '取消'] //按钮
            }, function (index) {
                // layer.msg("一键生成模块");
                // 切记采用FormData表单提交
                var formData = new FormData();
                formData.append("tables", item.join(","));
                // 调用内部方法
                var url = cUrl + "/batchGenerate";
                func.ajaxPost(url, formData, function (data, flag) {
                    // 关闭弹窗
                    layer.close(index);
                }, '模块文件生成中。。。');
            });
        });
    }
});
