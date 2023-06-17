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
	export class PlayerPageInfo {
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
	    aer: number;
	    cmd: number;
	    com: number;
	    ecc: number;
	    han: number;
	    kic: number;
	    ovo: number;
	    pun: number;
	    ref: number;
	    tro: number;
	    thr: number;
	    cor: number;
	    cro: number;
	    dri: number;
	    fin: number;
	    fir: number;
	    fre: number;
	    hea: number;
	    lon: number;
	    lth: number;
	    mar: number;
	    pas: number;
	    pen: number;
	    tck: number;
	    tec: number;
	    agg: number;
	    ant: number;
	    bra: number;
	    cmp: number;
	    cnt: number;
	    dec: number;
	    det: number;
	    fla: number;
	    ldr: number;
	    otb: number;
	    pos: number;
	    tea: number;
	    vis: number;
	    wor: number;
	    acc: number;
	    agi: number;
	    bal: number;
	    jum: number;
	    nat: number;
	    pac: number;
	    sta: number;
	    str: number;
	    saveName: string;
	    season: string;
	    teamName: string;
	    playerSeasonID: number;
	
	    static createFrom(source: any = {}) {
	        return new PlayerPageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
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
	        this.aer = source["aer"];
	        this.cmd = source["cmd"];
	        this.com = source["com"];
	        this.ecc = source["ecc"];
	        this.han = source["han"];
	        this.kic = source["kic"];
	        this.ovo = source["ovo"];
	        this.pun = source["pun"];
	        this.ref = source["ref"];
	        this.tro = source["tro"];
	        this.thr = source["thr"];
	        this.cor = source["cor"];
	        this.cro = source["cro"];
	        this.dri = source["dri"];
	        this.fin = source["fin"];
	        this.fir = source["fir"];
	        this.fre = source["fre"];
	        this.hea = source["hea"];
	        this.lon = source["lon"];
	        this.lth = source["lth"];
	        this.mar = source["mar"];
	        this.pas = source["pas"];
	        this.pen = source["pen"];
	        this.tck = source["tck"];
	        this.tec = source["tec"];
	        this.agg = source["agg"];
	        this.ant = source["ant"];
	        this.bra = source["bra"];
	        this.cmp = source["cmp"];
	        this.cnt = source["cnt"];
	        this.dec = source["dec"];
	        this.det = source["det"];
	        this.fla = source["fla"];
	        this.ldr = source["ldr"];
	        this.otb = source["otb"];
	        this.pos = source["pos"];
	        this.tea = source["tea"];
	        this.vis = source["vis"];
	        this.wor = source["wor"];
	        this.acc = source["acc"];
	        this.agi = source["agi"];
	        this.bal = source["bal"];
	        this.jum = source["jum"];
	        this.nat = source["nat"];
	        this.pac = source["pac"];
	        this.sta = source["sta"];
	        this.str = source["str"];
	        this.saveName = source["saveName"];
	        this.season = source["season"];
	        this.teamName = source["teamName"];
	        this.playerSeasonID = source["playerSeasonID"];
	    }
	}
	export class PlayerSquadView {
	    playerID: number;
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
	        this.playerID = source["playerID"];
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
	export class PlayerSumsAvgs {
	    playerID: number;
	    saveID: number;
	    playerName: string;
	    uniqueID: number;
	    nationality: string;
	    // Go type: sql
	    secondNationality: any;
	    birthdate: string;
	    position: string;
	    saveName: string;
	    gameVersion: number;
	    seasons: string;
	    teamName: string;
	    avgMin: number;
	    totMin: number;
	    avgStart: number;
	    totStart: number;
	    avgSubs: number;
	    totSubs: number;
	    avgGls: number;
	    totGls: number;
	    avgAst: number;
	    totAst: number;
	    avgYel: number;
	    totYel: number;
	    avgRed: number;
	    totRed: number;
	    avgRat: number;
	    avgPOM: number;
	    totPOM: number;
	    avgPasP: number;
	    avgWinP: number;
	    avgShutouts: number;
	    totShutouts: number;
	    avgSaveP: number;
	
	    static createFrom(source: any = {}) {
	        return new PlayerSumsAvgs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.playerID = source["playerID"];
	        this.saveID = source["saveID"];
	        this.playerName = source["playerName"];
	        this.uniqueID = source["uniqueID"];
	        this.nationality = source["nationality"];
	        this.secondNationality = this.convertValues(source["secondNationality"], null);
	        this.birthdate = source["birthdate"];
	        this.position = source["position"];
	        this.saveName = source["saveName"];
	        this.gameVersion = source["gameVersion"];
	        this.seasons = source["seasons"];
	        this.teamName = source["teamName"];
	        this.avgMin = source["avgMin"];
	        this.totMin = source["totMin"];
	        this.avgStart = source["avgStart"];
	        this.totStart = source["totStart"];
	        this.avgSubs = source["avgSubs"];
	        this.totSubs = source["totSubs"];
	        this.avgGls = source["avgGls"];
	        this.totGls = source["totGls"];
	        this.avgAst = source["avgAst"];
	        this.totAst = source["totAst"];
	        this.avgYel = source["avgYel"];
	        this.totYel = source["totYel"];
	        this.avgRed = source["avgRed"];
	        this.totRed = source["totRed"];
	        this.avgRat = source["avgRat"];
	        this.avgPOM = source["avgPOM"];
	        this.totPOM = source["totPOM"];
	        this.avgPasP = source["avgPasP"];
	        this.avgWinP = source["avgWinP"];
	        this.avgShutouts = source["avgShutouts"];
	        this.totShutouts = source["totShutouts"];
	        this.avgSaveP = source["avgSaveP"];
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
	export class PlayerTotalsView {
	    playerID: number;
	    playerName: string;
	    position: string;
	    teamName: string;
	    seasons: number;
	    minutes: number;
	    starts: number;
	    subs: number;
	    goals: number;
	    assists: number;
	    yellowCards: number;
	    redCards: number;
	    avgRating: number;
	    playerOfTheMatch: number;
	    avgPassP: number;
	    avgWinP: number;
	    shutouts: number;
	    avgSaveP: number;
	
	    static createFrom(source: any = {}) {
	        return new PlayerTotalsView(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.playerID = source["playerID"];
	        this.playerName = source["playerName"];
	        this.position = source["position"];
	        this.teamName = source["teamName"];
	        this.seasons = source["seasons"];
	        this.minutes = source["minutes"];
	        this.starts = source["starts"];
	        this.subs = source["subs"];
	        this.goals = source["goals"];
	        this.assists = source["assists"];
	        this.yellowCards = source["yellowCards"];
	        this.redCards = source["redCards"];
	        this.avgRating = source["avgRating"];
	        this.playerOfTheMatch = source["playerOfTheMatch"];
	        this.avgPassP = source["avgPassP"];
	        this.avgWinP = source["avgWinP"];
	        this.shutouts = source["shutouts"];
	        this.avgSaveP = source["avgSaveP"];
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
	export class Story {
	    saveID: number;
	    story: string;
	
	    static createFrom(source: any = {}) {
	        return new Story(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.saveID = source["saveID"];
	        this.story = source["story"];
	    }
	}
	export class TopResults {
	    playerID: number;
	    playerName: string;
	    goals: number;
	    assists: number;
	    apps: number;
	    avgRating: number;
	
	    static createFrom(source: any = {}) {
	        return new TopResults(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.playerID = source["playerID"];
	        this.playerName = source["playerName"];
	        this.goals = source["goals"];
	        this.assists = source["assists"];
	        this.apps = source["apps"];
	        this.avgRating = source["avgRating"];
	    }
	}
	export class TopTransfers {
	    teamName: string;
	    currency: string;
	    avgFee: number;
	    totFee: number;
	    numTransfers: number;
	
	    static createFrom(source: any = {}) {
	        return new TopTransfers(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.teamName = source["teamName"];
	        this.currency = source["currency"];
	        this.avgFee = source["avgFee"];
	        this.totFee = source["totFee"];
	        this.numTransfers = source["numTransfers"];
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
	export class Trophy {
	    // Go type: NullInt64
	    trophyID: any;
	    trophyName: string;
	    // Go type: NullString
	    trophyImage: any;
	    season: string;
	
	    static createFrom(source: any = {}) {
	        return new Trophy(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.trophyID = this.convertValues(source["trophyID"], null);
	        this.trophyName = source["trophyName"];
	        this.trophyImage = this.convertValues(source["trophyImage"], null);
	        this.season = source["season"];
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
	    OutTotals: backend.PlayerTotalsView[];
	    GKTotals: backend.PlayerTotalsView[];
	    TopGls: backend.TopResults[];
	    TopAsts: backend.TopResults[];
	    TopApps: backend.TopResults[];
	    TopAvg: backend.TopResults[];
	    TopTrfs: backend.TopTransfers[];
	    AvgInFee: number;
	    AvgOutFee: number;
	    Trophies: backend.Trophy[];
	    ImageFile: string;
	    b64Image: string;
	    OnePlayer: backend.PlayerPageInfo[];
	    PlayerAvgSum: backend.PlayerSumsAvgs;
	
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
	        this.OutTotals = this.convertValues(source["OutTotals"], backend.PlayerTotalsView);
	        this.GKTotals = this.convertValues(source["GKTotals"], backend.PlayerTotalsView);
	        this.TopGls = this.convertValues(source["TopGls"], backend.TopResults);
	        this.TopAsts = this.convertValues(source["TopAsts"], backend.TopResults);
	        this.TopApps = this.convertValues(source["TopApps"], backend.TopResults);
	        this.TopAvg = this.convertValues(source["TopAvg"], backend.TopResults);
	        this.TopTrfs = this.convertValues(source["TopTrfs"], backend.TopTransfers);
	        this.AvgInFee = source["AvgInFee"];
	        this.AvgOutFee = source["AvgOutFee"];
	        this.Trophies = this.convertValues(source["Trophies"], backend.Trophy);
	        this.ImageFile = source["ImageFile"];
	        this.b64Image = source["b64Image"];
	        this.OnePlayer = this.convertValues(source["OnePlayer"], backend.PlayerPageInfo);
	        this.PlayerAvgSum = this.convertValues(source["PlayerAvgSum"], backend.PlayerSumsAvgs);
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

