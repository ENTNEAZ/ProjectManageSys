function getAllWorker(){
    //在output_area生成一个表
    //jquery included
    //clear output_area
    $('#output_area').empty();
    var table = $('<table></table>');
    table.attr('id','worker_table');
    $('#output_area').append(table);
    var tr = $('<tr></tr>');
    tr.attr('id','worker_tr');
    $('#worker_table').append(tr);
    tr.append('<th>工号</th>');
    tr.append('<th>姓名</th>');
    tr.append('<th>性别</th>');
    tr.append('<th>生日</th>');
    tr.append('<th>入职日期</th>');
    tr.append('<th>职务/职称</th>');

    //get data
    $.ajax({
        url: '/api/getAllWorker',
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
                tr.attr('id','worker_tr_'+i);
                $('#worker_table').append(tr);
                tr.append('<td>'+data[i].WorkerId+'</td>');
                tr.append('<td>'+data[i].WorkerName+'</td>');
                tr.append('<td>'+data[i].WorkerGender+'</td>');
                tr.append('<td>'+data[i].WorkerBirth+'</td>');
                tr.append('<td>'+data[i].WorkerJoinTime+'</td>');
                tr.append('<td>'+data[i].WorkerJob+'</td>');
            }
        }
    });
}

function addOrUpdateWorkerTable(){
    var html = '<table><tr><th>工号</th><th>姓名</th><th>性别</th><th>生日</th><th>入职日期</th><th>职务/职称</th></tr><tr><td><input class="input_area_input" id="worker_id_input" type="text" placeholder="若添加用户此处留空"></td><td><input class="input_area_input" id="worker_name_input" type="text" placeholder="e.g. 张三"></td><td><input class="input_area_input" id="worker_gender_input" type="text" placeholder="e.g. 男"></td><td><input class="input_area_input" id="worker_birth_input" type="date" placeholder="生日"></td><td><input class="input_area_input" id="worker_join_time_input" type="date" placeholder="入职日期"></td><td><input class="input_area_input" id="worker_job_input" type="text" placeholder="e.g. 普通职员"></td></tr></table><button class="input_area_button" id="worker_submit_button" onclick="addOrUpdateWorkerSubmit()">添加或修改</button>';
    $('#input_area').empty();
    $('#input_area').append(html);
}

function addOrUpdateWorkerSubmit(){
    var worker = {};
    worker.WorkerId = $('#worker_id_input').val();
    worker.WorkerName = $('#worker_name_input').val();
    worker.WorkerGender = $('#worker_gender_input').val();
    worker.WorkerBirth = $('#worker_birth_input').val();
    worker.WorkerJoinTime = $('#worker_join_time_input').val();
    worker.WorkerJob = $('#worker_job_input').val();
    console.log(worker);

    $.ajax({
        url: '/api/addOrUpdateWorker',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        data: worker,
        error: function(data) {
            console.log(data);
            alertify.error('添加或修改失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加或修改成功');
        }
    });
}

function sleep (time) {
    return new Promise((resolve) => setTimeout(resolve, time));
  }