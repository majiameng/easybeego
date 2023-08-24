    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 演示二管理
 * @author 半城风雨
 * @since 2022-05-13
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
        
             
			, {field: 'Name', width: 100, title: '演示名称', align: 'center'}
            
        
            
            , {field: 'Status', width: 100, title: '状态', align: 'center', templet: function (d) {
				return  '<input type="checkbox" name="Status" value="'+d.Id+'" lay-skin="switch" lay-text="正常|停用" lay-filter="Status" '+(d.Status==1 ? 'checked' : '')+'>';
            }}
            
        
             
			, {field: 'Sort', width: 100, title: '排序号', align: 'center'}
            
        
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("演示二", 750, 450);

    
		
	
		
		//【设置状态】
        func.formSwitch('Status', null, function (data, res) {
            console.log("开关回调成功");
        });
		
	
		
	
    }
});
