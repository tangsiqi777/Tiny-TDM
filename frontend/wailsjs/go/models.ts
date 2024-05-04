export namespace service {
	
	export class ConnectionConfig {
	    id: number;
	    name: string;
	    addr?: string;
	    port?: number;
	    username?: string;
	    password?: string;
	
	    static createFrom(source: any = {}) {
	        return new ConnectionConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.addr = source["addr"];
	        this.port = source["port"];
	        this.username = source["username"];
	        this.password = source["password"];
	    }
	}
	export class Query {
	    timeOrder: number;
	    timeStart: string;
	    timeEnd: string;
	    primaryId: string;
	    size: number;
	    current: number;
	
	    static createFrom(source: any = {}) {
	        return new Query(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.timeOrder = source["timeOrder"];
	        this.timeStart = source["timeStart"];
	        this.timeEnd = source["timeEnd"];
	        this.primaryId = source["primaryId"];
	        this.size = source["size"];
	        this.current = source["current"];
	    }
	}

}

