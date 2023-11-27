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