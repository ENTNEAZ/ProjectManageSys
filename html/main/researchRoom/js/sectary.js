function getAllOrSpecifiedSectaryTable(){
    if (currentShow != "getAllOrSpecifiedSectaryTable"){
        var html = '<table><tbody><tr><th>办公室名称或办公室编号</th></tr><tr><td><input class="input_area_input" id="research_room_id_input" type="text" placeholder="留空则查询所有秘书"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="sectary_submit_button" onclick="getAllOrSpecifiedSectarySubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecifiedSectaryTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecifiedSectarySubmit(){
    var html = '<table id="sectary"><tr><th>工号</th><th>姓名</th><th>所在研究室名</th><th>职责</th></tr></table>';
    $('#output_area').empty();
    $('#output_area').append(html);
    var sectary = {};
    sectary.id_or_name = $('#research_room_id_input').val();
    $.ajax({
        url: '/api/getAllOrSpecifiedSectary',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: sectary,
        error: function(data) {
            alertify.error('获取失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('获取成功');
            console.log(data);
            data = data.data;
            for(var i = 0;i < data.length ;i++){
                var tr = $('<tr></tr>');
                tr.attr('id','sectary_tr_'+i);
                $('#sectary').append(tr);
                tr.append('<td>'+data[i].WorkerID+'</td>');
                tr.append('<td>'+data[i].WorkerName+'</td>');
                tr.append('<td>'+data[i].ResearchRoomName+'</td>');
                tr.append('<td>'+data[i].JobDetail+'</td>');
            }
        }
    });
}

function addOrUpdateOrDeleteSectaryTable(){
    // 工号 研究室号 职责
    if (currentShow != "addOrUpdateOrDeleteSectaryTable"){
        var html = '<table><tr><th>工号</th><th>研究室号</th><th>职责</th></tr><tr><td><input class="input_area_input" id="worker_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="research_room_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="job_detail_input" type="text" placeholder="e.g. 负责人"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateSectarySubmit()">添加或修改</button> <button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateOrDeleteSectaryDelete()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDeleteSectaryTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdateSectarySubmit(){
    var sectary = {};
    sectary.worker_id = $('#worker_id_input').val();
    sectary.research_room_id = $('#research_room_id_input').val();
    sectary.job_detail = $('#job_detail_input').val();
    $.ajax({
        url: '/api/addOrUpdateSectary',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: sectary,
        error: function(data) {
            alertify.error('添加或修改失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加或修改成功');
        }
    });
}

function addOrUpdateOrDeleteSectaryDelete(){
    var sectary = {};
    sectary.worker_id = $('#worker_id_input').val();
    sectary.research_room_id = $('#research_room_id_input').val();
    sectary.job_detail = $('#job_detail_input').val();
    $.ajax({
        url: '/api/deleteSectary',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: sectary,
        error: function(data) {
            alertify.error('删除失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('删除成功');
        }
    });
}