export namespace backend {
	
	export class Save {
	    // Go type: NullInt64
	    id: any;
	    managerName: string;
	    gameVersion: number;
	    saveName: string;
	    // Go type: NullString
	    saveImage: any;
	
	    static createFrom(source: any = {}) {
	        return new Save(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = this.convertValues(source["id"], null);
	        this.managerName = source["managerName"];
	        this.gameVersion = source["gameVersion"];
	        this.saveName = source["saveName"];
	        this.saveImage = this.convertValues(source["saveImage"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

