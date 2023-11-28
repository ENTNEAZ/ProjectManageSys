function getAllOrSpecifiedProjectFruitTable() {
    if (currentShow != "getAllOrSpecifiedProjectFruitTable"){
        var html = '<table><tbody><tr><th>项目名称或项目编号</th></tr><tr><td><input class="input_area_input" id="project_idname_input" type="text" placeholder="留空则查询所有项目"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="sectary_submit_button" onclick="getAllOrSpecifiedProjectFruitSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecifiedProjectFruitTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecifiedProjectFruitSubmit(){
    var project = {};
    project.idname = $('#project_idname_input').val();
    $.ajax({
        url: '/api/getAllOrSpecifiedProjectFruit',
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
            console.log(data);
            var html = '<table><tbody><tr><th>项目编号</th><th>项目名称</th><th>成果贡献人工号</th><th>成果贡献人姓名</th><th>项目成果编号</th><th>项目成果获得时间</th><th>项目成果排名</th><th>项目成果类型</th><th>项目成果详情</th></tr>';
            for (var i = 0; i < data.data.length; i++){
                html += '<tr><td>' + data.data[i].project_id + '</td><td>' + data.data[i].project_name + '</td><td>' + data.data[i].worker_id + '</td><td>' + data.data[i].worker_name + '</td><td>' + data.data[i].project_fruit_id + '</td><td>' + data.data[i].project_fruit_get_time + '</td><td>' + data.data[i].project_fruit_master_rank + '</td><td>' + data.data[i].project_fruit_type + '</td><td>' + data.data[i].project_fruit_detail + '</td></tr>';
            }

            html += '</tbody></table>';
            $('#output_area').empty();
            $('#output_area').append(html);

        }
    });
}

function addOrUpdateOrDeleteProjectFruitTable(){
    if (currentShow != "addOrUpdateOrDeleteProjectFruitTable"){
        var html = '<table><tr><th>项目编号</th><th>成果贡献人工号</th><th>项目成果编号</th><th>项目成果获得时间</th><th>项目成果排名</th><th>项目成果类型</th><th>项目成果详情</th></tr><tr><td><input class="input_area_input" id="project_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="worker_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="project_fruit_id_input" type="text" placeholder="e.g. 1001(删除时仅填此项)"></td><td><input class="input_area_input" id="project_fruit_get_time_input" type="text" placeholder="e.g. 2019-01-01"></td><td><input class="input_area_input" id="project_fruit_master_rank_input" type="text" placeholder="e.g. 1"></td><td><input class="input_area_input" id="project_fruit_type_input" type="text" placeholder="e.g. 专利"></td><td><input class="input_area_input" id="project_fruit_detail_input" type="text" placeholder="e.g. 专利详情"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateProjectFruitSubmit()">添加或修改</button> <button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateOrDeleteProjectFruitDelete()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDeleteProjectFruitTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdateProjectFruitSubmit(){
    var project = {};
    project.project_id = $('#project_id_input').val();
    project.worker_id = $('#worker_id_input').val();
    project.project_fruit_id = $('#project_fruit_id_input').val();
    project.project_fruit_get_time = $('#project_fruit_get_time_input').val();
    project.project_fruit_master_rank = $('#project_fruit_master_rank_input').val();
    project.project_fruit_type = $('#project_fruit_type_input').val();
    project.project_fruit_detail = $('#project_fruit_detail_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDeleteProjectFruit',
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

function addOrUpdateOrDeleteProjectFruitDelete(){
    var project = {};
    project.project_fruit_id = $('#project_fruit_id_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDeleteProjectFruit?project_fruit_id=' + project.project_fruit_id,
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

