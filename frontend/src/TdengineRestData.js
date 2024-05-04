export function parseNestedJsonAndGetData(jsonObject) {
    try {
        // 确保Result字段存在且不为空
        if (jsonObject.Result && jsonObject.Result.trim() !== '') {
            // 第二次解析Result字段中的JSON字符串
            let innerJsonObject = JSON.parse(jsonObject.Result);
            // 确保data和column_meta字段存在且不为空
            if (innerJsonObject.data && Array.isArray(innerJsonObject.data) &&
                innerJsonObject.column_meta && Array.isArray(innerJsonObject.column_meta[0]) &&
                innerJsonObject.column_meta[0].length > 0) {

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
                    console.log("rowData" + JSON.stringify(rowData))
                    for (let j = 0; j < columnMeta.length; j++) {
                        rowObj[toCamelCase(columnMeta[j])] = rowData[j]
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

export function toCamelCase(str) {
    // 使用正则表达式将下划线及其后面的字符转换为大写，并且去掉下划线
    return str.replace(/_(\w)/g, function (match, group1) {
        return group1.toUpperCase();
    });
}

export function getHead(jsonObject) {
    try {
        // 确保Result字段存在且不为空
        if (jsonObject.Result && jsonObject.Result.trim() !== '') {
            // 第二次解析Result字段中的JSON字符串
            let innerJsonObject = JSON.parse(jsonObject.Result);

            // 确保data字段存在且不为空，并构建结果数组
            if (innerJsonObject.column_meta && Array.isArray(innerJsonObject.column_meta) && innerJsonObject.column_meta.length > 0) {
                return innerJsonObject.column_meta.map(row => ({
                    title: row[0],
                    dataIndex: row[0]
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