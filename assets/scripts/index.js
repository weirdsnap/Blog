// window.onload = function() {
//     document.getElementById("login_button").onclick = function() {
//         //window.open("www.weirdsnap.top:8888/login")
//         window.location.href="/login"
//     }
// }
var name = new Vue({
    el: '#name',
    data: {
        massage : "Snap's Blog"
    },
    methods: {
        jumpList: function () {
            window.location.href='./htmls/list.html'
        }
    }
})

var header = new Vue({
    el: '#header',
    methods: {
        home : function() {
            window.location.href='https://weirdsnap.github.io'
        },
        github : function() {
            window.location.href='https://github.com/weirdsnap'
        }
    }
})