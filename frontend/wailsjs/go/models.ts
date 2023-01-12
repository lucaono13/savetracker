export namespace backend {
	
	export class Save {
	    id: number;
	    managerName: string;
	    gameVersion: number;
	
	    static createFrom(source: any = {}) {
	        return new Save(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.managerName = source["managerName"];
	        this.gameVersion = source["gameVersion"];
	    }
	}

}

