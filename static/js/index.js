console.log("hello world")


$(function () {
    $("#Get").click(function () {
        console.log("GET获取数据");
        $.getJSON("/getInfo", function (data) {
            console.log(data)
        });
    });

    $("#Post").click(function () {
        console.log("POST获取数据")
        $.post("/postInfo", {dn: "abc123", name: "张珊"}, function (data) {
            console.log(data)
        });
    });
});