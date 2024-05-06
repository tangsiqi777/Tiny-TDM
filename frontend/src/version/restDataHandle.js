import {getHead as getHeadV3, hasError as hasErrorV3, restDataToJsonObj as restDataToJsonObjV3} from "./rest_v3.js";
import {getHead as getHeadV2, hasError as hasErrorV2, restDataToJsonObj as restDataToJsonObjV2} from "./rest_v2.js";
import {Store} from "../util/store.js";
import pinia from '../util/myPinia.js'

const store = Store(pinia);

export const handleMap = new Map

handleMap.set("v3", {
    restDataToJsonObj: restDataToJsonObjV3,
    getHead: getHeadV3,
    hasError: hasErrorV3,
})

handleMap.set("v2", {
    restDataToJsonObj: restDataToJsonObjV2,
    getHead: getHeadV2,
    hasError: hasErrorV2,
})


const DEFAULT_VERSION = "v3"

export function restDataToJsonObj(jsonObject) {
    const version = store.conn.conn.Version
    return (handleMap.get(version) || handleMap.get(DEFAULT_VERSION)).restDataToJsonObj(jsonObject)
}


export function getHead(jsonObject) {
    const version = store.conn.conn.Version
    return (handleMap.get(version) || handleMap.get(DEFAULT_VERSION)).getHead(jsonObject)
}

export function hasError(jsonObject) {
    const version = store.conn.conn.Version
    const fun = handleMap.get(version) || handleMap.get(DEFAULT_VERSION);
    console.log("version:" + fun)
    return fun.hasError(jsonObject)
}
