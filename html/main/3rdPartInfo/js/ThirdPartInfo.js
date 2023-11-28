function getAllOrSpecified3rdPartInfoTable(){
    if(currentShow != "getAllOrSpecified3rdPartInfoTable"){
        var html = '<table><tbody><tr><th>第三方编号或名称</th></tr><tr><td><input class="input_area_input" id="3rd_part_idname_input" type="text" placeholder="留空则查询所有第三方信息"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="sectary_submit_button" onclick="getAllOrSpecified3rdPartInfoSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecified3rdPartInfoTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}   

function getAllOrSpecified3rdPartInfoSubmit(){
    var thirdPartInfo = {};
    thirdPartInfo.idname = $('#3rd_part_idname_input').val();
    $.ajax({
        url: '/api/getAllOrSpecified3rdPartInfo',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: thirdPartInfo,
        error: function(data) {
            alertify.error('查询失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('查询成功');
            data = data.data;
            console.log(data);
            var html = '<table><tbody><tr><th>第三方编号</th><th>第三方名称</th><th>第三方地址</th><th>第三方负责人编号</th><th>第三方负责人电话</th><th>第三方负责人手机</th><th>第三方负责人邮箱</th></tr>';
            for(var i = 0;i < data.length ;i++){
                var temp = '<tr><td>' + data[i].project_participant_id + '</td><td>' + data[i].project_participant_name + '</td><td>' + data[i].project_participant_address + '</td><td>' + data[i].project_participant_worker_id + '</td><td>' + data[i].project_participant_worker_telephone + '</td><td>' + data[i].project_participant_worker_mobile + '</td><td>' + data[i].project_participant_worker_email + '</td></tr>';
                html += temp;
            }

            html += '</tbody></table>';
            $('#output_area').empty();
            $('#output_area').append(html);

        }
    });
}

function addOrUpdateOrDelete3rdPartInfoTable(){
    // id name address worker_id
    if (currentShow != "addOrUpdateOrDelete3rdPartInfoTable"){
        var html = '<table><tr><th>第三方编号</th><th>第三方名称</th><th>第三方地址</th><th>第三方负责人编号</th></tr><tr><td><input class="input_area_input" id="3rd_part_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="3rd_part_name_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="3rd_part_address_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="3rd_part_worker_id_input" type="text" placeholder="e.g. 1001"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdate3rdPartInfoSubmit()">添加或修改</button> <button class="input_area_button" id="sectary_submit_button" onclick="addOrUpdateOrDelete3rdPartInfoDelete()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDelete3rdPartInfoTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}


function addOrUpdate3rdPartInfoSubmit(){
    var thirdPartInfo = {};
    thirdPartInfo.project_participant_id = $('#3rd_part_id_input').val();
    thirdPartInfo.project_participant_name = $('#3rd_part_name_input').val();
    thirdPartInfo.project_participant_address = $('#3rd_part_address_input').val();
    thirdPartInfo.project_participant_worker_id = $('#3rd_part_worker_id_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDelete3rdPartInfo',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: thirdPartInfo,
        error: function(data) {
            alertify.error('添加或修改失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加或修改成功');
        }
    });
}

function addOrUpdateOrDelete3rdPartInfoDelete(){
    var project_participant_id = $('#3rd_part_id_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDelete3rdPartInfo?project_participant_id=' + project_participant_id,
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

function getAllOrSpecified3rdPartContactTable() {
    if(currentShow != "getAllOrSpecified3rdPartContactTable"){
        var html = '<table><tbody><tr><th>第三方编号或名称</th></tr><tr><td><input class="input_area_input" id="3rd_part_idname_input" type="text" placeholder="留空则查询所有第三方联系人信息"></td></tr></tbody></table>';
        html += '<button class="input_area_button" id="sectary_submit_button" onclick="getAllOrSpecified3rdPartContactSubmit()">查询</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "getAllOrSpecified3rdPartContactTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function getAllOrSpecified3rdPartContactSubmit(){
    var thirdPartContact = {};
    thirdPartContact.idname = $('#3rd_part_idname_input').val();
    $.ajax({
        url: '/api/getAllOrSpecified3rdPartContact',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: thirdPartContact,
        error: function(data) {
            alertify.error('查询失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('查询成功');
            data = data.data;
            console.log(data);
            var html = '<table><tbody><tr><th>第三方编号</th><th>第三方名称</th><th>在第三方担任的职务</th><th>第三方负责人/联系人编号</th><th>第三方负责人/联系人电话</th><th>第三方联系人手机</th><th>第三方负责人/联系人邮箱</th></tr>';
            for(var i = 0;i < data.length ;i++){
                var temp = '<tr><td>' + data[i].project_participant_id;
                temp += '</td><td>' + data[i].project_participant_name;
                temp += '</td><td>' + data[i].project_participant_worker_work;
                temp += '</td><td>' + data[i].project_participant_worker_id;
                temp += '</td><td>' + data[i].project_participant_worker_telephone;
                temp += '</td><td>' + data[i].project_participant_worker_mobile;
                temp += '</td><td>' + data[i].project_participant_worker_email;
                temp += '</td></tr>';
                html += temp;
            }

            html += '</tbody></table>';
            $('#output_area').empty();
            $('#output_area').append(html);
            
        }   
    });
}


function addOrUpdateOrDelete3rdPartContactrTable(){
    if (currentShow != "addOrUpdateOrDelete3rdPartContactrTable"){
        var html = "<table><tr><th>第三方负责人/联系人编号</th><th>第三方负责人/联系人电话</th><th>第三方联系人手机</th><th>第三方负责人/联系人邮箱</th></tr><tr><td><input class='input_area_input' id='3rd_part_worker_id_input' type='text' placeholder='e.g. 1001'></td><td><input class='input_area_input' id='3rd_part_worker_telephone_input' type='text' placeholder='e.g. 1001'></td><td><input class='input_area_input' id='3rd_part_worker_mobile_input' type='text' placeholder='e.g. 1001'></td><td><input class='input_area_input' id='3rd_part_worker_email_input' type='text' placeholder='e.g. 1001'></td></tr></table><button class='input_area_button' id='sectary_submit_button' onclick='addOrUpdate3rdPartContactSubmit()'>添加或修改</button> <button class='input_area_button' id='sectary_submit_button' onclick='addOrUpdateOrDelete3rdPartContactDelete()'>删除</button>";
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrUpdateOrDelete3rdPartContactrTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addOrUpdate3rdPartContactSubmit(){
    var thirdPartContact = {};
    thirdPartContact.project_participant_worker_id = $('#3rd_part_worker_id_input').val();
    thirdPartContact.project_participant_worker_telephone = $('#3rd_part_worker_telephone_input').val();
    thirdPartContact.project_participant_worker_mobile = $('#3rd_part_worker_mobile_input').val();
    thirdPartContact.project_participant_worker_email = $('#3rd_part_worker_email_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDelete3rdPartContact',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: thirdPartContact,
        error: function(data) {
            alertify.error('添加或修改失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加或修改成功');
        }
    });
}

function addOrUpdateOrDelete3rdPartContactDelete(){
    var project_participant_worker_id = $('#3rd_part_worker_id_input').val();
    $.ajax({
        url: '/api/addOrUpdateOrDelete3rdPartContact?project_participant_worker_id=' + project_participant_worker_id,
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

function addOrDeleteContactRelationTable(){
    if (currentShow != "addOrDeleteContactRelationTable"){
        var html = '<table><tr><th>第三方编号</th><th>第三方负责人/联系人编号</th></tr><tr><td><input class="input_area_input" id="3rd_part_id_input" type="text" placeholder="e.g. 1001"></td><td><input class="input_area_input" id="3rd_part_worker_id_input" type="text" placeholder="e.g. 1001"></td></tr></table><button class="input_area_button" id="sectary_submit_button" onclick="addContactRelationSubmit()">添加</button> <button class="input_area_button" id="sectary_submit_button" onclick="deleteContactRelationSubmit()">删除</button>';
        $('#input_area').empty();
        $('#input_area').append(html);
        currentShow = "addOrDeleteContactRelationTable"
    } else {
        currentShow = "";
        $('#input_area').empty();
    }
}

function addContactRelationSubmit(){
    var thirdPartContact = {};
    thirdPartContact.project_participant_id = $('#3rd_part_id_input').val();
    thirdPartContact.project_participant_worker_id = $('#3rd_part_worker_id_input').val();
    $.ajax({
        url: '/api/addOrDeleteContactRelation',
        type: 'GET',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: thirdPartContact,
        error: function(data) {
            alertify.error('添加失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('添加成功');
        }
    });
}

function deleteContactRelationSubmit(){
    var thirdPartContact = {};
    thirdPartContact.project_participant_id = $('#3rd_part_id_input').val();
    thirdPartContact.project_participant_worker_id = $('#3rd_part_worker_id_input').val();
    $.ajax({
        url: '/api/addOrDeleteContactRelation?project_participant_id=' + thirdPartContact.project_participant_id + '&project_participant_worker_id=' + thirdPartContact.project_participant_worker_id,
        type: 'DELETE',
        dataType: 'json',
        timeout: 1000,
        cache: false,
        data: thirdPartContact,
        error: function(data) {
            alertify.error('删除失败:' + data.responseJSON.msg);
        },
        success: function(data) {
            alertify.success('删除成功');
        }
    });
}