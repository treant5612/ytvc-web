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
    fillVideoInfo(result.data.info)
    listCaptions(result.data.captions)
    listFiles(result.data.files, result.data.info.id)
    showDetails()

}

function fetchErr(xhr, status, error) {
    layer.msg("视频信息获取失败", {icon: 2})
}

function showDetails() {
    $("#details").show()
}

function fillVideoInfo(info) {
    $("#videoTitle").text(cutString(info.title));
    // $("#videoDescription").html(info.description);
    $("#videoUploader").text(info.uploader);
    $("#videoThumbnail").attr("src", info.ThumbnailUrl)


    $("#videoDuration").text(simpleDuration(info.duration / 10e5));
}

function listFiles(files, vid) {
    $("#fileList .file-row").remove()
    if (files.length == 0) {
        $("#fileList").append(`<div class="layui-row file-row"><b>暂无该视频相关文件信息</b></div>`)
    }
    files.forEach(function (f) {
        $("#fileList").append(createFileRow(f, vid))
    })
}

function createFileRow(f, vid) {
    let v = f.videoEncoding != "" ? f.videoEncoding : "-"
    let a = f.audioEncoding != "" ? f.audioEncoding : "-"
    let resolution = f.resolution != "" ? f.resolution : "-"
    let encoding = v + "&nbsp;/&nbsp;" + a
    let size = prettySize(f.size)
    let title = $("#videoTitle").text()
    let temp = `<div class="layui-row file-row">
        <div class="layui-col-md2">${f.extension}</div>
        <div class="layui-col-md2">${resolution}</div>
        <div class="layui-col-md2">${encoding}</div>
        <div class="layui-col-md2">${size}</div>
        <div class="layui-col-md2"><a href="${f.url}" target="_blank" style="color:deepskyblue">链接</a></div>
        <div class="layui-col-md2"><a href="./video/${vid}?no=${f.number}" download="${title}.${f.extension}" >下载</a></div>
     </div>
     <hr class="file-row">
`
    return temp
}

function listCaptions(captions) {
    $("#captionList .caption-row").remove()
    if (captions.length == 0) {
        $("#captionList").append(`<div class="layui-row caption-row"><b>暂无该视频相关字幕信息</b></div>`)
    }
    captions.forEach(function (c, i, arr) {
        $("#captionList").append(createCaptionRow(c, i, arr))
    })
}

function createCaptionRow(c, i, arr) {
    let lang = c.snippet.language
    let trackKind = cTrackKind(c.snippet.trackKind)

    let lastUpdated = new Date(c.snippet.lastUpdated).format("yyyy-MM-dd")
    let captionId = c.id
    let secondaryOptions = createSecondaryOptions(c, i, arr)
    let temp = `<div class="layui-row caption-row">
        <input id="caption_${i}" type="hidden" value="${captionId}">
        <div class="layui-col-md2">${lang}</div>
        <div class="layui-col-md2">${trackKind}</div>
        <div class="layui-col-md2">${lastUpdated}</div>
        <div class="layui-col-md2">
            <select id="cap_select_${i}">${secondaryOptions}</select>
        </div>
        <div class="layui-col-md2">    
            <input id="trans_${i}" type="checkbox" >机翻中文
        </div>
        <div class="layui-col-md2">
            <a href="javascript:downloadCaption('${i}')">下载</a>
        </div>
     </div>
     <hr class="caption-row">
`
    return temp
}

function createSecondaryOptions(c, i, arr) {
    let tmp = `<option value="">无</option>
                <option value="${c.id}">机翻(中文)</option>`
    for (j = 0; j < arr.length; j++) {
        if (j == i) {
            continue
        }
        tmp += `<option value="${arr[j].id}">${arr[j].snippet.language}</option>`
    }
    return tmp
}


function downloadCaption(i) {
    let captionId = $(`#caption_${i}`).val()
    let title = encodeURIComponent($("#videoTitle").text())
    let url = `./caption/${captionId}?fname=${title}`
    if ($(`#trans_${i}`).is(":checked")) {
        url += "&tlang=zh"
    }
    let secondary = $(`#cap_select_${i}`).val()
    if (secondary != "") {
        url += `&secondary=${secondary}`
        if (secondary == captionId) {
            url += "&secondary_tlang=zh"
        }
    }
    fDownload(url)
}


function cTrackKind(k) {
    switch (k) {
        case "ASR":
            return "自动生成"
        case "standard":
            return "标准"
    }
    return k
}

function cutString(s) {
    if (s.length < 64) {
        return s
    }
    return s.substr(0, 64) + "..."
}