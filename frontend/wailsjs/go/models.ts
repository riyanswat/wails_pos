export namespace backend {
	
	export class UserData {
	    website: string;
	    email: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new UserData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.website = source["website"];
	        this.email = source["email"];
	        this.password = source["password"];
	    }
	}

}

