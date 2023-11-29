function getAllOrSpecifiedSubProjectTable(){
    if (currentShow != "getAllOrSpecifiedSubProject"){
        var html = '<table><tbody><tr><th>项目名称或编号</th></tr><tr><td><input class="input_area_input" id="project_id_input" type="text" placeholder="留空则查询所有项目"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="sectary_submit_button" onclick="getAllOrSpecifiedSubProjectSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecifiedSubProject"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecifiedSubProjectSubmit(){
    var project = {};
    project.idname = $('#project_id_input').val();
    $.ajax({
        url: '/api/findAllSubProjectInProject',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: project,
        error: function(data) {
            alertify.error('查询失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('查询成功');
            var html = '<table><tbody><tr><th>项目编号</th><th>项目名称</th><th>子项目编号</th><th>子项目技术详情</th><th>负责人工号</th><th>负责人姓名</th><th>子项目结束时间</th><th>子项目资金</th></tr>';
            for (var i = 0; i < data.data.length; i++){
                html += '<tr><td>' + data.data[i].project_id + '</td><td>' + data.data[i].project_name + '</td><td>' + data.data[i].sub_project_id + '</td><td>' + data.data[i].sub_project_tech_detail + '</td><td>' + data.data[i].worker_id + '</td><td>' + data.data[i].worker_name + '</td><td>' + data.data[i].sub_project_end_time + '</td><td>' + data.data[i].sub_project_fund + '</td></tr>';
            }
            html += '</tbody></table>';
            $('#output_area').empty();
            $('#output_area').append(html);

        }
    });
}

function addOrUpdateOrDeleteSpecifiedSubProjectTable(){
    if (currentShow != "addOrUpdateOrDeleteSpecifiedSubProjectTable"){
        var html = '<table><tr><th>子项目编号</th><th>子项目所处项目编号</th><th>子项目技术详情</th><th>负责人工号</th><th>子项目结束时间</th><th>子项目资金</th></tr>';
        html += '<tr><td><input class="input_area_input" id="sub_project_id_input" type="text" placeholder="e.g. 1001(删除时仅填此项)"></td><td><input class="input_area_input" id="project_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="sub_project_tech_detail_input" type="text" placeholder="e.g. 项目1"></td><td><input class="input_area_input" id="worker_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="sub_project_end_time_input" type="text" placeholder="e.g. 2019-01-01"></td><td><input class="input_area_input" id="sub_project_fund_input" type="text" placeholder="e.g. 100"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateSpecifiedSubProjectSubmit()">添加或修改</button> <button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateOrDeleteSpecifiedSubProjectDelete()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDeleteSpecifiedSubProjectTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdateSpecifiedSubProjectSubmit(){
    var project = {};
    project.sub_project_id = $('#sub_project_id_input').val();
    project.project_id = $('#project_id_input').val();
    project.sub_project_tech_detail = $('#sub_project_tech_detail_input').val();
    project.worker_id = $('#worker_id_input').val();
    project.sub_project_end_time = $('#sub_project_end_time_input').val();
    project.sub_project_fund = $('#sub_project_fund_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDeleteSpecifiedSubProject',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: project,
        error: function(data) {
            alertify.error('添加或修改失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加或修改成功');
        }
    });
}

function addOrUpdateOrDeleteSpecifiedSubProjectDelete(){
    var project = {};
    project.sub_project_id = $('#sub_project_id_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDeleteSpecifiedSubProject?sub_project_id=' + project.sub_project_id,
        type: 'DELETE',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        error: function(data) {
            alertify.error('删除失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('删除成功');
        }
    });
}

function getAllOrSpecifiedSubProjectWorkerTable(){
    if (currentShow != "getAllOrSpecifiedSubProjectWorker"){
        var html = '<table><tbody><tr><th>子项目编号</th></tr><tr><td><input class="input_area_input" id="sub_project_id_input" type="text" placeholder="留空则查询所有子项目"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="sectary_submit_button" onclick="getAllOrSpecifiedSubProjectWorkerSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecifiedSubProjectWorker"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecifiedSubProjectWorkerSubmit(){
    var project = {};
    project.idname = $('#sub_project_id_input').val();
    $.ajax({
        url: '/api/findAllSubProjectInProjectForWorker',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: project,
        error: function(data) {
            alertify.error('查询失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('查询成功');
            var html = '<table><tbody><tr><th>子项目编号</th><th>子项目详细内容</th><th>子项目参与者工号</th><th>子项目参与者姓名</th><th>初次参与时间</th><th>经费</th><th>工作时长</th></tr>';
            //{"code": 0, "msg": "success", "data": [{"sub_project_id":"1","sub_project_tech_detail":"研究NFC技术","worker_id":"3","worker_name":"王五","join_time":"2023-11-28","sub_project_worker_fund":"100","workload":"2"},{"sub_project_id":"1","sub_project_tech_detail":"研究NFC技术","worker_id":"3","worker_name":"王五","join_time":"2023-11-28","sub_project_worker_fund":"100","workload":"2"}]}
            for (var i = 0; i < data.data.length; i++){
                html += '<tr><td>' + data.data[i].sub_project_id + '</td><td>' + data.data[i].sub_project_tech_detail + '</td><td>' + data.data[i].worker_id + '</td><td>' + data.data[i].worker_name + '</td><td>' + data.data[i].join_time + '</td><td>' + data.data[i].sub_project_worker_fund + '</td><td>' + data.data[i].workload + '</td></tr>';
            }
            html += '</tbody></table>';
            $('#output_area').empty();
            $('#output_area').append(html);

        }
    });
}

function addOrUpdateOrDeleteSpecifiedSubProjectWorkerTable(){
    if (currentShow != "addOrUpdateOrDeleteSpecifiedSubProjectWorkerTable"){
        var html = '<table><tr><th>子项目编号</th><th>子项目参与者工号</th><th>初次参与时间</th><th>经费</th><th>工作时长</th></tr>';
        html += '<tr><td><input class="input_area_input" id="sub_project_id_input" type="text" placeholder="e.g. 1001(删除时填此项)"></td><td><input class="input_area_input" id="worker_id_input" type="text" placeholder="e.g. 1001 (删除时填此项)"></td><td><input class="input_area_input" id="join_time_input" type="text" placeholder="e.g. 2019-01-01"></td><td><input class="input_area_input" id="sub_project_worker_fund_input" type="text" placeholder="e.g. 100"></td><td><input class="input_area_input" id="workload_input" type="text" placeholder="e.g. 100"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateSpecifiedSubProjectWorkerSubmit()">添加或修改</button> <button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateOrDeleteSpecifiedSubProjectWorkerDelete()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDeleteSpecifiedSubProjectWorkerTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdateSpecifiedSubProjectWorkerSubmit(){
    var project = {};
    project.sub_project_id = $('#sub_project_id_input').val();
    project.worker_id = $('#worker_id_input').val();
    project.join_time = $('#join_time_input').val();
    project.sub_project_worker_fund = $('#sub_project_worker_fund_input').val();
    project.workload = $('#workload_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDeleteSpecifiedSubProjectWorker',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: project,
        error: function(data) {
            alertify.error('添加或修改失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加或修改成功');
        }
    });
}

function addOrUpdateOrDeleteSpecifiedSubProjectWorkerDelete(){
    var project = {};
    project.sub_project_id = $('#sub_project_id_input').val();
    project.worker_id = $('#worker_id_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDeleteSpecifiedSubProjectWorker?sub_project_id=' + project.sub_project_id + '&worker_id=' + project.worker_id,
        type: 'DELETE',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        error: function(data) {
            alertify.error('删除失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('删除成功');
        }
    });
}