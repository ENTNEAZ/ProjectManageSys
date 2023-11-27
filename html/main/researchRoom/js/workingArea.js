function getAllOrSpecifiedWorkingAreaTable() {
    if (currentShow != "getAllOrSpecifiedWorkingAreaTable"){
        var html = '<table><tbody><tr><th>办公室名称或办公室编号</th></tr><tr><td><input class="input_area_input" id="research_room_id_input" type="text" placeholder="留空则查询所有工作地点"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="research_room_submit_button" onclick="getAllOrSpecifiedWorkingAreaSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecifiedWorkingAreaTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecifiedWorkingAreaSubmit() {
    var workingArea = {};
    var html = '<table id="working_area"><tbody><tr><th>研究室地点编号</th><th>此地办公研究室名称</th><th>研究室地点大小</th><th>研究室地点位置</th></tr></table>';
    $('#output_area').empty();
    $('#output_area').append(html); 
    workingArea.id_or_name = $('#research_room_id_input').val();
    $.ajax({
        url: '/api/getAllOrSpecifiedWorkingArea',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: workingArea,
        error: function(data) {
            alertify.error('查询失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('查询成功');
            console.log(data);
            data = data.data;
            for(var i = 0;i < data.length ;i++){
                var tr = $('<tr></tr>');
                tr.attr('id','working_area_tr_'+i);
                $('#working_area').append(tr);
                tr.append('<td>'+data[i].WorkingAreaID+'</td>');
                tr.append('<td>'+data[i].ResearchRoomName+'</td>');
                tr.append('<td>'+data[i].WorkingAreaSize+'</td>');
                tr.append('<td>'+data[i].WorkingAreaAddress+'</td>');
            }
        }
    });
}

function addOrUpdateWorkingAreaTable() {
    if (currentShow != "addOrUpdateWorkingAreaTable"){
        var html = '<table><tr><th>研究室地点编号</th><th>研究室地点大小</th><th>研究室地点位置</th></tr><tr><td><input class="input_area_input" id="working_area_id_input" type="text" placeholder="若添加研究室地点此处留空"></td></td><td><input class="input_area_input" id="working_area_size_input" type="text" placeholder="e.g. 100"></td><td><input class="input_area_input" id="working_area_address_input" type="text" placeholder="e.g. A1101"></td></tr></table><button class="input_area_button" id="working_area_submit_button" onclick="addOrUpdateWorkingAreaSubmit()">添加或修改</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateWorkingAreaTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdateWorkingAreaSubmit() {
    var workingArea = {};
    workingArea.id = $('#working_area_id_input').val();
    workingArea.size = $('#working_area_size_input').val();
    workingArea.address = $('#working_area_address_input').val();
    $.ajax({
        url: '/api/addOrUpdateWorkingArea',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: workingArea,
        error: function(data) {
            alertify.error('添加或修改失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加或修改成功');
        }
    });
}

function addOrDeleteWorkingAreaForResearchRoomTable() {
    if (currentShow != "addOrDeleteWorkingAreaForResearchRoomTable"){
        var html = '<table><tr><th>研究室地点编号</th><th>研究室编号</th></tr><tr><td><input class="input_area_input" id="working_area_id_input" type="text" placeholder="e.g. 1"></td><td><input class="input_area_input" id="research_room_id_input" type="text" placeholder="e.g. 1"></td></tr></table><button class="input_area_button" id="working_area_submit_button" onclick="addOrDeleteWorkingAreaForResearchRoomAddSubmit()">添加</button><button class="input_area_button" id="working_area_submit_button" onclick="addOrDeleteWorkingAreaForResearchRoomDeleteSubmit()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrDeleteWorkingAreaForResearchRoomTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrDeleteWorkingAreaForResearchRoomAddSubmit() {
    var workingArea = {};
    workingArea.working_area_id = $('#working_area_id_input').val();
    workingArea.research_room_id = $('#research_room_id_input').val();
    $.ajax({
        url: '/api/addWorkingAreaForResearchRoom',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: workingArea,
        error: function(data) {
            alertify.error('添加失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加成功');
        }
    });
}

function addOrDeleteWorkingAreaForResearchRoomDeleteSubmit(){
    var workingArea = {};
    workingArea.working_area_id = $('#working_area_id_input').val();
    workingArea.research_room_id = $('#research_room_id_input').val();
    $.ajax({
        url: '/api/deleteWorkingAreaForResearchRoom',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: workingArea,
        error: function(data) {
            alertify.error('删除失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('删除成功');
        }
    });
}