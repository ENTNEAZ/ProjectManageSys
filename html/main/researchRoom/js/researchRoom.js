function getAllResearchRoom(){
    var html = '<table id="research_room_table"><tr id="research_room_table_tr"><th>研究室编号</th><th>研究室名</th><th>研究室方向</th></tr></table>'
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
            }
        }
    });
}

function addOrUpdateResearchRoomTable() {
    if (currentShow != "addOrUpdateResearchRoomTable"){
        var html = '<table><tr><th>研究室编号</th><th>研究室名</th><th>研究室方向</th></tr><tr><td><input class="input_area_input" id="research_room_id_input" type="text" placeholder="若添加研究室此处留空"></td><td><input class="input_area_input" id="research_room_name_input" type="text" placeholder="e.g. 大语言模型研究室"></td><td><input class="input_area_input" id="research_room_direction_input" type="text" placeholder="e.g. 深入研究LLM"></td></tr></table><button class="input_area_button" id="research_room_submit_button" onclick="addOrUpdateResearchRoomSubmit()">添加或修改</button>';
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