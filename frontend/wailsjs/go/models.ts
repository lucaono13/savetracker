export namespace backend {
	
	export class Match {
	    seasonID: number;
	    // Go type: NullInt64
	    saveID: any;
	    // Go type: NullString
	    year: any;
	    date: string;
	    opposition: string;
	    venue: string;
	    stadium: string;
	    competition: string;
	    goalsAgainst: number;
	    goalsFor: number;
	    result: string;
	    extraTime: number;
	    penalties: number;
	
	    static createFrom(source: any = {}) {
	        return new Match(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.seasonID = source["seasonID"];
	        this.saveID = this.convertValues(source["saveID"], null);
	        this.year = this.convertValues(source["year"], null);
	        this.date = source["date"];
	        this.opposition = source["opposition"];
	        this.venue = source["venue"];
	        this.stadium = source["stadium"];
	        this.competition = source["competition"];
	        this.goalsAgainst = source["goalsAgainst"];
	        this.goalsFor = source["goalsFor"];
	        this.result = source["result"];
	        this.extraTime = source["extraTime"];
	        this.penalties = source["penalties"];
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
	export class Save {
	    // Go type: NullInt64
	    id: any;
	    managerName: string;
	    gameVersion: number;
	    saveName: string;
	    // Go type: NullString
	    saveImage: any;
	    currency: string;
	
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
	        this.currency = source["currency"];
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

export namespace main {
	
	export class ErrorReturn {
	    Error: string;
	    String: string;
	    Integer: number;
	    Save: backend.Save;
	    SaveList: backend.Save[];
	    Matches: backend.Match[];
	
	    static createFrom(source: any = {}) {
	        return new ErrorReturn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Error = source["Error"];
	        this.String = source["String"];
	        this.Integer = source["Integer"];
	        this.Save = this.convertValues(source["Save"], backend.Save);
	        this.SaveList = this.convertValues(source["SaveList"], backend.Save);
	        this.Matches = this.convertValues(source["Matches"], backend.Match);
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
	export class NewSeason {
	    teamName: string;
	    shortName: string;
	    season: string;
	    country: string;
	    trophiesWon: string[];
	    squadFile: string;
	    scheduleFile: string;
	    transfersFile: string;
	
	    static createFrom(source: any = {}) {
	        return new NewSeason(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.teamName = source["teamName"];
	        this.shortName = source["shortName"];
	        this.season = source["season"];
	        this.country = source["country"];
	        this.trophiesWon = source["trophiesWon"];
	        this.squadFile = source["squadFile"];
	        this.scheduleFile = source["scheduleFile"];
	        this.transfersFile = source["transfersFile"];
	    }
	}

}

