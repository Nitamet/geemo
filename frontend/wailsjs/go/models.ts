export namespace lcu {
	
	export class RunePage {
	    name: string;
	    primaryStyleId: number;
	    subStyleId: number;
	    selectedPerkIds: number[];
	    current: boolean;
	
	    static createFrom(source: any = {}) {
	        return new RunePage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.primaryStyleId = source["primaryStyleId"];
	        this.subStyleId = source["subStyleId"];
	        this.selectedPerkIds = source["selectedPerkIds"];
	        this.current = source["current"];
	    }
	}
	export class Summoner {
	    displayName: string;
	    profileIconId: number;
	
	    static createFrom(source: any = {}) {
	        return new Summoner(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.displayName = source["displayName"];
	        this.profileIconId = source["profileIconId"];
	    }
	}

}

export namespace lolbuild {
	
	export class Item {
	    id: number;
	    name: string;
	    slug: string;
	    iconUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.slug = source["slug"];
	        this.iconUrl = source["iconUrl"];
	    }
	}
	export class Items {
	    starting: Item[];
	    core: Item[];
	    mythic: Item;
	    fourth: Item[];
	    fifth: Item[];
	    sixth: Item[];
	
	    static createFrom(source: any = {}) {
	        return new Items(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.starting = this.convertValues(source["starting"], Item);
	        this.core = this.convertValues(source["core"], Item);
	        this.mythic = this.convertValues(source["mythic"], Item);
	        this.fourth = this.convertValues(source["fourth"], Item);
	        this.fifth = this.convertValues(source["fifth"], Item);
	        this.sixth = this.convertValues(source["sixth"], Item);
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
	export class Rune {
	    id: number;
	    name: string;
	    slug: string;
	    iconUrl: string;
	    path?: Rune;
	
	    static createFrom(source: any = {}) {
	        return new Rune(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.slug = source["slug"];
	        this.iconUrl = source["iconUrl"];
	        this.path = this.convertValues(source["path"], Rune);
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
	export class Build {
	    name: string;
	    winrate: string;
	    matches: string;
	    primary: Rune;
	    secondary: Rune;
	    selectedPerks: Rune[];
	    items: Items;
	
	    static createFrom(source: any = {}) {
	        return new Build(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.winrate = source["winrate"];
	        this.matches = source["matches"];
	        this.primary = this.convertValues(source["primary"], Rune);
	        this.secondary = this.convertValues(source["secondary"], Rune);
	        this.selectedPerks = this.convertValues(source["selectedPerks"], Rune);
	        this.items = this.convertValues(source["items"], Items);
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
	export class BuildCollection {
	    runes: Build[];
	    source: string;
	
	    static createFrom(source: any = {}) {
	        return new BuildCollection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.runes = this.convertValues(source["runes"], Build);
	        this.source = source["source"];
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
	
	
	
	export class RuneTree {
	    name: string;
	    keystones: Rune[];
	    perks: Rune[];
	    iconUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new RuneTree(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.keystones = this.convertValues(source["keystones"], Rune);
	        this.perks = this.convertValues(source["perks"], Rune);
	        this.iconUrl = source["iconUrl"];
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

