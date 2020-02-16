const $ = layui.$

$(function () {
    $("#comment_btn").on("click", function () {
        nickname = $("#nickname").val()
        comment = $("#comment_content").val()
        $.ajax({
            method: "POST",
            url: "/guestbook",
            data: {
                nickname: nickname,
                comment: comment
            },
            success: commentSucc,
            error: commentErr,
        })
    })
})

function commentSucc() {
    window.location.reload()
}

function commentErr() {
    layer.msg("评论失败", {icon: 2})

}