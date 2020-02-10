const $ = layui.$
$(function () {
    $("#fetch_btn").on("click", function () {
        let loading = layer.load(2, {shade: 0.3})
        $.ajax({
            url: "./video",
            type: "GET",
            data: {
                "url": $("#target_url").val()
            },
            complete: function () {
                layer.close(loading)
            },
            success: fetchSucc,
            error: fetchErr,

        })
    });
});

function fetchSucc(result, status, xhr) {
    console.log(result)
    fillVideoInfo(result.data.info)
    listFiles(result.data.files)
    showDetails()
}

function showDetails() {
    $("#details").show()
}

function fillVideoInfo(info) {
    $("#videoTitle").html(info.title);
    // $("#videoDescription").html(info.description);
    $("#videoUploader").html(info.uploader);
    $("#videoDuration").html(simpleDuration(info.duration / 10e5));

}

function listFiles(files) {
    $("#fileList .file-row").remove()
    files.forEach(function (f) {
        $("#fileList").append(createFileRow(f))
    })
}

function createFileRow(f) {
    let v = f.videoEncoding != "" ? f.videoEncoding : "-"
    let a = f.audioEncoding != "" ? f.audioEncoding : "-"
    let resolution = f.resolution != "" ? f.resolution : "-"
    let encoding = v + "&nbsp;/&nbsp;" + a
    let size = prettySize(f.size)
    let temp = `<div class="layui-row file-row">
        <div class="layui-col-md2">${f.extension}</div>
        <div class="layui-col-md2">${resolution}</div>
        <div class="layui-col-md2">${encoding}</div>
        <div class="layui-col-md2">${size}</div>
        <div class="layui-col-md2"><a href="${f.url}" style="color:deepskyblue">链接</a></div>
        <div class="layui-col-md2">下载</div>
     </div>
     <hr class="file-row">
`
    return temp
}

function fetchErr(xhr, status, error) {
    layer.msg("视频信息获取失败", {icon: 2})
}