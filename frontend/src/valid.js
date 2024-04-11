/*
 * 校验是否为空(null/空串)
 */
export const checkNull = function (str) {
    return !(str == null || str === "");
};

export const checkStrLen = function (str, len) {
    if (str === undefined || typeof str !== 'string' || str === '') {
        console.log("ss")
        return false
    }
    if (len <= 0) {
        console.log("ee")
        return false
    }
    if (str.length > len) {
        console.log("ff")
        return false
    }
    return true
}

/*
 * 校验是否为纯数字
 * js的isNaN函数
 */
export const checkNum = function (num) {
    return typeof num === 'number' && !isNaN(num)
};

/*
 * 校验是否为纯数字(正则)
 */
export const checkNum2 = function (num) {
    const re = /^[0-9]+.?[0-9]*$/; //判断字符串是否为数字 （判断正整数 /^[1-9]+[0-9]*]*$/）
    return re.test(num);

};

/*
 * 检验手机号
 */
export const checkPhone = function (phone) {
    const reg = /^1[3|4|5|7|8][0-9]{9}$/; //验证规则,第一位是【1】开头，第二位有【3,4,5,7,8】，第三位及以后可以是【0-9】
    //  var reg = /^1[0-9]{10}$/;//不验证第二位，防止几年后新增号码段
    if (!reg.test(phone)) {
        return false;
    }
    return true;
};

/*
 * 验证座机号
 */
export const checkTel = function (tel) {
    const reg = /^(\d3,4|\d{3,4}-)?\d{7,8}$/;
    if (!reg.test(tel)) {
        return false;
    }
    return true;
};

/*
 * 校验ip
 */
export const checkIp = function (ip) {
    const reSpaceCheck = /^(\d+)\.(\d+)\.(\d+)\.(\d+)$/;
    if (reSpaceCheck.test(ip)) {
        ip.match(reSpaceCheck);
        return RegExp.$1 <= 255 && RegExp.$1 >= 0 && RegExp.$2 <= 255 && RegExp.$2 >= 0
            && RegExp.$3 <= 255 && RegExp.$3 >= 0 && RegExp.$4 <= 255 && RegExp.$4 >= 0;
    } else {
        return false;
    }
};

/*
 * 检验url地址
 */
export const checkUrl = function (url) {
    const reg = /(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?/;
    if (!reg.test(url)) {
        return false;
    }
    return true;
};

/**
 * 完美判断是否为网址
 */
export function isURL(strUrl) {
    const regular = /^\b(((https?|ftp):\/\/)?[-a-z0-9]+(\.[-a-z0-9]+)*\.(?:com|edu|gov|int|mil|net|org|biz|info|name|museum|asia|coop|aero|[a-z][a-z]|((25[0-5])|(2[0-4]\d)|(1\d\d)|([1-9]\d)|\d))\b(\/[-a-z0-9_:\@&?=+,.!\/~%\$]*)?)$/i;
    if (regular.test(strUrl)) {
        return true;
    } else {
        return false;
    }
}

/*
 * 检验邮箱
 */
export const checkEmail = function (emailStr) {
    const reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/;
    if (!reg.test(emailStr)) {
        return false;
    }
    return true;
};

/*
 * 检验否是汉字
 */
export const checkCharacter = function (charValue) {
    const reg = /^[\u4e00-\u9fa5]{0,}$/;
    if (!reg.test(charValue)) {
        return false;
    }
    return true;
};