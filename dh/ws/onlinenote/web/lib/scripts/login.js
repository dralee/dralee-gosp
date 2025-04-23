(function($){
    const CACHE_DAYS = 7;
    $("#login").click(function () {
        var username = $("#userName").val();
        var password = $("#password").val();
        $.post("/login", {username: username, password: password}, function (data) {
            if (data.code == 0) {
                var result = data.data;
                console.log(result);
                //Cookies.set('token', data.token, { expires: CACHE_DAYS, path: '/' });
                window.location.href = "/";
            }
        })
    });
})(jQuery);