// 不使用vuex
// 使用传统的localStorage维持存储

// 使用cookie作为第二份存储
export function set_code_theme(theme) {
    localStorage.setItem("code_theme", theme)
}

// 传入参数 default theme
export function get_code_theme(theme) {
    let code_theme;
    if (localStorage.getItem("code_theme")) {
        code_theme = localStorage.getItem("code_theme");
        return code_theme
    }else {
        if (theme !== ""){
            return theme
        }else {
            return "github"
        }
    }
}

export function get_admin_cookie() {
    let cookie = document.cookie;
    let arrcookie = cookie.split(";");
    for (let i=0;i<arrcookie.length;i++) {
        let a = arrcookie[i].split("=");
        if (a[0] === "admin_token") {
            return a[1]
        }
    }
    return ""
}

export function set_admin_cookie(cook, timeout) {
    let d = new Date();
    d.setTime(d.getTime()+timeout);
    document.cookie = cook + ";" + "expires=" + d.toGMTString();
}

export function delete_admin_cookie() {
    document.cookie = "";
}