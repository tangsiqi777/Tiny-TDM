/**
 * v2将返回解析成对象形式
 * @param jsonObject
 * @returns {*[]}
 */
export function restDataToJsonObj(jsonObject) {
    try {
        // 确保Result字段存在且不为空
        if (jsonObject.Result && jsonObject.Result.trim() !== '') {
            // 第二次解析Result字段中的JSON字符串
            let innerJsonObject = JSON.parse(jsonObject.Result);
            // 确保data和column_meta字段存在且不为空
            if (innerJsonObject.data && Array.isArray(innerJsonObject.data) &&
                innerJsonObject.head && Array.isArray(innerJsonObject.head) &&
                innerJsonObject.head.length > 0) {

                // 获取第一列的列名
                let columnMeta = [];
                for (let i = 0; i < innerJsonObject.head.length; i++) {
                    columnMeta.push(innerJsonObject.head[i])
                }
                console.log("columnMeta" + JSON.stringify(columnMeta))
                let result = []
                for (let i = 0; i < innerJsonObject.data.length; i++) {
                    let rowObj = {}
                    let rowData = innerJsonObject.data[i]
                    console.log("rowData" + JSON.stringify(rowData))
                    for (let j = 0; j < columnMeta.length; j++) {
                        rowObj[columnMeta[j]] = rowData[j]
                    }
                    result.push(rowObj)
                }
                return result
            }

        }

        // 如果没有找到数据或数据无效，返回一个空数组
        return [];
    } catch (error) {
        console.error('Error parsing JSON:', error);
        // 解析错误时返回空数组
        return [];
    }
}


/**
 * v2从返回从获取列名信息
 * @param jsonObject
 * @returns {*|*[]}
 */
export function getHead(jsonObject) {
    try {
        // 确保Result字段存在且不为空
        if (jsonObject.Result && jsonObject.Result.trim() !== '') {
            // 第二次解析Result字段中的JSON字符串
            let innerJsonObject = JSON.parse(jsonObject.Result);

            // 确保data字段存在且不为空，并构建结果数组
            if (innerJsonObject.head && Array.isArray(innerJsonObject.head) && innerJsonObject.head.length > 0) {
                return innerJsonObject.head.map(item => ({
                    title: item,
                    dataIndex: item
                }));
            }
        }

        // 如果没有找到数据或数据无效，返回一个空数组
        return [];
    } catch (error) {
        console.error('Error parsing JSON:', error);
        // 解析错误时返回空数组
        return [];
    }
}

/**
 * v2判断返回是否有错误
 * @param jsonObjRet
 * @returns {*|SVGAttributes|string}
 */
export const hasError = function (jsonObjRet) {
    // console.log("check:" + JSON.stringify(jsonObjRet))
    if (jsonObjRet === undefined || jsonObjRet === null) {
        return "返回为空"
    }
    let connectionErrorMsg = jsonObjRet.Msg
    if (connectionErrorMsg !== undefined && connectionErrorMsg !== null && connectionErrorMsg !== "") {
        return connectionErrorMsg
    }
    const SUCCESS_CODE = "succ"
    if (jsonObjRet.Result && jsonObjRet.Result.trim() !== '') {
        let queryData = JSON.parse(jsonObjRet.Result);
        if (queryData && queryData.status === SUCCESS_CODE) {
            return ""
        }
        return queryData.desc
    }
    return "返回为空"
};