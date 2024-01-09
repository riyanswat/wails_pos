export namespace backend {
	
	export class UserData {
	    item: string;
	    quantity: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new UserData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.item = source["item"];
	        this.quantity = source["quantity"];
	        this.password = source["password"];
	    }
	}

}

