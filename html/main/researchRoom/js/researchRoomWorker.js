function getAllOrSpecifiedResearchRoomWorkerTable() {
    if (currentShow != "getAllOrSpecifiedResearchRoomWorkerTable") {
        var html = '<table><tbody><tr><th>办公室名称或办公室编号</th></tr><tr><td><input class="input_area_input" id="research_room_id_input" type="text" placeholder="留空则查询所有员工"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="research_room_submit_button" onclick="getAllOrSpecifiedResearchRoomWorkerSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecifiedResearchRoomWorkerTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecifiedResearchRoomWorkerSubmit() {
    var researchRoom = {};
    researchRoom.name_or_id = $('#research_room_id_input').val();
    $.ajax({
        url: '/api/getAllOrSpecifiedResearchRoomWorker',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: researchRoom,
        error: function(data) {
            alertify.error('查询失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('查询成功');
            console.log(data);
            data = data.data;
            var html = '<table><tbody><tr><th>工号</th><th>姓名</th><th>研究室名</th><th>方向</th></tr>';

            for (var i = 0; i < data.length; i++) {
                html += '<tr><td>' + data[i].worker_id + '</td><td>' + data[i].worker_name + '</td><td>' + data[i].research_room_name + '</td><td>' + data[i].direction + '</td></tr>';
            }

            html += '</tbody></table>';
            $('#output_area').empty();
            $('#output_area').append(html);

        }
    });
}

function addOrUpdateOrDeleteResearchRoomWorkerTable(){
    if (currentShow != "addOrUpdateOrDeleteResearchRoomWorkerTable"){
        var html = '<table><tr><th>工号</th><th>研究室号</th><th>职责</th></tr><tr><td><input class="input_area_input" id="worker_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="research_room_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="job_detail_input" type="text" placeholder="e.g. 负责人"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateResearchRoomWorkerSubmit()">添加或修改</button> <button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateOrDeleteResearchRoomWorkerDelete()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDeleteResearchRoomWorkerTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdateResearchRoomWorkerSubmit(){
    var researchRoomWorker = {};
    researchRoomWorker.worker_id = $('#worker_id_input').val();
    researchRoomWorker.research_room_id = $('#research_room_id_input').val();
    researchRoomWorker.direction = $('#job_detail_input').val();
    $.ajax({
        url: '/api/addOrUpdateResearchRoomWorker',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: researchRoomWorker,
        error: function(data) {
            alertify.error('添加或修改失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加或修改成功');
        }
    });
}

function addOrUpdateOrDeleteResearchRoomWorkerDelete(){
    var researchRoomWorker = {};
    researchRoomWorker.worker_id = $('#worker_id_input').val();
    researchRoomWorker.research_room_id = $('#research_room_id_input').val();

    $.ajax({
        url: '/api/deleteResearchRoomWorker',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: researchRoomWorker,
        error: function(data) {
            alertify.error('删除失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('删除成功');
        }
    });
}