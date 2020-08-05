var url = {
	baseUrl: "http://localhost:8000/",
}

/**
 * @param {number} time
 * @param {string} option
 * @returns {string}
 */
function formatTime(time) {
    if (('' + time).length === 10) {
        time = parseInt(time) * 1000
    } else {
        time = +time
    }
    const d = new Date(time)
    const now = Date.now()

    const diff = (now - d) / 1000

    if (diff < 30) {
        return '刚刚'
    } else if (diff < 3600) {
        // less 1 hour
        return Math.ceil(diff / 60) + '分钟前'
    } else if (diff < 3600 * 24) {
        return Math.ceil(diff / 3600) + '小时前'
    } else if (diff < 3600 * 24 * 2) {
        return '1天前'
    }
    return (
        d.getFullYear() +
        '年' +
        (d.getMonth() + 1) +
        '月' +
        d.getDate() +
        '日' +
        d.getHours() +
        '时' +
        d.getMinutes() +
        '分'
    )
}

//返回日期
function formatDate(time) {
    if (('' + time).length === 10) {
        time = parseInt(time) * 1000
    } else {
        time = +time
    }
    const d = new Date(time)
    return (
        d.getFullYear() +
        '年' +
        (d.getMonth() + 1) +
        '月' +
        d.getDate() +
        '日'
    )
}

/**
 * @param {string} url
 * @returns {Object}
 */
function param2Obj(url) {
    const search = url.split('?')[1]
    if (!search) {
      return {}
    }
    return JSON.parse(
      '{"' +
        decodeURIComponent(search)
          .replace(/"/g, '\\"')
          .replace(/&/g, '","')
          .replace(/=/g, '":"')
          .replace(/\+/g, ' ') +
        '"}'
    )
  }

  //name为获取的参数名称，比如URL为 xxxx?id=1,要获取id的值，则直接可以 getUrlParam('id')来获取
  function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    var r = window.location.search.substr(1).match(reg);  //匹配目标参数
    if (r != null) return unescape(r[2]); return null; //返回参数值
}

//随机颜色提交
function getColor(){  
    var colorElements = "0,1,2,3,4,5,6,7,8,9,a,b,c,d,e,f";  
    var colorArray = colorElements.split(",");  
    var color ="#";  
    for(var i =0;i<6;i++){  
        color+=colorArray[Math.floor(Math.random()*16)];  
    }  
return color;
}
  