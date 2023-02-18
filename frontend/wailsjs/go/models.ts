export namespace backend {
	
	export class Save {
	    id: number;
	    managerName: string;
	    gameVersion: number;
	    saveName: string;
	    saveImage?: string;
	
	    static createFrom(source: any = {}) {
	        return new Save(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.managerName = source["managerName"];
	        this.gameVersion = source["gameVersion"];
	        this.saveName = source["saveName"];
	        this.saveImage = source["saveImage"];
	    }
	}

}

