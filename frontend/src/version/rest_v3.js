/**
 * v3将返回解析成对象形式
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
            if (innerJsonObject.data && Array.isArray(innerJsonObject.data) && innerJsonObject.column_meta && Array.isArray(innerJsonObject.column_meta[0]) && innerJsonObject.column_meta[0].length > 0) {

                // 获取第一列的列名
                let columnMeta = [];
                for (let i = 0; i < innerJsonObject.column_meta.length; i++) {
                    columnMeta.push(innerJsonObject.column_meta[i][0])
                }
                console.log("columnMeta" + JSON.stringify(columnMeta))
                let result = []
                for (let i = 0; i < innerJsonObject.data.length; i++) {
                    let rowObj = {}
                    let rowData = innerJsonObject.data[i]
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
 * v3从返回从获取列名信息
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
            if (innerJsonObject.column_meta && Array.isArray(innerJsonObject.column_meta) && innerJsonObject.column_meta.length > 0) {
                return innerJsonObject.column_meta.map(row => ({
                    title: row[0], dataIndex: row[0]
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
 * v3判断返回是否有错误
 * @param jsonObjRet
 * @returns {*|SVGAttributes|string}
 */
export const hasError = function (jsonObjRet) {
    if (jsonObjRet === undefined || jsonObjRet === null) {
        return "返回为空"
    }
    let connectionErrorMsg = jsonObjRet.Msg
    if (connectionErrorMsg !== undefined && connectionErrorMsg !== null && connectionErrorMsg !== "") {
        return connectionErrorMsg
    }
    if (jsonObjRet.Result && jsonObjRet.Result.trim() !== '') {
        let queryData = JSON.parse(jsonObjRet.Result);
        if (queryData && queryData.code === 0) {
            return ""
        }
        return queryData.desc
    }
    return "返回为空"
};