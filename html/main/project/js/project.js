function getAllOrSpecifiedProjectTable(){
    if (currentShow != "getAllOrSpecifiedProjectTable"){
        var html = '<table><tbody><tr><th>项目名称或项目编号</th></tr><tr><td><input class="input_area_input" id="project_idname_input" type="text" placeholder="留空则查询所有项目"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="sectary_submit_button" onclick="getAllOrSpecifiedProjectSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecifiedProjectTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecifiedProjectSubmit(){
    var project = {};
    project.project_idname = $('#project_idname_input').val();
    $.ajax({
        url: '/api/getAllOrSpecifiedProject',
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
        var html = '<table><tbody><tr><th>项目编号</th><th>项目名称</th><th>项目详情</th><th>项目开始时间</th><th>项目结束时间</th><th>项目资金</th><th>项目负责人</th><th>项目委托方编号</th><th>项目委托方名</th><th>项目监管方编号</th><th>项目监管方名</th></tr>';
        for (var i = 0; i < data.data.length; i++){
            html += '<tr><td>' + data.data[i].project_id + '</td><td>' + data.data[i].project_name + '</td><td>' + data.data[i].project_detail + '</td><td>' + data.data[i].project_start_time + '</td><td>' + data.data[i].project_end_time + '</td><td>' + data.data[i].project_fund + '</td><td>' + data.data[i].worker_name + '</td><td>' + data.data[i].project_participant_id + '</td><td>' + data.data[i].project_participant_name + '</td><td>' + data.data[i].project_supervisor_id + '</td><td>'+ data.data[i].project_supervisor_name + '</td></tr>';
        }
        html += '</tbody></table>';
        $('#output_area').empty();
        $('#output_area').append(html);
        }
    });
}

