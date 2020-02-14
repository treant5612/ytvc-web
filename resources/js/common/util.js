/*
 fileSize --human
 */
function prettySize(bytes, separator = '', postFix = '') {
    if (bytes) {
        const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
        const i = Math.min(parseInt(Math.floor(Math.log(bytes) / Math.log(1024)), 10), sizes.length - 1);
        return `${(bytes / (1024 ** i)).toFixed(i ? 1 : 0)}${separator}${sizes[i]}${postFix}`;
    }
    return 'n/a';
}

/*
    将time.Duration 格式化为人类友好的格式
 */
function simpleDuration(duration, type) {
    if (type === 's') {
        duration = duration * 1000
    }
    let days = '', hours = '', minutes = '', seconds = ''
    let day = 24 * 60 * 60 * 1000,
        hour = 60 * 60 * 1000,
        minute = 60 * 1000,
        second = 1000
    if (duration >= day) {
        days = Math.floor(duration / day) + '天'
        hours = Math.floor(duration % day / hour) + '小时'
    } else if (duration >= hour && duration < day) {
        hours = Math.floor(duration / hour) + '小时'
        minutes = Math.floor(duration % hour / minute) + '分钟'
    } else if (duration > minute && duration < hour) {
        minutes = Math.floor(duration / minute) + '分钟'
        seconds = Math.floor(duration % minute / second) + '秒'
    } else if (duration < minute) {
        seconds = Math.floor(duration / second) + '秒'
    }
    return days + hours + minutes + seconds
}

/*
    为date原型增加format方法
 */
Date.prototype.format = function (fmt) {
    let o = {
        "M+": this.getMonth() + 1,                 //月份
        "d+": this.getDate(),                    //日
        "h+": this.getHours(),                   //小时
        "m+": this.getMinutes(),                 //分
        "s+": this.getSeconds(),                 //秒
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度
        "S": this.getMilliseconds()             //毫秒
    };
    if (/(y+)/.test(fmt)) {
        fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    }
    for (let k in o) {
        if (new RegExp("(" + k + ")").test(fmt)) {
            fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
        }
    }
    return fmt;
}


/*
下载文件
 */
function fDownload(url, filename) {
    var eleLink = document.createElement('a');
    eleLink.download = filename;
    eleLink.style.display = 'none';
    eleLink.href = url;
    document.body.appendChild(eleLink);
    eleLink.click();
    document.body.removeChild(eleLink);
}