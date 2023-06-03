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
	export class PlayerSquadView {
	    playerName: string;
	    teamName: string;
	    // Go type: NullString
	    year: any;
	    position: string;
	    starts: number;
	    subs: number;
	    goals: number;
	    assists: number;
	    playerOfTheMatch: number;
	    passPerc: number;
	    yellowCards: number;
	    redCards: number;
	    avgRating: number;
	    winPerc: number;
	    minutes: number;
	    shutouts: number;
	    savePerc: number;
	
	    static createFrom(source: any = {}) {
	        return new PlayerSquadView(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.playerName = source["playerName"];
	        this.teamName = source["teamName"];
	        this.year = this.convertValues(source["year"], null);
	        this.position = source["position"];
	        this.starts = source["starts"];
	        this.subs = source["subs"];
	        this.goals = source["goals"];
	        this.assists = source["assists"];
	        this.playerOfTheMatch = source["playerOfTheMatch"];
	        this.passPerc = source["passPerc"];
	        this.yellowCards = source["yellowCards"];
	        this.redCards = source["redCards"];
	        this.avgRating = source["avgRating"];
	        this.winPerc = source["winPerc"];
	        this.minutes = source["minutes"];
	        this.shutouts = source["shutouts"];
	        this.savePerc = source["savePerc"];
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
	export class Transfer {
	    transferID: number;
	    seasonID: number;
	    date: string;
	    player: string;
	    team: string;
	    fee: number;
	    // Go type: sql
	    potentialFee: any;
	    inTransfer: boolean;
	    loan: number;
	    free: number;
	    // Go type: NullString
	    year: any;
	
	    static createFrom(source: any = {}) {
	        return new Transfer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.transferID = source["transferID"];
	        this.seasonID = source["seasonID"];
	        this.date = source["date"];
	        this.player = source["player"];
	        this.team = source["team"];
	        this.fee = source["fee"];
	        this.potentialFee = this.convertValues(source["potentialFee"], null);
	        this.inTransfer = source["inTransfer"];
	        this.loan = source["loan"];
	        this.free = source["free"];
	        this.year = this.convertValues(source["year"], null);
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
	    InTransfers: backend.Transfer[];
	    OutTransfers: backend.Transfer[];
	    Currency: string;
	    Outfielders: backend.PlayerSquadView[];
	    Goalies: backend.PlayerSquadView[];
	
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
	        this.InTransfers = this.convertValues(source["InTransfers"], backend.Transfer);
	        this.OutTransfers = this.convertValues(source["OutTransfers"], backend.Transfer);
	        this.Currency = source["Currency"];
	        this.Outfielders = this.convertValues(source["Outfielders"], backend.PlayerSquadView);
	        this.Goalies = this.convertValues(source["Goalies"], backend.PlayerSquadView);
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