function addOrUpdateOrDeleteProjectTable(){
    if (currentShow != "addOrUpdateOrDeleteProjectTable"){
        var html = '<table><tr><th>项目编号</th><th>项目名称</th><th>项目详情</th><th>项目开始时间</th><th>项目结束时间</th><th>项目资金</th><th>项目负责人工号</th><th>项目委托方编号</th><th>项目监管方编号</th></tr><tr><td><input class="input_area_input" id="project_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="project_name_input" type="text" placeholder="e.g. 项目1"></td><td><input class="input_area_input" id="project_detail_input" type="text" placeholder="e.g. 项目1详情"></td><td><input class="input_area_input" id="project_start_time_input" type="text" placeholder="e.g. 2019-01-01"></td><td><input class="input_area_input" id="project_end_time_input" type="text" placeholder="e.g. 2019-01-01"></td><td><input class="input_area_input" id="project_fund_input" type="text" placeholder="e.g. 100"></td><td><input class="input_area_input" id="worker_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="project_participant_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="project_supervisor_id_input" type="text" placeholder="e.g. 1001"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateProjectSubmit()">添加或修改</button> <button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateOrDeleteProjectDelete()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDeleteProjectTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdateProjectSubmit(){
    var project = {};
    project.project_id = $('#project_id_input').val();
    project.project_name = $('#project_name_input').val();
    project.project_detail = $('#project_detail_input').val();
    project.project_start_time = $('#project_start_time_input').val();
    project.project_end_time = $('#project_end_time_input').val();
    project.project_fund = $('#project_fund_input').val();
    project.worker_id = $('#worker_id_input').val();
    project.project_participant_id = $('#project_participant_id_input').val();
    project.project_supervisor_id = $('#project_supervisor_id_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDeleteProject',
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

function addOrUpdateOrDeleteProjectDelete(){
    var project_id = $('#project_id_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDeleteProject?project_id=' + project_id,
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

function getAllOrSpecifiedProjectWorkerTable(){
    if (currentShow != "getAllOrSpecifiedProjectWorkerTable"){
        var html = '<table><tbody><tr><th>项目名称或编号</th></tr><tr><td><input class="input_area_input" id="project_id_input" type="text" placeholder="留空则查询所有项目"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="sectary_submit_button" onclick="getAllOrSpecifiedProjectWorkerSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecifiedProjectWorkerTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecifiedProjectWorkerSubmit(){
    var project = {};
    project.idname = $('#project_id_input').val();
    $.ajax({
        url: '/api/findAllWorkerInProject',
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
        var html = '<table><tbody><tr><th>项目编号</th><th>项目名称</th><th>科研人员工号</th><th>科研人员姓名</th></tr>';
        for (var i = 0; i < data.data.length; i++){
            html += '<tr><td>' + data.data[i].project_id + '</td><td>' + data.data[i].project_name + '</td><td>' + data.data[i].worker_id + '</td><td>' + data.data[i].worker_name + '</td></tr>';
        }
        html += '</tbody></table>';
        $('#output_area').empty();
        $('#output_area').append(html);
        }
    });
}

function addOrUpdateOrDeleteProjectWorkerTable(){
    if (currentShow != "addOrUpdateOrDeleteProjectWorkerTable"){
        var html = '<table><tr><th>项目编号</th><th>科研人员工号</th></tr><tr><td><input class="input_area_input" id="project_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="worker_id_input" type="text" placeholder="e.g. 1001"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateProjectWorkerSubmit()">添加</button> <button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateOrDeleteProjectWorkerDelete()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDeleteProjectWorkerTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdateProjectWorkerSubmit(){
    var project = {};
    project.project_id = $('#project_id_input').val();
    project.worker_id = $('#worker_id_input').val();
    $.ajax({
        url: '/api/addOrDeleteProjectWorker',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: project,
        error: function(data) {
            alertify.error('添加失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加成功');
        }
    });
}

function addOrUpdateOrDeleteProjectWorkerDelete(){
    var project = {};
    project.project_id = $('#project_id_input').val();
    project.worker_id = $('#worker_id_input').val();
    $.ajax({
        url: '/api/addOrDeleteProjectWorker?project_id=' + project.project_id + '&worker_id=' + project.worker_id,
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

function getAllOrSpecifiedProjectHelperTable(){
    if (currentShow != "getAllOrSpecifiedProjectParticipantTable"){
        var html = '<table><tbody><tr><th>项目名称或编号</th></tr><tr><td><input class="input_area_input" id="project_id_input" type="text" placeholder="留空则查询所有项目"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="sectary_submit_button" onclick="getAllOrSpecifiedProjectParticipantSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecifiedProjectParticipantTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecifiedProjectParticipantSubmit(){
    var project = {};
    project.idname = $('#project_id_input').val();
    $.ajax({
        url: '/api/findAllParticipantInProject',
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
        var html = '<table><tbody><tr><th>项目编号</th><th>项目名称</th><th>委托方编号</th><th>委托方名称</th></tr>';
        for (var i = 0; i < data.data.length; i++){
            html += '<tr><td>' + data.data[i].project_id + '</td><td>' + data.data[i].project_name + '</td><td>' + data.data[i].project_participant_id + '</td><td>' + data.data[i].project_participant_name + '</td></tr>';
        }
        html += '</tbody></table>';
        $('#output_area').empty();
        $('#output_area').append(html);
        }
    });
}

function addOrUpdateOrDeleteProjectHelperTable(){
    if (currentShow != "addOrUpdateOrDeleteProjectParticipantTable"){
        var html = '<table><tr><th>项目编号</th><th>委托方编号</th></tr><tr><td><input class="input_area_input" id="project_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="project_participant_id_input" type="text" placeholder="e.g. 1001"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateProjectParticipantSubmit()">添加</button> <button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateOrDeleteProjectParticipantDelete()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDeleteProjectParticipantTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdateProjectParticipantSubmit(){
    var project = {};
    project.project_id = $('#project_id_input').val();
    project.project_participant_id = $('#project_participant_id_input').val();
    $.ajax({
        url: '/api/addOrDeleteProjectParticipant',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: project,
        error: function(data) {
            alertify.error('添加失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加成功');
        }
    });
}

function addOrUpdateOrDeleteProjectParticipantDelete(){
    var project = {};
    project.project_id = $('#project_id_input').val();
    project.project_participant_id = $('#project_participant_id_input').val();
    $.ajax({
        url: '/api/addOrDeleteProjectParticipant?project_id=' + project.project_id + '&project_participant_id=' + project.project_participant_id,
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
