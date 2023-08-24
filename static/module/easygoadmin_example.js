    /**
     * @author: Tinymeng <666@majiameng.com>
     */

/**
 * 演示一管理
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
        
             
			, {field: 'Name', width: 100, title: '测试名称', align: 'center'}
            
        
            
            , {field: 'Avatar', width: 90, title: '头像', align: 'center', templet: function (d) {
                    var AvatarStr = "";
                    if (d.Avatar) {
                        AvatarStr = '<a href="' + d.Avatar + '" target="_blank"><img src="' + d.Avatar + '" height="26" /></a>';
                    }
                    return AvatarStr;
                }
            }
			
        
             
			, {field: 'Content', width: 100, title: '内容', align: 'center'}
            
        
            
            , {field: 'Status', width: 100, title: '状态', align: 'center', templet: function (d) {
				return  '<input type="checkbox" name="Status" value="'+d.Id+'" lay-skin="switch" lay-text="正常|停用" lay-filter="Status" '+(d.Status==1 ? 'checked' : '')+'>';
            }}
            
        
            
			, {field: 'Type', width: 100, title: '类型', align: 'center', templet(d) {
					
					if (d.Type == 1) {
						// 京东
						return '<span class="layui-btn layui-btn-normal layui-btn-xs">京东</span>';
					} 
					
					else if (d.Type == 2) {
						// 淘宝
						return '<span class="layui-btn layui-btn-danger layui-btn-xs">淘宝</span>';
					} 
					
					else if (d.Type == 3) {
						// 拼多多
						return '<span class="layui-btn layui-btn-warm layui-btn-xs">拼多多</span>';
					} 
					
					else if (d.Type == 4) {
						// 唯品会
						return '<span class="layui-btn layui-btn-primary layui-btn-xs">唯品会</span>';
					} 
					
				}
			}
			
        
            
            , {field: 'IsVip', width: 100, title: '是否VIP', align: 'center', templet: function (d) {
				return  '<input type="checkbox" name="IsVip" value="'+d.Id+'" lay-skin="switch" lay-text="是|否" lay-filter="IsVip" '+(d.IsVip==1 ? 'checked' : '')+'>';
            }}
            
        
             
			, {field: 'Sort', width: 100, title: '排序号', align: 'center'}
            
        
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("演示一", 750, 450);

    
		
	
		
	
		
	
		
		//【设置状态】
        func.formSwitch('Status', null, function (data, res) {
            console.log("开关回调成功");
        });
		
	
		
	
		
		//【设置是否VIP】
        func.formSwitch('IsVip', null, function (data, res) {
            console.log("开关回调成功");
        });
		
	
		
	
    }
});
