function login() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    var data = "username=" + username + "&hashPassword=" + generateHash(password);
    window.location.href = "/api/login?" + data;
}

function register() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    var data = "username=" + username + "&hashPassword=" + generateHash(password);
    window.location.href = "/api/register?" + data;
}


function toRegister(){
    window.location.href = "/login/register.html";
}

function toLogin(){
    window.location.href = "/login/login.html";
}

function generateHash(password) {
    if (password.length == 0) {
        return "";
    }
    password =  password + "saltysalt";
    var hash = 0;
    if (password.length == 0) return hash;
    for (i = 0; i < password.length; i++) {
        char = password.charCodeAt(i);
        hash = ((hash << 5) - hash) + char;
        hash = hash & hash; // Convert to 32bit integer
    }
    return hash;
}