function getAllResearchRoom(){
    var html = '<table id="research_room_table"><tr id="research_room_table_tr"><th>研究室编号</th><th>研究室名</th><th>研究室方向</th><th>主任工号</th><th>主任姓名</th><th>主任任期</th><th>主任任职时间</th></tr></table>'
    $('#output_area').empty();
    $('#output_area').append(html);

    // get data
    $.ajax({
        url: '/api/getAllResearchRoom',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        error: function() {
            alertify.error('获取失败');
        },
        success: function(data) {
            alertify.success('获取成功');
            console.log(data);
            data = data.data;
            for(var i = 0;i < data.length ;i++){
                var tr = $('<tr></tr>');
                tr.attr('id','research_room_table_tr_'+i);
                $('#research_room_table').append(tr);
                tr.append('<td>'+data[i].ResearchRoomID+'</td>');
                tr.append('<td>'+data[i].ResearchRoomName+'</td>');
                tr.append('<td>'+data[i].ResearchRoomDirection+'</td>');
                tr.append('<td>'+data[i].Worker_id+'</td>');
                tr.append('<td>'+data[i].Worker_name+'</td>');
                tr.append('<td>'+data[i].Term+'</td>');
                tr.append('<td>'+data[i].Join_date+'</td>');
            }
        }
    });
}

function addOrUpdateResearchRoomTable() {
    if (currentShow != "addOrUpdateResearchRoomTable"){
        var html = '<table><tr><th>研究室编号</th><th>研究室名</th><th>研究室方向</th><th>主任工号</th><th>主任任期</th><th>主任任职时间</th>';
        html += '<tr><td><input class="input_area_input" id="research_room_id_input" type="text" placeholder="e.g. 1"></td><td><input class="input_area_input" id="research_room_name_input" type="text" placeholder="e.g. 第一研究室"></td><td><input class="input_area_input" id="research_room_direction_input" type="text" placeholder="e.g. 研究高精尖设备"></td><td><input class="input_area_input" id="worker_id_input" type="text" placeholder="e.g. 1"></td><td><input class="input_area_input" id="term_input" type="text" placeholder="e.g. 10"></td><td><input class="input_area_input" id="join_date_input" type="text" placeholder="e.g. 2020-01-01"></td></tr></table><button class="input_area_button" id="research_room_submit_button" onclick="addOrUpdateResearchRoomSubmit()">添加或修改</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateResearchRoomTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }

}

function addOrUpdateResearchRoomSubmit(){
    var researchRoom = {};
    researchRoom.id = $('#research_room_id_input').val();
    researchRoom.name = $('#research_room_name_input').val();
    researchRoom.direction = $('#research_room_direction_input').val();
    researchRoom.worker_id = $('#worker_id_input').val();
    researchRoom.term = $('#term_input').val();
    researchRoom.join_date = $('#join_date_input').val();
    $.ajax({
        url: '/api/addOrUpdateResearchRoom',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: researchRoom,
        error: function(data) {
            alertify.error('添加或修改失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加或修改成功');
            console.log(data);
        }
    });
}

function deleteResearchRoomTable(){
    if (currentShow != "deleteResearchRoomTable"){
        var html = '<table><tr><th>研究室编号</th></tr><tr><td><input class="input_area_input" id="research_room_id_input" type="text" placeholder="e.g. 1"></td></tr></table><button class="input_area_button" id="research_room_submit_button" onclick="deleteResearchRoomSubmit()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "deleteResearchRoomTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function deleteResearchRoomSubmit(){
    var researchRoom = {};
    researchRoom.id = $('#research_room_id_input').val();
    $.ajax({
        url: '/api/deleteResearchRoom',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: researchRoom,
        error: function(data) {
            alertify.error('删除失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('删除成功');
            console.log(data);
        }
    });
}