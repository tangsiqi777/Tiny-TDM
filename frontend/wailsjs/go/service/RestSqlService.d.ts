// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {service} from '../models';
import {context} from '../models';

export function Check(arg1:string):Promise<any>;

export function CountChildTable(arg1:service.ConnectionConfig,arg2:string,arg3:string,arg4:string):Promise<service.RestSqlResult>;

export function CountData(arg1:service.ConnectionConfig,arg2:string,arg3:string,arg4:service.Query):Promise<service.RestSqlResult>;

export function DescTable(arg1:service.ConnectionConfig,arg2:string,arg3:string):Promise<service.RestSqlResult>;

export function DoPost(arg1:service.ConnectionConfig,arg2:string):Promise<service.RestSqlResult>;

export function ListChildTable(arg1:service.ConnectionConfig,arg2:string,arg3:string,arg4:string,arg5:number,arg6:number):Promise<service.RestSqlResult>;

export function ListDatabases(arg1:service.ConnectionConfig):Promise<service.RestSqlResult>;

export function ListSuperTable(arg1:service.ConnectionConfig,arg2:string,arg3:string):Promise<service.RestSqlResult>;

export function PageData(arg1:service.ConnectionConfig,arg2:string,arg3:string,arg4:service.Query):Promise<service.RestSqlResult>;

export function SqlQuery(arg1:service.ConnectionConfig,arg2:string):Promise<service.RestSqlResult>;

export function Startup(arg1:context.Context):Promise<void>;
